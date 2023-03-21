package side

type BlankProvider interface {
	Blank() Blank
}

type Blank interface {
	V1() string
}

type BlankV2 interface {
	V2() string
}
