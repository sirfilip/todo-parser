package main

import (
	"strings"
	"testing"
)

func TestValidText(t *testing.T) {
	input := strings.NewReader("TODO: Write some tests\nDONE: Write the implementation")
	s := NewScanner(input)
	tokens := []TokenType{
		TokenType{LABEL, "TODO"},
		TokenType{COLUMN, ":"},
		TokenType{WS, " "},
		TokenType{WORD, "Write"},
		TokenType{WS, " "},
		TokenType{WORD, "some"},
		TokenType{WS, " "},
		TokenType{WORD, "tests"},
		TokenType{WS, "\n"},
		TokenType{LABEL, "DONE"},
		TokenType{COLUMN, ":"},
		TokenType{WS, " "},
		TokenType{WORD, "Write"},
		TokenType{WS, " "},
		TokenType{WORD, "the"},
		TokenType{WS, " "},
		TokenType{WORD, "implementation"},
	}

	for _, tt := range tokens {
		token, litteral := s.Scan()
		if token != tt.Token {
			t.Errorf("Expected: %v but got: %v", tt.Token, token)
		}
		if litteral != tt.Litteral {
			t.Errorf("Expected: %v but got %v", tt.Litteral, litteral)
		}
	}
	ch, _ := s.Scan()
	if ch != EOF {
		t.Errorf("Expected eof but got: %v", ch)
	}
}
