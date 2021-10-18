package src

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type TwinkleSong struct {
	Lyrics []string
}

func (sing *TwinkleSong) Read() []string {

	fileBytes, err := ioutil.ReadFile("../lyric.txt")

	if err != nil {
		fmt.Errorf(err.Error())
	}

	sing.Lyrics = strings.Split(string(fileBytes), "\n")

	return strings.Split(string(fileBytes), "\n")

}

func (sing *TwinkleSong) Print() []string {
	return sing.Lyrics
}

func (sing *TwinkleSong) Process(index int, first int, lengthLyrics int, result string) string {

	if index == 0 {
		return result
	}

	lyric := sing.Lyrics[index-1]
	if index < lengthLyrics {
		if index == first {
			result += "This is " + lyric[0:len(lyric)-1]
		} else {
			result += " and " + lyric[0:len(lyric)-1]
		}
	}

	return sing.Process(index-1, lengthLyrics, index, result)
}

func randomize(max int) int {
	rand.Seed(time.Now().UnixNano())
	min := 1

	return rand.Intn(max-min+1) + min
}

func (sing *TwinkleSong) Recite(times ...int) string {
	counter := 0
	for _, n := range times {
		counter += n
	}

	if counter == 0 {
		counter = randomize(len(sing.Lyrics))
		return sing.Process(counter, counter, len(sing.Lyrics), "")
	}

	return sing.Process(counter, counter, len(sing.Lyrics), "")
}
