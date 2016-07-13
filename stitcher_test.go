package lieber

import (
	"github.com/Enriikke/lieber/marvel"
	"testing"
)

func TestStichingTwo(t *testing.T) {
	stories := []marvel.Story{
		marvel.Story{"Story1", "X-Men do stuff"},
		marvel.Story{"Story2", "Avengers do stuff"},
	}

	stitched := stitch(stories)
	expected := "X-Men do stuff, then\nAvengers do stuff.\nTHE END"
	if stitched != expected {
		t.Errorf("Wrong stitch: got %v, wanted %v", stitched, expected)
	}
}
