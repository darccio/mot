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
	for _, v := range wm.digraphs.data[key{0, 'l'}].children {
		fmt.Printf("%c: %d\n", v.value, v.counter)
	}
	// TODO check if both childrens have 2 hits in their counters
}
