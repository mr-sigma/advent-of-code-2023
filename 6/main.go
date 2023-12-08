package main

import (
  "fmt"
  "bufio"
  "os"
  "math"
  "strconv"
  "strings"
)

// velocity increases at 1mm/s/s
// t * t is just the equation?
// tHold * (tTotal - tHold) > distance
// -tHold * tHold - tHold * tTotal > distance
// tHold = x
// -x^2 -x * tTotal - distance = 0
// -x^2 - xt - d = 0
// Two roots, max and min button hold values
// Calculate the integer gap between the two

func Quadratic(t int, d int) (minRoot, maxRoot float64) {
  minRoot = (float64(t) - math.Sqrt(float64(t * t) - 4. * float64(d)))/2.
  maxRoot = (float64(t) + math.Sqrt(float64(t * t) - 4. * float64(d)))/2.
  return minRoot, maxRoot
}

func DistBetweenRoots(min, max float64) int {
  if math.Mod(max, 1) != 0 && math.Mod(min, 1) != 0 {
    return int(math.Floor(max) - math.Ceil(min)) + 1
  } else {
    return int(math.Ceil(max)) - int(math.Floor(min)) - 1
  }
}

func ParseInput(filename string) ([]int, []int) {
  file, err := os.Open(filename)
  if err != nil {
    panic("Problem parsing input")
  }

  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)

  tSlice := []int{}
  dSlice := []int{}

  lineNum := 0

  for scanner.Scan() {
    line := strings.ReplaceAll(scanner.Text(), "Time:", "")
    line = strings.ReplaceAll(line, "Distance:", "")
    nums := strings.Fields(line)
    for _, num := range(nums) {
      val, err := strconv.Atoi(num)

      if err != nil {
        panic(err)
      }

      if lineNum == 0 {
        tSlice = append(tSlice, val)
      } else {
        dSlice = append(dSlice, val)
      }
    }

    lineNum++
  }

  return tSlice, dSlice
}

func part1() {
  tSlice, dSlice := ParseInput("input.txt")

  result := 1

  for i, val := range(tSlice) {
    result *= DistBetweenRoots(Quadratic(val, dSlice[i]))
  }

  fmt.Println("Part1:", result)
}

func ParseInput2(filename string) (int, int) {
  file, err := os.Open(filename)
  if err != nil {
    panic("Problem parsing input")
  }

  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)

  lineNum := 0
  var tVal int
  var dVal int

  for scanner.Scan() {
    line := strings.ReplaceAll(scanner.Text(), "Time:", "")
    line = strings.ReplaceAll(line, "Distance:", "")
    nums := strings.Fields(line)
    num := strings.Join(nums, "")
    val, err := strconv.Atoi(num)

    if err != nil {
      panic(err)
    }

    if lineNum == 0 {
      tVal = val
    } else {
      dVal = val
    }

    lineNum++
  }

  return tVal, dVal
}

func part2() {
  tVal, dVal := ParseInput2("input.txt")

  result := DistBetweenRoots(Quadratic(tVal, dVal))

  fmt.Println("Part2:", result)
}

func main() {
  part1()
  part2()
}
