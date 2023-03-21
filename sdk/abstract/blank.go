package abstract

import "github.com/madvikinggod/intexplore/api/abstract"

func NewBlankProvider() abstract.BlankProvider {
	return &blankProvider{}
}

type blankProvider struct {
	abstract.BlankProvider
}

func (p *blankProvider) Blank() abstract.Blank {
	return &blank{}
}

type blank struct {
	abstract.Blank
}

func (b *blank) V1() string {
	return "v1"
}
