// Package inix is a parser for INI structured data
package inix

import (
	"bufio"
	"fmt"
	"strings"
)

type ParseError struct {
	LineNumber int
	Message    string
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("error on line %d: %s", e.LineNumber, e.Message)
}

func Parse(content string) (map[string]map[string]string, error) {
	var (
		document       map[string]map[string]string
		section        map[string]string
		sectionName    string
		writingSection bool
		lineNo         int
	)

	document = make(map[string]map[string]string)

	if len(content) == 0 {
		return document, nil
	}

	scanner := bufio.NewScanner(strings.NewReader(content))

	for scanner.Scan() {
		lineNo++
		rawLine := scanner.Text()
		line := strings.TrimSpace(rawLine)

		if line == "" {
			continue
		}

		// Ignore comment
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}

		// Parse section name
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			if writingSection {
				writingSection = false
				document[sectionName] = section
			}

			sectionName = strings.TrimPrefix(line, "[")
			sectionName = strings.TrimSuffix(sectionName, "]")

			section = make(map[string]string)
			writingSection = true
			continue
		}

		// Parse key and value
		if writingSection {
			left, right, ok := strings.Cut(line, "=")
			if !ok {
				return document, &ParseError{
					LineNumber: lineNo,
					Message:    "invalid syntax",
				}
			}

			key := strings.TrimSpace(left)
			value := strings.TrimSpace(right)

			_, ok = section[key]
			if ok {
				return document, &ParseError{
					LineNumber: lineNo,
					Message:    fmt.Sprintf("key '%s' in section '%s' is already defined", key, sectionName),
				}
			}

			section[key] = value
			continue
		} else {
			return document, &ParseError{
				LineNumber: lineNo,
				Message:    "key should be defined in section",
			}
		}
	}

	// If file ended but still writing section
	if writingSection {
		writingSection = false
		document[sectionName] = section
	}

	return document, nil
}

func Dump(document map[string]map[string]string) string {
	var builder strings.Builder

	for sectionName, section := range document {
		builder.WriteString(fmt.Sprintf("[%s]\n", sectionName))
		for key, value := range section {
			builder.WriteString(fmt.Sprintf("%s=%s\n", key, value))
		}
		builder.WriteString("\n")
	}

	return builder.String()
}
