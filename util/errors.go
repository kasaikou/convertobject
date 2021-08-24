// util/errors.go
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

package util

import "reflect"

type (
	errInvalidType struct {
		propName string
		wantType string
		hasType  string
	}
	errCannotFound struct {
		propName string
	}
)

func (e *errInvalidType) Error() string {
	return e.propName + " is invalid type (want: " + e.wantType + ", has: " + e.hasType + ")"
}
func ErrInvalidType(propName string, want interface{}, has interface{}) error {
	return &errInvalidType{
		propName: propName,
		wantType: TypeFullname(reflect.TypeOf(want)),
		hasType:  TypeFullname(reflect.TypeOf(has)),
	}
}

func (e *errCannotFound) Error() string {
	return e.propName + " is required property, but cannot found it"
}
func ErrCannotFound(propName string) error {
	return &errCannotFound{
		propName: propName,
	}
}
