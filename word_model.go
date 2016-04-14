package mot

var (
	zerune = '\000'
)

type key struct {
	position int
	value    rune
}

type node struct {
	value    rune
	counter  int
	children map[rune]*node
}

func newNode(value rune) *node {
	return &node{
		value,
		0,
		make(map[rune]*node),
	}
}

type chain struct {
	data map[key]*node
}

func newChain() (c *chain) {
	c = &chain{
		make(map[key]*node),
	}
	root := c.put(-1, zerune)
	root.counter++
	return
}

func (c *chain) get(position int, value rune) *node {
	return c.data[key{position, value}]
}

func (c *chain) put(position int, value rune) *node {
	k := key{position, value}
	if _, ok := c.data[k]; !ok {
		c.data[k] = newNode(value)
	}
	return c.data[k]
}

type WordModel struct {
	words    *chain
	digraphs *chain
}

func NewWordModel() *WordModel {
	return &WordModel{
		newChain(),
		newChain(),
	}
}

func (wm *WordModel) Add(word string) {
	var parent rune
	if len(word) == 0 {
		return
	}
	for ix, current := range word {
		wm.updateWords(ix, current, parent)
		wm.updateDigraphs(ix, current, parent)
		parent = current
	}
}

func (wm *WordModel) updateWords(position int, current, parent rune) {
	c := wm.words
	n := c.put(position, current)
	n.counter++
	p := c.get(position-1, parent)
	p.children[current] = n
}

func (wm *WordModel) updateDigraphs(position int, current, parent rune) {
	c := wm.digraphs
	c.put(0, current)
	parentPosition := 0
	if position == 0 {
		parentPosition = -1
	}
	p := c.get(parentPosition, parent)
	if _, ok := p.children[current]; !ok {
		p.children[current] = newNode(current)
	}
	p.children[current].counter++
}
