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

package slices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMake(t *testing.T) {
	if _cap := cap(Make[[]string](0, 0, 0)); _cap != 0 {
		t.Errorf("expect cap %d, but got %d", 0, _cap)
	}

	if _cap := cap(Make[[]string](0, 0, 1)); _cap != 1 {
		t.Errorf("expect cap %d, but got %d", 1, _cap)
	}
	if _len := len(Make[[]string](0, 0, 1)); _len != 0 {
		t.Errorf("expect len %d, but got %d", 0, _len)
	}

	if _cap := cap(Make[[]string](1, 0, 0)); _cap != 1 {
		t.Errorf("expect cap %d, but got %d", 1, _cap)
	}
	if _len := cap(Make[[]string](1, 0, 0)); _len != 1 {
		t.Errorf("expect len %d, but got %d", 1, _len)
	}
}

func TestMerge(t *testing.T) {
	s1 := []int{1, 2}
	s2 := []int{3, 4}

	if v := Merge[[]int](); v != nil {
		t.Errorf("expect a nil, but got %T", v)
	}
	if v := Merge[[]int](nil); v != nil {
		t.Errorf("expect a nil, but got %T", v)
	}

	if v := Merge[[]int](nil, nil); v != nil {
		t.Errorf("expect a nil, but got %T", v)
	} else if cap := cap(v); cap != 0 {
		t.Errorf("expect cap==%d, but got %d", 0, cap)
	}
	if v := Merge[[]int](nil, nil, nil); v != nil {
		t.Errorf("expect a nil, but got %T", v)
	} else if cap := cap(v); cap != 0 {
		t.Errorf("expect cap==%d, but got %d", 0, cap)
	}
	if v := Merge[[]int](nil, nil, []int{}); v == nil {
		t.Errorf("got unexpected nil")
	} else if cap := cap(v); cap != 0 {
		t.Errorf("expect cap==%d, but got %d", 0, cap)
	}

	if v := Merge[[]int](s1); !reflect.DeepEqual(s1, v) {
		t.Errorf("expect %v, but got %v", s1, v)
	}
	if v := Merge[[]int](s1, nil); !reflect.DeepEqual(s1, v) {
		t.Errorf("expect %v, but got %v", s1, v)
	}
	if v := Merge[[]int](nil, s1); !reflect.DeepEqual(s1, v) {
		t.Errorf("expect %v, but got %v", s1, v)
	}

	expect := []int{1, 2, 3, 4}
	if v := Merge[[]int](s1, s2); !reflect.DeepEqual(expect, v) {
		t.Errorf("expect %v, but got %v", expect, v)
	}

	expect = []int{3, 4, 1, 2}
	if v := Merge[[]int](s2, s1); !reflect.DeepEqual(expect, v) {
		t.Errorf("expect %v, but got %v", expect, v)
	}

	expect = []int{3, 4, 1, 2, 5, 6}
	if v := Merge[[]int](s2, s1, []int{5, 6}); !reflect.DeepEqual(expect, v) {
		t.Errorf("expect %v, but got %v", expect, v)
	}
}

func ExampleConvert() {
	type Ints []int

	ints1 := []int{1, 2, 3}
	ints2 := Ints{4, 5, 6}
	int64s1 := Convert(ints1, func(v int) int64 { return int64(v) })
	int64s2 := Convert(ints2, func(v int) int64 { return int64(v) })

	fmt.Println(int64s1)
	fmt.Println(int64s2)

	// Output:
	// [1 2 3]
	// [4 5 6]
}

func ExampleSetEqual() {
	s1 := []string{"a", "b", "c"}
	s2 := []string{"b", "c", "a"}
	s3 := []string{"a", "b", "b"}
	if SetEqual(s1, s2) {
		fmt.Printf("%v is equal to %v\n", s1, s2)
	} else {
		fmt.Printf("%v is not equal to %v\n", s1, s2)
	}

	if SetEqual(s1, s3) {
		fmt.Printf("%v is equal to %v\n", s1, s3)
	} else {
		fmt.Printf("%v is not equal to %v\n", s1, s3)
	}

	// Output:
	// [a b c] is equal to [b c a]
	// [a b c] is not equal to [a b b]
}

func ExampleInterfaces() {
	ss := []string{"a", "b", "c"}
	vs1 := Interfaces(ss)
	fmt.Printf("%T: %v\n", vs1, vs1)

	ints := []int{1, 2, 3}
	vs2 := Interfaces(ints)
	fmt.Printf("%T: %v\n", vs2, vs2)

	// Output:
	// []interface {}: [a b c]
	// []interface {}: [1 2 3]
}
