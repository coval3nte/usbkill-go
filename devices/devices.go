package devices

import (
	"strconv"
	"strings"
)

const sep = ":"

type Device map[string]int

func (_ *Device) format(s ...string) (string, bool) {
	var z string

	for i, v := range s {
		if len(v) == 0 {
			return "", false
		}

		z += strings.TrimPrefix(strings.TrimSpace(strings.Split(v, " ")[0]), "0x")
		if i+1 < len(s) {
			z += sep
		}
	}

	return z, true
}

func (d *Device) Equal(b Device) bool {
	if len(*d) != len(b) {
		return false
	}

	for i := range *d {
		if _, k := b[i]; !k {
			return false
		}
	}
	return true
}

func (d Device) Subi(b Device) Device {
	for i := range d {
		if _, j := b[i]; j { // ??
			delete(d, i)
		}
	}
	return d
}

func deepCopy[T Device](dst, src T) { // inplace
	for k, v := range src {
		dst[k] = v
	}
}

func (d Device) Sub(b Device) Device {
	l := make(Device)
	deepCopy(l, d)
	return l.Subi(b)
}

func (d Device) AreIn(b Device) bool {
	if len(b) == 0 && len(d) > 0 {
		return false
	}

	for k, v := range d {
		if _, j := b[k]; !(j && b[k] == v) {
			return false
		}
	}
	return true
}

func (d Device) Sumi(b Device) Device {
	for k, v := range b {
		if _, j := d[k]; !j {
			d[k] = v
		}
	}
	return d
}

func (d Device) Sum(b Device) Device {
	l := make(Device)
	deepCopy(l, d)
	return l.Sumi(b)
}

func (d Device) String() string {
	msg := *new(string)
	l, cnt := len(d), 0
	for k, v := range d {
		msg += k + ": " + strconv.Itoa(v)
		if cnt < l-1 {
			msg += "\n"
		}
		cnt++
	}

	if len(msg) > 0 {
		return msg
	}
	return "no device..."
}

func (d Device) Contains(productId, vendorId string) bool {
	if res, ok := d.format(productId, vendorId); ok {
		_, k := d[res]
		return k
	}
	return false
}
