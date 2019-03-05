package problem35

func sort(arr []string) []string {
	low, mid, high := 0, 0, len(arr)-1
	for mid <= high {
		if arr[mid] == "R" {
			arr[low], arr[mid] = arr[mid], arr[low]
			low++
			mid++
		} else if arr[mid] == "G" {
			mid++
		} else {
			arr[mid], arr[high] = arr[high], arr[mid]
			high--
		}
	}
	return arr
}
