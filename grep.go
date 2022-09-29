package main

import (
	"flag"
	"log"
	"os"

	"github.com/yudgxe/grep-custom/config"
	"github.com/yudgxe/grep-custom/serch"
)

func main() {
	flags := flag.NewFlagSet("grep", flag.ExitOnError)
	callback := config.Export(flags)

	if err := flags.Parse(os.Args[1:]); err != nil {
		log.Fatalln(err)
	}

	config := callback()

	path := os.Args[len(os.Args)-1]
	lookFor := os.Args[len(os.Args)-2]

	r, err := serch.Find(path, lookFor, config)
	if err != nil {
		log.Fatalln(err)
	}

	r.Print()
}
