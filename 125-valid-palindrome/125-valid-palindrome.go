func validByte(char byte) byte {
    if (char >= 97 && char <= 122) || (char >= '0' && char <= '9'){
        return char
    } 
    if char >= 65 && char <= 90 {
        return char + 32
    }
    return 0
}

func isPalindrome(s string) bool {
    low, high := 0, len(s)-1
    for low < high {
        left, right := validByte(s[low]), validByte(s[high])
        if left == 0 {
            low++
        }
        if right == 0 {
            high--
        }
        if left == 0 || right == 0 {
            continue
        }
        if left != right {
            return false
        }
        low++
        high--
    }
    return true
}