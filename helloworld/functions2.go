package main
import (
    "fmt"
    "io/ioutil"
)

func getFiles(path string) ([]string, error) {
    var files []string
    dirFiles, err := ioutil.ReadDir(fmt.Sprintf("%s", path))
    if err != nil {
        return nil, err 
    }
    
    for _, i := range dirFiles {
        files = append(files, i.Name())
    }
    return files, nil
}

func test(path string) ([]string, error) {
    return getFiles(path)
}

func returnRog() (rog int) {
    rog = 123
    return
}

func main() {
    // The name of the function is may be omited during declaraion
    // The name of the function would be privet()
    privet := func(x, y int) int {
        i := (x + y) * 2
        return i
    }
    
    f, _ := getFiles(".")  // '_' is an error, here we don't use it
    fmt.Println(f)
    fmt.Println(returnRog())
    fmt.Println(privet(1, 5))
    fmt.Println(test("/"))
}
 
// Functions
// We can declare function but may not use it
