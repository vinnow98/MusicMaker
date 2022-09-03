package rawMusicData

import (
	"fmt"
	"math/rand"
	"time"
)

var FinalOutput []int

func FirstNote() {
	rand.Seed(time.Now().UnixNano())
	note := rand.Intn(5 + 1)
	if note == 1 || note == 3 || note == 5 {
		FinalOutput = append(FinalOutput, note)
	} else {
		FirstNote()
	}
}

func SecondNote() {
	var noteSlice []int
	for i := range Psalm {
		for idx, elem := range Psalm[i] {
			if elem == FinalOutput[0] {
				if idx != len(Psalm[i])-1 {
					noteSlice = append(noteSlice, Psalm[i][idx+1])
				}
			}
		}
	}
	fmt.Printf("Notes in second slice %v\n", noteSlice)
	rand.Seed(time.Now().UnixNano())
	note := noteSlice[rand.Intn(len(noteSlice))]
	FinalOutput = append(FinalOutput, note)

}

func ThirdNote() {
	var noteSlice []int
	for i := range Psalm {
		for idx, elem := range Psalm[i] {
			if elem == FinalOutput[1] {
				if idx != len(Psalm[i])-1 {
					noteSlice = append(noteSlice, Psalm[i][idx+1])
				}
			}
		}
	}
	fmt.Printf("Notes in third slice %v\n", noteSlice)
	rand.Seed(time.Now().UnixNano())
	note := noteSlice[rand.Intn(len(noteSlice))]
	FinalOutput = append(FinalOutput, note)
}
