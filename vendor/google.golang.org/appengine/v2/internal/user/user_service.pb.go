// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: user_service.proto

package user

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

type UserServiceError_ErrorCode int32

const (
	UserServiceError_OK                    UserServiceError_ErrorCode = 0
	UserServiceError_REDIRECT_URL_TOO_LONG UserServiceError_ErrorCode = 1
	UserServiceError_NOT_ALLOWED           UserServiceError_ErrorCode = 2
	UserServiceError_OAUTH_INVALID_TOKEN   UserServiceError_ErrorCode = 3
	UserServiceError_OAUTH_INVALID_REQUEST UserServiceError_ErrorCode = 4
	UserServiceError_OAUTH_ERROR           UserServiceError_ErrorCode = 5
)

// Enum value maps for UserServiceError_ErrorCode.
var (
	UserServiceError_ErrorCode_name = map[int32]string{
		0: "OK",
		1: "REDIRECT_URL_TOO_LONG",
		2: "NOT_ALLOWED",
		3: "OAUTH_INVALID_TOKEN",
		4: "OAUTH_INVALID_REQUEST",
		5: "OAUTH_ERROR",
	}
	UserServiceError_ErrorCode_value = map[string]int32{
		"OK":                    0,
		"REDIRECT_URL_TOO_LONG": 1,
		"NOT_ALLOWED":           2,
		"OAUTH_INVALID_TOKEN":   3,
		"OAUTH_INVALID_REQUEST": 4,
		"OAUTH_ERROR":           5,
	}
)

func (x UserServiceError_ErrorCode) Enum() *UserServiceError_ErrorCode {
	p := new(UserServiceError_ErrorCode)
	*p = x
	return p
}

func (x UserServiceError_ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserServiceError_ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_user_service_proto_enumTypes[0].Descriptor()
}

func (UserServiceError_ErrorCode) Type() protoreflect.EnumType {
	return &file_user_service_proto_enumTypes[0]
}

func (x UserServiceError_ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *UserServiceError_ErrorCode) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = UserServiceError_ErrorCode(num)
	return nil
}

// Deprecated: Use UserServiceError_ErrorCode.Descriptor instead.
func (UserServiceError_ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{0, 0}
}

type UserServiceError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserServiceError) Reset() {
	*x = UserServiceError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserServiceError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserServiceError) ProtoMessage() {}

func (x *UserServiceError) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserServiceError.ProtoReflect.Descriptor instead.
func (*UserServiceError) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{0}
}

type CreateLoginURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DestinationUrl    *string `protobuf:"bytes,1,req,name=destination_url,json=destinationUrl" json:"destination_url,omitempty"`
	AuthDomain        *string `protobuf:"bytes,2,opt,name=auth_domain,json=authDomain" json:"auth_domain,omitempty"`
	FederatedIdentity *string `protobuf:"bytes,3,opt,name=federated_identity,json=federatedIdentity,def=" json:"federated_identity,omitempty"`
}

// Default values for CreateLoginURLRequest fields.
const (
	Default_CreateLoginURLRequest_FederatedIdentity = string("")
)

func (x *CreateLoginURLRequest) Reset() {
	*x = CreateLoginURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLoginURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLoginURLRequest) ProtoMessage() {}

func (x *CreateLoginURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLoginURLRequest.ProtoReflect.Descriptor instead.
func (*CreateLoginURLRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateLoginURLRequest) GetDestinationUrl() string {
	if x != nil && x.DestinationUrl != nil {
		return *x.DestinationUrl
	}
	return ""
}

func (x *CreateLoginURLRequest) GetAuthDomain() string {
	if x != nil && x.AuthDomain != nil {
		return *x.AuthDomain
	}
	return ""
}

func (x *CreateLoginURLRequest) GetFederatedIdentity() string {
	if x != nil && x.FederatedIdentity != nil {
		return *x.FederatedIdentity
	}
	return Default_CreateLoginURLRequest_FederatedIdentity
}

type CreateLoginURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoginUrl *string `protobuf:"bytes,1,req,name=login_url,json=loginUrl" json:"login_url,omitempty"`
}

func (x *CreateLoginURLResponse) Reset() {
	*x = CreateLoginURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLoginURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLoginURLResponse) ProtoMessage() {}

func (x *CreateLoginURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLoginURLResponse.ProtoReflect.Descriptor instead.
func (*CreateLoginURLResponse) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateLoginURLResponse) GetLoginUrl() string {
	if x != nil && x.LoginUrl != nil {
		return *x.LoginUrl
	}
	return ""
}

type CreateLogoutURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DestinationUrl *string `protobuf:"bytes,1,req,name=destination_url,json=destinationUrl" json:"destination_url,omitempty"`
	AuthDomain     *string `protobuf:"bytes,2,opt,name=auth_domain,json=authDomain" json:"auth_domain,omitempty"`
}

func (x *CreateLogoutURLRequest) Reset() {
	*x = CreateLogoutURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLogoutURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLogoutURLRequest) ProtoMessage() {}

func (x *CreateLogoutURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLogoutURLRequest.ProtoReflect.Descriptor instead.
func (*CreateLogoutURLRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateLogoutURLRequest) GetDestinationUrl() string {
	if x != nil && x.DestinationUrl != nil {
		return *x.DestinationUrl
	}
	return ""
}

func (x *CreateLogoutURLRequest) GetAuthDomain() string {
	if x != nil && x.AuthDomain != nil {
		return *x.AuthDomain
	}
	return ""
}

type CreateLogoutURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogoutUrl *string `protobuf:"bytes,1,req,name=logout_url,json=logoutUrl" json:"logout_url,omitempty"`
}

func (x *CreateLogoutURLResponse) Reset() {
	*x = CreateLogoutURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLogoutURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLogoutURLResponse) ProtoMessage() {}

func (x *CreateLogoutURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLogoutURLResponse.ProtoReflect.Descriptor instead.
func (*CreateLogoutURLResponse) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateLogoutURLResponse) GetLogoutUrl() string {
	if x != nil && x.LogoutUrl != nil {
		return *x.LogoutUrl
	}
	return ""
}

type GetOAuthUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scope  *string  `protobuf:"bytes,1,opt,name=scope" json:"scope,omitempty"`
	Scopes []string `protobuf:"bytes,2,rep,name=scopes" json:"scopes,omitempty"`
}

func (x *GetOAuthUserRequest) Reset() {
	*x = GetOAuthUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOAuthUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOAuthUserRequest) ProtoMessage() {}

func (x *GetOAuthUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOAuthUserRequest.ProtoReflect.Descriptor instead.
func (*GetOAuthUserRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetOAuthUserRequest) GetScope() string {
	if x != nil && x.Scope != nil {
		return *x.Scope
	}
	return ""
}

func (x *GetOAuthUserRequest) GetScopes() []string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

type GetOAuthUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email            *string  `protobuf:"bytes,1,req,name=email" json:"email,omitempty"`
	UserId           *string  `protobuf:"bytes,2,req,name=user_id,json=userId" json:"user_id,omitempty"`
	AuthDomain       *string  `protobuf:"bytes,3,req,name=auth_domain,json=authDomain" json:"auth_domain,omitempty"`
	UserOrganization *string  `protobuf:"bytes,4,opt,name=user_organization,json=userOrganization,def=" json:"user_organization,omitempty"`
	IsAdmin          *bool    `protobuf:"varint,5,opt,name=is_admin,json=isAdmin,def=0" json:"is_admin,omitempty"`
	ClientId         *string  `protobuf:"bytes,6,opt,name=client_id,json=clientId,def=" json:"client_id,omitempty"`
	Scopes           []string `protobuf:"bytes,7,rep,name=scopes" json:"scopes,omitempty"`
}

// Default values for GetOAuthUserResponse fields.
const (
	Default_GetOAuthUserResponse_UserOrganization = string("")
	Default_GetOAuthUserResponse_IsAdmin          = bool(false)
	Default_GetOAuthUserResponse_ClientId         = string("")
)

func (x *GetOAuthUserResponse) Reset() {
	*x = GetOAuthUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOAuthUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOAuthUserResponse) ProtoMessage() {}

func (x *GetOAuthUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOAuthUserResponse.ProtoReflect.Descriptor instead.
func (*GetOAuthUserResponse) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetOAuthUserResponse) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *GetOAuthUserResponse) GetUserId() string {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return ""
}

func (x *GetOAuthUserResponse) GetAuthDomain() string {
	if x != nil && x.AuthDomain != nil {
		return *x.AuthDomain
	}
	return ""
}

func (x *GetOAuthUserResponse) GetUserOrganization() string {
	if x != nil && x.UserOrganization != nil {
		return *x.UserOrganization
	}
	return Default_GetOAuthUserResponse_UserOrganization
}

func (x *GetOAuthUserResponse) GetIsAdmin() bool {
	if x != nil && x.IsAdmin != nil {
		return *x.IsAdmin
	}
	return Default_GetOAuthUserResponse_IsAdmin
}

func (x *GetOAuthUserResponse) GetClientId() string {
	if x != nil && x.ClientId != nil {
		return *x.ClientId
	}
	return Default_GetOAuthUserResponse_ClientId
}

func (x *GetOAuthUserResponse) GetScopes() []string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

type CheckOAuthSignatureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CheckOAuthSignatureRequest) Reset() {
	*x = CheckOAuthSignatureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckOAuthSignatureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckOAuthSignatureRequest) ProtoMessage() {}

func (x *CheckOAuthSignatureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckOAuthSignatureRequest.ProtoReflect.Descriptor instead.
func (*CheckOAuthSignatureRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{7}
}

type CheckOAuthSignatureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OauthConsumerKey *string `protobuf:"bytes,1,req,name=oauth_consumer_key,json=oauthConsumerKey" json:"oauth_consumer_key,omitempty"`
}

func (x *CheckOAuthSignatureResponse) Reset() {
	*x = CheckOAuthSignatureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckOAuthSignatureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckOAuthSignatureResponse) ProtoMessage() {}

func (x *CheckOAuthSignatureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckOAuthSignatureResponse.ProtoReflect.Descriptor instead.
func (*CheckOAuthSignatureResponse) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{8}
}

func (x *CheckOAuthSignatureResponse) GetOauthConsumerKey() string {
	if x != nil && x.OauthConsumerKey != nil {
		return *x.OauthConsumerKey
	}
	return ""
}

var File_user_service_proto protoreflect.FileDescriptor

var file_user_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x76, 0x32, 0x22, 0x99, 0x01, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x84, 0x01, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x19, 0x0a,
	0x15, 0x52, 0x45, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x54, 0x4f,
	0x4f, 0x5f, 0x4c, 0x4f, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x4f, 0x54, 0x5f,
	0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x41, 0x55,
	0x54, 0x48, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e,
	0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x4f, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x04, 0x12, 0x0f, 0x0a,
	0x0b, 0x4f, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x05, 0x22, 0x92,
	0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x52,
	0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x09, 0x52, 0x0e, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x72,
	0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x12, 0x2f, 0x0a, 0x12, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x3a, 0x00,
	0x52, 0x11, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x22, 0x35, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x72, 0x6c, 0x22, 0x62, 0x0a, 0x16, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x0e, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x1f, 0x0a,
	0x0b, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x38,
	0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x55, 0x52,
	0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x67,
	0x6f, 0x75, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x09, 0x6c,
	0x6f, 0x67, 0x6f, 0x75, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x43, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4f,
	0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x22, 0xee, 0x01,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x02, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x2d, 0x0a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x3a, 0x00, 0x52, 0x10, 0x75, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x07,
	0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1d, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x3a, 0x00, 0x52, 0x08, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x22, 0x1c,
	0x0a, 0x1a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4b, 0x0a, 0x1b,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x6f,
	0x61, 0x75, 0x74, 0x68, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x10, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x43, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x61,
	0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72,
}

var (
	file_user_service_proto_rawDescOnce sync.Once
	file_user_service_proto_rawDescData = file_user_service_proto_rawDesc
)

func file_user_service_proto_rawDescGZIP() []byte {
	file_user_service_proto_rawDescOnce.Do(func() {
		file_user_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_service_proto_rawDescData)
	})
	return file_user_service_proto_rawDescData
}

var file_user_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_user_service_proto_goTypes = []interface{}{
	(UserServiceError_ErrorCode)(0),     // 0: appengine.v2.UserServiceError.ErrorCode
	(*UserServiceError)(nil),            // 1: appengine.v2.UserServiceError
	(*CreateLoginURLRequest)(nil),       // 2: appengine.v2.CreateLoginURLRequest
	(*CreateLoginURLResponse)(nil),      // 3: appengine.v2.CreateLoginURLResponse
	(*CreateLogoutURLRequest)(nil),      // 4: appengine.v2.CreateLogoutURLRequest
	(*CreateLogoutURLResponse)(nil),     // 5: appengine.v2.CreateLogoutURLResponse
	(*GetOAuthUserRequest)(nil),         // 6: appengine.v2.GetOAuthUserRequest
	(*GetOAuthUserResponse)(nil),        // 7: appengine.v2.GetOAuthUserResponse
	(*CheckOAuthSignatureRequest)(nil),  // 8: appengine.v2.CheckOAuthSignatureRequest
	(*CheckOAuthSignatureResponse)(nil), // 9: appengine.v2.CheckOAuthSignatureResponse
}
var file_user_service_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_service_proto_init() }
func file_user_service_proto_init() {
	if File_user_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserServiceError); i {
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
		file_user_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLoginURLRequest); i {
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
		file_user_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLoginURLResponse); i {
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
		file_user_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLogoutURLRequest); i {
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
		file_user_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLogoutURLResponse); i {
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
		file_user_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOAuthUserRequest); i {
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
		file_user_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOAuthUserResponse); i {
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
		file_user_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckOAuthSignatureRequest); i {
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
		file_user_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckOAuthSignatureResponse); i {
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
			RawDescriptor: file_user_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_service_proto_goTypes,
		DependencyIndexes: file_user_service_proto_depIdxs,
		EnumInfos:         file_user_service_proto_enumTypes,
		MessageInfos:      file_user_service_proto_msgTypes,
	}.Build()
	File_user_service_proto = out.File
	file_user_service_proto_rawDesc = nil
	file_user_service_proto_goTypes = nil
	file_user_service_proto_depIdxs = nil
}