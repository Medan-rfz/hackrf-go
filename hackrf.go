package hackrf

import (
	"encoding/binary"

	"github.com/google/gousb"
)

type HackRF struct {
	device      *gousb.Device
	inter       *gousb.Interface
	inEndpoint  *gousb.InEndpoint
	outEndpoint *gousb.OutEndpoint

	mode       uint16
	centerFreq uint64
	sampleRate uint32

	lna   uint16
	vga   uint16
	txVga uint16

	rxBuffer []byte
}

// NewHackRF constructor
func NewHackRF(dev *gousb.Device) *HackRF {
	return &HackRF{
		device: dev,
	}
}

// Init init hackrf device (read interfaces and endpoints)
func (d *HackRF) Init() error {
	var err error

	d.inter, _, err = d.device.DefaultInterface()
	if err != nil {
		return err
	}

	d.inEndpoint, err = d.inter.InEndpoint(inEndpointAddr)
	if err != nil {
		return err
	}

	d.outEndpoint, err = d.inter.OutEndpoint(outEndpointAddr)
	if err != nil {
		return err
	}

	return nil
}

// SetCenterFrequency sets the central operating frequency
func (d *HackRF) SetCenterFrequency(freq uint64) error {
	freqMHz, freqHz := сonvertFreqHzToMHz(freq)

	freqMHzBytes := make([]uint8, 4)
	binary.LittleEndian.PutUint32(freqMHzBytes, freqMHz)

	freqHzBytes := make([]uint8, 4)
	binary.LittleEndian.PutUint32(freqHzBytes, freqHz)

	packet := append(freqMHzBytes, freqHzBytes...)
	_, err := d.device.Control(vendorRequestTypeToDevice, VendorRequestSetFreq, 0, 0, packet)
	return err
}

// SetSampleRateManual manually sets the values ​​of the synthesizer frequency and the divider coefficient
func (d *HackRF) SetSampleRateManual(freq uint32, div uint32) error {
	freqBytes := make([]uint8, 4)
	binary.LittleEndian.PutUint32(freqBytes, freq)

	divBytes := make([]uint8, 4)
	binary.LittleEndian.PutUint32(divBytes, div)

	packet := append(freqBytes, divBytes...)
	_, err := d.device.Control(vendorRequestTypeToDevice, VendorRequestSampleRateSet, 0, 0, packet)
	return err
}

// SetSampleRate setting the sampling frequency with automatic calculation of the synthesizer frequency and the divider coefficient
func (d *HackRF) SetSampleRate(sr float64) error {
	return d.SetSampleRateManual(calculateFrequencyConfig(sr))
}

// SetLNA sets the gain factor for the LNA
func (d *HackRF) SetLNA(lna uint16) error {
	if lna > lnaMax {
		return errInvalidLna
	}

	retval := make([]uint8, 1)
	lna &^= 0x7
	n, err := d.device.Control(vendorRequestTypeFromDevice, VendorRequestSetLNAGain, 0, lna, retval)

	if n != 1 || retval[0] == 0 {
		err = errInvalidParam
	}

	d.lna = lna
	return err
}

// SetVGA sets the gain for the VGA of the receive path
func (d *HackRF) SetVGA(vga uint16) error {
	if vga > vgaMax {
		return errInvalidVga
	}

	retval := make([]uint8, 1)
	vga &^= 0x1
	n, err := d.device.Control(vendorRequestTypeFromDevice, VendorRequestSetLNAGain, 0, vga, retval)

	if n != 1 || retval[0] == 0 {
		err = errInvalidParam
	}

	d.vga = vga
	return err
}

// SetTxVGA sets the gain for the VGA of the transmit path
func (d *HackRF) SetTxVGA(txVga uint16) error {
	if txVga > txVgaMax {
		return errInvalidTxVga
	}

	retval := make([]uint8, 1)
	n, err := d.device.Control(vendorRequestTypeFromDevice, VendorRequestSetLNAGain, 0, txVga, retval)

	if n != 1 || retval[0] == 0 {
		err = errInvalidParam
	}

	d.txVga = txVga
	return err
}

// EnableRx turns on the receiver
func (d *HackRF) EnableRx(callback func([]byte)) error {
	_, err := d.device.Control(vendorRequestTypeToDevice, VendorRequestSetTransceiverMode, TranceiverModeReceive, 0, nil)

	if err == nil {
		d.mode = TranceiverModeReceive
		d.readReceivedDataStart(callback)
	}

	return err
}

// EnableTx turns on the transmission
func (d *HackRF) EnableTx() error {
	_, err := d.device.Control(vendorRequestTypeToDevice, VendorRequestSetTransceiverMode, TranceiverModeTransmit, 0, nil)

	if err == nil {
		d.mode = TranceiverModeTransmit
	}

	return err
}

// Disable turns off any interaction with the environment
func (d *HackRF) Disable() error {
	_, err := d.device.Control(vendorRequestTypeToDevice, VendorRequestSetTransceiverMode, TranceiverModeOff, 0, nil)

	if err == nil {
		d.mode = TranceiverModeOff
	}

	return err
}
