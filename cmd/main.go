package main

import "github.com/RulezKT/findf"

const DIR = "files"
const CERES_FILE = "ceres.bin"
const CHIRON_FILE = "chiron.bin"

func main() {
	dir := findf.Dir(DIR)
	ceres := floatsfile.LoadBinary(findf.File(dir, CERES_FILE), 457984)
	chiron := floatsfile.LoadBinary(findf.File(dir, CHIRON_FILE), 458880)

}
