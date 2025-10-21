package msgx

import (
	"sdxx/server/internal/code"
	corev1 "sdxx/server/protobuf/gen/core/v1"

	"github.com/dobyte/due/v2/cluster"
	"github.com/dobyte/due/v2/log"
	proto "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func Cluster(route int32, message proto.Message) *cluster.Message {
	data, err := anypb.New(message)
	if err != nil {
		log.Errorf("anypb.New failed: %v", err)
		return nil
	}
	return &cluster.Message{
		Route: route,
		Data: &corev1.Response{
			Code: int32(code.OK.Code()),
			Msg:  code.OK.Message(),
			Data: data,
		},
	}
}
