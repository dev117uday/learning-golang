func numJewelsInStones(J string, S string) int {
	result := 0

	for i:=0; i<len(J); i++ {
		for j:=0; j<len(S); j++ {
			if S[j] == J[i] {
				result++
			}
		}
	}
	return result
}