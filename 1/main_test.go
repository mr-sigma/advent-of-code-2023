package main

import "testing"

func TestParseDigit(t *testing.T) {
  tests := []struct {
    name string
    input  string
    want string
  }{
    {
      name: "single digit",
      input:  "9",
      want: "9",
    },
    {
      name: "name nine",
      input:  "nine",
      want: "9",
    },
  }
    for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      if got := ParseDigit(tt.input); got != tt.want {
        t.Errorf("ParseDigit(\"%s\") = \"%s\", want \"%s\"", tt.input, got, tt.want)
      }
    })
  }
}

func TestGetDigits(t *testing.T) {
  tests := []struct {
    name string
    input  string
    want int
  }{
    {
      name: "both digits",
      input:  "1staoheusnthaoeur4tsoe",
      want: 14,
    },
    {
      name: "both words",
      input:  "ninestaoheustehuneightastehu",
      want: 98,
    },
    {
      name: "first word second digit",
      input:  "ninestaoheustehun8astehu",
      want: 98,
    },
    {
      name: "first digit second word",
      input:  "sta9oheustehunasteighthu",
      want: 98,
    },
    {
      name: "both word overlap",
      input:  "oneight",
      want: 18,
    },
    {
      name: "one word",
      input:  "asoteuhaoennineththuhet",
      want: 99,
    },
    {
      name: "one digit",
      input:  "asoteuhaoen8ththuhet",
      want: 88,
    },
  }
    for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      if got := GetDigits(tt.input); got != tt.want {
        t.Errorf("GetDigits(\"%s\") = \"%v\", want \"%v\"", tt.input, got, tt.want)
      }
    })
  }
}
