package main

import "github.com/madvikinggod/intexplore/sdk/side"

func main() {
	bp := side.NewBlankProvider()
	b := bp.Blank()
	println(b.V1())
}
