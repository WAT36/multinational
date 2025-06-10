package notes

import (
	"fmt"
	"reflect"
)

type Point2 struct {
	X int "X座標"
	Y int "Y座標"
}

func Tag() {
	p := Point2{X: 1, Y: 2}

	t := reflect.TypeOf(p)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Tag)
	}
}
