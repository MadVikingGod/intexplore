package main

import "github.com/madvikinggod/intexplore/sdk/direct"

func main() {
	bp := direct.NewBlankProvider()
	b := bp.Blank()
	println(b.V1())
}
