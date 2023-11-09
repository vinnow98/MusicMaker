package main

import (
	"bufio"
	"fmt"
	"music/logic"
	"net/http"
	"os"
	"strings"

	_ "github.com/vinnow98/MusicMaker/rawMusicData"
)

func main() {
	http.HandleFunc("/",handlerFunc)
	http.HandleFunc("/test",anotherHandler)
	http.ListenAndServe("",nil)
	fmt.Println("Instructions:")
	fmt.Println("Type in the song lyrics. Syllables should be separated by a space. If the syllable is important, capitalise it.")
	// the Lord is My she Pard There is No thing I shall Want
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	//line := "the La zy Dog Jumps o Ver the Big Brown Dog"
	fmt.Println("length of lyric syllables")
	fmt.Println(length(line))
	//characteristic := 1
	//logic.Characteristics(characteristic)
	logic.MainMelodyLogic(length(line))
	logic.MainRhythmLogic(line)
	fmt.Println("Line:")
	fmt.Println(line)
	fmt.Println("Result:")
	fmt.Println("Melody:", logic.FinalMelodyOutput)
	fmt.Println("Rhythm:", logic.FinalRhythmOutput)

}

func length(line string) int {
	sliceLyrics := strings.Split(line, " ")
	return len(sliceLyrics)
}

func handlerFunc(w http.ResponseWriter , r *http.Request) {
 fmt.Fprint(w, "<h1>Hello!</h1>")
}

func anotherHandler(w http.ResponseWriter , r *http.Request) {
 fmt.Fprint(w, "<h1>this is another</h1>")
}