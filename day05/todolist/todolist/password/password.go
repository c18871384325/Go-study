package password

import (
	"fmt"
	"crypto/sha256"
	"os"
)

func Password() {
	passinit()
	passverity()
}

func passinit(
	file, err 
)


func fileidexists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else if 
}

