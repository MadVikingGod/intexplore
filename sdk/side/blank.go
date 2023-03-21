package side

import "github.com/madvikinggod/intexplore/api/side"

func NewBlankProvider() side.BlankProvider {
	return &blankProvider{}
}

type blankProvider struct{}

func (p *blankProvider) Blank() side.Blank {
	return &blank{}
}

type blank struct{}

func (b *blank) V1() string {
	return "v1"
}

func (b *blank) V2() string {
	return "v2"
}
