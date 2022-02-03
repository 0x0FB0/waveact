package waveact

import (
	"fmt"
	"time"

	"gitlab.com/gomidi/midi/writer"
	driver "gitlab.com/gomidi/rtmididrv"
)

var wr writer.ChannelWriter
var drv *driver.Driver

var ControlChannelX uint8 = 30
var ControlChannelY uint8 = 0
var ControlChannelZ uint8 = 0
var midiOutChannel uint8 = 11
var midiOutIntf uint8 = 1

func scaleTo8(unscalled int, min_allowed int, max_allowed int, min int, max int) uint8 {
	if unscalled >= 1000 {
		unscalled = 1000
	} else if unscalled <= -1000 {
		unscalled = -1000
	}
	normalized := unscalled + 1000
	return uint8((max_allowed-min_allowed)*(normalized-min)/(max-min) + min_allowed)
}

func midiWriter(cc uint8, val uint8) {

	drv, err := driver.New()
	handleError(err)

	outs, err := drv.Outs()
	handleError(err)

	out := outs[midiOutIntf]
	handleError(out.Open())

	wr = writer.New(out)

	wr.SetChannel(midiOutChannel)

	writer.ControlChange(wr, cc, val)

	time.Sleep(time.Millisecond * 13)

	out.Close()
	out.Close()
}

func ProcessDataToMidi(vector Vector) {

	fX := vector.X
	fY := vector.Y
	fZ := vector.Z

	integral := lowest(len(fX), len(fY), len(fZ))

	// fmt.Println(midiOutChannel)

	for i := 0; i < integral; i++ {
		maX.Add(float64(fX[i]))
		maY.Add(float64(fY[i]))
		maZ.Add(float64(fZ[i]))
		//fmt.Println(maX.Avg())
		//fmt.Println(int(maX.Avg()))
		fmt.Println(scaleTo8(int(maX.Avg()), 0, 255, 0, 2000))
		midiWriter(ControlChannelX, scaleTo8(int(maX.Avg()), 0, 255, 0, 2000))
		// MIDI channel throughput saved
		if ControlChannelY != 0 {
			midiWriter(ControlChannelY, scaleTo8(int(maY.Avg()), 0, 255, 0, 2000))
		}
		if ControlChannelZ != 0 {
			midiWriter(ControlChannelZ, scaleTo8(int(maZ.Avg()), 0, 255, 0, 2000))
		}
	}
}

func SetupMidi(outIntf int, midiChannel uint8, ccX uint8, ccY uint8, ccZ uint8) {

	drv, err := driver.New()
	handleError(err)

	outs, err := drv.Outs()
	handleError(err)

	fmt.Println(outs)

	drv.Close()

	processing = "midi"

	ControlChannelX = ccX
	ControlChannelY = ccY
	ControlChannelZ = ccZ

	midiOutChannel = midiChannel
	midiOutIntf = uint8(outIntf)
}
