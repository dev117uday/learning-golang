import "math"

func getDecimalValue(head *ListNode) int {
	var list []int
	var sum float64 = 0
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	length := len(list) - 1
	for i := 0; i < len(list); i++ {
		if list[i] == 1 {
			sum += math.Pow(float64(2), float64(length-i))
		}
	}
	return int(sum)
}