package mot

import (
	"testing"
)

func TestAdd(t *testing.T) {
	wm := NewWordModel()
	wm.Add("hello")
}
