package svix

import (
	"net/url"
	"reflect"
	"strconv"
)

type apiClient struct {
}

// parameterAddToHeaderOrQuery adds the provided object to the request header or url query
// supporting deep object syntax
func parameterAddToHeaderOrQuery(headerOrQueryParams interface{}, keyPrefix string, obj interface{}) {
	var v = reflect.ValueOf(obj)
	var value = ""
	if v == reflect.ValueOf(nil) {
		value = "null"
	} else {
		switch v.Kind() {
		case reflect.Invalid:
			value = "invalid"

		case reflect.Struct:
			panic("unexpect query or header parameter struct value")

		// both queryName=x&queryName=y and queryName=x,y are supported server side
		case reflect.Slice:
			var indValue = reflect.ValueOf(obj)
			if indValue == reflect.ValueOf(nil) {
				return
			}
			var lenIndValue = indValue.Len()
			for i := 0; i < lenIndValue; i++ {
				var arrayValue = indValue.Index(i)
				var keyPrefixForCollectionType = keyPrefix
				parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefixForCollectionType, arrayValue.Interface())
			}
			return

		case reflect.Map:
			panic("unexpect query or header parameter map value")

		case reflect.Interface:
			fallthrough
		case reflect.Ptr:
			parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, v.Elem().Interface())
			return

		case reflect.Int, reflect.Int8, reflect.Int16,
			reflect.Int32, reflect.Int64:
			value = strconv.FormatInt(v.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16,
			reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			value = strconv.FormatUint(v.Uint(), 10)
		case reflect.Float32, reflect.Float64:
			value = strconv.FormatFloat(v.Float(), 'g', -1, 32)
		case reflect.Bool:
			value = strconv.FormatBool(v.Bool())
		case reflect.String:
			value = v.String()
		default:
			value = v.Type().String() + " value"
		}
	}

	switch valuesMap := headerOrQueryParams.(type) {
	case url.Values:
		valuesMap.Add(keyPrefix, value)
		break
	case map[string]string:
		valuesMap[keyPrefix] = value
		break
	}
}
