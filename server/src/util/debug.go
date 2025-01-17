package util

import (
	"fmt"
)

const DEBUG_FLAG = true

func Log(message string) {
	if !DEBUG_FLAG {
		return
	}

	fmt.Println("[INFO] " + message)
}
