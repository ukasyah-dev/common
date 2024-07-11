package handler

import (
	"reflect"
	"runtime"
	"strings"

	"github.com/iancoleman/strcase"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func getFuncName(f any) string {
	funcName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	funcNameSegments := strings.Split(funcName, ".")
	funcName = funcNameSegments[len(funcNameSegments)-1]
	funcName = strings.Replace(funcName, "-fm", "", 1)
	funcName = strings.ReplaceAll(strcase.ToSnake(funcName), "_", " ")
	funcName = cases.Title(language.English, cases.Compact).String(funcName)

	return funcName
}
