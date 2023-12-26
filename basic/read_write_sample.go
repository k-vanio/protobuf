package basic

import (
	"fmt"
	"os"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"hands.on/protogen/users"
)

func WriteProtoToFile(message proto.Message, filename string, json bool) error {
	b, err := proto.Marshal(message)
	if err != nil {
		return err
	}

	if json {
		b, err = protojson.Marshal(message)
	} else {
		b, err = proto.Marshal(message)
	}

	if err := os.WriteFile(filename, b, 0644); err != nil {
		return err
	}

	return nil
}

func ReadProtoFromFile(name string, dest proto.Message, json bool) error {
	b, err := os.ReadFile(name)
	if err != nil {
		return err
	}

	if json {
		err = protojson.Unmarshal(b, dest)
	} else {
		err = proto.Unmarshal(b, dest)
	}

	if err != nil {
		return err
	}

	return nil
}

func WriterSample() {
	u := &users.User{
		Id:       1,
		Name:     "John Doe",
		IsActive: true,
		Address:  &users.Address{City: "New York", Country: "USA"},
	}

	if err := WriteProtoToFile(u, "user.json", true); err != nil {
		panic(err)
	}
}

func ReadSample() {
	u := &users.User{}
	if err := ReadProtoFromFile("user.json", u, true); err != nil {
		panic(err)
	}

	fmt.Println(u)
}
