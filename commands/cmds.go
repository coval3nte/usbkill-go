package commands

import (
	"log"
	"runtime"
)

type CmdDesc struct {
	Fun  func()
	Sudo bool
}

func GenCmds(dry bool) map[string]CmdDesc {
	dryFun := func() { log.Println("dry run harmless") }
	def := map[string]CmdDesc{
		"poweroff": {
			Fun:  dryFun,
			Sudo: true,
		},
		"sleep": {
			Fun:  dryFun,
			Sudo: false,
		},
	}

	if dry {
		return def
	}

	switch runtime.GOOS {
	case `darwin`:
		def["poweroff"] = CmdDesc{
			Fun:  DarwinPowerOff,
			Sudo: true,
		}
		def["sleep"] = CmdDesc{
			Fun:  DarwinSleep,
			Sudo: false,
		}
	}
	return def
}
