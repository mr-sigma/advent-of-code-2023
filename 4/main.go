package main

import (
  "fmt"
  "bufio"
  "math"
  "os"
  "regexp"
  "sort"
  "strings"
  "strconv"
)

var GamePrefix = regexp.MustCompile(`Card[\s]+[\d]+:`)

func ConvertAndSortSlice(s string) []int {
  res := []int{}

  scanner := bufio.NewScanner(strings.NewReader(s))
  scanner.Split(bufio.ScanWords)
  for scanner.Scan() {
      val, err := strconv.Atoi(scanner.Text())

      if err != nil {
        panic(err)
      }

      res = append(res, val)
  }

  sort.Ints(res)

  return res
}

func GetNums(line string) ([]int, []int) {
  // Remove the game header
  line = GamePrefix.ReplaceAllString(line, "")
  // split on | into winning numbers and ticket numbers
  splitGames := strings.Split(line, "|")

  // sort each slice
  winningNumbers := ConvertAndSortSlice(splitGames[0])
  ticketNumbers := ConvertAndSortSlice(splitGames[1])

  return winningNumbers, ticketNumbers
}

func CalcWinningMatches(line string) int {
  winningNumbers, ticketNumbers := GetNums(line)

  match := 0
  // iterate through winning numbers and check against all ticket numbers
  for _, tixNum := range(ticketNumbers) {
    for _, winNum := range(winningNumbers) {
      if tixNum == winNum {
        match ++
      }
    }
  }

  return match
}

func CalcWinningPoints(line string) int {
  pow := CalcWinningMatches(line)

  if pow == 0 {
    return 0
  } else {
    return int(math.Pow(2, float64(pow - 1)))
  }
}

func part1() {
  file, err := os.Open("input.txt")

  if err != nil {
    panic("problem opening file")
  }

  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)

  sum := 0

  for scanner.Scan() {
    sum += CalcWinningPoints(scanner.Text())
  }

  fmt.Println("Part 1 -- Sum of Winning Numbers:", sum)
}

// part two

func ParseFile(filename string) string {
  content, err := os.ReadFile(filename)

  if err != nil {
    panic("problem parsing input file")
  }

  return string(content)
}

func CalcNumTix(lines []string) int {
  if len(lines) == 1 {
    return 1
  } else {
    // calculate the matches in the top line
    fmt.Println(lines[0])
    matches := CalcWinningMatches(lines[0])

    if matches == 0 {
      return 1 + CalcNumTix(lines[1:])
    }

    // slice the lines and do recursion
    return 1 + CalcNumTix(lines[1:(matches + 1)])
  }
}

func part2() {
  lines := strings.Split(ParseFile("example_input.txt"), "\n")
  numTix := 0
  for i, _ := range(lines) {
    numTix += CalcNumTix(lines[i:])

  }
  // Read in input as array of strings
  // read a line, calculate number of matches
  // pass those lines to the same function with slice as argument

  fmt.Println("Part 2, number of tickets:", numTix)
}

func main() {
  // part1()
  part2()
}
