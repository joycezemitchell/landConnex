package sorter

import (
	"reflect"
	"testing"
)

func TestNameParser_Parse(t *testing.T) {
	parser := NewNameParser()

	tests := []struct {
		name  string
		input []string
		want  []Name
	}{
		{
			name:  "normal use case",
			input: []string{"John Doe", "Jane Smith"},
			want: []Name{
				{First: "John", Last: "Doe", Full: "John Doe"},
				{First: "Jane", Last: "Smith", Full: "Jane Smith"},
			},
		},
		{
			name:  "single name",
			input: []string{"Cher"},
			want: []Name{
				{First: "", Last: "Cher", Full: "Cher"},
			},
		},
		{
			name:  "multiple spaces",
			input: []string{"John   Doe", "  Jane   Smith  "},
			want: []Name{
				{First: "John", Last: "Doe", Full: "John   Doe"},
				{First: "Jane", Last: "Smith", Full: "  Jane   Smith  "},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parser.Parse(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NameParser.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
