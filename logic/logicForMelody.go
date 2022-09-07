package logic

import (
	"fmt"
	"math/rand"
	"music/rawMusicData"
	"time"
)

var FinalOutput []int

func FirstNote() {
	var firstNoteSlice []int
	for i := range rawMusicData.Melody {
		firstNoteSlice = append(firstNoteSlice, rawMusicData.Melody[i][0])
	}
	fmt.Printf("#1 note options, %v\n", firstNoteSlice)
	rand.Seed(time.Now().UnixNano())
	note := firstNoteSlice[rand.Intn(len(firstNoteSlice))]
	FinalOutput = append(FinalOutput, note)
}

func FirstHalfNotes(noOfNotes int) {
	for x := 0; x < noOfNotes-1; x++ {
		var noteSlice []int
		for i := range rawMusicData.Melody {
			for idx, elem := range rawMusicData.Melody[i] {
				if elem == FinalOutput[x] {
					if idx != len(rawMusicData.Melody[i])-1 {
						noteSlice = append(noteSlice, rawMusicData.Melody[i][idx+1])
					}

				}
			}
		}

		fmt.Printf("#%v note options, %v\n", x+2, noteSlice)
		rand.Seed(time.Now().UnixNano())
		note := noteSlice[rand.Intn(len(noteSlice))]
		FinalOutput = append(FinalOutput, note)
	}
}

func SecondHalfNotes(noOfNotes int) {
	for x := 0; x < noOfNotes-1; x++ {
		var noteSlice []int
		for i := range rawMusicData.Melody {
			for idx, elem := range rawMusicData.Melody[i] {
				if elem == FinalOutput[x] {
					if idx != len(rawMusicData.Melody[i])-1 {
						noteSlice = append(noteSlice, rawMusicData.Melody[i][idx+1])
					}

				}
			}
		}

		fmt.Printf("#%v note options, %v\n", x+2, noteSlice)
		rand.Seed(time.Now().UnixNano())
		note := noteSlice[rand.Intn(len(noteSlice))]
		FinalOutput = append(FinalOutput, note)
	}
}

func LastNote() {
	var lastNoteSlice []int
	for i := range rawMusicData.Melody {
		lastNoteSlice = append(lastNoteSlice, rawMusicData.Melody[i][len(rawMusicData.Melody[i])-1])
	}
	fmt.Printf("#last note options, %v\n", lastNoteSlice)
	rand.Seed(time.Now().UnixNano())
	note := lastNoteSlice[rand.Intn(len(lastNoteSlice))]
	FinalOutput = append(FinalOutput, note)
}
