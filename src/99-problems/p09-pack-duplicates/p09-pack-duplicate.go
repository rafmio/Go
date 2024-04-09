package packdup

// PackDupStr accepts a string slice ([]string) pack duplicates into it's own
// slice and do similar with unique elements - pack them into own slice with
// single one element
func PackDupStr(list []string) [][]string {
	packedDupListStr := make([][]string, 0)
	uniqueList := make([]string, 0)

	for i := 0; i < len(list)-1; i++ {
		if len(uniqueList) == 0 {
			uniqueList = append(uniqueList, list[i])
		} else if list[i] == uniqueList[0] {
			uniqueList = append(uniqueList, list[i])
		} else if list[i] != uniqueList[0] {
			packedDupListStr = append(packedDupListStr, uniqueList)
			uniqueList = uniqueList[:0]
		}
	}

	return packedDupListStr
}
