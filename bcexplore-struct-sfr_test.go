package main

import (
	"os"
	"testing"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestStoreEvent(t *testing.T) {
	setGlobalVarForTest(false, "../src/testOutput")
	var message = "this is a message to write in file output."
	storeEvent(message)

	f, err := os.Open("../src/testOutput")
	check(err)
	binaryWrittenMessage := make([]byte, len(message))
	nBits, err := f.Read(binaryWrittenMessage)
	check(err)
	writtenMessage :=string(binaryWrittenMessage[:nBits])

	if(writtenMessage != message){
		t.Errorf("'%v' != '%v'", message, writtenMessage )
	}
}
