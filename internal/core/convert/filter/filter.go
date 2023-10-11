package filter

import (
	"github.com/kelvinkuo/crud/internal/db"
)

type StringFilter struct {
	words map[string]bool
}

func NewStringFilter(words []string) *StringFilter {
	wordMap := make(map[string]bool, 0)
	for _, word := range words {
		wordMap[word] = true
	}
	return &StringFilter{words: wordMap}
}

func (f *StringFilter) FilterOut(col db.Column) bool {
	if _, ok := f.words[col.Name()]; ok {
		return true
	}
	return false
}

func (f *StringFilter) AddWord(word string) {
	f.words[word] = true
}
