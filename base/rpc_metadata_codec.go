package base

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/haupc/cartransplant/grpcproto"
	"google.golang.org/grpc/metadata"
)

var (
	headerRPCMetadata = "rpc_metadata"
)

// RPCMetadataFromMD func
func RPCMetadataFromMD(md metadata.MD) (*grpcproto.UserProfile, error) {
	val := metautils.NiceMD(md).Get(headerRPCMetadata)
	if val == "" {
		return nil, nil
	}

	// proto.Marshal()
	buf, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		return nil, fmt.Errorf("base64 decode error, rpc_metadata: %s, error: %v", val, err)
		//panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_OTHER2),
		//	fmt.Sprintf("Base64 decode error, rpc_metadata: %s", val)))
	}

	rpcMetadata := &grpcproto.UserProfile{}
	err = proto.Unmarshal(buf, rpcMetadata)
	if err != nil {
		return nil, fmt.Errorf("RpcMetadata unmarshal error, rpc_metadata: %s, error: %v", val, err)
		//panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_OTHER2),
		//	fmt.Sprintf("Unmarshal error, rpc_metadata: %s", val)))
	}

	return rpcMetadata, nil
}

// RPCMetadataFromIncoming func
func RPCMetadataFromIncoming(ctx context.Context) *grpcproto.UserProfile {
	md2, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil
	}

	md, _ := RPCMetadataFromMD(md2)
	return md
}

// RPCMetadataToOutgoing func
func RPCMetadataToOutgoing(ctx context.Context, md *grpcproto.UserProfile) (context.Context, error) {
	buf, err := proto.Marshal(md)
	if err != nil {
		return nil, err
	}

	return metadata.NewOutgoingContext(ctx, metadata.Pairs(headerRPCMetadata, base64.StdEncoding.EncodeToString(buf))), nil
}

// RPCMetadataToOutgoingForInternal func
// For send internal server
func RPCMetadataToOutgoingForInternal(ctx context.Context, md *grpcproto.UserProfile) (context.Context, error) {
	buf, err := proto.Marshal(md)
	if err != nil {
		return nil, err
	}

	return metadata.NewIncomingContext(ctx, metadata.Pairs(headerRPCMetadata, base64.StdEncoding.EncodeToString(buf))), nil
}
