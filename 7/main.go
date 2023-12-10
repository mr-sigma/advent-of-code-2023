package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "sort"
  "strconv"
)

func CardVals(char rune, wildcard bool)int {
  switch char {
  case 'A':
    return 14
  case 'K':
    return 13
  case 'Q':
    return 12
  case 'J':
    if wildcard {
      return 1
    } else {
      return  11
    }
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

func (h Hand) Gt(oh Hand, wildcard bool) bool {
  if h.Strength > oh.Strength {
    return true
  } else if h.Strength < oh.Strength {
    return false
  } else {
    for i, val := range(h.Cards) {
      card1 := CardVals(val, wildcard)
      card2 := CardVals([]rune(oh.Cards)[i], wildcard)
      if card1 > card2 {
        return true
      } else if card1 < card2 {
        return false
      }
    }
  }

  return false
}

func (h Hand) Lt(oh Hand, wildcard bool) bool {
  if h.Strength < oh.Strength {
    return true
  } else if h.Strength > oh.Strength {
    return false
  } else {
    for i, val := range(h.Cards) {
      card1 := CardVals(val, wildcard)
      card2 := CardVals([]rune(oh.Cards)[i], wildcard)
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

func (h *Hand) CalcStrengthWildcard() {
  h.CalcStrength()

  if h.Strength == 6 { // five of a kind can't get better
    return
  }

  jokers := strings.Count(h.Cards, "J")

  switch h.Strength {
  case 5:
    if jokers > 0 {
      h.Strength = h.Strength + 1
    }
  case 4:
    if jokers > 0 {
      h.Strength = 6
    }
  case 3:
    if jokers > 0 {
      if jokers == 3 {
        h.Strength = 5
      } else {
        h.Strength = h.Strength + jokers + 1
      }
    }
  case 2:
    if jokers == 2 {
      h.Strength = 5
    } else if jokers == 1 {
      h.Strength = 4
    }
  case 1:
    if jokers > 0 {
      h.Strength = 3
    }
  default:
    if jokers > 0 {
      h.Strength = 1
    }
  }
}

func MakeHand(s string, wildcard bool) Hand {
  attrs := strings.Split(s, " ")
  bid, err := strconv.Atoi(attrs[1])

  if err != nil {
    panic(err)
  }

  hand := Hand{Cards: attrs[0], Bid: bid}

  if wildcard {
    hand.CalcStrengthWildcard()
  } else {
    hand.CalcStrength()
  }

  return hand
}

func MakeHands(file *os.File, wildcard bool) []Hand {
  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)

  var hands []Hand

  for scanner.Scan() {
    h := MakeHand(scanner.Text(), wildcard)
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

func mergeSort(items []Hand, wildcard bool) []Hand {
    if len(items) < 2 {
        return items
    }
    first := mergeSort(items[:len(items)/2], wildcard)
    second := mergeSort(items[len(items)/2:], wildcard)
    return merge(first, second, wildcard)
}

func merge(a []Hand, b []Hand, wildcard bool) []Hand {
    final := []Hand{}
    i := 0
    j := 0
    for i < len(a) && j < len(b) {
        if a[i].Lt(b[j], wildcard) {
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



func SortHandsByStrength(hands []Hand, wildcard bool) []Hand {
  return mergeSort(hands, wildcard)
}

func part1() {
  file := OpenFile("input.txt")
  hands := MakeHands(file, false)
  hands = SortHandsByStrength(hands, false)

  sum := 0

  for i, hand := range(hands) {
    sum += hand.Bid * (i + 1)
  }

  fmt.Println("Part 1:", sum)
}

func part2() {
  file := OpenFile("input.txt")
  hands := MakeHands(file, true)
  hands = SortHandsByStrength(hands, true)

  sum := 0

  for i, hand := range(hands) {
    sum += hand.Bid * (i + 1)
  }

  fmt.Println("Part 1:", sum)

}

func main() {
  part1()
  part2()
}
