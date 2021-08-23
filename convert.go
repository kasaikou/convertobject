// convert.go
// Copyright (C) 2021 Kasai Koji

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 	http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package convertobject

import (
	"reflect"

	"github.com/streamwest-1629/convertobject/standard"
	"github.com/streamwest-1629/convertobject/util"
)

// TODO: WRITE COMMENT
func GeneratePtr(src interface{}, converter Convert, __type reflect.Type, property string) (destPtr interface{}, err error) {

	destPtr = reflect.New(__type).Interface()
	err = converter.Convert(src, destPtr, property)
	return
}

func (c ConvertFunc) Convert(src, dst interface{}, property string) error {
	return c(src, dst, property)
}

func selectConvert(__type reflect.Type, cache *map[string]Struct) (Convert, error) {
	switch __type.Kind() {
	case reflect.Int64:
		return ConvertFunc(standard.ConvertoInt64), nil
	case reflect.Int32:
		return ConvertFunc(standard.ConvertoInt32), nil
	case reflect.Int16:
		return ConvertFunc(standard.ConvertoInt16), nil
	case reflect.Int8:
		return ConvertFunc(standard.ConvertoInt8), nil
	case reflect.Int:
		return ConvertFunc(standard.ConvertoInt), nil
	case reflect.Uint64:
		return ConvertFunc(standard.ConvertoUint64), nil
	case reflect.Uint32:
		return ConvertFunc(standard.ConvertoUint32), nil
	case reflect.Uint16:
		return ConvertFunc(standard.ConvertoUint16), nil
	case reflect.Uint8:
		return ConvertFunc(standard.ConvertoUint8), nil
	case reflect.Uint:
		return ConvertFunc(standard.ConvertoUint), nil
	case reflect.Float64:
		return ConvertFunc(standard.ConvertoFloat64), nil
	case reflect.Float32:
		return ConvertFunc(standard.ConvertoFloat32), nil
	case reflect.Bool:
		return ConvertFunc(standard.ConvertoBool), nil
	case reflect.String:
		return ConvertFunc(standard.ConvertoString), nil
	case reflect.Ptr:
		elem := __type.Elem()
		if gen, err := selectConvert(elem, cache); err != nil {
			return nil, err
		} else {
			return &Ptr{
				gen:      elem,
				Internal: gen,
			}, nil
		}
	case reflect.Slice:
		elem := __type.Elem()
		if gen, err := selectConvert(elem, cache); err != nil {
			return nil, err
		} else {
			return &Slice{
				gen:      elem,
				Internal: gen,
			}, nil
		}
	case reflect.Struct:

		__name := util.TypeFullname(__type)

		// check cache
		if val, exist := (*cache)[__name]; exist {
			return &val, nil
		} else {

			// compile
			if compiled, err := compileStruct(__type, cache); err != nil {
				return nil, err
			} else {
				(*cache)[__name] = *compiled
				return compiled, nil
			}
		}
	}

	panic("not supported type: " + util.TypeFullname(__type))
}
