package main_test

import (
	"bytes"
	"testing"

	main "devtidbits.com/yn"
)

func TestRead(t *testing.T) {
	var stdin bytes.Buffer
	tests := []struct {
		name     string
		keypress string
		want     rune
		wantErr  bool
	}{
		{"empty", "", 0, true},
		{"newline", "\n", rune('\n'), false},
		{"space", " ", rune(' '), false},
		{"upper y", "Y", rune('Y'), false},
		{"lower y", "y", rune('y'), false},
		{"multibyte", "😃", rune('😃'), false},
		{"Yes", "Yes", rune('Y'), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdin.WriteString(tt.keypress)
			got, err := main.Read(&stdin)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Read() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		r       rune
		want    bool
		wantErr bool
	}{
		// incorrect values
		{"empty", 0, false, true},
		{"newline", rune('\n'), false, true},
		{"space", rune(' '), false, true},
		{"multibyte", rune('😃'), false, true},
		// correct values
		{"Yes", rune('Y'), true, false},
		{"yes", rune('y'), true, false},
		{"No", rune('N'), false, false},
		{"no", rune('n'), false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := main.Parse(tt.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
