// standard/number_boolean.go
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
	"strconv"

	"github.com/streamwest-1629/convertobject/util"
)

func ConvertoFloat64(src, dst interface{}, property string) error {
	if destination, ok := dst.(*float64); !ok {
		panic(util.ErrInvalidType(property, destination, dst).Error())
	} else if val, ok := src.(*float64); ok {
		*destination = float64(*val)
	} else if val, ok := src.(*float32); ok {
		*destination = float64(*val)
	} else if val, ok := src.(*int64); ok {
		*destination = float64(*val)
	} else if val, ok := src.(*int32); ok {
		*destination = float64(*val)
	} else if val, ok := src.(*int16); ok {
		*destination = float64(*val)
	} else if val, ok := src.(*int8); ok {
		*destination = float64(*val)
	} else if val, ok := src.(*int); ok {
		*destination = float64(*val)
	} else if val, ok := src.(*uint64); ok {
		*destination = float64(*val)
	} else if val, ok := src.(*uint32); ok {
		*destination = float64(*val)
	} else if val, ok := src.(*uint16); ok {
		*destination = float64(*val)
	} else if val, ok := src.(*uint8); ok {
		*destination = float64(*val)
	} else if val, ok := src.(*uint); ok {
		*destination = float64(*val)
	} else if val, ok := src.(*byte); ok {
		*destination = float64(*val)
	} else if val, ok := src.(*string); ok {
		if val, err := strconv.ParseFloat(*val, 64); err != nil {
			return err
		} else {
			*destination = val
		}
	} else {
		return util.ErrInvalidType(property, destination, src)
	}
	return nil

}

func ConvertoFloat32(src, dst interface{}, property string) error {
	buf := float64(0)

	if destination, ok := dst.(*float32); !ok {
		panic(util.ErrInvalidType(property, destination, dst).Error())
	} else if err := ConvertoFloat64(src, &buf, property); err != nil {
		return err
	} else {
		*destination = float32(buf)
		return nil
	}
}

func ConvertoBool(src, dst interface{}, property string) error {
	if destination, ok := dst.(*bool); !ok {
		panic(util.ErrInvalidType(property, destination, dst).Error())
	} else if val, ok := src.(bool); ok {
		*destination = val
	} else if val, ok := src.(string); ok {
		if val, err := strconv.ParseBool(val); err != nil {
			return err
		} else {
			*destination = val
		}
	} else {
		return util.ErrInvalidType(property, destination, src)
	}
	return nil
}
