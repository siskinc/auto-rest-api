package autorestapi

import "reflect"

func getStructName(t interface{}) (name string) {
	ty := reflect.TypeOf(t)
	name = ty.Elem().Name()
	switch ty.Kind() {
	case reflect.Ptr:
		name = ty.Elem().Name()
	default:
		panic("[getStructName] t parameter have to be a point")
	}
	return
}
