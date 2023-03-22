package noop

type BlankProvider interface {
	Blank() Blank
}

type Blank interface {
	V1() string
	V2() string
}

type NoopBlankProvider struct{}

func (NoopBlankProvider) Blank() Blank {
	return NoopBlank{}
}

type NoopBlank struct{}

func (NoopBlank) V1() string {
	return "noop"
}

func (NoopBlank) V2() string {
	return "noop"
}
