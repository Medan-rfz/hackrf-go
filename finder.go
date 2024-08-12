package hackrf

import (
	"hackrf-driver/internal/helpers"

	"github.com/google/gousb"
)

// FindAllDevices finds all devices HackRF available for interaction and returns a list of them
func FindAllDevices() ([]*HackRF, error) {
	ctx := gousb.NewContext()
	defer ctx.Close()

	devs, _ := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
		return true
	})

	rawDevices := helpers.Where(devs, func(d *gousb.Device) bool {
		name, _ := d.Product()
		return name == "HackRF One"
	})

	hackrfs := helpers.Select(rawDevices, func(d *gousb.Device) *HackRF {
		return NewHackRF(d)
	})

	return hackrfs, nil
}
