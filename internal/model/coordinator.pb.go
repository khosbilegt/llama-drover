package model

import (
	// protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	// reflect "reflect"
	// sync "sync"
	// unsafe "unsafe"
)

// This is a compile-time assertion that a sufficiently up-to-date version is being used.
const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RequestType int32

const (
	RequestType_UNSPECIFIED RequestType = 0
	RequestType_REGISTER    RequestType = 1
	RequestType_UNREGISTER  RequestType = 2
	RequestType_ERROR       RequestType = 3
	RequestType_PROMPT      RequestType = 4
)

type RegisterRequest struct {
	State  protoimpl.MessageState `protogen:"open.v1"`
	NodeId string                 `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
}

type UnregisterRequest struct {
	State  protoimpl.MessageState `protogen:"open.v1"`
	NodeId string                 `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
}

type PromptRequest struct {
	State         protoimpl.MessageState `protogen:"open.v1"`
	Prompt        string                 `protobuf:"bytes,2,opt,name=prompt,proto3" json:"prompt,omitempty"`
	PromptHistory []string               `protobuf:"bytes,3,rep,name=prompt_history,json=promptHistory,proto3" json:"prompt_history,omitempty"`
}

type MessageResponse struct {
	State   protoimpl.MessageState `protogen:"open.v1"`
	NodeId  string                 `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Message string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

type UnregisterResponse struct {
	State  protoimpl.MessageState `protogen:"open.v1"`
	NodeId string                 `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
}

type PromptResponse struct {
	State         protoimpl.MessageState `protogen:"open.v1"`
	Response      string                 `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	PromptHistory []string               `protobuf:"bytes,2,rep,name=prompt_history,json=promptHistory,proto3" json:"prompt_history,omitempty"`
	IsComplete    bool                   `protobuf:"varint,3,opt,name=is_complete,json=isComplete,proto3" json:"is_complete,omitempty"`
}

type ErrorResponse struct {
	State   protoimpl.MessageState `protogen:"open.v1"`
	Message string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

type ConnectRequest struct {
	State     protoimpl.MessageState `protogen:"open.v1"`
	Type      RequestType            `protobuf:"varint,1,opt,name=type,proto3,enum=routeguide.RequestType" json:"type,omitempty"`
	Timestamp int64                  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Types that are assignable to Request:
	//
	//	*ConnectRequest_Register
	//	*ConnectRequest_Unregister
	//	*ConnectRequest_Prompt
	Request isConnectRequest_Request `protobuf_oneof:"request"`
}

type isConnectRequest_Request interface {
	isConnectRequest_Request()
}

type ConnectRequest_Register struct {
	Register *RegisterRequest `protobuf:"bytes,3,opt,name=register,proto3,oneof"`
}

type ConnectRequest_Unregister struct {
	Unregister *UnregisterRequest `protobuf:"bytes,4,opt,name=unregister,proto3,oneof"`
}

type ConnectRequest_Prompt struct {
	Prompt *PromptRequest `protobuf:"bytes,5,opt,name=prompt,proto3,oneof"`
}

func (*ConnectRequest_Register) isConnectRequest_Request()   {}
func (*ConnectRequest_Unregister) isConnectRequest_Request() {}
func (*ConnectRequest_Prompt) isConnectRequest_Request()     {}

// GetRegister returns the register request if present
func (x *ConnectRequest) GetRegister() *RegisterRequest {
	if x, ok := x.GetRequest().(*ConnectRequest_Register); ok {
		return x.Register
	}
	return nil
}

// GetUnregister returns the unregister request if present
func (x *ConnectRequest) GetUnregister() *UnregisterRequest {
	if x, ok := x.GetRequest().(*ConnectRequest_Unregister); ok {
		return x.Unregister
	}
	return nil
}

// GetPrompt returns the prompt request if present
func (x *ConnectRequest) GetPrompt() *PromptRequest {
	if x, ok := x.GetRequest().(*ConnectRequest_Prompt); ok {
		return x.Prompt
	}
	return nil
}

// GetRequest returns the oneof request field
func (x *ConnectRequest) GetRequest() isConnectRequest_Request {
	if x != nil {
		return x.Request
	}
	return nil
}

// GetType returns the request type
func (x *ConnectRequest) GetType() RequestType {
	if x != nil {
		return x.Type
	}
	return RequestType_UNSPECIFIED
}

// GetTimestamp returns the timestamp
func (x *ConnectRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type ConnectResponse struct {
	State     protoimpl.MessageState `protogen:"open.v1"`
	Type      RequestType            `protobuf:"varint,1,opt,name=type,proto3,enum=routeguide.RequestType" json:"type,omitempty"`
	Timestamp int64                  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Types that are assignable to Response:
	//
	//	*ConnectResponse_Register
	//	*ConnectResponse_Unregister
	//	*ConnectResponse_Prompt
	//	*ConnectResponse_Error
	Response isConnectResponse_Response `protobuf_oneof:"response"`
}

type isConnectResponse_Response interface {
	isConnectResponse_Response()
}

type ConnectResponse_Register struct {
	Register *MessageResponse `protobuf:"bytes,3,opt,name=register,proto3,oneof"`
}

type ConnectResponse_Unregister struct {
	Unregister *UnregisterResponse `protobuf:"bytes,4,opt,name=unregister,proto3,oneof"`
}

type ConnectResponse_Prompt struct {
	Prompt *PromptResponse `protobuf:"bytes,5,opt,name=prompt,proto3,oneof"`
}

type ConnectResponse_Error struct {
	Error *ErrorResponse `protobuf:"bytes,6,opt,name=error,proto3,oneof"`
}

func (*ConnectResponse_Register) isConnectResponse_Response()   {}
func (*ConnectResponse_Unregister) isConnectResponse_Response() {}
func (*ConnectResponse_Prompt) isConnectResponse_Response()     {}
func (*ConnectResponse_Error) isConnectResponse_Response()      {}

// GetRegister returns the register response if present
func (x *ConnectResponse) GetRegister() *MessageResponse {
	if x, ok := x.GetResponse().(*ConnectResponse_Register); ok {
		return x.Register
	}
	return nil
}

// GetUnregister returns the unregister response if present
func (x *ConnectResponse) GetUnregister() *UnregisterResponse {
	if x, ok := x.GetResponse().(*ConnectResponse_Unregister); ok {
		return x.Unregister
	}
	return nil
}

// GetPrompt returns the prompt response if present
func (x *ConnectResponse) GetPrompt() *PromptResponse {
	if x, ok := x.GetResponse().(*ConnectResponse_Prompt); ok {
		return x.Prompt
	}
	return nil
}

// GetError returns the error response if present
func (x *ConnectResponse) GetError() *ErrorResponse {
	if x, ok := x.GetResponse().(*ConnectResponse_Error); ok {
		return x.Error
	}
	return nil
}

// GetResponse returns the oneof response field
func (x *ConnectResponse) GetResponse() isConnectResponse_Response {
	if x != nil {
		return x.Response
	}
	return nil
}

// GetType returns the response type
func (x *ConnectResponse) GetType() RequestType {
	if x != nil {
		return x.Type
	}
	return RequestType_UNSPECIFIED
}

// GetTimestamp returns the timestamp
func (x *ConnectResponse) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}
