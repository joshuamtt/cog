package main

import (
  "fmt"
  "os"
  "io/ioutil"
  "log"
  "regexp"
)

type TOKEN_TYPE int

const (
  TOKEN_STRING TOKEN_TYPE = iota
  TOKEN_NUMBER
  TOKEN_LEFT_BRACE
  TOKEN_RIGHT_BRACE
  TOKEN_LEFT_PAREN
  TOKEN_RIGHT_PAREN
  TOKEN_KEYWORD
  TOKEN_IDENTIFIER
  TOKEN_SEMICOLON
  TOKEN_EOF
)

type Token struct {
  token_type    TOKEN_TYPE
  value         string
  regex         *regexp.Regexp
  position      int
}

type Lexer struct {
  file_input    string
  position      int
  regex         []string // #might remove
}

type Token_Regex struct {
  tok_type      TOKEN_TYPE
  tok_regex     *regexp.Regexp
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

  var token_regex = []Token_Regex{
    {TOKEN_STRING, regexp.MustCompile(`\s`)},
    {TOKEN_NUMBER, regexp.MustCompile(`\s`)},
    {TOKEN_LEFT_BRACE, regexp.MustCompile(`{`)},
    {TOKEN_RIGHT_BRACE, regexp.MustCompile(`}`)},
    {TOKEN_LEFT_PAREN, regexp.MustCompile(`\(`)},
    {TOKEN_RIGHT_PAREN, regexp.MustCompile(`\)`)},
    {TOKEN_KEYWORD, regexp.MustCompile(``)},
    {TOKEN_IDENTIFIER, regexp.MustCompile(`[a-zA-Z]\w*`)},
    {TOKEN_RETURN, regexp.MustCompile(`return`)},
  }


  tokens := lex(string(content))
  fmt.Println(tokens)
}
// ##Update this to return tokens:
func lex(file string) []string {
  fmt.Println(file)

  // Initial regex structure for getting tokens
  // #Don't think this is the way to go. How should I relate the regex patterns, to tokens?
  //   - I Think a global array of "Token" structure, where each element is the token_type,
  //     and it has its compiled regex
  //
  token_regex_strings := []string{"{", "}", `\(`, `\)`, `[a-zA-Z]\w*`, "[0-9]+"}
  //

  var token_regex_objs []*regexp.Regexp

  for i := 0; i < len(token_regex_strings); i++  {
    re, err := regexp.Compile(token_regex_strings[i])
    if err != nil {
      fmt.Println("Error compiling regex...", err)
    }

    token_regex_objs = append(token_regex_objs, re)
  }

  // We have each of the regex objects compiled now.
  // The question: How should the tokens be organized?
  // If I use the regex on the file (string), what will be returned is the list of all of them...?

  for j := 0; j < len(token_regex_objs); j++ {
    tok := token_regex_objs[j].FindAllString(file, 1)
    fmt.Println(tok)
  }


  // fmt.Println(token_regex_objs)

  tokens := []string{"i"}

  return tokens

}
