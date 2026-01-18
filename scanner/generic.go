package scanner

import (
	"fmt"
	"strconv"
	"strings"
)

func Read[T InputType](s *Scanner, prompt string) T {
	for {
		value := s.readLine(prompt)

		var result any
		var err error

		switch any(*new(T)).(type) {
		case int:
			result, err = strconv.Atoi(value)

		case int64:
			result, err = strconv.ParseInt(value, 10, 64)

		case float64:
			result, err = strconv.ParseFloat(value, 64)

		case bool:
			switch strings.ToLower(value) {
			case "true", "yes", "y", "1":
				result = true
			case "false", "no", "n", "0":
				result = false
			default:
				err = fmt.Errorf("invalid boolean")
			}

		case string:
			result = value

		default:
			panic("unsupported type")
		}

		if err == nil {
			return result.(T)
		}

		fmt.Println("❌ Invalid input, try again!")
	}
}
