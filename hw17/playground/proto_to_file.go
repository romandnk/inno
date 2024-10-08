package playground

import (
	"fmt"
	"os"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf(
			"cannot marshal proto message to binary: %w",
			err,
		)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("cannot write binary data to file: %w", err)
	}

	return nil
}

func ReadProtobufFromBinaryFile(filename string, message proto.Message) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read binary data from file: %w", err)
	}

	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf(
			"canot unmarshal binary to proto message: %w",
			err,
		)
	}

	return nil
}

func ProtobufToJSON(message proto.Message) ([]byte, error) {
	marshaler := protojson.MarshalOptions{
		UseEnumNumbers:    false,
		EmitDefaultValues: true,
		Indent:            " ",
		UseProtoNames:     true,
	}

	return marshaler.Marshal(message)
}

func WriteProtobufToJSONFile(message proto.Message, filename string) error {
	data, err := ProtobufToJSON(message)
	if err != nil {
		return fmt.Errorf(
			"cannot marshal proto message to JSON: %w",
			err,
		)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("cannot write JSON data to file: %w", err)
	}

	return nil
}
