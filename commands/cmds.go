package commands

import (
	"log"
	"runtime"
)

func GenCmds(dry bool) map[string]func() {
	dryFun := func() { log.Println("dry run harmless") }
	def := map[string]func(){
		"poweroff": dryFun,
		"sleep":    dryFun,
	}

	if dry {
		return def
	}

	switch runtime.GOOS {
	case `darwin`:
		def["poweroff"] = DarwinPowerOff
		def["sleep"] = DarwinSleep
	}
	return def
}
