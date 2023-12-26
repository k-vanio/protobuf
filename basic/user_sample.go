package basic

import (
	"log"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"hands.on/protogen/users"
)

func BasicUser() {
	var m anypb.Any
	anypb.MarshalFrom(&m, &users.InstantMessage{Message: "Hi message"}, proto.MarshalOptions{})

	metadata := make(map[string]*users.Address)
	metadata["home"] = &users.Address{City: "Anytown"}
	metadata["work"] = &users.Address{City: "Othertown"}

	u := &users.User{
		Id:                   1,
		Name:                 "John Doe",
		IsActive:             true,
		Password:             []byte("secret"),
		Roles:                []string{"admin", "user"},
		Address:              &users.Address{},
		CommunicationChannel: &m,
		Metadata:             metadata,
	}

	log.Println(u)
}

func ToJsonUser() {
	var m anypb.Any
	anypb.MarshalFrom(&m, &users.InstantMessage{Message: "Hi message"}, proto.MarshalOptions{})

	metadata := make(map[string]*users.Address)
	metadata["home"] = &users.Address{City: "Anytown"}
	metadata["work"] = &users.Address{City: "Othertown"}

	u := &users.User{
		Id:       1,
		Name:     "John Doe",
		IsActive: true,
		Password: []byte("secret"),
		Roles:    []string{"admin", "user"},
		Address: &users.Address{
			Street:  "123 Main St",
			City:    "Anytown",
			Country: "USA",
			Coordinate: &users.Address_Coordinate{
				Latitude:  1.234,
				Longitude: 5.678,
			},
		},
		CommunicationChannel: &m,
		Metadata:             metadata,
	}

	jsonBytes, _ := protojson.Marshal(u)
	log.Println(string(jsonBytes))
}
