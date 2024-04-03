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

// TryAdd tries to add the key-value pair into the maps
// if the key does not exist. Or, do nothing and return false.
func TryAdd[M ~map[K]V, K comparable, V any](maps M, k K, v V) (ok bool) {
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

func FromSliceWithIndex[S ~[]E, K comparable, V, E any](s S, convert func(int, E) (K, V)) map[K]V {
	_len := len(s)
	maps := make(map[K]V, _len)
	for i := 0; i < _len; i++ {
		k, v := convert(i, s[i])
		maps[k] = v
	}
	return maps
}

func FromSlice[S ~[]E, K comparable, V, E any](s S, convert func(E) (K, V)) map[K]V {
	return FromSliceWithIndex(s, func(_ int, e E) (K, V) { return convert(e) })
}

// SetMap converts a slice s to a set map.
func SetMap[S ~[]T, T comparable](s S) map[T]struct{} {
	return FromSlice(s, func(e T) (T, struct{}) { return e, struct{}{} })
}

// BoolMap converts a slice s to a bool map.
func BoolMap[S ~[]T, T comparable](s S) map[T]bool {
	return FromSlice(s, func(e T) (T, bool) { return e, true })
}

// SetMapFunc converts a slice s to a set map by a conversion function.
func SetMapFunc[S ~[]T, K comparable, T any](s S, convert func(T) K) map[K]struct{} {
	return FromSlice(s, func(e T) (K, struct{}) { return convert(e), struct{}{} })
}

// BoolMapFunc converts a slice s to a bool map by a conversion function.
func BoolMapFunc[S ~[]T, K comparable, T any](s S, convert func(T) K) map[K]bool {
	return FromSlice(s, func(e T) (K, bool) { return convert(e), true })
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

// KeysFunc returns all the keys of the map by the conversion function.
func KeysFunc[M ~map[K]V, T any, K comparable, V any](maps M, convert func(K) T) []T {
	keys := make([]T, 0, len(maps))
	for k := range maps {
		keys = append(keys, convert(k))
	}
	return keys
}

// Values returns all the values of the map by the conversion function.
func ValuesFunc[M ~map[K]V, T any, K comparable, V any](maps M, convert func(V) T) []T {
	values := make([]T, 0, len(maps))
	for _, v := range maps {
		values = append(values, convert(v))
	}
	return values
}
