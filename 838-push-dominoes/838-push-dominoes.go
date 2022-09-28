func pushDominoes(dominoes string) string {
    pivot := ' '
    dot := 0
    var res []rune
    for _, val := range dominoes {
        if val == '.' {
            dot++
        } else if val == 'L' {
            if pivot == ' ' || pivot == 'L' {
                for i := 0; i <= dot; i++ {
                    res = append(res, val)
                }
            } else {
                mid := dot / 2
                for i := 0; i <= dot; i++ {
                    if mid == i {
                        pivot = val
                        if dot % 2 == 1 {
                            res = append(res, '.')
                            continue
                        }
                    }
                    res = append(res, pivot)                    
                }
            }
            pivot = val
            dot = 0
        } else {
            if pivot == ' ' {
                for i := 0; i < dot; i++ {
                    res = append(res, '.')
                }
                res = append(res, val)
            } else if pivot == 'R' {
                for i := 0; i <= dot; i++ {
                    res = append(res, val)
                }
            } else {
                for i := 0; i < dot; i++ {
                    res = append(res, '.')                    
                }
                res = append(res, val)
            }
            pivot = val
            dot = 0
        }
    }
    if dot != 0 {
        if pivot == ' ' || pivot == 'L' {
            pivot = '.'
        } else {
            pivot = 'R'
        }
        for i := 0; i < dot; i++ {
            res = append(res, pivot)                    
        }
    }
    return string(res)
}