package abstract

import "github.com/madvikinggod/intexplore/api/abstract"

func NewBlankProvider() abstract.BlankProvider {
	return &blankProvider{}
}

type blankProvider struct{}

func (p *blankProvider) Blank() abstract.Blank {
	return &blank{}
}

type blank struct{}

func (b *blank) V1() string {
	return "v1"
}
