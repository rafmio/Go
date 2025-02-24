package randperm

import "math/rand"

func RandPerm(list []string) []string {
	for i := 0; i < len(list); i++ {
		j := i + rand.Intn(len(list)-i)
		list[i], list[j] = list[j], list[i]
	}

	return list
}
