// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: p2p_network_service.proto

package rpc

import (
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role Role       `protobuf:"varint,1,opt,name=role,proto3,enum=darcher.Role" json:"role,omitempty"`
	Url  string     `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Head *ChainHead `protobuf:"bytes,3,opt,name=head,proto3" json:"head,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_network_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_network_service_proto_msgTypes[0]
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
	return file_p2p_network_service_proto_rawDescGZIP(), []int{0}
}

func (x *Node) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_DOER
}

func (x *Node) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Node) GetHead() *ChainHead {
	if x != nil {
		return x.Head
	}
	return nil
}

type AddPeerControlMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role   Role   `protobuf:"varint,1,opt,name=role,proto3,enum=darcher.Role" json:"role,omitempty"`
	Id     string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Url    string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	PeerId string `protobuf:"bytes,4,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Err    Error  `protobuf:"varint,5,opt,name=err,proto3,enum=darcher.Error" json:"err,omitempty"`
}

func (x *AddPeerControlMsg) Reset() {
	*x = AddPeerControlMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_network_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPeerControlMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPeerControlMsg) ProtoMessage() {}

func (x *AddPeerControlMsg) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_network_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPeerControlMsg.ProtoReflect.Descriptor instead.
func (*AddPeerControlMsg) Descriptor() ([]byte, []int) {
	return file_p2p_network_service_proto_rawDescGZIP(), []int{1}
}

func (x *AddPeerControlMsg) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_DOER
}

func (x *AddPeerControlMsg) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AddPeerControlMsg) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *AddPeerControlMsg) GetPeerId() string {
	if x != nil {
		return x.PeerId
	}
	return ""
}

func (x *AddPeerControlMsg) GetErr() Error {
	if x != nil {
		return x.Err
	}
	return Error_NilErr
}

type RemovePeerControlMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role   Role   `protobuf:"varint,1,opt,name=role,proto3,enum=darcher.Role" json:"role,omitempty"`
	Id     string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Url    string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	PeerId string `protobuf:"bytes,4,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Err    Error  `protobuf:"varint,5,opt,name=err,proto3,enum=darcher.Error" json:"err,omitempty"`
}

func (x *RemovePeerControlMsg) Reset() {
	*x = RemovePeerControlMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_network_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePeerControlMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePeerControlMsg) ProtoMessage() {}

func (x *RemovePeerControlMsg) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_network_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePeerControlMsg.ProtoReflect.Descriptor instead.
func (*RemovePeerControlMsg) Descriptor() ([]byte, []int) {
	return file_p2p_network_service_proto_rawDescGZIP(), []int{2}
}

func (x *RemovePeerControlMsg) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_DOER
}

func (x *RemovePeerControlMsg) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RemovePeerControlMsg) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *RemovePeerControlMsg) GetPeerId() string {
	if x != nil {
		return x.PeerId
	}
	return ""
}

func (x *RemovePeerControlMsg) GetErr() Error {
	if x != nil {
		return x.Err
	}
	return Error_NilErr
}

var File_p2p_network_service_proto protoreflect.FileDescriptor

var file_p2p_network_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x32, 0x70, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x64, 0x61, 0x72,
	0x63, 0x68, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x63, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x64, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x26, 0x0a,
	0x04, 0x68, 0x65, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x61,
	0x72, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x52,
	0x04, 0x68, 0x65, 0x61, 0x64, 0x22, 0x93, 0x01, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x50, 0x65, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x4d, 0x73, 0x67, 0x12, 0x21, 0x0a, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x64, 0x61, 0x72, 0x63,
	0x68, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x12, 0x17, 0x0a, 0x07, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x03, 0x65, 0x72, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x64, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x96, 0x01, 0x0a, 0x14,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x4d, 0x73, 0x67, 0x12, 0x21, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x64, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6c,
	0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x65, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0e, 0x2e, 0x64, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x03, 0x65, 0x72, 0x72, 0x32, 0xf8, 0x01, 0x0a, 0x11, 0x50, 0x32, 0x50, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x0f, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x0d, 0x2e,
	0x64, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0e, 0x61, 0x64, 0x64, 0x50, 0x65, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x1a, 0x2e, 0x64, 0x61, 0x72, 0x63, 0x68,
	0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x65, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x4d, 0x73, 0x67, 0x1a, 0x1a, 0x2e, 0x64, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x41,
	0x64, 0x64, 0x50, 0x65, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x4d, 0x73, 0x67,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x57, 0x0a, 0x11, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x50, 0x65, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x1d, 0x2e, 0x64, 0x61,
	0x72, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x4d, 0x73, 0x67, 0x1a, 0x1d, 0x2e, 0x64, 0x61, 0x72,
	0x63, 0x68, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x4d, 0x73, 0x67, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42,
	0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x72,
	0x6f, 0x75, 0x62, 0x6c, 0x6f, 0x72, 0x2f, 0x64, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x2d, 0x67,
	0x6f, 0x2d, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2f, 0x65, 0x74, 0x68, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_p2p_network_service_proto_rawDescOnce sync.Once
	file_p2p_network_service_proto_rawDescData = file_p2p_network_service_proto_rawDesc
)

func file_p2p_network_service_proto_rawDescGZIP() []byte {
	file_p2p_network_service_proto_rawDescOnce.Do(func() {
		file_p2p_network_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_p2p_network_service_proto_rawDescData)
	})
	return file_p2p_network_service_proto_rawDescData
}

var file_p2p_network_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_p2p_network_service_proto_goTypes = []interface{}{
	(*Node)(nil),                 // 0: darcher.Node
	(*AddPeerControlMsg)(nil),    // 1: darcher.AddPeerControlMsg
	(*RemovePeerControlMsg)(nil), // 2: darcher.RemovePeerControlMsg
	(Role)(0),                    // 3: darcher.Role
	(*ChainHead)(nil),            // 4: darcher.ChainHead
	(Error)(0),                   // 5: darcher.Error
	(*empty.Empty)(nil),          // 6: google.protobuf.Empty
}
var file_p2p_network_service_proto_depIdxs = []int32{
	3, // 0: darcher.Node.role:type_name -> darcher.Role
	4, // 1: darcher.Node.head:type_name -> darcher.ChainHead
	3, // 2: darcher.AddPeerControlMsg.role:type_name -> darcher.Role
	5, // 3: darcher.AddPeerControlMsg.err:type_name -> darcher.Error
	3, // 4: darcher.RemovePeerControlMsg.role:type_name -> darcher.Role
	5, // 5: darcher.RemovePeerControlMsg.err:type_name -> darcher.Error
	0, // 6: darcher.P2PNetworkService.notifyNodeStart:input_type -> darcher.Node
	1, // 7: darcher.P2PNetworkService.addPeerControl:input_type -> darcher.AddPeerControlMsg
	2, // 8: darcher.P2PNetworkService.removePeerControl:input_type -> darcher.RemovePeerControlMsg
	6, // 9: darcher.P2PNetworkService.notifyNodeStart:output_type -> google.protobuf.Empty
	1, // 10: darcher.P2PNetworkService.addPeerControl:output_type -> darcher.AddPeerControlMsg
	2, // 11: darcher.P2PNetworkService.removePeerControl:output_type -> darcher.RemovePeerControlMsg
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_p2p_network_service_proto_init() }
func file_p2p_network_service_proto_init() {
	if File_p2p_network_service_proto != nil {
		return
	}
	file_common_proto_init()
	file_blockchain_status_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_p2p_network_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_p2p_network_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPeerControlMsg); i {
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
		file_p2p_network_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePeerControlMsg); i {
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
			RawDescriptor: file_p2p_network_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_p2p_network_service_proto_goTypes,
		DependencyIndexes: file_p2p_network_service_proto_depIdxs,
		MessageInfos:      file_p2p_network_service_proto_msgTypes,
	}.Build()
	File_p2p_network_service_proto = out.File
	file_p2p_network_service_proto_rawDesc = nil
	file_p2p_network_service_proto_goTypes = nil
	file_p2p_network_service_proto_depIdxs = nil
}
