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
	line := "the lazy dog jumps over the big brown dog"
	fmt.Println("length of lyric syllables")
	fmt.Println(length(line))
	characteristic := 1
	logic.Characteristics(characteristic)
	numberOfNotes := 10
	logic.FirstNote()
	logic.FirstHalfNotes(numberOfNotes)
	logic.LastNote()

	fmt.Println("Final Output")
	fmt.Println(logic.FinalOutput)

}

func length(line string) int {
	sliceLyrics := strings.Split(line, " ")
	return len(sliceLyrics)
}
