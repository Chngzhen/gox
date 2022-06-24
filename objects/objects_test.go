package xobjects

import (
	"fmt"
	"testing"
)

type TestObject struct{}

func (t *TestObject) sayHello(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func TestCall(t *testing.T) {
	testObject := &TestObject{}
	results, err := Call(testObject.sayHello, "world")
	if err != nil {
		fmt.Printf("%+v", err)
	} else {
		fmt.Println(results[0].String())
	}
}

func TestGetCallerInfo(t *testing.T) {
	packageName, fileName, funcName, lineNo, ok := GetCallerInfo(0)
	fmt.Println(packageName, fileName, funcName, lineNo, ok)

	packageName, fileName, funcName, lineNo, ok = GetCallerInfo(1)
	fmt.Println(packageName, fileName, funcName, lineNo, ok)

	packageName, fileName, funcName, lineNo, ok = GetCallerInfo(2)
	fmt.Println(packageName, fileName, funcName, lineNo, ok)
}
