package main

func groupAnagramsEff(str []string) [][]string {
	res := [][]string{}
	tmp := map[[26]int][]string{}

	for _, s := range str {
		chars := [26]int{}

		for _, c := range s {
			chars[c-'a']++
		}
		tmp[chars] = append(tmp[chars], s)
	}

	for _, v := range tmp {
		res = append(res, v)
	}

	return res
}
