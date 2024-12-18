// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: grpc/pb.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StatusValue int32

const (
	StatusValue_FAULT       StatusValue = 0
	StatusValue_ACCEPTED    StatusValue = 1
	StatusValue_UNDERBID    StatusValue = 2
	StatusValue_NOT_STARTED StatusValue = 3
	StatusValue_IN_PROGRESS StatusValue = 4
	StatusValue_SOLD        StatusValue = 5
	StatusValue_CLOSED      StatusValue = 6
	StatusValue_REJECTED    StatusValue = 7
)

// Enum value maps for StatusValue.
var (
	StatusValue_name = map[int32]string{
		0: "FAULT",
		1: "ACCEPTED",
		2: "UNDERBID",
		3: "NOT_STARTED",
		4: "IN_PROGRESS",
		5: "SOLD",
		6: "CLOSED",
		7: "REJECTED",
	}
	StatusValue_value = map[string]int32{
		"FAULT":       0,
		"ACCEPTED":    1,
		"UNDERBID":    2,
		"NOT_STARTED": 3,
		"IN_PROGRESS": 4,
		"SOLD":        5,
		"CLOSED":      6,
		"REJECTED":    7,
	}
)

func (x StatusValue) Enum() *StatusValue {
	p := new(StatusValue)
	*p = x
	return p
}

func (x StatusValue) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusValue) Descriptor() protoreflect.EnumDescriptor {
	return file_grpc_pb_proto_enumTypes[0].Descriptor()
}

func (StatusValue) Type() protoreflect.EnumType {
	return &file_grpc_pb_proto_enumTypes[0]
}

func (x StatusValue) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusValue.Descriptor instead.
func (StatusValue) EnumDescriptor() ([]byte, []int) {
	return file_grpc_pb_proto_rawDescGZIP(), []int{0}
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_pb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_pb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_grpc_pb_proto_rawDescGZIP(), []int{0}
}

func (x *Node) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

type NodeState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopBid    *Bid       `protobuf:"bytes,1,opt,name=top_bid,json=topBid,proto3" json:"top_bid,omitempty"`
	BidTime   int64      `protobuf:"varint,2,opt,name=bid_time,json=bidTime,proto3" json:"bid_time,omitempty"`
	Bidders   int64      `protobuf:"varint,3,opt,name=bidders,proto3" json:"bidders,omitempty"`
	Addr      string     `protobuf:"bytes,4,opt,name=addr,proto3" json:"addr,omitempty"`
	Discovery *Discovery `protobuf:"bytes,5,opt,name=discovery,proto3,oneof" json:"discovery,omitempty"`
}

func (x *NodeState) Reset() {
	*x = NodeState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_pb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeState) ProtoMessage() {}

func (x *NodeState) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_pb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeState.ProtoReflect.Descriptor instead.
func (*NodeState) Descriptor() ([]byte, []int) {
	return file_grpc_pb_proto_rawDescGZIP(), []int{1}
}

func (x *NodeState) GetTopBid() *Bid {
	if x != nil {
		return x.TopBid
	}
	return nil
}

func (x *NodeState) GetBidTime() int64 {
	if x != nil {
		return x.BidTime
	}
	return 0
}

func (x *NodeState) GetBidders() int64 {
	if x != nil {
		return x.Bidders
	}
	return 0
}

func (x *NodeState) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *NodeState) GetDiscovery() *Discovery {
	if x != nil {
		return x.Discovery
	}
	return nil
}

type Lot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AskingPrice int64                  `protobuf:"varint,3,opt,name=asking_price,json=askingPrice,proto3" json:"asking_price,omitempty"`
	StartingBid int64                  `protobuf:"varint,4,opt,name=starting_bid,json=startingBid,proto3" json:"starting_bid,omitempty"`
	CurrentBid  int64                  `protobuf:"varint,5,opt,name=current_bid,json=currentBid,proto3" json:"current_bid,omitempty"`
	ClosingTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=closing_time,json=closingTime,proto3" json:"closing_time,omitempty"`
}

func (x *Lot) Reset() {
	*x = Lot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_pb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lot) ProtoMessage() {}

func (x *Lot) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_pb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lot.ProtoReflect.Descriptor instead.
func (*Lot) Descriptor() ([]byte, []int) {
	return file_grpc_pb_proto_rawDescGZIP(), []int{2}
}

func (x *Lot) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Lot) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Lot) GetAskingPrice() int64 {
	if x != nil {
		return x.AskingPrice
	}
	return 0
}

func (x *Lot) GetStartingBid() int64 {
	if x != nil {
		return x.StartingBid
	}
	return 0
}

func (x *Lot) GetCurrentBid() int64 {
	if x != nil {
		return x.CurrentBid
	}
	return 0
}

func (x *Lot) GetClosingTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ClosingTime
	}
	return nil
}

type Bid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount   int64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	BidderId int64 `protobuf:"varint,2,opt,name=bidder_id,json=bidderId,proto3" json:"bidder_id,omitempty"`
}

func (x *Bid) Reset() {
	*x = Bid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_pb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bid) ProtoMessage() {}

func (x *Bid) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_pb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bid.ProtoReflect.Descriptor instead.
func (*Bid) Descriptor() ([]byte, []int) {
	return file_grpc_pb_proto_rawDescGZIP(), []int{3}
}

func (x *Bid) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Bid) GetBidderId() int64 {
	if x != nil {
		return x.BidderId
	}
	return 0
}

type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount   int64       `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Result   StatusValue `protobuf:"varint,2,opt,name=result,proto3,enum=StatusValue" json:"result,omitempty"`
	BidderId int64       `protobuf:"varint,3,opt,name=bidder_id,json=bidderId,proto3" json:"bidder_id,omitempty"`
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_pb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_pb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_grpc_pb_proto_rawDescGZIP(), []int{4}
}

func (x *Ack) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Ack) GetResult() StatusValue {
	if x != nil {
		return x.Result
	}
	return StatusValue_FAULT
}

func (x *Ack) GetBidderId() int64 {
	if x != nil {
		return x.BidderId
	}
	return 0
}

type Discovery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Others []string `protobuf:"bytes,1,rep,name=others,proto3" json:"others,omitempty"`
	Leader string   `protobuf:"bytes,2,opt,name=leader,proto3" json:"leader,omitempty"`
}

func (x *Discovery) Reset() {
	*x = Discovery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_pb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Discovery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Discovery) ProtoMessage() {}

func (x *Discovery) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_pb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Discovery.ProtoReflect.Descriptor instead.
func (*Discovery) Descriptor() ([]byte, []int) {
	return file_grpc_pb_proto_rawDescGZIP(), []int{5}
}

func (x *Discovery) GetOthers() []string {
	if x != nil {
		return x.Others
	}
	return nil
}

func (x *Discovery) GetLeader() string {
	if x != nil {
		return x.Leader
	}
	return ""
}

type ElectionBallot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BidTime int64  `protobuf:"varint,1,opt,name=bid_time,json=bidTime,proto3" json:"bid_time,omitempty"`
	Addr    string `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
}

func (x *ElectionBallot) Reset() {
	*x = ElectionBallot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_pb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ElectionBallot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ElectionBallot) ProtoMessage() {}

func (x *ElectionBallot) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_pb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ElectionBallot.ProtoReflect.Descriptor instead.
func (*ElectionBallot) Descriptor() ([]byte, []int) {
	return file_grpc_pb_proto_rawDescGZIP(), []int{6}
}

func (x *ElectionBallot) GetBidTime() int64 {
	if x != nil {
		return x.BidTime
	}
	return 0
}

func (x *ElectionBallot) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

type ElectionAnswer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result StatusValue `protobuf:"varint,1,opt,name=result,proto3,enum=StatusValue" json:"result,omitempty"`
}

func (x *ElectionAnswer) Reset() {
	*x = ElectionAnswer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_pb_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ElectionAnswer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ElectionAnswer) ProtoMessage() {}

func (x *ElectionAnswer) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_pb_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ElectionAnswer.ProtoReflect.Descriptor instead.
func (*ElectionAnswer) Descriptor() ([]byte, []int) {
	return file_grpc_pb_proto_rawDescGZIP(), []int{7}
}

func (x *ElectionAnswer) GetResult() StatusValue {
	if x != nil {
		return x.Result
	}
	return StatusValue_FAULT
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_pb_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_pb_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_grpc_pb_proto_rawDescGZIP(), []int{8}
}

var File_grpc_pb_proto protoreflect.FileDescriptor

var file_grpc_pb_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x1a, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x22, 0xb0, 0x01, 0x0a,
	0x09, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x74, 0x6f,
	0x70, 0x5f, 0x62, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x04, 0x2e, 0x42, 0x69,
	0x64, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x42, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x69, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x69, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64,
	0x64, 0x72, 0x12, 0x2d, 0x0a, 0x09, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x48, 0x00, 0x52, 0x09, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x88, 0x01,
	0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x22,
	0xcf, 0x01, 0x0a, 0x03, 0x4c, 0x6f, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61,
	0x73, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x61, 0x73, 0x6b, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x69,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42,
	0x69, 0x64, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x6c, 0x6f, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x6c, 0x6f, 0x73, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x3a, 0x0a, 0x03, 0x42, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x60, 0x0a,
	0x03, 0x41, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x3b, 0x0a, 0x09, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x74, 0x68, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x3f, 0x0a, 0x0e,
	0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x6c, 0x6c, 0x6f, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x62, 0x69, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x62, 0x69, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x22, 0x36, 0x0a,
	0x0e, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12,
	0x24, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0c, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x2a, 0x7a,
	0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x09, 0x0a,
	0x05, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x43, 0x43, 0x45,
	0x50, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x4e, 0x44, 0x45, 0x52, 0x42,
	0x49, 0x44, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x52,
	0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47,
	0x52, 0x45, 0x53, 0x53, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x4f, 0x4c, 0x44, 0x10, 0x05,
	0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08,
	0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x07, 0x32, 0xa6, 0x02, 0x0a, 0x07, 0x41,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x06, 0x50, 0x75, 0x74, 0x42, 0x69, 0x64,
	0x12, 0x04, 0x2e, 0x42, 0x69, 0x64, 0x1a, 0x04, 0x2e, 0x41, 0x63, 0x6b, 0x12, 0x20, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x04, 0x2e, 0x41, 0x63, 0x6b, 0x12, 0x16,
	0x0a, 0x06, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x74, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x04, 0x2e, 0x4c, 0x6f, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a,
	0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0a, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x17, 0x0a, 0x08,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x05, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x1a,
	0x04, 0x2e, 0x4c, 0x6f, 0x74, 0x12, 0x16, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x06, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2c, 0x0a,
	0x08, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0f, 0x2e, 0x45, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x6c, 0x6c, 0x6f, 0x74, 0x1a, 0x0f, 0x2e, 0x45, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0b, 0x43,
	0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x0f, 0x2e, 0x45, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x6c, 0x6c, 0x6f, 0x74, 0x1a, 0x06, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x42, 0x20, 0x5a, 0x1e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_pb_proto_rawDescOnce sync.Once
	file_grpc_pb_proto_rawDescData = file_grpc_pb_proto_rawDesc
)

func file_grpc_pb_proto_rawDescGZIP() []byte {
	file_grpc_pb_proto_rawDescOnce.Do(func() {
		file_grpc_pb_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_pb_proto_rawDescData)
	})
	return file_grpc_pb_proto_rawDescData
}

var file_grpc_pb_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_grpc_pb_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_grpc_pb_proto_goTypes = []any{
	(StatusValue)(0),              // 0: StatusValue
	(*Node)(nil),                  // 1: Node
	(*NodeState)(nil),             // 2: NodeState
	(*Lot)(nil),                   // 3: Lot
	(*Bid)(nil),                   // 4: Bid
	(*Ack)(nil),                   // 5: Ack
	(*Discovery)(nil),             // 6: Discovery
	(*ElectionBallot)(nil),        // 7: ElectionBallot
	(*ElectionAnswer)(nil),        // 8: ElectionAnswer
	(*Empty)(nil),                 // 9: Empty
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_grpc_pb_proto_depIdxs = []int32{
	4,  // 0: NodeState.top_bid:type_name -> Bid
	6,  // 1: NodeState.discovery:type_name -> Discovery
	10, // 2: Lot.closing_time:type_name -> google.protobuf.Timestamp
	0,  // 3: Ack.result:type_name -> StatusValue
	0,  // 4: ElectionAnswer.result:type_name -> StatusValue
	4,  // 5: Auction.PutBid:input_type -> Bid
	9,  // 6: Auction.GetAuctionStatus:input_type -> Empty
	9,  // 7: Auction.GetLot:input_type -> Empty
	9,  // 8: Auction.GetDiscovery:input_type -> Empty
	2,  // 9: Auction.UpdateNode:input_type -> NodeState
	1,  // 10: Auction.Register:input_type -> Node
	9,  // 11: Auction.Ping:input_type -> Empty
	7,  // 12: Auction.Election:input_type -> ElectionBallot
	7,  // 13: Auction.Coordinator:input_type -> ElectionBallot
	5,  // 14: Auction.PutBid:output_type -> Ack
	5,  // 15: Auction.GetAuctionStatus:output_type -> Ack
	3,  // 16: Auction.GetLot:output_type -> Lot
	6,  // 17: Auction.GetDiscovery:output_type -> Discovery
	9,  // 18: Auction.UpdateNode:output_type -> Empty
	3,  // 19: Auction.Register:output_type -> Lot
	9,  // 20: Auction.Ping:output_type -> Empty
	8,  // 21: Auction.Election:output_type -> ElectionAnswer
	9,  // 22: Auction.Coordinator:output_type -> Empty
	14, // [14:23] is the sub-list for method output_type
	5,  // [5:14] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_grpc_pb_proto_init() }
func file_grpc_pb_proto_init() {
	if File_grpc_pb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_pb_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Node); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_pb_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*NodeState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_pb_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Lot); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_pb_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Bid); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_pb_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Ack); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_pb_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Discovery); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_pb_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ElectionBallot); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_pb_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*ElectionAnswer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_pb_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_grpc_pb_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_pb_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_pb_proto_goTypes,
		DependencyIndexes: file_grpc_pb_proto_depIdxs,
		EnumInfos:         file_grpc_pb_proto_enumTypes,
		MessageInfos:      file_grpc_pb_proto_msgTypes,
	}.Build()
	File_grpc_pb_proto = out.File
	file_grpc_pb_proto_rawDesc = nil
	file_grpc_pb_proto_goTypes = nil
	file_grpc_pb_proto_depIdxs = nil
}
