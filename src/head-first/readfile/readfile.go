package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {
    file, err := os.Open("data.txt")
    if err != nil {
        log.Fatal(err)
    }
    
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
    
    err = file.Close()
    if err != nil {
        log.Fatal(err)
    }
    
    if scanner.Err() != nil {
        log.Fatal(scanner.Err())
    }
}

/*
Метод Scan значения bufio.Scanner создан для использования в циклах for. Он читает одну строку текста из файла и возвращает true, если данные прочитаны успешно, или false в случае ошибки. Если Scan используется как условие цикла for, цикл продолжит выполняться, пока остаются данные для чтения. При достижении конца файла (или возникновении ошибки) Scan вернет false и цикл завершится.
*/

// метод Text возвращает строку с прочитанными данными
