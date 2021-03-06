// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages/messages.proto

package messages

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RequestStatus int32

const (
	RequestStatus_REQUEST_SUCCESSFUL RequestStatus = 0
	RequestStatus_REQUEST_FAILED     RequestStatus = 1
)

var RequestStatus_name = map[int32]string{
	0: "REQUEST_SUCCESSFUL",
	1: "REQUEST_FAILED",
}

var RequestStatus_value = map[string]int32{
	"REQUEST_SUCCESSFUL": 0,
	"REQUEST_FAILED":     1,
}

func (x RequestStatus) String() string {
	return proto.EnumName(RequestStatus_name, int32(x))
}

func (RequestStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_83994550f81e9f35, []int{0}
}

type LaunchLogStatus int32

const (
	LaunchLogStatus_CREATED              LaunchLogStatus = 0
	LaunchLogStatus_RETRIED              LaunchLogStatus = 1
	LaunchLogStatus_PENDING              LaunchLogStatus = 2
	LaunchLogStatus_SUCCESS              LaunchLogStatus = 3
	LaunchLogStatus_FAILED               LaunchLogStatus = 4
	LaunchLogStatus_SIGN_FAILED          LaunchLogStatus = 5
	LaunchLogStatus_SEND_FAILED          LaunchLogStatus = 6
	LaunchLogStatus_ESTIMATED_GAS_FAILED LaunchLogStatus = 7
)

var LaunchLogStatus_name = map[int32]string{
	0: "CREATED",
	1: "RETRIED",
	2: "PENDING",
	3: "SUCCESS",
	4: "FAILED",
	5: "SIGN_FAILED",
	6: "SEND_FAILED",
	7: "ESTIMATED_GAS_FAILED",
}

var LaunchLogStatus_value = map[string]int32{
	"CREATED":              0,
	"RETRIED":              1,
	"PENDING":              2,
	"SUCCESS":              3,
	"FAILED":               4,
	"SIGN_FAILED":          5,
	"SEND_FAILED":          6,
	"ESTIMATED_GAS_FAILED": 7,
}

func (x LaunchLogStatus) String() string {
	return proto.EnumName(LaunchLogStatus_name, int32(x))
}

func (LaunchLogStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_83994550f81e9f35, []int{1}
}

// The request message containing the user's name.
type CreateMessage struct {
	From                 string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Data                 []byte   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	GasPrice             string   `protobuf:"bytes,5,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	GasLimit             uint32   `protobuf:"varint,6,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	ItemType             string   `protobuf:"bytes,7,opt,name=item_type,json=itemType,proto3" json:"item_type,omitempty"`
		ItemId               string   `protobuf:"bytes,8,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	IsUrgent             bool     `protobuf:"varint,9,opt,name=isUrgent,proto3" json:"isUrgent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMessage) Reset()         { *m = CreateMessage{} }
func (m *CreateMessage) String() string { return proto.CompactTextString(m) }
func (*CreateMessage) ProtoMessage()    {}
func (*CreateMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_83994550f81e9f35, []int{0}
}

func (m *CreateMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMessage.Unmarshal(m, b)
}
func (m *CreateMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMessage.Marshal(b, m, deterministic)
}
func (m *CreateMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMessage.Merge(m, src)
}
func (m *CreateMessage) XXX_Size() int {
	return xxx_messageInfo_CreateMessage.Size(m)
}
func (m *CreateMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMessage.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMessage proto.InternalMessageInfo

func (m *CreateMessage) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *CreateMessage) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *CreateMessage) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *CreateMessage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *CreateMessage) GetGasPrice() string {
	if m != nil {
		return m.GasPrice
	}
	return ""
}

func (m *CreateMessage) GetGasLimit() uint32 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *CreateMessage) GetItemType() string {
	if m != nil {
		return m.ItemType
	}
	return ""
}

func (m *CreateMessage) GetItemId() string {
	if m != nil {
		return m.ItemId
	}
	return ""
}

func (m *CreateMessage) GetIsUrgent() bool {
	if m != nil {
		return m.IsUrgent
	}
	return false
}

// The response message containing the greetings
type CreateReply struct {
	Status               RequestStatus `protobuf:"varint,1,opt,name=status,proto3,enum=messages.RequestStatus" json:"status,omitempty"`
	ErrMsg               string        `protobuf:"bytes,2,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
	Data                 *Log          `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateReply) Reset()         { *m = CreateReply{} }
func (m *CreateReply) String() string { return proto.CompactTextString(m) }
func (*CreateReply) ProtoMessage()    {}
func (*CreateReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_83994550f81e9f35, []int{1}
}

func (m *CreateReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateReply.Unmarshal(m, b)
}
func (m *CreateReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateReply.Marshal(b, m, deterministic)
}
func (m *CreateReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateReply.Merge(m, src)
}
func (m *CreateReply) XXX_Size() int {
	return xxx_messageInfo_CreateReply.Size(m)
}
func (m *CreateReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateReply.DiscardUnknown(m)
}

var xxx_messageInfo_CreateReply proto.InternalMessageInfo

func (m *CreateReply) GetStatus() RequestStatus {
	if m != nil {
		return m.Status
	}
	return RequestStatus_REQUEST_SUCCESSFUL
}

func (m *CreateReply) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *CreateReply) GetData() *Log {
	if m != nil {
		return m.Data
	}
	return nil
}

type HelloMessage struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloMessage) Reset()         { *m = HelloMessage{} }
func (m *HelloMessage) String() string { return proto.CompactTextString(m) }
func (*HelloMessage) ProtoMessage()    {}
func (*HelloMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_83994550f81e9f35, []int{2}
}

func (m *HelloMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloMessage.Unmarshal(m, b)
}
func (m *HelloMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloMessage.Marshal(b, m, deterministic)
}
func (m *HelloMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloMessage.Merge(m, src)
}
func (m *HelloMessage) XXX_Size() int {
	return xxx_messageInfo_HelloMessage.Size(m)
}
func (m *HelloMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloMessage.DiscardUnknown(m)
}

var xxx_messageInfo_HelloMessage proto.InternalMessageInfo

type HelloReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_83994550f81e9f35, []int{3}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

type NotifyMessage struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Status               bool     `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyMessage) Reset()         { *m = NotifyMessage{} }
func (m *NotifyMessage) String() string { return proto.CompactTextString(m) }
func (*NotifyMessage) ProtoMessage()    {}
func (*NotifyMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_83994550f81e9f35, []int{4}
}

func (m *NotifyMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyMessage.Unmarshal(m, b)
}
func (m *NotifyMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyMessage.Marshal(b, m, deterministic)
}
func (m *NotifyMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyMessage.Merge(m, src)
}
func (m *NotifyMessage) XXX_Size() int {
	return xxx_messageInfo_NotifyMessage.Size(m)
}
func (m *NotifyMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyMessage.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyMessage proto.InternalMessageInfo

func (m *NotifyMessage) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *NotifyMessage) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

type NotifyReply struct {
	Status               RequestStatus `protobuf:"varint,1,opt,name=status,proto3,enum=messages.RequestStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *NotifyReply) Reset()         { *m = NotifyReply{} }
func (m *NotifyReply) String() string { return proto.CompactTextString(m) }
func (*NotifyReply) ProtoMessage()    {}
func (*NotifyReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_83994550f81e9f35, []int{5}
}

func (m *NotifyReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyReply.Unmarshal(m, b)
}
func (m *NotifyReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyReply.Marshal(b, m, deterministic)
}
func (m *NotifyReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyReply.Merge(m, src)
}
func (m *NotifyReply) XXX_Size() int {
	return xxx_messageInfo_NotifyReply.Size(m)
}
func (m *NotifyReply) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyReply.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyReply proto.InternalMessageInfo

func (m *NotifyReply) GetStatus() RequestStatus {
	if m != nil {
		return m.Status
	}
	return RequestStatus_REQUEST_SUCCESSFUL
}

type Log struct {
	Hash                 string          `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	ItemType             string          `protobuf:"bytes,2,opt,name=item_type,json=itemType,proto3" json:"item_type,omitempty"`
	ItemId               string          `protobuf:"bytes,3,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Status               LaunchLogStatus `protobuf:"varint,4,opt,name=status,proto3,enum=messages.LaunchLogStatus" json:"status,omitempty"`
	GasPrice             string          `protobuf:"bytes,5,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	GasLimit             string          `protobuf:"bytes,6,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	GasUsed              uint64          `protobuf:"varint,7,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
	ExecutedAt           uint64          `protobuf:"varint,8,opt,name=executed_at,json=executedAt,proto3" json:"executed_at,omitempty"`
	Error                string          `protobuf:"bytes,9,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Log) Reset()         { *m = Log{} }
func (m *Log) String() string { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()    {}
func (*Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_83994550f81e9f35, []int{6}
}

func (m *Log) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Log.Unmarshal(m, b)
}
func (m *Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Log.Marshal(b, m, deterministic)
}
func (m *Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Log.Merge(m, src)
}
func (m *Log) XXX_Size() int {
	return xxx_messageInfo_Log.Size(m)
}
func (m *Log) XXX_DiscardUnknown() {
	xxx_messageInfo_Log.DiscardUnknown(m)
}

var xxx_messageInfo_Log proto.InternalMessageInfo

func (m *Log) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Log) GetItemType() string {
	if m != nil {
		return m.ItemType
	}
	return ""
}

func (m *Log) GetItemId() string {
	if m != nil {
		return m.ItemId
	}
	return ""
}

func (m *Log) GetStatus() LaunchLogStatus {
	if m != nil {
		return m.Status
	}
	return LaunchLogStatus_CREATED
}

func (m *Log) GetGasPrice() string {
	if m != nil {
		return m.GasPrice
	}
	return ""
}

func (m *Log) GetGasLimit() string {
	if m != nil {
		return m.GasLimit
	}
	return ""
}

func (m *Log) GetGasUsed() uint64 {
	if m != nil {
		return m.GasUsed
	}
	return 0
}

func (m *Log) GetExecutedAt() uint64 {
	if m != nil {
		return m.ExecutedAt
	}
	return 0
}

func (m *Log) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type GetMessage struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	ItemType             string   `protobuf:"bytes,2,opt,name=item_type,json=itemType,proto3" json:"item_type,omitempty"`
	ItemId               string   `protobuf:"bytes,3,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMessage) Reset()         { *m = GetMessage{} }
func (m *GetMessage) String() string { return proto.CompactTextString(m) }
func (*GetMessage) ProtoMessage()    {}
func (*GetMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_83994550f81e9f35, []int{7}
}

func (m *GetMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMessage.Unmarshal(m, b)
}
func (m *GetMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMessage.Marshal(b, m, deterministic)
}
func (m *GetMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMessage.Merge(m, src)
}
func (m *GetMessage) XXX_Size() int {
	return xxx_messageInfo_GetMessage.Size(m)
}
func (m *GetMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMessage.DiscardUnknown(m)
}

var xxx_messageInfo_GetMessage proto.InternalMessageInfo

func (m *GetMessage) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *GetMessage) GetItemType() string {
	if m != nil {
		return m.ItemType
	}
	return ""
}

func (m *GetMessage) GetItemId() string {
	if m != nil {
		return m.ItemId
	}
	return ""
}

type GetReply struct {
	Status               RequestStatus `protobuf:"varint,1,opt,name=status,proto3,enum=messages.RequestStatus" json:"status,omitempty"`
	Data                 []*Log        `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetReply) Reset()         { *m = GetReply{} }
func (m *GetReply) String() string { return proto.CompactTextString(m) }
func (*GetReply) ProtoMessage()    {}
func (*GetReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_83994550f81e9f35, []int{8}
}

func (m *GetReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReply.Unmarshal(m, b)
}
func (m *GetReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReply.Marshal(b, m, deterministic)
}
func (m *GetReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReply.Merge(m, src)
}
func (m *GetReply) XXX_Size() int {
	return xxx_messageInfo_GetReply.Size(m)
}
func (m *GetReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetReply proto.InternalMessageInfo

func (m *GetReply) GetStatus() RequestStatus {
	if m != nil {
		return m.Status
	}
	return RequestStatus_REQUEST_SUCCESSFUL
}

func (m *GetReply) GetData() []*Log {
	if m != nil {
		return m.Data
	}
	return nil
}

type SubscribeMessage struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	ItemType             string   `protobuf:"bytes,2,opt,name=item_type,json=itemType,proto3" json:"item_type,omitempty"`
	ItemId               string   `protobuf:"bytes,3,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeMessage) Reset()         { *m = SubscribeMessage{} }
func (m *SubscribeMessage) String() string { return proto.CompactTextString(m) }
func (*SubscribeMessage) ProtoMessage()    {}
func (*SubscribeMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_83994550f81e9f35, []int{9}
}

func (m *SubscribeMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeMessage.Unmarshal(m, b)
}
func (m *SubscribeMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeMessage.Marshal(b, m, deterministic)
}
func (m *SubscribeMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeMessage.Merge(m, src)
}
func (m *SubscribeMessage) XXX_Size() int {
	return xxx_messageInfo_SubscribeMessage.Size(m)
}
func (m *SubscribeMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeMessage proto.InternalMessageInfo

func (m *SubscribeMessage) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *SubscribeMessage) GetItemType() string {
	if m != nil {
		return m.ItemType
	}
	return ""
}

func (m *SubscribeMessage) GetItemId() string {
	if m != nil {
		return m.ItemId
	}
	return ""
}

type SubscribeReply struct {
	Status               LaunchLogStatus `protobuf:"varint,1,opt,name=status,proto3,enum=messages.LaunchLogStatus" json:"status,omitempty"`
	Hash                 string          `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	ItemType             string          `protobuf:"bytes,3,opt,name=item_type,json=itemType,proto3" json:"item_type,omitempty"`
	ItemId               string          `protobuf:"bytes,4,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	ErrMsg               string          `protobuf:"bytes,5,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SubscribeReply) Reset()         { *m = SubscribeReply{} }
func (m *SubscribeReply) String() string { return proto.CompactTextString(m) }
func (*SubscribeReply) ProtoMessage()    {}
func (*SubscribeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_83994550f81e9f35, []int{10}
}

func (m *SubscribeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeReply.Unmarshal(m, b)
}
func (m *SubscribeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeReply.Marshal(b, m, deterministic)
}
func (m *SubscribeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeReply.Merge(m, src)
}
func (m *SubscribeReply) XXX_Size() int {
	return xxx_messageInfo_SubscribeReply.Size(m)
}
func (m *SubscribeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeReply.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeReply proto.InternalMessageInfo

func (m *SubscribeReply) GetStatus() LaunchLogStatus {
	if m != nil {
		return m.Status
	}
	return LaunchLogStatus_CREATED
}

func (m *SubscribeReply) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *SubscribeReply) GetItemType() string {
	if m != nil {
		return m.ItemType
	}
	return ""
}

func (m *SubscribeReply) GetItemId() string {
	if m != nil {
		return m.ItemId
	}
	return ""
}

func (m *SubscribeReply) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func init() {
	proto.RegisterEnum("messages.RequestStatus", RequestStatus_name, RequestStatus_value)
	proto.RegisterEnum("messages.LaunchLogStatus", LaunchLogStatus_name, LaunchLogStatus_value)
	proto.RegisterType((*CreateMessage)(nil), "messages.CreateMessage")
	proto.RegisterType((*CreateReply)(nil), "messages.CreateReply")
	proto.RegisterType((*HelloMessage)(nil), "messages.HelloMessage")
	proto.RegisterType((*HelloReply)(nil), "messages.HelloReply")
	proto.RegisterType((*NotifyMessage)(nil), "messages.NotifyMessage")
	proto.RegisterType((*NotifyReply)(nil), "messages.NotifyReply")
	proto.RegisterType((*Log)(nil), "messages.Log")
	proto.RegisterType((*GetMessage)(nil), "messages.GetMessage")
	proto.RegisterType((*GetReply)(nil), "messages.GetReply")
	proto.RegisterType((*SubscribeMessage)(nil), "messages.SubscribeMessage")
	proto.RegisterType((*SubscribeReply)(nil), "messages.SubscribeReply")
}

func init() { proto.RegisterFile("messages/messages.proto", fileDescriptor_83994550f81e9f35) }

var fileDescriptor_83994550f81e9f35 = []byte{
	// 724 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xd1, 0x6e, 0xda, 0x4a,
	0x10, 0xc5, 0x06, 0x0c, 0x0c, 0x81, 0xa0, 0x11, 0x97, 0x38, 0xdc, 0x87, 0xcb, 0xf5, 0x13, 0xca,
	0x43, 0x72, 0x93, 0xfb, 0x50, 0x29, 0x91, 0x2a, 0x21, 0x70, 0x28, 0x12, 0x41, 0xa9, 0x0d, 0x7d,
	0xaa, 0x8a, 0x1c, 0xd8, 0x38, 0x96, 0x20, 0xa6, 0xbb, 0xeb, 0xaa, 0xf4, 0x1f, 0xda, 0xbf, 0xe8,
	0x37, 0xf4, 0x9b, 0xfa, 0x17, 0xd5, 0xae, 0x6d, 0xb0, 0x51, 0xa8, 0xda, 0xa8, 0x6f, 0x7b, 0x66,
	0x76, 0x67, 0xce, 0xcc, 0x1c, 0x8f, 0xe1, 0x68, 0x49, 0x18, 0x73, 0x5c, 0xc2, 0xce, 0xe2, 0xc3,
	0xe9, 0x8a, 0xfa, 0xdc, 0xc7, 0x62, 0x8c, 0x8d, 0xef, 0x0a, 0x54, 0xba, 0x94, 0x38, 0x9c, 0xdc,
	0x84, 0x26, 0x44, 0xc8, 0xdd, 0x53, 0x7f, 0xa9, 0x2b, 0x2d, 0xa5, 0x5d, 0xb2, 0xe4, 0x19, 0xab,
	0xa0, 0x72, 0x5f, 0x57, 0xa5, 0x45, 0xe5, 0x3e, 0xd6, 0x21, 0xff, 0xc1, 0x59, 0x04, 0x44, 0xcf,
	0x4a, 0x53, 0x08, 0xc4, 0xcb, 0xb9, 0xc3, 0x1d, 0x3d, 0xd7, 0x52, 0xda, 0x07, 0x96, 0x3c, 0xe3,
	0xdf, 0x50, 0x72, 0x1d, 0x36, 0x5d, 0x51, 0x6f, 0x46, 0xf4, 0xbc, 0xbc, 0x5d, 0x74, 0x1d, 0x76,
	0x2b, 0x70, 0xec, 0x5c, 0x78, 0x4b, 0x8f, 0xeb, 0x5a, 0x4b, 0x69, 0x57, 0xa4, 0x73, 0x28, 0xb0,
	0x70, 0x7a, 0x9c, 0x2c, 0xa7, 0x7c, 0xbd, 0x22, 0x7a, 0x21, 0x7c, 0x29, 0x0c, 0xe3, 0xf5, 0x8a,
	0xe0, 0x11, 0x14, 0xa4, 0xd3, 0x9b, 0xeb, 0x45, 0xe9, 0xd2, 0x04, 0x1c, 0xcc, 0xb1, 0x09, 0x45,
	0x8f, 0x4d, 0xa8, 0x4b, 0x1e, 0xb9, 0x5e, 0x6a, 0x29, 0xed, 0xa2, 0xb5, 0xc1, 0xc6, 0x27, 0x28,
	0x87, 0xa5, 0x5a, 0x64, 0xb5, 0x58, 0xe3, 0x19, 0x68, 0x8c, 0x3b, 0x3c, 0x60, 0xb2, 0xd4, 0xea,
	0xc5, 0xd1, 0xe9, 0xa6, 0x4b, 0x16, 0x79, 0x1f, 0x10, 0xc6, 0x6d, 0xe9, 0xb6, 0xa2, 0x6b, 0x22,
	0x29, 0xa1, 0x74, 0xba, 0x64, 0x6e, 0xd4, 0x0a, 0x8d, 0x50, 0x7a, 0xc3, 0x5c, 0xfc, 0x37, 0x2a,
	0x5c, 0x74, 0xa3, 0x7c, 0x51, 0xd9, 0xc6, 0x19, 0xfa, 0x6e, 0xd8, 0x07, 0xa3, 0x0a, 0x07, 0xaf,
	0xc8, 0x62, 0xe1, 0x47, 0x5d, 0x36, 0x0e, 0x00, 0x24, 0x96, 0x54, 0x8c, 0x2b, 0xa8, 0x8c, 0x7c,
	0xee, 0xdd, 0xaf, 0x13, 0x43, 0x78, 0x70, 0xd8, 0x43, 0x3c, 0x04, 0x71, 0xc6, 0xc6, 0x86, 0xaf,
	0x2a, 0x0b, 0x8b, 0x90, 0xf1, 0x12, 0xca, 0xe1, 0xe3, 0xe7, 0x95, 0x65, 0x7c, 0x51, 0x21, 0x3b,
	0xf4, 0xdd, 0x27, 0x73, 0xa6, 0x86, 0xa0, 0xee, 0x1f, 0x42, 0x36, 0x35, 0x84, 0xf3, 0x0d, 0x85,
	0x9c, 0xa4, 0x70, 0x9c, 0xe8, 0x88, 0x13, 0x3c, 0xce, 0x1e, 0x86, 0xbe, 0xbb, 0xd3, 0xdb, 0xdf,
	0xd3, 0x49, 0x29, 0xa1, 0x93, 0x63, 0x10, 0xe7, 0x69, 0xc0, 0xc8, 0x5c, 0xca, 0x24, 0x67, 0x15,
	0x5c, 0x87, 0x4d, 0x18, 0x99, 0xe3, 0x3f, 0x50, 0x26, 0x1f, 0xc9, 0x2c, 0xe0, 0x64, 0x3e, 0x75,
	0xb8, 0x54, 0x4a, 0xce, 0x82, 0xd8, 0xd4, 0xe1, 0x42, 0xc7, 0x84, 0x52, 0x9f, 0x4a, 0xa9, 0x94,
	0xac, 0x10, 0x18, 0x6f, 0x00, 0xfa, 0x84, 0xff, 0x6c, 0x14, 0xcf, 0x6a, 0x8b, 0xf1, 0x0e, 0x8a,
	0x7d, 0xc2, 0x9f, 0x29, 0xbe, 0x58, 0x63, 0x6a, 0x2b, 0xbb, 0x4f, 0x63, 0x6f, 0xa1, 0x66, 0x07,
	0x77, 0x6c, 0x46, 0xbd, 0x3b, 0xf2, 0xe7, 0xd9, 0x7f, 0x55, 0xa0, 0xba, 0x09, 0x1f, 0x16, 0x71,
	0xbe, 0x53, 0xc4, 0x2f, 0xcc, 0x39, 0xe6, 0xa3, 0xee, 0xe3, 0x93, 0xdd, 0xcf, 0x27, 0x97, 0x12,
	0x59, 0x03, 0xa2, 0xcf, 0x2f, 0x92, 0x4b, 0x84, 0x4e, 0xae, 0xa0, 0x92, 0xea, 0x20, 0x36, 0x00,
	0x2d, 0xf3, 0xf5, 0xc4, 0xb4, 0xc7, 0x53, 0x7b, 0xd2, 0xed, 0x9a, 0xb6, 0x7d, 0x3d, 0x19, 0xd6,
	0x32, 0x88, 0x50, 0x8d, 0xed, 0xd7, 0x9d, 0xc1, 0xd0, 0xec, 0xd5, 0x94, 0x93, 0xcf, 0x0a, 0x1c,
	0xee, 0x50, 0xc7, 0x32, 0x14, 0xba, 0x96, 0xd9, 0x19, 0x9b, 0xbd, 0x5a, 0x46, 0x00, 0xcb, 0x1c,
	0x5b, 0x03, 0x71, 0x5b, 0x80, 0x5b, 0x73, 0xd4, 0x1b, 0x8c, 0xfa, 0x35, 0x55, 0x80, 0x28, 0x7c,
	0x2d, 0x8b, 0x00, 0x5a, 0x14, 0x33, 0x87, 0x87, 0x50, 0xb6, 0x07, 0xfd, 0x51, 0x9c, 0x24, 0x2f,
	0x0d, 0xe6, 0xa8, 0x17, 0x1b, 0x34, 0xd4, 0xa1, 0x6e, 0xda, 0xe3, 0xc1, 0x8d, 0xc8, 0x31, 0xed,
	0x77, 0xec, 0xd8, 0x53, 0xb8, 0xf8, 0xa6, 0x42, 0x31, 0xe4, 0x43, 0x28, 0x5e, 0x82, 0x16, 0xee,
	0x2f, 0x4c, 0xa8, 0x25, 0xb5, 0xbc, 0x9b, 0x7f, 0xed, 0x3a, 0xc2, 0xfd, 0x92, 0xc1, 0x73, 0xc8,
	0xf6, 0x09, 0xc7, 0xfa, 0xd6, 0xbf, 0x95, 0x78, 0x13, 0x53, 0xd6, 0xf8, 0xc9, 0x25, 0x68, 0xe1,
	0x5e, 0x49, 0xa6, 0x4b, 0xad, 0xa9, 0x64, 0xba, 0xc4, 0x0a, 0x32, 0x32, 0xf8, 0x02, 0xf2, 0x72,
	0xbd, 0x61, 0x63, 0x7b, 0x23, 0xb9, 0xff, 0x9a, 0xf5, 0x1d, 0x7b, 0xfc, 0xb0, 0x0f, 0xa5, 0x8d,
	0xc8, 0xb0, 0xb9, 0xbd, 0xb4, 0x2b, 0xec, 0xa6, 0xfe, 0x84, 0x2f, 0x0a, 0xd2, 0x56, 0xfe, 0x53,
	0xee, 0x34, 0xf9, 0xa7, 0xfb, 0xff, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0x82, 0xd5, 0x9f, 0x63,
	0x04, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LauncherClient is the client API for Launcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LauncherClient interface {
	Create(ctx context.Context, in *CreateMessage, opts ...grpc.CallOption) (*CreateReply, error)
	Get(ctx context.Context, in *GetMessage, opts ...grpc.CallOption) (*GetReply, error)
	Notify(ctx context.Context, in *NotifyMessage, opts ...grpc.CallOption) (*NotifyReply, error)
	Hello(ctx context.Context, in *HelloMessage, opts ...grpc.CallOption) (*HelloReply, error)
	// will push every changed launch log info to the subscriber
	Subscribe(ctx context.Context, opts ...grpc.CallOption) (Launcher_SubscribeClient, error)
}

type launcherClient struct {
	cc *grpc.ClientConn
}

func NewLauncherClient(cc *grpc.ClientConn) LauncherClient {
	return &launcherClient{cc}
}

func (c *launcherClient) Create(ctx context.Context, in *CreateMessage, opts ...grpc.CallOption) (*CreateReply, error) {
	out := new(CreateReply)
	err := c.cc.Invoke(ctx, "/messages.Launcher/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *launcherClient) Get(ctx context.Context, in *GetMessage, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := c.cc.Invoke(ctx, "/messages.Launcher/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *launcherClient) Notify(ctx context.Context, in *NotifyMessage, opts ...grpc.CallOption) (*NotifyReply, error) {
	out := new(NotifyReply)
	err := c.cc.Invoke(ctx, "/messages.Launcher/Notify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *launcherClient) Hello(ctx context.Context, in *HelloMessage, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/messages.Launcher/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *launcherClient) Subscribe(ctx context.Context, opts ...grpc.CallOption) (Launcher_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Launcher_serviceDesc.Streams[0], "/messages.Launcher/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &launcherSubscribeClient{stream}
	return x, nil
}

type Launcher_SubscribeClient interface {
	Send(*SubscribeMessage) error
	Recv() (*SubscribeReply, error)
	grpc.ClientStream
}

type launcherSubscribeClient struct {
	grpc.ClientStream
}

func (x *launcherSubscribeClient) Send(m *SubscribeMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *launcherSubscribeClient) Recv() (*SubscribeReply, error) {
	m := new(SubscribeReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LauncherServer is the server API for Launcher service.
type LauncherServer interface {
	Create(context.Context, *CreateMessage) (*CreateReply, error)
	Get(context.Context, *GetMessage) (*GetReply, error)
	Notify(context.Context, *NotifyMessage) (*NotifyReply, error)
	Hello(context.Context, *HelloMessage) (*HelloReply, error)
	// will push every changed launch log info to the subscriber
	Subscribe(Launcher_SubscribeServer) error
}

// UnimplementedLauncherServer can be embedded to have forward compatible implementations.
type UnimplementedLauncherServer struct {
}

func (*UnimplementedLauncherServer) Create(ctx context.Context, req *CreateMessage) (*CreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedLauncherServer) Get(ctx context.Context, req *GetMessage) (*GetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedLauncherServer) Notify(ctx context.Context, req *NotifyMessage) (*NotifyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Notify not implemented")
}
func (*UnimplementedLauncherServer) Hello(ctx context.Context, req *HelloMessage) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (*UnimplementedLauncherServer) Subscribe(srv Launcher_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

func RegisterLauncherServer(s *grpc.Server, srv LauncherServer) {
	s.RegisterService(&_Launcher_serviceDesc, srv)
}

func _Launcher_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LauncherServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messages.Launcher/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LauncherServer).Create(ctx, req.(*CreateMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Launcher_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LauncherServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messages.Launcher/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LauncherServer).Get(ctx, req.(*GetMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Launcher_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LauncherServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messages.Launcher/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LauncherServer).Notify(ctx, req.(*NotifyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Launcher_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LauncherServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messages.Launcher/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LauncherServer).Hello(ctx, req.(*HelloMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Launcher_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LauncherServer).Subscribe(&launcherSubscribeServer{stream})
}

type Launcher_SubscribeServer interface {
	Send(*SubscribeReply) error
	Recv() (*SubscribeMessage, error)
	grpc.ServerStream
}

type launcherSubscribeServer struct {
	grpc.ServerStream
}

func (x *launcherSubscribeServer) Send(m *SubscribeReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *launcherSubscribeServer) Recv() (*SubscribeMessage, error) {
	m := new(SubscribeMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Launcher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "messages.Launcher",
	HandlerType: (*LauncherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Launcher_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Launcher_Get_Handler,
		},
		{
			MethodName: "Notify",
			Handler:    _Launcher_Notify_Handler,
		},
		{
			MethodName: "Hello",
			Handler:    _Launcher_Hello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _Launcher_Subscribe_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "messages/messages.proto",
}
