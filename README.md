# Integer Overflow in Go Average Calculation

This Go program demonstrates a potential integer overflow when calculating the average of a slice of integers.  The `processData` function calculates the sum of integers and divides it by the number of elements to find the average. However, if the sum of the integers exceeds the maximum value of an `int`, an integer overflow occurs, resulting in an incorrect average.  The solution provided addresses this by using `int64` to handle larger sums, preventing integer overflow.

## How to Run

1. Save the code in `bug.go` and `bugSolution.go`.
2. Run the program using `go run bug.go` and `go run bugSolution.go`.
3. Observe the output to see the difference in behavior between the original and corrected code.