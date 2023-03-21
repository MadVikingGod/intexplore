package side

type BlankProvider interface {
	Blank() Blank
}

type Blank interface {
	V1() string
}
