package mot

type WordModel struct {
	structure []map[rune]int
	follows   map[rune]map[rune]int
}

func NewWordModel() *WordModel {
	structure := make([]map[rune]int, 0)
	follows := make(map[rune]map[rune]int)
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
			wm.follows[current] = make(map[rune]int)
		}
		previous = current
	}
	wm.hit(wm.follows[previous], '\000')
}

func (wm *WordModel) grow(target int) {
	required := target - len(wm.structure)
	if required < 1 {
		return
	}
	moreRoom := make([]map[rune]int, required)
	for ix := range moreRoom {
		moreRoom[ix] = make(map[rune]int)
	}
	wm.structure = append(wm.structure, moreRoom...)
}

func (wm *WordModel) hit(data map[rune]int, target rune) {
	if _, ok := data[target]; ok {
		data[target] = 0
	}
	data[target]++
}
