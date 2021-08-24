// types.go
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

import "reflect"

type (

	// The interface to convert from unknown interface{}, builtin types, to value of the identifiered type.
	Convert interface {
		Convert(src, dst interface{}, property string) error
	}

	// The function to convert from unknown interface{} to value of the identifiered type.
	ConvertFunc func(src, dst interface{}, property string) error

	// Defines to convert from unknown interface{} to the structure object.
	// map[interface{}]interface{}, map[string]interface{} and map[int64]interface{}(optionally) are allowed types as src interface{}'s type.
	//
	// Source map object's key value are defined by member's label: `map-to:"keyname"`.
	// Keyname is valid with regular expression [a-zA-Z0-9][a-zA-Z0-9_-]*.
	// If a member is require value, append '!' to keyname. For example, `map-to:"dirname!"`
	Struct struct {
		// Defines rules assigning to member value.
		Members         []Member
		Type            reflect.Type
		allowIntegerKey bool
	}

	// Defines to convert from unknown interface{} to the structure's member object.
	Member struct {
		// Converter function, this function converts from unknown interface{}, builtin types, to value of the identifiered type.
		Convert
		// The source map object side's key name.
		Keyname string
		// The source map object side's key name.
		Keynumber int64
		// The destination structure object side's member number.
		MemberAt int
		// Shows whether converter occers error when source map doesn't have value with same keyname.
		Required bool
		// Embed value, uses same map as given source value.
		Embed bool
	}

	// Defines to buffer instance and convert from interface{}, uses only structure instance.
	Ptr struct {
		gen      reflect.Type
		Internal Convert
	}

	// Defines to buffer instance and convert from interface{}, uses only structure instance.
	Slice struct {
		gen      reflect.Type
		Internal Convert
	}
)

var (
	// Pre-compiled converters from maps to struct.
	PreCompiled = make(map[string]*Struct)
)
