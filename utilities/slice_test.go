package utilities

import (
	"testing"
	"reflect"
)

func TestFilter(t *testing.T) {
	// Test case 1: Filter integers with a simple condition (even numbers)
	t.Run("filter even integers", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		expected := []int{2, 4, 6, 8, 10}

		result := Filter(input, func(item int, index int) bool {
			return item%2 == 0
		})

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Filter() = %v, want %v", result, expected)
		}
	})

	// Test case 2: Filter strings by length
	t.Run("filter strings by length", func(t *testing.T) {
		input := []string{"apple", "banana", "kiwi", "strawberry", "fig"}
		expected := []string{"apple", "kiwi", "fig"}

		result := Filter(input, func(item string, index int) bool {
			return len(item) <= 5
		})

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Filter() = %v, want %v", result, expected)
		}
	})

	// Test case 3: Filter with empty slice
	t.Run("filter empty slice", func(t *testing.T) {
		var input []int

		result := Filter(input, func(item int, index int) bool {
			return item > 0
		})

		if len(result) != 0 {
			t.Errorf("Filter() returned %d elements, want 0", len(result))
		}
	})

	// Test case 4: Filter where no elements match the predicate
	t.Run("no elements match predicate", func(t *testing.T) {
		input := []int{1, 3, 5, 7, 9}

		result := Filter(input, func(item int, index int) bool {
			return item%2 == 0
		})

		if len(result) != 0 {
			t.Errorf("Filter() returned %d elements, want 0", len(result))
		}
	})

	// Test case 5: Filter where all elements match the predicate
	t.Run("all elements match predicate", func(t *testing.T) {
		input := []int{2, 4, 6, 8, 10}
		expected := []int{2, 4, 6, 8, 10}

		result := Filter(input, func(item int, index int) bool {
			return item%2 == 0
		})

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Filter() = %v, want %v", result, expected)
		}
	})

	// Test case 6: Filter using the index parameter
	t.Run("filter using index", func(t *testing.T) {
		input := []string{"a", "b", "c", "d", "e"}
		expected := []string{"a", "c", "e"}

		result := Filter(input, func(item string, index int) bool {
			return index%2 == 0
		})

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Filter() = %v, want %v", result, expected)
		}
	})

	// Test case 7: Filter with custom struct type
	t.Run("filter custom struct", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}

		input := []Person{
			{Name: "Alice", Age: 25},
			{Name: "Bob", Age: 17},
			{Name: "Charlie", Age: 30},
			{Name: "David", Age: 15},
		}

		expected := []Person{
			{Name: "Alice", Age: 25},
			{Name: "Charlie", Age: 30},
		}

		result := Filter(input, func(item Person, index int) bool {
			return item.Age >= 18
		})

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Filter() = %v, want %v", result, expected)
		}
	})
}
