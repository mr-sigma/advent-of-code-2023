package main

import (
  "fmt"
  "bufio"
  "os"
  "regexp"
  "strings"
  "strconv"
  "time"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

var Map = regexp.MustCompile(`(\d+)\s+(\d+)\s+(\d+)`)
var Seeds = regexp.MustCompile(`seeds:\s+`)
var MapHeader = regexp.MustCompile(`\w+-\w+-\w+`)

func GetSeeds(line string) []int {
  seeds := []int{}

  line = Seeds.ReplaceAllString(line, "")

  scanner := bufio.NewScanner(strings.NewReader(line))
  scanner.Split(bufio.ScanWords)

  for scanner.Scan() {
    val, err := strconv.Atoi(scanner.Text())
    
    if err != nil {
      panic(err)
    }

    seeds = append(seeds, val)
  }

  return seeds
}

func ParseFile(filename string) string {
  content, err := os.ReadFile(filename)

  if err != nil {
    panic("problem parsing input file")
  }

  return string(content)
}

type Interval struct {
  Start int
  End int
}

type LineMap struct {
  From Interval
  To Interval
}

func (i Interval) Contains(num int) bool {
  return num >= i.Start && num <= i.End
}

func (lm LineMap) Transform(num int) int {
  if lm.From.Contains(num) {
    transformValue := num - lm.From.Start
    return lm.To.Start + transformValue
  } else {
    return num
  }
}

func (lm LineMap) ReverseTransform(num int) int {
  if lm.To.Contains(num) {
    transformValue := num - lm.To.Start
    return lm.From.Start + transformValue
  } else {
    return num
  }
}

func BuildLineMap(line string) LineMap {
  matches := Map.FindAllStringSubmatch(line, -1) 

  fromStart, err := strconv.Atoi(matches[0][2])
  if err != nil {
    panic(err)
  }

  toStart, err := strconv.Atoi(matches[0][1])
  if err != nil {
    panic(err)
  }

  interval, err := strconv.Atoi(matches[0][3])
  if err != nil {
    panic(err)
  }

  fromEnd := fromStart + interval
  toEnd := toStart + interval

  from := Interval{Start: fromStart, End: fromEnd}
  to := Interval{Start: toStart, End: toEnd}

  return LineMap{From: from, To: to}
}

// ReadLineMap returns the number after possible transform and whether the 
// value changed during the transform
func ReadLineMap(val int, line string) (int, bool) {
  if Map.MatchString(line) {
    lm := BuildLineMap(line)
    newVal := lm.Transform(val)
    return newVal, newVal != val
  } else {
    return val, false
  }
}

func WalkMaps(seed int, lines []string) int {
  transform := seed
  skip := false
  var changed bool

  for i, line := range(lines) {
    if skip {
      // check that line matches the next section
      if MapHeader.MatchString(line) {
        // toggle skip
        skip = false
      }
      fmt.Println("skipping line", i+2)
    } else {
      transform, changed = ReadLineMap(transform, line)
      fmt.Println("Line ", i + 2, "Map Value:", transform)
      if changed {
        skip = true
      }
    }
  }

  return transform
}

func part1() {
  startTime := time.Now()
  lines := strings.Split(ParseFile("input.txt"), "\n")

  seeds := GetSeeds(lines[0])
  minLocation := MaxInt
  var location int

  for _, seed := range(seeds) {
    fmt.Println("Seed: ", seed)
    location = WalkMaps(seed, lines[1:])
    fmt.Println("Location: ", location)
    fmt.Println("Current Min Location: ", minLocation)

    if location < minLocation {
      minLocation = location
    }
  }

  fmt.Println("Part 1 -- Min Location:", minLocation)
  fmt.Println("took", time.Since(startTime))
}

// part two
func GetSeedsPartTwo(line string) []Interval {
  seeds := []Interval{}

  line = Seeds.ReplaceAllString(line, "")

  scanner := bufio.NewScanner(strings.NewReader(line))
  scanner.Split(bufio.ScanWords)

  isLenInterval := false

  var startNum int
  var lenInterval int

  for scanner.Scan() {
    val, err := strconv.Atoi(scanner.Text())
    
    if err != nil {
      panic(err)
    }

    if isLenInterval {
      lenInterval = val
      seeds = append(seeds, Interval{Start: startNum, End: startNum + lenInterval - 1})
    } else {
      startNum = val
    }

    isLenInterval = !isLenInterval
  }

  return seeds
}

func ReverseWalkMaps(location int, lines []string) int {
  transform := location
  skip := false
  var changed bool

  for i := len(lines) - 1; i >= 0; i-- {
    if skip {
      // check that line matches the next section
      if MapHeader.MatchString(lines[i]) {
        // toggle skip
        skip = false
      }
      fmt.Println("skipping line", i+2)
    } else {
      transform, changed = ReverseReadLineMap(transform, lines[i])
      fmt.Println("Line ", i + 2, "Map Value:", transform)
      if changed {
        skip = true
      }
    }
  }

  return transform
}

// ReverseReadLineMap returns the number after possible transform and whether the 
// value changed during the transform
func ReverseReadLineMap(val int, line string) (int, bool) {
  if Map.MatchString(line) {
    lm := BuildLineMap(line)
    newVal := lm.ReverseTransform(val)
    return newVal, newVal != val
  } else {
    return val, false
  }
}

func SeedInSeeds(seed int, seeds []Interval) bool {
  for _, s := range(seeds) {
    if s.Contains(seed) {
      return true
    }
  }
  return false
}

func part2() {
  startTime := time.Now()
  lines := strings.Split(ParseFile("input.txt"), "\n")

  seeds := GetSeedsPartTwo(lines[0])

  location := 1

  for {
    seed := ReverseWalkMaps(location, lines[1:])
    if SeedInSeeds(seed, seeds) {
      break
    }
    location ++
  }

    fmt.Println("Part 2 -- Min Location:", location)
    fmt.Println("took", time.Since(startTime))
}

func main() {
  part1()
  part2()
}
