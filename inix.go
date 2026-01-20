// Package inix is a parser for INI structured data
package inix

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
)

type IniDocument struct {
	sections map[string]map[string]string
}

func (d *IniDocument) GetSection(name string) (*map[string]string, bool) {
	section, ok := d.sections[name]
	if !ok {
		return nil, false
	}

	return &section, true
}

func (d *IniDocument) AddSection(name string, section map[string]string) error {
	_, ok := d.sections[name]
	if ok {
		return errors.New("section already exist")
	}

	d.sections[name] = section

	return nil
}

func (d *IniDocument) DeleteSection(name string) error {
	_, ok := d.sections[name]
	if !ok {
		return errors.New("section not found")
	}

	delete(d.sections, name)
	return nil
}

func (d *IniDocument) GetKey(section, key string) (string, bool) {
	s, ok := d.sections[section]
	if !ok {
		return "", false
	}

	value, ok := s[key]
	return value, ok
}

func (d *IniDocument) SetKey(section, key, value string) error {
	s, ok := d.sections[section]
	if !ok {
		return errors.New("section not found")
	}

	s[key] = value
	return nil
}

func (d *IniDocument) DeleteKey(section, key string) error {
	s, ok := d.sections[section]
	if !ok {
		return errors.New("section not found")
	}

	_, ok = s[key]
	if !ok {
		return errors.New("key not found")
	}

	delete(s, key)

	return nil
}

type ParseError struct {
	LineNumber int
	Message    string
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("error on line %d: %s", e.LineNumber, e.Message)
}

func Parse(content string) (IniDocument, error) {
	var (
		document       IniDocument
		section        map[string]string
		sectionName    string
		writingSection bool
		lineNo         int
	)

	document.sections = make(map[string]map[string]string)

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
		if strings.HasPrefix(line, ";") {
			continue
		}

		// Parse section name
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			if writingSection {
				writingSection = false
				document.sections[sectionName] = section
			}

			sectionName = strings.TrimPrefix(line, "[")
			sectionName = strings.TrimSuffix(sectionName, "]")

			if strings.Contains(sectionName, " ") {
				return document, &ParseError{
					LineNumber: lineNo,
					Message:    "section name cannot use spaces",
				}
			}

			section = make(map[string]string)
			writingSection = true
			continue
		}

		// Parse key and value
		if writingSection {
			left, right, ok := strings.Cut(line, "=")
			fmt.Printf("%s %s\n", left, right)
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
		document.sections[sectionName] = section
	}

	return document, nil
}
