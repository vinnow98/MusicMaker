package main

import (
	"fmt"
	"music/logic"
	"strings"

	_ "github.com/vinnow98/MusicMaker/rawMusicData"
)

func main() {
	//fmt.Println("Type in the lyrics")
	//in := bufio.NewReader(os.Stdin)
	//line, _ := in.ReadString('\n')
	line := "the La zy Dog Jumps o Ver the Big Brown Dog"
	fmt.Println("length of lyric syllables")
	fmt.Println(length(line))
	//characteristic := 1
	//logic.Characteristics(characteristic)
	logic.MainMelodyLogic(length(line))
	logic.MainRhythmLogic(line)
	fmt.Println("Final Output")
	fmt.Println(line)
	fmt.Println("Melody:", logic.FinalMelodyOutput)
	fmt.Println("Rhythm:", logic.FinalRhythmOutput)

}

func length(line string) int {
	sliceLyrics := strings.Split(line, " ")
	return len(sliceLyrics)
}
