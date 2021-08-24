// example_test.go
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

package convertobject_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/streamwest-1629/convertobject"
)

func TestExample(t *testing.T) {

	Example()

}

func Example() {

	type Data struct {
		Version float64 `map-to:"version"`
	}

	// Definates target structure type
	type Person struct {
		Name     string   `map-to:"name!"`
		Date     int64    `map-to:"date!"`
		Old      int      `map-to:"old"`
		Children []Person `map-to:"children"`
		Partner  *Person  `map-to:"partner"`
		Data     `map-to:"info"`
	}

	// Source map and Destination object defines
	src, dest := map[string]interface{}{
		"name": "John",
		"date": "19950216",
		"old":  26,
		"children": []interface{}{
			map[interface{}]interface{}{
				"name": "Amy",
				"date": 20171104,
			},
		},
		"partner": map[string]interface{}{
			"name": "Ada",
			"date": 19950425,
		},
		"info": map[string]interface{}{
			"version": 0.11,
		},
	}, Person{Old: -1}

	// Compile and Convert
	if err := convertobject.DirectConvert(src, &dest); err != nil {
		panic(err.Error())
	}

	// output struct with json
	bytes, _ := json.MarshalIndent(dest, "", "    ")
	fmt.Println(string(bytes))
	// {
	//     "Name": "John",
	//     "Date": 19950216,
	//     "Old": 26,
	//     "Children": [
	//         {
	//             "Name": "Amy",
	//             "Date": 20171104,
	//             "Old": 0,
	//             "Children": null,
	//             "Partner": null
	//         }
	//     ],
	//     "Partner": {
	//         "Name": "Ada",
	//         "Date": 19950425,
	//         "Old": 0,
	//         "Children": null,
	//         "Partner": null
	//     }
	// }

}
