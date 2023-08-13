package devices

import (
	"bytes"
	"encoding/json"
	"log"
	"os/exec"
)

type Items struct {
	Items            []Items `json:"_items,omitempty"`
	Name             string  `json:"_name"`
	LocationId       string  `json:"location_id"`
	BcdDevice        string  `json:"bcd_device,omitempty"`
	BusPower         string  `json:"bus_power,omitempty"`
	BusPowerUsed     string  `json:"bus_power_used,omitempty"`
	DeviceSpeed      string  `json:"device_speed,omitempty"`
	ExtraCurrentUsed string  `json:"extra_current_used,omitempty"`
	Manufacturer     string  `json:"manufacturer,omitempty"`
	ProductId        string  `json:"product_id,omitempty"`
	VendorId         string  `json:"vendor_id,omitempty"`
}

type UsbDarwin struct {
	SPUSBDataType []struct {
		Items          []Items `json:"_items,omitempty"`
		Name           string  `json:"_name"`
		HostController string  `json:"host_controller"`
	} `json:"SPUSBDataType"`
}

func (_ *UsbDarwin) runTool() []byte {
	var out bytes.Buffer
	cmd := exec.Command(`/usr/sbin/system_profiler`,
		`SPUSBDataType`, `-json`, `-detailLevel`, `mini`)
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		log.Panicln(err)
	}
	return out.Bytes()
}

func parseItem(d Device, i []Items) Device {
	for _, item := range i {
		if len(item.Items) > 0 {
			d = parseItem(d, item.Items)
		}

		identifier, valid := d.format(item.ProductId, item.VendorId)
		if !valid {
			continue
		}
		if _, k := d[identifier]; !k {
			d[identifier] = 1
			continue
		}
		d[identifier] += 1
	}
	return d
}
func (ld *UsbDarwin) Get() Device {
	d := make(Device)

	result := *new(UsbDarwin)
	if err := json.Unmarshal(ld.runTool(), &result); err != nil {
		log.Panicln(err)
	}

	for _, r := range result.SPUSBDataType {
		if len(r.Items) == 0 {
			continue
		}

		d = parseItem(d, r.Items)
	}

	return d
}
