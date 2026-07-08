package stringutil

import "testing"

func TestUnquote(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "Original string is returned when there are no quotes",
			s:    "not quoted",
			want: "not quoted",
		},
		{
			name: "Single quotes",
			s:    "'I am quoted'",
			want: "I am quoted",
		},
		{
			name: "Double quotes",
			s:    `"I am quoted"`,
			want: "I am quoted",
		},
		{
			name: "Backtick quotes",
			s:    "`I am quoted`",
			want: "I am quoted",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unquote(tt.s); got != tt.want {
				t.Fatalf(`Unquote() = %q, want %q`, got, tt.want)
			}
		})
	}
}
