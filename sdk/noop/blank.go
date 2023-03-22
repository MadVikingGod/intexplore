package noop

import "github.com/madvikinggod/intexplore/api/noop"

func NewBlankProvider() noop.BlankProvider {
	return &blankProvider{}
}

type blankProvider struct {
	noop.NoopBlankProvider
}

func (p *blankProvider) Blank() noop.Blank {
	return &blank{}
}

type blank struct {
	noop.NoopBlank
}

func (b *blank) V1() string {
	return "v1"
}
