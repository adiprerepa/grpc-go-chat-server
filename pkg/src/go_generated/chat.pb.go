// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

package chat

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// username
// chat room id
type RegisterClient struct {
	Username             *string  `protobuf:"bytes,1,req,name=username" json:"username,omitempty"`
	RoomId               *string  `protobuf:"bytes,2,req,name=roomId" json:"roomId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterClient) Reset()         { *m = RegisterClient{} }
func (m *RegisterClient) String() string { return proto.CompactTextString(m) }
func (*RegisterClient) ProtoMessage()    {}
func (*RegisterClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_1f142050368a6cc6, []int{0}
}
func (m *RegisterClient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterClient.Unmarshal(m, b)
}
func (m *RegisterClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterClient.Marshal(b, m, deterministic)
}
func (dst *RegisterClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterClient.Merge(dst, src)
}
func (m *RegisterClient) XXX_Size() int {
	return xxx_messageInfo_RegisterClient.Size(m)
}
func (m *RegisterClient) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterClient.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterClient proto.InternalMessageInfo

func (m *RegisterClient) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func (m *RegisterClient) GetRoomId() string {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return ""
}

// response :
// chat authentication response
// (optional) current users in room
// (maybe) - all messages so far in given chat room
type RegisterClientResponse struct {
	// 1 is ok.
	// 2 is disallowed.
	Status               *int32            `protobuf:"varint,1,req,name=status" json:"status,omitempty"`
	Users                []*RoomUserEntity `protobuf:"bytes,2,rep,name=users" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RegisterClientResponse) Reset()         { *m = RegisterClientResponse{} }
func (m *RegisterClientResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterClientResponse) ProtoMessage()    {}
func (*RegisterClientResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_1f142050368a6cc6, []int{1}
}
func (m *RegisterClientResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterClientResponse.Unmarshal(m, b)
}
func (m *RegisterClientResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterClientResponse.Marshal(b, m, deterministic)
}
func (dst *RegisterClientResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterClientResponse.Merge(dst, src)
}
func (m *RegisterClientResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterClientResponse.Size(m)
}
func (m *RegisterClientResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterClientResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterClientResponse proto.InternalMessageInfo

func (m *RegisterClientResponse) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *RegisterClientResponse) GetUsers() []*RoomUserEntity {
	if m != nil {
		return m.Users
	}
	return nil
}

// user entity for registration
// maybe add custom user features in future?
type RoomUserEntity struct {
	Username             *string  `protobuf:"bytes,1,req,name=username" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomUserEntity) Reset()         { *m = RoomUserEntity{} }
func (m *RoomUserEntity) String() string { return proto.CompactTextString(m) }
func (*RoomUserEntity) ProtoMessage()    {}
func (*RoomUserEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_1f142050368a6cc6, []int{2}
}
func (m *RoomUserEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomUserEntity.Unmarshal(m, b)
}
func (m *RoomUserEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomUserEntity.Marshal(b, m, deterministic)
}
func (dst *RoomUserEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomUserEntity.Merge(dst, src)
}
func (m *RoomUserEntity) XXX_Size() int {
	return xxx_messageInfo_RoomUserEntity.Size(m)
}
func (m *RoomUserEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomUserEntity.DiscardUnknown(m)
}

var xxx_messageInfo_RoomUserEntity proto.InternalMessageInfo

func (m *RoomUserEntity) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

// chat payload
// username
// timestamp
type ChatMessage struct {
	Username             *string  `protobuf:"bytes,1,req,name=username" json:"username,omitempty"`
	Payload              *string  `protobuf:"bytes,2,req,name=payload" json:"payload,omitempty"`
	Time                 *string  `protobuf:"bytes,3,req,name=time" json:"time,omitempty"`
	RoomId               *string  `protobuf:"bytes,4,req,name=roomId" json:"roomId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatMessage) Reset()         { *m = ChatMessage{} }
func (m *ChatMessage) String() string { return proto.CompactTextString(m) }
func (*ChatMessage) ProtoMessage()    {}
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_1f142050368a6cc6, []int{3}
}
func (m *ChatMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatMessage.Unmarshal(m, b)
}
func (m *ChatMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatMessage.Marshal(b, m, deterministic)
}
func (dst *ChatMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatMessage.Merge(dst, src)
}
func (m *ChatMessage) XXX_Size() int {
	return xxx_messageInfo_ChatMessage.Size(m)
}
func (m *ChatMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ChatMessage proto.InternalMessageInfo

func (m *ChatMessage) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func (m *ChatMessage) GetPayload() string {
	if m != nil && m.Payload != nil {
		return *m.Payload
	}
	return ""
}

func (m *ChatMessage) GetTime() string {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return ""
}

func (m *ChatMessage) GetRoomId() string {
	if m != nil && m.RoomId != nil {
		return *m.RoomId
	}
	return ""
}

// Message Response
type ChatMessageResponse struct {
	// 1 means ok
	// 2 means err
	MesssageStatus       *int32   `protobuf:"varint,1,req,name=messsageStatus" json:"messsageStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatMessageResponse) Reset()         { *m = ChatMessageResponse{} }
func (m *ChatMessageResponse) String() string { return proto.CompactTextString(m) }
func (*ChatMessageResponse) ProtoMessage()    {}
func (*ChatMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_1f142050368a6cc6, []int{4}
}
func (m *ChatMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatMessageResponse.Unmarshal(m, b)
}
func (m *ChatMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatMessageResponse.Marshal(b, m, deterministic)
}
func (dst *ChatMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatMessageResponse.Merge(dst, src)
}
func (m *ChatMessageResponse) XXX_Size() int {
	return xxx_messageInfo_ChatMessageResponse.Size(m)
}
func (m *ChatMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ChatMessageResponse proto.InternalMessageInfo

func (m *ChatMessageResponse) GetMesssageStatus() int32 {
	if m != nil && m.MesssageStatus != nil {
		return *m.MesssageStatus
	}
	return 0
}

// for pushing chat updates later
type MessageUpdateReq struct {
	ChatRoomId           *string  `protobuf:"bytes,1,req,name=chatRoomId" json:"chatRoomId,omitempty"`
	Username             *string  `protobuf:"bytes,2,req,name=username" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageUpdateReq) Reset()         { *m = MessageUpdateReq{} }
func (m *MessageUpdateReq) String() string { return proto.CompactTextString(m) }
func (*MessageUpdateReq) ProtoMessage()    {}
func (*MessageUpdateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_1f142050368a6cc6, []int{5}
}
func (m *MessageUpdateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageUpdateReq.Unmarshal(m, b)
}
func (m *MessageUpdateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageUpdateReq.Marshal(b, m, deterministic)
}
func (dst *MessageUpdateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageUpdateReq.Merge(dst, src)
}
func (m *MessageUpdateReq) XXX_Size() int {
	return xxx_messageInfo_MessageUpdateReq.Size(m)
}
func (m *MessageUpdateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageUpdateReq.DiscardUnknown(m)
}

var xxx_messageInfo_MessageUpdateReq proto.InternalMessageInfo

func (m *MessageUpdateReq) GetChatRoomId() string {
	if m != nil && m.ChatRoomId != nil {
		return *m.ChatRoomId
	}
	return ""
}

func (m *MessageUpdateReq) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

type UpdateChatMessage struct {
	Message              *ChatMessage      `protobuf:"bytes,1,req,name=message" json:"message,omitempty"`
	RoomEntities         []*RoomUserEntity `protobuf:"bytes,2,rep,name=roomEntities" json:"roomEntities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *UpdateChatMessage) Reset()         { *m = UpdateChatMessage{} }
func (m *UpdateChatMessage) String() string { return proto.CompactTextString(m) }
func (*UpdateChatMessage) ProtoMessage()    {}
func (*UpdateChatMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_1f142050368a6cc6, []int{6}
}
func (m *UpdateChatMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateChatMessage.Unmarshal(m, b)
}
func (m *UpdateChatMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateChatMessage.Marshal(b, m, deterministic)
}
func (dst *UpdateChatMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateChatMessage.Merge(dst, src)
}
func (m *UpdateChatMessage) XXX_Size() int {
	return xxx_messageInfo_UpdateChatMessage.Size(m)
}
func (m *UpdateChatMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateChatMessage.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateChatMessage proto.InternalMessageInfo

func (m *UpdateChatMessage) GetMessage() *ChatMessage {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *UpdateChatMessage) GetRoomEntities() []*RoomUserEntity {
	if m != nil {
		return m.RoomEntities
	}
	return nil
}

type ExitRequest struct {
	Username             *string  `protobuf:"bytes,1,req,name=username" json:"username,omitempty"`
	ChatId               *string  `protobuf:"bytes,2,req,name=chatId" json:"chatId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExitRequest) Reset()         { *m = ExitRequest{} }
func (m *ExitRequest) String() string { return proto.CompactTextString(m) }
func (*ExitRequest) ProtoMessage()    {}
func (*ExitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_1f142050368a6cc6, []int{7}
}
func (m *ExitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExitRequest.Unmarshal(m, b)
}
func (m *ExitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExitRequest.Marshal(b, m, deterministic)
}
func (dst *ExitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExitRequest.Merge(dst, src)
}
func (m *ExitRequest) XXX_Size() int {
	return xxx_messageInfo_ExitRequest.Size(m)
}
func (m *ExitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExitRequest proto.InternalMessageInfo

func (m *ExitRequest) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func (m *ExitRequest) GetChatId() string {
	if m != nil && m.ChatId != nil {
		return *m.ChatId
	}
	return ""
}

type ExitResponse struct {
	// 1 is permitted
	// 2 is not permitted
	ExitStatus           *int32   `protobuf:"varint,1,req,name=exitStatus" json:"exitStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExitResponse) Reset()         { *m = ExitResponse{} }
func (m *ExitResponse) String() string { return proto.CompactTextString(m) }
func (*ExitResponse) ProtoMessage()    {}
func (*ExitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_1f142050368a6cc6, []int{8}
}
func (m *ExitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExitResponse.Unmarshal(m, b)
}
func (m *ExitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExitResponse.Marshal(b, m, deterministic)
}
func (dst *ExitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExitResponse.Merge(dst, src)
}
func (m *ExitResponse) XXX_Size() int {
	return xxx_messageInfo_ExitResponse.Size(m)
}
func (m *ExitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExitResponse proto.InternalMessageInfo

func (m *ExitResponse) GetExitStatus() int32 {
	if m != nil && m.ExitStatus != nil {
		return *m.ExitStatus
	}
	return 0
}

type RegisterChatRoomRequest struct {
	// register chat room with username as admin
	ChatRoomName         *string  `protobuf:"bytes,1,req,name=chatRoomName" json:"chatRoomName,omitempty"`
	Username             *string  `protobuf:"bytes,2,req,name=username" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterChatRoomRequest) Reset()         { *m = RegisterChatRoomRequest{} }
func (m *RegisterChatRoomRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterChatRoomRequest) ProtoMessage()    {}
func (*RegisterChatRoomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_1f142050368a6cc6, []int{9}
}
func (m *RegisterChatRoomRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterChatRoomRequest.Unmarshal(m, b)
}
func (m *RegisterChatRoomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterChatRoomRequest.Marshal(b, m, deterministic)
}
func (dst *RegisterChatRoomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterChatRoomRequest.Merge(dst, src)
}
func (m *RegisterChatRoomRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterChatRoomRequest.Size(m)
}
func (m *RegisterChatRoomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterChatRoomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterChatRoomRequest proto.InternalMessageInfo

func (m *RegisterChatRoomRequest) GetChatRoomName() string {
	if m != nil && m.ChatRoomName != nil {
		return *m.ChatRoomName
	}
	return ""
}

func (m *RegisterChatRoomRequest) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

type RegisterChatRoomResponse struct {
	// 1 is ok
	// 2 is the chat room already exists
	RegistrationStatus   *int32   `protobuf:"varint,1,req,name=registrationStatus" json:"registrationStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterChatRoomResponse) Reset()         { *m = RegisterChatRoomResponse{} }
func (m *RegisterChatRoomResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterChatRoomResponse) ProtoMessage()    {}
func (*RegisterChatRoomResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chat_1f142050368a6cc6, []int{10}
}
func (m *RegisterChatRoomResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterChatRoomResponse.Unmarshal(m, b)
}
func (m *RegisterChatRoomResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterChatRoomResponse.Marshal(b, m, deterministic)
}
func (dst *RegisterChatRoomResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterChatRoomResponse.Merge(dst, src)
}
func (m *RegisterChatRoomResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterChatRoomResponse.Size(m)
}
func (m *RegisterChatRoomResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterChatRoomResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterChatRoomResponse proto.InternalMessageInfo

func (m *RegisterChatRoomResponse) GetRegistrationStatus() int32 {
	if m != nil && m.RegistrationStatus != nil {
		return *m.RegistrationStatus
	}
	return 0
}

func init() {
	proto.RegisterType((*RegisterClient)(nil), "chat.RegisterClient")
	proto.RegisterType((*RegisterClientResponse)(nil), "chat.RegisterClientResponse")
	proto.RegisterType((*RoomUserEntity)(nil), "chat.RoomUserEntity")
	proto.RegisterType((*ChatMessage)(nil), "chat.ChatMessage")
	proto.RegisterType((*ChatMessageResponse)(nil), "chat.ChatMessageResponse")
	proto.RegisterType((*MessageUpdateReq)(nil), "chat.MessageUpdateReq")
	proto.RegisterType((*UpdateChatMessage)(nil), "chat.UpdateChatMessage")
	proto.RegisterType((*ExitRequest)(nil), "chat.ExitRequest")
	proto.RegisterType((*ExitResponse)(nil), "chat.ExitResponse")
	proto.RegisterType((*RegisterChatRoomRequest)(nil), "chat.RegisterChatRoomRequest")
	proto.RegisterType((*RegisterChatRoomResponse)(nil), "chat.RegisterChatRoomResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServiceClient interface {
	// register user
	RegisterClient(ctx context.Context, in *RegisterClient, opts ...grpc.CallOption) (*RegisterClientResponse, error)
	// messaging
	SendMessage(ctx context.Context, in *ChatMessage, opts ...grpc.CallOption) (*ChatMessageResponse, error)
	// updating messages
	UpdateMessage(ctx context.Context, in *MessageUpdateReq, opts ...grpc.CallOption) (ChatService_UpdateMessageClient, error)
	// leaving
	ExitChat(ctx context.Context, in *ExitRequest, opts ...grpc.CallOption) (*ExitResponse, error)
}

type chatServiceClient struct {
	cc *grpc.ClientConn
}

func NewChatServiceClient(cc *grpc.ClientConn) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) RegisterClient(ctx context.Context, in *RegisterClient, opts ...grpc.CallOption) (*RegisterClientResponse, error) {
	out := new(RegisterClientResponse)
	err := c.cc.Invoke(ctx, "/chat.ChatService/registerClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) SendMessage(ctx context.Context, in *ChatMessage, opts ...grpc.CallOption) (*ChatMessageResponse, error) {
	out := new(ChatMessageResponse)
	err := c.cc.Invoke(ctx, "/chat.ChatService/sendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) UpdateMessage(ctx context.Context, in *MessageUpdateReq, opts ...grpc.CallOption) (ChatService_UpdateMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatService_serviceDesc.Streams[0], "/chat.ChatService/updateMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceUpdateMessageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatService_UpdateMessageClient interface {
	Recv() (*UpdateChatMessage, error)
	grpc.ClientStream
}

type chatServiceUpdateMessageClient struct {
	grpc.ClientStream
}

func (x *chatServiceUpdateMessageClient) Recv() (*UpdateChatMessage, error) {
	m := new(UpdateChatMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) ExitChat(ctx context.Context, in *ExitRequest, opts ...grpc.CallOption) (*ExitResponse, error) {
	out := new(ExitResponse)
	err := c.cc.Invoke(ctx, "/chat.ChatService/exitChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	// register user
	RegisterClient(context.Context, *RegisterClient) (*RegisterClientResponse, error)
	// messaging
	SendMessage(context.Context, *ChatMessage) (*ChatMessageResponse, error)
	// updating messages
	UpdateMessage(*MessageUpdateReq, ChatService_UpdateMessageServer) error
	// leaving
	ExitChat(context.Context, *ExitRequest) (*ExitResponse, error)
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_RegisterClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterClient)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).RegisterClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/RegisterClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).RegisterClient(ctx, req.(*RegisterClient))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SendMessage(ctx, req.(*ChatMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_UpdateMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MessageUpdateReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).UpdateMessage(m, &chatServiceUpdateMessageServer{stream})
}

type ChatService_UpdateMessageServer interface {
	Send(*UpdateChatMessage) error
	grpc.ServerStream
}

type chatServiceUpdateMessageServer struct {
	grpc.ServerStream
}

func (x *chatServiceUpdateMessageServer) Send(m *UpdateChatMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _ChatService_ExitChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).ExitChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/ExitChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).ExitChat(ctx, req.(*ExitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "registerClient",
			Handler:    _ChatService_RegisterClient_Handler,
		},
		{
			MethodName: "sendMessage",
			Handler:    _ChatService_SendMessage_Handler,
		},
		{
			MethodName: "exitChat",
			Handler:    _ChatService_ExitChat_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "updateMessage",
			Handler:       _ChatService_UpdateMessage_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chat.proto",
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor_chat_1f142050368a6cc6) }

var fileDescriptor_chat_1f142050368a6cc6 = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x55, 0xd3, 0x2d, 0x2d, 0x93, 0x25, 0xa2, 0x06, 0x6d, 0x43, 0x84, 0xaa, 0xca, 0x07, 0x54,
	0x01, 0x8a, 0x60, 0x4f, 0x48, 0x88, 0x03, 0xb4, 0x3d, 0x80, 0x44, 0x0f, 0xae, 0x7a, 0x40, 0xe2,
	0x62, 0x75, 0x47, 0xad, 0xa5, 0x26, 0x4e, 0x6d, 0x07, 0xb5, 0xfc, 0x00, 0xbf, 0x8d, 0x62, 0xc7,
	0xa9, 0xbd, 0xa9, 0xd2, 0x5b, 0xe6, 0xd9, 0x9e, 0x79, 0xef, 0xcd, 0x0b, 0xc0, 0xc5, 0x15, 0x37,
	0x65, 0xa3, 0xa4, 0x91, 0x64, 0xd6, 0x7d, 0xd3, 0x63, 0xc8, 0x18, 0x5e, 0x0a, 0x6d, 0x50, 0x1d,
	0x5d, 0x0b, 0xac, 0x0d, 0x29, 0x60, 0xa7, 0xd5, 0xa8, 0x6a, 0x5e, 0x61, 0xbe, 0x71, 0x90, 0x1c,
	0x3e, 0x65, 0x43, 0x4d, 0x16, 0xf0, 0x44, 0x49, 0x59, 0x7d, 0x5f, 0xe5, 0x89, 0x3d, 0xe9, 0x2b,
	0xfa, 0x1b, 0x16, 0x71, 0x17, 0x86, 0xba, 0x91, 0xb5, 0xb6, 0x2f, 0xb4, 0xe1, 0xa6, 0xd5, 0xb6,
	0xd7, 0x16, 0xeb, 0x2b, 0xf2, 0x16, 0xb6, 0xba, 0xae, 0x3a, 0x4f, 0x0e, 0x36, 0x0f, 0xd3, 0xe5,
	0xcb, 0xd2, 0x32, 0x63, 0x52, 0x56, 0xe7, 0x1a, 0xd5, 0x49, 0x6d, 0x84, 0xb9, 0x63, 0xee, 0x0a,
	0x7d, 0x0f, 0x59, 0x7c, 0x30, 0xc5, 0x91, 0x4a, 0x48, 0x8f, 0xae, 0xb8, 0xf9, 0x89, 0x5a, 0xf3,
	0x4b, 0x9c, 0x94, 0x93, 0xc3, 0x76, 0xc3, 0xef, 0xae, 0x25, 0xf7, 0x7a, 0x7c, 0x49, 0x08, 0xcc,
	0x8c, 0xa8, 0x30, 0xdf, 0xb4, 0xb0, 0xfd, 0x0e, 0xc4, 0xcf, 0x22, 0xf1, 0x5f, 0xe0, 0x45, 0x30,
	0x70, 0x50, 0xfe, 0x06, 0xb2, 0x0a, 0xb5, 0xc5, 0xce, 0x42, 0x07, 0xd6, 0x50, 0x7a, 0x0a, 0xcf,
	0xfb, 0xa7, 0xe7, 0xcd, 0x8a, 0x1b, 0x64, 0x78, 0x43, 0xf6, 0xdd, 0xa6, 0x98, 0x1b, 0xe7, 0x68,
	0x07, 0x48, 0x24, 0x2a, 0x59, 0xd3, 0xff, 0x17, 0x76, 0x5d, 0xa3, 0xd0, 0x85, 0x77, 0xb0, 0x5d,
	0xb9, 0x4f, 0xdb, 0x2d, 0x5d, 0xee, 0x3a, 0xc3, 0x43, 0xe2, 0xfe, 0x06, 0xf9, 0x04, 0xf3, 0x4e,
	0x9a, 0xf5, 0x5a, 0xe0, 0xf4, 0x8a, 0xa2, 0x9b, 0xf4, 0x2b, 0xa4, 0x27, 0xb7, 0xc2, 0x30, 0xbc,
	0x69, 0x51, 0x3f, 0x1a, 0xa5, 0xae, 0xdf, 0x7d, 0x94, 0x5c, 0x45, 0x4b, 0x98, 0xbb, 0x16, 0xbd,
	0x8d, 0xfb, 0x00, 0x78, 0x2b, 0x4c, 0x64, 0x61, 0x80, 0xd0, 0x5f, 0xb0, 0x37, 0x44, 0xaf, 0x37,
	0xc8, 0x8f, 0xa7, 0x30, 0xf7, 0x9e, 0x9d, 0xde, 0x53, 0x88, 0xb0, 0x49, 0x27, 0x7f, 0x40, 0x3e,
	0x6e, 0xdd, 0xd3, 0x2a, 0x81, 0x28, 0x7b, 0xa6, 0xb8, 0x11, 0xb2, 0x8e, 0xe8, 0x3d, 0x70, 0xb2,
	0xfc, 0x97, 0xb8, 0x58, 0x9e, 0xa1, 0xfa, 0x23, 0x2e, 0x90, 0x1c, 0x43, 0xa6, 0xe2, 0xff, 0xce,
	0xfb, 0x1b, 0xa1, 0xc5, 0xeb, 0x87, 0xd0, 0x81, 0xc5, 0x67, 0x48, 0x35, 0xd6, 0x2b, 0xbf, 0xe5,
	0xf1, 0x52, 0x8b, 0x57, 0xe3, 0x3d, 0xfb, 0xc7, 0xdf, 0xe0, 0x59, 0x6b, 0x83, 0xe2, 0x9f, 0x2f,
	0xdc, 0xdd, 0xf5, 0x34, 0x16, 0x7b, 0x0e, 0x1f, 0xa5, 0xea, 0xc3, 0x06, 0xf9, 0x08, 0x3b, 0xdd,
	0x2e, 0x3a, 0xd0, 0x4f, 0x0f, 0x02, 0x50, 0x90, 0x10, 0x72, 0x63, 0xff, 0x07, 0x00, 0x00, 0xff,
	0xff, 0x90, 0x8d, 0x71, 0x0e, 0x84, 0x04, 0x00, 0x00,
}
