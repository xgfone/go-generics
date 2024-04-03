// Copyright 2023 xgfone
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package maps

import (
	"fmt"
	"slices"
	"sort"
	"testing"
)

func ExampleConvert() {
	type Maps map[string]int

	intmap1 := Maps{"a": 1, "b": 2}
	intmap2 := map[string]int{"a": 3, "b": 4}

	int64map1 := Convert(intmap1, func(k string, v int) (string, int64) { return k, int64(v) })
	int64map2 := Convert(intmap2, func(k string, v int) (string, int64) { return k, int64(v) })

	fmt.Printf("%T\n", int64map1)
	fmt.Printf("%T\n", int64map2)
	fmt.Printf("%s=%v\n", "a", int64map1["a"])
	fmt.Printf("%s=%v\n", "b", int64map1["b"])
	fmt.Printf("%s=%v\n", "a", int64map2["a"])
	fmt.Printf("%s=%v\n", "b", int64map2["b"])

	// Output:
	// map[string]int64
	// map[string]int64
	// a=1
	// b=2
	// a=3
	// b=4
}

func ExampleAddSlice() {
	type Map map[string]int
	type Slice []byte

	maps1 := Map{"a": 1}
	maps2 := map[string]int{"a": 1}
	AddSlice(maps1, []byte{'b'}, func(v byte) (string, int) { return string(v), int(v) })
	AddSlice(maps2, Slice{'b'}, func(v byte) (string, int) { return string(v), int(v) })

	fmt.Printf("%s=%v\n", "a", maps1["a"])
	fmt.Printf("%s=%v\n", "b", maps1["b"])
	fmt.Printf("%s=%v\n", "a", maps2["a"])
	fmt.Printf("%s=%v\n", "b", maps2["b"])

	// Output:
	// a=1
	// b=98
	// a=1
	// b=98
}

func ExampleDeleteSlice() {
	type Map map[string]int
	type Slice []string

	maps1 := Map{"a": 1, "b": 2, "c": 3}
	maps2 := map[string]int{"a": 1, "b": 2, "c": 3}
	DeleteSlice(maps1, []string{"a", "b"})
	DeleteSlice(maps2, Slice{"a", "b"})

	fmt.Println(maps1)
	fmt.Println(maps2)

	// Output:
	// map[c:3]
	// map[c:3]
}

func ExampleDeleteSliceFunc() {
	type Map map[string]int
	type Slice []byte

	maps1 := Map{"a": 1, "b": 2, "c": 3}
	maps2 := map[string]int{"a": 1, "b": 2, "c": 3}
	DeleteSliceFunc(maps1, []byte{'a', 'b'}, func(b byte) string { return string(b) })
	DeleteSliceFunc(maps2, Slice{'a', 'b'}, func(b byte) string { return string(b) })

	fmt.Println(maps1)
	fmt.Println(maps2)

	// Output:
	// map[c:3]
	// map[c:3]
}

func TestKeys(t *testing.T) {
	expectints := []int{1, 2}
	intmap := map[int]int{1: 11, 2: 22}
	ints := Keys(intmap)
	sort.Ints(ints)
	if !slices.Equal(expectints, ints) {
		t.Errorf("expect %v, but got %v", expectints, ints)
	}

	expectstrs := []string{"a", "b"}
	strmap := map[string]string{"a": "aa", "b": "bb"}
	strs := Keys(strmap)
	sort.Strings(strs)
	if !slices.Equal(expectstrs, strs) {
		t.Errorf("expect %v, but got %v", expectstrs, strs)
	}
}

func TestValues(t *testing.T) {
	expectints := []int{11, 22}
	intmap := map[int]int{1: 11, 2: 22}
	ints := Values(intmap)
	sort.Ints(ints)
	if !slices.Equal(expectints, ints) {
		t.Errorf("expect %v, but got %v", expectints, ints)
	}

	expectstrs := []string{"aa", "bb"}
	strmap := map[string]string{"a": "aa", "b": "bb"}
	strs := Values(strmap)
	sort.Strings(strs)
	if !slices.Equal(expectstrs, strs) {
		t.Errorf("expect %v, but got %v", expectstrs, strs)
	}
}

func ExampleKeysFunc() {
	type Key struct {
		K string
		V int32
	}
	maps := map[Key]bool{
		{K: "a", V: 1}: true,
		{K: "b", V: 2}: true,
		{K: "c", V: 3}: true,
	}

	keys := KeysFunc(maps, func(k Key) string { return k.K })
	slices.Sort(keys)
	fmt.Println(keys)

	// Output:
	// [a b c]
}

func ExampleValuesFunc() {
	type Value struct {
		V int
	}
	maps := map[string]Value{
		"a": {V: 1},
		"b": {V: 2},
		"c": {V: 3},
	}

	values := ValuesFunc(maps, func(v Value) int { return v.V })
	slices.Sort(values)
	fmt.Println(values)

	// Output:
	// [1 2 3]
}

func ExampleSetMap() {
	setmap := SetMap([]string{"a", "b", "c"})
	fmt.Println(setmap)

	// Output:
	// map[a:{} b:{} c:{}]
}

func ExampleBoolMap() {
	boolmap := BoolMap([]string{"a", "b", "c"})
	fmt.Println(boolmap)

	// Output:
	// map[a:true b:true c:true]
}

func ExampleSetMapFunc() {
	type S struct {
		K string
		V int32
	}

	values := []S{{K: "a", V: 1}, {K: "b", V: 2}, {K: "c", V: 3}}
	setmap := SetMapFunc(values, func(s S) string { return s.K })
	fmt.Println(setmap)

	// Output:
	// map[a:{} b:{} c:{}]
}

func ExampleBoolMapFunc() {
	type S struct {
		K string
		V int32
	}

	values := []S{{K: "a", V: 1}, {K: "b", V: 2}, {K: "c", V: 3}}
	setmap := BoolMapFunc(values, func(s S) string { return s.K })
	fmt.Println(setmap)

	// Output:
	// map[a:true b:true c:true]
}
