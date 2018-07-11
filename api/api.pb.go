// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

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

type AccountReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=lastName,proto3" json:"lastName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountReq) Reset()         { *m = AccountReq{} }
func (m *AccountReq) String() string { return proto.CompactTextString(m) }
func (*AccountReq) ProtoMessage()    {}
func (*AccountReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_0bf5ef62db4ac6b4, []int{0}
}
func (m *AccountReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountReq.Unmarshal(m, b)
}
func (m *AccountReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountReq.Marshal(b, m, deterministic)
}
func (dst *AccountReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountReq.Merge(dst, src)
}
func (m *AccountReq) XXX_Size() int {
	return xxx_messageInfo_AccountReq.Size(m)
}
func (m *AccountReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountReq.DiscardUnknown(m)
}

var xxx_messageInfo_AccountReq proto.InternalMessageInfo

func (m *AccountReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AccountReq) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type ID struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ID) Reset()         { *m = ID{} }
func (m *ID) String() string { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()    {}
func (*ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_0bf5ef62db4ac6b4, []int{1}
}
func (m *ID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ID.Unmarshal(m, b)
}
func (m *ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ID.Marshal(b, m, deterministic)
}
func (dst *ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ID.Merge(dst, src)
}
func (m *ID) XXX_Size() int {
	return xxx_messageInfo_ID.Size(m)
}
func (m *ID) XXX_DiscardUnknown() {
	xxx_messageInfo_ID.DiscardUnknown(m)
}

var xxx_messageInfo_ID proto.InternalMessageInfo

func (m *ID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Credit struct {
	Credit               int64    `protobuf:"varint,1,opt,name=credit,proto3" json:"credit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Credit) Reset()         { *m = Credit{} }
func (m *Credit) String() string { return proto.CompactTextString(m) }
func (*Credit) ProtoMessage()    {}
func (*Credit) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_0bf5ef62db4ac6b4, []int{2}
}
func (m *Credit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Credit.Unmarshal(m, b)
}
func (m *Credit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Credit.Marshal(b, m, deterministic)
}
func (dst *Credit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Credit.Merge(dst, src)
}
func (m *Credit) XXX_Size() int {
	return xxx_messageInfo_Credit.Size(m)
}
func (m *Credit) XXX_DiscardUnknown() {
	xxx_messageInfo_Credit.DiscardUnknown(m)
}

var xxx_messageInfo_Credit proto.InternalMessageInfo

func (m *Credit) GetCredit() int64 {
	if m != nil {
		return m.Credit
	}
	return 0
}

type Account struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Credit               int64    `protobuf:"varint,4,opt,name=credit,proto3" json:"credit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_0bf5ef62db4ac6b4, []int{3}
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

func (m *Account) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Account) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Account) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Account) GetCredit() int64 {
	if m != nil {
		return m.Credit
	}
	return 0
}

type TransactionReq struct {
	SenderID             string   `protobuf:"bytes,1,opt,name=senderID,proto3" json:"senderID,omitempty"`
	RecvID               string   `protobuf:"bytes,2,opt,name=recvID,proto3" json:"recvID,omitempty"`
	Amount               int64    `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionReq) Reset()         { *m = TransactionReq{} }
func (m *TransactionReq) String() string { return proto.CompactTextString(m) }
func (*TransactionReq) ProtoMessage()    {}
func (*TransactionReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_0bf5ef62db4ac6b4, []int{4}
}
func (m *TransactionReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionReq.Unmarshal(m, b)
}
func (m *TransactionReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionReq.Marshal(b, m, deterministic)
}
func (dst *TransactionReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionReq.Merge(dst, src)
}
func (m *TransactionReq) XXX_Size() int {
	return xxx_messageInfo_TransactionReq.Size(m)
}
func (m *TransactionReq) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionReq.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionReq proto.InternalMessageInfo

func (m *TransactionReq) GetSenderID() string {
	if m != nil {
		return m.SenderID
	}
	return ""
}

func (m *TransactionReq) GetRecvID() string {
	if m != nil {
		return m.RecvID
	}
	return ""
}

func (m *TransactionReq) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type Transaction struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SenderID             string   `protobuf:"bytes,2,opt,name=senderID,proto3" json:"senderID,omitempty"`
	RecvID               string   `protobuf:"bytes,3,opt,name=recvID,proto3" json:"recvID,omitempty"`
	Amount               int64    `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_0bf5ef62db4ac6b4, []int{5}
}
func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (dst *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(dst, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Transaction) GetSenderID() string {
	if m != nil {
		return m.SenderID
	}
	return ""
}

func (m *Transaction) GetRecvID() string {
	if m != nil {
		return m.RecvID
	}
	return ""
}

func (m *Transaction) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func init() {
	proto.RegisterType((*AccountReq)(nil), "api.AccountReq")
	proto.RegisterType((*ID)(nil), "api.ID")
	proto.RegisterType((*Credit)(nil), "api.Credit")
	proto.RegisterType((*Account)(nil), "api.Account")
	proto.RegisterType((*TransactionReq)(nil), "api.TransactionReq")
	proto.RegisterType((*Transaction)(nil), "api.Transaction")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BankClient is the client API for Bank service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BankClient interface {
	NewAccount(ctx context.Context, in *AccountReq, opts ...grpc.CallOption) (*Account, error)
	GetCredit(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Credit, error)
	Transfer(ctx context.Context, in *TransactionReq, opts ...grpc.CallOption) (*Transaction, error)
}

type bankClient struct {
	cc *grpc.ClientConn
}

func NewBankClient(cc *grpc.ClientConn) BankClient {
	return &bankClient{cc}
}

func (c *bankClient) NewAccount(ctx context.Context, in *AccountReq, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/api.Bank/NewAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankClient) GetCredit(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Credit, error) {
	out := new(Credit)
	err := c.cc.Invoke(ctx, "/api.Bank/GetCredit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankClient) Transfer(ctx context.Context, in *TransactionReq, opts ...grpc.CallOption) (*Transaction, error) {
	out := new(Transaction)
	err := c.cc.Invoke(ctx, "/api.Bank/Transfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BankServer is the server API for Bank service.
type BankServer interface {
	NewAccount(context.Context, *AccountReq) (*Account, error)
	GetCredit(context.Context, *ID) (*Credit, error)
	Transfer(context.Context, *TransactionReq) (*Transaction, error)
}

func RegisterBankServer(s *grpc.Server, srv BankServer) {
	s.RegisterService(&_Bank_serviceDesc, srv)
}

func _Bank_NewAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServer).NewAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Bank/NewAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServer).NewAccount(ctx, req.(*AccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bank_GetCredit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServer).GetCredit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Bank/GetCredit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServer).GetCredit(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bank_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Bank/Transfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServer).Transfer(ctx, req.(*TransactionReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Bank_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Bank",
	HandlerType: (*BankServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewAccount",
			Handler:    _Bank_NewAccount_Handler,
		},
		{
			MethodName: "GetCredit",
			Handler:    _Bank_GetCredit_Handler,
		},
		{
			MethodName: "Transfer",
			Handler:    _Bank_Transfer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_api_0bf5ef62db4ac6b4) }

var fileDescriptor_api_0bf5ef62db4ac6b4 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x41, 0x4b, 0xfc, 0x30,
	0x10, 0xc5, 0x69, 0x5a, 0xba, 0xed, 0xec, 0x9f, 0xfd, 0x4b, 0x14, 0x91, 0x9c, 0xd6, 0x9c, 0x04,
	0x61, 0x41, 0xbd, 0x7a, 0x51, 0x0b, 0xd2, 0xcb, 0x1e, 0x8a, 0x47, 0x2f, 0x31, 0x8d, 0x10, 0x74,
	0xd3, 0x9a, 0x46, 0xfd, 0x0c, 0x7e, 0x6b, 0xc9, 0x10, 0x6b, 0xba, 0xd2, 0xdb, 0xbc, 0x37, 0x24,
	0xbf, 0x37, 0x99, 0x40, 0x29, 0x7a, 0xbd, 0xe9, 0x6d, 0xe7, 0x3a, 0x9a, 0x8a, 0x5e, 0xf3, 0x6b,
	0x80, 0x1b, 0x29, 0xbb, 0x77, 0xe3, 0x1a, 0xf5, 0x46, 0x29, 0x64, 0x46, 0xec, 0xd4, 0x49, 0xb2,
	0x4e, 0xce, 0xca, 0x06, 0x6b, 0xca, 0xa0, 0x78, 0x15, 0x83, 0xdb, 0x7a, 0x9f, 0xa0, 0x3f, 0x6a,
	0x7e, 0x04, 0xa4, 0xae, 0xe8, 0x0a, 0x88, 0x6e, 0xc3, 0x19, 0xa2, 0x5b, 0xbe, 0x86, 0xfc, 0xce,
	0xaa, 0x56, 0x3b, 0x7a, 0x0c, 0xb9, 0xc4, 0x0a, 0xbb, 0x69, 0x13, 0x14, 0x17, 0xb0, 0x08, 0xd4,
	0xfd, 0xc3, 0x63, 0x04, 0x32, 0x13, 0x21, 0x9d, 0x46, 0x88, 0x10, 0xd9, 0x04, 0xf1, 0x08, 0xab,
	0x07, 0x2b, 0xcc, 0x20, 0xa4, 0xd3, 0x9d, 0xf1, 0xc3, 0x31, 0x28, 0x06, 0x65, 0x5a, 0x65, 0xeb,
	0x2a, 0xf0, 0x46, 0xed, 0x6f, 0xb1, 0x4a, 0x7e, 0xd4, 0x55, 0xe0, 0x06, 0xe5, 0x7d, 0xb1, 0xf3,
	0x39, 0x91, 0x9b, 0x36, 0x41, 0x71, 0x0d, 0xcb, 0xe8, 0xf6, 0x3f, 0x43, 0xc4, 0x28, 0x32, 0x8b,
	0x4a, 0x67, 0x50, 0x59, 0x8c, 0xba, 0xfc, 0x4a, 0x20, 0xbb, 0x15, 0xe6, 0x85, 0x9e, 0x03, 0x6c,
	0xd5, 0xe7, 0xcf, 0xbb, 0xfd, 0xdf, 0xf8, 0x4d, 0xfe, 0xee, 0x8e, 0xfd, 0x8b, 0x0d, 0x7a, 0x0a,
	0xe5, 0xbd, 0x72, 0x61, 0x0d, 0x0b, 0x6c, 0xd5, 0x15, 0x5b, 0x62, 0x11, 0xdc, 0x0b, 0x28, 0x70,
	0x86, 0x67, 0x65, 0xe9, 0x21, 0x36, 0xa6, 0x0f, 0xc6, 0x0e, 0xf6, 0xcd, 0xa7, 0x1c, 0x7f, 0xce,
	0xd5, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x68, 0xe9, 0xc9, 0x46, 0x02, 0x00, 0x00,
}