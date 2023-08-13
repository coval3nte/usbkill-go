package commands

import (
	"log"
	"os/exec"
)

func DarwinPowerOff() {
	toKill := []string{
		"Finder",
		"loginwindow",
	}
	for _, c := range toKill {
		if err := exec.Command("/usr/bin/killall", c).Run(); err != nil {
			log.Fatalln(err)
		}
	}

	if err := exec.Command(`/sbin/halt`, `-q`).Run(); err != nil {
		log.Fatalln(err)
	}
}

func DarwinSleep() {
	if err := exec.Command("/usr/bin/pmset", "sleepnow").Run(); err != nil {
		log.Fatalln(err)
	}
}
