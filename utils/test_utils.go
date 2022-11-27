package utils

import (
	"os"
	"strings"
)

func TestsAreRunning() bool {
	return strings.HasSuffix(os.Args[0], ".test")
}
