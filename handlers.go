package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Vino: the wine's hangman game!\n")
}

func GameIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(games); err != nil {
		panic(err)
	}
}

func GameShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var gameID int
	var err error
	if gameID, err = strconv.Atoi(vars["id"]); err != nil {
		panic(err)
	}
	game := RepoFindGame(gameID)
	if game.ID > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(game); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}

func GameCreate(w http.ResponseWriter, r *http.Request) {
	g := RepoCreateGame()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(g); err != nil {
		panic(err)
	}
}

func GameGuessLetter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var gameID int
	var letter Letter
	var err error

	if gameID, err = strconv.Atoi(vars["id"]); err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &letter); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(jsonErr{Code: 422, Text: "Malformed JSON"}); err != nil {
			panic(err)
		}
		return
	}

	if err := letter.isValid(); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(jsonErr{Code: 422, Text: err.Error()}); err != nil {
			panic(err)
		}
		return
	}

	game := RepoFindGame(gameID)

	if game.ID > 0 {

		game.GuessLetter(letter)
		RepoUpdateGame(game)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(game); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Game Not Found"}); err != nil {
		panic(err)
	}
}
