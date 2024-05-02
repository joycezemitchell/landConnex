package sorter

import (
	"reflect"
	"testing"
)

func TestLastNameSorting(t *testing.T) {
	names := []Name{
		{First: "John", Last: "Doe", Full: "John Doe"},
		{First: "Jane", Last: "Doe", Full: "Jane Doe"},
		{First: "Alice", Last: "Smith", Full: "Alice Smith"},
	}
	lastNameSorting := LastNameSorting{}
	sorted := lastNameSorting.Sort(names, false)
	expected := []Name{
		{First: "John", Last: "Doe", Full: "John Doe"},
		{First: "Jane", Last: "Doe", Full: "Jane Doe"},
		{First: "Alice", Last: "Smith", Full: "Alice Smith"},
	}
	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Expected %v, got %v", expected, sorted)
	}
}

func TestFirstNameSorting(t *testing.T) {
	names := []Name{
		{First: "Alice", Last: "Smith", Full: "Alice Smith"},
		{First: "Jane", Last: "Doe", Full: "Jane Doe"},
		{First: "John", Last: "Doe", Full: "John Doe"},
	}
	firstNameSorting := FirstNameSorting{}
	sorted := firstNameSorting.Sort(names, false)
	expected := []Name{
		{First: "Alice", Last: "Smith", Full: "Alice Smith"},
		{First: "Jane", Last: "Doe", Full: "Jane Doe"},
		{First: "John", Last: "Doe", Full: "John Doe"},
	}
	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Expected %v, got %v", expected, sorted)
	}
}
