package main

import (
	"os"
	"regexp"
)

func main() {
	str := `package svgPlot

const (
`
	check := regexp.MustCompile(".*\\.svg")
	repl := regexp.MustCompile("\\.svg")
	files, _ := os.ReadDir(".")
	for _, val := range files {
		if check.MatchString(val.Name()) {
			file, _ := os.ReadFile(val.Name())
			name := repl.ReplaceAllString(val.Name(), "=`")
			str += name + string(file) + "`\n"
		}
	}
	str += ")\n"
	_ = os.WriteFile("big_outputs_test.go", []byte(str), 777)
}
