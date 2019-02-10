package problem11

import (
	"strings"
)

func autocomplete(query string, set []string) []string {
	out := make([]string, 0)
	query = strings.ToLower(query)
	for _, s := range set {
		if strings.HasPrefix(strings.ToLower(s), query) {
			out = append(out, s)
		}
	}
	return out
}

func autocompleteRadix(query string, set []string) []string {
	root := newNode()
	for _, w := range set {
		insert(root, strings.ToLower(w))
	}
	return printAutoSuggestions(root, strings.ToLower(query))
}

const alphabetSize = 26

func charToIndex(r rune) int {
	i := int(r)
	if i >= 97 && i <= 122 {
		return i - 97
	}
	return -1
}

type node struct {
	children  []*node
	isWordEnd bool
}

func newNode() *node {
	n := &node{
		children:  make([]*node, alphabetSize),
		isWordEnd: false,
	}
	for i := 0; i < alphabetSize; i++ {
		n.children[i] = nil
	}
	return n
}

func insert(root *node, key string) {
	crawl := root
	for _, r := range key {
		index := charToIndex(r)
		if crawl.children[index] == nil {
			crawl.children[index] = newNode()
		}
		crawl = crawl.children[index]
	}
	crawl.isWordEnd = true
}

func search(root *node, key string) bool {
	crawl := root
	for _, r := range key {
		index := charToIndex(r)
		if crawl.children[index] == nil {
			return false
		}
		crawl = crawl.children[index]
	}
	return crawl != nil && crawl.isWordEnd
}

func isLastNode(root *node) bool {
	for i := 0; i < alphabetSize; i++ {
		if root.children[i] != nil {
			return false
		}
	}
	return true
}

func suggestionsRec(root *node, prefix string) []string {
	out := make([]string, 0)
	if root.isWordEnd {
		return append(out, prefix)
	}
	if isLastNode(root) {
		return out
	}
	for i := 0; i < alphabetSize; i++ {
		if root.children[i] != nil {
			newPrefix := prefix + string(97+i)
			out = append(suggestionsRec(root.children[i], newPrefix), out...)
		}
	}
	return out
}

func printAutoSuggestions(root *node, query string) []string {
	out := make([]string, 0)
	crawl := root
	for _, r := range query {
		index := charToIndex(r)
		if crawl.children[index] == nil {
			return out
		}
		crawl = crawl.children[index]
	}

	isWord := crawl.isWordEnd == true
	isLast := isLastNode(crawl)

	if isWord && isLast {
		return nil
	}

	if !isLast {
		return suggestionsRec(crawl, query)
	}
	return out
}
