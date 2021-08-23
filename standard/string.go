// standard/string.go
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

import "github.com/streamwest-1629/convertobject/util"

func ConvertoString(src, dst interface{}, property string) error {
	if destination, ok := dst.(*string); !ok {
		panic(util.ErrInvalidType(property, destination, dst))
	} else if val, ok := src.(*string); ok {
		*destination = *val
	} else {
		return util.ErrInvalidType(property, destination, src)
	}
	return nil
}
