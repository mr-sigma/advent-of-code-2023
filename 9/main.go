package main

import (
  "fmt"
  "os"
  "strconv"
  "strings"
  "time"
)

func ReadLines(filename string) []string {
  content, err := os.ReadFile(filename)

  if err != nil {
    panic("problem parsing input file")
  }

  return strings.Split(string(content), "\n")
}

func AllZeroes(nums []int) bool {
  for _, x := range(nums) {
    if x != 0 {
      return false
    }
  }

  return true
}

func CalcDifferences(nums []int) []int {
  differences := []int{}
  for i := 1; i < len(nums); i++ {
    differences = append(differences, nums[i] - nums[i-1])
  }

  return differences
}

func ExtrapolateValue(nums []int) int {
  differences := CalcDifferences(nums)

  if AllZeroes(differences) {
    return nums[len(nums) - 1]
  }

  return nums[len(nums) - 1] + ExtrapolateValue(differences)
}

func ParseLine(line string) []int {
  nums := []int{}
  values := strings.Split(line, " ")

  for _, v := range(values) {
    val, err := strconv.Atoi(v)

    if err != nil {
      panic(err)
    }

    nums = append(nums, val)
  }

  return nums
}

func part1() {
  startTime := time.Now()
  lines := ReadLines("input.txt")
  sum := 0

  for _, line := range(lines) {
    if line == "" {
      break
    }
    nums := ParseLine(line)
    sum += ExtrapolateValue(nums)
  }

  fmt.Println("Part 1:", sum)
  fmt.Println("Took:", time.Since(startTime))
}

func ExtrapolateReverseValue (nums []int) int {
  differences := CalcDifferences(nums)

  if AllZeroes(differences) {
    return nums[0]
  }

  return nums[0] - ExtrapolateReverseValue(differences)
}

func part2() {
  startTime := time.Now()
  lines := ReadLines("input.txt")
  sum := 0

  for _, line := range(lines) {
    if line == "" {
      break
    }
    nums := ParseLine(line)
    sum += ExtrapolateReverseValue(nums)
  }

  fmt.Println("Part 2:", sum)
  fmt.Println("Took:", time.Since(startTime))
}

func main() {
  part1()
  part2()
}
