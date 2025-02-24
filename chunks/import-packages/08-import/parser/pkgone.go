package parser

import "fmt"

const (
	portTest                 string = "8081"
	portBlackOxygenium       string = "5432"
	portCuteGanymede         string = "5432"
	serverNameBlackOxygenium string = "BlackOxygenium"
	serverNameCuteGanymede   string = "CuteGanymede"
	ipBlackOxygenium         string = "194.58.102.129"
	ipCuteGanymede           string = "147.45.226.19"
	fileDBConfigName         string = "../config/databaseConfig.json"
	fileConfigName           string = "../config/fileConfig.json"
)

func PrintData() {
	fmt.Println("portTest:", portTest)
	fmt.Println("portBlackOxygenium:", portBlackOxygenium)
	fmt.Println("portCuteGanymede:", portCuteGanymede)
	fmt.Println("serverNameBlackOxygenium:", serverNameBlackOxygenium)
	fmt.Println("serverNameCuteGanymede:", serverNameCuteGanymede)
	fmt.Println("ipBlackOxygenium:", ipBlackOxygenium)
	fmt.Println("ipCuteGanymede:", ipCuteGanymede)
	fmt.Println("fileDBConfigName:", fileDBConfigName)
	fmt.Println("fileConfigName:", fileConfigName)
	fmt.Println()

}
