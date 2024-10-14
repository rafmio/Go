package main

type collection struct {
	name   string
	strSls []string
	intSls []string
	mapMap map[string]int
}

func main() {
	inst1 := &collection{
		name:   "inst1",
		strSls: []string{"one", "two", "three", "four"},
		intSls: []string{"1", "2", "3", "4"},
	}

}
