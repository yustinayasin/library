// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: author.proto

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

type AuthorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Bio       string                 `protobuf:"bytes,3,opt,name=bio,proto3" json:"bio,omitempty"`
	Birthdate *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=birthdate,proto3" json:"birthdate,omitempty"`
}

func (x *AuthorRequest) Reset() {
	*x = AuthorRequest{}
	mi := &file_author_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorRequest) ProtoMessage() {}

func (x *AuthorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorRequest.ProtoReflect.Descriptor instead.
func (*AuthorRequest) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{0}
}

func (x *AuthorRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AuthorRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AuthorRequest) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *AuthorRequest) GetBirthdate() *timestamppb.Timestamp {
	if x != nil {
		return x.Birthdate
	}
	return nil
}

type AuthorIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AuthorIdRequest) Reset() {
	*x = AuthorIdRequest{}
	mi := &file_author_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthorIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorIdRequest) ProtoMessage() {}

func (x *AuthorIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorIdRequest.ProtoReflect.Descriptor instead.
func (*AuthorIdRequest) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{1}
}

func (x *AuthorIdRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type AuthorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Bio       string                 `protobuf:"bytes,3,opt,name=bio,proto3" json:"bio,omitempty"`
	Birthdate *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=birthdate,proto3" json:"birthdate,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *AuthorResponse) Reset() {
	*x = AuthorResponse{}
	mi := &file_author_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorResponse) ProtoMessage() {}

func (x *AuthorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorResponse.ProtoReflect.Descriptor instead.
func (*AuthorResponse) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{2}
}

func (x *AuthorResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AuthorResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AuthorResponse) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *AuthorResponse) GetBirthdate() *timestamppb.Timestamp {
	if x != nil {
		return x.Birthdate
	}
	return nil
}

func (x *AuthorResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *AuthorResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type BaseResponseAuthor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string          `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Author  *AuthorResponse `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *BaseResponseAuthor) Reset() {
	*x = BaseResponseAuthor{}
	mi := &file_author_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BaseResponseAuthor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResponseAuthor) ProtoMessage() {}

func (x *BaseResponseAuthor) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResponseAuthor.ProtoReflect.Descriptor instead.
func (*BaseResponseAuthor) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{3}
}

func (x *BaseResponseAuthor) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *BaseResponseAuthor) GetAuthor() *AuthorResponse {
	if x != nil {
		return x.Author
	}
	return nil
}

type AuthorResponseExist struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists bool `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
}

func (x *AuthorResponseExist) Reset() {
	*x = AuthorResponseExist{}
	mi := &file_author_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthorResponseExist) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorResponseExist) ProtoMessage() {}

func (x *AuthorResponseExist) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorResponseExist.ProtoReflect.Descriptor instead.
func (*AuthorResponseExist) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{4}
}

func (x *AuthorResponseExist) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

type AuthorsBooksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorId int32 `protobuf:"varint,1,opt,name=authorId,proto3" json:"authorId,omitempty"`
	BookId   int32 `protobuf:"varint,2,opt,name=bookId,proto3" json:"bookId,omitempty"`
}

func (x *AuthorsBooksRequest) Reset() {
	*x = AuthorsBooksRequest{}
	mi := &file_author_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthorsBooksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorsBooksRequest) ProtoMessage() {}

func (x *AuthorsBooksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorsBooksRequest.ProtoReflect.Descriptor instead.
func (*AuthorsBooksRequest) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{5}
}

func (x *AuthorsBooksRequest) GetAuthorId() int32 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *AuthorsBooksRequest) GetBookId() int32 {
	if x != nil {
		return x.BookId
	}
	return 0
}

type AuthorsBooksIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AuthorsBooksIdRequest) Reset() {
	*x = AuthorsBooksIdRequest{}
	mi := &file_author_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthorsBooksIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorsBooksIdRequest) ProtoMessage() {}

func (x *AuthorsBooksIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorsBooksIdRequest.ProtoReflect.Descriptor instead.
func (*AuthorsBooksIdRequest) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{6}
}

func (x *AuthorsBooksIdRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type AuthorsBooksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorId  int32                  `protobuf:"varint,1,opt,name=authorId,proto3" json:"authorId,omitempty"`
	BookId    int32                  `protobuf:"varint,2,opt,name=bookId,proto3" json:"bookId,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *AuthorsBooksResponse) Reset() {
	*x = AuthorsBooksResponse{}
	mi := &file_author_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthorsBooksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorsBooksResponse) ProtoMessage() {}

func (x *AuthorsBooksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorsBooksResponse.ProtoReflect.Descriptor instead.
func (*AuthorsBooksResponse) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{7}
}

func (x *AuthorsBooksResponse) GetAuthorId() int32 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *AuthorsBooksResponse) GetBookId() int32 {
	if x != nil {
		return x.BookId
	}
	return 0
}

func (x *AuthorsBooksResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *AuthorsBooksResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type BaseResponseAuthorBooks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message      string                `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Authorsbooks *AuthorsBooksResponse `protobuf:"bytes,2,opt,name=authorsbooks,proto3" json:"authorsbooks,omitempty"`
}

func (x *BaseResponseAuthorBooks) Reset() {
	*x = BaseResponseAuthorBooks{}
	mi := &file_author_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BaseResponseAuthorBooks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResponseAuthorBooks) ProtoMessage() {}

func (x *BaseResponseAuthorBooks) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResponseAuthorBooks.ProtoReflect.Descriptor instead.
func (*BaseResponseAuthorBooks) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{8}
}

func (x *BaseResponseAuthorBooks) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *BaseResponseAuthorBooks) GetAuthorsbooks() *AuthorsBooksResponse {
	if x != nil {
		return x.Authorsbooks
	}
	return nil
}

var File_author_proto protoreflect.FileDescriptor

var file_author_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x62, 0x69, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x69, 0x6f, 0x12,
	0x38, 0x0a, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x74, 0x65, 0x22, 0x21, 0x0a, 0x0f, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0xf4, 0x01, 0x0a,
	0x0e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x62, 0x69, 0x6f, 0x12, 0x38, 0x0a, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x5f, 0x0a, 0x12, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x22, 0x2d, 0x0a, 0x13, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x78, 0x69, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x65,
	0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x78, 0x69,
	0x73, 0x74, 0x73, 0x22, 0x49, 0x0a, 0x13, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x42, 0x6f,
	0x6f, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x22, 0x27,
	0x0a, 0x15, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0xbe, 0x01, 0x0a, 0x14, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x6f,
	0x6f, 0x6b, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38,
	0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x76, 0x0a, 0x17, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x42, 0x6f,
	0x6f, 0x6b, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x41, 0x0a,
	0x0c, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x32, 0xe9, 0x02, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12,
	0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x73, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x12, 0x41, 0x0a, 0x0a, 0x45, 0x64, 0x69, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x12, 0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x73, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x45, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x42,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x18, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x12, 0x48, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x45,
	0x78, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x78, 0x69, 0x73, 0x74, 0x32, 0xe9, 0x02, 0x0a,
	0x13, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x1c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x52, 0x0a, 0x10, 0x45, 0x64, 0x69, 0x74, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x1c, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x42, 0x6f, 0x6f,
	0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x73, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x56, 0x0a, 0x12, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x42, 0x6f, 0x6f, 0x6b,
	0x73, 0x12, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x2e, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x42, 0x6f,
	0x6f, 0x6b, 0x73, 0x12, 0x53, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73,
	0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_author_proto_rawDescOnce sync.Once
	file_author_proto_rawDescData = file_author_proto_rawDesc
)

func file_author_proto_rawDescGZIP() []byte {
	file_author_proto_rawDescOnce.Do(func() {
		file_author_proto_rawDescData = protoimpl.X.CompressGZIP(file_author_proto_rawDescData)
	})
	return file_author_proto_rawDescData
}

var file_author_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_author_proto_goTypes = []any{
	(*AuthorRequest)(nil),           // 0: authors.AuthorRequest
	(*AuthorIdRequest)(nil),         // 1: authors.AuthorIdRequest
	(*AuthorResponse)(nil),          // 2: authors.AuthorResponse
	(*BaseResponseAuthor)(nil),      // 3: authors.BaseResponseAuthor
	(*AuthorResponseExist)(nil),     // 4: authors.AuthorResponseExist
	(*AuthorsBooksRequest)(nil),     // 5: authors.AuthorsBooksRequest
	(*AuthorsBooksIdRequest)(nil),   // 6: authors.AuthorsBooksIdRequest
	(*AuthorsBooksResponse)(nil),    // 7: authors.AuthorsBooksResponse
	(*BaseResponseAuthorBooks)(nil), // 8: authors.BaseResponseAuthorBooks
	(*timestamppb.Timestamp)(nil),   // 9: google.protobuf.Timestamp
}
var file_author_proto_depIdxs = []int32{
	9,  // 0: authors.AuthorRequest.birthdate:type_name -> google.protobuf.Timestamp
	9,  // 1: authors.AuthorResponse.birthdate:type_name -> google.protobuf.Timestamp
	9,  // 2: authors.AuthorResponse.CreatedAt:type_name -> google.protobuf.Timestamp
	9,  // 3: authors.AuthorResponse.UpdatedAt:type_name -> google.protobuf.Timestamp
	2,  // 4: authors.BaseResponseAuthor.author:type_name -> authors.AuthorResponse
	9,  // 5: authors.AuthorsBooksResponse.CreatedAt:type_name -> google.protobuf.Timestamp
	9,  // 6: authors.AuthorsBooksResponse.UpdatedAt:type_name -> google.protobuf.Timestamp
	7,  // 7: authors.BaseResponseAuthorBooks.authorsbooks:type_name -> authors.AuthorsBooksResponse
	0,  // 8: authors.AuthorService.AddAuthor:input_type -> authors.AuthorRequest
	0,  // 9: authors.AuthorService.EditAuthor:input_type -> authors.AuthorRequest
	1,  // 10: authors.AuthorService.DeleteAuthor:input_type -> authors.AuthorIdRequest
	1,  // 11: authors.AuthorService.GetAuthor:input_type -> authors.AuthorIdRequest
	1,  // 12: authors.AuthorService.GetAuthorExist:input_type -> authors.AuthorIdRequest
	5,  // 13: authors.AuthorsBooksService.AddAuthorsBooks:input_type -> authors.AuthorsBooksRequest
	5,  // 14: authors.AuthorsBooksService.EditAuthorsBooks:input_type -> authors.AuthorsBooksRequest
	6,  // 15: authors.AuthorsBooksService.DeleteAuthorsBooks:input_type -> authors.AuthorsBooksIdRequest
	6,  // 16: authors.AuthorsBooksService.GetAuthorsBooks:input_type -> authors.AuthorsBooksIdRequest
	3,  // 17: authors.AuthorService.AddAuthor:output_type -> authors.BaseResponseAuthor
	3,  // 18: authors.AuthorService.EditAuthor:output_type -> authors.BaseResponseAuthor
	3,  // 19: authors.AuthorService.DeleteAuthor:output_type -> authors.BaseResponseAuthor
	3,  // 20: authors.AuthorService.GetAuthor:output_type -> authors.BaseResponseAuthor
	4,  // 21: authors.AuthorService.GetAuthorExist:output_type -> authors.AuthorResponseExist
	8,  // 22: authors.AuthorsBooksService.AddAuthorsBooks:output_type -> authors.BaseResponseAuthorBooks
	8,  // 23: authors.AuthorsBooksService.EditAuthorsBooks:output_type -> authors.BaseResponseAuthorBooks
	8,  // 24: authors.AuthorsBooksService.DeleteAuthorsBooks:output_type -> authors.BaseResponseAuthorBooks
	8,  // 25: authors.AuthorsBooksService.GetAuthorsBooks:output_type -> authors.BaseResponseAuthorBooks
	17, // [17:26] is the sub-list for method output_type
	8,  // [8:17] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_author_proto_init() }
func file_author_proto_init() {
	if File_author_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_author_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_author_proto_goTypes,
		DependencyIndexes: file_author_proto_depIdxs,
		MessageInfos:      file_author_proto_msgTypes,
	}.Build()
	File_author_proto = out.File
	file_author_proto_rawDesc = nil
	file_author_proto_goTypes = nil
	file_author_proto_depIdxs = nil
}
