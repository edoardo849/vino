package main

import (
	"math/rand"
	"strings"
	"time"
)

const hiddenChar = "."

//Game is the model for our game object
type Game struct {
	ID        int    `json:"id"`
	Word      string `json:"word"`
	TriesLeft int    `json:"tries_left"`
	Status    string `json:"status"` //busy, fail, success

	// For simplicity we will store the challenge here as a hidden value because
	// we are using a data structure for a NRDB. For a RDB-like system
	// we should could have had created an associative table like:
	// ```go
	// type GameToChallenge struct {
	// 		GameID	int
	// 		ChallengeID	string
	// }
	//```
	Challenge string `json:"-"` // do not show the challenge!
}

//NewGame is a constructor-like factory function
func NewGame(ID int) Game {
	var g Game
	g.ID = ID
	g.TriesLeft = 11
	g.Status = "busy"
	g.Begin()

	return g
}

//Begin : randomly selects a new Word to guess for the current game
//and hides the word
func (g *Game) Begin() {

	//generate random seed off UTC time
	rand.Seed(time.Now().UTC().UnixNano())

	//words that are available to use
	//words are wines because... reasons!
	words := []string{
		"merlot",
		"cabernet",
		"pinot",
		"barbera",
		"barolo",
		"nebbiolo",
		"chianti",
		"nerodavola",
		"gewurtztraminer",
		"dolcetto",
		"sauvignon",
		"chardonnay",
		"brunello",
		"amarone",
	}

	//generate random integer for game
	randInt := rand.Intn(int(len(words)))

	//select a random word
	randWord := words[randInt]

	//store the selected word in the Game object
	//and masks it for the user
	g.Word = hideWord(randWord)
	g.Challenge = randWord

}

//GuessLetter is the actual Hangman game implementation
func (g *Game) GuessLetter(l Letter) {

	// Do not bother to execute code if the game is not busy
	if g.Status != "busy" {
		return
	}

	//Keeps track if the guessed letter is included
	//into one of the challenge's
	isCorrect := false
	guessedLetter := l.Char

	//create two slices, one to keep track of the challenge word (split into chars)
	challengeCharSlice := strings.Split(g.Challenge, "")

	//another to keep track of the words already guessed
	wordCharSlice := strings.Split(g.Word, "")

	for key, value := range challengeCharSlice {
		if guessedLetter == value {
			// track here a correct answer
			isCorrect = true

			//replace the guessed letter in the masked word
			wordCharSlice[key] = guessedLetter
		}
	}

	//count the revealed words
	correctWords := 0
	for _, value := range wordCharSlice {
		if value != hiddenChar {
			correctWords++
		}
	}

	//Update the masked word
	g.Word = strings.Join(wordCharSlice, "")

	// Guessing a correct letter doesn’t decrement the amount of tries left
	if g.TriesLeft > 0 && !isCorrect {
		g.TriesLeft--
	}

	if g.TriesLeft == 0 {
		g.Status = "fail"
	} else if correctWords == len(challengeCharSlice) {
		g.Status = "success"
	}
}

func hideWord(word string) string {
	var w string
	//iterate over words and split them into characters
	for i := 0; i < int(len(word)); i++ {
		w += hiddenChar
	}

	return w
}

//Games is a list of games
type Games []Game
