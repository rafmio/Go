package client

import (
	"fmt"
	"io"
	"log"
	"testing"
	"time"
)

func TestRunClient(t *testing.T) {
	currentDate := time.Now()
	want := currentDate.Format("01-02-2006")

	go runServer()

	got := runClient()
	formatGot, err := io.ReadAll(got)
	if err != nil {
		log.Fatalln(err)
	}

	if string(formatGot) != want {
		t.Errorf("incorrect return - format or value, wanted: %q, got: %q\n", want, formatGot)
	} else {
		fmt.Println("Everithing is OK")
	}
}
