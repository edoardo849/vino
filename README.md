# Vino: a Hangman-like game written in Go
Vino is a Hangman game that runs on the server and communicates with the client with a RESTful API. The game is a quick showcase of a Go API and should **NOT** be used in any production environment.

The words to guess are picked from a list of some of the best wines out there because... good wines are awesome!

| Method  | URL  | Description  |
|---|---|---|
| POST  | /games  |Start a new game   |
| GET  | /games  | Overview of all games  |
| GET  | /games/{id}  | JSON response that should at least include: **word**: representation of the word that is being guessed. Should contain dots for letters that have not been guessed yet (e.g. aw.so..) **tries_left**: the number of tries left to guess the word (starts at 11) **status**: current status of the game (busy/fail/success)   |
| POST  | games/{id}   | Guessing a letter, POST body: char=a |

## Rules
- Guessing a correct letter doesn’t decrement the amount of tries left
- Only valid characters are *a-­z* (lowercase)

## Requisites
- Go must be correctly installed on your system. Please check that you have both declared your GOPATH as well as your GOBIN environment variables.
- Git
- Curl or a UI-like tool for sending and receiving REST calls. I recommend [Postman](https://www.getpostman.com/).

## Instructions
Open up a new Terminal window and clone the repo:

```bash
# clone the repo
git clone git@github.com:edoardo849/vino.git

# enter in the vino directory
cd vino

#install dependencies
go get

#run the game
go run ./*.go
```

## Implementation Notes
- The directory structure is flat because of its simple implementation. For a production-ready environment you should organize your code into packages and folders
- The `POST /games/{id}` body must be formatted as an `application/json`
- The game does not register a repeated letter as failed attempt. Example:
if I've guessed so far "a..s.d...""
when I try to put another "a" then it's not considered incorrect. The logic is that it could be an error of intentions from the user's part.
- The router listens on port 8080: if you need to change it you can do it in the `main.go` file
- The game mimics a DB by keeping everything in-memory. This is **ABSOLUTELY** not a production-ready solution, of course...
- The data structure mimics a NRDB system, so it should be fairly easy to implement MongoDb logic for a more "robust" implementation
- The game automatically creates 2 games from the start to have something to play with.
