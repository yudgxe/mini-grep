package serch

import (
	"bufio"
	"os"
	"strings"

	"github.com/yudgxe/grep-custom/config"
	"github.com/yudgxe/grep-custom/serch/result"
)

func Find(path string, lookFor string, c *config.Config) (*result.Result, error) {
	result := result.New(c)

	file, err := os.Open(path)
	if err != nil {
		return result, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if c.IgnoreCase {
		lookFor = strings.ToLower(lookFor)
	}

	if c.Fixed {
		lookFor = strings.TrimSpace(lookFor)
	}

	for scanner.Scan() {
		str := scanner.Text()
		if c.IgnoreCase {
			str = strings.ToLower(str)
		}

		result.AddData(scanner.Text())

		if c.Fixed {
			strs := strings.Fields(str)
			for _, v := range strs {
				if v == lookFor {
					result.AddFound(result.DataLastIndex())
					break
				}
			}
		} else {
			if strings.Contains(str, lookFor) {
				result.AddFound(result.DataLastIndex())
			}
		}
	}

	return result, nil
}
