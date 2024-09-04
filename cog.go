package main

import (
  "fmt"
  "os"
  "io/ioutil"
  "log"
)

type TOKEN_TYPE int iota
const (
  TOKEN_STRING TOKEN_TYPE
  TOKEN_NUMBER
  TOKEN_LEFT_BRACE
  TOKEN_RIGHT_BRACE
  TOKEN_LEFT_PAREN
  TOKEN_RIGHT_PAREN
  TOKEN_KEYWORD
  TOKEN_IDENTIFIER
  TOKEN_SEMICOLON
)

type Token struct {
  type    Token_Type
  value   string
}

type Lexer struct {
  file_input    string
  position      int
  

}

func main() {
  if len(os.Args) < 2 {
    fmt.Println("ERROR: expected a file input.")
    os.Exit(1)
  }

  content, err := ioutil.ReadFile(os.Args[1])
  if err != nil {
    log.Fatal(err)
  }

  lex(string(content))
}

func lex(file string) []string {
  fmt.Println(file)

  tokens := []string{"i"}

  return tokens

}
