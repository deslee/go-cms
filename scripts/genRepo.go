package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
	"github.com/deslee/cms/data"
	"reflect"
)

type tabler interface {
	TableName() string
}

func main() {
	generate(
		&data.User{},
		&data.Site{},
		&data.SiteUser{},
		&data.ItemGroup{},
		&data.ItemAsset{},
		&data.Item{},
		&data.Group{},
		&data.Asset{},
	)
}

func generate(modelTypes ... interface{}) {
	f := NewFile("data")
	for _, modelType := range modelTypes {
		model := getModel(modelType)
		generateRepositoryMethods(model, f)
	}

	fmt.Printf("%#v", f)

}

func generateRepositoryMethods(model Model, f *File) {
	// generate upsert
	upsertSql := `INSERT INTO ` + model.TableName + ` VALUES (`
	for _, field := range model.Fields {
		upsertSql += fmt.Sprintf(":%s,", field.ColumnName)
	}
	upsertSql += `) ON CONFLICT(Id) DO UPDATE SET `

	for _, field := range model.Fields {
		upsertSql += fmt.Sprintf("%s=excluded.%s,", field.ColumnName, field.ColumnName)
	}

	f.Func().Id(fmt.Sprintf("Upsert%s", model.StructName)).Params(
		Id("ctx").Qual("context", "Context"), Id("db").Add(Op("*")).Qual("sqlx","DB"), Id("obj").Id(model.StructName),
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