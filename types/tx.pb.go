// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: noble/authority/v1/tx.proto

package types

import (
	context "context"
	types1 "cosmossdk.io/x/upgrade/types"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgExecute struct {
	Signer   string       `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	Messages []*types.Any `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (m *MsgExecute) Reset()         { *m = MsgExecute{} }
func (m *MsgExecute) String() string { return proto.CompactTextString(m) }
func (*MsgExecute) ProtoMessage()    {}
func (*MsgExecute) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7f9675ad166e349, []int{0}
}
func (m *MsgExecute) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgExecute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgExecute.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgExecute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgExecute.Merge(m, src)
}
func (m *MsgExecute) XXX_Size() int {
	return m.Size()
}
func (m *MsgExecute) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgExecute.DiscardUnknown(m)
}

var xxx_messageInfo_MsgExecute proto.InternalMessageInfo

type MsgExecuteResponse struct {
	Results [][]byte `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (m *MsgExecuteResponse) Reset()         { *m = MsgExecuteResponse{} }
func (m *MsgExecuteResponse) String() string { return proto.CompactTextString(m) }
func (*MsgExecuteResponse) ProtoMessage()    {}
func (*MsgExecuteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7f9675ad166e349, []int{1}
}
func (m *MsgExecuteResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgExecuteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgExecuteResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgExecuteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgExecuteResponse.Merge(m, src)
}
func (m *MsgExecuteResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgExecuteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgExecuteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgExecuteResponse proto.InternalMessageInfo

func (m *MsgExecuteResponse) GetResults() [][]byte {
	if m != nil {
		return m.Results
	}
	return nil
}

type MsgTransferOwnership struct {
	Signer   string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	NewOwner string `protobuf:"bytes,2,opt,name=new_owner,json=newOwner,proto3" json:"new_owner,omitempty"`
}

func (m *MsgTransferOwnership) Reset()         { *m = MsgTransferOwnership{} }
func (m *MsgTransferOwnership) String() string { return proto.CompactTextString(m) }
func (*MsgTransferOwnership) ProtoMessage()    {}
func (*MsgTransferOwnership) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7f9675ad166e349, []int{2}
}
func (m *MsgTransferOwnership) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTransferOwnership) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTransferOwnership.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTransferOwnership) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTransferOwnership.Merge(m, src)
}
func (m *MsgTransferOwnership) XXX_Size() int {
	return m.Size()
}
func (m *MsgTransferOwnership) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTransferOwnership.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTransferOwnership proto.InternalMessageInfo

type MsgTransferOwnershipResponse struct {
}

func (m *MsgTransferOwnershipResponse) Reset()         { *m = MsgTransferOwnershipResponse{} }
func (m *MsgTransferOwnershipResponse) String() string { return proto.CompactTextString(m) }
func (*MsgTransferOwnershipResponse) ProtoMessage()    {}
func (*MsgTransferOwnershipResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7f9675ad166e349, []int{3}
}
func (m *MsgTransferOwnershipResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTransferOwnershipResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTransferOwnershipResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTransferOwnershipResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTransferOwnershipResponse.Merge(m, src)
}
func (m *MsgTransferOwnershipResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgTransferOwnershipResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTransferOwnershipResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTransferOwnershipResponse proto.InternalMessageInfo

type MsgAcceptOwnership struct {
	Signer string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
}

func (m *MsgAcceptOwnership) Reset()         { *m = MsgAcceptOwnership{} }
func (m *MsgAcceptOwnership) String() string { return proto.CompactTextString(m) }
func (*MsgAcceptOwnership) ProtoMessage()    {}
func (*MsgAcceptOwnership) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7f9675ad166e349, []int{4}
}
func (m *MsgAcceptOwnership) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAcceptOwnership) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAcceptOwnership.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAcceptOwnership) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAcceptOwnership.Merge(m, src)
}
func (m *MsgAcceptOwnership) XXX_Size() int {
	return m.Size()
}
func (m *MsgAcceptOwnership) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAcceptOwnership.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAcceptOwnership proto.InternalMessageInfo

type MsgAcceptOwnershipResponse struct {
}

func (m *MsgAcceptOwnershipResponse) Reset()         { *m = MsgAcceptOwnershipResponse{} }
func (m *MsgAcceptOwnershipResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAcceptOwnershipResponse) ProtoMessage()    {}
func (*MsgAcceptOwnershipResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7f9675ad166e349, []int{5}
}
func (m *MsgAcceptOwnershipResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAcceptOwnershipResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAcceptOwnershipResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAcceptOwnershipResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAcceptOwnershipResponse.Merge(m, src)
}
func (m *MsgAcceptOwnershipResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAcceptOwnershipResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAcceptOwnershipResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAcceptOwnershipResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgExecute)(nil), "noble.authority.v1.MsgExecute")
	proto.RegisterType((*MsgExecuteResponse)(nil), "noble.authority.v1.MsgExecuteResponse")
	proto.RegisterType((*MsgTransferOwnership)(nil), "noble.authority.v1.MsgTransferOwnership")
	proto.RegisterType((*MsgTransferOwnershipResponse)(nil), "noble.authority.v1.MsgTransferOwnershipResponse")
	proto.RegisterType((*MsgAcceptOwnership)(nil), "noble.authority.v1.MsgAcceptOwnership")
	proto.RegisterType((*MsgAcceptOwnershipResponse)(nil), "noble.authority.v1.MsgAcceptOwnershipResponse")
}

func init() { proto.RegisterFile("noble/authority/v1/tx.proto", fileDescriptor_a7f9675ad166e349) }

var fileDescriptor_a7f9675ad166e349 = []byte{
	// 549 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4f, 0x8f, 0xd2, 0x5e,
	0x14, 0xa5, 0x43, 0x7e, 0xf3, 0xe7, 0xfd, 0x34, 0x93, 0x69, 0x30, 0x53, 0xeb, 0xa4, 0x90, 0x26,
	0x2a, 0x21, 0xe1, 0x15, 0x30, 0x6e, 0xd8, 0x31, 0x89, 0x4b, 0x62, 0x64, 0x74, 0xe3, 0x66, 0x52,
	0xe0, 0xf2, 0x68, 0x42, 0xdf, 0xab, 0xef, 0xbe, 0xc2, 0xb0, 0x33, 0xae, 0x8c, 0x2b, 0x77, 0x6e,
	0xe7, 0x13, 0x18, 0x16, 0x7e, 0x08, 0x97, 0x13, 0x57, 0x2e, 0x0d, 0x2c, 0xf0, 0x63, 0x18, 0xfa,
	0x07, 0x0c, 0x88, 0x12, 0xdd, 0x34, 0xbd, 0x3d, 0xe7, 0xbe, 0x73, 0xee, 0xbd, 0xef, 0x96, 0xdc,
	0xe3, 0xa2, 0x3d, 0x00, 0xc7, 0x0d, 0x55, 0x5f, 0x48, 0x4f, 0x8d, 0x9d, 0x61, 0xd5, 0x51, 0x57,
	0x34, 0x90, 0x42, 0x09, 0x5d, 0x8f, 0x40, 0xba, 0x04, 0xe9, 0xb0, 0x6a, 0x9e, 0xb8, 0xbe, 0xc7,
	0x85, 0x13, 0x3d, 0x63, 0x9a, 0x79, 0xda, 0x11, 0xe8, 0x0b, 0x74, 0x7c, 0x64, 0x8b, 0x74, 0x1f,
	0x59, 0x02, 0xe4, 0x13, 0x20, 0x0c, 0x98, 0x74, 0xbb, 0xe0, 0x0c, 0xab, 0x6d, 0x50, 0xee, 0x4a,
	0xc0, 0xbc, 0x1b, 0x13, 0x2e, 0xa3, 0xc8, 0x89, 0x83, 0x04, 0xca, 0x31, 0xc1, 0x44, 0xfc, 0x7d,
	0xf1, 0x96, 0x26, 0x30, 0x21, 0xd8, 0x00, 0x9c, 0x28, 0x6a, 0x87, 0x3d, 0xc7, 0xe5, 0xe3, 0x18,
	0xb2, 0x3f, 0x68, 0x84, 0x34, 0x91, 0x3d, 0xb9, 0x82, 0x4e, 0xa8, 0x40, 0xaf, 0x90, 0x7d, 0xf4,
	0x18, 0x07, 0x69, 0x68, 0x05, 0xad, 0x78, 0x74, 0x6e, 0x7c, 0xf9, 0x54, 0xce, 0x25, 0x0a, 0x8d,
	0x6e, 0x57, 0x02, 0xe2, 0x85, 0x92, 0x1e, 0x67, 0xad, 0x84, 0xa7, 0x57, 0xc8, 0xa1, 0x0f, 0x88,
	0x2e, 0x03, 0x34, 0xf6, 0x0a, 0xd9, 0xe2, 0xff, 0xb5, 0x1c, 0x8d, 0xe5, 0x68, 0x2a, 0x47, 0x1b,
	0x7c, 0xdc, 0x5a, 0xb2, 0xea, 0xf7, 0xdf, 0x5e, 0xe7, 0x33, 0xdf, 0xaf, 0xf3, 0x99, 0x37, 0xf3,
	0x49, 0x29, 0x39, 0xe6, 0xdd, 0x7c, 0x52, 0xba, 0x1d, 0x37, 0x35, 0xb1, 0x62, 0x53, 0xa2, 0xaf,
	0x8c, 0xb5, 0x00, 0x03, 0xc1, 0x11, 0x74, 0x83, 0x1c, 0x48, 0xc0, 0x70, 0xa0, 0xd0, 0xd0, 0x0a,
	0xd9, 0xe2, 0xad, 0x56, 0x1a, 0xda, 0x13, 0x8d, 0xe4, 0x9a, 0xc8, 0x9e, 0x4b, 0x97, 0x63, 0x0f,
	0xe4, 0xd3, 0x11, 0x07, 0x89, 0x7d, 0x2f, 0xf8, 0x8b, 0x9a, 0x1e, 0x93, 0x23, 0x0e, 0xa3, 0x4b,
	0xb1, 0x38, 0xc2, 0xd8, 0xfb, 0x43, 0xd2, 0x21, 0x87, 0x51, 0x24, 0x56, 0x77, 0xb6, 0x14, 0x76,
	0x1a, 0x17, 0xb6, 0xe1, 0xcc, 0xb6, 0xc8, 0xd9, 0xaf, 0x1c, 0xa7, 0xc5, 0xda, 0x61, 0xd4, 0x82,
	0x46, 0xa7, 0x03, 0x81, 0xfa, 0x87, 0x7a, 0xea, 0xe5, 0x2d, 0xc6, 0xee, 0xc4, 0xc6, 0xd6, 0x04,
	0xec, 0x33, 0x62, 0x6e, 0xca, 0xa6, 0xa6, 0x6a, 0x1f, 0xb3, 0x24, 0xdb, 0x44, 0xa6, 0x3f, 0x23,
	0x07, 0xe9, 0xad, 0xb1, 0xe8, 0xe6, 0x95, 0xa7, 0xab, 0xe1, 0x99, 0x0f, 0x7e, 0x8f, 0x2f, 0x87,
	0x2b, 0xc8, 0xc9, 0xe6, 0xf8, 0x8a, 0x5b, 0x92, 0x37, 0x98, 0x66, 0x65, 0x57, 0xe6, 0x52, 0xd0,
	0x23, 0xc7, 0xeb, 0xdd, 0xdd, 0xe6, 0x75, 0x8d, 0x67, 0xd2, 0xdd, 0x78, 0x4b, 0xa9, 0x57, 0xe4,
	0xf8, 0x42, 0xf4, 0xd4, 0xc8, 0x95, 0xf0, 0x22, 0x5e, 0x6c, 0xbd, 0x44, 0x93, 0xa9, 0x25, 0x9b,
	0x4e, 0x93, 0x4d, 0x5f, 0x1c, 0xb3, 0xc6, 0x35, 0x6b, 0xbb, 0x73, 0x53, 0x49, 0xf3, 0xbf, 0xd7,
	0xf3, 0x49, 0x49, 0x3b, 0x6f, 0x7c, 0x9e, 0x5a, 0xda, 0xcd, 0xd4, 0xd2, 0xbe, 0x4d, 0x2d, 0xed,
	0xfd, 0xcc, 0xca, 0xdc, 0xcc, 0xac, 0xcc, 0xd7, 0x99, 0x95, 0x79, 0xf9, 0x90, 0x79, 0xaa, 0x1f,
	0xb6, 0x69, 0x47, 0xf8, 0x4e, 0x54, 0x4d, 0xd9, 0x45, 0x04, 0x85, 0x3f, 0xfd, 0xd8, 0xd4, 0x38,
	0x00, 0x6c, 0xef, 0x47, 0xab, 0xfc, 0xe8, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdf, 0xc6, 0x6d,
	0xd4, 0xf8, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	Execute(ctx context.Context, in *MsgExecute, opts ...grpc.CallOption) (*MsgExecuteResponse, error)
	TransferOwnership(ctx context.Context, in *MsgTransferOwnership, opts ...grpc.CallOption) (*MsgTransferOwnershipResponse, error)
	AcceptOwnership(ctx context.Context, in *MsgAcceptOwnership, opts ...grpc.CallOption) (*MsgAcceptOwnershipResponse, error)
	SoftwareUpgrade(ctx context.Context, in *types1.MsgSoftwareUpgrade, opts ...grpc.CallOption) (*types1.MsgSoftwareUpgradeResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Execute(ctx context.Context, in *MsgExecute, opts ...grpc.CallOption) (*MsgExecuteResponse, error) {
	out := new(MsgExecuteResponse)
	err := c.cc.Invoke(ctx, "/noble.authority.v1.Msg/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) TransferOwnership(ctx context.Context, in *MsgTransferOwnership, opts ...grpc.CallOption) (*MsgTransferOwnershipResponse, error) {
	out := new(MsgTransferOwnershipResponse)
	err := c.cc.Invoke(ctx, "/noble.authority.v1.Msg/TransferOwnership", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AcceptOwnership(ctx context.Context, in *MsgAcceptOwnership, opts ...grpc.CallOption) (*MsgAcceptOwnershipResponse, error) {
	out := new(MsgAcceptOwnershipResponse)
	err := c.cc.Invoke(ctx, "/noble.authority.v1.Msg/AcceptOwnership", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SoftwareUpgrade(ctx context.Context, in *types1.MsgSoftwareUpgrade, opts ...grpc.CallOption) (*types1.MsgSoftwareUpgradeResponse, error) {
	out := new(types1.MsgSoftwareUpgradeResponse)
	err := c.cc.Invoke(ctx, "/noble.authority.v1.Msg/SoftwareUpgrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Execute(context.Context, *MsgExecute) (*MsgExecuteResponse, error)
	TransferOwnership(context.Context, *MsgTransferOwnership) (*MsgTransferOwnershipResponse, error)
	AcceptOwnership(context.Context, *MsgAcceptOwnership) (*MsgAcceptOwnershipResponse, error)
	SoftwareUpgrade(context.Context, *types1.MsgSoftwareUpgrade) (*types1.MsgSoftwareUpgradeResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Execute(ctx context.Context, req *MsgExecute) (*MsgExecuteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (*UnimplementedMsgServer) TransferOwnership(ctx context.Context, req *MsgTransferOwnership) (*MsgTransferOwnershipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferOwnership not implemented")
}
func (*UnimplementedMsgServer) AcceptOwnership(ctx context.Context, req *MsgAcceptOwnership) (*MsgAcceptOwnershipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptOwnership not implemented")
}
func (*UnimplementedMsgServer) SoftwareUpgrade(ctx context.Context, req *types1.MsgSoftwareUpgrade) (*types1.MsgSoftwareUpgradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SoftwareUpgrade not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
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
		FullMethod: "/noble.authority.v1.Msg/Execute",
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
		FullMethod: "/noble.authority.v1.Msg/TransferOwnership",
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
		FullMethod: "/noble.authority.v1.Msg/AcceptOwnership",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AcceptOwnership(ctx, req.(*MsgAcceptOwnership))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SoftwareUpgrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types1.MsgSoftwareUpgrade)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SoftwareUpgrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/noble.authority.v1.Msg/SoftwareUpgrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SoftwareUpgrade(ctx, req.(*types1.MsgSoftwareUpgrade))
	}
	return interceptor(ctx, in, info, handler)
}

var Msg_serviceDesc = _Msg_serviceDesc
var _Msg_serviceDesc = grpc.ServiceDesc{
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
		{
			MethodName: "SoftwareUpgrade",
			Handler:    _Msg_SoftwareUpgrade_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "noble/authority/v1/tx.proto",
}

func (m *MsgExecute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgExecute) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgExecute) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Messages) > 0 {
		for iNdEx := len(m.Messages) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Messages[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgExecuteResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgExecuteResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgExecuteResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Results) > 0 {
		for iNdEx := len(m.Results) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Results[iNdEx])
			copy(dAtA[i:], m.Results[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Results[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MsgTransferOwnership) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTransferOwnership) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTransferOwnership) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NewOwner) > 0 {
		i -= len(m.NewOwner)
		copy(dAtA[i:], m.NewOwner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.NewOwner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgTransferOwnershipResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTransferOwnershipResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTransferOwnershipResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgAcceptOwnership) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAcceptOwnership) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAcceptOwnership) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAcceptOwnershipResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAcceptOwnershipResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAcceptOwnershipResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgExecute) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Messages) > 0 {
		for _, e := range m.Messages {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgExecuteResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Results) > 0 {
		for _, b := range m.Results {
			l = len(b)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgTransferOwnership) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.NewOwner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgTransferOwnershipResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgAcceptOwnership) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgAcceptOwnershipResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgExecute) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgExecute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgExecute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Messages", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Messages = append(m.Messages, &types.Any{})
			if err := m.Messages[len(m.Messages)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgExecuteResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgExecuteResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgExecuteResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Results", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Results = append(m.Results, make([]byte, postIndex-iNdEx))
			copy(m.Results[len(m.Results)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgTransferOwnership) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgTransferOwnership: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTransferOwnership: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewOwner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgTransferOwnershipResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgTransferOwnershipResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTransferOwnershipResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgAcceptOwnership) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgAcceptOwnership: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAcceptOwnership: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgAcceptOwnershipResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgAcceptOwnershipResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAcceptOwnershipResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
