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

type parentLocator func(int) int

type chain struct {
	data map[key]*node
}

func newChain() (c *chain) {
	c = &chain{
		make(map[key]*node),
	}
	c.put(-1, zerune)
	return
}

func (c *chain) get(position int, value rune) (n *node) {
	return c.data[key{position, value}]
}

func (c *chain) put(position int, value rune) (n *node) {
	k := key{position, value}
	if _, ok := c.data[k]; !ok {
		c.data[k] = &node{
			value,
			0,
			make(map[rune]*node),
		}
	}
	n = c.data[k]
	n.counter++
	return
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
	p := c.get(position-1, parent)
	p.children[current] = n
}

func (wm *WordModel) updateDigraphs(position int, current, parent rune) {
	c := wm.digraphs
	n := c.put(0, current)
	parentPosition := 0
	if position == 0 {
		parentPosition = -1
	}
	p := c.get(parentPosition, parent)
	p.children[current] = n
}
