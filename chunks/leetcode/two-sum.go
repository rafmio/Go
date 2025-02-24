func twoSum(nums []int, target int) []int {
    result := make([]int, 0)

    for idx, value := range nums {
        desired := target - value
        for iidx, summand := range nums {
            if summand == desired {
                if idx == iidx {
                    continue
                } else {
                    result = append(result, idx)
                    result = append(result, iidx)
                    break
                }
                
            }
        }
        
        if len(result) == 2 {
            break
        }
    }

    return result

    // for idx, value := range nums {
    //     if value < target && target != 0 {
    //         desired := target - value
    //         for iidx, summand := range nums {
    //             if summand == desired {
    //                 if idx == iidx {
    //                     continue
    //                 } else {
    //                     result = append(result, idx)
    //                     result = append(result, iidx)
    //                     break
    //                 }
                    
    //             }
    //         }
    //     } else {

    //     }
        
    //     if len(result) == 2 {
    //         break
    //     }
    // }

    // return result
}
