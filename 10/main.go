package main

import (
  "fmt"
  "math"
  "os"
  "strings"
)

type Point struct {
  X int
  Y int
}

func GetLines(filename string) []string {
  content, err := os.ReadFile(filename)

  if err != nil {
    panic(err)
  }

  return strings.Split(string(content), "\n")
}

func FindStartingPoint(lines []string) Point {
  var startingPoint Point
  exit := false

  for i, line := range(lines) {
    if exit {
      break
    }

    for j, char := range(line) {
      if char == 'S' {
        startingPoint = Point{X: j, Y: i}
        exit = true
        break
      }
    }
  }

  return startingPoint
}

func (p Point) Eq(op Point) bool {
  return p.X == op.X && p.Y == op.Y
}

func (p Point) Valid() bool {
  return p.X > -1 && p.Y > -1
}

func NextPipe(lines []string, currentPoint Point, previousPoint Point, dirs string) (Point, bool) {
  var nextPoint Point
  for _, char := range(dirs) {
    switch char {
    case 'U':
      nextPoint = Point{X: currentPoint.X, Y: currentPoint.Y - 1}
    case 'D':
      nextPoint = Point{X: currentPoint.X, Y: currentPoint.Y + 1}
    case 'L':
      nextPoint = Point{X: currentPoint.X - 1, Y: currentPoint.Y}
    case 'R':
      nextPoint = Point{X: currentPoint.X + 1, Y: currentPoint.Y}
    default:
      panic("invalid direction instruction")
    }

    if !nextPoint.Eq(previousPoint) && lines[nextPoint.Y][nextPoint.X] != '.' && nextPoint.Valid() {
      return nextPoint, true
    }
  }

  return Point{X: -1, Y: -1}, false
}

// returns the Point of the next pipe we can travel to
func FindAdjacentPipe(lines []string, currentPoint Point, previousPoint Point) Point {
  var nextPoint Point
  var ok bool

  // check the current symbol and then determine which directions to check

  currentPipe := lines[currentPoint.Y][currentPoint.X]

  switch currentPipe {
  case '|':
    nextPoint, ok = NextPipe(lines, currentPoint, previousPoint, "UD")
  case '-':
    nextPoint, ok = NextPipe(lines, currentPoint, previousPoint, "LR")
  case 'L':
    nextPoint, ok = NextPipe(lines, currentPoint, previousPoint, "UR")
  case 'J':
    nextPoint, ok = NextPipe(lines, currentPoint, previousPoint, "UL")
  case '7':
    nextPoint, ok = NextPipe(lines, currentPoint, previousPoint, "DL")
  case 'F':
    nextPoint, ok = NextPipe(lines, currentPoint, previousPoint, "RD")
  case 'S':
    nextPoint, ok = NextPipe(lines, currentPoint, previousPoint, "UDLR")
  default:
    panic("Hit ground symbol .")
  }

  if !ok {
    panic("failed to find next point")
  }

  return nextPoint
}

// returns the number of steps it takes to reach the starting point again
func WalkPipes(lines []string, startingPoint Point) (int, []Point) {
  points := []Point{startingPoint}
  currentPoint := startingPoint
  previousPoint := startingPoint
  exit := false

  steps := 0

  for {
    if exit {
      break
    }

    nextPoint := FindAdjacentPipe(lines, currentPoint, previousPoint)
    fmt.Println(nextPoint, string(lines[nextPoint.Y][nextPoint.X]))

    steps ++
    previousPoint = currentPoint
    currentPoint = nextPoint

    if currentPoint.Eq(startingPoint) {
      exit = true
    }

    points = append(points, currentPoint)

  }

  return steps, points
}

func part1() {
  lines := GetLines("input.txt")
  
  // find start
  startingPoint := FindStartingPoint(lines)
  fmt.Println("Starting point:", startingPoint)
  // look at each adjacent tile and find a way to walk
  // repeat until back at start
  // furthest point is entire trip/2

  steps, _ := WalkPipes(lines, startingPoint)

  fmt.Println("Part 1:", steps/2)
}

// Part 2
func ShoelaceArea(points []Point) int {
  // https://en.wikipedia.org/wiki/Shoelace_formula#Shoelace_formula
  sum := 0
  var j int

  for i := 0; i < len(points); i++ {
    switch i {
    case len(points) - 1:
      j = 0
    default:
      j = i + 1
    }

    sum += (points[i].X * points[j].Y) - (points[j].X * points[i].Y)
  }

  return int(math.Abs(float64(sum / 2)))
}

func FindInteriorPoints(area int, steps int) int {
  // A = i + 1/2 * b - 1
  // A - 1/2 * b + 1 = i
  // https://en.wikipedia.org/wiki/Pick%27s_theorem

  return area - (steps / 2) + 1

}

func part2() {
  lines := GetLines("input.txt")
  
  // find start
  startingPoint := FindStartingPoint(lines)
  fmt.Println("Starting point:", startingPoint)
  // look at each adjacent tile and find a way to walk
  // repeat until back at start
  // furthest point is entire trip/2

  b, points := WalkPipes(lines, startingPoint)
  area := ShoelaceArea(points)

  i := FindInteriorPoints(area, b)
  fmt.Println("Interior Points:", i)
}

func main() {
  // part1()
  part2()
}
