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
	pdf := typst.CompileToPdf(string(file))
	fmt.Println("cost:", time.Since(now))

	err = os.WriteFile("output.pdf", pdf, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("PDF compiled successfully: output.pdf")
}
