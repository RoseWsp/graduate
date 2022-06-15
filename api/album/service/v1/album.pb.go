// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: album.proto

package v1

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

type GetAlbumByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAlbumByIdReq) Reset() {
	*x = GetAlbumByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_album_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAlbumByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlbumByIdReq) ProtoMessage() {}

func (x *GetAlbumByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_album_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlbumByIdReq.ProtoReflect.Descriptor instead.
func (*GetAlbumByIdReq) Descriptor() ([]byte, []int) {
	return file_album_proto_rawDescGZIP(), []int{0}
}

func (x *GetAlbumByIdReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetAlbumByIdReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Album *GetAlbumByIdReply_Album `protobuf:"bytes,1,opt,name=album,proto3" json:"album,omitempty"`
}

func (x *GetAlbumByIdReply) Reset() {
	*x = GetAlbumByIdReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_album_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAlbumByIdReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlbumByIdReply) ProtoMessage() {}

func (x *GetAlbumByIdReply) ProtoReflect() protoreflect.Message {
	mi := &file_album_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlbumByIdReply.ProtoReflect.Descriptor instead.
func (*GetAlbumByIdReply) Descriptor() ([]byte, []int) {
	return file_album_proto_rawDescGZIP(), []int{1}
}

func (x *GetAlbumByIdReply) GetAlbum() *GetAlbumByIdReply_Album {
	if x != nil {
		return x.Album
	}
	return nil
}

type ListAlbumReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNum  int64 `protobuf:"varint,1,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListAlbumReq) Reset() {
	*x = ListAlbumReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_album_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAlbumReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAlbumReq) ProtoMessage() {}

func (x *ListAlbumReq) ProtoReflect() protoreflect.Message {
	mi := &file_album_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAlbumReq.ProtoReflect.Descriptor instead.
func (*ListAlbumReq) Descriptor() ([]byte, []int) {
	return file_album_proto_rawDescGZIP(), []int{2}
}

func (x *ListAlbumReq) GetPageNum() int64 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *ListAlbumReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListAlbumReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count  int64                   `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Albums []*ListAlbumReply_Album `protobuf:"bytes,2,rep,name=albums,proto3" json:"albums,omitempty"`
}

func (x *ListAlbumReply) Reset() {
	*x = ListAlbumReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_album_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAlbumReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAlbumReply) ProtoMessage() {}

func (x *ListAlbumReply) ProtoReflect() protoreflect.Message {
	mi := &file_album_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAlbumReply.ProtoReflect.Descriptor instead.
func (*ListAlbumReply) Descriptor() ([]byte, []int) {
	return file_album_proto_rawDescGZIP(), []int{3}
}

func (x *ListAlbumReply) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ListAlbumReply) GetAlbums() []*ListAlbumReply_Album {
	if x != nil {
		return x.Albums
	}
	return nil
}

type GetAlbumByIdReply_Album struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title    string  `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Artist   string  `protobuf:"bytes,3,opt,name=artist,proto3" json:"artist,omitempty"`
	Price    float64 `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	CreateAt string  `protobuf:"bytes,5,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
}

func (x *GetAlbumByIdReply_Album) Reset() {
	*x = GetAlbumByIdReply_Album{}
	if protoimpl.UnsafeEnabled {
		mi := &file_album_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAlbumByIdReply_Album) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlbumByIdReply_Album) ProtoMessage() {}

func (x *GetAlbumByIdReply_Album) ProtoReflect() protoreflect.Message {
	mi := &file_album_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlbumByIdReply_Album.ProtoReflect.Descriptor instead.
func (*GetAlbumByIdReply_Album) Descriptor() ([]byte, []int) {
	return file_album_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetAlbumByIdReply_Album) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetAlbumByIdReply_Album) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetAlbumByIdReply_Album) GetArtist() string {
	if x != nil {
		return x.Artist
	}
	return ""
}

func (x *GetAlbumByIdReply_Album) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *GetAlbumByIdReply_Album) GetCreateAt() string {
	if x != nil {
		return x.CreateAt
	}
	return ""
}

type ListAlbumReply_Album struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title    string  `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Artist   string  `protobuf:"bytes,3,opt,name=artist,proto3" json:"artist,omitempty"`
	Price    float64 `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	CreateAt string  `protobuf:"bytes,5,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
}

func (x *ListAlbumReply_Album) Reset() {
	*x = ListAlbumReply_Album{}
	if protoimpl.UnsafeEnabled {
		mi := &file_album_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAlbumReply_Album) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAlbumReply_Album) ProtoMessage() {}

func (x *ListAlbumReply_Album) ProtoReflect() protoreflect.Message {
	mi := &file_album_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAlbumReply_Album.ProtoReflect.Descriptor instead.
func (*ListAlbumReply_Album) Descriptor() ([]byte, []int) {
	return file_album_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ListAlbumReply_Album) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListAlbumReply_Album) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ListAlbumReply_Album) GetArtist() string {
	if x != nil {
		return x.Artist
	}
	return ""
}

func (x *ListAlbumReply_Album) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ListAlbumReply_Album) GetCreateAt() string {
	if x != nil {
		return x.CreateAt
	}
	return ""
}

var File_album_proto protoreflect.FileDescriptor

var file_album_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0xbd, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52,
	0x05, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x1a, 0x78, 0x0a, 0x05, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74,
	0x22, 0x46, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x71,
	0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xcf, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x2d, 0x0a, 0x06, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x06, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x73,
	0x1a, 0x78, 0x0a, 0x05, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x32, 0x6e, 0x0a, 0x05, 0x41, 0x6c,
	0x62, 0x75, 0x6d, 0x12, 0x2d, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d,
	0x12, 0x0d, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x1a,
	0x0f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x36, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x42, 0x79,
	0x49, 0x64, 0x12, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x72,
	0x61, 0x64, 0x75, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6c, 0x62, 0x75, 0x6d,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_album_proto_rawDescOnce sync.Once
	file_album_proto_rawDescData = file_album_proto_rawDesc
)

func file_album_proto_rawDescGZIP() []byte {
	file_album_proto_rawDescOnce.Do(func() {
		file_album_proto_rawDescData = protoimpl.X.CompressGZIP(file_album_proto_rawDescData)
	})
	return file_album_proto_rawDescData
}

var file_album_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_album_proto_goTypes = []interface{}{
	(*GetAlbumByIdReq)(nil),         // 0: GetAlbumByIdReq
	(*GetAlbumByIdReply)(nil),       // 1: GetAlbumByIdReply
	(*ListAlbumReq)(nil),            // 2: ListAlbumReq
	(*ListAlbumReply)(nil),          // 3: ListAlbumReply
	(*GetAlbumByIdReply_Album)(nil), // 4: GetAlbumByIdReply.Album
	(*ListAlbumReply_Album)(nil),    // 5: ListAlbumReply.Album
}
var file_album_proto_depIdxs = []int32{
	4, // 0: GetAlbumByIdReply.album:type_name -> GetAlbumByIdReply.Album
	5, // 1: ListAlbumReply.albums:type_name -> ListAlbumReply.Album
	2, // 2: Album.ListAlbum:input_type -> ListAlbumReq
	0, // 3: Album.GetAlbumById:input_type -> GetAlbumByIdReq
	3, // 4: Album.ListAlbum:output_type -> ListAlbumReply
	1, // 5: Album.GetAlbumById:output_type -> GetAlbumByIdReply
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_album_proto_init() }
func file_album_proto_init() {
	if File_album_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_album_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAlbumByIdReq); i {
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
		file_album_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAlbumByIdReply); i {
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
		file_album_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAlbumReq); i {
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
		file_album_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAlbumReply); i {
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
		file_album_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAlbumByIdReply_Album); i {
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
		file_album_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAlbumReply_Album); i {
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
			RawDescriptor: file_album_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_album_proto_goTypes,
		DependencyIndexes: file_album_proto_depIdxs,
		MessageInfos:      file_album_proto_msgTypes,
	}.Build()
	File_album_proto = out.File
	file_album_proto_rawDesc = nil
	file_album_proto_goTypes = nil
	file_album_proto_depIdxs = nil
}
