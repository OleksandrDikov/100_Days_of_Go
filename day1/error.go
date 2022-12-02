package main

import (
	"fmt"
	"reflect"
	// "errors"
)

func main() {
	// err := errors.New("The new error massage!")
	// fmt.Println(err.Error())

	err := fmt.Errorf("The new error massage for %s!", "Server")
	fmt.Println(err)
	fmt.Println(reflect.TypeOf(err))
	fmt.Println(err.Error())
	fmt.Println(reflect.TypeOf(err.Error()))
}
