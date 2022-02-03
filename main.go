package main

import (
	"flag"
	wva "waveact/waveact"
)

var processingType string
var window *int

var bindAddr string
var certFile string
var keyFile string

var midiInterface *int
var midiChannel *int
var midiCcX *int
var midiCcY *int
var midiCcZ *int

func main() {

	flag.StringVar(&processingType, "process", "keys", "Type of processing for accelerometer data (midi/keys)")
	flag.StringVar(&bindAddr, "bind", "0.0.0.0:443", "Address to listen on for data receiver")
	flag.StringVar(&certFile, "cert", "fullchain.pem", "TLS certificate to use")
	flag.StringVar(&keyFile, "key", "privkey.pem", "TLS private key to use")

	window = flag.Int("w", 5, "Moving average window")
	midiInterface = flag.Int("i", 1, "Midi output interface index")
	midiChannel = flag.Int("c", 11, "Midi output channel")
	midiCcX = flag.Int("x", 30, "Midi control channel (CC) for X axis")
	midiCcY = flag.Int("y", 31, "Midi control channel (CC) for Y axis")
	midiCcZ = flag.Int("z", 32, "Midi control channel (CC) for Z axis")

	flag.Parse()

	if processingType == "keys" {
		wva.SetupKeyboard()
		wva.VectorServerListen(bindAddr, certFile, keyFile)
	} else if processingType == "midi" {
		wva.SetupMidi(*midiInterface, uint8(*midiChannel), uint8(*midiCcX), uint8(*midiCcY), uint8(*midiCcZ))
		wva.VectorServerListen(bindAddr, certFile, keyFile)
	}
}
