// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: taskpb/api.proto

package taskpb

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

type GetAllTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllTaskRequest) Reset() {
	*x = GetAllTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_taskpb_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTaskRequest) ProtoMessage() {}

func (x *GetAllTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_taskpb_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTaskRequest.ProtoReflect.Descriptor instead.
func (*GetAllTaskRequest) Descriptor() ([]byte, []int) {
	return file_taskpb_api_proto_rawDescGZIP(), []int{0}
}

type GetAllTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*Task `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *GetAllTaskResponse) Reset() {
	*x = GetAllTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_taskpb_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTaskResponse) ProtoMessage() {}

func (x *GetAllTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_taskpb_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTaskResponse.ProtoReflect.Descriptor instead.
func (*GetAllTaskResponse) Descriptor() ([]byte, []int) {
	return file_taskpb_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllTaskResponse) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type CreateTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateTaskRequest) Reset() {
	*x = CreateTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_taskpb_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskRequest) ProtoMessage() {}

func (x *CreateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_taskpb_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskRequest.ProtoReflect.Descriptor instead.
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return file_taskpb_api_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTaskRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateTaskRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task *Task `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *CreateTaskResponse) Reset() {
	*x = CreateTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_taskpb_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskResponse) ProtoMessage() {}

func (x *CreateTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_taskpb_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskResponse.ProtoReflect.Descriptor instead.
func (*CreateTaskResponse) Descriptor() ([]byte, []int) {
	return file_taskpb_api_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTaskResponse) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type TaskDoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TaskDoneRequest) Reset() {
	*x = TaskDoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_taskpb_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskDoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskDoneRequest) ProtoMessage() {}

func (x *TaskDoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_taskpb_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskDoneRequest.ProtoReflect.Descriptor instead.
func (*TaskDoneRequest) Descriptor() ([]byte, []int) {
	return file_taskpb_api_proto_rawDescGZIP(), []int{4}
}

func (x *TaskDoneRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TaskDoneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task *Task `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *TaskDoneResponse) Reset() {
	*x = TaskDoneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_taskpb_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskDoneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskDoneResponse) ProtoMessage() {}

func (x *TaskDoneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_taskpb_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskDoneResponse.ProtoReflect.Descriptor instead.
func (*TaskDoneResponse) Descriptor() ([]byte, []int) {
	return file_taskpb_api_proto_rawDescGZIP(), []int{5}
}

func (x *TaskDoneResponse) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type TaskInProgressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TaskInProgressRequest) Reset() {
	*x = TaskInProgressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_taskpb_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskInProgressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskInProgressRequest) ProtoMessage() {}

func (x *TaskInProgressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_taskpb_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskInProgressRequest.ProtoReflect.Descriptor instead.
func (*TaskInProgressRequest) Descriptor() ([]byte, []int) {
	return file_taskpb_api_proto_rawDescGZIP(), []int{6}
}

func (x *TaskInProgressRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TaskInProgressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task *Task `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *TaskInProgressResponse) Reset() {
	*x = TaskInProgressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_taskpb_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskInProgressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskInProgressResponse) ProtoMessage() {}

func (x *TaskInProgressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_taskpb_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskInProgressResponse.ProtoReflect.Descriptor instead.
func (*TaskInProgressResponse) Descriptor() ([]byte, []int) {
	return file_taskpb_api_proto_rawDescGZIP(), []int{7}
}

func (x *TaskInProgressResponse) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

var File_taskpb_api_proto protoreflect.FileDescriptor

var file_taskpb_api_proto_rawDesc = []byte{
	0x0a, 0x10, 0x74, 0x61, 0x73, 0x6b, 0x70, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x70, 0x62, 0x1a, 0x15, 0x74, 0x61, 0x73, 0x6b,
	0x70, 0x62, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x38, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05,
	0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x61,
	0x73, 0x6b, 0x70, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73,
	0x22, 0x4b, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x36, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x70, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x21, 0x0a, 0x0f, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x6f, 0x6e,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x34, 0x0a, 0x10, 0x54, 0x61, 0x73, 0x6b,
	0x44, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x04,
	0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x70, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x27,
	0x0a, 0x15, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3a, 0x0a, 0x16, 0x54, 0x61, 0x73, 0x6b, 0x49,
	0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x20, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x70, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x74,
	0x61, 0x73, 0x6b, 0x32, 0xaf, 0x02, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x19, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x19, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x70,
	0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3f, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x6f, 0x6e, 0x65, 0x12, 0x17, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x70, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x6f, 0x6e, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x70, 0x62, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x44, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x51, 0x0a, 0x0e, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x70, 0x62, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x49, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x70, 0x62, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x49, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x6c, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x70, 0x62, 0x42, 0x08, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x1c, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x74, 0x61, 0x73, 0x6b, 0x70, 0x62, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x70, 0x62, 0xa2, 0x02, 0x03,
	0x54, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x54, 0x61, 0x73, 0x6b, 0x70, 0x62, 0xca, 0x02, 0x06, 0x54,
	0x61, 0x73, 0x6b, 0x70, 0x62, 0xe2, 0x02, 0x12, 0x54, 0x61, 0x73, 0x6b, 0x70, 0x62, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x06, 0x54, 0x61, 0x73,
	0x6b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_taskpb_api_proto_rawDescOnce sync.Once
	file_taskpb_api_proto_rawDescData = file_taskpb_api_proto_rawDesc
)

func file_taskpb_api_proto_rawDescGZIP() []byte {
	file_taskpb_api_proto_rawDescOnce.Do(func() {
		file_taskpb_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_taskpb_api_proto_rawDescData)
	})
	return file_taskpb_api_proto_rawDescData
}

var file_taskpb_api_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_taskpb_api_proto_goTypes = []interface{}{
	(*GetAllTaskRequest)(nil),      // 0: taskpb.GetAllTaskRequest
	(*GetAllTaskResponse)(nil),     // 1: taskpb.GetAllTaskResponse
	(*CreateTaskRequest)(nil),      // 2: taskpb.CreateTaskRequest
	(*CreateTaskResponse)(nil),     // 3: taskpb.CreateTaskResponse
	(*TaskDoneRequest)(nil),        // 4: taskpb.TaskDoneRequest
	(*TaskDoneResponse)(nil),       // 5: taskpb.TaskDoneResponse
	(*TaskInProgressRequest)(nil),  // 6: taskpb.TaskInProgressRequest
	(*TaskInProgressResponse)(nil), // 7: taskpb.TaskInProgressResponse
	(*Task)(nil),                   // 8: taskpb.Task
}
var file_taskpb_api_proto_depIdxs = []int32{
	8, // 0: taskpb.GetAllTaskResponse.tasks:type_name -> taskpb.Task
	8, // 1: taskpb.CreateTaskResponse.task:type_name -> taskpb.Task
	8, // 2: taskpb.TaskDoneResponse.task:type_name -> taskpb.Task
	8, // 3: taskpb.TaskInProgressResponse.task:type_name -> taskpb.Task
	0, // 4: taskpb.TaskService.GetAllTask:input_type -> taskpb.GetAllTaskRequest
	2, // 5: taskpb.TaskService.CreateTask:input_type -> taskpb.CreateTaskRequest
	4, // 6: taskpb.TaskService.TaskDone:input_type -> taskpb.TaskDoneRequest
	6, // 7: taskpb.TaskService.TaskInProgress:input_type -> taskpb.TaskInProgressRequest
	1, // 8: taskpb.TaskService.GetAllTask:output_type -> taskpb.GetAllTaskResponse
	3, // 9: taskpb.TaskService.CreateTask:output_type -> taskpb.CreateTaskResponse
	5, // 10: taskpb.TaskService.TaskDone:output_type -> taskpb.TaskDoneResponse
	7, // 11: taskpb.TaskService.TaskInProgress:output_type -> taskpb.TaskInProgressResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_taskpb_api_proto_init() }
func file_taskpb_api_proto_init() {
	if File_taskpb_api_proto != nil {
		return
	}
	file_taskpb_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_taskpb_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllTaskRequest); i {
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
		file_taskpb_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllTaskResponse); i {
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
		file_taskpb_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTaskRequest); i {
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
		file_taskpb_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTaskResponse); i {
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
		file_taskpb_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskDoneRequest); i {
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
		file_taskpb_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskDoneResponse); i {
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
		file_taskpb_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskInProgressRequest); i {
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
		file_taskpb_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskInProgressResponse); i {
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
			RawDescriptor: file_taskpb_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_taskpb_api_proto_goTypes,
		DependencyIndexes: file_taskpb_api_proto_depIdxs,
		MessageInfos:      file_taskpb_api_proto_msgTypes,
	}.Build()
	File_taskpb_api_proto = out.File
	file_taskpb_api_proto_rawDesc = nil
	file_taskpb_api_proto_goTypes = nil
	file_taskpb_api_proto_depIdxs = nil
}
