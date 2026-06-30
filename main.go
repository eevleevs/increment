package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	decRE = regexp.MustCompile(`^\d+$`)
	hexRE = regexp.MustCompile(`^0x[\da-fA-F]+$`)
	binRE = regexp.MustCompile(`^0b[01]+$`)
	octRE = regexp.MustCompile(`^0o[0-7]+$`)
)

var toggles = map[string]string{
	"true": "false", "false": "true",
	"True": "False", "False": "True",
	"TRUE": "FALSE", "FALSE": "TRUE",
	"yes": "no", "no": "yes",
	"Yes": "No", "No": "Yes",
	"YES": "NO", "NO": "YES",
	"on": "off", "off": "on",
	"On": "Off", "Off": "On",
	"ON": "OFF", "OFF": "ON",
	"enable": "disable", "disable": "enable",
	"enabled": "disabled", "disabled": "enabled",
}

func Run(input string, amount int) string {
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return input
	}

	hasSign := strings.HasPrefix(input, "-")
	absInput := input
	if hasSign {
		absInput = input[1:]
	}

	var format string
	switch {
	case decRE.MatchString(absInput):
		format = "dec"
	case hexRE.MatchString(absInput):
		format = "hex"
	case binRE.MatchString(absInput):
		format = "bin"
	case octRE.MatchString(absInput):
		format = "oct"
	}

	if format != "" {
		val, err := strconv.ParseInt(absInput, 0, 64)
		if err != nil {
			return input
		}
		signedVal := val
		if hasSign {
			signedVal = -val
		}
		result := signedVal + int64(amount)
		absResult := result
		negOut := false
		if result < 0 {
			absResult = -result
			negOut = true
		}
		prefix := ""
		if negOut {
			prefix = "-"
		}

		switch format {
		case "dec":
			return fmt.Sprintf("%d", result)
		case "hex":
			return fmt.Sprintf("%s0x%X", prefix, absResult)
		case "bin":
			return fmt.Sprintf("%s0b%b", prefix, absResult)
		case "oct":
			return fmt.Sprintf("%s0o%o", prefix, absResult)
		}
	}

	if amount%2 != 0 {
		if val, ok := toggles[input]; ok {
			return val
		}
	}

	return input
}

func main() {
	amount := 1
	if len(os.Args) > 1 {
		n, err := strconv.Atoi(os.Args[1])
		if err == nil {
			amount = n
		}
	}

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		os.Exit(1)
	}

	os.Stdout.WriteString(Run(string(data), amount))
}
