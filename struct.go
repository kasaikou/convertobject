// struct.go
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
	"regexp"
	"strconv"

	"github.com/streamwest-1629/convertobject/util"
)

const (
	Label              = `map-to`
	LabelRegexp        = `^(?P<key>[a-zA-Z0-9][a-zA-Z0-9_-]*)(\!)?$`
	LabelRequireRegexp = `[a-zA-Z0-9][a-zA-Z0-9_-]*(\!)$`
	LabelEmbed         = `<-`
)

var (
	labelMatches        = regexp.MustCompile(LabelRegexp)
	labelRequireMatches = regexp.MustCompile(LabelRequireRegexp)
)

// Build converter from interface{} to structure object.
//
// Built object is managed by the module mapping, to use as cache converter building.
// When object type has already build before, This function returns cached converter.
func CompileStruct(target interface{}) (compiled *Struct, err error) {

	// initialize to check cache
	__type := formatStructType(reflect.TypeOf(target))
	if convert, err := selectConvert(__type, &PreCompiled); err != nil {
		return nil, err
	} else {
		return convert.(*Struct), nil
	}

}

// Build converter from interface{} to structure object.
//
// Built object isn't managed by the module mapping.
func CompileStructIndepended(target interface{}) (compiled *Struct, err error) {

	// compile
	__type := formatStructType(reflect.TypeOf(target))
	if convert, err := selectConvert(__type, &map[string]*Struct{}); err != nil {
		return nil, err
	} else {
		return convert.(*Struct), nil
	}

}

// Force to build converter from interface{} to structure object.
// If function failed to make it, occer panic().
//
// Built object is managed by the module mapping, to use as cache converter building.
// When object type has already build before, This function returns cached converter.
func CompileStructForce(target interface{}) *Struct {

	if compiled, err := CompileStruct(target); err != nil {
		panic(err.Error())
	} else {
		return compiled
	}
}

// Force to build converter from interface{} to structure object.
// If function failed to make it, occer panic().
//
// Built object isn't managed by the module mapping.
func CompileStructIndependedForce(target interface{}) *Struct {

	if compiled, err := CompileStructIndepended(target); err != nil {
		panic(err.Error())
	} else {
		return compiled
	}

}

// TODO: WRITE COMMENT
func (s *Struct) Generate(src interface{}) (dst interface{}, err error) {
	return GeneratePtr(src, s, s.Type, "")
}

func formatStructType(__type reflect.Type) reflect.Type {
	switch __type.Kind() {
	case reflect.Struct:
		return __type
	case reflect.Ptr:
		return formatStructType(__type.Elem())
	default:
		panic("the type used in compile is invalid, allow structure or structure's pointer")
	}
}

func compileStruct(__type reflect.Type, compiled *Struct, cache *map[string]*Struct) error {

	(*compiled) = Struct{
		Members:         make([]Member, 0),
		Type:            __type,
		allowIntegerKey: true,
	}

	for i, l := 0, __type.NumField(); i < l; i++ {

		field := __type.Field(i)
		label := field.Tag.Get(Label)

		if matches := labelMatches.FindStringSubmatchIndex(label); matches != nil {

			keyname := string(labelMatches.ExpandString([]byte{}, `${key}`, label, matches))
			keynumber := int64(0)

			if id, err := strconv.ParseInt(keyname, 0, 64); err != nil {
				compiled.allowIntegerKey = false
			} else {
				keynumber = id
			}

			if convert, err := selectConvert(field.Type, cache); err != nil {
				return err
			} else {
				compiled.Members = append(compiled.Members,
					Member{
						Convert:   convert,
						Keyname:   keyname,
						Keynumber: keynumber,
						MemberAt:  i,
						Required:  labelRequireMatches.MatchString(keyname),
					})
			}
		} else if label == LabelEmbed {

			if convert, err := selectConvert(field.Type, cache); err != nil {
				return err
			} else {
				compiled.Members = append(compiled.Members,
					Member{
						Convert:  convert,
						Embed:    true,
						MemberAt: i,
					})
			}
		}
	}
	return nil
}

func (c *Struct) Convert(src, dst interface{}, property string) error {

	var (
		val reflect.Value
	)

	if ptr := reflect.ValueOf(dst); ptr.Kind() == reflect.Ptr {
		val = ptr.Elem()
	} else {
		val = ptr
	}

	// check type are convertible
	if ty := val.Type(); !c.Type.AssignableTo(ty) || !ty.AssignableTo(c.Type) {
		panic(util.ErrInvalidType(property, reflect.New(c.Type).Interface(), dst))
	}

	// initialize convert function
	var MemberProperty func(member *Member) string
	var AssignToMember = func(member *Member, src interface{}, property string) error {
		dst := val.Field(member.MemberAt).Addr().Interface()
		return member.Convert.Convert(src, dst, property)
	}
	if len(property) > 0 {
		MemberProperty = func(member *Member) string {
			return property + "." + member.Keyname
		}
	} else {
		MemberProperty = func(member *Member) string {
			return member.Keyname
		}
	}

	// convert from map[interface{}]interface{}
	if mapped, ok := src.(map[interface{}]interface{}); ok {
		for _, member := range c.Members {

			memProperty := MemberProperty(&member)

			if member.Embed {
				if err := AssignToMember(&member, src, memProperty); err != nil {
					return err
				}
				continue
			} else if buf, exist := mapped[member.Keyname]; exist {
				if err := AssignToMember(&member, buf, memProperty); err != nil {
					return err
				}
				continue
			} else if c.allowIntegerKey {
				if buf, exist := mapped[int(member.Keynumber)]; exist {
					if err := AssignToMember(&member, buf, memProperty); err != nil {
						return err
					}
					continue
				} else if buf, exist := mapped[int64(member.Keynumber)]; exist {
					if err := AssignToMember(&member, buf, memProperty); err != nil {
						return err
					}
					continue
				}
			}

			// check property is required member
			if member.Required {
				return util.ErrorCannotFound(memProperty)
			}
		}
		return nil
	}

	// convert from map[string]interface{}
	if mapped, ok := src.(map[string]interface{}); ok {
		for _, member := range c.Members {

			memProperty := MemberProperty(&member)
			if member.Embed {
				if err := AssignToMember(&member, src, memProperty); err != nil {
					return err
				}
				continue
			} else if buf, exist := mapped[member.Keyname]; exist {
				if err := AssignToMember(&member, buf, memProperty); err != nil {
					return err
				}
			} else if member.Required {
				return util.ErrorCannotFound(memProperty)
			}
		}

		return nil
	}

	if c.allowIntegerKey {

		// convert from map[int]interface{}
		if mapped, ok := src.(map[int]interface{}); ok {

			for _, member := range c.Members {

				memProperty := MemberProperty(&member)

				if member.Embed {
					if err := AssignToMember(&member, src, memProperty); err != nil {
						return err
					}
					continue
				} else if buf, exist := mapped[int(member.Keynumber)]; exist {
					if err := AssignToMember(&member, buf, memProperty); err != nil {
						return err
					}
				} else if member.Required {
					return util.ErrorCannotFound(memProperty)
				}
			}

			return nil
		}

		// convert from map[int64]interface{}
		if mapped, ok := src.(map[int64]interface{}); ok {

			for _, member := range c.Members {

				memProperty := MemberProperty(&member)

				if member.Embed {
					if err := AssignToMember(&member, src, memProperty); err != nil {
						return err
					}
					continue
				} else if buf, exist := mapped[int64(member.Keynumber)]; exist {
					if err := AssignToMember(&member, buf, memProperty); err != nil {
						return err
					}
				} else if member.Required {
					return util.ErrorCannotFound(memProperty)
				}
			}

			return nil
		}
	}

	return util.ErrInvalidType(property, &map[string]interface{}{}, src)
}
