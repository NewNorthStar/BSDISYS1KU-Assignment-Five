// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: grpc/pb.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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
	StatusValue_SOLD        StatusValue = 3
	StatusValue_IN_PROGRESS StatusValue = 4
)

// Enum value maps for StatusValue.
var (
	StatusValue_name = map[int32]string{
		0: "FAULT",
		1: "ACCEPTED",
		2: "UNDERBID",
		3: "SOLD",
		4: "IN_PROGRESS",
	}
	StatusValue_value = map[string]int32{
		"FAULT":       0,
		"ACCEPTED":    1,
		"UNDERBID":    2,
		"SOLD":        3,
		"IN_PROGRESS": 4,
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

type Lot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AskingPrice int64  `protobuf:"varint,3,opt,name=asking_price,json=askingPrice,proto3" json:"asking_price,omitempty"`
	StartingBid int64  `protobuf:"varint,4,opt,name=starting_bid,json=startingBid,proto3" json:"starting_bid,omitempty"`
	CurrentBid  int64  `protobuf:"varint,5,opt,name=current_bid,json=currentBid,proto3" json:"current_bid,omitempty"`
}

func (x *Lot) Reset() {
	*x = Lot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_pb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lot) ProtoMessage() {}

func (x *Lot) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Lot.ProtoReflect.Descriptor instead.
func (*Lot) Descriptor() ([]byte, []int) {
	return file_grpc_pb_proto_rawDescGZIP(), []int{0}
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
		mi := &file_grpc_pb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bid) ProtoMessage() {}

func (x *Bid) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Bid.ProtoReflect.Descriptor instead.
func (*Bid) Descriptor() ([]byte, []int) {
	return file_grpc_pb_proto_rawDescGZIP(), []int{1}
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
		mi := &file_grpc_pb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_grpc_pb_proto_rawDescGZIP(), []int{2}
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

	IpAddresses []string `protobuf:"bytes,1,rep,name=ip_addresses,json=ipAddresses,proto3" json:"ip_addresses,omitempty"`
}

func (x *Discovery) Reset() {
	*x = Discovery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_pb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Discovery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Discovery) ProtoMessage() {}

func (x *Discovery) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Discovery.ProtoReflect.Descriptor instead.
func (*Discovery) Descriptor() ([]byte, []int) {
	return file_grpc_pb_proto_rawDescGZIP(), []int{3}
}

func (x *Discovery) GetIpAddresses() []string {
	if x != nil {
		return x.IpAddresses
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_pb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_grpc_pb_proto_rawDescGZIP(), []int{4}
}

var File_grpc_pb_proto protoreflect.FileDescriptor

var file_grpc_pb_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x90, 0x01, 0x0a, 0x03, 0x4c, 0x6f, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61,
	0x73, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x61, 0x73, 0x6b, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x69,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42,
	0x69, 0x64, 0x22, 0x3a, 0x0a, 0x03, 0x42, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x60,
	0x0a, 0x03, 0x41, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x2e, 0x0a, 0x09, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x12, 0x21, 0x0a,
	0x0c, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x2a, 0x4f, 0x0a, 0x0b, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x41, 0x55, 0x4c,
	0x54, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x4e, 0x44, 0x45, 0x52, 0x42, 0x49, 0x44, 0x10, 0x02, 0x12,
	0x08, 0x0a, 0x04, 0x53, 0x4f, 0x4c, 0x44, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x5f,
	0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x04, 0x32, 0x85, 0x01, 0x0a, 0x0f, 0x41,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x45, 0x6e, 0x64, 0x12, 0x14,
	0x0a, 0x06, 0x50, 0x75, 0x74, 0x42, 0x69, 0x64, 0x12, 0x04, 0x2e, 0x42, 0x69, 0x64, 0x1a, 0x04,
	0x2e, 0x41, 0x63, 0x6b, 0x12, 0x20, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x04, 0x2e, 0x41, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x74,
	0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x04, 0x2e, 0x4c, 0x6f, 0x74, 0x12, 0x22,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x12, 0x06,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x32, 0x59, 0x0a, 0x07, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a,
	0x06, 0x50, 0x75, 0x74, 0x42, 0x69, 0x64, 0x12, 0x04, 0x2e, 0x42, 0x69, 0x64, 0x1a, 0x04, 0x2e,
	0x41, 0x63, 0x6b, 0x12, 0x20, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x04, 0x2e, 0x41, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x74, 0x12,
	0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x04, 0x2e, 0x4c, 0x6f, 0x74, 0x42, 0x20, 0x5a,
	0x1e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
var file_grpc_pb_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_grpc_pb_proto_goTypes = []any{
	(StatusValue)(0),  // 0: StatusValue
	(*Lot)(nil),       // 1: Lot
	(*Bid)(nil),       // 2: Bid
	(*Ack)(nil),       // 3: Ack
	(*Discovery)(nil), // 4: Discovery
	(*Empty)(nil),     // 5: Empty
}
var file_grpc_pb_proto_depIdxs = []int32{
	0, // 0: Ack.result:type_name -> StatusValue
	2, // 1: AuctionFrontEnd.PutBid:input_type -> Bid
	5, // 2: AuctionFrontEnd.GetAuctionStatus:input_type -> Empty
	5, // 3: AuctionFrontEnd.GetLot:input_type -> Empty
	5, // 4: AuctionFrontEnd.GetDiscovery:input_type -> Empty
	2, // 5: Auction.PutBid:input_type -> Bid
	5, // 6: Auction.GetAuctionStatus:input_type -> Empty
	5, // 7: Auction.GetLot:input_type -> Empty
	3, // 8: AuctionFrontEnd.PutBid:output_type -> Ack
	3, // 9: AuctionFrontEnd.GetAuctionStatus:output_type -> Ack
	1, // 10: AuctionFrontEnd.GetLot:output_type -> Lot
	4, // 11: AuctionFrontEnd.GetDiscovery:output_type -> Discovery
	3, // 12: Auction.PutBid:output_type -> Ack
	3, // 13: Auction.GetAuctionStatus:output_type -> Ack
	1, // 14: Auction.GetLot:output_type -> Lot
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_grpc_pb_proto_init() }
func file_grpc_pb_proto_init() {
	if File_grpc_pb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_pb_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_grpc_pb_proto_msgTypes[1].Exporter = func(v any, i int) any {
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
		file_grpc_pb_proto_msgTypes[2].Exporter = func(v any, i int) any {
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
		file_grpc_pb_proto_msgTypes[3].Exporter = func(v any, i int) any {
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
		file_grpc_pb_proto_msgTypes[4].Exporter = func(v any, i int) any {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_pb_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   2,
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
