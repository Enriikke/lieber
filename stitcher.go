package lieber

import (
	"bytes"
	"github.com/Enriikke/lieber/marvel"
)

func stitch(stories []marvel.Story) string {

	var text bytes.Buffer
	for ind, story := range stories {
		text.WriteString(story.Description)
		if len(stories) > ind+1 {
			text.WriteString(", then")
			text.WriteString("\n")
		}
	}
	text.WriteString(".\n")
	text.WriteString("THE END")
	return text.String()
}
