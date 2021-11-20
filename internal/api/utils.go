package api

import (
	"errors"
	"reflect"
)

func StructToMap(input interface{}, output map[string]interface{}) error {

	//val := reflect.ValueOf(output)

	//if val.Kind() != reflect.Ptr {
	//	return errors.New("result must be a pointer")
	//}

	valueOf := reflect.ValueOf(input)
	if valueOf.Kind() != reflect.Struct {
		return errors.New("input must be a struct")
	}
	typeOf := reflect.TypeOf(input)
	for i := 0; i < valueOf.NumField(); i++ {
		value := valueOf.Field(i)
		field := typeOf.Field(i)

		switch value.Kind() {
		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int, reflect.Int64:
			output[field.Name] = value.Int()
		case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint, reflect.Uint64:
			output[field.Name] = value.Uint()
		case reflect.Float32, reflect.Float64:
			output[field.Name] = value.Float()
		case reflect.String:
			output[field.Name] = value.String()
		case reflect.Bool:
			output[field.Name] = value.Bool()
		case reflect.Struct:
			output[field.Name] = value
		case reflect.Map:
			output[field.Name] = value
		case reflect.Slice:
			output[field.Name] = value
		case reflect.Array:
			output[field.Name] = value
		case reflect.Ptr:
			//addMetaKey, err = d.decodePtr(name, input, outVal)
		case reflect.Func:
			//err = d.decodeFunc(name, input, outVal)
		default:
		}
	}

	return nil
}
func MapToStruct(input interface{}, output interface{}) error {
	return nil
}
func decodeStructFromMap() {}
func decodeMapFromStruct() {}
