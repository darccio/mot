package mot

import (
	"fmt"
	"testing"
)

func TestNewWordModel(t *testing.T) {
	wm := NewWordModel()
	if wm.words.get(-1, zerune) == nil {
		t.Errorf("root not found in wm.words")
	}
	if wm.digraphs.get(-1, zerune) == nil {
		t.Errorf("root not found in wm.digraphs")
	}
}

func TestAdd(t *testing.T) {
	wm := NewWordModel()
	wm.Add("hello")
	wm.Add("hello")
	var root *node
	root = wm.words.get(-1, zerune)
	if root == nil {
		t.Errorf("root not found in wm.words")
	}
	root = wm.digraphs.get(-1, zerune)
	if root == nil {
		t.Errorf("root not found in wm.words")
	}
	// TODO treat digraph different, we should count only how many times after parent a rune appears.
}
