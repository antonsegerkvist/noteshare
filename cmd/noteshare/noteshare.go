package main

import (
	"fmt"

	"github.com/noteshare/internal/noteshare"
)

func main() {
	fmt.Println(`==> Starting noteshare server`)
	noteshare.Run()
	fmt.Println(`==> Ending noteshare server`)
}
