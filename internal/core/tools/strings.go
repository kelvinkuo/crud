package tools

import (
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// LowerCamelCase return string with format "camelCase"
func LowerCamelCase(str string) string {
	titleCase := cases.Title(language.English)
	lowerCase := cases.Lower(language.English)

	words := split(str)
	if len(words) == 0 {
		return ""
	}

	words[0] = lowerCase.String(words[0])
	for i := 1; i < len(words); i++ {
		words[i] = titleCase.String(words[i])
	}

	return strings.Join(words, "")
}

// UpperCamelCase return string with format "CamelCase"
func UpperCamelCase(str string) string {
	titleCase := cases.Title(language.English)

	words := split(str)
	if len(words) == 0 {
		return ""
	}

	for i := 0; i < len(words); i++ {
		words[i] = titleCase.String(words[i])
	}

	return strings.Join(words, "")
}

// LowerUnderline return string with format "camel_case"
func LowerUnderline(str string) string {
	lowerCase := cases.Lower(language.English)

	words := split(str)
	if len(words) == 0 {
		return ""
	}

	for i := 0; i < len(words); i++ {
		words[i] = lowerCase.String(words[i])
	}

	return strings.Join(words, "_")
}

func split(str string) []string {
	sep := '='
	origin := []rune(str)
	result := make([]rune, 0)

	for i, r := range origin {
		if i != 0 && (unicode.IsUpper(r) || r == '_') && result[len(result)-1] != sep {
			result = append(result, sep)
		}

		if unicode.IsLetter(r) {
			result = append(result, r)
		}
	}

	return strings.Split(string(result), string(sep))
}

// Blank return string consists of n whitespace
func Blank(n int) string {
	builder := strings.Builder{}
	for i := 0; i < n; i++ {
		builder.WriteString(" ")
	}

	return builder.String()
}
