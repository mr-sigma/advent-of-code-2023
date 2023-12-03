package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "strconv"
  "regexp"
)

var GameTag = regexp.MustCompile(`Game\s[\d]+:\s+`)
var RedRegexp = regexp.MustCompile(`red`)
var GreenRegexp = regexp.MustCompile(`green`)
var BlueRegexp = regexp.MustCompile(`blue`)
var DigitsRegexp = regexp.MustCompile(`[\d]+`)
var MaxRed int = 12
var MaxGreen int = 13
var MaxBlue int = 14

func ParseGame(game string) []string {
  game = GameTag.ReplaceAllString(game, "" )
  return strings.Split(game, ";")
}

func RoundImpossible(round string) bool {
  impossible := false

  colors := strings.Split(round, ",")

  for _, color := range(colors) {
    digit, err := strconv.Atoi(DigitsRegexp.FindString(color)) 
    
    if err != nil {
      panic("couldn't match digit")
    }

    switch {
    case RedRegexp.MatchString(color):
      if digit > MaxRed {
        impossible = true
        break
      }
    case GreenRegexp.MatchString(color):
      if digit > MaxGreen {
        impossible = true
        break
      }
    case BlueRegexp.MatchString(color):
      if digit > MaxBlue {
        impossible = true
        break
      }
    default:
    }
  }

  return impossible
}

func GamePossible(game string) bool {
  rounds := ParseGame(game)
  possible := true

  for _, round := range(rounds) {
    if RoundImpossible(round) {
      possible = false
      break
    }
  }

  return possible
}

func part1() {
  file, err := os.Open("input.txt")

  if err != nil {
    panic("problem opening input file")
  }

  scanner := bufio.NewScanner(file)

  scanner.Split(bufio.ScanLines)

  gameCount := 0

  gameNum := 1

  for scanner.Scan() {
    if GamePossible(scanner.Text()) {
      gameCount += gameNum
    }

    gameNum ++
  }

  fmt.Println("Sum of possible game numbers: ", gameCount)

  file.Close()
}

func CalcMinCubes(round string, minRed, minBlue, minGreen int) (int, int, int) {
  colors := strings.Split(round, ",")

  for _, color := range(colors) {
    digit, err := strconv.Atoi(DigitsRegexp.FindString(color)) 
    
    if err != nil {
      panic("couldn't match digit")
    }

    switch {
    case RedRegexp.MatchString(color):
      if digit > minRed {
        minRed = digit
      }
    case GreenRegexp.MatchString(color):
      if digit > minGreen {
        minGreen = digit
      }
    case BlueRegexp.MatchString(color):
      if digit > minBlue {
        minBlue = digit
      }
    default:
    }
  }

  return minRed, minBlue, minGreen
}

func CalcGameCube(game string) int {
  rounds := ParseGame(game)

  red, blue, green := 0, 0, 0

  for _, round := range(rounds) {
    red, blue, green = CalcMinCubes(round, red, blue, green)
  }

  return red * blue * green

}

func part2() {
  file, err := os.Open("input.txt")

  if err != nil {
    panic("problem opening input file")
  }

  scanner := bufio.NewScanner(file)

  scanner.Split(bufio.ScanLines)

  powerSum := 0

  for scanner.Scan() {
    powerSum += CalcGameCube(scanner.Text())
  }

  fmt.Println("Sum of game cubes: ", powerSum)

  file.Close()
}

func main() {
  part1()
  part2()
}
