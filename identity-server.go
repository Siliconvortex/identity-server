package main

import (
  "fmt"
  "time"
  "encoding/json"
  "net/http"
  "github.com/dchest/uniuri"
)

func main() {
  http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    s, _ := json.Marshal(TokenGenerator())
    fmt.Fprint(w, string(s))
    return
  })

  http.ListenAndServe(":8080", nil)
}

const Token_Length = 256
const Token_Duration = time.Hour

type Token struct {
  Id string
  Expiration time.Time
}

func TokenGenerator() Token {
  answer := Token{
    Id : uniuri.NewLen(Token_Length),
    Expiration: time.Now().Add(Token_Duration).UTC(),
  }
  return answer
}
