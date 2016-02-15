package main

import "fmt"

var currentID int

var games Games

// Give us some seed data
func init() {
	RepoCreateGame()
	RepoCreateGame()
}

func RepoFindGame(id int) Game {
	for _, g := range games {
		if g.ID == id {
			return g
		}
	}
	// return empty Game if not found
	return Game{}
}

func RepoCreateGame() Game {
	currentID++ //this is bad, I don't think it passes race condtions, but ok for example

	g := NewGame(currentID)
	games = append(games, g)

	return g
}

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

func RepoDestroyGame(id int) error {
	for i, g := range games {
		if g.ID == id {
			games = append(games[:i], games[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Game with id of %d to delete", id)
}
