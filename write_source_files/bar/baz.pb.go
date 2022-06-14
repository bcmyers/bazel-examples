// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.14.0
// source: bar/baz.proto

package bar

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

type Baz struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Baz *string `protobuf:"bytes,1,req,name=baz" json:"baz,omitempty"`
	Id  *int32  `protobuf:"varint,2,req,name=id" json:"id,omitempty"`
}

func (x *Baz) Reset() {
	*x = Baz{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bar_baz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Baz) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Baz) ProtoMessage() {}

func (x *Baz) ProtoReflect() protoreflect.Message {
	mi := &file_bar_baz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Baz.ProtoReflect.Descriptor instead.
func (*Baz) Descriptor() ([]byte, []int) {
	return file_bar_baz_proto_rawDescGZIP(), []int{0}
}

func (x *Baz) GetBaz() string {
	if x != nil && x.Baz != nil {
		return *x.Baz
	}
	return ""
}

func (x *Baz) GetId() int32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

var File_bar_baz_proto protoreflect.FileDescriptor

var file_bar_baz_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x62, 0x61, 0x72, 0x2f, 0x62, 0x61, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x62, 0x61, 0x72, 0x22, 0x27, 0x0a, 0x03, 0x42, 0x61, 0x7a, 0x12, 0x10, 0x0a, 0x03, 0x62,
	0x61, 0x7a, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x03, 0x62, 0x61, 0x7a, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
}

var (
	file_bar_baz_proto_rawDescOnce sync.Once
	file_bar_baz_proto_rawDescData = file_bar_baz_proto_rawDesc
)

func file_bar_baz_proto_rawDescGZIP() []byte {
	file_bar_baz_proto_rawDescOnce.Do(func() {
		file_bar_baz_proto_rawDescData = protoimpl.X.CompressGZIP(file_bar_baz_proto_rawDescData)
	})
	return file_bar_baz_proto_rawDescData
}

var file_bar_baz_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_bar_baz_proto_goTypes = []interface{}{
	(*Baz)(nil), // 0: bar.Baz
}
var file_bar_baz_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bar_baz_proto_init() }
func file_bar_baz_proto_init() {
	if File_bar_baz_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bar_baz_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Baz); i {
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
			RawDescriptor: file_bar_baz_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bar_baz_proto_goTypes,
		DependencyIndexes: file_bar_baz_proto_depIdxs,
		MessageInfos:      file_bar_baz_proto_msgTypes,
	}.Build()
	File_bar_baz_proto = out.File
	file_bar_baz_proto_rawDesc = nil
	file_bar_baz_proto_goTypes = nil
	file_bar_baz_proto_depIdxs = nil
}
