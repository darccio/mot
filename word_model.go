package mot

import (
	"fmt"
)

type node struct {
	counter  int
	children []*node
}

func newNode() *node {
	children := make([]*node, 0)
	return &node{
		counter:  0,
		children: children,
	}
}

type WordModel struct {
	structure []map[rune]*node
	follows   map[rune]map[rune]*node
}

func NewWordModel() *WordModel {
	structure := make([]map[rune]*node, 0)
	follows := make(map[rune]map[rune]*node)
	return &WordModel{
		structure,
		follows,
	}
}

func (wm *WordModel) Add(word string) {
	var previous rune
	if len(word) == 0 {
		return
	}
	wm.grow(len(word))
	for ix, current := range word {
		wm.hit(wm.structure[ix], current)
		if previous != '\000' {
			wm.hit(wm.follows[previous], current)
		}
		if _, ok := wm.follows[current]; !ok {
			wm.follows[current] = make(map[rune]*node)
		}
		previous = current
	}
	wm.hit(wm.follows[previous], '\000')
	fmt.Printf("%v\n", wm.structure)
	fmt.Printf("%v\n", wm.follows)
}

func (wm *WordModel) grow(target int) {
	required := target - len(wm.structure)
	if required < 1 {
		return
	}
	moreRoom := make([]map[rune]*node, required)
	for ix := range moreRoom {
		moreRoom[ix] = make(map[rune]*node)
	}
	wm.structure = append(wm.structure, moreRoom...)
}

func (wm *WordModel) hit(data map[rune]*node, target rune) {
	if _, ok := data[target]; !ok {
		data[target] = newNode()
	}
	data[target].counter++
}
