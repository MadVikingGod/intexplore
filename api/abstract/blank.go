package abstract

type BlankProvider interface {
	blankProvider()

	Blank() Blank
}

type Blank interface {
	blank()
	V1() string
	V2() string
}
