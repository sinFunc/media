// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: msapi.proto

package rpc

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

type CodecType int32

const (
	CodecType_RAW             CodecType = 0 // raw packet data
	CodecType_TELEPHONE_EVENT CodecType = 1
	CodecType_PCM_ALAW        CodecType = 2
	CodecType_AMRNB           CodecType = 3
	CodecType_AMRWB           CodecType = 4
)

// Enum value maps for CodecType.
var (
	CodecType_name = map[int32]string{
		0: "RAW",
		1: "TELEPHONE_EVENT",
		2: "PCM_ALAW",
		3: "AMRNB",
		4: "AMRWB",
	}
	CodecType_value = map[string]int32{
		"RAW":             0,
		"TELEPHONE_EVENT": 1,
		"PCM_ALAW":        2,
		"AMRNB":           3,
		"AMRWB":           4,
	}
)

func (x CodecType) Enum() *CodecType {
	p := new(CodecType)
	*p = x
	return p
}

func (x CodecType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CodecType) Descriptor() protoreflect.EnumDescriptor {
	return file_msapi_proto_enumTypes[0].Descriptor()
}

func (CodecType) Type() protoreflect.EnumType {
	return &file_msapi_proto_enumTypes[0]
}

func (x CodecType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CodecType.Descriptor instead.
func (CodecType) EnumDescriptor() ([]byte, []int) {
	return file_msapi_proto_rawDescGZIP(), []int{0}
}

type SystemCommand int32

const (
	SystemCommand_USER_EVENT   SystemCommand = 0 // used by other subsystem
	SystemCommand_REGISTER     SystemCommand = 1
	SystemCommand_KEEPALIVE    SystemCommand = 2
	SystemCommand_SESSION_INFO SystemCommand = 3
)

// Enum value maps for SystemCommand.
var (
	SystemCommand_name = map[int32]string{
		0: "USER_EVENT",
		1: "REGISTER",
		2: "KEEPALIVE",
		3: "SESSION_INFO",
	}
	SystemCommand_value = map[string]int32{
		"USER_EVENT":   0,
		"REGISTER":     1,
		"KEEPALIVE":    2,
		"SESSION_INFO": 3,
	}
)

func (x SystemCommand) Enum() *SystemCommand {
	p := new(SystemCommand)
	*p = x
	return p
}

func (x SystemCommand) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SystemCommand) Descriptor() protoreflect.EnumDescriptor {
	return file_msapi_proto_enumTypes[1].Descriptor()
}

func (SystemCommand) Type() protoreflect.EnumType {
	return &file_msapi_proto_enumTypes[1]
}

func (x SystemCommand) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SystemCommand.Descriptor instead.
func (SystemCommand) EnumDescriptor() ([]byte, []int) {
	return file_msapi_proto_rawDescGZIP(), []int{1}
}

type CodecInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PayloadNumber uint32    `protobuf:"varint,1,opt,name=payload_number,json=payloadNumber,proto3" json:"payload_number,omitempty"`              // negotiated payload type, dynamic(96 ~ 127) or fixed type
	PayloadType   CodecType `protobuf:"varint,2,opt,name=payload_type,json=payloadType,proto3,enum=rpc.CodecType" json:"payload_type,omitempty"` // used to identify mime type, like "AMR","PCM_ALAW"
	CodecParam    string    `protobuf:"bytes,3,opt,name=codec_param,json=codecParam,proto3" json:"codec_param,omitempty"`                        // parameter of codec, like fmtp: ...
}

func (x *CodecInfo) Reset() {
	*x = CodecInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodecInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodecInfo) ProtoMessage() {}

func (x *CodecInfo) ProtoReflect() protoreflect.Message {
	mi := &file_msapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodecInfo.ProtoReflect.Descriptor instead.
func (*CodecInfo) Descriptor() ([]byte, []int) {
	return file_msapi_proto_rawDescGZIP(), []int{0}
}

func (x *CodecInfo) GetPayloadNumber() uint32 {
	if x != nil {
		return x.PayloadNumber
	}
	return 0
}

func (x *CodecInfo) GetPayloadType() CodecType {
	if x != nil {
		return x.PayloadType
	}
	return CodecType_RAW
}

func (x *CodecInfo) GetCodecParam() string {
	if x != nil {
		return x.CodecParam
	}
	return ""
}

type CreateParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerIp     string       `protobuf:"bytes,1,opt,name=peer_ip,json=peerIp,proto3" json:"peer_ip,omitempty"`        // remote rtp ip
	PeerPort   uint32       `protobuf:"varint,2,opt,name=peer_port,json=peerPort,proto3" json:"peer_port,omitempty"` // remote rtp port
	Codecs     []*CodecInfo `protobuf:"bytes,3,rep,name=codecs,proto3" json:"codecs,omitempty"`
	GraphDesc  string       `protobuf:"bytes,4,opt,name=graph_desc,json=graphDesc,proto3" json:"graph_desc,omitempty"`    // used to describe event graph
	InstanceId string       `protobuf:"bytes,5,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"` // which instance creates this session
}

func (x *CreateParam) Reset() {
	*x = CreateParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateParam) ProtoMessage() {}

func (x *CreateParam) ProtoReflect() protoreflect.Message {
	mi := &file_msapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateParam.ProtoReflect.Descriptor instead.
func (*CreateParam) Descriptor() ([]byte, []int) {
	return file_msapi_proto_rawDescGZIP(), []int{1}
}

func (x *CreateParam) GetPeerIp() string {
	if x != nil {
		return x.PeerIp
	}
	return ""
}

func (x *CreateParam) GetPeerPort() uint32 {
	if x != nil {
		return x.PeerPort
	}
	return 0
}

func (x *CreateParam) GetCodecs() []*CodecInfo {
	if x != nil {
		return x.Codecs
	}
	return nil
}

func (x *CreateParam) GetGraphDesc() string {
	if x != nil {
		return x.GraphDesc
	}
	return ""
}

func (x *CreateParam) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

type StartParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *StartParam) Reset() {
	*x = StartParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartParam) ProtoMessage() {}

func (x *StartParam) ProtoReflect() protoreflect.Message {
	mi := &file_msapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartParam.ProtoReflect.Descriptor instead.
func (*StartParam) Descriptor() ([]byte, []int) {
	return file_msapi_proto_rawDescGZIP(), []int{2}
}

func (x *StartParam) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

type StopParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *StopParam) Reset() {
	*x = StopParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msapi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopParam) ProtoMessage() {}

func (x *StopParam) ProtoReflect() protoreflect.Message {
	mi := &file_msapi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopParam.ProtoReflect.Descriptor instead.
func (*StopParam) Descriptor() ([]byte, []int) {
	return file_msapi_proto_rawDescGZIP(), []int{3}
}

func (x *StopParam) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msapi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_msapi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_msapi_proto_rawDescGZIP(), []int{4}
}

func (x *Status) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId    string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	LocalIp      string `protobuf:"bytes,2,opt,name=local_ip,json=localIp,proto3" json:"local_ip,omitempty"`
	LocalRtpPort uint32 `protobuf:"varint,3,opt,name=local_rtp_port,json=localRtpPort,proto3" json:"local_rtp_port,omitempty"`
	PeerIp       string `protobuf:"bytes,4,opt,name=peer_ip,json=peerIp,proto3" json:"peer_ip,omitempty"`
	PeerRtpPort  uint32 `protobuf:"varint,5,opt,name=peer_rtp_port,json=peerRtpPort,proto3" json:"peer_rtp_port,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msapi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_msapi_proto_msgTypes[5]
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
	return file_msapi_proto_rawDescGZIP(), []int{5}
}

func (x *Session) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *Session) GetLocalIp() string {
	if x != nil {
		return x.LocalIp
	}
	return ""
}

func (x *Session) GetLocalRtpPort() uint32 {
	if x != nil {
		return x.LocalRtpPort
	}
	return 0
}

func (x *Session) GetPeerIp() string {
	if x != nil {
		return x.PeerIp
	}
	return ""
}

func (x *Session) GetPeerRtpPort() uint32 {
	if x != nil {
		return x.PeerRtpPort
	}
	return 0
}

type Action struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Cmd       string `protobuf:"bytes,2,opt,name=cmd,proto3" json:"cmd,omitempty"`
	CmdArg    string `protobuf:"bytes,3,opt,name=cmd_arg,json=cmdArg,proto3" json:"cmd_arg,omitempty"`
}

func (x *Action) Reset() {
	*x = Action{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msapi_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

func (x *Action) ProtoReflect() protoreflect.Message {
	mi := &file_msapi_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action.ProtoReflect.Descriptor instead.
func (*Action) Descriptor() ([]byte, []int) {
	return file_msapi_proto_rawDescGZIP(), []int{6}
}

func (x *Action) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *Action) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

func (x *Action) GetCmdArg() string {
	if x != nil {
		return x.CmdArg
	}
	return ""
}

type ActionResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	State     string `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *ActionResult) Reset() {
	*x = ActionResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msapi_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionResult) ProtoMessage() {}

func (x *ActionResult) ProtoReflect() protoreflect.Message {
	mi := &file_msapi_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionResult.ProtoReflect.Descriptor instead.
func (*ActionResult) Descriptor() ([]byte, []int) {
	return file_msapi_proto_rawDescGZIP(), []int{7}
}

func (x *ActionResult) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *ActionResult) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type ActionEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Event     string `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *ActionEvent) Reset() {
	*x = ActionEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msapi_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionEvent) ProtoMessage() {}

func (x *ActionEvent) ProtoReflect() protoreflect.Message {
	mi := &file_msapi_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionEvent.ProtoReflect.Descriptor instead.
func (*ActionEvent) Descriptor() ([]byte, []int) {
	return file_msapi_proto_rawDescGZIP(), []int{8}
}

func (x *ActionEvent) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *ActionEvent) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

type SystemEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd        SystemCommand `protobuf:"varint,1,opt,name=cmd,proto3,enum=rpc.SystemCommand" json:"cmd,omitempty"`
	InstanceId string        `protobuf:"bytes,2,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	SessionId  string        `protobuf:"bytes,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Event      string        `protobuf:"bytes,4,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *SystemEvent) Reset() {
	*x = SystemEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msapi_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemEvent) ProtoMessage() {}

func (x *SystemEvent) ProtoReflect() protoreflect.Message {
	mi := &file_msapi_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemEvent.ProtoReflect.Descriptor instead.
func (*SystemEvent) Descriptor() ([]byte, []int) {
	return file_msapi_proto_rawDescGZIP(), []int{9}
}

func (x *SystemEvent) GetCmd() SystemCommand {
	if x != nil {
		return x.Cmd
	}
	return SystemCommand_USER_EVENT
}

func (x *SystemEvent) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *SystemEvent) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *SystemEvent) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

var File_msapi_proto protoreflect.FileDescriptor

var file_msapi_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x73, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72,
	0x70, 0x63, 0x22, 0x86, 0x01, 0x0a, 0x09, 0x43, 0x6f, 0x64, 0x65, 0x63, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x63, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f,
	0x64, 0x65, 0x63, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x22, 0xab, 0x01, 0x0a, 0x0b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x70,
	0x65, 0x65, 0x72, 0x5f, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65,
	0x65, 0x72, 0x49, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x65, 0x65, 0x72, 0x50, 0x6f, 0x72,
	0x74, 0x12, 0x26, 0x0a, 0x06, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x63, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x06, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x22, 0x2b, 0x0a, 0x0a, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x2a, 0x0a, 0x09, 0x53, 0x74, 0x6f, 0x70, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x22, 0x20, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0xa6, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x49, 0x70, 0x12, 0x24, 0x0a, 0x0e, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x5f, 0x72, 0x74, 0x70, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x74, 0x70, 0x50, 0x6f, 0x72, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x70, 0x12, 0x22, 0x0a, 0x0d, 0x70, 0x65, 0x65,
	0x72, 0x5f, 0x72, 0x74, 0x70, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x70, 0x65, 0x65, 0x72, 0x52, 0x74, 0x70, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x52, 0x0a,
	0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x6d, 0x64, 0x5f,
	0x61, 0x72, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6d, 0x64, 0x41, 0x72,
	0x67, 0x22, 0x43, 0x0a, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x42, 0x0a, 0x0b, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x89, 0x01, 0x0a, 0x0b, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x03, 0x63, 0x6d,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x03, 0x63, 0x6d, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2a, 0x4d, 0x0a, 0x09, 0x43, 0x6f, 0x64, 0x65, 0x63, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x41, 0x57, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f,
	0x54, 0x45, 0x4c, 0x45, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x10,
	0x01, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x43, 0x4d, 0x5f, 0x41, 0x4c, 0x41, 0x57, 0x10, 0x02, 0x12,
	0x09, 0x0a, 0x05, 0x41, 0x4d, 0x52, 0x4e, 0x42, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x4d,
	0x52, 0x57, 0x42, 0x10, 0x04, 0x2a, 0x4e, 0x0a, 0x0d, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x45,
	0x56, 0x45, 0x4e, 0x54, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54,
	0x45, 0x52, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x4b, 0x45, 0x45, 0x50, 0x41, 0x4c, 0x49, 0x56,
	0x45, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x49,
	0x4e, 0x46, 0x4f, 0x10, 0x03, 0x32, 0xc8, 0x02, 0x0a, 0x08, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x41,
	0x70, 0x69, 0x12, 0x32, 0x0a, 0x0e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x0c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x0b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x70, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x6f, 0x70,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x0b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x11, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x17, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x12, 0x0b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x10, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x39, 0x0a, 0x0d, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x10, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x10, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01,
	0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x70, 0x70, 0x63, 0x72, 0x61, 0x73, 0x68, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msapi_proto_rawDescOnce sync.Once
	file_msapi_proto_rawDescData = file_msapi_proto_rawDesc
)

func file_msapi_proto_rawDescGZIP() []byte {
	file_msapi_proto_rawDescOnce.Do(func() {
		file_msapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_msapi_proto_rawDescData)
	})
	return file_msapi_proto_rawDescData
}

var file_msapi_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_msapi_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_msapi_proto_goTypes = []interface{}{
	(CodecType)(0),       // 0: rpc.CodecType
	(SystemCommand)(0),   // 1: rpc.SystemCommand
	(*CodecInfo)(nil),    // 2: rpc.CodecInfo
	(*CreateParam)(nil),  // 3: rpc.CreateParam
	(*StartParam)(nil),   // 4: rpc.StartParam
	(*StopParam)(nil),    // 5: rpc.StopParam
	(*Status)(nil),       // 6: rpc.Status
	(*Session)(nil),      // 7: rpc.Session
	(*Action)(nil),       // 8: rpc.Action
	(*ActionResult)(nil), // 9: rpc.ActionResult
	(*ActionEvent)(nil),  // 10: rpc.ActionEvent
	(*SystemEvent)(nil),  // 11: rpc.SystemEvent
}
var file_msapi_proto_depIdxs = []int32{
	0,  // 0: rpc.CodecInfo.payload_type:type_name -> rpc.CodecType
	2,  // 1: rpc.CreateParam.codecs:type_name -> rpc.CodecInfo
	1,  // 2: rpc.SystemEvent.cmd:type_name -> rpc.SystemCommand
	3,  // 3: rpc.MediaApi.PrepareSession:input_type -> rpc.CreateParam
	4,  // 4: rpc.MediaApi.StartSession:input_type -> rpc.StartParam
	5,  // 5: rpc.MediaApi.StopSession:input_type -> rpc.StopParam
	8,  // 6: rpc.MediaApi.ExecuteAction:input_type -> rpc.Action
	8,  // 7: rpc.MediaApi.ExecuteActionWithNotify:input_type -> rpc.Action
	11, // 8: rpc.MediaApi.SystemChannel:input_type -> rpc.SystemEvent
	7,  // 9: rpc.MediaApi.PrepareSession:output_type -> rpc.Session
	6,  // 10: rpc.MediaApi.StartSession:output_type -> rpc.Status
	6,  // 11: rpc.MediaApi.StopSession:output_type -> rpc.Status
	9,  // 12: rpc.MediaApi.ExecuteAction:output_type -> rpc.ActionResult
	10, // 13: rpc.MediaApi.ExecuteActionWithNotify:output_type -> rpc.ActionEvent
	11, // 14: rpc.MediaApi.SystemChannel:output_type -> rpc.SystemEvent
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_msapi_proto_init() }
func file_msapi_proto_init() {
	if File_msapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodecInfo); i {
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
		file_msapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateParam); i {
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
		file_msapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartParam); i {
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
		file_msapi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopParam); i {
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
		file_msapi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_msapi_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_msapi_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Action); i {
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
		file_msapi_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionResult); i {
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
		file_msapi_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionEvent); i {
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
		file_msapi_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemEvent); i {
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
			RawDescriptor: file_msapi_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_msapi_proto_goTypes,
		DependencyIndexes: file_msapi_proto_depIdxs,
		EnumInfos:         file_msapi_proto_enumTypes,
		MessageInfos:      file_msapi_proto_msgTypes,
	}.Build()
	File_msapi_proto = out.File
	file_msapi_proto_rawDesc = nil
	file_msapi_proto_goTypes = nil
	file_msapi_proto_depIdxs = nil
}
