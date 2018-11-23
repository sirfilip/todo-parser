package main

import (
	"bytes"
	"fmt"
	"io"
)

func NewParser(r io.Reader) *Parser {
	p := &Parser{s: NewScanner(r)}
	p.buf.Unprocessed = false
	return p
}

type Parser struct {
	s   *Scanner
	buf struct {
		Token       Token
		Litteral    string
		Unprocessed bool
	}
}

func (p *Parser) Parse() ([]Todo, error) {
	todos := make([]Todo, 0)

	for {
		token, label := p.next()
		if token != LABEL {
			return nil, fmt.Errorf("Expected label but got: %v", label)
		}
		token, litteral := p.next()
		if token != COLUMN {
			return nil, fmt.Errorf("Expected column but got: %v", litteral)
		}
		token, litteral = p.next()
		if token != WS {
			return nil, fmt.Errorf("Expected whitespace but got: %v", litteral)
		}
		task := p.readWordsWithWhitespace()
		todos = append(todos, p.buildTodo(label, task))

		token, litteral = p.next()
		if token == EOF {
			return todos, nil
		} else {
			p.prev()
		}
	}
}

func (p *Parser) next() (Token, string) {
	if p.buf.Unprocessed == true {
		p.buf.Unprocessed = false
		return p.buf.Token, p.buf.Litteral
	}
	token, litteral := p.s.Scan()
	p.buf.Token = token
	p.buf.Litteral = litteral
	p.buf.Unprocessed = false
	return token, litteral
}

func (p *Parser) prev() {
	p.buf.Unprocessed = true
}

func (p *Parser) readWordsWithWhitespace() string {
	var buf bytes.Buffer
	var ws = ""
	for {
		tok, lit := p.next()
		if tok == WS {
			ws = lit
		} else if tok == WORD {
			buf.WriteString(ws)
			buf.WriteString(lit)
		} else {
			p.prev()
			return buf.String()
		}
	}
}

func (p *Parser) buildTodo(label string, task string) Todo {
	todo := Todo{}
	todo.Task = task
	if label == "TODO" {
		todo.Finished = false
	} else {
		todo.Finished = true
	}
	return todo
}
