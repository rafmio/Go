func getCommon(nums1 []int, nums2 []int) int {
    commonNums := make([]int, 0)
    var min int
    
    if len(nums1) == 0 || len(nums2) == 0 {
        return -1
    } else {
        for _, val1 := range nums1 {
            for _, val2 := range nums2 {
                if val1 == val2 {
                    commonNums = append(commonNums, val2)
                    break
                }
            }
        }
    }
    
    if len(commonNums) == 0 {
        return -1
    } else {
        min = commonNums[0]
        for _, val := range commonNums {
            if val < min {
                min = val
            }
        }
        return min
    }
}
