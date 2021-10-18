package src

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRecite(t *testing.T) {

	// Read file lyric and return array / slice
	t.Run("Read and print lyrics", func(t *testing.T) {

		sing := TwinkleSong{}
		sing.Read()
		lyrics := sing.Print()
		expected := []string{
			"Twinkle, twinkle, little star,",
			"How I wonder what you are!",
			"Up above the world so high,",
			"Like a diamond in the sky.",
			"Twinkle, twinkle, little star,",
			"How I wonder what you are!",
		}

		if !reflect.DeepEqual(lyrics, expected) {
			t.Fail()
		}
	})

	// Print recite with param 1 return string
	t.Run("Recite by 1", func(t *testing.T) {
		sing := TwinkleSong{}
		sing.Read()
		recite := sing.Recite(1)
		expected := "This is Twinkle, twinkle, little star"

		fmt.Println(recite)

		if recite != expected {
			t.Fail()
		}
	})

	// Print recite with param 2 return string
	t.Run("Recite by 2", func(t *testing.T) {
		sing := TwinkleSong{}
		sing.Read()
		recite := sing.Recite(2)
		expected := "This is How I wonder what you are and Twinkle, twinkle, little star"

		fmt.Println(recite)

		if recite != expected {
			t.Fail()
		}
	})

	// Print recite with param 4 return string
	t.Run("Recite by 4", func(t *testing.T) {
		sing := TwinkleSong{}
		sing.Read()
		recite := sing.Recite(4)
		expected := "This is Like a diamond in the sky and Up above the world so high and How I wonder what you are and Twinkle, twinkle, little star"

		fmt.Println(recite)

		if recite != expected {
			t.Fail()
		}
	})

	// Print recite with empty return string
	t.Run("Recite by empty param / random", func(t *testing.T) {
		sing := TwinkleSong{}
		sing.Read()
		recite := sing.Recite()
		expected := "This is Like a diamond in the sky and Up above the world so high and How I wonder what you are and Twinkle, twinkle, little star"
		if recite != expected {
			t.Fail()
		}
	})

}
