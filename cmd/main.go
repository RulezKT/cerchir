package main

import (
	"fmt"

	"github.com/RulezKT/findf"
	"github.com/RulezKT/floatsfile"

	"github.com/RulezKT/cerchir"
)

const DIR = "files"
const CERES_FILE = "ceres.bin"
const CHIRON_FILE = "chiron.bin"

func main() {

	// мое время -682470731.47  [ 1978, 5, 17, 12, 47, 0 ]
	date_in_seconds := int64(-682470731)

	dir := findf.Dir(DIR)
	ceres := floatsfile.LoadBinary(findf.File(dir, CERES_FILE), 457984)
	chiron := floatsfile.LoadBinary(findf.File(dir, CHIRON_FILE), 458880)

	chironCoords := cerchir.Chiron(float64(date_in_seconds), chiron)
	fmt.Println("chironCoords", chironCoords)

	ceresCoords := cerchir.Ceres(float64(date_in_seconds), ceres)
	fmt.Println("ceresCoords", ceresCoords)

}
