package main

import "fmt"

func twoSum(nums []int, target int) []int {

	m := make(map[int]int)

	for i, num := range nums {
		if val, found := m[target-num]; found {
			return []int{i, val}
		}
		m[num] = i
	}
	return nil
}

func main() {
	nums := []int{2, 3, 4}
	fmt.Println(twoSum(nums, 7))

	mapIterate()

}

func mapIterate() {
	sample := map[string]string{
		"ravi":  "shilpa",
		"kiran": "ammu",
	}

	for k, v := range sample {
		fmt.Printf("key :%s value: %s\n", k, v)
	}

	for _, v := range sample {
		fmt.Printf("value %s\n", v)
	}

	keys := getAllKeys(sample)
	fmt.Println(keys)

	LenofMap()
}

func getAllKeys(sample map[string]string) []string {
	var keys []string
	for k := range sample {
		keys = append(keys, k)
	}
	return keys
}

func LenofMap() {
	employeeSal := make(map[string]int)

	employeeSal["Raki"] = 2000
	employeeSal["shi"] = 1200

	LenofMap := len(employeeSal)
	fmt.Printf("Lenght of map %d", LenofMap)
}
