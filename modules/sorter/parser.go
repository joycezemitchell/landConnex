package sorter

import "strings"

type Name struct {
	First string
	Last  string
	Full  string
}

type NameParser struct{}

func NewNameParser() *NameParser {
	return &NameParser{}
}

func (np *NameParser) Parse(names []string) []Name {
	var parsed []Name
	for _, name := range names {
		parts := strings.Fields(name)
		if len(parts) == 0 {
			continue
		}
		last := parts[len(parts)-1]
		first := ""
		if len(parts) > 1 {
			first = strings.Join(parts[:len(parts)-1], " ")
		}
		parsed = append(parsed, Name{First: first, Last: last, Full: name})
	}
	return parsed
}
