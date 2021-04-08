package main

import (
	"log"

	"github.com/zlingqu/go-template-tool/cmd"

)

func main() {

	newGenetateYml := cmd.NewGenetateYmlCommand()
	if err := newGenetateYml.Execute(); err != nil {
		log.Fatal(err)
	}
}
