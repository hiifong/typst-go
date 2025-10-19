# typst-go

## Usage

```shell
go get github.com/hiifong/typst-go
```

## Example

```go
package main

import (
	"fmt"
	"os"

	"github.com/hiifong/typst-go"
)

func main() {
	file, err := os.ReadFile("./main.typ")
	if err != nil {
		panic(err)
	}
	
	pdf := typst.CompileToPdf(string(file))

	err = os.WriteFile("output.pdf", pdf, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("PDF compiled successfully: output.pdf")
}
```