package config

import (
	"flag"
)

type Config struct {
	After   int
	Before  int
	Context int

	Count      bool
	IgnoreCase bool
	LineNum    bool
	Invert     bool
	Fixed      bool
}

func Export(flag *flag.FlagSet) func() *Config {
	var c Config

	flag.IntVar(&c.After, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(&c.Before, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(&c.Context, "C", 0, "печатать ±N строк вокруг совпадения")

	flag.BoolVar(&c.Count, "c", false, "напечатать количество строк")
	flag.BoolVar(&c.IgnoreCase, "i", false, "игнорировать регистр")
	flag.BoolVar(&c.LineNum, "n", false, "печатать номер строки")
	flag.BoolVar(&c.Invert, "v", false, "вместо совпадения, исключать")
	flag.BoolVar(&c.Fixed, "f", false, "точное совпадение со строкой")

	return func() *Config {
		if c.Context != 0 {
			c.After = c.Context
			c.Before = c.Context
		}

		return &c
	}
}
