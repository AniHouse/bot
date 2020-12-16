package dstype

import (
	"errors"

	"github.com/rivo/uniseg"
)

type Grapheme struct {
	Value string
}

func (g *Grapheme) Scan(s string, i *int) error {
	gr := uniseg.NewGraphemes(s[*i:])
	if !gr.Next() {
		return errors.New("grapheme parse error")
	}

	g.Value = gr.Str()
	_, *i = gr.Positions()
	return nil
}
