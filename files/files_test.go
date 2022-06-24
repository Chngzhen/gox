package files

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRegExp(t *testing.T) {
	sql := "sqsds.3gp"
	rst := regexp.MustCompile("\\.[a-zA-z1-9]*$").MatchString(sql)
	if !rst {
		t.Error("期望结果：;", true, "实际结果：", rst)
	}
}

func TestExistsFile(t *testing.T) {
	if _, ok := IsExisted("files.go"); !ok {
		t.Error("期望结果：;", true, "实际结果：", false)
	}
}

func TestCheckDir(t *testing.T) {
	isExisted, err := CheckDir("D:/home/lmode", false)
	if err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		fmt.Println(isExisted)
	}
}

func TestExtendsIn(t *testing.T) {
	extends := []string{"txt", "jpg"}
	fmt.Println(ExtendsIn("test.jpeg", extends))
}

func TestRetrieveFiles(t *testing.T) {
	fileChannel := make(chan string, 1000)
	err := RetrieveFiles("D:/home", fileChannel)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	close(fileChannel)

	for file := range fileChannel {
		fmt.Println(file)
	}
}
