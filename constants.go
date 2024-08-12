package hackrf

import "github.com/google/gousb"

const (
	VendorRequestSetTransceiverMode uint8 = iota + 1
	VendorRequestMax2837Write
	VendorRequestMax2837Read
	VendorRequestSi5351CWrite
	VendorRequestSi5351CRead
	VendorRequestSampleRateSet
	VendorRequestBasebandFilterBandwidthSet
	VendorRequestRFFC5071Write
	VendorRequestRFFC5071Read
	VendorRequestSPIFlashErase
	VendorRequestSPIFlashWrite
	VendorRequestSPIFlashRead
	VendorRequestBoardIDRead = iota + 2
	VendorRequestVersionStringRead
	VendorRequestSetFreq
	VendorRequestAmpEnable
	VendorRequestBoardPartIDSerialNoRead
	VendorRequestSetLNAGain
	VendorRequestSetVGAGain
	VendorRequestSetTXVGAGain
	VendorRequestAntennaEnable = iota + 3
	VendorRequestSetFreqExplicit
	VendorRequestUSBWCIDVendorReq
	VendorRequestInitSweep
	VendorRequestOperacakeGetBoards
	VendorRequestOperacakeSetPorts
	VendorRequestSetHWSyncMode
	VendorRequestReset
	VendorRequestOperacakeSetRanges
	VendorRequestClkoutEnable
	VendorRequestSPIFlashStatus
	VendorRequestSPIFlashClearStatus
	VendorRequestOperacakeGPIOTest
	VendorRequestCPLDChecksum
	VendorRequestUIEnable
	VendorRequestOperacakeSetMode
	VendorRequestOperacakeGetMode
	VendorRequestOperacakeSetDwellTimes
	VendorRequestGetM0State
	VendorRequestSetTXUnderrunLimit
	VendorRequestSetRXOverrunLimit
	VendorRequestGetCLKINStatus
	VendorRequestBoardRevRead
	VendorRequestSupportedPlatformRead
)

const (
	TranceiverModeOff uint16 = iota
	TranceiverModeReceive
	TranceiverModeTransmit
	TranceiverModeSS
	TranceiverModeCPLDUpdate
	TranceiverModeRxSweep
)

const (
	vendorRequestTypeToDevice   = gousb.ControlVendor | gousb.ControlOut | gousb.ControlDevice
	vendorRequestTypeFromDevice = gousb.ControlVendor | gousb.ControlIn | gousb.ControlDevice
)

const (
	inEndpointAddr  = 0x81
	outEndpointAddr = 0x02
)

const (
	transferBufferSize = 262144
	deviceBufferSize   = 32768
)

const (
	lnaMax   = 40
	vgaMax   = 62
	txVgaMax = 47
)
