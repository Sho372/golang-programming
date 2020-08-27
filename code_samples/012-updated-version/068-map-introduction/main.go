package main

import "fmt"

func main() {

	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}

	fmt.Println(m)

	fmt.Println(m["James"])

	fmt.Println(m["Miss Moneypenny"])

	v, ok := m["Barnabas"]
	fmt.Println(v)  // If the key doesn't exist, the value is zero value
	fmt.Println(ok) // Check if the key exist -> false

	v2, ok2 := m["James"]
	fmt.Println(v2)
	fmt.Println(ok2)

	// comma ok idiom
	// ifの条件文の中の変数定義はスコープが違う?ここでのvは上とは違う、たぶん
	if v, ok := m["James"]; ok {
		fmt.Println(v)
	}
}
