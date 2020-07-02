package reflectutil

import "reflect"

// Util Reflect Utility

// CreateSlice .
func CreateSlice(reflectType reflect.Type) reflect.Value {
	t := reflect.SliceOf(reflectType)
	makedSlice := reflect.MakeSlice(t, 0, 16)
	reflectionValue := reflect.New(t)
	reflectionValue.Elem().Set(makedSlice)
	slicePtr := reflect.ValueOf(reflectionValue.Interface())
	return slicePtr.Elem()
}
