package common

import (
	"bytes"
	"fmt"
)

const letterWidth = 5

type AsciiOutput struct {
	letters []uint64
	pos     int
	width   int
}

func NewAsciiOutput(width int) *AsciiOutput {
	return &AsciiOutput{
		letters: make([]uint64, width),
		width:   width * letterWidth,
	}
}

func (a *AsciiOutput) Write(on bool) {
	if on {
		currentLine := a.pos / a.width
		currentLetter := (a.pos % a.width) / letterWidth
		value := currentLine*letterWidth + (a.pos % a.width % letterWidth)
		a.letters[currentLetter] |= 1 << value
	}
	a.pos++
}

var letters = map[uint64]byte{
	311928102: 'A',
	244620583: 'B',
	210797862: 'C',
	244622631: 'D',
	504405039: 'E',
	34642991:  'F',
	479626534: 'G',
	311737641: 'H',
	474091662: 'I',
	211034380: 'J',
	307399849: 'K',
	504398881: 'L',
	311737833: 'M',
	311735657: 'N',
	211068198: 'O',
	34841895:  'P',
	341091622: 'Q',
	307471655: 'R',
	243467310: 'S',
	138547359: 'T',
	211068201: 'U',
	145049137: 'V',
	318219561: 'W',
	581046609: 'X',
	138553905: 'Y',
	504434959: 'Z',
	0:         ' ',
}

func (a *AsciiOutput) Read() string {
	out := &bytes.Buffer{}
	for i := range a.letters {
		v, ok := letters[a.letters[i]]
		if !ok {
			panic(fmt.Sprintf("Unknown letter at index %d: %d", i, a.letters[i]))
		}
		out.WriteByte(v)
	}
	return out.String()
}
