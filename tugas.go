package tugas7

import (
	"fmt"
	"reflect"
)

func ReflectInt(a int) {
	reflectValue := reflect.ValueOf(a)

	fmt.Println("tipe Variable :", reflectValue.Type())
	if reflectValue.Kind() == reflect.Int {

		fmt.Println("Nilai Variable :", reflectValue.Int())

	}
}

func ReflectString(a string) {
	reflectValue := reflect.ValueOf(a)

	fmt.Println("tipe Variable :", reflectValue.Type())
	if reflectValue.Kind() == reflect.Int {

		fmt.Println("Nilai Variable :", reflectValue.String())

	}
}
