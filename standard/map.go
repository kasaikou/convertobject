// standard/map.go
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

package standard

import (
	"reflect"
	"strconv"

	"github.com/streamwest-1629/convertobject/util"
)

func ConvertoInterfaceKeyInterfaceMap(src, dst interface{}, property string) error {

	var destination *map[interface{}]interface{}

	if dest, ok := dst.(*map[interface{}]interface{}); !ok {
		panic(util.ErrInvalidType(property, destination, dst).Error())
	} else {
		destination = dest
		AllocateNewMapOnNil(dest)
	}

	if mapped, ok := src.(map[string]interface{}); ok {
		for key, val := range mapped {
			(*destination)[key] = val
		}
	} else if mapped, ok := src.(map[int]interface{}); ok {
		for key, val := range mapped {
			(*destination)[key] = val
		}
	} else if mapped, ok := src.(map[int64]interface{}); ok {
		for key, val := range mapped {
			(*destination)[key] = val
		}
	} else if mapped, ok := src.(map[interface{}]interface{}); ok {
		for key, val := range mapped {
			(*destination)[key] = val
		}
	} else {
		return util.ErrInvalidType(property, *destination, src)
	}
	return nil
}

func ConvertoStringKeyInterfaceMap(src, dst interface{}, property string) error {

	var destination *map[string]interface{}

	if dest, ok := dst.(*map[string]interface{}); !ok {
		panic(util.ErrInvalidType(property, destination, dst).Error())
	} else {
		destination = dest
		AllocateNewMapOnNil(dest)
	}

	if mapped, ok := src.(map[string]interface{}); ok {
		for key, val := range mapped {
			(*destination)[key] = val
		}
	} else if mapped, ok := src.(map[int]interface{}); ok {
		for key, val := range mapped {
			(*destination)[strconv.Itoa(key)] = val
		}
	} else if mapped, ok := src.(map[int64]interface{}); ok {
		for key, val := range mapped {
			(*destination)[strconv.FormatInt(key, 10)] = val
		}
	} else if mapped, ok := src.(map[interface{}]interface{}); ok {
		for key, val := range mapped {
			keyStr := ""
			if err := ConvertoString(key, &keyStr, property+".(key)"); err != nil {
				return err
			}
			(*destination)[keyStr] = val
		}
	} else {
		return util.ErrInvalidType(property, *destination, src)
	}
	return nil
}

func ConvertoStringKeyStringMap(src, dst interface{}, property string) error {

	var destination *map[string]string

	if dest, ok := dst.(*map[string]string); !ok {
		panic(util.ErrInvalidType(property, destination, dst).Error())
	} else {
		destination = dest
		AllocateNewMapOnNil(dest)
	}

	if mapped, ok := src.(map[string]string); ok {
		for key, val := range mapped {
			(*destination)[key] = val
		}
	} else if mapped, ok := src.(map[string]interface{}); ok {
		for key, val := range mapped {
			valStr := ""
			if err := ConvertoString(val, &valStr, property+".(key)"); err != nil {
				return err
			}
			(*destination)[key] = valStr
		}
	} else if mapped, ok := src.(map[interface{}]interface{}); ok {
		for key, val := range mapped {
			keyStr, valStr := "", ""
			if err := ConvertoString(key, &keyStr, property+".(key)"); err != nil {
				return err
			} else if err := ConvertoString(val, &valStr, property+".(value)"); err != nil {
				return err
			}
			(*destination)[keyStr] = valStr
		}
	}
	return nil
}

func AllocateNewMapOnNil(dst interface{}) {
	if val := reflect.ValueOf(dst).Elem(); val.IsNil() {
		val.Set(reflect.MakeMap(val.Type()))
	}
}
