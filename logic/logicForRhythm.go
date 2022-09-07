package logic

import (
	"fmt"
	"math/rand"
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

// Based on important words
//time signature :6/8 (Neutral)
//4-5bars (random)
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
