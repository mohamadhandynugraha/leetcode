func numJewelsInStones(jewels string, stones string) int {
  frequency := make(map[rune]int)
  result := 0
  for _, jewel := range jewels {
    frequency[jewel]++
  }

  for _, stone := range stones {
    result += frequency[stone]
  }

  return result
}
