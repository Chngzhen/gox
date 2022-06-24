package xobjects

import (
	"errors"
	"path"
	"reflect"
	"runtime"
	"strings"
)

// GetCallerInfo 获取调用者信息，包括包名、文件名、方法名盒行号。若无法获取，则返回false。
// 当skip=0时，获取被调用者的信息；当skip>0时，根据调用栈获取指定调用者的信息。
func GetCallerInfo(skip int) (packageName, fileNameWithoutExt, funcName string, lineNo int, successful bool) {
	pc, filePath, lineNo, ok := runtime.Caller(skip)
	if ok {
		fileName := path.Base(filePath)
		fileExt := path.Ext(filePath)
		fileNameWithoutExt = fileName[0 : len(fileName)-len(fileExt)]

		function := runtime.FuncForPC(pc)
		functionFullName := function.Name()
		delimiterIndexBetweenPkgAndFunc := strings.LastIndex(functionFullName, ".")
		packageName = functionFullName[0:delimiterIndexBetweenPkgAndFunc]
		funcName = functionFullName[delimiterIndexBetweenPkgAndFunc+1:]
		return packageName, fileNameWithoutExt, funcName, lineNo, true
	}
	return "", "", "", 0, false
}

// Call 利用反射调用未知结构体的函数。
func Call(method interface{}, params ...interface{}) (result []reflect.Value, err error) {
	function := reflect.ValueOf(method)
	// 校验实参数量是否与函数的形参数量一致
	if len(params) != function.Type().NumIn() {
		err = errors.New("函数的实参数量不匹配")
		return nil, err
	}

	// 将参数构建成reflect.Value数组，供反射调用
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	return function.Call(in), nil
}
