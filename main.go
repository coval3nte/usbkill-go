//go:build (darwin && cgo) || (linux && cgo)

package main

import (
	"log"
	"math"
	"os"
	"runtime"
	"time"
	"usbkill-go/actions"
	"usbkill-go/configs"
	"usbkill-go/devices"
)

var (
	conf       = configs.Config{}
	getDevices = map[string]func() devices.Device{
		"darwin": func() devices.Device {
			usbDarwin := devices.UsbDarwin{}
			return usbDarwin.Get()
		},
	}[runtime.GOOS]

	localDevices       = make(devices.Device)
	whitelistedDevices = make(devices.Device)
)

func helper(old devices.Device, new devices.Device) bool {
	if !old.Equal(new) {
		return true
	}

	time.Sleep(time.Second / time.Duration(math.Pow(2, 3)))

	return false
}

func detectChanges(c devices.Device) {
	if helper(localDevices, c) {
		if detachedDevices := localDevices.Sub(c); len(localDevices) > len(c) {
			log.Println("detached", detachedDevices)
			actions.MissingDevices(detachedDevices)
		}

		if unknownDevices := c.Sub(localDevices); !unknownDevices.AreIn(whitelistedDevices) {
			log.Println("attached unknown", unknownDevices.Sub(whitelistedDevices))
			actions.NewDevices(unknownDevices)
		}
	}

	localDevices = c
}

func main() {
	// initialisation
	if os.Getuid() != 0 {
		log.Println("only root")
		os.Exit(0)
	}
	conf.Read()
	actions.Config = conf
	actions.Init()

	localDevices = getDevices()
	whitelistedDevices = localDevices.Sum(conf.Whitelisted)
	for {
		detectChanges(getDevices())
	}
}
