package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	_ "github.com/MusicMaker/rawMusicData"
)

func main() {
	fmt.Println("Type in the lyrics")
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	fmt.Println("length of lyric syllables")
	fmt.Println(length(line))
	for i := 0; i < 5; i++ {

	}
	rawMusicData.FirstNote()
	rawMusicData.SecondNote()
	rawMusicData.ThirdNote()
	fmt.Println("Final Output")
	fmt.Println(rawMusicData.FinalOutput)

}

func length(line string) int {
	sliceLyrics := strings.Split(line, " ")
	return len(sliceLyrics)
}
