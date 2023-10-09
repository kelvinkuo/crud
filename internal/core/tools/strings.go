package tools

import (
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// supporting convert between lower camel case and upper camel case and underline style
// example: userName UserName user_name User_Name

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
