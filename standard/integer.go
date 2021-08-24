// standard/integer.go
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
	"math"
	"strconv"

	"github.com/streamwest-1629/convertobject/util"
)

func ConvertoInt64(src, dst interface{}, property string) error {
	if destination, ok := dst.(*int64); !ok {
		panic(util.ErrInvalidType(property, destination, dst).Error())
	} else if val, ok := src.(int64); ok {
		*destination = int64(val)
	} else if val, ok := src.(int32); ok {
		*destination = int64(val)
	} else if val, ok := src.(int16); ok {
		*destination = int64(val)
	} else if val, ok := src.(int8); ok {
		*destination = int64(val)
	} else if val, ok := src.(int); ok {
		*destination = int64(val)
	} else if val, ok := src.(uint64); ok {
		if val > math.MaxInt64 {
			return util.ErrInvalidType(property, *destination, val)
		}
		*destination = int64(val)
	} else if val, ok := src.(uint32); ok {
		*destination = int64(val)
	} else if val, ok := src.(uint16); ok {
		*destination = int64(val)
	} else if val, ok := src.(uint8); ok {
		*destination = int64(val)
	} else if val, ok := src.(uint); ok {
		*destination = int64(val)
	} else if val, ok := src.(byte); ok {
		*destination = int64(val)
	} else if val, ok := src.(string); ok {
		if val, err := strconv.ParseInt(val, 0, 64); err != nil {
			return err
		} else {
			*destination = val
		}
	} else {
		util.ErrInvalidType(property, *destination, src)
	}
	return nil
}

func ConvertoUint64(src, dst interface{}, property string) error {

	if destination, ok := dst.(*uint64); !ok {
		panic(util.ErrInvalidType(property, destination, dst).Error())
	} else if val, ok := src.(int64); ok {
		if val < 0 {
			return util.ErrInvalidType(property, *destination, val)
		}
		*destination = uint64(val)
	} else if val, ok := src.(int32); ok {
		if val < 0 {
			return util.ErrInvalidType(property, *destination, val)
		}
		*destination = uint64(val)
	} else if val, ok := src.(int16); ok {
		if val < 0 {
			return util.ErrInvalidType(property, *destination, val)
		}
		*destination = uint64(val)
	} else if val, ok := src.(int8); ok {
		if val < 0 {
			return util.ErrInvalidType(property, *destination, val)
		}
		*destination = uint64(val)
	} else if val, ok := src.(int); ok {
		if val < 0 {
			return util.ErrInvalidType(property, *destination, val)
		}
		*destination = uint64(val)
	} else if val, ok := src.(uint64); ok {
		*destination = uint64(val)
	} else if val, ok := src.(uint32); ok {
		*destination = uint64(val)
	} else if val, ok := src.(uint16); ok {
		*destination = uint64(val)
	} else if val, ok := src.(uint8); ok {
		*destination = uint64(val)
	} else if val, ok := src.(uint); ok {
		*destination = uint64(val)
	} else if val, ok := src.(byte); ok {
		*destination = uint64(val)
	} else if val, ok := src.(string); ok {
		if val, err := strconv.ParseUint(val, 0, 64); err != nil {
			return err
		} else {
			*destination = val
		}
	} else {
		util.ErrInvalidType(property, *destination, src)
	}
	return nil
}

func ConvertoInt32(src, dst interface{}, property string) error {
	buf := int64(0)
	if destination, ok := dst.(*int32); !ok {
		panic(util.ErrInvalidType(property, destination, dst).Error())
	} else if err := ConvertoInt64(src, &buf, property); err != nil {
		return err
	} else {
		*destination = int32(buf)
		return nil
	}
}
func ConvertoInt16(src, dst interface{}, property string) error {
	buf := int64(0)
	if destination, ok := dst.(*int16); !ok {
		panic(util.ErrInvalidType(property, destination, dst).Error())
	} else if err := ConvertoInt64(src, &buf, property); err != nil {
		return err
	} else {
		*destination = int16(buf)
		return nil
	}
}
func ConvertoInt8(src, dst interface{}, property string) error {
	buf := int64(0)
	if destination, ok := dst.(*int8); !ok {
		panic(util.ErrInvalidType(property, destination, dst).Error())
	} else if err := ConvertoInt64(src, &buf, property); err != nil {
		return err
	} else {
		*destination = int8(buf)
		return nil
	}
}
func ConvertoInt(src, dst interface{}, property string) error {
	buf := int64(0)
	if destination, ok := dst.(*int); !ok {
		panic(util.ErrInvalidType(property, destination, dst).Error())
	} else if err := ConvertoInt64(src, &buf, property); err != nil {
		return err
	} else {
		*destination = int(buf)
		return nil
	}
}

func ConvertoUint32(src, dst interface{}, property string) error {
	buf := uint64(0)
	if destination, ok := dst.(*uint32); !ok {
		panic(util.ErrInvalidType(property, destination, dst).Error())
	} else if err := ConvertoUint64(src, &buf, property); err != nil {
		return err
	} else {
		*destination = uint32(buf)
		return nil
	}
}

func ConvertoUint16(src, dst interface{}, property string) error {
	buf := uint64(0)
	if destination, ok := dst.(*uint16); !ok {
		panic(util.ErrInvalidType(property, destination, dst).Error())
	} else if err := ConvertoUint64(src, &buf, property); err != nil {
		return err
	} else {
		*destination = uint16(buf)
		return nil
	}
}

func ConvertoUint8(src, dst interface{}, property string) error {
	buf := uint64(0)
	if destination, ok := dst.(*uint8); !ok {
		panic(util.ErrInvalidType(property, destination, dst).Error())
	} else if err := ConvertoUint64(src, &buf, property); err != nil {
		return err
	} else {
		*destination = uint8(buf)
		return nil
	}
}

func ConvertoUint(src, dst interface{}, property string) error {
	buf := uint64(0)
	if destination, ok := dst.(*uint); !ok {
		panic(util.ErrInvalidType(property, destination, dst).Error())
	} else if err := ConvertoUint64(src, &buf, property); err != nil {
		return err
	} else {
		*destination = uint(buf)
		return nil
	}
}

func ConvertoByte(src, dst interface{}, property string) error {
	buf := uint64(0)
	if destination, ok := dst.(*byte); !ok {
		panic(util.ErrInvalidType(property, destination, dst).Error())
	} else if err := ConvertoUint64(src, &buf, property); err != nil {
		return err
	} else {
		*destination = byte(buf)
		return nil
	}
}
