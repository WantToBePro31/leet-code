func isValid(s string) bool {
	var data []byte
	for i := 0; i < len(s); i++ {
		if len(data) == 0 {
			data = append(data, s[i])
		} else {
			if s[i] - data[len(data)-1] == 1 || s[i] - data[len(data)-1] == 2 {
				data = data[:len(data)-1]
			} else {
				data = append(data, s[i]) 
			}
		}
	}
	if len(data) == 0 {
		return true
	}
	return false
}