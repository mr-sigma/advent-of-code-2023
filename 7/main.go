package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "sort"
  "strconv"
)

func CardVals(char rune)int {
  switch char {
  case 'A':
    return 14
  case 'K':
    return 13
  case 'Q':
    return 12
  case 'J':
    return  11
  case 'T':
    return 10
  case '9':
    return 9
  case '8':
    return 8
  case '7':
    return 7
  case '6':
    return 6
  case '5':
    return 5
  case '4':
    return 4
  case '3':
    return 3
  case '2':
    return 2
  default:
    panic("passed unknown card")
  }

}

func Unique(s string) string {
  keys := make(map[rune]bool)
  uniq := []rune{}
  for _, entry := range(s) {
    if _, value := keys[entry]; !value {
      keys[entry] = true
      uniq = append(uniq, entry)
    }
  }

  return string(uniq)
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
    return len(s)
}

func SortString(s string) string {
    r := []rune(s)
    sort.Sort(sortRunes(r))
    return string(r)
}

type Hand struct {
  Cards string
  Strength int
  Bid int
}

func (h Hand) Gt(oh Hand) bool {
  if h.Strength > oh.Strength {
    return true
  } else if h.Strength < oh.Strength {
    return false
  } else {
    for i, val := range(h.Cards) {
      card1 := CardVals(val)
      card2 := CardVals([]rune(oh.Cards)[i]) 
      if card1 > card2 {
        return true
      } else if card1 < card2 {
        return false
      }
    }
  }

  return false
}

func (h Hand) Lt(oh Hand) bool {
  if h.Strength < oh.Strength {
    return true
  } else if h.Strength > oh.Strength {
    return false
  } else {
    for i, val := range(h.Cards) {
      card1 := CardVals(val)
      card2 := CardVals([]rune(oh.Cards)[i]) 
      if card1 < card2 {
        return true
      } else if card1 > card2 {
        return false
      }
    }
  }

  return false
}

func (h Hand) IsFiveOfAKind() bool {
  return len(Unique(h.Cards)) == 1
}

func (h Hand) IsFourOfAKind() bool {
  sorted := SortString(h.Cards)
  return (len(Unique(sorted[:4])) == 1 || len(Unique(sorted[1:])) == 1)
}

func (h Hand) IsFullHouse() bool {
  return len(Unique(h.Cards)) == 2
}

func (h Hand) IsThreeOfAKind() bool {
  counts := map[rune]int {}
  is := false

  for _, card := range(h.Cards) {
    counts[card]++

    if counts[card] == 3 {
      is = true
    } else if counts[card] > 3 {
      is = false
    }
  }

  return is
}

func (h Hand) IsTwoPair() bool {
  return len(Unique(h.Cards)) == 3
}

func (h Hand) IsOnePair() bool {
  return len(Unique(h.Cards)) == 4
}

func (h *Hand) CalcStrength() {
  var strength int

  switch {
  case h.IsFiveOfAKind():
    strength = 6
  case h.IsFourOfAKind():
    strength = 5
  case h.IsFullHouse():
    strength = 4
  case h.IsThreeOfAKind():
    strength = 3
  case h.IsTwoPair():
    strength = 2
  case h.IsOnePair():
    strength = 1
  default:
    strength = 0
  }

  h.Strength = strength
}

func MakeHand(s string) Hand {
  attrs := strings.Split(s, " ")
  bid, err := strconv.Atoi(attrs[1])

  if err != nil {
    panic(err)
  }

  hand := Hand{Cards: attrs[0], Bid: bid}

  hand.CalcStrength()

  return hand
}

func MakeHands(file *os.File) []Hand {
  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)

  var hands []Hand

  for scanner.Scan() {
    h := MakeHand(scanner.Text())
    hands = append(hands, h)
  }

  return hands
}

func OpenFile(filename string) *os.File {
  file, err := os.Open(filename)

  if err != nil {
    panic(err)
  }

  return file
}

func mergeSort(items []Hand) []Hand {
    if len(items) < 2 {
        return items
    }
    first := mergeSort(items[:len(items)/2])
    second := mergeSort(items[len(items)/2:])
    return merge(first, second)
}

func merge(a []Hand, b []Hand) []Hand {
    final := []Hand{}
    i := 0
    j := 0
    for i < len(a) && j < len(b) {
        if a[i].Lt(b[j]) {
            final = append(final, a[i])
            i++
        } else {
            final = append(final, b[j])
            j++
        }
    }

    for ; i < len(a); i++ {
        final = append(final, a[i])
    }

    for ; j < len(b); j++ {
        final = append(final, b[j])
    }

    return final
}



func SortHandsByStrength(hands []Hand) []Hand {
  return mergeSort(hands)
}

func part1() {
  file := OpenFile("input.txt")
  hands := MakeHands(file)
  hands = SortHandsByStrength(hands)

  sum := 0

  for i, hand := range(hands) {
    fmt.Println("cards:", hand.Cards)
    fmt.Println("strength:", hand.Strength)
    fmt.Println("bid:", hand.Bid)
    sum += hand.Bid * (i + 1)
  }

  fmt.Println("Part 1:", sum)
}

func main() {
  part1()
}
