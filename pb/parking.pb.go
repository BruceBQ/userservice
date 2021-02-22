// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: pb/parking.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

type StatusCode int32

const (
	StatusCode_UNKNOW           StatusCode = 0
	StatusCode_OK               StatusCode = 200
	StatusCode_FORBIDDEN        StatusCode = 403
	StatusCode_NOT_FOUND        StatusCode = 404
	StatusCode_METHOD_NOT_ALLOW StatusCode = 405
	StatusCode_INTERNAL_ERROR   StatusCode = 500
)

// Enum value maps for StatusCode.
var (
	StatusCode_name = map[int32]string{
		0:   "UNKNOW",
		200: "OK",
		403: "FORBIDDEN",
		404: "NOT_FOUND",
		405: "METHOD_NOT_ALLOW",
		500: "INTERNAL_ERROR",
	}
	StatusCode_value = map[string]int32{
		"UNKNOW":           0,
		"OK":               200,
		"FORBIDDEN":        403,
		"NOT_FOUND":        404,
		"METHOD_NOT_ALLOW": 405,
		"INTERNAL_ERROR":   500,
	}
)

func (x StatusCode) Enum() *StatusCode {
	p := new(StatusCode)
	*p = x
	return p
}

func (x StatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_parking_proto_enumTypes[0].Descriptor()
}

func (StatusCode) Type() protoreflect.EnumType {
	return &file_pb_parking_proto_enumTypes[0]
}

func (x StatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusCode.Descriptor instead.
func (StatusCode) EnumDescriptor() ([]byte, []int) {
	return file_pb_parking_proto_rawDescGZIP(), []int{0}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_parking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_pb_parking_proto_msgTypes[0]
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
	return file_pb_parking_proto_rawDescGZIP(), []int{0}
}

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token          string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserId         string `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
	PermissionName string `protobuf:"bytes,2,opt,name=permissionName,proto3" json:"permissionName,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_parking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_pb_parking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_pb_parking_proto_rawDescGZIP(), []int{1}
}

func (x *Session) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Session) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Session) GetPermissionName() string {
	if x != nil {
		return x.PermissionName
	}
	return ""
}

type SessionCamera struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token          string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserId         string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	PermissionName string `protobuf:"bytes,3,opt,name=permissionName,proto3" json:"permissionName,omitempty"`
	CameraId       string `protobuf:"bytes,4,opt,name=cameraId,proto3" json:"cameraId,omitempty"`
}

func (x *SessionCamera) Reset() {
	*x = SessionCamera{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_parking_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionCamera) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionCamera) ProtoMessage() {}

func (x *SessionCamera) ProtoReflect() protoreflect.Message {
	mi := &file_pb_parking_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionCamera.ProtoReflect.Descriptor instead.
func (*SessionCamera) Descriptor() ([]byte, []int) {
	return file_pb_parking_proto_rawDescGZIP(), []int{2}
}

func (x *SessionCamera) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SessionCamera) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SessionCamera) GetPermissionName() string {
	if x != nil {
		return x.PermissionName
	}
	return ""
}

func (x *SessionCamera) GetCameraId() string {
	if x != nil {
		return x.CameraId
	}
	return ""
}

type AuthorizationResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode StatusCode `protobuf:"varint,1,opt,name=statusCode,proto3,enum=parking.StatusCode" json:"statusCode,omitempty"`
	Message    string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AuthorizationResult) Reset() {
	*x = AuthorizationResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_parking_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizationResult) ProtoMessage() {}

func (x *AuthorizationResult) ProtoReflect() protoreflect.Message {
	mi := &file_pb_parking_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizationResult.ProtoReflect.Descriptor instead.
func (*AuthorizationResult) Descriptor() ([]byte, []int) {
	return file_pb_parking_proto_rawDescGZIP(), []int{3}
}

func (x *AuthorizationResult) GetStatusCode() StatusCode {
	if x != nil {
		return x.StatusCode
	}
	return StatusCode_UNKNOW
}

func (x *AuthorizationResult) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UserId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *UserId) Reset() {
	*x = UserId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_parking_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserId) ProtoMessage() {}

func (x *UserId) ProtoReflect() protoreflect.Message {
	mi := &file_pb_parking_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserId.ProtoReflect.Descriptor instead.
func (*UserId) Descriptor() ([]byte, []int) {
	return file_pb_parking_proto_rawDescGZIP(), []int{4}
}

func (x *UserId) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type CameraId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CameraId string `protobuf:"bytes,1,opt,name=cameraId,proto3" json:"cameraId,omitempty"`
}

func (x *CameraId) Reset() {
	*x = CameraId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_parking_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CameraId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CameraId) ProtoMessage() {}

func (x *CameraId) ProtoReflect() protoreflect.Message {
	mi := &file_pb_parking_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CameraId.ProtoReflect.Descriptor instead.
func (*CameraId) Descriptor() ([]byte, []int) {
	return file_pb_parking_proto_rawDescGZIP(), []int{5}
}

func (x *CameraId) GetCameraId() string {
	if x != nil {
		return x.CameraId
	}
	return ""
}

type CameraList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cameras []string `protobuf:"bytes,1,rep,name=cameras,proto3" json:"cameras,omitempty"`
}

func (x *CameraList) Reset() {
	*x = CameraList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_parking_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CameraList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CameraList) ProtoMessage() {}

func (x *CameraList) ProtoReflect() protoreflect.Message {
	mi := &file_pb_parking_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CameraList.ProtoReflect.Descriptor instead.
func (*CameraList) Descriptor() ([]byte, []int) {
	return file_pb_parking_proto_rawDescGZIP(), []int{6}
}

func (x *CameraList) GetCameras() []string {
	if x != nil {
		return x.Cameras
	}
	return nil
}

type GetCameraRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids map[string]string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetCameraRequest) Reset() {
	*x = GetCameraRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_parking_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCameraRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCameraRequest) ProtoMessage() {}

func (x *GetCameraRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_parking_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCameraRequest.ProtoReflect.Descriptor instead.
func (*GetCameraRequest) Descriptor() ([]byte, []int) {
	return file_pb_parking_proto_rawDescGZIP(), []int{7}
}

func (x *GetCameraRequest) GetIds() map[string]string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type GetCameraReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cameras map[string]*CameraData `protobuf:"bytes,1,rep,name=cameras,proto3" json:"cameras,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetCameraReply) Reset() {
	*x = GetCameraReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_parking_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCameraReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCameraReply) ProtoMessage() {}

func (x *GetCameraReply) ProtoReflect() protoreflect.Message {
	mi := &file_pb_parking_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCameraReply.ProtoReflect.Descriptor instead.
func (*GetCameraReply) Descriptor() ([]byte, []int) {
	return file_pb_parking_proto_rawDescGZIP(), []int{8}
}

func (x *GetCameraReply) GetCameras() map[string]*CameraData {
	if x != nil {
		return x.Cameras
	}
	return nil
}

type CameraData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address string  `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Lat     float64 `protobuf:"fixed64,4,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng     float64 `protobuf:"fixed64,5,opt,name=lng,proto3" json:"lng,omitempty"`
}

func (x *CameraData) Reset() {
	*x = CameraData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_parking_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CameraData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CameraData) ProtoMessage() {}

func (x *CameraData) ProtoReflect() protoreflect.Message {
	mi := &file_pb_parking_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CameraData.ProtoReflect.Descriptor instead.
func (*CameraData) Descriptor() ([]byte, []int) {
	return file_pb_parking_proto_rawDescGZIP(), []int{9}
}

func (x *CameraData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CameraData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CameraData) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CameraData) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *CameraData) GetLng() float64 {
	if x != nil {
		return x.Lng
	}
	return 0
}

var File_pb_parking_proto protoreflect.FileDescriptor

var file_pb_parking_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0x07, 0x0a, 0x05, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x5f, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a,
	0x0e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x49, 0x64, 0x22, 0x64, 0x0a, 0x13, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x33, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x20, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x26, 0x0a, 0x08, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x49, 0x64, 0x22, 0x26, 0x0a, 0x0a, 0x43, 0x61, 0x6d,
	0x65, 0x72, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x6d, 0x65, 0x72,
	0x61, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61,
	0x73, 0x22, 0x80, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x49,
	0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x69, 0x64, 0x73, 0x1a, 0x36, 0x0a, 0x08,
	0x49, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0xa1, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6d, 0x65,
	0x72, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3e, 0x0a, 0x07, 0x63, 0x61, 0x6d, 0x65, 0x72,
	0x61, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x2e, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07,
	0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x73, 0x1a, 0x4f, 0x0a, 0x0c, 0x43, 0x61, 0x6d, 0x65, 0x72,
	0x61, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x29, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x6e, 0x0a, 0x0a, 0x43, 0x61, 0x6d, 0x65,
	0x72, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6e, 0x67, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x2a, 0x6d, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x10, 0x00, 0x12, 0x07, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0xc8, 0x01, 0x12, 0x0e, 0x0a, 0x09, 0x46,
	0x4f, 0x52, 0x42, 0x49, 0x44, 0x44, 0x45, 0x4e, 0x10, 0x93, 0x03, 0x12, 0x0e, 0x0a, 0x09, 0x4e,
	0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x94, 0x03, 0x12, 0x15, 0x0a, 0x10, 0x4d,
	0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x10,
	0x95, 0x03, 0x12, 0x13, 0x0a, 0x0e, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0xf4, 0x03, 0x32, 0xac, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x16, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x48, 0x61, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54,
	0x6f, 0x12, 0x10, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x1c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x61,
	0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x43, 0x61, 0x6d,
	0x65, 0x72, 0x61, 0x12, 0x16, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x1a, 0x1c, 0x2e, 0x70, 0x61,
	0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x0f, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x1a, 0x13, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x61, 0x6d,
	0x65, 0x72, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x14, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x46, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x11, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x61, 0x6d, 0x65,
	0x72, 0x61, 0x49, 0x64, 0x1a, 0x0e, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x32, 0x0e, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x08, 0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x32, 0x45, 0x0a, 0x06, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x12, 0x3b, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x19, 0x2e, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x6d, 0x65, 0x72, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70,
	0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6d, 0x65, 0x72, 0x61,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x32, 0x09, 0x0a, 0x07, 0x50, 0x61, 0x72, 0x6b, 0x69,
	0x6e, 0x67, 0x32, 0x06, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x10, 0x5a, 0x0e, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_parking_proto_rawDescOnce sync.Once
	file_pb_parking_proto_rawDescData = file_pb_parking_proto_rawDesc
)

func file_pb_parking_proto_rawDescGZIP() []byte {
	file_pb_parking_proto_rawDescOnce.Do(func() {
		file_pb_parking_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_parking_proto_rawDescData)
	})
	return file_pb_parking_proto_rawDescData
}

var file_pb_parking_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_parking_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_pb_parking_proto_goTypes = []interface{}{
	(StatusCode)(0),             // 0: parking.StatusCode
	(*Empty)(nil),               // 1: parking.Empty
	(*Session)(nil),             // 2: parking.Session
	(*SessionCamera)(nil),       // 3: parking.SessionCamera
	(*AuthorizationResult)(nil), // 4: parking.AuthorizationResult
	(*UserId)(nil),              // 5: parking.UserId
	(*CameraId)(nil),            // 6: parking.CameraId
	(*CameraList)(nil),          // 7: parking.CameraList
	(*GetCameraRequest)(nil),    // 8: parking.GetCameraRequest
	(*GetCameraReply)(nil),      // 9: parking.GetCameraReply
	(*CameraData)(nil),          // 10: parking.CameraData
	nil,                         // 11: parking.GetCameraRequest.IdsEntry
	nil,                         // 12: parking.GetCameraReply.CamerasEntry
}
var file_pb_parking_proto_depIdxs = []int32{
	0,  // 0: parking.AuthorizationResult.statusCode:type_name -> parking.StatusCode
	11, // 1: parking.GetCameraRequest.ids:type_name -> parking.GetCameraRequest.IdsEntry
	12, // 2: parking.GetCameraReply.cameras:type_name -> parking.GetCameraReply.CamerasEntry
	10, // 3: parking.GetCameraReply.CamerasEntry.value:type_name -> parking.CameraData
	2,  // 4: parking.UserService.SessionHasPermissionTo:input_type -> parking.Session
	3,  // 5: parking.UserService.SessionHasPermissionToCamera:input_type -> parking.SessionCamera
	5,  // 6: parking.UserService.GetCamerasByUserId:input_type -> parking.UserId
	6,  // 7: parking.UserService.DeleteCameraFromUser:input_type -> parking.CameraId
	8,  // 8: parking.Camera.Get:input_type -> parking.GetCameraRequest
	4,  // 9: parking.UserService.SessionHasPermissionTo:output_type -> parking.AuthorizationResult
	4,  // 10: parking.UserService.SessionHasPermissionToCamera:output_type -> parking.AuthorizationResult
	7,  // 11: parking.UserService.GetCamerasByUserId:output_type -> parking.CameraList
	1,  // 12: parking.UserService.DeleteCameraFromUser:output_type -> parking.Empty
	9,  // 13: parking.Camera.Get:output_type -> parking.GetCameraReply
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_pb_parking_proto_init() }
func file_pb_parking_proto_init() {
	if File_pb_parking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_parking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_pb_parking_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
		file_pb_parking_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionCamera); i {
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
		file_pb_parking_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizationResult); i {
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
		file_pb_parking_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserId); i {
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
		file_pb_parking_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CameraId); i {
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
		file_pb_parking_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CameraList); i {
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
		file_pb_parking_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCameraRequest); i {
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
		file_pb_parking_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCameraReply); i {
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
		file_pb_parking_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CameraData); i {
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
			RawDescriptor: file_pb_parking_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   6,
		},
		GoTypes:           file_pb_parking_proto_goTypes,
		DependencyIndexes: file_pb_parking_proto_depIdxs,
		EnumInfos:         file_pb_parking_proto_enumTypes,
		MessageInfos:      file_pb_parking_proto_msgTypes,
	}.Build()
	File_pb_parking_proto = out.File
	file_pb_parking_proto_rawDesc = nil
	file_pb_parking_proto_goTypes = nil
	file_pb_parking_proto_depIdxs = nil
}