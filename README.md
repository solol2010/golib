# Go Utils Library

This repository provides a collection of commonly used generic data structures and utility functions for Go programming. It aims to offer reusable and efficient components to simplify development.

Currently, the library includes:
*   **Data Structures:** Thread-safe generic implementations of Queue and Stack.
*   **Slice Utilities:** A suite of generic functions for common slice operations like finding, mapping, filtering, and reducing.
*   **String Utilities:** A set of helper functions for string manipulations.

## Data Structures (`datastruct` package)

This package provides generic and thread-safe implementations of fundamental data structures.

### Queue

A generic First-In-First-Out (FIFO) queue.

*   `NewQueue[T any]() *queue[T]`: Creates and returns a new empty queue.
*   `Push(e T)`: Adds an element to the end of the queue.
*   `Pop() (T, error)`: Removes and returns the element from the front of the queue. Returns an error if the queue is empty.
*   `Get() (T, error)`: Returns the element at the front of the queue without removing it. Returns an error if the queue is empty.
*   `Length() int`: Returns the number of elements in the queue.

The queue is thread-safe and can be used concurrently.

### Stack

A generic Last-In-First-Out (LIFO) stack.

*   `NewStack[T any]() *stack[T]`: Creates and returns a new empty stack.
*   `Push(e T)`: Adds an element to the top of the stack.
*   `Pop() (T, error)`: Removes and returns the element from the top of the stack. Returns an error if the stack is empty.
*   `Get() (T, error)`: Returns the element at the top of the stack without removing it. Returns an error if the stack is empty.
*   `Length() int`: Returns the number of elements in the stack.

The stack is thread-safe and can be used concurrently.

## Slice Utilities (`slicefuncs` package)

This package offers a collection of generic functions to perform common operations on slices.

*   `FindAll[T any](slice []T, condition func(T) bool) []T`: Returns a new slice containing all elements from the input slice that satisfy the given condition function.
*   `FindFirst[T any](slice []T, condition func(T) bool) (int, T)`: Returns the index and value of the first element in the slice that satisfies the condition. If no such element is found, it returns -1 and the zero value for type `T`.
*   `FindLast[T any](slice []T, condition func(T) bool) (int, T)`: Returns the index and value of the last element in the slice that satisfies the condition. If no such element is found, it returns -1 and the zero value for type `T`.
*   `ForEach[T any](slice []T, action func(e *T))`: Executes the given action function for each element in the slice. The action function receives a pointer to the element.
*   `ForEachParal[T any](slice []T, n int, action func(e *T))`: Executes the given action function for each element in the slice in parallel using `n` goroutines. The action function receives a pointer to the element.
*   `Map[T1 any, T2 any](slice []T1, mapper func(T1) T2) []T2`: Transforms a slice of type `T1` into a slice of type `T2` by applying the `mapper` function to each element.
*   `MapMethod[T1 any, T2 any](slice []T1, method func(*T1) T2) []T2`: Transforms a slice by calling a method (defined on `*T1`) on each element, returning a new slice of type `T2`.
*   `MapParal[T1 any, T2 any](slice []T1, n int, mapper func(T1) T2) []T2`: Transforms a slice of type `T1` to `T2` in parallel using `n` goroutines, applying the `mapper` function to each element.
*   `Reduce[T1 any, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2`: Reduces a slice `s` to a single value of type `T2`. It starts with an `initializer` value and applies the function `f` cumulatively to the elements.
*   `Filter[T any](s []T, f func(T) bool) []T`: Returns a new slice containing only the elements from slice `s` for which the filter function `f` returns `true`.
*   `FilterBySlice[T1 any, T2 any](s1 []T1, s2 []T2, match func(T1, T2) bool) []T1`: Returns a new slice containing elements from `s1` that have a corresponding matching element in `s2`, based on the `match` function.
*   `Reverse[T any](s []T) []T`: Returns a new slice with the elements of `s` in reverse order.
*   `GroupCount[T comparable](s []T) map[T]int`: Counts the occurrences of each element in slice `s` and returns the counts as a map. (Note: `T` must be comparable).
*   `Max[T any](slice []T, smaller func(T, T) bool) (index int, max T)`: Finds the maximum element in a slice. Requires a `smaller` function that returns `true` if the first argument is less than the second. Returns the index and the maximum value. If the slice is empty, returns -1 and the zero value for `T`.
*   `Min[T any](slice []T, smaller func(T, T) bool) (index int, min T)`: Finds the minimum element in a slice. Requires a `smaller` function that returns `true` if the first argument is less than the second. Returns the index and the minimum value. If the slice is empty, returns -1 and the zero value for `T`.

All functions in this package are generic and operate on slices of any type, provided the type constraints (e.g., `comparable` for `GroupCount`, or the provision of a comparison function for `Max`/`Min`) are met.

## String Utilities (`strs` package)

This package contains helper functions for common string manipulations.

*   `SplitLast(s, sep string) string`: Splits string `s` by separator `sep` and returns the last part.
*   `SplitFirst(s, sep string) string`: Splits string `s` by separator `sep` and returns the first part.
*   `Reverse(s string) string`: Reverses the input string `s`.
*   `TrimMultiSpace(s string) string`: Removes all leading and trailing spaces from string `s`. (Note: This function repeatedly trims spaces from both ends. It does not reduce multiple internal spaces to a single space.)

## How to Use

To use these utilities in your Go project, you can import the desired packages. For example:

```go
import (
	"fmt"
	"your_repository_path/datastruct" // Replace your_repository_path with the actual import path
	"your_repository_path/slicefuncs"
	"your_repository_path/strs"
)

func main() {
	// Using datastruct.Queue
	q := datastruct.NewQueue[int]()
	q.Push(1)
	q.Push(2)
	val, _ := q.Pop()
	fmt.Println("Popped from queue:", val) // Output: 1

	// Using slicefuncs.Map
	nums := []int{1, 2, 3}
	doubled := slicefuncs.Map(nums, func(n int) int {
		return n * 2
	})
	fmt.Println("Doubled numbers:", doubled) // Output: [2 4 6]

	// Using strs.Reverse
	original := "hello"
	reversed := strs.Reverse(original)
	fmt.Println("Reversed string:", reversed) // Output: olleh
}
```

**Note:** Replace `your_repository_path` with the actual import path for this library (e.g., `github.com/yourusername/your-repo-name`). You'll also need to ensure your project is set up to fetch modules from this path.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please feel free to:

1.  Open an issue to discuss the changes.
2.  Fork the repository, make your changes, and submit a pull request.

Please ensure your code follows the existing style and includes tests where appropriate.
