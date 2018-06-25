// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package testproto // import "github.com/mennanov/fieldmask-utils/testproto"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"
import field_mask "google.golang.org/genproto/protobuf/field_mask"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Role int32

const (
	Role_UNKNOWN Role = 0
	Role_REGULAR Role = 1
	Role_ADMIN   Role = 2
)

var Role_name = map[int32]string{
	0: "UNKNOWN",
	1: "REGULAR",
	2: "ADMIN",
}
var Role_value = map[string]int32{
	"UNKNOWN": 0,
	"REGULAR": 1,
	"ADMIN":   2,
}

func (x Role) String() string {
	return proto.EnumName(Role_name, int32(x))
}
func (Role) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_test_5498107a3fa265dd, []int{0}
}

type Permission int32

const (
	Permission_READ    Permission = 0
	Permission_WRITE   Permission = 1
	Permission_EXECUTE Permission = 2
)

var Permission_name = map[int32]string{
	0: "READ",
	1: "WRITE",
	2: "EXECUTE",
}
var Permission_value = map[string]int32{
	"READ":    0,
	"WRITE":   1,
	"EXECUTE": 2,
}

func (x Permission) String() string {
	return proto.EnumName(Permission_name, int32(x))
}
func (Permission) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_test_5498107a3fa265dd, []int{1}
}

type Image struct {
	OriginalUrl          string   `protobuf:"bytes,1,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
	ResizedUrl           string   `protobuf:"bytes,2,opt,name=resized_url,json=resizedUrl,proto3" json:"resized_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_5498107a3fa265dd, []int{0}
}
func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (dst *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(dst, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetOriginalUrl() string {
	if m != nil {
		return m.OriginalUrl
	}
	return ""
}

func (m *Image) GetResizedUrl() string {
	if m != nil {
		return m.ResizedUrl
	}
	return ""
}

type Metrics struct {
	Height               uint32   `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Weight               uint32   `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metrics) Reset()         { *m = Metrics{} }
func (m *Metrics) String() string { return proto.CompactTextString(m) }
func (*Metrics) ProtoMessage()    {}
func (*Metrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_5498107a3fa265dd, []int{1}
}
func (m *Metrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metrics.Unmarshal(m, b)
}
func (m *Metrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metrics.Marshal(b, m, deterministic)
}
func (dst *Metrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metrics.Merge(dst, src)
}
func (m *Metrics) XXX_Size() int {
	return xxx_messageInfo_Metrics.Size(m)
}
func (m *Metrics) XXX_DiscardUnknown() {
	xxx_messageInfo_Metrics.DiscardUnknown(m)
}

var xxx_messageInfo_Metrics proto.InternalMessageInfo

func (m *Metrics) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Metrics) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

type User struct {
	Id          uint32            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username    string            `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Role        Role              `protobuf:"varint,3,opt,name=role,proto3,enum=Role" json:"role,omitempty"`
	Meta        map[string]string `protobuf:"bytes,4,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Deactivated bool              `protobuf:"varint,5,opt,name=deactivated,proto3" json:"deactivated,omitempty"`
	Permissions []Permission      `protobuf:"varint,6,rep,packed,name=permissions,proto3,enum=Permission" json:"permissions,omitempty"`
	// Types that are valid to be assigned to Name:
	//	*User_MaleName
	//	*User_FemaleName
	Name                 isUser_Name `protobuf_oneof:"name"`
	Details              []*any.Any  `protobuf:"bytes,9,rep,name=details,proto3" json:"details,omitempty"`
	Images               []*Image    `protobuf:"bytes,10,rep,name=images,proto3" json:"images,omitempty"`
	Avatar               *Image      `protobuf:"bytes,11,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Tags                 []string    `protobuf:"bytes,12,rep,name=tags,proto3" json:"tags,omitempty"`
	Friends              []*User     `protobuf:"bytes,13,rep,name=friends,proto3" json:"friends,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_5498107a3fa265dd, []int{2}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

type isUser_Name interface {
	isUser_Name()
}

type User_MaleName struct {
	MaleName string `protobuf:"bytes,7,opt,name=male_name,json=maleName,proto3,oneof"`
}
type User_FemaleName struct {
	FemaleName string `protobuf:"bytes,8,opt,name=female_name,json=femaleName,proto3,oneof"`
}

func (*User_MaleName) isUser_Name()   {}
func (*User_FemaleName) isUser_Name() {}

func (m *User) GetName() isUser_Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *User) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetRole() Role {
	if m != nil {
		return m.Role
	}
	return Role_UNKNOWN
}

func (m *User) GetMeta() map[string]string {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *User) GetDeactivated() bool {
	if m != nil {
		return m.Deactivated
	}
	return false
}

func (m *User) GetPermissions() []Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *User) GetMaleName() string {
	if x, ok := m.GetName().(*User_MaleName); ok {
		return x.MaleName
	}
	return ""
}

func (m *User) GetFemaleName() string {
	if x, ok := m.GetName().(*User_FemaleName); ok {
		return x.FemaleName
	}
	return ""
}

func (m *User) GetDetails() []*any.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *User) GetImages() []*Image {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *User) GetAvatar() *Image {
	if m != nil {
		return m.Avatar
	}
	return nil
}

func (m *User) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *User) GetFriends() []*User {
	if m != nil {
		return m.Friends
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*User) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _User_OneofMarshaler, _User_OneofUnmarshaler, _User_OneofSizer, []interface{}{
		(*User_MaleName)(nil),
		(*User_FemaleName)(nil),
	}
}

func _User_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*User)
	// name
	switch x := m.Name.(type) {
	case *User_MaleName:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.MaleName)
	case *User_FemaleName:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.FemaleName)
	case nil:
	default:
		return fmt.Errorf("User.Name has unexpected type %T", x)
	}
	return nil
}

func _User_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*User)
	switch tag {
	case 7: // name.male_name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Name = &User_MaleName{x}
		return true, err
	case 8: // name.female_name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Name = &User_FemaleName{x}
		return true, err
	default:
		return false, nil
	}
}

func _User_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*User)
	// name
	switch x := m.Name.(type) {
	case *User_MaleName:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.MaleName)))
		n += len(x.MaleName)
	case *User_FemaleName:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.FemaleName)))
		n += len(x.FemaleName)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type UpdateUserRequest struct {
	User                 *User                 `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	FieldMask            *field_mask.FieldMask `protobuf:"bytes,2,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateUserRequest) Reset()         { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()    {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_5498107a3fa265dd, []int{3}
}
func (m *UpdateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserRequest.Unmarshal(m, b)
}
func (m *UpdateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserRequest.Merge(dst, src)
}
func (m *UpdateUserRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserRequest.Size(m)
}
func (m *UpdateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserRequest proto.InternalMessageInfo

func (m *UpdateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *UpdateUserRequest) GetFieldMask() *field_mask.FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func init() {
	proto.RegisterType((*Image)(nil), "Image")
	proto.RegisterType((*Metrics)(nil), "Metrics")
	proto.RegisterType((*User)(nil), "User")
	proto.RegisterMapType((map[string]string)(nil), "User.MetaEntry")
	proto.RegisterType((*UpdateUserRequest)(nil), "UpdateUserRequest")
	proto.RegisterEnum("Role", Role_name, Role_value)
	proto.RegisterEnum("Permission", Permission_name, Permission_value)
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_test_5498107a3fa265dd) }

var fileDescriptor_test_5498107a3fa265dd = []byte{
	// 593 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x93, 0x5d, 0x6f, 0xd3, 0x3c,
	0x14, 0xc7, 0x97, 0x36, 0x7d, 0xc9, 0xc9, 0xb6, 0x27, 0x8f, 0x35, 0xa1, 0xac, 0x12, 0x2c, 0x2b,
	0x37, 0xd5, 0xd0, 0x52, 0xa9, 0x5c, 0xc0, 0xb8, 0xeb, 0x58, 0x80, 0x69, 0xb4, 0x20, 0x6b, 0xd1,
	0x10, 0x37, 0x93, 0xb7, 0x9c, 0x66, 0xd6, 0xf2, 0x32, 0x6c, 0xa7, 0xa8, 0x7c, 0x65, 0xbe, 0x04,
	0xb2, 0x93, 0x76, 0x13, 0x5c, 0xc5, 0xe7, 0x77, 0xfe, 0x3e, 0x3e, 0xf1, 0xff, 0x18, 0x40, 0xa1,
	0x54, 0xe1, 0x83, 0x28, 0x55, 0x39, 0xd8, 0x4f, 0xcb, 0x32, 0xcd, 0x70, 0x6c, 0xa2, 0x9b, 0x6a,
	0x31, 0x66, 0xc5, 0xaa, 0x49, 0x05, 0x7f, 0xa7, 0x16, 0x1c, 0xb3, 0xe4, 0x3a, 0x67, 0xf2, 0xbe,
	0x56, 0x0c, 0x2f, 0xa0, 0x73, 0x9e, 0xb3, 0x14, 0xc9, 0x21, 0x6c, 0x97, 0x82, 0xa7, 0xbc, 0x60,
	0xd9, 0x75, 0x25, 0x32, 0xdf, 0x0a, 0xac, 0x91, 0x43, 0xdd, 0x35, 0x8b, 0x45, 0x46, 0x0e, 0xc0,
	0x15, 0x28, 0xf9, 0x2f, 0x4c, 0x8c, 0xa2, 0x65, 0x14, 0xd0, 0xa0, 0x58, 0x64, 0xc3, 0x13, 0xe8,
	0xcd, 0x50, 0x09, 0x7e, 0x2b, 0xc9, 0x33, 0xe8, 0xde, 0x21, 0x4f, 0xef, 0x94, 0x29, 0xb4, 0x43,
	0x9b, 0x48, 0xf3, 0x9f, 0x35, 0x6f, 0xd5, 0xbc, 0x8e, 0x86, 0xbf, 0xdb, 0x60, 0xc7, 0x12, 0x05,
	0xd9, 0x85, 0x16, 0x4f, 0x9a, 0x4d, 0x2d, 0x9e, 0x90, 0x01, 0xf4, 0x2b, 0x89, 0xa2, 0x60, 0x39,
	0x36, 0x27, 0x6e, 0x62, 0xb2, 0x0f, 0xb6, 0x28, 0x33, 0xf4, 0xdb, 0x81, 0x35, 0xda, 0x9d, 0x74,
	0x42, 0x5a, 0x66, 0x48, 0x0d, 0x22, 0x2f, 0xc1, 0xce, 0x51, 0x31, 0xdf, 0x0e, 0xda, 0x23, 0x77,
	0xf2, 0x5f, 0xa8, 0x6b, 0x87, 0x33, 0x54, 0x2c, 0x2a, 0x94, 0x58, 0x51, 0x93, 0x24, 0x01, 0xb8,
	0x09, 0xb2, 0x5b, 0xc5, 0x97, 0x4c, 0x61, 0xe2, 0x77, 0x02, 0x6b, 0xd4, 0xa7, 0x4f, 0x11, 0x39,
	0x06, 0xf7, 0x01, 0x45, 0xce, 0xa5, 0xe4, 0x65, 0x21, 0xfd, 0x6e, 0xd0, 0x1e, 0xed, 0x4e, 0xdc,
	0xf0, 0xeb, 0x86, 0xd1, 0xa7, 0x79, 0xf2, 0x1c, 0x9c, 0x9c, 0x65, 0x78, 0x6d, 0xba, 0xed, 0xe9,
	0x6e, 0x3f, 0x6d, 0xd1, 0xbe, 0x46, 0x73, 0xdd, 0xef, 0x21, 0xb8, 0x0b, 0x7c, 0x14, 0xf4, 0x1b,
	0x01, 0xd4, 0xd0, 0x48, 0x42, 0xe8, 0x25, 0xa8, 0x18, 0xcf, 0xa4, 0xef, 0x98, 0xd6, 0xf7, 0xc2,
	0xda, 0xc3, 0x70, 0xed, 0x61, 0x38, 0x2d, 0x56, 0x74, 0x2d, 0x22, 0x2f, 0xa0, 0xcb, 0xb5, 0x7f,
	0xd2, 0x07, 0x23, 0xef, 0x86, 0xc6, 0x4e, 0xda, 0x50, 0x9d, 0x67, 0x4b, 0xa6, 0x98, 0xf0, 0xdd,
	0xc0, 0x7a, 0x9a, 0xaf, 0x29, 0x21, 0x60, 0x2b, 0x96, 0x4a, 0x7f, 0x3b, 0x68, 0x8f, 0x1c, 0x6a,
	0xd6, 0xe4, 0x00, 0x7a, 0x0b, 0xc1, 0xb1, 0x48, 0xa4, 0xbf, 0x63, 0x8a, 0x76, 0xcc, 0xf5, 0xd1,
	0x35, 0x1d, 0xbc, 0x01, 0x67, 0x73, 0x95, 0xc4, 0x83, 0xf6, 0x3d, 0xae, 0x9a, 0x79, 0xd1, 0x4b,
	0xb2, 0x07, 0x9d, 0x25, 0xcb, 0xaa, 0xb5, 0x5f, 0x75, 0xf0, 0xae, 0xf5, 0xd6, 0x3a, 0xed, 0x82,
	0xad, 0xff, 0x7c, 0xc8, 0xe1, 0xff, 0xf8, 0x21, 0x61, 0x0a, 0x4d, 0x5d, 0xfc, 0x51, 0xa1, 0x54,
	0xda, 0x4d, 0xed, 0xac, 0xa9, 0xb4, 0x39, 0xd3, 0x20, 0x72, 0x02, 0xf0, 0x38, 0xb9, 0xa6, 0xac,
	0x3b, 0x19, 0xfc, 0x73, 0x31, 0x1f, 0xb4, 0x64, 0xc6, 0xe4, 0x3d, 0x75, 0x16, 0xeb, 0xe5, 0xd1,
	0x2b, 0xb0, 0xf5, 0x58, 0x10, 0x17, 0x7a, 0xf1, 0xfc, 0x62, 0xfe, 0xe5, 0x6a, 0xee, 0x6d, 0xe9,
	0x80, 0x46, 0x1f, 0xe3, 0xcf, 0x53, 0xea, 0x59, 0xc4, 0x81, 0xce, 0xf4, 0x6c, 0x76, 0x3e, 0xf7,
	0x5a, 0x47, 0x21, 0xc0, 0xa3, 0xb5, 0xa4, 0x0f, 0x36, 0x8d, 0xa6, 0x67, 0xde, 0x96, 0x96, 0x5c,
	0xd1, 0xf3, 0xcb, 0xc8, 0xb3, 0xf4, 0xd6, 0xe8, 0x5b, 0xf4, 0x3e, 0xbe, 0x8c, 0xbc, 0xd6, 0xe9,
	0xf8, 0xfb, 0x71, 0xca, 0xd5, 0x5d, 0x75, 0x13, 0xde, 0x96, 0xf9, 0x38, 0xc7, 0xa2, 0x60, 0x45,
	0xb9, 0xac, 0x5f, 0x99, 0x6e, 0xf5, 0xb8, 0x52, 0x3c, 0x93, 0x63, 0xfd, 0x58, 0xeb, 0x2e, 0xbb,
	0xe6, 0xf3, 0xfa, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x20, 0x4e, 0xbb, 0xc0, 0x03, 0x00,
	0x00,
}