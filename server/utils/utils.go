package utils

func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func Concat(slice1 []string, slice2 []string) (result []string){
	result = append(result, slice1...)
	result = append(result, slice2...)
	return
}