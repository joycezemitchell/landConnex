package sorter

import (
	"reflect"
	"testing"
)

func sampleNames() []Name {
	return []Name{
		{First: "John", Last: "Doe", Full: "John Doe"},
		{First: "Jane", Last: "Smith", Full: "Jane Smith"},
		{First: "Alice", Last: "Brown", Full: "Alice Brown"},
	}
}

func TestLastNameSorting(t *testing.T) {
	names := sampleNames()
	sorter := NewLastNameSorting()
	sortedAsc := sorter.Sort(names, false)
	expectedAsc := []Name{
		{First: "Alice", Last: "Brown", Full: "Alice Brown"},
		{First: "John", Last: "Doe", Full: "John Doe"},
		{First: "Jane", Last: "Smith", Full: "Jane Smith"},
	}
	if !reflect.DeepEqual(sortedAsc, expectedAsc) {
		t.Errorf("Ascending LastNameSorting failed, got %v, want %v", sortedAsc, expectedAsc)
	}

	sortedDesc := sorter.Sort(names, true)
	expectedDesc := []Name{
		{First: "Jane", Last: "Smith", Full: "Jane Smith"},
		{First: "John", Last: "Doe", Full: "John Doe"},
		{First: "Alice", Last: "Brown", Full: "Alice Brown"},
	}
	if !reflect.DeepEqual(sortedDesc, expectedDesc) {
		t.Errorf("Descending LastNameSorting failed, got %v, want %v", sortedDesc, expectedDesc)
	}
}

func TestFirstNameSorting(t *testing.T) {
	names := sampleNames()
	sorter := NewFirstNameSorting()
	sortedAsc := sorter.Sort(names, false)
	expectedAsc := []Name{
		{First: "Alice", Last: "Brown", Full: "Alice Brown"},
		{First: "Jane", Last: "Smith", Full: "Jane Smith"},
		{First: "John", Last: "Doe", Full: "John Doe"},
	}
	if !reflect.DeepEqual(sortedAsc, expectedAsc) {
		t.Errorf("Ascending FirstNameSorting failed, got %v, want %v", sortedAsc, expectedAsc)
	}

	sortedDesc := sorter.Sort(names, true)
	expectedDesc := []Name{
		{First: "John", Last: "Doe", Full: "John Doe"},
		{First: "Jane", Last: "Smith", Full: "Jane Smith"},
		{First: "Alice", Last: "Brown", Full: "Alice Brown"},
	}
	if !reflect.DeepEqual(sortedDesc, expectedDesc) {
		t.Errorf("Descending FirstNameSorting failed, got %v, want %v", sortedDesc, expectedDesc)
	}
}

func TestNameSorter(t *testing.T) {
	names := []string{"John Doe", "Jane Smith", "Alice Brown"}
	parser := NewNameParser()
	strategy := NewFirstNameSorting()
	sorter := NewNameSorter(strategy, parser)

	sortedNames := sorter.Sort(names, false)
	expected := []string{"Alice Brown", "Jane Smith", "John Doe"}
	if !reflect.DeepEqual(sortedNames, expected) {
		t.Errorf("NameSorter failed, got %v, want %v", sortedNames, expected)
	}
}
