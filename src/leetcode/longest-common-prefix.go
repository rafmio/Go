func longestCommonPrefix(strs []string) string {
    var minPrefix string

    if len(strs) == 1 {
        return strs[0]
    }
       
    var minWord string
    for i, value := range strs {
        if i == 0 {
            minWord = value
        } else {
            if len(value) <= len(minWord) {
                minWord = value
            } else {
                continue
            }
        }
    }

    if len(minWord) == 1 {
        var commonChar string
        for _, value := range strs {
            if string(value[0]) == string(minWord[0]) {
                commonChar = string(value[0])
            } else {
                commonChar = ""
            }
        }
        if commonChar != "" {
            minPrefix += commonChar
        }
    } else {
        for i := 0; i < len(minWord); i++ {
            var commonChar string
            for _, value := range strs {
                if string(value[i]) == string(minWord[i]) {
                    commonChar = string(value[i])
                } else if string(value[i]) != string(minWord[i]) {
                    commonChar = ""
                    break
                }
            }

            if commonChar != "" {
                minPrefix += commonChar
            } else {
                break
            }
        }
    }


    return minPrefix
}
