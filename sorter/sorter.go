package sorter

import (
	"sort"
	"strings"
)

type Name struct {
	First string
	Last  string
	Full  string
}

type Sorter interface {
	Sort(names []string, desc bool) []string
}

type SortingStrategy interface {
	Sort(names []Name, desc bool) []Name
}

type LastNameSorting struct{}

func (l *LastNameSorting) Sort(names []Name, desc bool) []Name {
	sort.Slice(names, func(i, j int) bool {
		if desc {
			return names[i].Last > names[j].Last
		}
		return names[i].Last < names[j].Last
	})
	return names
}

type FirstNameSorting struct{}

func (f *FirstNameSorting) Sort(names []Name, desc bool) []Name {
	sort.Slice(names, func(i, j int) bool {
		if desc {
			return names[i].First > names[j].First
		}
		return names[i].First < names[j].First
	})
	return names
}

type NameSorter struct {
	Strategy SortingStrategy
}

func NewNameSorter(strategy SortingStrategy) Sorter {
	return &NameSorter{Strategy: strategy}
}

func (s *NameSorter) Sort(names []string, desc bool) []string {
	parsedNames := parseNames(names)
	sortedNames := s.Strategy.Sort(parsedNames, desc)
	return namesFromStructs(sortedNames)
}

func parseNames(names []string) []Name {
	var parsed []Name
	for _, name := range names {
		parts := strings.Fields(name)
		last := parts[len(parts)-1]
		first := strings.Join(parts[:len(parts)-1], " ")
		parsed = append(parsed, Name{First: first, Last: last, Full: name})
	}
	return parsed
}

func namesFromStructs(names []Name) []string {
	sorted := make([]string, len(names))
	for i, name := range names {
		sorted[i] = name.Full
	}
	return sorted
}
