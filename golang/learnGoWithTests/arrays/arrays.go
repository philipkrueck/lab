package arrays

func sum(inp []int) (sum int) {
	for _, num := range inp {
		sum += num
	}
	return
}
