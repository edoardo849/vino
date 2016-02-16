package main

import "fmt"

var currentID int

var games Games

// Give us some seed data
func init() {
	RepoCreateGame()
	RepoCreateGame()
}

//RepoFindGame is the "R" in CRUD
func RepoFindGame(id int) Game {
	for _, g := range games {
		if g.ID == id {
			return g
		}
	}
	// return empty Game if not found
	return Game{}
}

//RepoCreateGame is the "C" in CRUD: it stores the games
//inside an in-memory struct.
func RepoCreateGame() Game {
	currentID++ //this is bad, I don't think it passes race condtions, but ok for example

	g := NewGame(currentID)
	games = append(games, g)

	return g
}

//RepoUpdateGame is the "U" in CRUD: it deletes the old
//reference in the struct and replaces with a new one
func RepoUpdateGame(g Game) error {
	for i, gi := range games {
		if gi.ID == g.ID {

			// Delete the old Game
			games = append(games[:i], games[i+1:]...)
			// https://github.com/golang/go/wiki/SliceTricks
			games = append(games, Game{})
			copy(games[i+1:], games[i:])
			games[i] = g

			return nil
		}
	}
	return fmt.Errorf("Could not find Game with id of %d to delete", g.ID)
}

//RepoDestroyGame is the "D" in CRUD: currently not implemented
//elsewhere
func RepoDestroyGame(id int) error {
	for i, g := range games {
		if g.ID == id {
			games = append(games[:i], games[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Game with id of %d to delete", id)
}
