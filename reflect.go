package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	t := reflect.TypeOf(3)
	fmt.Println(t.String())
	fmt.Println(t)

	var w io.Writer
	w = os.Stdout
	fmt.Println(reflect.TypeOf(w))

	v := reflect.ValueOf(3)
	fmt.Println(v)

	user := User{Name: "nonocd", Password: "123456", NickName: "hhahah"}

	v = reflect.ValueOf(user)
	t = v.Type()
	fmt.Printf("type: %s\n", t)
	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldType := t.Field(i)
		fmt.Printf("field name: %10s, field type: %s, field tag: %20s, field value: %s\n", fieldType.Name, fieldType.Type, fieldType.Tag, fieldValue.String())
	}
}

// User ...
type User struct {
	Name     string `json:"name"`
	Password string `json:"name"`
	NickName string `json:"nick_name"`
}
