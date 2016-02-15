# Hangman in Go with a REST API

| Method  | URL  | Description  |
|---|---|---|
| POST  | /games  |Start a new game   |
| GET  | /games  | Overview of all games  |
| GET  | /games/{id}  | JSON response that should at least include: word: representation of the word that is being guessed. Should contain dots for letters that have not been guessed yet (e.g. aw.so..) tries ​_left: the number of tries left to guess the word (starts at 11) status: current status of the game (busy/fail/success)   |
| POST  | games/{id}   | Guessing a letter, POST body: char=a |


- Guessing a correct letter doesn’t decrement the amount of tries left
- Only valid characters are *a­z* (lowercase)

## Running the game:
```bash
go run ./*.go
```

## Notes
The game does not register a repeated letter as failed attempt. Example:
a_sad_ _ _

If I put another "a" then it's not incorrect. The logic is that it could be an error of intentions from the user part.
