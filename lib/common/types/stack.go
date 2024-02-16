package common_types

// Taken from: https://www.educative.io/answers/how-to-implement-a-stack-in-golang

type Stack[T any] []T

// IsEmpty: check if stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack[T]) Push(str T) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	index := 0             // Get the index of the topmost element.
	element := (*s)[index] // Index into the slice and obtain the element.
	*s = (*s)[1:]          // Remove it from the stack by slicing it off.
	return element, true
}
