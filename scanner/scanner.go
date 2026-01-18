package scanner

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Scanner struct {
	reader *bufio.Reader
}

func New() *Scanner {
	return &Scanner{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (s *Scanner) readLine(prompt string) string {
	if prompt != "" {
		fmt.Print(prompt)
	}

	input, err := s.reader.ReadString('\n')
	if err != nil {
		panic("failed to read input")
	}

	return strings.TrimSpace(input)
}

func (s *Scanner) String(prompt string) string {
	return s.readLine(prompt)
}

func (s *Scanner) Int(prompt string) int {
	for {
		value := s.readLine(prompt)
		if i, err := strconv.Atoi(value); err == nil {
			return i
		}
		fmt.Println("❌ Invalid integer, try again")
	}
}

func (s *Scanner) Float(prompt string) float64 {
	for {
		value := s.readLine(prompt)
		if f, err := strconv.ParseFloat(value, 64); err == nil {
			return f
		}
		fmt.Println("❌ Invalid float, try again")
	}
}

func (s *Scanner) Bool(prompt string) bool {
	for {
		switch strings.ToLower(s.readLine(prompt)) {
		case "true", "yes", "y", "1":
			return true
		case "false", "no", "n", "0":
			return false
		default:
			fmt.Println("❌ Enter true/false, yes/no")
		}
	}
}
