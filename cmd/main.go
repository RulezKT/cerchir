package main

import (
	"fmt"

	"github.com/RulezKT/cerchir"
)

const DIR = "files"
const CERES_FILE = "ceres.bin"
const CHIRON_FILE = "chiron.bin"

func main() {

	// мое время -682470731.47  [ 1978, 5, 17, 12, 47, 0 ]
	date_in_seconds := float64(-682470731)

	cc := cerchir.CerChir{}
	cc.Load(DIR)

	fmt.Println("chironCoords", cc.CalcChiron(date_in_seconds))

	fmt.Println("ceresCoords", cc.CalcCeres(date_in_seconds))

}
