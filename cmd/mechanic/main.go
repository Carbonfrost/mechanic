package main

import (
	"os"

	"github.com/Carbonfrost/mechanic/internal/cmd/mechanic"
)

func main() {
	mechanic.Run(os.Args)
}
