// Code generated by protoc-gen-go. DO NOT EDIT.
// source: surface/surface.proto

package surface_v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type FieldKind int32

const (
	FieldKind_SCALAR    FieldKind = 0
	FieldKind_MAP       FieldKind = 1
	FieldKind_ARRAY     FieldKind = 2
	FieldKind_REFERENCE FieldKind = 3
	FieldKind_ANY       FieldKind = 4
)

var FieldKind_name = map[int32]string{
	0: "SCALAR",
	1: "MAP",
	2: "ARRAY",
	3: "REFERENCE",
	4: "ANY",
}

var FieldKind_value = map[string]int32{
	"SCALAR":    0,
	"MAP":       1,
	"ARRAY":     2,
	"REFERENCE": 3,
	"ANY":       4,
}

func (x FieldKind) String() string {
	return proto.EnumName(FieldKind_name, int32(x))
}

func (FieldKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_68892468a0db0ac7, []int{0}
}

type TypeKind int32

const (
	TypeKind_STRUCT TypeKind = 0
	TypeKind_OBJECT TypeKind = 1
)

var TypeKind_name = map[int32]string{
	0: "STRUCT",
	1: "OBJECT",
}

var TypeKind_value = map[string]int32{
	"STRUCT": 0,
	"OBJECT": 1,
}

func (x TypeKind) String() string {
	return proto.EnumName(TypeKind_name, int32(x))
}

func (TypeKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_68892468a0db0ac7, []int{1}
}

type Position int32

const (
	Position_BODY     Position = 0
	Position_HEADER   Position = 1
	Position_FORMDATA Position = 2
	Position_QUERY    Position = 3
	Position_PATH     Position = 4
)

var Position_name = map[int32]string{
	0: "BODY",
	1: "HEADER",
	2: "FORMDATA",
	3: "QUERY",
	4: "PATH",
}

var Position_value = map[string]int32{
	"BODY":     0,
	"HEADER":   1,
	"FORMDATA": 2,
	"QUERY":    3,
	"PATH":     4,
}

func (x Position) String() string {
	return proto.EnumName(Position_name, int32(x))
}

func (Position) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_68892468a0db0ac7, []int{2}
}

// Field is a field in a definition and can be associated with
// a position in a request structure.
type Field struct {
	Name                 string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string    `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Kind                 FieldKind `protobuf:"varint,3,opt,name=kind,proto3,enum=surface.v1.FieldKind" json:"kind,omitempty"`
	Format               string    `protobuf:"bytes,4,opt,name=format,proto3" json:"format,omitempty"`
	Position             Position  `protobuf:"varint,5,opt,name=position,proto3,enum=surface.v1.Position" json:"position,omitempty"`
	NativeType           string    `protobuf:"bytes,6,opt,name=nativeType,proto3" json:"nativeType,omitempty"`
	FieldName            string    `protobuf:"bytes,7,opt,name=fieldName,proto3" json:"fieldName,omitempty"`
	ParameterName        string    `protobuf:"bytes,8,opt,name=parameterName,proto3" json:"parameterName,omitempty"`
	Serialize            bool      `protobuf:"varint,9,opt,name=serialize,proto3" json:"serialize,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Field) Reset()         { *m = Field{} }
func (m *Field) String() string { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()    {}
func (*Field) Descriptor() ([]byte, []int) {
	return fileDescriptor_68892468a0db0ac7, []int{0}
}

func (m *Field) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Field.Unmarshal(m, b)
}
func (m *Field) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Field.Marshal(b, m, deterministic)
}
func (m *Field) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Field.Merge(m, src)
}
func (m *Field) XXX_Size() int {
	return xxx_messageInfo_Field.Size(m)
}
func (m *Field) XXX_DiscardUnknown() {
	xxx_messageInfo_Field.DiscardUnknown(m)
}

var xxx_messageInfo_Field proto.InternalMessageInfo

func (m *Field) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Field) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Field) GetKind() FieldKind {
	if m != nil {
		return m.Kind
	}
	return FieldKind_SCALAR
}

func (m *Field) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *Field) GetPosition() Position {
	if m != nil {
		return m.Position
	}
	return Position_BODY
}

func (m *Field) GetNativeType() string {
	if m != nil {
		return m.NativeType
	}
	return ""
}

func (m *Field) GetFieldName() string {
	if m != nil {
		return m.FieldName
	}
	return ""
}

func (m *Field) GetParameterName() string {
	if m != nil {
		return m.ParameterName
	}
	return ""
}

func (m *Field) GetSerialize() bool {
	if m != nil {
		return m.Serialize
	}
	return false
}

// Type typically corresponds to a definition, parameter, or response
// in an API and is represented by a type in generated code.
type Type struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Kind                 TypeKind `protobuf:"varint,2,opt,name=kind,proto3,enum=surface.v1.TypeKind" json:"kind,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	ContentType          string   `protobuf:"bytes,4,opt,name=contentType,proto3" json:"contentType,omitempty"`
	Fields               []*Field `protobuf:"bytes,5,rep,name=fields,proto3" json:"fields,omitempty"`
	TypeName             string   `protobuf:"bytes,6,opt,name=typeName,proto3" json:"typeName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Type) Reset()         { *m = Type{} }
func (m *Type) String() string { return proto.CompactTextString(m) }
func (*Type) ProtoMessage()    {}
func (*Type) Descriptor() ([]byte, []int) {
	return fileDescriptor_68892468a0db0ac7, []int{1}
}

func (m *Type) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Type.Unmarshal(m, b)
}
func (m *Type) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Type.Marshal(b, m, deterministic)
}
func (m *Type) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Type.Merge(m, src)
}
func (m *Type) XXX_Size() int {
	return xxx_messageInfo_Type.Size(m)
}
func (m *Type) XXX_DiscardUnknown() {
	xxx_messageInfo_Type.DiscardUnknown(m)
}

var xxx_messageInfo_Type proto.InternalMessageInfo

func (m *Type) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Type) GetKind() TypeKind {
	if m != nil {
		return m.Kind
	}
	return TypeKind_STRUCT
}

func (m *Type) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Type) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *Type) GetFields() []*Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *Type) GetTypeName() string {
	if m != nil {
		return m.TypeName
	}
	return ""
}

// Method is an operation of an API and typically has associated client and server code.
type Method struct {
	Operation            string   `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Method               string   `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Name                 string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	HandlerName          string   `protobuf:"bytes,6,opt,name=handlerName,proto3" json:"handlerName,omitempty"`
	ProcessorName        string   `protobuf:"bytes,7,opt,name=processorName,proto3" json:"processorName,omitempty"`
	ClientName           string   `protobuf:"bytes,8,opt,name=clientName,proto3" json:"clientName,omitempty"`
	ParametersTypeName   string   `protobuf:"bytes,9,opt,name=parametersTypeName,proto3" json:"parametersTypeName,omitempty"`
	ResponsesTypeName    string   `protobuf:"bytes,10,opt,name=responsesTypeName,proto3" json:"responsesTypeName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Method) Reset()         { *m = Method{} }
func (m *Method) String() string { return proto.CompactTextString(m) }
func (*Method) ProtoMessage()    {}
func (*Method) Descriptor() ([]byte, []int) {
	return fileDescriptor_68892468a0db0ac7, []int{2}
}

func (m *Method) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Method.Unmarshal(m, b)
}
func (m *Method) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Method.Marshal(b, m, deterministic)
}
func (m *Method) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Method.Merge(m, src)
}
func (m *Method) XXX_Size() int {
	return xxx_messageInfo_Method.Size(m)
}
func (m *Method) XXX_DiscardUnknown() {
	xxx_messageInfo_Method.DiscardUnknown(m)
}

var xxx_messageInfo_Method proto.InternalMessageInfo

func (m *Method) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func (m *Method) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Method) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Method) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Method) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Method) GetHandlerName() string {
	if m != nil {
		return m.HandlerName
	}
	return ""
}

func (m *Method) GetProcessorName() string {
	if m != nil {
		return m.ProcessorName
	}
	return ""
}

func (m *Method) GetClientName() string {
	if m != nil {
		return m.ClientName
	}
	return ""
}

func (m *Method) GetParametersTypeName() string {
	if m != nil {
		return m.ParametersTypeName
	}
	return ""
}

func (m *Method) GetResponsesTypeName() string {
	if m != nil {
		return m.ResponsesTypeName
	}
	return ""
}

// Model represents an API for code generation.
type Model struct {
	Name                 string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Types                []*Type   `protobuf:"bytes,2,rep,name=types,proto3" json:"types,omitempty"`
	Methods              []*Method `protobuf:"bytes,3,rep,name=methods,proto3" json:"methods,omitempty"`
	SymbolicReferences   []string  `protobuf:"bytes,4,rep,name=symbolic_references,json=symbolicReferences,proto3" json:"symbolic_references,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Model) Reset()         { *m = Model{} }
func (m *Model) String() string { return proto.CompactTextString(m) }
func (*Model) ProtoMessage()    {}
func (*Model) Descriptor() ([]byte, []int) {
	return fileDescriptor_68892468a0db0ac7, []int{3}
}

func (m *Model) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Model.Unmarshal(m, b)
}
func (m *Model) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Model.Marshal(b, m, deterministic)
}
func (m *Model) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Model.Merge(m, src)
}
func (m *Model) XXX_Size() int {
	return xxx_messageInfo_Model.Size(m)
}
func (m *Model) XXX_DiscardUnknown() {
	xxx_messageInfo_Model.DiscardUnknown(m)
}

var xxx_messageInfo_Model proto.InternalMessageInfo

func (m *Model) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Model) GetTypes() []*Type {
	if m != nil {
		return m.Types
	}
	return nil
}

func (m *Model) GetMethods() []*Method {
	if m != nil {
		return m.Methods
	}
	return nil
}

func (m *Model) GetSymbolicReferences() []string {
	if m != nil {
		return m.SymbolicReferences
	}
	return nil
}

func init() {
	proto.RegisterEnum("surface.v1.FieldKind", FieldKind_name, FieldKind_value)
	proto.RegisterEnum("surface.v1.TypeKind", TypeKind_name, TypeKind_value)
	proto.RegisterEnum("surface.v1.Position", Position_name, Position_value)
	proto.RegisterType((*Field)(nil), "surface.v1.Field")
	proto.RegisterType((*Type)(nil), "surface.v1.Type")
	proto.RegisterType((*Method)(nil), "surface.v1.Method")
	proto.RegisterType((*Model)(nil), "surface.v1.Model")
}

func init() { proto.RegisterFile("surface/surface.proto", fileDescriptor_68892468a0db0ac7) }

var fileDescriptor_68892468a0db0ac7 = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0xef, 0x6e, 0xd3, 0x3e,
	0x14, 0x5d, 0xfe, 0x75, 0xc9, 0xdd, 0x6f, 0x3f, 0x79, 0x86, 0xa1, 0x08, 0x21, 0x14, 0x55, 0x08,
	0x75, 0xd3, 0xd4, 0xc1, 0x78, 0x82, 0xac, 0xcb, 0x34, 0x01, 0xdd, 0x86, 0xc9, 0x3e, 0xf4, 0x13,
	0xca, 0x12, 0x57, 0x8b, 0x48, 0xe3, 0xc8, 0x0e, 0x93, 0xc6, 0xe3, 0xf0, 0x81, 0xc7, 0x81, 0x57,
	0x42, 0xbe, 0x49, 0xda, 0x6c, 0xed, 0xa7, 0xda, 0xe7, 0x9e, 0x5c, 0xfb, 0x9c, 0x73, 0x5d, 0xd8,
	0x57, 0x3f, 0xe4, 0x3c, 0x49, 0xf9, 0x71, 0xfb, 0x3b, 0xae, 0xa4, 0xa8, 0x05, 0x85, 0x6e, 0x7b,
	0xff, 0x7e, 0xf8, 0xdb, 0x04, 0xe7, 0x3c, 0xe7, 0x45, 0x46, 0x29, 0xd8, 0x65, 0xb2, 0xe0, 0xbe,
	0x11, 0x18, 0x23, 0x8f, 0xe1, 0x5a, 0x63, 0xf5, 0x43, 0xc5, 0x7d, 0xb3, 0xc1, 0xf4, 0x9a, 0x1e,
	0x80, 0xfd, 0x3d, 0x2f, 0x33, 0xdf, 0x0a, 0x8c, 0xd1, 0xff, 0x27, 0xfb, 0xe3, 0x55, 0xb3, 0x31,
	0x36, 0xfa, 0x94, 0x97, 0x19, 0x43, 0x0a, 0x7d, 0x01, 0x83, 0xb9, 0x90, 0x8b, 0xa4, 0xf6, 0x6d,
	0x6c, 0xd0, 0xee, 0xe8, 0x3b, 0x70, 0x2b, 0xa1, 0xf2, 0x3a, 0x17, 0xa5, 0xef, 0x60, 0x9b, 0xe7,
	0xfd, 0x36, 0xd7, 0x6d, 0x8d, 0x2d, 0x59, 0xf4, 0x35, 0x40, 0x99, 0xd4, 0xf9, 0x3d, 0x8f, 0xf5,
	0x75, 0x06, 0xd8, 0xad, 0x87, 0xd0, 0x57, 0xe0, 0xcd, 0xf5, 0xe1, 0x97, 0x5a, 0xc1, 0x36, 0x96,
	0x57, 0x00, 0x7d, 0x03, 0xbb, 0x55, 0x22, 0x93, 0x05, 0xaf, 0xb9, 0x44, 0x86, 0x8b, 0x8c, 0xc7,
	0xa0, 0xee, 0xa1, 0xb8, 0xcc, 0x93, 0x22, 0xff, 0xc9, 0x7d, 0x2f, 0x30, 0x46, 0x2e, 0x5b, 0x01,
	0xc3, 0xbf, 0x06, 0xd8, 0x78, 0xd4, 0x26, 0x9f, 0x46, 0xad, 0x27, 0xe6, 0xba, 0x18, 0xfd, 0x4d,
	0xcf, 0x92, 0x00, 0x76, 0x32, 0xae, 0x52, 0x99, 0x57, 0xa8, 0xde, 0xc2, 0x26, 0x7d, 0x48, 0x33,
	0x52, 0x51, 0xd6, 0xbc, 0xac, 0x51, 0x6b, 0xe3, 0x5c, 0x1f, 0xa2, 0x07, 0x30, 0x40, 0x6d, 0xca,
	0x77, 0x02, 0x6b, 0xb4, 0x73, 0xb2, 0xb7, 0x96, 0x01, 0x6b, 0x09, 0xf4, 0x25, 0xb8, 0x3a, 0x34,
	0x14, 0xdd, 0xb8, 0xb6, 0xdc, 0x0f, 0xff, 0x98, 0x30, 0x98, 0xf2, 0xfa, 0x4e, 0x64, 0x5a, 0xba,
	0xa8, 0xb8, 0x4c, 0xf0, 0x4e, 0x8d, 0xb0, 0x15, 0xa0, 0x15, 0x57, 0x49, 0x7d, 0xd7, 0x4d, 0x81,
	0x5e, 0xeb, 0x68, 0x17, 0xf8, 0x6d, 0x2b, 0xa1, 0xdd, 0x3d, 0xd5, 0x67, 0xaf, 0xeb, 0xeb, 0xfc,
	0x73, 0x7a, 0xfe, 0x05, 0xb0, 0x73, 0x97, 0x94, 0x59, 0xd1, 0xc6, 0xd3, 0xdc, 0xb4, 0x0f, 0x61,
	0x84, 0x52, 0xa4, 0x5c, 0x29, 0x21, 0x7b, 0x21, 0x3f, 0x06, 0xf5, 0x98, 0xa4, 0x45, 0xce, 0xcb,
	0xba, 0x97, 0x72, 0x0f, 0xa1, 0x63, 0xa0, 0xcb, 0xcc, 0x55, 0xdc, 0x19, 0xe3, 0x21, 0x6f, 0x43,
	0x85, 0x1e, 0xc1, 0x9e, 0xe4, 0xaa, 0x12, 0xa5, 0xe2, 0x2b, 0x3a, 0x20, 0x7d, 0xbd, 0x30, 0xfc,
	0x65, 0x80, 0x33, 0x15, 0x19, 0x2f, 0x36, 0xce, 0xc8, 0x5b, 0x70, 0xb4, 0xf5, 0xca, 0x37, 0x31,
	0x34, 0xf2, 0x74, 0x48, 0x58, 0x53, 0xa6, 0x47, 0xb0, 0xdd, 0x78, 0xa9, 0x7c, 0x0b, 0x99, 0xb4,
	0xcf, 0x6c, 0x02, 0x63, 0x1d, 0x85, 0x1e, 0xc3, 0x33, 0xf5, 0xb0, 0xb8, 0x15, 0x45, 0x9e, 0x7e,
	0x93, 0x7c, 0xce, 0x25, 0x2f, 0x53, 0xae, 0x7c, 0x3b, 0xb0, 0xb4, 0xa4, 0xae, 0xc4, 0x96, 0x95,
	0xc3, 0x09, 0x78, 0xcb, 0x67, 0x4a, 0x01, 0x06, 0x5f, 0x27, 0xe1, 0xe7, 0x90, 0x91, 0x2d, 0xba,
	0x0d, 0xd6, 0x34, 0xbc, 0x26, 0x06, 0xf5, 0xc0, 0x09, 0x19, 0x0b, 0x67, 0xc4, 0xa4, 0xbb, 0xe0,
	0xb1, 0xe8, 0x3c, 0x62, 0xd1, 0xe5, 0x24, 0x22, 0x96, 0xa6, 0x84, 0x97, 0x33, 0x62, 0x1f, 0x0e,
	0xc1, 0xed, 0xe6, 0x1a, 0x7b, 0xc4, 0xec, 0x66, 0x12, 0x93, 0x2d, 0xbd, 0xbe, 0x3a, 0xfd, 0x18,
	0x4d, 0x62, 0x62, 0x1c, 0x4e, 0xc0, 0xed, 0x1e, 0x32, 0x75, 0xc1, 0x3e, 0xbd, 0x3a, 0x9b, 0x35,
	0x8c, 0x8b, 0x28, 0x3c, 0x8b, 0x18, 0x31, 0xe8, 0x7f, 0xe0, 0x9e, 0x5f, 0xb1, 0xe9, 0x59, 0x18,
	0x87, 0xc4, 0xd4, 0xc7, 0x7e, 0xb9, 0x89, 0xd8, 0x8c, 0x58, 0x9a, 0x7e, 0x1d, 0xc6, 0x17, 0xc4,
	0xbe, 0x1d, 0xe0, 0x3f, 0xd6, 0x87, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x07, 0xac, 0x60, 0x12,
	0xca, 0x04, 0x00, 0x00,
}
