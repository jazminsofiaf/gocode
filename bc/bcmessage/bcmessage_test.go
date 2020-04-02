package bcmessage

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"testing"
)


func TestSendRequest(t *testing.T) {
	var payload = getPayloadWithCurrentDate()

	byteArray := []byte{0, 1}
	if( reflect.TypeOf(payload) != reflect.TypeOf(byteArray)){
		t.Errorf("wrong type payload. expected '%T', but was: '%T'", payload, byteArray)
	}

	if( len(payload) != 105){
		t.Errorf("wrong type payload. expected '%v', but was: '%v'",  len(payload), 105)
	}

	binaryNum := []byte{0x7F, 0x11, 0x01, 0x00}
	num := binary.LittleEndian.Uint32(binaryNum)
	fmt.Printf("decimal num %v = %v  in binary little indian\n", num, binaryNum)

	expectedBitcoinVersion := 70015
	binaryBitcoinVersion := payload[:4]
	bitcoinVersion := int(binary.LittleEndian.Uint32(binaryBitcoinVersion))
	if(expectedBitcoinVersion != bitcoinVersion ){
		t.Errorf("wrong version. expected '%v', but was: '%v (binary little indian = %v)'",
			expectedBitcoinVersion, bitcoinVersion, binaryBitcoinVersion)
	}

}