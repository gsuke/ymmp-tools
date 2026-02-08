package ymmp

import (
	"fmt"
	"reflect"
)

func PrintSummary(v interface{}) {
	rv := reflect.ValueOf(v)
	rt := reflect.TypeOf(v)

	// ポインタ対応
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
		rt = rt.Elem()
	}

	if rv.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < rv.NumField(); i++ {
		fv := rv.Field(i)
		ft := rt.Field(i)

		name := ft.Tag.Get("json")
		if name == "" {
			name = ft.Name
		}

		switch fv.Kind() {
		case reflect.String:
			fmt.Printf("%s = %q\n", name, fv.String())

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fmt.Printf("%s = %d\n", name, fv.Int())

		case reflect.Bool:
			fmt.Printf("%s = %v\n", name, fv.Bool())

		case reflect.Float32, reflect.Float64:
			fmt.Printf("%s = %f\n", name, fv.Float())

		case reflect.Slice, reflect.Array:
			fmt.Printf("%s = [%d items]\n", name, fv.Len())

		case reflect.Map:
			fmt.Printf("%s = { %d items }\n", name, fv.Len())

		default:
			// 要検証: 本当に消して良いのか？ 値が入っているポインタも (ptr) となって見逃している可能性もある
			// fmt.Printf("%s = (%s)\n", name, fv.Kind())
		}
	}
}
