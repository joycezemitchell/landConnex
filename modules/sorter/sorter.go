package sorter

import "sort"

type NameSorter struct {
	Strategy SortingStrategy
	Parser   *NameParser
}

type Sorter interface {
	Sort(names []string, desc bool) []string
}

type SortingStrategy interface {
	Sort(names []Name, desc bool) []Name
}

type LastNameSorting struct{}

func NewLastNameSorting() SortingStrategy {
	return &LastNameSorting{}
}

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

func NewFirstNameSorting() SortingStrategy {
	return &FirstNameSorting{}
}

func (f *FirstNameSorting) Sort(names []Name, desc bool) []Name {
	sort.Slice(names, func(i, j int) bool {
		if desc {
			return names[i].First > names[j].First
		}
		return names[i].First < names[j].First
	})
	return names
}

func NewNameSorter(strategy SortingStrategy, parser *NameParser) Sorter {
	return &NameSorter{Strategy: strategy, Parser: parser}
}

func (s *NameSorter) Sort(names []string, desc bool) []string {
	parsedNames := s.Parser.Parse(names)
	sortedNames := s.Strategy.Sort(parsedNames, desc)
	return namesFromStructs(sortedNames)
}

func namesFromStructs(names []Name) []string {
	sorted := make([]string, len(names))
	for i, name := range names {
		sorted[i] = name.Full
	}
	return sorted
}
