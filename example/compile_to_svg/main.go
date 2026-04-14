package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hiifong/typst-go"
)

func main() {
	file, err := os.ReadFile("./main.typ")
	if err != nil {
		panic(err)
	}

	now := time.Now()
	svg := typst.CompileToSvg(string(file))
	fmt.Println("cost:", time.Since(now))

	err = os.WriteFile("output.svg", svg, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("PDF compiled successfully: output.svg")
}
