package main

import "github.com/madvikinggod/intexplore/sdk/abstract"

func main() {
	bp := abstract.NewBlankProvider()
	b := bp.Blank()
	println(b.V1())
}
