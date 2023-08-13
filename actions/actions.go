package actions

import (
	"usbkill-go/commands"
	"usbkill-go/configs"
	"usbkill-go/devices"
)

var (
	Config configs.Config
	cmds   map[string]func()
)

func NewDevices(unknownDevices devices.Device) {
	if Config.HasKillSwitch() {
		return
	}
	cmds["poweroff"]()
}

func MissingDevices(detachedDevices devices.Device) {
	if Config.HasKillSwitch() {
		if detachedDevices.Contains(Config.KillSwitch.ProductId, Config.KillSwitch.VendorId) {
			cmds["poweroff"]()
		}
		return
	}
	cmds["poweroff"]()
}

func Init() {
	cmds = commands.GenCmds(Config.DryRun)
}
