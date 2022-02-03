package waveact

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"

	movingaverage "github.com/RobinUS2/golang-moving-average"
)

var processing string = "keys"

var bindAddr string = "0.0.0.0:443"
var certificateChain string = "fullchain.pem"
var privateKey string = "privkey.pem"

var MovingAverageWindow int = 5

var maX *movingaverage.MovingAverage = movingaverage.New(MovingAverageWindow)
var maY *movingaverage.MovingAverage = movingaverage.New(MovingAverageWindow)
var maZ *movingaverage.MovingAverage = movingaverage.New(MovingAverageWindow)

type Vector struct {
	X []int
	Y []int
	Z []int
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func lowest(a int, b int, c int) int {
	ans := math.Min(float64(a), float64(b))
	return int(math.Min(ans, float64(c)))
}

func findDiv(x int) int {
	for i := 4; i > 0; i-- {
		if (x % i) == 0 {
			return i
		}
	}
	return x
}

func chunks(xs []int, chunkSize int) [][]int {
	if len(xs) == 0 {
		return nil
	}
	divided := make([][]int, (len(xs)+chunkSize-1)/chunkSize)
	prev := 0
	i := 0
	till := len(xs) - chunkSize
	for prev < till {
		next := prev + chunkSize
		divided[i] = xs[prev:next]
		prev = next
		i++
	}
	divided[i] = xs[prev:]
	return divided
}

func evenSplit(v []int) [][]int {
	div := findDiv(len(v))
	chunks := chunks(v, div)
	return chunks
}

func healthCheck(w http.ResponseWriter, req *http.Request) {
	fmt.Println("[+] Client connected")
}

func receive(w http.ResponseWriter, req *http.Request) {
	var v Vector

	err := json.NewDecoder(req.Body).Decode(&v)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if processing == "keys" {
		ProcessDataToKeys(v)
	} else {
		ProcessDataToMidi(v)
	}

}

func VectorServerListen(bindAddr string, certificateChain string, privateKey string) {
	http.HandleFunc("/data", receive)
	http.HandleFunc("/health", healthCheck)

	fmt.Printf("[+] Starting accelerometer receiver for %s\n", processing)

	err := http.ListenAndServeTLS(bindAddr, certificateChain, privateKey, nil)
	handleError(err)
}
