package main

func ArrayContains(arr []string, str string) bool {
	for _, k := range arr {
		if k == str {
			return true
			break
		}
	}
	return false
}
