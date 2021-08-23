// ptr_slice.go
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
	"strconv"

	"github.com/streamwest-1629/convertobject/util"
)

func (p *Ptr) Convert(src, dst interface{}, property string) error {

	if ptr := reflect.ValueOf(dst).Elem(); ptr.Kind() != reflect.Ptr {
		panic(util.ErrInvalidType(property, reflect.New(p.gen).Addr().Addr(), dst))
	} else {
		ptr.Set(reflect.New(p.gen))
		return p.Internal.Convert(src, ptr.Interface(), property)
	}
}

func (s *Slice) Convert(src, dst interface{}, property string) error {

	if destination := reflect.ValueOf(dst).Elem(); destination.Kind() != reflect.Slice {
		panic(util.ErrInvalidType(property, reflect.MakeSlice(s.gen, 0, 0).Addr(), dst))
	} else if buf, ok := src.(*[]interface{}); ok {

		destination.Set(reflect.MakeSlice(s.gen, len(*buf), 0))
		for i, val := range *buf {
			if err := s.Internal.Convert(val, destination.Index(i).Addr().Interface(), property+"["+strconv.Itoa(i)+"]"); err != nil {
				return err
			}
		}
	}

	return nil
}
