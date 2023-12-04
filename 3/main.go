package main

import (
  "fmt"
  "os"
  "regexp"
  "strconv"
  "strings"
)

var Numbers = regexp.MustCompile(`\d+`)
var Symbol = regexp.MustCompile(`[^\s\d\w.]`)
var Gear = regexp.MustCompile(`\*`)
var LineMultiplication = regexp.MustCompile(`(\d+)\*(\d+)`)

func SymbolsInLineRange(line string, startIndex, endIndex int) bool {
  if Symbol.MatchString(line[startIndex:endIndex]) {
    return true
  }

  return false
}

func AnyAdjacentSymbols(lines *[]string, lineNum int, indices[]int) bool {
  startIndex := indices[0]
  endIndex := indices[1]

  if startIndex > 0 {
    startIndex --
  }

  if endIndex < len((*lines)[lineNum]) {
    endIndex ++
  }

  // check left and right
  if SymbolsInLineRange((*lines)[lineNum], startIndex, endIndex) {
    return true
  }
  if lineNum != 0 {
    // check line above
    if SymbolsInLineRange((*lines)[lineNum - 1], startIndex, endIndex) {
      return true
    }
  }

  if lineNum != len(*lines) - 1 {
    // check line below
    if SymbolsInLineRange((*lines)[lineNum + 1], startIndex, endIndex) {
      return true
    }
  }

  return false
}

func PartNumbersWithAdjacentSymbols(lines *[]string, lineNum int, indicesList [][]int) int {
  sum := 0

  for _, indices := range(indicesList) {
    if AnyAdjacentSymbols(lines, lineNum, indices) {
      val, err := strconv.Atoi((*lines)[lineNum][indices[0]: indices[1]])

      if err != nil {
        panic("problem converting string to integer in PartNumbersWithAdjacentSymbols")
      }

      sum += val
    }
  }

  return sum
}

func part1() {
  content, err := os.ReadFile("input.txt")

  if err != nil {
    panic("problem parsing input file")
  }

  lines := strings.Split(string(content), "\n")

  // Drop empty last line
  lines = lines[:len(lines) - 1]

  sum := 0

  for lineNum, line := range(lines) {
    indices := Numbers.FindAllStringSubmatchIndex(line, -1)
    if len(indices) != 0 {
      sum += PartNumbersWithAdjacentSymbols(&lines, lineNum, indices)
    }
  }

  fmt.Println("Sum of part numbers: ", sum)
}

// Part 2
func ParseFile(filename string) string {
  content, err := os.ReadFile(filename)

  if err != nil {
    panic("problem parsing input file")
  }

  return string(content)
}

func GearRatiosSameLine(line string) int {
  gearRatios := 0

  if LineMultiplication.MatchString(line) {
    matches := LineMultiplication.FindAllStringSubmatch(line, -1)
    fmt.Println("matches", matches)

    for _, match := range(matches) {
      val1, err := strconv.Atoi(match[1])
      
      if err != nil {
        panic("problem converting string to int in PartsNumbersWithAdjancentGears")
      }

      val2, err := strconv.Atoi(match[2])

      if err != nil {
        panic("problem converting string to int in PartsNumbersWithAdjancentGears")
      }

      gearRatios += val1 * val2
    }
  }

  return gearRatios
}

func NumberOverlapsIndex(numStart int, numEnd int, gearIndex int) bool {
  return numEnd == gearIndex || // left
  numStart == gearIndex + 1 || // right
  (gearIndex > numStart && gearIndex <= numEnd) // overlaps on another line
}

func FindAdjacentNumbersByLine(line string, gearIndex int) []int {
  nums = make([]int, 0, 2) // can only ever be two adjacent nums per line
  matches := Numbers.FindAllStringIndex(line)
  
  for _, match := range(matches) {
    if NumberOverlapsIndex(match[0], match[1]) {
      num, err := strconv.Atoi(line[match[0]:match[1]])
      if err != nil {
        panic("problem converting string in FindAdjacentNumbersByLine")
      }

      append(nums, num)
    }
  }

  return nums
}

func AdjacentNumbersToGear(lines *[]string, lineNum int, gearIndex int) []int {
  nums := make([]int, 0, 6)
  if lineNum > 0 {
  }

  current := Numbers.FindAllStringMatchIndex((*lines)[lineNum])

  if lineNum < len((*lines)) - 1 {
    lower := Numbers.FindAllStringMatchIndex((*lines)[lineNum + 1])
  }
}

func PartNumbersWithAdjacentGears(lines *[]string, lineNum int, indices [][]int) int {
  gearRatios := 0

  // loop through the gears on a line
  for _, i := range(indices) {
    // search for numbers touching the gear
    nums := AdjacentNumbersToGear(lines, lineNum, i)
  }

  // Locate a gear
  // Find numbers which are "touching" the gear
    // Upper
    // Lower
    // Same line
  // find the product of all the combinations of the numbers touching the gear

  // // check the current line
  // gearRatios += GearRatiosSameLine((*lines)[lineNum])


  // // check the current line and above line
  // if lineNum != 0 {
  //   GearRatiosSplitLines()
  // }

  return gearRatios
}

func part2() {

  lines := strings.Split(ParseFile("input.txt"), "\n")

  // Drop empty last line
  lines = lines[:len(lines) - 1]

  sum := 0

  for lineNum, line := range(lines) {
    // find all gears
    indices := Gear.FindAllStringSubmatchIndex(line, -1)
    // fmt.Println("gear indices:",indices, lineNum )
    if len(indices) != 0 {
      sum += PartNumbersWithAdjacentGears(&lines, lineNum, indices)
    }
  }

  fmt.Println("Sum of gear numbers: ", sum)
}

func main() {
  // part1()
  part2()
}
