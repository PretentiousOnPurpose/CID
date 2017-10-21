package CID

import (
	"math"
	"strconv"
	"os"
	"log"
	"bufio"
)

func DistancCal(InV , Comp [][]float64) float64 {
	dist := 0.0
	for y := 0; y < 32; y++ {
		for x := 0; x < 32; x++ {
			dist += (InV[y][x] - Comp[y][x])*(InV[y][x] - Comp[y][x])
		}
	}
	return math.Sqrt(dist)
}

func str2float(str string) []float64 {
	p := []float64{}
	for i := 0; i < len(str); i++ {
		val , _ := strconv.Atoi(string(str[i]))
		p = append(p ,float64(val))
	}
	return p
}

//k is 5

func KNNClassify(InV [][]float64) int {
	closePoint := [][]float64{[]float64{1e10, 0},[]float64{1e10, 0},[]float64{1e10, 0},[]float64{1e10, 0},[]float64{1e10, 0}}
	for n := 0; n < 9; n++ {
		num := strconv.Itoa(n)
		for i := 0; i < 150; i++ {
			Comp := OpenImg("trainingDigits/" + num + "_" + strconv.Itoa(i) +".txt")
			dist := DistancCal(InV , Comp)
			for j := 0; j < 5; j++ {
				if closePoint[j][0] > dist {
					closePoint[j] = []float64{dist, float64(n)}
					break
				}
			}
		}
	}
	return Max_Frequent(closePoint)
}

func Max_Frequent(closePoint [][]float64) int {
	p := 0.0
	Map := map[float64]int{}
	for i := 0; i < len(closePoint); i++ {
		if _, ok := Map[closePoint[i][1]]; ok {
			Map[closePoint[i][1]]++
		} else {
			Map[closePoint[i][1]] = 1
		}
	}
	big := 0
	for keys , vals := range Map {
		if vals > big {
			p = keys
		}
	}
	return int(p)
}

func OpenImg(filename string) [][]float64 {
	CompFile , err  := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	CompScan := bufio.NewScanner(CompFile)
	CompScan.Split(bufio.ScanLines)
	Comp := [][]float64{}
	for CompScan.Scan() {
		Comp = append(Comp, str2float(CompScan.Text()))
	}
	return Comp
}
