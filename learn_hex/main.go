package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	b, _ := hex.DecodeString("7b2254797065223a22736563703235366b31222c22507269766174654b6579223a226f71545a446c2b6f6c5145466f705264796a785551424e63733861695a305163383145434e667a433856493d227d")
	fmt.Println(string(b))
}
