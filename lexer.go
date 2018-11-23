package main

import (
	"bufio"
	"bytes"
	"io"
)

type Token int

const (
	LABEL Token = iota
	COLUMN
	WORD
	WS
	EOF
	ILLEGAL
)

var eof = rune(0)

type TokenType struct {
	Token    Token
	Litteral string
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

type Scanner struct {
	r *bufio.Reader
}

func (s *Scanner) Scan() (token Token, litteral string) {
	ch := s.read()

	if ch == eof {
		return EOF, ""
	} else if ch == ':' {
		return COLUMN, ":"
	} else if isWhitespace(ch) {
		s.unread()
		return s.readWhitespace()
	} else if isLetter(ch) {
		s.unread()
		return s.readText()
	}

	return ILLEGAL, ""
}

func (s *Scanner) read() rune {
	ch, _, _ := s.r.ReadRune()
	return ch
}

func (s *Scanner) unread() { s.r.UnreadRune() }

func (s *Scanner) readWhitespace() (token Token, litteral string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		ch := s.read()
		if !isWhitespace(ch) {
			s.unread()
			return WS, buf.String()
		} else {
			buf.WriteRune(ch)
		}
	}
}

func (s *Scanner) readText() (token Token, litteral string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		ch := s.read()
		if !isLetter(ch) {
			s.unread()
			text := buf.String()
			if text == "TODO" || text == "DONE" {
				return LABEL, text
			} else {
				return WORD, text
			}
		}
		buf.WriteRune(ch)
	}
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}
