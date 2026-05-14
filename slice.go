package lo

// Filter iterates over elements of collection, returning an array of
// all elements predicate returns truthy for.
func Filter[V any](collection []V, predicate func(item V, index int) bool) []V {
	result := make([]V, 0, len(collection))
	for i, item := range collection {
		if predicate(item, i) {
			result = append(result, item)
		}
	}
	return result
}

// Map manipulates a slice and transforms it to a slice of another type.
func Map[T any, R any](collection []T, iteratee func(item T, index int) R) []R {
	result := make([]R, len(collection))
	for i, item := range collection {
		result[i] = iteratee(item, i)
	}
	return result
}

// Reduce reduces collection to a value which is the accumulated result of
// running each element in collection thru accumulator, where each successive
// invocation is supplied the return value of the previous.
func Reduce[T any, R any](collection []T, accumulator func(agg R, item T, index int) R, initial R) R {
	for i, item := range collection {
		initial = accumulator(initial, item, i)
	}
	return initial
}

// ForEach iterates over elements of collection and invokes iteratee for each element.
func ForEach[T any](collection []T, iteratee func(item T, index int)) {
	for i, item := range collection {
		iteratee(item, i)
	}
}

// Uniq returns a duplicate-free version of an array, in which only the first
// occurrence of each element is kept. The order of result values is determined
// by the order they occur in the array.
func Uniq[T comparable](collection []T) []T {
	seen := make(map[T]struct{}, len(collection))
	result := make([]T, 0, len(collection))
	for _, item := range collection {
		if _, ok := seen[item]; !ok {
			seen[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// Contains returns true if an element is present in a collection.
func Contains[T comparable](collection []T, element T) bool {
	for _, item := range collection {
		if item == element {
			return true
		}
	}
	return false
}

// Flatten returns an array a single level deep.
func Flatten[T any](collection [][]T) []T {
	totalLen := 0
	for _, inner := range collection {
		totalLen += len(inner)
	}
	result := make([]T, 0, totalLen)
	for _, inner := range collection {
		result = append(result, inner...)
	}
	return result
}

// Chunk returns an array of elements split into groups the length of size.
// If array can't be split evenly, the final chunk will be the remaining elements.
func Chunk[T any](collection []T, size int) [][]T {
	if size <= 0 {
		panic("lo: chunk size must be greater than 0")
	}
	result := make([][]T, 0, (len(collection)+size-1)/size)
	for i := 0; i < len(collection); i += size {
		end := i + size
		if end > len(collection) {
			end = len(collection)
		}
		chunk := make([]T, end-i)
		copy(chunk, collection[i:end])
		result = append(result, chunk)
	}
	return result
}

// Reverse reverses a slice in place and returns it.
// Note: this modifies the original slice. If you need a copy, clone it first.
func Reverse[T any](collection []T) []T {
	length := len(collection)
	half := length / 2
	for i := 0; i < half; i++ {
		j := length - 1 - i
		collection[i], collection[j] = collection[j], collection[i]
	}
	return collection
}
