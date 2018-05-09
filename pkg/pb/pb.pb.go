// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_ebdceedf41c8d1ba, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type PublicKey struct {
	PublicKey            string   `protobuf:"bytes,1,opt,name=public_key,json=publicKey" json:"public_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublicKey) Reset()         { *m = PublicKey{} }
func (m *PublicKey) String() string { return proto.CompactTextString(m) }
func (*PublicKey) ProtoMessage()    {}
func (*PublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_ebdceedf41c8d1ba, []int{1}
}
func (m *PublicKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublicKey.Unmarshal(m, b)
}
func (m *PublicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublicKey.Marshal(b, m, deterministic)
}
func (dst *PublicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicKey.Merge(dst, src)
}
func (m *PublicKey) XXX_Size() int {
	return xxx_messageInfo_PublicKey.Size(m)
}
func (m *PublicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicKey.DiscardUnknown(m)
}

var xxx_messageInfo_PublicKey proto.InternalMessageInfo

func (m *PublicKey) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

type Account struct {
	Meta                 *AccountMeta `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	PublicKey            string       `protobuf:"bytes,2,opt,name=public_key,json=publicKey" json:"public_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_ebdceedf41c8d1ba, []int{2}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (dst *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(dst, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetMeta() *AccountMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *Account) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

type AccountMeta struct {
	Email                string   `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	FullName             string   `protobuf:"bytes,2,opt,name=full_name,json=fullName" json:"full_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountMeta) Reset()         { *m = AccountMeta{} }
func (m *AccountMeta) String() string { return proto.CompactTextString(m) }
func (*AccountMeta) ProtoMessage()    {}
func (*AccountMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_ebdceedf41c8d1ba, []int{3}
}
func (m *AccountMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountMeta.Unmarshal(m, b)
}
func (m *AccountMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountMeta.Marshal(b, m, deterministic)
}
func (dst *AccountMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountMeta.Merge(dst, src)
}
func (m *AccountMeta) XXX_Size() int {
	return xxx_messageInfo_AccountMeta.Size(m)
}
func (m *AccountMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountMeta.DiscardUnknown(m)
}

var xxx_messageInfo_AccountMeta proto.InternalMessageInfo

func (m *AccountMeta) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AccountMeta) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

type Candidate struct {
	Email                string   `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	FullName             string   `protobuf:"bytes,2,opt,name=full_name,json=fullName" json:"full_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Candidate) Reset()         { *m = Candidate{} }
func (m *Candidate) String() string { return proto.CompactTextString(m) }
func (*Candidate) ProtoMessage()    {}
func (*Candidate) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_ebdceedf41c8d1ba, []int{4}
}
func (m *Candidate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Candidate.Unmarshal(m, b)
}
func (m *Candidate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Candidate.Marshal(b, m, deterministic)
}
func (dst *Candidate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Candidate.Merge(dst, src)
}
func (m *Candidate) XXX_Size() int {
	return xxx_messageInfo_Candidate.Size(m)
}
func (m *Candidate) XXX_DiscardUnknown() {
	xxx_messageInfo_Candidate.DiscardUnknown(m)
}

var xxx_messageInfo_Candidate proto.InternalMessageInfo

func (m *Candidate) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Candidate) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

type Block struct {
	Index                int64    `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
	Hash                 string   `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	PrevHash             string   `protobuf:"bytes,3,opt,name=prev_hash,json=prevHash" json:"prev_hash,omitempty"`
	Timestamp            int64    `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_ebdceedf41c8d1ba, []int{5}
}
func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (dst *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(dst, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Block) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Block) GetPrevHash() string {
	if m != nil {
		return m.PrevHash
	}
	return ""
}

func (m *Block) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type NetOpts struct {
	Broker               *BrokerOpts     `protobuf:"bytes,1,opt,name=broker" json:"broker,omitempty"`
	Membership           *MembershipOpts `protobuf:"bytes,2,opt,name=membership" json:"membership,omitempty"`
	Master               *MasterOpts     `protobuf:"bytes,3,opt,name=master" json:"master,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *NetOpts) Reset()         { *m = NetOpts{} }
func (m *NetOpts) String() string { return proto.CompactTextString(m) }
func (*NetOpts) ProtoMessage()    {}
func (*NetOpts) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_ebdceedf41c8d1ba, []int{6}
}
func (m *NetOpts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetOpts.Unmarshal(m, b)
}
func (m *NetOpts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetOpts.Marshal(b, m, deterministic)
}
func (dst *NetOpts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetOpts.Merge(dst, src)
}
func (m *NetOpts) XXX_Size() int {
	return xxx_messageInfo_NetOpts.Size(m)
}
func (m *NetOpts) XXX_DiscardUnknown() {
	xxx_messageInfo_NetOpts.DiscardUnknown(m)
}

var xxx_messageInfo_NetOpts proto.InternalMessageInfo

func (m *NetOpts) GetBroker() *BrokerOpts {
	if m != nil {
		return m.Broker
	}
	return nil
}

func (m *NetOpts) GetMembership() *MembershipOpts {
	if m != nil {
		return m.Membership
	}
	return nil
}

func (m *NetOpts) GetMaster() *MasterOpts {
	if m != nil {
		return m.Master
	}
	return nil
}

type BrokerOpts struct {
	Endpoint             string   `protobuf:"bytes,1,opt,name=endpoint" json:"endpoint,omitempty"`
	User                 string   `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BrokerOpts) Reset()         { *m = BrokerOpts{} }
func (m *BrokerOpts) String() string { return proto.CompactTextString(m) }
func (*BrokerOpts) ProtoMessage()    {}
func (*BrokerOpts) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_ebdceedf41c8d1ba, []int{7}
}
func (m *BrokerOpts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BrokerOpts.Unmarshal(m, b)
}
func (m *BrokerOpts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BrokerOpts.Marshal(b, m, deterministic)
}
func (dst *BrokerOpts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BrokerOpts.Merge(dst, src)
}
func (m *BrokerOpts) XXX_Size() int {
	return xxx_messageInfo_BrokerOpts.Size(m)
}
func (m *BrokerOpts) XXX_DiscardUnknown() {
	xxx_messageInfo_BrokerOpts.DiscardUnknown(m)
}

var xxx_messageInfo_BrokerOpts proto.InternalMessageInfo

func (m *BrokerOpts) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *BrokerOpts) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *BrokerOpts) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type MembershipOpts struct {
	Endpoint             string   `protobuf:"bytes,1,opt,name=endpoint" json:"endpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MembershipOpts) Reset()         { *m = MembershipOpts{} }
func (m *MembershipOpts) String() string { return proto.CompactTextString(m) }
func (*MembershipOpts) ProtoMessage()    {}
func (*MembershipOpts) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_ebdceedf41c8d1ba, []int{8}
}
func (m *MembershipOpts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MembershipOpts.Unmarshal(m, b)
}
func (m *MembershipOpts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MembershipOpts.Marshal(b, m, deterministic)
}
func (dst *MembershipOpts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MembershipOpts.Merge(dst, src)
}
func (m *MembershipOpts) XXX_Size() int {
	return xxx_messageInfo_MembershipOpts.Size(m)
}
func (m *MembershipOpts) XXX_DiscardUnknown() {
	xxx_messageInfo_MembershipOpts.DiscardUnknown(m)
}

var xxx_messageInfo_MembershipOpts proto.InternalMessageInfo

func (m *MembershipOpts) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

type MasterOpts struct {
	Endpoint             string   `protobuf:"bytes,1,opt,name=endpoint" json:"endpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MasterOpts) Reset()         { *m = MasterOpts{} }
func (m *MasterOpts) String() string { return proto.CompactTextString(m) }
func (*MasterOpts) ProtoMessage()    {}
func (*MasterOpts) Descriptor() ([]byte, []int) {
	return fileDescriptor_pb_ebdceedf41c8d1ba, []int{9}
}
func (m *MasterOpts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MasterOpts.Unmarshal(m, b)
}
func (m *MasterOpts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MasterOpts.Marshal(b, m, deterministic)
}
func (dst *MasterOpts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MasterOpts.Merge(dst, src)
}
func (m *MasterOpts) XXX_Size() int {
	return xxx_messageInfo_MasterOpts.Size(m)
}
func (m *MasterOpts) XXX_DiscardUnknown() {
	xxx_messageInfo_MasterOpts.DiscardUnknown(m)
}

var xxx_messageInfo_MasterOpts proto.InternalMessageInfo

func (m *MasterOpts) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "pb.Empty")
	proto.RegisterType((*PublicKey)(nil), "pb.PublicKey")
	proto.RegisterType((*Account)(nil), "pb.Account")
	proto.RegisterType((*AccountMeta)(nil), "pb.AccountMeta")
	proto.RegisterType((*Candidate)(nil), "pb.Candidate")
	proto.RegisterType((*Block)(nil), "pb.Block")
	proto.RegisterType((*NetOpts)(nil), "pb.NetOpts")
	proto.RegisterType((*BrokerOpts)(nil), "pb.BrokerOpts")
	proto.RegisterType((*MembershipOpts)(nil), "pb.MembershipOpts")
	proto.RegisterType((*MasterOpts)(nil), "pb.MasterOpts")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Membership service

type MembershipClient interface {
	Register(ctx context.Context, in *Candidate, opts ...grpc.CallOption) (*Empty, error)
	Fetch(ctx context.Context, in *PublicKey, opts ...grpc.CallOption) (*Account, error)
}

type membershipClient struct {
	cc *grpc.ClientConn
}

func NewMembershipClient(cc *grpc.ClientConn) MembershipClient {
	return &membershipClient{cc}
}

func (c *membershipClient) Register(ctx context.Context, in *Candidate, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/pb.Membership/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *membershipClient) Fetch(ctx context.Context, in *PublicKey, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := grpc.Invoke(ctx, "/pb.Membership/Fetch", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Membership service

type MembershipServer interface {
	Register(context.Context, *Candidate) (*Empty, error)
	Fetch(context.Context, *PublicKey) (*Account, error)
}

func RegisterMembershipServer(s *grpc.Server, srv MembershipServer) {
	s.RegisterService(&_Membership_serviceDesc, srv)
}

func _Membership_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Candidate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MembershipServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Membership/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MembershipServer).Register(ctx, req.(*Candidate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Membership_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublicKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MembershipServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Membership/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MembershipServer).Fetch(ctx, req.(*PublicKey))
	}
	return interceptor(ctx, in, info, handler)
}

var _Membership_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Membership",
	HandlerType: (*MembershipServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Membership_Register_Handler,
		},
		{
			MethodName: "Fetch",
			Handler:    _Membership_Fetch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb.proto",
}

// Client API for Master service

type MasterClient interface {
	Candidate(ctx context.Context, in *Block, opts ...grpc.CallOption) (*Empty, error)
	Node(ctx context.Context, in *PublicKey, opts ...grpc.CallOption) (*Empty, error)
}

type masterClient struct {
	cc *grpc.ClientConn
}

func NewMasterClient(cc *grpc.ClientConn) MasterClient {
	return &masterClient{cc}
}

func (c *masterClient) Candidate(ctx context.Context, in *Block, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/pb.Master/Candidate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) Node(ctx context.Context, in *PublicKey, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/pb.Master/Node", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Master service

type MasterServer interface {
	Candidate(context.Context, *Block) (*Empty, error)
	Node(context.Context, *PublicKey) (*Empty, error)
}

func RegisterMasterServer(s *grpc.Server, srv MasterServer) {
	s.RegisterService(&_Master_serviceDesc, srv)
}

func _Master_Candidate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Block)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).Candidate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Master/Candidate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).Candidate(ctx, req.(*Block))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_Node_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublicKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).Node(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Master/Node",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).Node(ctx, req.(*PublicKey))
	}
	return interceptor(ctx, in, info, handler)
}

var _Master_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Master",
	HandlerType: (*MasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Candidate",
			Handler:    _Master_Candidate_Handler,
		},
		{
			MethodName: "Node",
			Handler:    _Master_Node_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb.proto",
}

// Client API for Discovery service

type DiscoveryClient interface {
	Network(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NetOpts, error)
}

type discoveryClient struct {
	cc *grpc.ClientConn
}

func NewDiscoveryClient(cc *grpc.ClientConn) DiscoveryClient {
	return &discoveryClient{cc}
}

func (c *discoveryClient) Network(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NetOpts, error) {
	out := new(NetOpts)
	err := grpc.Invoke(ctx, "/pb.Discovery/Network", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Discovery service

type DiscoveryServer interface {
	Network(context.Context, *Empty) (*NetOpts, error)
}

func RegisterDiscoveryServer(s *grpc.Server, srv DiscoveryServer) {
	s.RegisterService(&_Discovery_serviceDesc, srv)
}

func _Discovery_Network_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServer).Network(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Discovery/Network",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServer).Network(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Discovery_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Discovery",
	HandlerType: (*DiscoveryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Network",
			Handler:    _Discovery_Network_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb.proto",
}

func init() { proto.RegisterFile("pb.proto", fileDescriptor_pb_ebdceedf41c8d1ba) }

var fileDescriptor_pb_ebdceedf41c8d1ba = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x4b, 0x8f, 0xd3, 0x30,
	0x10, 0xc7, 0xdb, 0xdd, 0xbe, 0x32, 0x15, 0x8b, 0x64, 0x71, 0x58, 0x15, 0x90, 0x90, 0x2b, 0xd0,
	0x0a, 0xa1, 0x0a, 0x85, 0x3b, 0x82, 0xe5, 0x21, 0x24, 0x94, 0x02, 0x39, 0x21, 0x71, 0xa8, 0x9c,
	0x64, 0xa0, 0x56, 0xe3, 0xd8, 0xb2, 0xdd, 0x5d, 0xfa, 0x21, 0xf8, 0xce, 0xc8, 0x8f, 0x4d, 0xb2,
	0x3d, 0xec, 0x61, 0x6f, 0x9e, 0xf9, 0xff, 0xe7, 0xe7, 0xf1, 0x64, 0x02, 0x33, 0x55, 0xac, 0x94,
	0x96, 0x56, 0x92, 0x13, 0x55, 0xd0, 0x29, 0x8c, 0x3f, 0x09, 0x65, 0x0f, 0xf4, 0x25, 0x24, 0xdf,
	0xf7, 0x45, 0xcd, 0xcb, 0xaf, 0x78, 0x20, 0x4f, 0x01, 0x94, 0x0f, 0x36, 0x3b, 0x3c, 0x9c, 0x0f,
	0x9f, 0x0d, 0x2f, 0x92, 0x3c, 0x51, 0x37, 0x32, 0xcd, 0x60, 0xfa, 0xbe, 0x2c, 0xe5, 0xbe, 0xb1,
	0x64, 0x09, 0x23, 0x81, 0x96, 0x79, 0xcf, 0x3c, 0x7d, 0xb8, 0x52, 0xc5, 0x2a, 0x4a, 0x19, 0x5a,
	0x96, 0x7b, 0xf1, 0x08, 0x77, 0x72, 0x8c, 0x7b, 0x07, 0xf3, 0x5e, 0x0d, 0x79, 0x04, 0x63, 0x14,
	0x8c, 0xd7, 0xf1, 0xde, 0x10, 0x90, 0xc7, 0x90, 0xfc, 0xde, 0xd7, 0xf5, 0xa6, 0x61, 0x02, 0x23,
	0x62, 0xe6, 0x12, 0x6b, 0x26, 0x90, 0xbe, 0x85, 0xe4, 0x03, 0x6b, 0x2a, 0x5e, 0x31, 0x8b, 0xf7,
	0xa9, 0xaf, 0x61, 0x7c, 0x59, 0xcb, 0x72, 0xe7, 0x6a, 0x79, 0x53, 0xe1, 0x5f, 0x5f, 0x7b, 0x9a,
	0x87, 0x80, 0x10, 0x18, 0x6d, 0x99, 0xd9, 0xc6, 0x32, 0x7f, 0x76, 0x3c, 0xa5, 0xf1, 0x6a, 0xe3,
	0x85, 0xd3, 0xc0, 0x73, 0x89, 0x2f, 0x4e, 0x7c, 0x02, 0x89, 0xe5, 0x02, 0x8d, 0x65, 0x42, 0x9d,
	0x8f, 0x3c, 0xaa, 0x4b, 0xd0, 0x7f, 0x43, 0x98, 0xae, 0xd1, 0x7e, 0x53, 0xd6, 0x90, 0x17, 0x30,
	0x29, 0xb4, 0xdc, 0xa1, 0x8e, 0x13, 0x3c, 0x73, 0x13, 0xbc, 0xf4, 0x19, 0xa7, 0xe7, 0x51, 0x25,
	0x29, 0x80, 0x40, 0x51, 0xa0, 0x36, 0x5b, 0xae, 0x7c, 0x23, 0xf3, 0x94, 0x38, 0x6f, 0xd6, 0x66,
	0xbd, 0xbf, 0xe7, 0x72, 0x6c, 0xc1, 0x8c, 0x45, 0xed, 0xfb, 0x8b, 0xec, 0xcc, 0x67, 0x02, 0x3b,
	0xa8, 0xf4, 0x27, 0x40, 0x77, 0x23, 0x59, 0xc0, 0x0c, 0x9b, 0x4a, 0x49, 0xde, 0xd8, 0x38, 0xc1,
	0x36, 0x76, 0x83, 0xd8, 0x1b, 0xd4, 0x37, 0x83, 0x70, 0x67, 0xe7, 0x57, 0xcc, 0x98, 0x6b, 0xa9,
	0xab, 0x76, 0x0e, 0x31, 0xa6, 0xaf, 0xe0, 0xec, 0x76, 0x7f, 0x77, 0xd1, 0xe9, 0x05, 0x40, 0xd7,
	0xdd, 0x5d, 0xce, 0xf4, 0x17, 0x40, 0xd6, 0x7f, 0xe7, 0x2c, 0xc7, 0x3f, 0xdc, 0x55, 0x92, 0x07,
	0xee, 0x8d, 0xed, 0x2e, 0x2c, 0x12, 0x17, 0x86, 0x05, 0x1f, 0x90, 0xe7, 0x30, 0xfe, 0x8c, 0xb6,
	0xdc, 0x06, 0x53, 0xbb, 0xed, 0x8b, 0x79, 0x6f, 0x6b, 0xe9, 0x20, 0xfd, 0x01, 0x93, 0xd0, 0x06,
	0x59, 0xf6, 0xd7, 0xca, 0xa3, 0xfc, 0x96, 0xdc, 0xa6, 0x52, 0x18, 0xad, 0x65, 0x85, 0xc7, 0xd0,
	0xbe, 0x27, 0x7d, 0x0d, 0xc9, 0x47, 0x6e, 0x4a, 0x79, 0x85, 0xfa, 0x40, 0x96, 0xfe, 0xeb, 0x5f,
	0x4b, 0xbd, 0x23, 0x9d, 0x29, 0x34, 0x11, 0xb7, 0x82, 0x0e, 0x8a, 0x89, 0xff, 0x45, 0xdf, 0xfc,
	0x0f, 0x00, 0x00, 0xff, 0xff, 0x83, 0x83, 0xc0, 0xe7, 0xae, 0x03, 0x00, 0x00,
}
