package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
	"github.com/deslee/cms/model"
	"os"
	"reflect"
	"strings"
)

type tabler interface {
	TableName() string
}

func main() {
	generate(
		&model.User{},
		&model.Site{},
		&model.Item{},
		&model.Group{},
		&model.Asset{},
		&model.SiteUser{},
		&model.ItemGroup{},
		&model.ItemAsset{},
	)
}

func generate(modelTypes ...interface{}) {
	f := NewFile("repository")

	f.HeaderComment("Code generated by genRepo.go, DO NOT EDIT.")

	for _, modelType := range modelTypes {
		model := getModel(modelType)
		generateRepositoryMethods(model, f)
	}

	repoFile, err := os.Create("./repository/generated.go")
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
	writeQueryFunctionsForModel(model, f)
	writeGetterFunctionsForModel(model, model.primaryKeyFields(), f)
	for _, field := range model.Fields {
		if field.needsGetter() {
			writeGetterFunctionsForModel(model, []Field{field}, f)
		}
	}
	writeUpsertFunctionsForModel(model, f)
	writeDeleteFunctionsForModel(model, model.primaryKeyFields(), f)

	f.Line()
}

func writeQueryFunctionsForModel(model Model, f *File) {
	f.Line()

	f.Func().Id(fmt.Sprintf("Scan%sList", model.StructName)).Params(
		Id("ctx").Qual("context", "Context"),
		Id("db").Add(Op("*")).Qual("github.com/jmoiron/sqlx", "DB"),
		Id("query").Id("string"),
		Id("args").Id("...interface{}"),
	).Params(Index().Qual("github.com/deslee/cms/model", model.StructName), Error()).Block(
		Var().Id("list").Index().Qual("github.com/deslee/cms/model", model.StructName),
		List(Id("rows"), Id("err")).Op(":=").Id("db.Queryx").Call(
			Id("query"),
			Id("args..."),
		),
		If(Id("err").Op("!=").Id("nil")).Block(
			Return(List(Id("nil"), Id("err"))),
		),
		Defer().Id("rows.Close").Call(),
		For(Id("rows.Next").Call()).Block(
			Var().Id("obj").Qual("github.com/deslee/cms/model", model.StructName),
			Id("err").Op("=").Id("rows.StructScan").Call(Add(Op("&")).Id("obj")),
			If(Id("err").Op("!=").Id("nil")).Block(
				Return(List(Id("nil"), Id("err"))),
			),
			Id("list").Op("=").Id("append").Call(Id("list"), Id("obj")),
		),
		Return(Id("list"), Id("nil")),
	)
}

func writeDeleteFunctionsForModel(model Model, fieldsToQueryOn []Field, f *File) {
	writeDeleteFunctionsForModelForCtxCreator(model, fieldsToQueryOn, "DB", f)
	writeDeleteFunctionsForModelForCtxCreator(model, fieldsToQueryOn, "Tx", f)
}

func writeDeleteFunctionsForModelForCtxCreator(model Model, fieldsToQueryOn []Field, ctxCreator string, f *File) {
	// create params of the delete method
	params := []Code{
		Id("ctx").Qual("context", "Context"),
		Id(strings.ToLower(ctxCreator)).Add(Op("*")).Qual("github.com/jmoiron/sqlx", ctxCreator),
	}

	// begin creating delete query
	deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE ", model.TableName)

	// append the keys as arguments
	// append the keys as query params
	for idx, keyField := range fieldsToQueryOn {
		params = append(params, Id(fmt.Sprintf("key%s", keyField.FieldName)).Id("string"))
		if idx == 0 {
			deleteQuery += fmt.Sprintf("%s=?", keyField.FieldName)
		} else {
			deleteQuery += fmt.Sprintf(" AND %s=?", keyField.FieldName)
		}
	}

	// create the arguments code block for the actual db call
	dbExecArguments := []Code{
		Lit(deleteQuery),
	}

	methodNameEnding := ""

	// append the keys as sql parameterized arguments
	for idx, keyField := range fieldsToQueryOn {
		dbExecArguments = append(dbExecArguments, Id(fmt.Sprintf("key%s", keyField.FieldName)))
		if idx == 0 {
			methodNameEnding += keyField.FieldName
		} else {
			methodNameEnding += fmt.Sprintf("And%s", keyField.FieldName)
		}
	}

	functionName := fmt.Sprintf("Delete%sBy%s", model.StructName, methodNameEnding)

	if ctxCreator == "DB" {

	} else if ctxCreator == "Tx" {
		functionName += "Tx"
	} else {
		panic(fmt.Sprintf("Unrecognized argument to writeUpsertMethodForModelForCtxCreator %s", ctxCreator))
	}

	f.Line()

	f.Func().Id(functionName).Params(
		params...,
	).Params(Error()).Block(
		List(Id("_"), Id("err")).Op(":=").Id(fmt.Sprintf("%s.Exec", strings.ToLower(ctxCreator))).Call(
			dbExecArguments...,
		),
		Line(),
		If(Id("err").Op("!=").Id("nil")).Block(
			Return(Id("err")),
		),
		Line(),
		Return(Id("nil")),
	)
}

func writeGetterFunctionsForModel(model Model, fieldsToQueryOn []Field, f *File) {
	f.Line()

	// create params of the getter method
	params := []Code{
		Id("ctx").Qual("context", "Context"),
		Id("db").Add(Op("*")).Qual("github.com/jmoiron/sqlx", "DB"),
	}

	// begin creating select query
	selectQuery := fmt.Sprintf("SELECT * FROM %s WHERE ", model.TableName)

	// append the keys as arguments
	// append the keys as query params
	for idx, keyField := range fieldsToQueryOn {
		params = append(params, Id(fmt.Sprintf("key%s", keyField.FieldName)).Id("string"))
		if idx == 0 {
			selectQuery += fmt.Sprintf("%s=?", keyField.FieldName)
		} else {
			selectQuery += fmt.Sprintf(" AND %s=?", keyField.FieldName)
		}
	}

	// create the arguments code block for the actual db call
	queryRowArguments := []Code{
		Lit(selectQuery),
	}

	methodNameEnding := ""

	// append the keys as sql parameterized arguments
	for idx, keyField := range fieldsToQueryOn {
		queryRowArguments = append(queryRowArguments, Id(fmt.Sprintf("key%s", keyField.FieldName)))
		if idx == 0 {
			methodNameEnding += keyField.FieldName
		} else {
			methodNameEnding += fmt.Sprintf("And%s", keyField.FieldName)
		}
	}

	f.Func().Id(fmt.Sprintf("Find%sBy%s", model.StructName, methodNameEnding)).Params(
		params...,
	).Params(Op("*").Qual("github.com/deslee/cms/model", model.StructName), Error()).Block(
		Id("obj").Op(":=").Qual("github.com/deslee/cms/model", model.StructName).Values(),
		Line(),
		Id("err").Op(":=").Id("db.QueryRowx").Call(
			queryRowArguments...,
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

func writeUpsertFunctionsForModel(model Model, f *File) {
	writeUpsertMethodForModelForCtxCreator(model, f, "DB")
	f.Line()
	writeUpsertMethodForModelForCtxCreator(model, f, "Tx")
}

func writeUpsertMethodForModelForCtxCreator(model Model, f *File, ctxCreator string) {
	var functionName = fmt.Sprintf("Upsert%s", model.StructName)

	if ctxCreator == "DB" {

	} else if ctxCreator == "Tx" {
		functionName += "Tx"
	} else {
		panic(fmt.Sprintf("Unrecognized argument to writeUpsertMethodForModelForCtxCreator %s", ctxCreator))
	}

	// generate upsert sql
	upsertSql := `INSERT INTO ` + model.TableName + ` VALUES (`
	for idx, field := range model.Fields {
		upsertSql += fmt.Sprintf(":%s", field.columnName())
		if idx < (len(model.Fields) - 1) {
			upsertSql += ","
		}
	}
	upsertSql += `) ON CONFLICT(`

	for idx, field := range model.primaryKeyFields() {
		if idx == 0 {
			upsertSql += field.columnName()
		} else {
			upsertSql += fmt.Sprintf(",%s", field.columnName())
		}
	}

	upsertSql += `) DO UPDATE SET `
	for idx, field := range model.Fields {
		upsertSql += fmt.Sprintf("`%s`=excluded.`%s`", field.columnName(), field.columnName())
		if idx < (len(model.Fields) - 1) {
			upsertSql += ","
		}
	}

	f.Func().Id(functionName).Params(
		Id("ctx").Qual("context", "Context"), Id(strings.ToLower(ctxCreator)).Add(Op("*")).Qual("github.com/jmoiron/sqlx", ctxCreator), Id("obj").Qual("github.com/deslee/cms/model", model.StructName),
	).Error().Block(
		List(Id("stmt"), Id("err")).Op(":=").Id(fmt.Sprintf("%s.PrepareNamedContext", strings.ToLower(ctxCreator))).Call(Id("ctx"), Lit(upsertSql)),
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

func (model Model) primaryKeyFields() []Field {
	var primaryKeyFields []Field
	for _, field := range model.Fields {
		if field.isPk() {
			primaryKeyFields = append(primaryKeyFields, field)
		}
	}
	return primaryKeyFields
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

func (f Field) needsGetter() bool {
	needsGetter := f.TagMap["needsGetter"]
	return len(needsGetter) > 0
}

func getModel(i interface{}) Model {
	columns := recursivelyGetColumns(reflect.TypeOf(i).Elem())
	structName := reflect.TypeOf(i).Elem().Name()
	modelTabler, ok := i.(tabler)
	if ok == false {
		panic(fmt.Sprintf("Model %s must implement tabler! Write a Table() method e.g: \nfunc (%s) TableName() string {\n\treturn \"%ss\"\n}", structName, structName, structName))
	}

	return Model{
		StructName: structName,
		TableName:  modelTabler.TableName(),
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
