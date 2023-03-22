package main

import "github.com/madvikinggod/intexplore/sdk/noop"

func main() {
	bp := noop.NewBlankProvider()
	b := bp.Blank()
	println(b.V1())
	println(b.V2())
}
