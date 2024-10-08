// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: api/check/v1/check.proto

package check

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId  int64  `protobuf:"varint,1,opt,name=postId,proto3" json:"postId,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Title   string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	UserId  int64  `protobuf:"varint,4,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *CreateCheckRequest) Reset() {
	*x = CreateCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_check_v1_check_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCheckRequest) ProtoMessage() {}

func (x *CreateCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_check_v1_check_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCheckRequest.ProtoReflect.Descriptor instead.
func (*CreateCheckRequest) Descriptor() ([]byte, []int) {
	return file_api_check_v1_check_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCheckRequest) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *CreateCheckRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateCheckRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateCheckRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CreateCheckReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *CreateCheckReply) Reset() {
	*x = CreateCheckReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_check_v1_check_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCheckReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCheckReply) ProtoMessage() {}

func (x *CreateCheckReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_check_v1_check_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCheckReply.ProtoReflect.Descriptor instead.
func (*CreateCheckReply) Descriptor() ([]byte, []int) {
	return file_api_check_v1_check_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCheckReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CreateCheckReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type DeleteCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckId int64 `protobuf:"varint,1,opt,name=checkId,proto3" json:"checkId,omitempty"`
}

func (x *DeleteCheckRequest) Reset() {
	*x = DeleteCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_check_v1_check_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCheckRequest) ProtoMessage() {}

func (x *DeleteCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_check_v1_check_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCheckRequest.ProtoReflect.Descriptor instead.
func (*DeleteCheckRequest) Descriptor() ([]byte, []int) {
	return file_api_check_v1_check_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteCheckRequest) GetCheckId() int64 {
	if x != nil {
		return x.CheckId
	}
	return 0
}

type DeleteCheckReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *DeleteCheckReply) Reset() {
	*x = DeleteCheckReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_check_v1_check_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCheckReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCheckReply) ProtoMessage() {}

func (x *DeleteCheckReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_check_v1_check_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCheckReply.ProtoReflect.Descriptor instead.
func (*DeleteCheckReply) Descriptor() ([]byte, []int) {
	return file_api_check_v1_check_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteCheckReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DeleteCheckReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type GetCheckByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckId int64 `protobuf:"varint,1,opt,name=checkId,proto3" json:"checkId,omitempty"`
}

func (x *GetCheckByIdRequest) Reset() {
	*x = GetCheckByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_check_v1_check_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCheckByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCheckByIdRequest) ProtoMessage() {}

func (x *GetCheckByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_check_v1_check_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCheckByIdRequest.ProtoReflect.Descriptor instead.
func (*GetCheckByIdRequest) Descriptor() ([]byte, []int) {
	return file_api_check_v1_check_proto_rawDescGZIP(), []int{4}
}

func (x *GetCheckByIdRequest) GetCheckId() int64 {
	if x != nil {
		return x.CheckId
	}
	return 0
}

type GetCheckByIdReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32           `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string          `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data *ListOrGetCheck `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetCheckByIdReply) Reset() {
	*x = GetCheckByIdReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_check_v1_check_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCheckByIdReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCheckByIdReply) ProtoMessage() {}

func (x *GetCheckByIdReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_check_v1_check_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCheckByIdReply.ProtoReflect.Descriptor instead.
func (*GetCheckByIdReply) Descriptor() ([]byte, []int) {
	return file_api_check_v1_check_proto_rawDescGZIP(), []int{5}
}

func (x *GetCheckByIdReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetCheckByIdReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetCheckByIdReply) GetData() *ListOrGetCheck {
	if x != nil {
		return x.Data
	}
	return nil
}

type ListChecksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   int32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size   int64  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Status uint32 `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ListChecksRequest) Reset() {
	*x = ListChecksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_check_v1_check_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListChecksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChecksRequest) ProtoMessage() {}

func (x *ListChecksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_check_v1_check_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChecksRequest.ProtoReflect.Descriptor instead.
func (*ListChecksRequest) Descriptor() ([]byte, []int) {
	return file_api_check_v1_check_proto_rawDescGZIP(), []int{6}
}

func (x *ListChecksRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListChecksRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListChecksRequest) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type ListChecksReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string            `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data []*ListOrGetCheck `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListChecksReply) Reset() {
	*x = ListChecksReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_check_v1_check_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListChecksReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChecksReply) ProtoMessage() {}

func (x *ListChecksReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_check_v1_check_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChecksReply.ProtoReflect.Descriptor instead.
func (*ListChecksReply) Descriptor() ([]byte, []int) {
	return file_api_check_v1_check_proto_rawDescGZIP(), []int{7}
}

func (x *ListChecksReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ListChecksReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ListChecksReply) GetData() []*ListOrGetCheck {
	if x != nil {
		return x.Data
	}
	return nil
}

type SubmitCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckId int64  `protobuf:"varint,1,opt,name=checkId,proto3" json:"checkId,omitempty"`
	Status  uint32 `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	Remark  string `protobuf:"bytes,4,opt,name=remark,proto3" json:"remark,omitempty"`
}

func (x *SubmitCheckRequest) Reset() {
	*x = SubmitCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_check_v1_check_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitCheckRequest) ProtoMessage() {}

func (x *SubmitCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_check_v1_check_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitCheckRequest.ProtoReflect.Descriptor instead.
func (*SubmitCheckRequest) Descriptor() ([]byte, []int) {
	return file_api_check_v1_check_proto_rawDescGZIP(), []int{8}
}

func (x *SubmitCheckRequest) GetCheckId() int64 {
	if x != nil {
		return x.CheckId
	}
	return 0
}

func (x *SubmitCheckRequest) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SubmitCheckRequest) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type SubmitCheckReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SubmitCheckReply) Reset() {
	*x = SubmitCheckReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_check_v1_check_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitCheckReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitCheckReply) ProtoMessage() {}

func (x *SubmitCheckReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_check_v1_check_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitCheckReply.ProtoReflect.Descriptor instead.
func (*SubmitCheckReply) Descriptor() ([]byte, []int) {
	return file_api_check_v1_check_proto_rawDescGZIP(), []int{9}
}

func (x *SubmitCheckReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SubmitCheckReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ListOrGetCheck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                               // 审核ID
	PostId    int64                  `protobuf:"varint,2,opt,name=postId,proto3" json:"postId,omitempty"`                       // 帖子ID
	Content   string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`                      // 审核内容
	Title     string                 `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`                          // 审核标签
	UserId    int64                  `protobuf:"varint,5,opt,name=userId,proto3" json:"userId,omitempty"`                       // 提交审核的用户ID
	Status    uint32                 `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`                       // 审核状态
	Remark    string                 `protobuf:"bytes,7,opt,name=remark,proto3" json:"remark,omitempty"`                        // 审核备注
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"` // 创建时间
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"` // 更新时间
}

func (x *ListOrGetCheck) Reset() {
	*x = ListOrGetCheck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_check_v1_check_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrGetCheck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrGetCheck) ProtoMessage() {}

func (x *ListOrGetCheck) ProtoReflect() protoreflect.Message {
	mi := &file_api_check_v1_check_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrGetCheck.ProtoReflect.Descriptor instead.
func (*ListOrGetCheck) Descriptor() ([]byte, []int) {
	return file_api_check_v1_check_proto_rawDescGZIP(), []int{10}
}

func (x *ListOrGetCheck) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListOrGetCheck) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *ListOrGetCheck) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ListOrGetCheck) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ListOrGetCheck) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ListOrGetCheck) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ListOrGetCheck) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *ListOrGetCheck) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ListOrGetCheck) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_api_check_v1_check_proto protoreflect.FileDescriptor

var file_api_check_v1_check_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70,
	0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x38, 0x0a,
	0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x2e, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x22, 0x2f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x49, 0x64, 0x22, 0x6b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x30, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f,
	0x72, 0x47, 0x65, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x53, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x69, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x30, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f,
	0x72, 0x47, 0x65, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x5e, 0x0a, 0x12, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72,
	0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x22,
	0x38, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0xa6, 0x02, 0x0a, 0x0e, 0x4c, 0x69,
	0x73, 0x74, 0x4f, 0x72, 0x47, 0x65, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f,
	0x73, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x32, 0xf5, 0x03, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x4f, 0x0a, 0x0b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x20, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x6a, 0x0a,
	0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x20, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x19,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x2a, 0x11, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2f,
	0x7b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x64, 0x7d, 0x12, 0x6a, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x79, 0x49, 0x64, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x7b, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x49, 0x64, 0x7d, 0x12, 0x5e, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x73, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x3a, 0x01, 0x2a, 0x22, 0x05,
	0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x63, 0x0a, 0x0b, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x3a, 0x01,
	0x2a, 0x22, 0x07, 0x2f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x6f, 0x53, 0x69, 0x6d, 0x70, 0x6c,
	0x69, 0x63, 0x69, 0x74, 0x79, 0x2f, 0x4c, 0x69, 0x6e, 0x6b, 0x4d, 0x65, 0x2d, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_check_v1_check_proto_rawDescOnce sync.Once
	file_api_check_v1_check_proto_rawDescData = file_api_check_v1_check_proto_rawDesc
)

func file_api_check_v1_check_proto_rawDescGZIP() []byte {
	file_api_check_v1_check_proto_rawDescOnce.Do(func() {
		file_api_check_v1_check_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_check_v1_check_proto_rawDescData)
	})
	return file_api_check_v1_check_proto_rawDescData
}

var file_api_check_v1_check_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_check_v1_check_proto_goTypes = []any{
	(*CreateCheckRequest)(nil),    // 0: api.check.v1.CreateCheckRequest
	(*CreateCheckReply)(nil),      // 1: api.check.v1.CreateCheckReply
	(*DeleteCheckRequest)(nil),    // 2: api.check.v1.DeleteCheckRequest
	(*DeleteCheckReply)(nil),      // 3: api.check.v1.DeleteCheckReply
	(*GetCheckByIdRequest)(nil),   // 4: api.check.v1.GetCheckByIdRequest
	(*GetCheckByIdReply)(nil),     // 5: api.check.v1.GetCheckByIdReply
	(*ListChecksRequest)(nil),     // 6: api.check.v1.ListChecksRequest
	(*ListChecksReply)(nil),       // 7: api.check.v1.ListChecksReply
	(*SubmitCheckRequest)(nil),    // 8: api.check.v1.SubmitCheckRequest
	(*SubmitCheckReply)(nil),      // 9: api.check.v1.SubmitCheckReply
	(*ListOrGetCheck)(nil),        // 10: api.check.v1.ListOrGetCheck
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
}
var file_api_check_v1_check_proto_depIdxs = []int32{
	10, // 0: api.check.v1.GetCheckByIdReply.data:type_name -> api.check.v1.ListOrGetCheck
	10, // 1: api.check.v1.ListChecksReply.data:type_name -> api.check.v1.ListOrGetCheck
	11, // 2: api.check.v1.ListOrGetCheck.created_at:type_name -> google.protobuf.Timestamp
	11, // 3: api.check.v1.ListOrGetCheck.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 4: api.check.v1.Check.CreateCheck:input_type -> api.check.v1.CreateCheckRequest
	2,  // 5: api.check.v1.Check.DeleteCheck:input_type -> api.check.v1.DeleteCheckRequest
	4,  // 6: api.check.v1.Check.GetCheckById:input_type -> api.check.v1.GetCheckByIdRequest
	6,  // 7: api.check.v1.Check.ListChecks:input_type -> api.check.v1.ListChecksRequest
	8,  // 8: api.check.v1.Check.SubmitCheck:input_type -> api.check.v1.SubmitCheckRequest
	1,  // 9: api.check.v1.Check.CreateCheck:output_type -> api.check.v1.CreateCheckReply
	3,  // 10: api.check.v1.Check.DeleteCheck:output_type -> api.check.v1.DeleteCheckReply
	5,  // 11: api.check.v1.Check.GetCheckById:output_type -> api.check.v1.GetCheckByIdReply
	7,  // 12: api.check.v1.Check.ListChecks:output_type -> api.check.v1.ListChecksReply
	9,  // 13: api.check.v1.Check.SubmitCheck:output_type -> api.check.v1.SubmitCheckReply
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_api_check_v1_check_proto_init() }
func file_api_check_v1_check_proto_init() {
	if File_api_check_v1_check_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_check_v1_check_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateCheckRequest); i {
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
		file_api_check_v1_check_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateCheckReply); i {
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
		file_api_check_v1_check_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteCheckRequest); i {
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
		file_api_check_v1_check_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteCheckReply); i {
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
		file_api_check_v1_check_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetCheckByIdRequest); i {
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
		file_api_check_v1_check_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetCheckByIdReply); i {
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
		file_api_check_v1_check_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ListChecksRequest); i {
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
		file_api_check_v1_check_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*ListChecksReply); i {
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
		file_api_check_v1_check_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*SubmitCheckRequest); i {
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
		file_api_check_v1_check_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*SubmitCheckReply); i {
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
		file_api_check_v1_check_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*ListOrGetCheck); i {
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
			RawDescriptor: file_api_check_v1_check_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_check_v1_check_proto_goTypes,
		DependencyIndexes: file_api_check_v1_check_proto_depIdxs,
		MessageInfos:      file_api_check_v1_check_proto_msgTypes,
	}.Build()
	File_api_check_v1_check_proto = out.File
	file_api_check_v1_check_proto_rawDesc = nil
	file_api_check_v1_check_proto_goTypes = nil
	file_api_check_v1_check_proto_depIdxs = nil
}
