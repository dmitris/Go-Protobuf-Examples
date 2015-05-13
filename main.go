package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
	protobuftest "github.com/minaandrawos/Go-Protobuf-Examples/ProtobufTest"
)

//go:generate protoc --go_out=./ProtobufTest ProtoTest.proto
func main() {
	test := protobuftest.TestMessage{
		ClientName:  proto.String("testname"),
		ClientId:    proto.Int32(1),
		Description: proto.String("test description"),
	}
	data, err := proto.Marshal(&test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	newTest := &protobuftest.TestMessage{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	// Now test and newTest contain the same data.
	if test.GetClientName() != newTest.GetClientName() {
		log.Fatalf("data mismatch %q != %q", test.GetClientName(), newTest.GetClientName())
	}
	fmt.Println("OK")
}
