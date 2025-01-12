func processData(data []int) (int64, error) {
  if len(data) == 0 {
    return 0, errors.New("data slice is empty")
  }

  sum := int64(0)
  for _, v := range data {
    sum += int64(v)
  }

  avg := float64(sum) / float64(len(data))

  return int64(avg), nil
} 