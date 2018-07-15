package main

import (
	"fmt"

	"github.com/noteshare/internal/noteshare"
)

func main() {
	fmt.Println(``)
	fmt.Println(`==> Starting noteshare server`)

	noteshare.Run()

	fmt.Println(`==> Ending`)
	fmt.Println(``)
}
