package actions

import (
	"log"
	"slices"
	"usbkill-go/commands"
	"usbkill-go/configs"
	"usbkill-go/devices"
	"usbkill-go/utils"
)

var (
	Config configs.Config
	cmds   map[string]func()
)

func NewDevices(unknownDevices devices.Device) {
	if Config.HasKillSwitch() {
		return
	}
	cmds[Config.Action]()
}

func MissingDevices(detachedDevices devices.Device) {
	if Config.HasKillSwitch() {
		if detachedDevices.Contains(Config.KillSwitch.ProductId, Config.KillSwitch.VendorId) {
			cmds[Config.Action]()
		}
		return
	}
	cmds[Config.Action]()
}

func Init() {
	cmds = commands.GenCmds(Config.DryRun)

	if !slices.Contains(utils.MapKeys(cmds), Config.Action) {
		log.Fatalln("action", Config.Action, "doesn't exists.\n", "Available actions:", utils.MapKeys(cmds))
	}
}
