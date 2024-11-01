package internal

import "reflect"

func ReflectDeeperEqual(expected, actual any) bool {
	v1 := reflect.ValueOf(expected)
	v2 := reflect.ValueOf(actual)

	if v1.Kind() == reflect.Ptr {
		v1 = v1.Elem()
	}

	if v2.Kind() == reflect.Ptr {
		v2 = v2.Elem()
	}

	if !v1.IsValid() && !v2.IsValid() {
		return true
	}

	switch v1.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		if v1.IsNil() {
			v1 = reflect.ValueOf(nil)
		}
	}

	switch v2.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		if v2.IsNil() {
			v2 = reflect.ValueOf(nil)
		}
	}

	v1Underlying := reflect.Zero(reflect.TypeOf(v1)).Interface()
	v2Underlying := reflect.Zero(reflect.TypeOf(v2)).Interface()

	if v1 == v1Underlying {
		if v2 == v2Underlying {
			return reflect.DeepEqual(v1, v2)
		} else {
			return reflect.DeepEqual(v1, v2.Interface())
		}
	} else {
		if v2 == v2Underlying {
			return reflect.DeepEqual(v1.Interface(), v2)
		} else {
			return reflect.DeepEqual(v1.Interface(), v2.Interface())
		}
	}
}
