// Loops can be used in the different forms:

// 1.As simple for loop - it similar to C, C++, Java, C#, etc
for i := 0; i < 4; i++ {
        fmt.Printf("Print the string 4 times")
    }

    // 2.For loop as infinite loop 
for {
    fmt.Printf("Types infinite more times")
}

// 3.for loop as while loop:
i := 0
for i < 3 {
    i += 2
}

// 4.Simple range in for loop:
rvariable := []string{"GFG", "Geeks", "Programmers"}
for i, j := range rvariable {
    fmt.Println(i, j)
}

// 5.Using for loop for strings
for i, j := range "Combination" {
    fmr.Printf("The index number of %U is %d\n", j, i)
}

// 6.for maps
mmap := map[int]string{
    22:"Geek",
    33:"GFG",
    44:"Developer",
}

for key, value := range mmap {
    fmt.Println(key, value)
}

// 7.for channel:
chnl := make(chan int)
go func() {
    chnl <- 100
    chnl <- 1000
    chnl <- 10000
    chnl <- 100000
    close(chnl)
} ()

for i := range chnl {
    fmt.Println(i)
}
