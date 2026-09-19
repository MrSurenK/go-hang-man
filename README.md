# go-hang-man
Full stack Hangman game written in Golang.

This game is meant to experiment with golang and learn the net/http module offeredn in go. It is intentionally kept simple in version 1.


Architecture: 

- Backend: Golang
- Frontend: React + Typescript
- Public APIs: 
    - Random Word API (https://random-word-api.herokuapp.com/home)


Design: 

- Go Backend will provide the user with a word list called with the public API
- User will have the option to choose difficulty between easy, medium and hard which will determine the commonality of the word
- Users will have 8 wrong attempts to guess the full word
- When user successfully guesses the word they will be awarded a point. (Points are not persistent and will be aggregated within the session only)

Things to note about v1: 
- There is no user authentication or authorization system in version 1
- There is no database connection


Potential Updates for V2: 

- Database implmentation
- User authentication 
- Persistent Highscore and Global scoreboard
