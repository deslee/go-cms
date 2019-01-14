package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
	"github.com/deslee/cms/models"
	"os"
	"reflect"
	"strings"
)

type tabler interface {
	TableName() string
}

func main() {
	generate(
		&models.User{},
		&models.Site{},
		&models.Item{},
		&models.Group{},
		&models.Asset{},
		&models.SiteUser{},
		&models.ItemGroup{},
		&models.ItemAsset{},
	)
}

func generate(modelTypes ...interface{}) {
	f := NewFile("repository")
	for _, modelType := range modelTypes {
		model := getModel(modelType)
		generateRepositoryMethods(model, f)
	}

	repoFile, err := os.Create("./repository/repository.go")
	if err != nil {
		panic(err)
	}
	defer repoFile.Close()

	err = f.Render(repoFile)
	if err != nil {
		panic(err)
	}
}

func generateRepositoryMethods(model Model, f *File) {
	// generate upsert sql
	upsertSql := `INSERT INTO ` + model.TableName + ` VALUES (`
	for idx, field := range model.Fields {
		upsertSql += fmt.Sprintf(":%s", field.columnName())
		if idx < (len(model.Fields) - 1) {
			upsertSql += ","
		}
	}
	upsertSql += `) ON CONFLICT(Id) DO UPDATE SET `
	for idx, field := range model.Fields {
		upsertSql += fmt.Sprintf("`%s`=excluded.`%s`", field.columnName(), field.columnName())
		if idx < (len(model.Fields) - 1) {
			upsertSql += ","
		}
	}

	f.Func().Id(fmt.Sprintf("Upsert%s", model.StructName)).Params(
		Id("ctx").Qual("context", "Context"), Id("db").Add(Op("*")).Qual("github.com/jmoiron/sqlx", "DB"), Id("obj").Qual("github.com/deslee/cms/models", model.StructName),
	).Error().Block(
		List(Id("stmt"), Id("err")).Op(":=").Id("db.PrepareNamedContext").Call(Id("ctx"), Lit(upsertSql)),
		If(Id("err").Op("!=").Id("nil")).Block(
			Return(Id("err")),
		),
		Line(),
		List(Id("_"), Id("err")).Op("=").Id("stmt.Exec").Call(Id("obj")),

		If(Id("err").Op("!=").Id("nil")).Block(
			Return(Id("err")),
		),
		Line(),
		Return(Id("err")),
	)

	if model.singlePrimaryKeyField() != nil {
		f.Line()

		f.Func().Id(fmt.Sprintf("Find%sBy%s", model.StructName, model.singlePrimaryKeyField().FieldName)).Params(
			Id("ctx").Qual("context", "Context"), Id("db").Add(Op("*")).Qual("github.com/jmoiron/sqlx", "DB"), Id("val").Id("string"),
		).Params(Op("*").Qual("github.com/deslee/cms/models", model.StructName), Error()).Block(
			Id("obj").Op(":=").Qual("github.com/deslee/cms/models", model.StructName).Values(),
			Line(),
			Id("err").Op(":=").Id("db.QueryRowx").Call(
				Lit(fmt.Sprintf("SELECT * FROM %s WHERE %s=?", model.TableName, model.singlePrimaryKeyField().FieldName)), Id("val"),
			).Add(Op(".")).Add(Id("StructScan")).Call(Op("&").Id("obj")),
			Line(),
			If(Id("err").Op("!=").Id("nil")).Block(
				If(Id("err")).Op("==").Qual("database/sql", "ErrNoRows").Block(
					Return(List(Id("nil"), Id("nil"))),
				),
				Return(List(Id("nil"), Id("err"))),
			),
			Line(),
			Return(List(Op("&").Id("obj"), Id("nil"))),
		)
	}

	f.Line()
}

type Model struct {
	StructName string
	TableName  string
	Fields     []Field
}

type Field struct {
	FieldName string
	TagMap    map[string]string
}

func (f Field) columnName() string {
	columnName, ok := f.TagMap["column"]
	if !ok {
		panic(fmt.Sprintf("Field %s has no column defined!", f.FieldName))
	}
	return columnName
}

func (f Field) isPk() bool {
	isPk := f.TagMap["Pk"]
	return len(isPk) > 0
}

func (m Model) singlePrimaryKeyField() *Field {
	var primaryKeys []Field
	for _, field := range m.Fields {
		if field.isPk() {
			primaryKeys = append(primaryKeys, field)
		}
	}
	if len(primaryKeys) > 1 {
		return nil
	}
	return &primaryKeys[0]
}

func getModel(i interface{}) Model {
	columns := recursivelyGetColumns(reflect.TypeOf(i).Elem())
	return Model{
		StructName: reflect.TypeOf(i).Elem().Name(),
		TableName:  (i.(tabler)).TableName(),
		Fields:     columns,
	}
}

func recursivelyGetColumns(t reflect.Type) []Field {
	var columns []Field
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Type.Kind() == reflect.Struct {
			for _, col := range recursivelyGetColumns(field.Type) {
				columns = append(columns, col)
			}
			continue
		}

		dbGen := field.Tag.Get("dbGen")
		if len(dbGen) > 0 {
			column := Field{
				FieldName: field.Name,
				TagMap:    make(map[string]string),
			}
			for _, seg := range strings.Split(dbGen, ";") {
				definition := strings.Split(seg, ":")
				if len(definition) != 2 {
					continue
				}
				column.TagMap[definition[0]] = definition[1]
			}
			columns = append(columns, column)
		}

	}
	return columns
}
