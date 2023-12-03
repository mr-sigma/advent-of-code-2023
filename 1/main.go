package main

import (
  "os"
  "fmt"
  "bufio"
  "unicode/utf8"
  "regexp"
  "strconv"
  "strings"
)

// returns the first character corresponding to the digits 0-9 it encounters
// on a line
func getFirstDigit(line string) string {
  for _, c := range(line) {
    if c > 47 && c < 58 {
      return string(c)
    }
  }
  return "0"
}

// returns the first character corresponding to the digits 0-9 it encounters
// on a line
func getLastDigit(line string) string {
  lineLength := utf8.RuneCountInString(line)

  for i := lineLength - 1; i > -1; i-- {
    if line[i] > 47 && line[i] < 58 {
      return string(line[i])
    }
  }

  return "0"
}

func partOne() {
  // iterate through the string forward to find a digit
  // iterate through the string backwards to find a digit
  // concat the digits
  // convert to int
  // add to sum
  // read in the input
  file, err := os.Open("part1_input.txt")

  if err != nil {
    fmt.Println("Problem parsing input")
  }

  count := 0

  scanner := bufio.NewScanner(file)

  scanner.Split(bufio.ScanLines)

  // iterate through lines
  for scanner.Scan() {
    text := scanner.Text()

    firstDigit := getFirstDigit(text)
    lastDigit := getLastDigit(text)

    val, err := strconv.Atoi((firstDigit + lastDigit))

    if err != nil {
      return 
    }
    count += val

  }

  file.Close()

  fmt.Println("Part 1:", count)

}

// Part 2
var DigitMap = map[string]string{
  "zero": "0",
  "one": "1",
  "two": "2",
  "three": "3",
  "four": "4",
  "five": "5",
  "six": "6",
  "seven": "7",
  "eight": "8",
  "nine": "9",
}

var Digit = regexp.MustCompile(`(zero|one|two|three|four|five|six|seven|eight|nine|[\d])`)
var DigitReverse = regexp.MustCompile(`(orez|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin|[\d])`)

func ParseDigits(firstDigit, lastDigit string) int {
  val, err := strconv.Atoi((ParseDigit(firstDigit) + ParseDigit(lastDigit)))

  if err != nil {
    panic("couldn't parse digits")
  }

  return val
}

func ParseDigit(digit string) string {
  if utf8.RuneCountInString(digit) > 1 {
    return DigitMap[digit]
  } else {
    return digit
  }
}

func ReverseString(in string) string {
    var sb strings.Builder
    runes := []rune(in)
    for i := len(runes) - 1; 0 <= i; i-- {
        sb.WriteRune(runes[i])
    }
    return sb.String()
}

func GetDigits(line string) int {
  resultString := ""

  resultString += ParseDigit(Digit.FindString(line))

  revLine := ReverseString(line)

  secondMatch := DigitReverse.FindString(revLine)

  if len(secondMatch) == 1 {
    resultString += secondMatch
  } else {
    resultString += ParseDigit(ReverseString(secondMatch))
  }

  val, err := strconv.Atoi(resultString)

  if err != nil {
    panic("problem converting result in GetDigits")
  }

  return val
}

func partTwo() {
  file, err := os.Open("input.txt")

  if err != nil {
    panic("Problem parsing input")
  }

  count := 0
  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)

  // iterate through lines
  for scanner.Scan() {
    digits := GetDigits(scanner.Text())

    count +=  digits
  }

  fmt.Println("Part 2:", count)
}

func main() {
  partOne()
  partTwo()
}
