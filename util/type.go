// util/type.go
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

import (
	"fmt"
	"reflect"
)

// Get fullname of type, including package name and type name.
func TypeFullname(__type reflect.Type) string {
	if pkg, name := Typename(__type); len(pkg) > 0 {
		return pkg + "." + name
	} else {
		return name
	}
}

// Get type name and containing package name.
func Typename(__type reflect.Type) (pkg string, name string) {

	const (
		Ptr   = "(*%s)"
		Arr   = "[%d]%s"
		Slice = "[]%s"
		Map   = "map[%s]%s"
		Chan  = "chan <- %s"
		Other = "%s"
	)

	switch __type.Kind() {
	case reflect.Ptr:
		pkg, name := Typename(__type.Elem())
		return pkg, fmt.Sprintf(Ptr, name)
	case reflect.Array:
		pkg, name := Typename(__type.Elem())
		return pkg, fmt.Sprintf(Arr, __type.Len(), name)
	case reflect.Slice:
		pkg, name := Typename(__type.Elem())
		return pkg, fmt.Sprintf(Slice, name)
	case reflect.Map:
		return "", fmt.Sprintf(Map, TypeFullname(__type.Key()), TypeFullname(__type.Elem()))
	case reflect.Chan:
		return "", fmt.Sprintf(Chan, TypeFullname(__type.Elem()))
	default:
		return __type.PkgPath(), __type.Name()
	}
}
