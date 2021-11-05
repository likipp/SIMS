package utils

import "reflect"

func StructConvMap(params interface{}) map[string]interface{} {
	t := reflect.TypeOf(params)
	v := reflect.ValueOf(params)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}