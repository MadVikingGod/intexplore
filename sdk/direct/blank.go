package direct

import "github.com/madvikinggod/intexplore/api/direct"

func NewBlankProvider() direct.BlankProvider {
	return &blankProvider{}
}

type blankProvider struct{}

func (p *blankProvider) Blank() direct.Blank {
	return &blank{}
}

type blank struct{}

func (b *blank) V1() string {
	return "v1"
}
