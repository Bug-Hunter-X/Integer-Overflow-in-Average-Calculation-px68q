func processData(data []int) (int, error) {
  if len(data) == 0 {
    return 0, errors.New("data slice is empty")
  }

  sum := 0
  for _, v := range data {
    sum += v
  }

  avg := float64(sum) / float64(len(data))

  // potential bug: integer overflow if sum is large
  return int(avg), nil
}