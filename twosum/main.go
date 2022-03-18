package main

func main() {
	nums := []int{11, 2, 15, 7}
	target := 9
	result := twoSum(nums, target)
	print(result)
}
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, v := range nums {

		if key, ok := m[target-v]; ok {
			return []int{index, key}
		}
		m[v] = index
	}
	return nil
}
