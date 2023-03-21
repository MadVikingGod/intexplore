package main

import (
	"github.com/madvikinggod/intexplore/api/side"
	sdk "github.com/madvikinggod/intexplore/sdk/side"
)

func main() {
	bp := sdk.NewBlankProvider()
	b := bp.Blank()
	println(b.V1())
	if b, ok := b.(side.BlankV2); ok {
		println(b.V2())
	} else {
		println("blank is not V2")
	}
}
