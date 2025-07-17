package main

import "fmt"

type TrieNode struct {
    children    map[rune]*TrieNode
    isEndOfWord bool
}

type Trie struct {
    root *TrieNode
}

func NewTrie() *Trie {
    return &Trie{
        root: &TrieNode{children: make(map[rune]*TrieNode)},
    }
}

func (t *Trie) Insert(word string) {
    current := t.root
    for _, char := range word {
        if _, exists := current.children[char]; !exists {
            current.children[char] = &TrieNode{children: make(map[rune]*TrieNode)}
        }
        current = current.children[char]
    }
    current.isEndOfWord = true
}

func (t *Trie) Search(word string) bool {
    current := t.root
    for _, char := range word {
        if _, exists := current.children[char]; !exists {
            return false
        }
        current = current.children[char]
    }
    return current.isEndOfWord
}

func (t *Trie) Autocomplete(prefix string) []string {
    current := t.root
    for _, char := range prefix {
        if _, exists := current.children[char]; !exists {
            return []string{}
        }
        current = current.children[char]
    }

    results := []string{}
    t.collectWords(current, prefix, &results)
    return results
}

func (t *Trie) collectWords(node *TrieNode, prefix string, results *[]string) {
    if node.isEndOfWord {
        *results = append(*results, prefix)
    }
    for char, child := range node.children {
        t.collectWords(child, prefix+string(char), results)
    }
}

func main() {
    trie := NewTrie()
    kata := []string{"go", "golang", "game", "gopher", "gopay", "gold", "gone"}
    for _, word := range kata {
        trie.Insert(word)
    }

    fmt.Println("Autocomplete untuk 'go':")
    saran := trie.Autocomplete("go")
    for _, s := range saran {
        fmt.Println(" -", s)
    }
}