package main

import (
  "fmt"
  "math"
  "os"
  "regexp"
  "strings"
  "time"
)

var Instructions = regexp.MustCompile(`[RL]+`)

type Node struct {
  Name string
  Left string
  Right string
}

func MakeNode(line string) Node {
  line = strings.ReplaceAll(line, "(", "")
  line = strings.ReplaceAll(line, ")", "")
  line = strings.ReplaceAll(line, " =", "")
  line = strings.ReplaceAll(line, ",", "")

  attrs := strings.Split(line, " ")

  return Node{Name: attrs[0], Left: attrs[1], Right: attrs[2]}
}

func MakeNodes(lines []string) []Node {
  nodes := []Node{}

  for _, line := range(lines) {
    nodes = append(nodes, MakeNode(line))
  }

  return nodes
}

func MakeInstructions(s string) []string {
  return strings.Split(s, "")
}

func ParseInput(filename string) ([]Node, []string) {
  content, err := os.ReadFile(filename)

  if err != nil {
    panic("problem parsing input file")
  }

  lines := strings.Split(string(content), "\n")

  instructions := MakeInstructions(lines[0])
  nodes := MakeNodes(lines[2:len(lines) - 1])

  return nodes, instructions
}

func ReadInstruction(s string, node Node) string {
  switch s {
  case "R":
    return node.Right
  case "L":
    return node.Left
  default:
    panic("invalid directions")
  }
}

func part1() {
  startTime := time.Now()
  var numSteps float64 = 0
  var exit bool = false

  nodes, instructions := ParseInput("input.txt")

  var nextNodeName = "AAA"
  var instructionIndex int

  lengthInstructions := float64(len(instructions))
  instructionIndex = int(math.Mod(numSteps, lengthInstructions))

  for {
    if exit {
      break
    }


    for _, node := range(nodes) {
      instructionIndex = int(math.Mod(numSteps, lengthInstructions))

      if nextNodeName == "ZZZ" {
        exit = true
        break
      } else if nextNodeName == node.Name {
        nextNodeName = ReadInstruction(instructions[instructionIndex], node)
        numSteps += 1
      }
    }
  }

  fmt.Println("Part 1:", numSteps)
  fmt.Println("Took:", time.Since(startTime))
}

func GetStartingNodes(nodes []Node) []Node {
  result := []Node{}

  for _, node := range(nodes) {
    if node.Name[2] == 'A' {
      result = append(result, node)
    }
  }

  return result
}

func NextNode(node Node, mapNodes []Node, instruction string) Node {
  var nextNodeName string

  switch instruction {
  case "R":
    nextNodeName = node.Right
  case "L":
    nextNodeName = node.Left
  default:
    panic("invalid instruction")
  }

  for {
    for _, n := range(mapNodes) {
      if nextNodeName == n.Name {
        return n
      }
    }
  }
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
      for b != 0 {
              t := b
              b = a % b
              a = t
      }
      return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
      result := a * b / GCD(a, b)

      for i := 0; i < len(integers); i++ {
              result = LCM(result, integers[i])
      }

      return result
}

func part2() {
  startTime := time.Now()
  nodes, instructions := ParseInput("input.txt")

  startingNodes := GetStartingNodes(nodes)
  var cycleCount int
  var cycleCounts = []int{}
  var instructionsIndex int

  for _, node := range(startingNodes) {
    nextNode := node
    cycleCount = 0

    for {
      if nextNode.Name[2] == 'Z' {
        cycleCounts = append(cycleCounts, cycleCount)
        break
      } else {
        instructionsIndex = int(math.Mod(float64(cycleCount),float64(len(instructions))))

        nextNode = NextNode(nextNode, nodes, instructions[instructionsIndex])
        cycleCount++
      }
    }
  }

  lcm := 1

  for _, count := range(cycleCounts) {
    lcm = LCM(lcm, count)
  }

  fmt.Println("Part 2:", lcm)
  fmt.Println("Took:", time.Since(startTime))
}

func main() {
  part1()
  part2()
}
