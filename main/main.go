package main

import (
	"github.com/chawat/CID"
	"fmt"
)

func main() {
	Test := CID.OpenImg("testDigits/8_12.txt")
	fmt.Println(CID.KNNClassify(Test))
}
