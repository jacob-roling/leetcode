package leetcode

type Key [26]byte

func GroupAnagrams(strs []string) [][]string {
	groupMap := make(map[Key][]string)
	for _, str := range strs {
		key := strKey(str)
		groupMap[key] = append(groupMap[key], str)
	}
	groups := make([][]string, 0, len(groupMap))
	for _, group := range groupMap {
		groups = append(groups, group)
	}
	return groups
}

func strKey(str string) Key {
	var key Key

	for i := range str {
		key[str[i]-'a']++
	}

	return key
}

// var alphabet string = "abcdefghijklmnopqrstuvwxyz"
// var primes []float64 = []float64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}

// func hash(s string) int {
// 	var product float64 = 1
// 	var mod float64 = 100000007
// 	for _, r := range s {
// 		product *= math.Mod(primes[strings.IndexRune(alphabet, r)], mod)
// 	}
// 	return int(product)
// }
