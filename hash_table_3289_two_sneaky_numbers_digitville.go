func getSneakyNumbers(nums []int) []int {
  result := []int{}
  frequency := make(map[int]int)

  for _, num := range nums {
    frequency[num]++
  }

  for key, val := range frequency {
    if val >= 2 {
      result = append(result, key)
    }
  }
  return result
}
