package noop

import "github.com/madvikinggod/intexplore/api/noop"

func NewBlankProvider() noop.BlankProvider {
	return &blankProvider{}
}

type blankProvider struct {
	noop.BlankProvider
}

func (p *blankProvider) Blank() noop.Blank {
	return &blank{}
}

type blank struct {
	noop.Blank
}

func (b *blank) V1() string {
	return "v1"
}
