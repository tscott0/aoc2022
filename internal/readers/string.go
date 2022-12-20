package readers

import (
	"fmt"
	"os"
)

func ReadAll(file string) string {
	b, err := os.ReadFile(file)
	if err != nil {
		fmt.Print(err)
	}

	return string(b)
}
