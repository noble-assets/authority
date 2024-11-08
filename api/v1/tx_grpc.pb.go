// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
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
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Msg_Execute_FullMethodName           = "/noble.authority.v1.Msg/Execute"
	Msg_TransferOwnership_FullMethodName = "/noble.authority.v1.Msg/TransferOwnership"
	Msg_AcceptOwnership_FullMethodName   = "/noble.authority.v1.Msg/AcceptOwnership"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	Execute(ctx context.Context, in *MsgExecute, opts ...grpc.CallOption) (*MsgExecuteResponse, error)
	TransferOwnership(ctx context.Context, in *MsgTransferOwnership, opts ...grpc.CallOption) (*MsgTransferOwnershipResponse, error)
	AcceptOwnership(ctx context.Context, in *MsgAcceptOwnership, opts ...grpc.CallOption) (*MsgAcceptOwnershipResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Execute(ctx context.Context, in *MsgExecute, opts ...grpc.CallOption) (*MsgExecuteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgExecuteResponse)
	err := c.cc.Invoke(ctx, Msg_Execute_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) TransferOwnership(ctx context.Context, in *MsgTransferOwnership, opts ...grpc.CallOption) (*MsgTransferOwnershipResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgTransferOwnershipResponse)
	err := c.cc.Invoke(ctx, Msg_TransferOwnership_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AcceptOwnership(ctx context.Context, in *MsgAcceptOwnership, opts ...grpc.CallOption) (*MsgAcceptOwnershipResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MsgAcceptOwnershipResponse)
	err := c.cc.Invoke(ctx, Msg_AcceptOwnership_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility.
type MsgServer interface {
	Execute(context.Context, *MsgExecute) (*MsgExecuteResponse, error)
	TransferOwnership(context.Context, *MsgTransferOwnership) (*MsgTransferOwnershipResponse, error)
	AcceptOwnership(context.Context, *MsgAcceptOwnership) (*MsgAcceptOwnershipResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMsgServer struct{}

func (UnimplementedMsgServer) Execute(context.Context, *MsgExecute) (*MsgExecuteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedMsgServer) TransferOwnership(context.Context, *MsgTransferOwnership) (*MsgTransferOwnershipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferOwnership not implemented")
}
func (UnimplementedMsgServer) AcceptOwnership(context.Context, *MsgAcceptOwnership) (*MsgAcceptOwnershipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptOwnership not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}
func (UnimplementedMsgServer) testEmbeddedByValue()             {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	// If the following call pancis, it indicates UnimplementedMsgServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
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

func _Msg_TransferOwnership_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTransferOwnership)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).TransferOwnership(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_TransferOwnership_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).TransferOwnership(ctx, req.(*MsgTransferOwnership))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AcceptOwnership_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAcceptOwnership)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AcceptOwnership(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_AcceptOwnership_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AcceptOwnership(ctx, req.(*MsgAcceptOwnership))
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
			MethodName: "TransferOwnership",
			Handler:    _Msg_TransferOwnership_Handler,
		},
		{
			MethodName: "AcceptOwnership",
			Handler:    _Msg_AcceptOwnership_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "noble/authority/v1/tx.proto",
}
