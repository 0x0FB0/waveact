package waveact

import (
	"fmt"
	"time"

	"github.com/micmonay/keybd_event"
)

var kb keybd_event.KeyBonding

var RightSwipeKey int = keybd_event.VK_RIGHT
var LeftSwipeKey int = keybd_event.VK_LEFT

func ProcessDataToKeys(vector Vector) {
	integral := lowest(len(vector.X), len(vector.Y), len(vector.Z))

	var i, j int

	var armed bool = false

	var trimmedX []int
	var trimmedY []int
	var trimmedZ []int

	for i := 0; i < integral; i++ {
		trimmedX = append(trimmedX, vector.X[i])
		trimmedY = append(trimmedY, vector.Y[i])
		trimmedZ = append(trimmedZ, vector.Z[i])
	}

	chunksX := evenSplit(vector.X)
	chunksY := evenSplit(vector.Y)
	chunksZ := evenSplit(vector.Z)

	for j = range chunksX {
		for i = 0; i < (integral / len(chunksX)); i++ {

			maX.Add(float64(chunksX[j][i]))
			maY.Add(float64(chunksY[j][i]))
			maZ.Add(float64(chunksZ[j][i]))

			time.Sleep(time.Millisecond * 40)

		}

		baX := maX.Avg()
		baY := maY.Avg()
		baZ := maZ.Avg()

		// fmt.Printf("Chunk stats    X: %f Y: %f Z: %f\n", maX.Avg(), maY.Avg(), maZ.Avg())

		if baX < -450 && baX > -850 && baY > 350 && baY < 750 && baZ > -650 && baZ < 0 {
			// hand in ~60 degrees straight positon initiates the gesture
			armed = true
		}
		if armed && baX < -150 && baX > -450 && baY > 450 && baY < 650 && baZ > -850 && baZ < -450 {
			fmt.Println("[>] Hand RIGHT SWIPE")
			kb.SetKeys(RightSwipeKey)
			kb.Press()
			time.Sleep(10 * time.Millisecond)
			kb.Release()
			// wait for return to straight position
			armed = false
		}
		if armed && baX < -450 && baX > -850 && baY > 100 && baY < 650 && baZ > 100 && baZ < 650 {
			fmt.Println("[<] Hand LEFT SWIPE")
			kb.SetKeys(LeftSwipeKey)
			kb.Press()
			time.Sleep(10 * time.Millisecond)
			kb.Release()
			// wait for return to straight position
			armed = false
		}

	}

}

func SetupKeyboard() {

	processing = "keys"

	var err error

	kb, err = keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}
}
