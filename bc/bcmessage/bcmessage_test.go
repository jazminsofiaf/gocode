package bcmessage

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"testing"
)

const LENGTH = 105
const VERSION_END = 4
const SERVICE_END  = 12


const NODE_NETWORK = 1
const NODE_BLOOM = 4
const NODE_WITNESS = 8
const NODE_NETWORK_LIMITED = 1024




func TestPayloadWithCurrentDate(t *testing.T) {
	var payload = getPayloadWithCurrentDate()

	byteArray := []byte{0, 1}
	if( reflect.TypeOf(payload) != reflect.TypeOf(byteArray)){
		t.Errorf("wrong type payload. expected '%T', but was: '%T'", payload, byteArray)
	}

	if( len(payload) != LENGTH){
		t.Errorf("wrong type payload. expected '%v', but was: '%v'",  len(payload), LENGTH)
	}

	binaryNum := []byte{0x7F, 0x11, 0x01, 0x00}
	num := binary.LittleEndian.Uint32(binaryNum)
	fmt.Printf("decimal num %v = %v  in binary little indian\n", num, binaryNum)

	expectedBitcoinVersion := 70015
	binaryBitcoinVersion := payload[:VERSION_END]
	bitcoinVersion := int(binary.LittleEndian.Uint32(binaryBitcoinVersion))
	if(expectedBitcoinVersion != bitcoinVersion ){
		t.Errorf("wrong version. expected '%v', but was: '%v (binary little indian = %v)'",
			expectedBitcoinVersion, bitcoinVersion, binaryBitcoinVersion)
	}

	fmt.Println(int64(binary.LittleEndian.Uint64([]byte{13,4,0,0,0,0,0,0})))


	expectedServices := int64(NODE_NETWORK | NODE_BLOOM | NODE_WITNESS | NODE_NETWORK_LIMITED)
	binaryService := payload[VERSION_END:SERVICE_END]
	service := int64(binary.LittleEndian.Uint64(binaryService))
	if(expectedServices != service ){
		t.Errorf("wrong version. expected '%v', but was: '%v (binary little indian = %v)'",
			expectedServices, service, binaryService)
	}

}