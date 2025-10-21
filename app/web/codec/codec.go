package codec

import (
	"encoding/json"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var marshalOptions = protojson.MarshalOptions{
	UseProtoNames:   true,
	UseEnumNumbers:  true,
	EmitUnpopulated: true,
}

var unmarshalOptions = protojson.UnmarshalOptions{}

func JsonEncoder(v any) ([]byte, error) {
	if m, ok := v.(proto.Message); ok {
		return marshalOptions.Marshal(m)
	} else {
		return json.Marshal(v)
	}
}

func JsonDecoder(data []byte, v any) error {
	if m, ok := v.(proto.Message); ok {
		return unmarshalOptions.Unmarshal(data, m)
	} else {
		return json.Unmarshal(data, v)
	}
}
