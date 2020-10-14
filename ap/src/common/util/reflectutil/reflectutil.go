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

// IsPointer 指定したReflectTypeがPointerかどうかを判定
// true: pointer, false: 値
func IsPointer(reflectType reflect.Type) bool {
	return reflectType.Kind() == reflect.Ptr
}

// IsNil 確実にnilを検出する
// golangのnilは型を持つため、interface{}でnilを受け取ったときに、キャストしないとnil判定ができない
// そんな時のために、nilかどうかの判定ができるようにしたもの
// structの場合は、初期値があるのでnilではない判定にします
func IsNil(x interface{}) bool {

	// reflectはコストが高いので、先にreflectを使わない判定をします
	if x == nil {
		return true
	}

	// structの場合は、reflect.ValueOf(x).IsNil()でpanicするため先に判定
	if !IsPointer(reflect.TypeOf(x)) {
		// structなので、nilではありません
		return false
	}

	return reflect.ValueOf(x).IsNil()
}
