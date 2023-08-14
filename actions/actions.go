package actions

import (
	"log"
	"os"
	"slices"
	"usbkill-go/commands"
	"usbkill-go/configs"
	"usbkill-go/devices"
	"usbkill-go/utils"
)

var (
	Config configs.Config
	cmds   map[string]commands.CmdDesc
)

func NewDevices(unknownDevices devices.Device) {
	if Config.HasKillSwitch() {
		return
	}

	commands.UserCommands(Config.Commands)
	cmds[Config.Action].Fun()
}

func MissingDevices(detachedDevices devices.Device) {
	if Config.HasKillSwitch() {
		if detachedDevices.Contains(Config.KillSwitch.ProductId, Config.KillSwitch.VendorId) {
			cmds[Config.Action].Fun()
		}
		return
	}

	commands.UserCommands(Config.Commands)
	cmds[Config.Action].Fun()
}

func Init() {
	cmds = commands.GenCmds(Config.DryRun)

	if !slices.Contains(utils.MapKeys(cmds), Config.Action) {
		log.Fatalln("action", Config.Action, "doesn't exists.\n", "Available actions:", utils.MapKeys(cmds))
	}

	if cmds[Config.Action].Sudo && os.Geteuid() != 0 && !Config.DryRun {
		log.Fatalln("action", Config.Action, "needs more privileges")
	}
}
