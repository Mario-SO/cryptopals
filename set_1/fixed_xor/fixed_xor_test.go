package fixedxor

import (
	"testing"
)

func TestFixedXOR(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{
			name:     "Cryptopals Challenge 2",
			input:    "1c0111001f010100061a024b53535009181c",
			expected: "746865206b696420646f6e277420706c6179",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FixedXOR(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("fixedXOR() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("fixedXOR() = %v, want %v", got, tt.expected)
			}
		})
	}
}
