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

// Package maps provides some convenient map generic functions.
package maps

import (
	"github.com/xgfone/go-generics/set"
	"golang.org/x/exp/constraints"
)

// Make returns a new map.
//
// If both cap and defaultCap are equal to 0, it is equal to make(M).
// If cap is equal to 0, use defaultCap as cap instead, which is equal to
// make(M, defaultCap).
func Make[M ~map[K]V, K comparable, V any, I constraints.Integer](cap, defaultCap I) M {
	if cap == 0 {
		if cap = defaultCap; cap == 0 {
			return make(M)
		}
	}
	return make(M, cap)
}

// Add adds the key-value pair into the maps if the key does not exist.
// Or, do nothing and return false.
func Add[M ~map[K]V, K comparable, V any](maps M, k K, v V) (ok bool) {
	_, exist := maps[k]
	if ok = !exist; ok {
		maps[k] = v
	}
	return
}

// AddSlice adds each element of the slices into maps, which convert the slice
// element to the key-value pair by the convert fucntion.
func AddSlice[M ~map[K]V, S ~[]E, K comparable, V, E any](maps M, slices S, convert func(E) (K, V)) {
	for _, value := range slices {
		k, v := convert(value)
		maps[k] = v
	}
}

// AddSliceAsValue adds each element of the slices as the value into maps,
// which gets the key by the fucntion getkey.
func AddSliceAsValue[M ~map[K]V, S ~[]V, K comparable, V any](maps M, slices S, getkey func(V) K) {
	for _, v := range slices {
		maps[getkey(v)] = v
	}
}

// Pop removes the element by the key and returns the removed value.
func Pop[M ~map[K]V, K comparable, V any](maps M, k K) (v V, ok bool) {
	if v, ok = maps[k]; ok {
		delete(maps, k)
	}
	return
}

// Delete removes the element by the key.
func Delete[M ~map[K]V, K comparable, V any](maps M, k K) (ok bool) {
	if _, ok = maps[k]; ok {
		delete(maps, k)
	}
	return
}

// DeleteSlice deletes a set of values as the keys.
func DeleteSlice[M ~map[K]V, S ~[]K, K comparable, V any](maps M, keys S) {
	for _, key := range keys {
		delete(maps, key)
	}
}

// DeleteSliceFunc deletes a set of values, which converts the key from K2 to K1
// by the function getkey.
func DeleteSliceFunc[M ~map[K1]V, S ~[]K2, K1, K2 comparable, V any](maps M, keys S, getkey func(K2) K1) {
	for _, key := range keys {
		delete(maps, getkey(key))
	}
}

// Values returns all the values of the map.
func Values[M ~map[K]V, K comparable, V any](maps M) []V {
	values := make([]V, 0, len(maps))
	for _, v := range maps {
		values = append(values, v)
	}
	return values
}

// Keys returns all the keys of the map.
func Keys[M ~map[K]V, K comparable, V any](maps M) []K {
	keys := make([]K, 0, len(maps))
	for k := range maps {
		keys = append(keys, k)
	}
	return keys
}

// Clone clones the map and returns the new.
func Clone[M ~map[K]V, K comparable, V any](maps M) M {
	if maps == nil {
		return nil
	}

	newmap := make(M, len(maps))
	for k, v := range maps {
		newmap[k] = v
	}
	return newmap
}

// Convert converts the map from map[K1]V1 to map[K1]V2.
func Convert[M ~map[K1]V1, K1, K2 comparable, V1, V2 any](maps M, convert func(K1, V1) (K2, V2)) map[K2]V2 {
	if maps == nil {
		return nil
	}

	newmap := make(map[K2]V2, len(maps))
	for k1, v1 := range maps {
		k2, v2 := convert(k1, v1)
		newmap[k2] = v2
	}
	return newmap
}

// ConvertValues clones the map, converts the value from V1 to V2, and returns the new.
func ConvertValues[M ~map[K]V1, K comparable, V1, V2 any](maps M, convert func(V1) V2) map[K]V2 {
	if maps == nil {
		return nil
	}

	newmap := make(map[K]V2, len(maps))
	for k, v := range maps {
		newmap[k] = convert(v)
	}
	return newmap
}

// KeysToSet collects all the keys and converts them to a set.
func KeysToSet[M ~map[K]V, K comparable, V any](maps M) set.Set[K] {
	kset := set.NewSetWithCap[K](len(maps))
	for k := range maps {
		kset.Add(k)
	}
	return kset
}
