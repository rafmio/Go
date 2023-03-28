package main

import (
	"log"
	"os"
	"io"
)

func main() {
	file, err := os.OpenFile("sloth.txt", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalf("error opening sloth.txt: %v", err)
	}
	defer file.Close()

	bytesRead := make([]byte, 33)
	n, err := file.Read(bytesRead)
	if err != nil {
		log.Fatalf("error reading from sloth.txt: %v", err)
	}

	log.Printf("We read \"%s\" into bytesRead (%d bytes)", string(bytesRead), n)

// ---------------------------------------------------------

	n, err = file.Read(bytesRead)
    if err != nil {
        log.Fatalf("error reading from sloth.txt: %v", err)
    }

    log.Printf("We read \"%s\" into bytesRead (%d bytes)",
        string(bytesRead), n)

 // ---------------------------------------------------------

 	n, err = file.Read(bytesRead)
 	if err == io.EOF {
 		log.Printf("End of file reached; we got %d bytes ot the last read", n)
 	} else {
 		log.Fatalf("unexpected error reading from the file: %v\n", err)
 	}
}


// I'm a sloth and hibiscus flowers are tasty!
