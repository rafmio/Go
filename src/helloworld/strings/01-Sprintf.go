package main

import "fmt"

func main() {
	witSs := `
		SELECT AVG(daily_count) AS average_records_per_day
		FROM (
			SELECT COUNT(*) AS daily_count
			FROM %s
			GROUP BY DATE(%s)

			`
	lg_tab := "lg_tab"
	tmstmp := "tmstmp"

	res := fmt.Sprintf(witSs, lg_tab, tmstmp)

	fmt.Println(res)

	str1 := `hello\nmello`
	// str2 := 'hello\nmello'
	str3 := "hello\nmello"

	fmt.Println(str1)
	// fmt.Println(str2)
	fmt.Println(str3)

	fmt.Println("--------------------")

	fmt.Printf("%s", str1)
	fmt.Println()
	// fmt.Printf("%s", str2)
	fmt.Printf("%s", str3)

	fmt.Println()
	fmt.Println("Done")
}
