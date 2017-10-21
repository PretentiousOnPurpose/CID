package main

import (
	"github.com/chawat/CID"
	"fmt"
	"strconv"
)

func main() {
	corr := 0.0
	for i := 0; i < 9; i++ {
		for j := 0; j < 80; j++ { // 80 because we have only around 87 test example for "0"
			filename := "testDigits/" + strconv.Itoa(i) + "_" + strconv.Itoa(j)+".txt"
			if CID.KNNClassify(CID.OpenImg(filename)) == i {
				corr++
			}
		}
	}
	fmt.Println("Accuray : ", corr/float64(150*9))
}
