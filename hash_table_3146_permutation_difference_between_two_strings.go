func findPermutationDifference(s string, t string) int {
  frequency := make(map[rune]int)
  result := 0
  var res float64
  for index, val := range s {
    frequency[val] = index
  }
  for index, val := range t {
    value, ok := frequency[val]
    if ok {
      var indexS float64
      var indexT float64
      indexS = float64(index)
      indexT = float64(value)
      res += math.Abs(indexS - indexT)
    }
  }
  result = int(res)
}
