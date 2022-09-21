package logic

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//number of bars
//affected by style-
//characteristics
//1.Type of beats
//2. Length of bars
//3. Time signature
//4.Type of key (favours which type?)
//input what style you want
//grand: more straight down beats, longer bars, 3/4 or 4/4 favours major key
//hopeful/happy: more upbeats, average bar, favours a 6/8 beat, major key
//melancholy more down beats, average bars, 3/4 or 4/4, minor key
// Neutral (worship statements)random beats, average bars,random time, major/minor equal

type Rhythm struct {
	beats         []int
	noOfBars      int
	timeSignature int
}

func Characteristics(c int) {
	switch c {
	default:
		Neutral()
	}
}

func Neutral() int {
	timeSignatureList := []int{3, 4, 6}
	rand.Seed(time.Now().UnixNano())
	timeSignature := timeSignatureList[rand.Intn(len(timeSignatureList))]
	switch timeSignature {
	case 3:
		fmt.Println("Time Signature:3/4")
	case 4:
		fmt.Println("Time Signature:4/4")
	case 6:
		fmt.Println("Time Signature:6/8")
	}
	return timeSignature
}

// Input the dummy data fromm the current psalms
//1:   {5, 1, 2, 2, 1, 1, 4, 3, 1, 1, 1},
//1:   {3,1,2}
var FinalRhythmOutput []int
var line = "the La zy Dog Jumps o Ver the Big Brown Dog"
var binaryLine []int

func MainRhythmLogic() {
	sliceLine := strings.Split(line, " ")
	for _, element := range sliceLine {
		if strings.ToUpper(element[0:1]) == element[0:1] {
			binaryLine = append(binaryLine, 1)
		} else {
			binaryLine = append(binaryLine, 0)
		}
	}
	fmt.Println(binaryLine)

	FinalRhythmOutput = append(FinalRhythmOutput, pickUpNote(binaryLine)...)

	fmt.Println(FinalRhythmOutput)
}

func pickUpNote(binaryLine []int) []int {
	noOfBeats := 6
	zeroCount := 0
	pickUpNote := []int{1, 2, 3, 4, 5, 6}
	for _, element := range binaryLine {
		if element == 0 {
			zeroCount++
		} else {
			break
		}
		if element == 1 {
			return nil
		}
	}
	binaryLine = binaryLine[zeroCount:]
	fmt.Println(binaryLine)
	return pickUpNote[noOfBeats-zeroCount:]
}

func importantNote(binaryLine []int) []int {
	if binaryLine[0] == 1 {
		if FinalOutput[len(FinalOutput)-1] == 6 {
			FinalOutput = append(FinalOutput, 1)
		}
		noOfBeats := 6
		oneCount := 0
		for _, element := range binaryLine {
			if element == 1 {
				oneCount++
			} else {
				break
			}
			if element == 0 {
				return nil
			}
		}
		//binaryLine = binaryLine[zeroCount:]
		//fmt.Println(binaryLine)
		//return pickUpNote[noOfBeats-zeroCount:]
	} else {
		unimportantNotes()
	}

}

func unimportantNotes() {

}

//	res1 := bytes.Replace(slice_1, []byte("E"), []byte("e"), 2)
//for _, element := range binaryLine {
//	// fmt.Println(element)
//	if element != 0 {
//		fmt.Println("This is rando number", rando)
//		//newSlice is when everytime there is a 1,
//		//I want to check for all previous 0s
//		// so like [0 0 1]
//		// or like [0 0 5]
//		var newSlice []int
//		lengthOfArray := 0
//		for lengthOfArray < rando {
//			rand.Seed(time.Now().UnixNano())
//			newSlice = append(newSlice, rand.Intn(noOfBeats-1)+1)
//			fmt.Println("This is the slice after append", newSlice)
//			newSlice = removeDuplicateValues(newSlice)
//			fmt.Println("This is the slice after delete dupls", newSlice)
//			lengthOfArray = len(newSlice)
//		}
//		sort.Sort(sort.Reverse(sort.IntSlice(newSlice)))
//		fmt.Println("This is the slice after reverse", newSlice)
//		//reset the rando number
//		rando = 0
//	}
//	if element == 0 {
//		rando++
//	}
//}

//func removeDuplicateValues(intSlice []int) []int {
//	keys := make(map[int]bool)
//	list := []int{}
//
//	// If the key(values of the slice) is not equal
//	// to the already present value in new slice (list)
//	// then we append it. else we jump on another element.
//	for _, entry := range intSlice {
//		if _, value := keys[entry]; !value {
//			keys[entry] = true
//			list = append(list, entry)
//		}
//	}
//	return list
//}

//func MainLogic() {
//	//var outputRhythm []int
//	line := "in the La zy Dog Jumps o Ver the Big Brown Dog"
//	sliceLine := strings.Split(line, " ")
//	//This is make a binary slice to show the position of capital syllables
//	var binarySlice []int
//	for i := range sliceLine {
//		if strings.ToUpper(sliceLine[i][0:1]) != sliceLine[i][0:1] {
//			binarySlice = append(binarySlice, 0)
//		} else {
//			binarySlice = append(binarySlice, 1)
//		}
//	}
//	fmt.Println(binarySlice)
//
//	//for i, e := range sliceLine {
//	//	if strings.ToUpper(e[0:1]) == e[0:1] {
//	//		if outputRhythm != nil {
//	//			if outputRhythm[0] == 6 {
//	//
//	//			}
//	//		}
//	//	} else {
//	//		if outputRhythm == nil {
//	//			outputRhythm = append(outputRhythm, 6)
//	//		} else {
//	//			if outputRhythm[i-1] == 1 {
//	//				outputRhythm = append(outputRhythm, 2)
//	//			} else {
//	//				if outputRhythm[i-1] == 2 {
//	//					outputRhythm = append(outputRhythm, 3)
//	//				}
//	//			}
//	//		}
//	//
//	//	}
//	//}
//}

// Based on important words
//time signature :6/8 (Neutral)
//the La zy Dog Jumps o Ver the Big Brown Dog
// 6 |1  2/3 4 | 1    2/3 4  5/6| 1   4   |  1
//3/4
// 6| 1  2  3   5     6|1   2    3   5    6
//if downbeats 1 and 5
// 6| 1  2-4  5  | 1   2-4 5  6|1    5     |1

//(Grand)
//4/4
//8-9 bars
//the La zy Dog Jumps o Ver the Big Brown Dog (11)
//    1     1    1      1        1   1     1
//randomize a higher number of 1s
//

//int perCent = s_Random.Next(0, 100);
//
//if (perCent < 50)               //  0 .. 49
//{
//// return Item of size 1
//}
//else if (perCent < 50 + 20)     // 50 .. 69
//{
//// return Item of size 5
//}
//else if (perCent < 50 + 20 + 5) // 70 .. 74
//{
//// return Item of size 10
//}
//...
