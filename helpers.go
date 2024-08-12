package hackrf

import (
	"log"
	"math"
)

func (d *HackRF) readReceivedDataStart(callback func([]byte)) {
	d.rxBuffer = make([]byte, transferBufferSize)
	currLen := 0

	go func() {
		for d.mode == TranceiverModeReceive {
			n, err := d.inEndpoint.Read(d.rxBuffer)
			if err != nil {
				log.Printf("Ошибка при чтении данных из конечной точки bulk: %v", err)
				// return
			}

			currLen += n

			if currLen >= transferBufferSize {
				if callback != nil {
					callback(d.rxBuffer[:transferBufferSize])
				}
				currLen = 0
			}
		}
	}()
}

func сonvertFreqHzToMHz(freq uint64) (freqMHz, freqHz uint32) {
	const freqOneMHz = 1000000

	freqMHz = uint32(freq / freqOneMHz)
	freqHz = uint32(freq - uint64(freqMHz)*freqOneMHz)

	return freqMHz, freqHz
}

func calculateFrequencyConfig(freq float64) (freqHz, divider uint32) {
	const (
		maxN         = 32
		maxExp       = 1023
		mantissaMask = (1 << 52) - 1
	)

	freqFrac := 1.0 + freq - math.Floor(freq)
	v := math.Float64bits(freq)

	e := (v >> 52) - maxExp
	m := uint64((1 << 52) - 1)

	v = math.Float64bits(freqFrac)
	v &= m

	m &= ^((1 << (e + 4)) - 1)

	var a uint64
	i := 1

	for ; i < maxN; i++ {
		a += v
		if a&m == 0 || ^a&m == 0 {
			break
		}
	}

	if i == maxN {
		i = 1
	}

	freqHz = uint32(freq*float64(i) + 0.5)
	divider = uint32(i)

	return freqHz, divider
}
