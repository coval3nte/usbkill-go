package commands

import (
	"log"
	"runtime"
)

func GenCmds(dry bool) map[string]func() {
	def := map[string]func(){
		"poweroff": func() { log.Println("dry run harmless") },
	}

	if dry {
		return def
	}

	switch runtime.GOOS {
	case `darwin`:
		def["poweroff"] = DarwinPowerOff
	}
	return def
}
