package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
	"github.com/deslee/cms/data"
	"os"
	"reflect"
)

type tabler interface {
	TableName() string
}

func main() {
	generate(
		&data.User{},
		&data.Site{},
		&data.Item{},
		&data.Group{},
		&data.Asset{},
		&data.SiteUser{},
		&data.ItemGroup{},
		&data.ItemAsset{},
	)
}

func generate(modelTypes ... interface{}) {
	f := NewFile("data")
	for _, modelType := range modelTypes {
		model := getModel(modelType)
		generateRepositoryMethods(model, f)
	}

	repoFile, err := os.Create("./data/repository.go")
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
	// generate upsert
	upsertSql := `INSERT INTO ` + model.TableName + ` VALUES (`
	for idx, field := range model.Fields {
		upsertSql += fmt.Sprintf(":%s", field.ColumnName)
		if idx < (len(model.Fields) - 1) {
			upsertSql += ","
		}
	}
	upsertSql += `) ON CONFLICT(Id) DO UPDATE SET `

	for idx, field := range model.Fields {
		upsertSql += fmt.Sprintf("`%s`=excluded.`%s`", field.ColumnName, field.ColumnName)
		if idx < (len(model.Fields) - 1) {
			upsertSql += ","
		}
	}

	f.Func().Id(fmt.Sprintf("RepoUpsert%s", model.StructName)).Params(
		Id("ctx").Qual("context", "Context"), Id("db").Add(Op("*")).Qual("github.com/jmoiron/sqlx","DB"), Id("obj").Id(model.StructName),
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
	f.Line()
	f.Func().Id(fmt.Sprintf("RepoFind%sById", model.StructName)).Params(
		Id("ctx").Qual("context", "Context"), Id("db").Add(Op("*")).Qual("github.com/jmoiron/sqlx","DB"), Id("id").Id("string"),
	).Params(Op("*").Id(model.StructName), Error()).Block(
		Id("obj").Op(":=").Id(model.StructName).Values(),
		Line(),
		Id("err").Op(":=").Id("db.QueryRowx").Call(
			Lit(fmt.Sprintf("SELECT * FROM %s WHERE Id=?", model.TableName)), Id("id"),
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
	f.Line()
}

type Model struct {
	StructName string
	TableName  string
	Fields     []Field
}

type Field struct {
	FieldName string
	ColumnName string
}

func getModel(i interface{}) Model {
	var columns []Field
	columns = recursivelyGetColumns(reflect.TypeOf(i).Elem())
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

		column := Field{field.Name, field.Tag.Get("db")}
		columns = append(columns, column)
	}
	return columns
}