// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: auth.proto

package auth

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

// example with same name
type Gender int32

const (
	Gender_GENDER_UNKNOWN Gender = 0
	Gender_GENDER_MALE    Gender = 1
	Gender_GENDER_FEMALE  Gender = 2
	Gender_GENDER_OTHER   Gender = 3
)

// Enum value maps for Gender.
var (
	Gender_name = map[int32]string{
		0: "GENDER_UNKNOWN",
		1: "GENDER_MALE",
		2: "GENDER_FEMALE",
		3: "GENDER_OTHER",
	}
	Gender_value = map[string]int32{
		"GENDER_UNKNOWN": 0,
		"GENDER_MALE":    1,
		"GENDER_FEMALE":  2,
		"GENDER_OTHER":   3,
	}
)

func (x Gender) Enum() *Gender {
	p := new(Gender)
	*p = x
	return p
}

func (x Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_auth_proto_enumTypes[0].Descriptor()
}

func (Gender) Type() protoreflect.EnumType {
	return &file_auth_proto_enumTypes[0]
}

func (x Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gender.Descriptor instead.
func (Gender) EnumDescriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

type UserRole int32

const (
	UserRole_USER_ROLE_UNKNOWN   UserRole = 0
	UserRole_USER_ROLE_ADMIN     UserRole = 1
	UserRole_USER_ROLE_USER      UserRole = 2
	UserRole_USER_ROLE_MODERATOR UserRole = 3
)

// Enum value maps for UserRole.
var (
	UserRole_name = map[int32]string{
		0: "USER_ROLE_UNKNOWN",
		1: "USER_ROLE_ADMIN",
		2: "USER_ROLE_USER",
		3: "USER_ROLE_MODERATOR",
	}
	UserRole_value = map[string]int32{
		"USER_ROLE_UNKNOWN":   0,
		"USER_ROLE_ADMIN":     1,
		"USER_ROLE_USER":      2,
		"USER_ROLE_MODERATOR": 3,
	}
)

func (x UserRole) Enum() *UserRole {
	p := new(UserRole)
	*p = x
	return p
}

func (x UserRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserRole) Descriptor() protoreflect.EnumDescriptor {
	return file_auth_proto_enumTypes[1].Descriptor()
}

func (UserRole) Type() protoreflect.EnumType {
	return &file_auth_proto_enumTypes[1]
}

func (x UserRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserRole.Descriptor instead.
func (UserRole) EnumDescriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{1}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string        `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name    string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Gender  Gender        `protobuf:"varint,3,opt,name=gender,proto3,enum=user.Gender" json:"gender,omitempty"`
	Role    UserRole      `protobuf:"varint,4,opt,name=role,proto3,enum=user.UserRole" json:"role,omitempty"`
	Email   string        `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Address *User_Address `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	// Types that are assignable to ContactMethod:
	//
	//	*User_PhoneNumber
	//	*User_EmailContact
	ContactMethod isUser_ContactMethod `protobuf_oneof:"contact_method"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetGender() Gender {
	if x != nil {
		return x.Gender
	}
	return Gender_GENDER_UNKNOWN
}

func (x *User) GetRole() UserRole {
	if x != nil {
		return x.Role
	}
	return UserRole_USER_ROLE_UNKNOWN
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetAddress() *User_Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (m *User) GetContactMethod() isUser_ContactMethod {
	if m != nil {
		return m.ContactMethod
	}
	return nil
}

func (x *User) GetPhoneNumber() string {
	if x, ok := x.GetContactMethod().(*User_PhoneNumber); ok {
		return x.PhoneNumber
	}
	return ""
}

func (x *User) GetEmailContact() string {
	if x, ok := x.GetContactMethod().(*User_EmailContact); ok {
		return x.EmailContact
	}
	return ""
}

type isUser_ContactMethod interface {
	isUser_ContactMethod()
}

type User_PhoneNumber struct {
	PhoneNumber string `protobuf:"bytes,7,opt,name=phone_number,json=phoneNumber,proto3,oneof"`
}

type User_EmailContact struct {
	EmailContact string `protobuf:"bytes,8,opt,name=email_contact,json=emailContact,proto3,oneof"`
}

func (*User_PhoneNumber) isUser_ContactMethod() {}

func (*User_EmailContact) isUser_ContactMethod() {}

// Request and Response messages for user registration
type RegisterUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User     *User  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RegisterUserRequest) Reset() {
	*x = RegisterUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUserRequest) ProtoMessage() {}

func (x *RegisterUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUserRequest.ProtoReflect.Descriptor instead.
func (*RegisterUserRequest) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterUserRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *RegisterUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	UserId  string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *RegisterUserResponse) Reset() {
	*x = RegisterUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUserResponse) ProtoMessage() {}

func (x *RegisterUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUserResponse.ProtoReflect.Descriptor instead.
func (*RegisterUserResponse) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterUserResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RegisterUserResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type LoginUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to LoginMethod:
	//
	//	*LoginUserRequest_Email
	//	*LoginUserRequest_PhoneNumber
	LoginMethod isLoginUserRequest_LoginMethod `protobuf_oneof:"login_method"`
	Password    string                         `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginUserRequest) Reset() {
	*x = LoginUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginUserRequest) ProtoMessage() {}

func (x *LoginUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginUserRequest.ProtoReflect.Descriptor instead.
func (*LoginUserRequest) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{3}
}

func (m *LoginUserRequest) GetLoginMethod() isLoginUserRequest_LoginMethod {
	if m != nil {
		return m.LoginMethod
	}
	return nil
}

func (x *LoginUserRequest) GetEmail() string {
	if x, ok := x.GetLoginMethod().(*LoginUserRequest_Email); ok {
		return x.Email
	}
	return ""
}

func (x *LoginUserRequest) GetPhoneNumber() string {
	if x, ok := x.GetLoginMethod().(*LoginUserRequest_PhoneNumber); ok {
		return x.PhoneNumber
	}
	return ""
}

func (x *LoginUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type isLoginUserRequest_LoginMethod interface {
	isLoginUserRequest_LoginMethod()
}

type LoginUserRequest_Email struct {
	Email string `protobuf:"bytes,1,opt,name=email,proto3,oneof"`
}

type LoginUserRequest_PhoneNumber struct {
	PhoneNumber string `protobuf:"bytes,2,opt,name=phone_number,json=phoneNumber,proto3,oneof"`
}

func (*LoginUserRequest_Email) isLoginUserRequest_LoginMethod() {}

func (*LoginUserRequest_PhoneNumber) isLoginUserRequest_LoginMethod() {}

type LoginUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LoginUserResponse) Reset() {
	*x = LoginUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginUserResponse) ProtoMessage() {}

func (x *LoginUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginUserResponse.ProtoReflect.Descriptor instead.
func (*LoginUserResponse) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{4}
}

func (x *LoginUserResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UserInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserInfoRequest) Reset() {
	*x = UserInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoRequest) ProtoMessage() {}

func (x *UserInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoRequest.ProtoReflect.Descriptor instead.
func (*UserInfoRequest) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{5}
}

type UserInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UserInfoResponse) Reset() {
	*x = UserInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoResponse) ProtoMessage() {}

func (x *UserInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoResponse.ProtoReflect.Descriptor instead.
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{6}
}

func (x *UserInfoResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type User_Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Street  string `protobuf:"bytes,1,opt,name=street,proto3" json:"street,omitempty"`
	City    string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	State   string `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Zipcode string `protobuf:"bytes,4,opt,name=zipcode,proto3" json:"zipcode,omitempty"`
}

func (x *User_Address) Reset() {
	*x = User_Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User_Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User_Address) ProtoMessage() {}

func (x *User_Address) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User_Address.ProtoReflect.Descriptor instead.
func (*User_Address) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0, 0}
}

func (x *User_Address) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *User_Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *User_Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *User_Address) GetZipcode() string {
	if x != nil {
		return x.Zipcode
	}
	return ""
}

var File_auth_proto protoreflect.FileDescriptor

var file_auth_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x22, 0x86, 0x03, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x22,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2c, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0d, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x1a, 0x65, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x7a, 0x69, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x7a, 0x69, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x51, 0x0a, 0x13, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x49,
	0x0a, 0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x7b, 0x0a, 0x10, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x23, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x29, 0x0a, 0x11, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x11, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x32, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x2a, 0x52, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x0e, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52,
	0x5f, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x47, 0x45, 0x4e, 0x44, 0x45,
	0x52, 0x5f, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x47, 0x45,
	0x4e, 0x44, 0x45, 0x52, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x03, 0x2a, 0x63, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x53, 0x45, 0x52,
	0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x13, 0x0a, 0x0f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x41, 0x44, 0x4d,
	0x49, 0x4e, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x52, 0x4f, 0x4c,
	0x45, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x55, 0x53, 0x45, 0x52,
	0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x10,
	0x03, 0x32, 0xcd, 0x01, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x45, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_auth_proto_rawDescOnce sync.Once
	file_auth_proto_rawDescData = file_auth_proto_rawDesc
)

func file_auth_proto_rawDescGZIP() []byte {
	file_auth_proto_rawDescOnce.Do(func() {
		file_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_proto_rawDescData)
	})
	return file_auth_proto_rawDescData
}

var file_auth_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_auth_proto_goTypes = []any{
	(Gender)(0),                  // 0: user.Gender
	(UserRole)(0),                // 1: user.UserRole
	(*User)(nil),                 // 2: user.User
	(*RegisterUserRequest)(nil),  // 3: user.RegisterUserRequest
	(*RegisterUserResponse)(nil), // 4: user.RegisterUserResponse
	(*LoginUserRequest)(nil),     // 5: user.LoginUserRequest
	(*LoginUserResponse)(nil),    // 6: user.LoginUserResponse
	(*UserInfoRequest)(nil),      // 7: user.UserInfoRequest
	(*UserInfoResponse)(nil),     // 8: user.UserInfoResponse
	(*User_Address)(nil),         // 9: user.User.Address
}
var file_auth_proto_depIdxs = []int32{
	0, // 0: user.User.gender:type_name -> user.Gender
	1, // 1: user.User.role:type_name -> user.UserRole
	9, // 2: user.User.address:type_name -> user.User.Address
	2, // 3: user.RegisterUserRequest.user:type_name -> user.User
	2, // 4: user.UserInfoResponse.user:type_name -> user.User
	3, // 5: user.AuthService.RegisterUser:input_type -> user.RegisterUserRequest
	5, // 6: user.AuthService.LoginUser:input_type -> user.LoginUserRequest
	7, // 7: user.AuthService.UserInfo:input_type -> user.UserInfoRequest
	4, // 8: user.AuthService.RegisterUser:output_type -> user.RegisterUserResponse
	6, // 9: user.AuthService.LoginUser:output_type -> user.LoginUserResponse
	8, // 10: user.AuthService.UserInfo:output_type -> user.UserInfoResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_auth_proto_init() }
func file_auth_proto_init() {
	if File_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*User); i {
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
		file_auth_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RegisterUserRequest); i {
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
		file_auth_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*RegisterUserResponse); i {
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
		file_auth_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*LoginUserRequest); i {
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
		file_auth_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*LoginUserResponse); i {
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
		file_auth_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UserInfoRequest); i {
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
		file_auth_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*UserInfoResponse); i {
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
		file_auth_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*User_Address); i {
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
	file_auth_proto_msgTypes[0].OneofWrappers = []any{
		(*User_PhoneNumber)(nil),
		(*User_EmailContact)(nil),
	}
	file_auth_proto_msgTypes[3].OneofWrappers = []any{
		(*LoginUserRequest_Email)(nil),
		(*LoginUserRequest_PhoneNumber)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_proto_goTypes,
		DependencyIndexes: file_auth_proto_depIdxs,
		EnumInfos:         file_auth_proto_enumTypes,
		MessageInfos:      file_auth_proto_msgTypes,
	}.Build()
	File_auth_proto = out.File
	file_auth_proto_rawDesc = nil
	file_auth_proto_goTypes = nil
	file_auth_proto_depIdxs = nil
}
