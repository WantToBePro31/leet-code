func validUtf8(data []int) bool {
    var strBinary []string
    for _, val := range data {
        binary := fmt.Sprintf("%08s", strconv.FormatInt(int64(val), 2))
        strBinary = append(strBinary, binary)
    }
    i := 0
    for i < len(strBinary) {
        if strBinary[i][0] == '0' {
            i++
        } else {
            if strBinary[i][1] == '0' {
                return false
            } else {
                if strBinary[i][2] == '0' {
                    if i + 1 >= len(strBinary) {
                        return false
                    } else {
                        if strBinary[i+1][0] == '1' && strBinary[i+1][1] == '0' {
                            i += 2
                        } else {
                            return false
                        }
                    }
                } else {
                    if strBinary[i][3] == '0' {
                        if i + 2 >= len(strBinary) {
                            return false
                        } else {
                            if (strBinary[i+1][0] == '1' && strBinary[i+1][1] == '0') && (strBinary[i+2][0] == '1' && strBinary[i+2][1] == '0') {
                                i += 3
                            } else {
                                return false
                            }
                        }
                    } else {
                        if strBinary[i][4] == '0' {
                            if i + 3 >= len(strBinary) {
                                return false
                            } else {
                                if (strBinary[i+1][0] == '1' && strBinary[i+1][1] == '0') && (strBinary[i+2][0] == '1' && strBinary[i+2][1] == '0') && (strBinary[i+3][0] == '1' && strBinary[i+3][1] == '0') {
                                i += 4
                                } else {
                                    return false
								}
							}
                        } else {
                            return false
                        }
                    }
                }
            }
        }
    }
    return true
}