// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: noble/authority/v1/tx.proto

package authorityv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Msg_Execute_FullMethodName           = "/noble.authority.v1.Msg/Execute"
	Msg_TransferAuthority_FullMethodName = "/noble.authority.v1.Msg/TransferAuthority"
	Msg_AcceptAuthority_FullMethodName   = "/noble.authority.v1.Msg/AcceptAuthority"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	Execute(ctx context.Context, in *MsgExecute, opts ...grpc.CallOption) (*MsgExecuteResponse, error)
	TransferAuthority(ctx context.Context, in *MsgTransferAuthority, opts ...grpc.CallOption) (*MsgTransferAuthorityResponse, error)
	AcceptAuthority(ctx context.Context, in *MsgAcceptAuthority, opts ...grpc.CallOption) (*MsgAcceptAuthorityResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Execute(ctx context.Context, in *MsgExecute, opts ...grpc.CallOption) (*MsgExecuteResponse, error) {
	out := new(MsgExecuteResponse)
	err := c.cc.Invoke(ctx, Msg_Execute_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) TransferAuthority(ctx context.Context, in *MsgTransferAuthority, opts ...grpc.CallOption) (*MsgTransferAuthorityResponse, error) {
	out := new(MsgTransferAuthorityResponse)
	err := c.cc.Invoke(ctx, Msg_TransferAuthority_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AcceptAuthority(ctx context.Context, in *MsgAcceptAuthority, opts ...grpc.CallOption) (*MsgAcceptAuthorityResponse, error) {
	out := new(MsgAcceptAuthorityResponse)
	err := c.cc.Invoke(ctx, Msg_AcceptAuthority_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	Execute(context.Context, *MsgExecute) (*MsgExecuteResponse, error)
	TransferAuthority(context.Context, *MsgTransferAuthority) (*MsgTransferAuthorityResponse, error)
	AcceptAuthority(context.Context, *MsgAcceptAuthority) (*MsgAcceptAuthorityResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) Execute(context.Context, *MsgExecute) (*MsgExecuteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedMsgServer) TransferAuthority(context.Context, *MsgTransferAuthority) (*MsgTransferAuthorityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferAuthority not implemented")
}
func (UnimplementedMsgServer) AcceptAuthority(context.Context, *MsgAcceptAuthority) (*MsgAcceptAuthorityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptAuthority not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgExecute)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Execute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Execute(ctx, req.(*MsgExecute))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_TransferAuthority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTransferAuthority)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).TransferAuthority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_TransferAuthority_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).TransferAuthority(ctx, req.(*MsgTransferAuthority))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AcceptAuthority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAcceptAuthority)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AcceptAuthority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_AcceptAuthority_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AcceptAuthority(ctx, req.(*MsgAcceptAuthority))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "noble.authority.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _Msg_Execute_Handler,
		},
		{
			MethodName: "TransferAuthority",
			Handler:    _Msg_TransferAuthority_Handler,
		},
		{
			MethodName: "AcceptAuthority",
			Handler:    _Msg_AcceptAuthority_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "noble/authority/v1/tx.proto",
}
