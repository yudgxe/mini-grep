package result

import (
	"fmt"

	"github.com/yudgxe/grep-custom/cmath"
	"github.com/yudgxe/grep-custom/config"
)

type Result struct {
	data  []string
	found []int

	c *config.Config
}

func New(c *config.Config) *Result {
	return &Result{
		make([]string, 0),
		make([]int, 0),
		c,
	}
}

func (r *Result) AddData(in string) {
	r.data = append(r.data, in)
}

func (r *Result) AddFound(index int) {
	r.found = append(r.found, index)
}

func (r *Result) Print() {
	if r.c.Count {
		fmt.Println(len(r.found))
		return
	}

	isPrinted := make(map[int]bool)
	for _, v := range r.found {
		for i := cmath.Max(0, v-r.c.Before); i < v; i++ {
			isPrinted[i] = true
		}

		isPrinted[v] = true

		for i := v + 1; i < cmath.Min(v+r.c.After+1, len(r.data)); i++ {
			isPrinted[i] = true
		}
	}

	for i := range r.data {
		if r.c.Invert {
			if !isPrinted[i] {
				if r.c.LineNum {
					fmt.Printf("%d ", i+1)
				}
				fmt.Println(r.data[i])
			}

			continue
		}

		if isPrinted[i] {
			if r.c.LineNum {
				fmt.Printf("%d ", i+1)
			}
			fmt.Println(r.data[i])
		}
	}
}

func (r *Result) DataLastIndex() int {
	return len(r.data) - 1
}
