package hello

import (
	context "context"
	"fmt"
	"log"
)

// HelloService gRPC 服務接口，需要實作對應的 func
type HelloService struct{}

// SayHello 接收 gRPC 內容，並回傳對應 response
func (h *HelloService) SayHello(ctx context.Context, request *SayHelloRequest) (*SayHelloResponse, error) {
	log.Printf("透過 Request 收到訊息: Name:%s", request.Name)

	return &SayHelloResponse{
		Message: fmt.Sprintf("哈囉，%s。 很高興認識你！", request.Name),
	}, nil
}

func (h *HelloService) mustEmbedUnimplementedHelloServiceServer() {
	panic("implement me")
}
