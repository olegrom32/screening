package main

import (
	"log"
	"math/rand"
)

// Generator defines the API for generation of random strings
type Generator interface {
	Generate(cfg Config) string
}

// Config mimics some well-known password generators
type Config struct {
	Length    int
	Uppercase bool
	Lowercase bool
	Numbers   bool
	Symbols   bool
}

func (c *Config) Valid() bool {
	return c.Length > 0 && (c.Uppercase || c.Lowercase || c.Numbers || c.Symbols)
}

// UnseededRandGenerator is an implementation of Generator. Since the task description states that you
// "are mostly interested in the interface and less interested in the implementation" - this is the
// simplest implementation I could think of.
type UnseededRandGenerator struct {
}

const (
	charsetLower   = "abcdefghijklmnopqrstuvwxyz"
	charsetUpper   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charsetNumbers = "0123456789"
	charsetSymbols = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
)

func (g *UnseededRandGenerator) Generate(cfg Config) string {
	if !cfg.Valid() {
		// No errors to keep the solution concise
		return ""
	}

	var charset string

	if cfg.Uppercase {
		charset = charset + charsetUpper
	}
	if cfg.Lowercase {
		charset = charset + charsetLower
	}
	if cfg.Numbers {
		charset = charset + charsetNumbers
	}
	if cfg.Symbols {
		charset = charset + charsetSymbols
	}

	b := make([]byte, cfg.Length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

func main() {
	g := UnseededRandGenerator{}

	log.Print(g.Generate(Config{
		Length:    24,
		Uppercase: true,
		Lowercase: true,
		Numbers:   false,
		Symbols:   true,
	}))
}
