// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contentkm.proto

package grpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Document_Language int32

const (
	Document_ENGLISH    Document_Language = 0
	Document_FRENCH     Document_Language = 1
	Document_PORTUGUESE Document_Language = 2
	Document_SPANISH    Document_Language = 3
	Document_SWAHILI    Document_Language = 4
)

var Document_Language_name = map[int32]string{
	0: "ENGLISH",
	1: "FRENCH",
	2: "PORTUGUESE",
	3: "SPANISH",
	4: "SWAHILI",
}

var Document_Language_value = map[string]int32{
	"ENGLISH":    0,
	"FRENCH":     1,
	"PORTUGUESE": 2,
	"SPANISH":    3,
	"SWAHILI":    4,
}

func (x Document_Language) String() string {
	return proto.EnumName(Document_Language_name, int32(x))
}

func (Document_Language) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_63e96ee2b70a1102, []int{2, 0}
}

type Document_Importance int32

const (
	Document_LOW      Document_Importance = 0
	Document_MEDIUM   Document_Importance = 1
	Document_HIGH     Document_Importance = 2
	Document_CRITICAL Document_Importance = 3
)

var Document_Importance_name = map[int32]string{
	0: "LOW",
	1: "MEDIUM",
	2: "HIGH",
	3: "CRITICAL",
}

var Document_Importance_value = map[string]int32{
	"LOW":      0,
	"MEDIUM":   1,
	"HIGH":     2,
	"CRITICAL": 3,
}

func (x Document_Importance) String() string {
	return proto.EnumName(Document_Importance_name, int32(x))
}

func (Document_Importance) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_63e96ee2b70a1102, []int{2, 1}
}

type Document_MediaFilter int32

const (
	Document_IG_WILLOW    Document_MediaFilter = 0
	Document_IG_WALDEN    Document_MediaFilter = 1
	Document_IG_VALENCIA  Document_MediaFilter = 2
	Document_IG_TOASTER   Document_MediaFilter = 3
	Document_IG_SUTRO     Document_MediaFilter = 4
	Document_IG_SIERRA    Document_MediaFilter = 5
	Document_IG_RISE      Document_MediaFilter = 6
	Document_IG_NASHVILLE Document_MediaFilter = 7
	Document_IG_MAYFAIR   Document_MediaFilter = 8
	Document_IG_LOFI      Document_MediaFilter = 9
	Document_IG_KELVIN    Document_MediaFilter = 10
	Document_IG_INKWELL   Document_MediaFilter = 11
	Document_IG_HUDSON    Document_MediaFilter = 12
	Document_IG_HEFE      Document_MediaFilter = 13
	Document_IG_EARLYBIRD Document_MediaFilter = 14
	Document_IG_BRANNAN   Document_MediaFilter = 15
	Document_IG_AMARO     Document_MediaFilter = 16
	Document_IG_1977      Document_MediaFilter = 17
)

var Document_MediaFilter_name = map[int32]string{
	0:  "IG_WILLOW",
	1:  "IG_WALDEN",
	2:  "IG_VALENCIA",
	3:  "IG_TOASTER",
	4:  "IG_SUTRO",
	5:  "IG_SIERRA",
	6:  "IG_RISE",
	7:  "IG_NASHVILLE",
	8:  "IG_MAYFAIR",
	9:  "IG_LOFI",
	10: "IG_KELVIN",
	11: "IG_INKWELL",
	12: "IG_HUDSON",
	13: "IG_HEFE",
	14: "IG_EARLYBIRD",
	15: "IG_BRANNAN",
	16: "IG_AMARO",
	17: "IG_1977",
}

var Document_MediaFilter_value = map[string]int32{
	"IG_WILLOW":    0,
	"IG_WALDEN":    1,
	"IG_VALENCIA":  2,
	"IG_TOASTER":   3,
	"IG_SUTRO":     4,
	"IG_SIERRA":    5,
	"IG_RISE":      6,
	"IG_NASHVILLE": 7,
	"IG_MAYFAIR":   8,
	"IG_LOFI":      9,
	"IG_KELVIN":    10,
	"IG_INKWELL":   11,
	"IG_HUDSON":    12,
	"IG_HEFE":      13,
	"IG_EARLYBIRD": 14,
	"IG_BRANNAN":   15,
	"IG_AMARO":     16,
	"IG_1977":      17,
}

func (x Document_MediaFilter) String() string {
	return proto.EnumName(Document_MediaFilter_name, int32(x))
}

func (Document_MediaFilter) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_63e96ee2b70a1102, []int{2, 2}
}

type Media_Type int32

const (
	Media_AUDIO    Media_Type = 0
	Media_DOCUMENT Media_Type = 1
	Media_IMAGE    Media_Type = 2
	Media_VIDEO    Media_Type = 3
)

var Media_Type_name = map[int32]string{
	0: "AUDIO",
	1: "DOCUMENT",
	2: "IMAGE",
	3: "VIDEO",
}

var Media_Type_value = map[string]int32{
	"AUDIO":    0,
	"DOCUMENT": 1,
	"IMAGE":    2,
	"VIDEO":    3,
}

func (x Media_Type) String() string {
	return proto.EnumName(Media_Type_name, int32(x))
}

func (Media_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_63e96ee2b70a1102, []int{3, 0}
}

type Content struct {
	Documents            []*Document `protobuf:"bytes,1,rep,name=documents,proto3" json:"documents,omitempty"`
	MediaItems           []*Media    `protobuf:"bytes,2,rep,name=media_items,json=mediaItems,proto3" json:"media_items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Content) Reset()         { *m = Content{} }
func (m *Content) String() string { return proto.CompactTextString(m) }
func (*Content) ProtoMessage()    {}
func (*Content) Descriptor() ([]byte, []int) {
	return fileDescriptor_63e96ee2b70a1102, []int{0}
}

func (m *Content) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Content.Unmarshal(m, b)
}
func (m *Content) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Content.Marshal(b, m, deterministic)
}
func (m *Content) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Content.Merge(m, src)
}
func (m *Content) XXX_Size() int {
	return xxx_messageInfo_Content.Size(m)
}
func (m *Content) XXX_DiscardUnknown() {
	xxx_messageInfo_Content.DiscardUnknown(m)
}

var xxx_messageInfo_Content proto.InternalMessageInfo

func (m *Content) GetDocuments() []*Document {
	if m != nil {
		return m.Documents
	}
	return nil
}

func (m *Content) GetMediaItems() []*Media {
	if m != nil {
		return m.MediaItems
	}
	return nil
}

type ContentHandles struct {
	ItemIds              []*Identification `protobuf:"bytes,1,rep,name=item_ids,json=itemIds,proto3" json:"item_ids,omitempty"`
	Message              string            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ContentHandles) Reset()         { *m = ContentHandles{} }
func (m *ContentHandles) String() string { return proto.CompactTextString(m) }
func (*ContentHandles) ProtoMessage()    {}
func (*ContentHandles) Descriptor() ([]byte, []int) {
	return fileDescriptor_63e96ee2b70a1102, []int{1}
}

func (m *ContentHandles) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentHandles.Unmarshal(m, b)
}
func (m *ContentHandles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentHandles.Marshal(b, m, deterministic)
}
func (m *ContentHandles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentHandles.Merge(m, src)
}
func (m *ContentHandles) XXX_Size() int {
	return xxx_messageInfo_ContentHandles.Size(m)
}
func (m *ContentHandles) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentHandles.DiscardUnknown(m)
}

var xxx_messageInfo_ContentHandles proto.InternalMessageInfo

func (m *ContentHandles) GetItemIds() []*Identification {
	if m != nil {
		return m.ItemIds
	}
	return nil
}

func (m *ContentHandles) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Document struct {
	Title                string               `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Slug                 string               `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
	Publish              bool                 `protobuf:"varint,3,opt,name=publish,proto3" json:"publish,omitempty"`
	Body                 string               `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	Language             Document_Language    `protobuf:"varint,5,opt,name=language,proto3,enum=grpc.Document_Language" json:"language,omitempty"`
	Importance           Document_Importance  `protobuf:"varint,6,opt,name=importance,proto3,enum=grpc.Document_Importance" json:"importance,omitempty"`
	MediaFilter          Document_MediaFilter `protobuf:"varint,7,opt,name=media_filter,json=mediaFilter,proto3,enum=grpc.Document_MediaFilter" json:"media_filter,omitempty"`
	Metadata             *MetaData            `protobuf:"bytes,8,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Document) Reset()         { *m = Document{} }
func (m *Document) String() string { return proto.CompactTextString(m) }
func (*Document) ProtoMessage()    {}
func (*Document) Descriptor() ([]byte, []int) {
	return fileDescriptor_63e96ee2b70a1102, []int{2}
}

func (m *Document) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Document.Unmarshal(m, b)
}
func (m *Document) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Document.Marshal(b, m, deterministic)
}
func (m *Document) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Document.Merge(m, src)
}
func (m *Document) XXX_Size() int {
	return xxx_messageInfo_Document.Size(m)
}
func (m *Document) XXX_DiscardUnknown() {
	xxx_messageInfo_Document.DiscardUnknown(m)
}

var xxx_messageInfo_Document proto.InternalMessageInfo

func (m *Document) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Document) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Document) GetPublish() bool {
	if m != nil {
		return m.Publish
	}
	return false
}

func (m *Document) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Document) GetLanguage() Document_Language {
	if m != nil {
		return m.Language
	}
	return Document_ENGLISH
}

func (m *Document) GetImportance() Document_Importance {
	if m != nil {
		return m.Importance
	}
	return Document_LOW
}

func (m *Document) GetMediaFilter() Document_MediaFilter {
	if m != nil {
		return m.MediaFilter
	}
	return Document_IG_WILLOW
}

func (m *Document) GetMetadata() *MetaData {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Media struct {
	Name                 string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 Media_Type `protobuf:"varint,2,opt,name=type,proto3,enum=grpc.Media_Type" json:"type,omitempty"`
	FileUrl              string     `protobuf:"bytes,3,opt,name=file_url,json=fileUrl,proto3" json:"file_url,omitempty"`
	Metadata             *MetaData  `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Media) Reset()         { *m = Media{} }
func (m *Media) String() string { return proto.CompactTextString(m) }
func (*Media) ProtoMessage()    {}
func (*Media) Descriptor() ([]byte, []int) {
	return fileDescriptor_63e96ee2b70a1102, []int{3}
}

func (m *Media) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Media.Unmarshal(m, b)
}
func (m *Media) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Media.Marshal(b, m, deterministic)
}
func (m *Media) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Media.Merge(m, src)
}
func (m *Media) XXX_Size() int {
	return xxx_messageInfo_Media.Size(m)
}
func (m *Media) XXX_DiscardUnknown() {
	xxx_messageInfo_Media.DiscardUnknown(m)
}

var xxx_messageInfo_Media proto.InternalMessageInfo

func (m *Media) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Media) GetType() Media_Type {
	if m != nil {
		return m.Type
	}
	return Media_AUDIO
}

func (m *Media) GetFileUrl() string {
	if m != nil {
		return m.FileUrl
	}
	return ""
}

func (m *Media) GetMetadata() *MetaData {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type MetaData struct {
	Identification       *Identification     `protobuf:"bytes,1,opt,name=identification,proto3" json:"identification,omitempty"`
	Timestamps           *TimeStamps         `protobuf:"bytes,2,opt,name=timestamps,proto3" json:"timestamps,omitempty"`
	M                    map[string]*any.Any `protobuf:"bytes,3,rep,name=m,proto3" json:"m,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MetaData) Reset()         { *m = MetaData{} }
func (m *MetaData) String() string { return proto.CompactTextString(m) }
func (*MetaData) ProtoMessage()    {}
func (*MetaData) Descriptor() ([]byte, []int) {
	return fileDescriptor_63e96ee2b70a1102, []int{4}
}

func (m *MetaData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetaData.Unmarshal(m, b)
}
func (m *MetaData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetaData.Marshal(b, m, deterministic)
}
func (m *MetaData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaData.Merge(m, src)
}
func (m *MetaData) XXX_Size() int {
	return xxx_messageInfo_MetaData.Size(m)
}
func (m *MetaData) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaData.DiscardUnknown(m)
}

var xxx_messageInfo_MetaData proto.InternalMessageInfo

func (m *MetaData) GetIdentification() *Identification {
	if m != nil {
		return m.Identification
	}
	return nil
}

func (m *MetaData) GetTimestamps() *TimeStamps {
	if m != nil {
		return m.Timestamps
	}
	return nil
}

func (m *MetaData) GetM() map[string]*any.Any {
	if m != nil {
		return m.M
	}
	return nil
}

type Identification struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Identifier           string   `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Tags                 []string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Identification) Reset()         { *m = Identification{} }
func (m *Identification) String() string { return proto.CompactTextString(m) }
func (*Identification) ProtoMessage()    {}
func (*Identification) Descriptor() ([]byte, []int) {
	return fileDescriptor_63e96ee2b70a1102, []int{5}
}

func (m *Identification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Identification.Unmarshal(m, b)
}
func (m *Identification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Identification.Marshal(b, m, deterministic)
}
func (m *Identification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identification.Merge(m, src)
}
func (m *Identification) XXX_Size() int {
	return xxx_messageInfo_Identification.Size(m)
}
func (m *Identification) XXX_DiscardUnknown() {
	xxx_messageInfo_Identification.DiscardUnknown(m)
}

var xxx_messageInfo_Identification proto.InternalMessageInfo

func (m *Identification) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Identification) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *Identification) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type TimeStamps struct {
	Created              *timestamp.Timestamp `protobuf:"bytes,1,opt,name=created,proto3" json:"created,omitempty"`
	Updated              *timestamp.Timestamp `protobuf:"bytes,2,opt,name=updated,proto3" json:"updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TimeStamps) Reset()         { *m = TimeStamps{} }
func (m *TimeStamps) String() string { return proto.CompactTextString(m) }
func (*TimeStamps) ProtoMessage()    {}
func (*TimeStamps) Descriptor() ([]byte, []int) {
	return fileDescriptor_63e96ee2b70a1102, []int{6}
}

func (m *TimeStamps) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeStamps.Unmarshal(m, b)
}
func (m *TimeStamps) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeStamps.Marshal(b, m, deterministic)
}
func (m *TimeStamps) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeStamps.Merge(m, src)
}
func (m *TimeStamps) XXX_Size() int {
	return xxx_messageInfo_TimeStamps.Size(m)
}
func (m *TimeStamps) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeStamps.DiscardUnknown(m)
}

var xxx_messageInfo_TimeStamps proto.InternalMessageInfo

func (m *TimeStamps) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *TimeStamps) GetUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.Updated
	}
	return nil
}

func init() {
	proto.RegisterEnum("grpc.Document_Language", Document_Language_name, Document_Language_value)
	proto.RegisterEnum("grpc.Document_Importance", Document_Importance_name, Document_Importance_value)
	proto.RegisterEnum("grpc.Document_MediaFilter", Document_MediaFilter_name, Document_MediaFilter_value)
	proto.RegisterEnum("grpc.Media_Type", Media_Type_name, Media_Type_value)
	proto.RegisterType((*Content)(nil), "grpc.Content")
	proto.RegisterType((*ContentHandles)(nil), "grpc.ContentHandles")
	proto.RegisterType((*Document)(nil), "grpc.Document")
	proto.RegisterType((*Media)(nil), "grpc.Media")
	proto.RegisterType((*MetaData)(nil), "grpc.MetaData")
	proto.RegisterMapType((map[string]*any.Any)(nil), "grpc.MetaData.MEntry")
	proto.RegisterType((*Identification)(nil), "grpc.Identification")
	proto.RegisterType((*TimeStamps)(nil), "grpc.TimeStamps")
}

func init() { proto.RegisterFile("contentkm.proto", fileDescriptor_63e96ee2b70a1102) }

var fileDescriptor_63e96ee2b70a1102 = []byte{
	// 979 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xae, 0xf3, 0xeb, 0x9c, 0xb4, 0xe9, 0x30, 0x2a, 0x5a, 0x37, 0x17, 0x50, 0x19, 0x2e, 0xaa,
	0x55, 0x95, 0x42, 0x17, 0x58, 0x8a, 0xe0, 0xc2, 0x9b, 0xb8, 0xce, 0x50, 0xc7, 0x59, 0x4d, 0x92,
	0x56, 0x2b, 0x84, 0x2c, 0x37, 0x9e, 0x06, 0x6b, 0x6d, 0x27, 0xb2, 0x27, 0x88, 0x3c, 0x0c, 0xcf,
	0xc2, 0x4b, 0xf0, 0x0e, 0x5c, 0xf0, 0x12, 0x68, 0xc6, 0xe3, 0x36, 0x29, 0x54, 0xcb, 0xdd, 0x9c,
	0x39, 0xdf, 0xf7, 0x9d, 0x9f, 0x39, 0x73, 0xe0, 0x70, 0xbe, 0x4c, 0x39, 0x4b, 0xf9, 0xfb, 0xa4,
	0xb7, 0xca, 0x96, 0x7c, 0x89, 0x6b, 0x8b, 0x6c, 0x35, 0xef, 0x1e, 0x2f, 0x96, 0xcb, 0x45, 0xcc,
	0xce, 0xe5, 0xdd, 0xdd, 0xfa, 0xfe, 0x3c, 0x48, 0x37, 0x05, 0xa0, 0xfb, 0xe9, 0x53, 0x17, 0x8f,
	0x12, 0x96, 0xf3, 0x20, 0x59, 0x15, 0x00, 0x93, 0x41, 0xb3, 0x5f, 0x88, 0xe2, 0x33, 0x68, 0x85,
	0xcb, 0xf9, 0x3a, 0x61, 0x29, 0xcf, 0x0d, 0xed, 0xa4, 0x7a, 0xda, 0xbe, 0xe8, 0xf4, 0x44, 0x80,
	0xde, 0x40, 0x5d, 0xd3, 0x47, 0x00, 0x3e, 0x83, 0x76, 0xc2, 0xc2, 0x28, 0xf0, 0x23, 0xce, 0x92,
	0xdc, 0xa8, 0x48, 0x7c, 0xbb, 0xc0, 0x8f, 0x84, 0x83, 0x82, 0xf4, 0x13, 0xe1, 0x36, 0x7f, 0x82,
	0x8e, 0x0a, 0x33, 0x0c, 0xd2, 0x30, 0x66, 0x39, 0x3e, 0x07, 0x5d, 0x30, 0xfd, 0x28, 0x2c, 0x83,
	0x1d, 0x15, 0x64, 0x12, 0xb2, 0x94, 0x47, 0xf7, 0xd1, 0x3c, 0xe0, 0xd1, 0x32, 0xa5, 0x4d, 0x81,
	0x22, 0x61, 0x8e, 0x0d, 0x68, 0x26, 0x2c, 0xcf, 0x83, 0x05, 0x33, 0x2a, 0x27, 0xda, 0x69, 0x8b,
	0x96, 0xa6, 0xf9, 0x57, 0x1d, 0xf4, 0x32, 0x45, 0x7c, 0x04, 0x75, 0x1e, 0xf1, 0x98, 0x19, 0x9a,
	0x04, 0x15, 0x06, 0xc6, 0x50, 0xcb, 0xe3, 0xf5, 0x42, 0x31, 0xe5, 0x59, 0x08, 0xae, 0xd6, 0x77,
	0x71, 0x94, 0xff, 0x62, 0x54, 0x4f, 0xb4, 0x53, 0x9d, 0x96, 0xa6, 0x40, 0xdf, 0x2d, 0xc3, 0x8d,
	0x51, 0x2b, 0xd0, 0xe2, 0x8c, 0x5f, 0x81, 0x1e, 0x07, 0xe9, 0x62, 0x2d, 0xe2, 0xd7, 0x4f, 0xb4,
	0xd3, 0xce, 0xc5, 0x8b, 0xdd, 0xe6, 0xf4, 0x5c, 0xe5, 0xa6, 0x0f, 0x40, 0x7c, 0x09, 0x10, 0x25,
	0xab, 0x65, 0xc6, 0x83, 0x74, 0xce, 0x8c, 0x86, 0xa4, 0x1d, 0x3f, 0xa1, 0x91, 0x07, 0x00, 0xdd,
	0x02, 0xe3, 0x1f, 0x60, 0xbf, 0xe8, 0xef, 0x7d, 0x14, 0x73, 0x96, 0x19, 0x4d, 0x49, 0xee, 0x3e,
	0x21, 0xcb, 0x4e, 0x5f, 0x49, 0x04, 0x2d, 0xde, 0xa3, 0x30, 0xf0, 0x4b, 0xd0, 0x13, 0xc6, 0x83,
	0x30, 0xe0, 0x81, 0xa1, 0x9f, 0x68, 0x8f, 0x6f, 0x39, 0x62, 0x3c, 0x18, 0x04, 0x3c, 0xa0, 0x0f,
	0x7e, 0x73, 0x04, 0x7a, 0x99, 0x3b, 0x6e, 0x43, 0xd3, 0xf6, 0x1c, 0x97, 0x4c, 0x86, 0x68, 0x0f,
	0x03, 0x34, 0xae, 0xa8, 0xed, 0xf5, 0x87, 0x48, 0xc3, 0x1d, 0x80, 0xb7, 0x63, 0x3a, 0x9d, 0x39,
	0x33, 0x7b, 0x62, 0xa3, 0x8a, 0x00, 0x4e, 0xde, 0x5a, 0x9e, 0x00, 0x56, 0xa5, 0x71, 0x6b, 0x0d,
	0x89, 0x4b, 0x50, 0xcd, 0xbc, 0x04, 0x78, 0xac, 0x09, 0x37, 0xa1, 0xea, 0x8e, 0x6f, 0x0b, 0xb1,
	0x91, 0x3d, 0x20, 0xb3, 0x11, 0xd2, 0xb0, 0x0e, 0xb5, 0x21, 0x71, 0x86, 0xa8, 0x82, 0xf7, 0x41,
	0xef, 0x53, 0x32, 0x25, 0x7d, 0xcb, 0x45, 0x55, 0xf3, 0xf7, 0x0a, 0xb4, 0xb7, 0x4a, 0xc2, 0x07,
	0xd0, 0x22, 0x8e, 0x7f, 0x4b, 0xdc, 0x42, 0x42, 0x99, 0x96, 0x3b, 0xb0, 0x3d, 0xa4, 0xe1, 0x43,
	0x68, 0x13, 0xc7, 0xbf, 0xb1, 0x5c, 0xdb, 0xeb, 0x13, 0x0b, 0x55, 0x44, 0x8e, 0xc4, 0xf1, 0xa7,
	0x63, 0x6b, 0x32, 0xb5, 0x29, 0xaa, 0x0a, 0x71, 0xe2, 0xf8, 0x93, 0xd9, 0x94, 0x8e, 0x51, 0x4d,
	0xb1, 0x27, 0xc4, 0xa6, 0xd4, 0x42, 0x75, 0x91, 0x33, 0x71, 0x7c, 0x4a, 0x26, 0x36, 0x6a, 0x60,
	0x04, 0xfb, 0xc4, 0xf1, 0x3d, 0x6b, 0x32, 0xbc, 0x21, 0xae, 0x6b, 0xa3, 0xa6, 0xd2, 0x1a, 0x59,
	0xef, 0xae, 0x2c, 0x42, 0x91, 0xae, 0xe0, 0xee, 0xf8, 0x8a, 0xa0, 0x96, 0x92, 0xba, 0xb6, 0xdd,
	0x1b, 0xe2, 0x21, 0x50, 0x58, 0xe2, 0x5d, 0xdf, 0xda, 0xae, 0x8b, 0xda, 0xca, 0x3d, 0x9c, 0x0d,
	0x26, 0x63, 0x0f, 0xed, 0x2b, 0xea, 0xd0, 0xbe, 0xb2, 0xd1, 0x81, 0x8a, 0x64, 0x5b, 0xd4, 0x7d,
	0xf7, 0x86, 0xd0, 0x01, 0xea, 0x28, 0xf6, 0x1b, 0x6a, 0x79, 0x9e, 0xe5, 0xa1, 0x43, 0x95, 0xb5,
	0x35, 0xb2, 0xe8, 0x18, 0x21, 0x45, 0xfe, 0xf2, 0xf2, 0xf5, 0x6b, 0xf4, 0x91, 0xf9, 0x87, 0x06,
	0x75, 0xd9, 0x1f, 0x31, 0xa2, 0x69, 0x90, 0x94, 0x53, 0x2e, 0xcf, 0xf8, 0x73, 0xa8, 0xf1, 0xcd,
	0xaa, 0xf8, 0x1e, 0x9d, 0x0b, 0xb4, 0xf5, 0x17, 0x7b, 0xd3, 0xcd, 0x8a, 0x51, 0xe9, 0xc5, 0xc7,
	0xa0, 0xdf, 0x47, 0x31, 0xf3, 0xd7, 0x59, 0x2c, 0xe7, 0xbe, 0x45, 0x9b, 0xc2, 0x9e, 0x65, 0xf1,
	0xce, 0xd0, 0xd4, 0x3e, 0x30, 0x34, 0x5f, 0x43, 0x4d, 0x88, 0xe2, 0x16, 0xd4, 0xad, 0xd9, 0x80,
	0x8c, 0xd1, 0x9e, 0x48, 0x7c, 0x30, 0xee, 0xcf, 0x46, 0xb6, 0x37, 0x45, 0x9a, 0x70, 0x90, 0x91,
	0xe5, 0x88, 0x59, 0x69, 0x41, 0xfd, 0x86, 0x0c, 0xec, 0x31, 0xaa, 0x9a, 0x7f, 0x6b, 0xa0, 0x97,
	0x6a, 0xf8, 0x7b, 0xe8, 0x44, 0x3b, 0xbf, 0x5d, 0x96, 0xf3, 0xdc, 0x26, 0x78, 0x82, 0xc5, 0x5f,
	0x00, 0x3c, 0x6c, 0xb3, 0x5c, 0x16, 0xdd, 0x2e, 0x8b, 0x9e, 0x46, 0x09, 0x9b, 0xc8, 0x7b, 0xba,
	0x85, 0xc1, 0x9f, 0x81, 0x96, 0x18, 0x55, 0xb9, 0x6c, 0x3e, 0xde, 0x2d, 0xac, 0x37, 0xb2, 0x53,
	0x9e, 0x6d, 0xa8, 0x96, 0x74, 0x7f, 0x14, 0x73, 0x2a, 0x0c, 0x8c, 0xa0, 0xfa, 0x9e, 0x6d, 0x54,
	0x8b, 0xc5, 0x11, 0xbf, 0x84, 0xfa, 0xaf, 0x41, 0xbc, 0x66, 0x2a, 0xda, 0x51, 0xaf, 0x58, 0xaf,
	0xbd, 0x72, 0xbd, 0xf6, 0xac, 0x74, 0x43, 0x0b, 0xc8, 0x77, 0x95, 0x6f, 0x35, 0xf3, 0x67, 0xe8,
	0xec, 0x16, 0x81, 0x5f, 0x40, 0x73, 0x9d, 0xb3, 0xcc, 0x8f, 0x42, 0xa5, 0xdb, 0x10, 0x26, 0x09,
	0xf1, 0x27, 0x00, 0x65, 0x7d, 0x2c, 0x53, 0x7b, 0x6a, 0xeb, 0x46, 0x3c, 0x38, 0x0f, 0x16, 0xb9,
	0x4c, 0xbf, 0x45, 0xe5, 0xd9, 0xfc, 0x0d, 0xe0, 0xb1, 0x52, 0xfc, 0x15, 0x34, 0xe7, 0x19, 0x0b,
	0x38, 0x0b, 0x55, 0x1b, 0xbb, 0xff, 0x4a, 0x6f, 0x5a, 0xf6, 0x82, 0x96, 0x50, 0xc1, 0x5a, 0xaf,
	0x42, 0xc9, 0xaa, 0x7c, 0x98, 0xa5, 0xa0, 0x17, 0x7f, 0x6a, 0x00, 0x6a, 0xa1, 0x5f, 0x07, 0x09,
	0x3e, 0x87, 0x46, 0x5f, 0xea, 0xe1, 0x83, 0xa2, 0xaf, 0xca, 0xd7, 0x3d, 0xda, 0x31, 0xd5, 0xee,
	0x37, 0xf7, 0x04, 0x61, 0x26, 0xa5, 0xfe, 0x2f, 0xe1, 0x0c, 0xaa, 0x0e, 0xe3, 0xf8, 0x3f, 0xdd,
	0xdd, 0x5d, 0x0d, 0x73, 0x0f, 0x7f, 0x03, 0x8d, 0x01, 0x8b, 0x19, 0x67, 0xcf, 0x10, 0x9e, 0x89,
	0x72, 0xd7, 0x90, 0x35, 0xbf, 0xfa, 0x27, 0x00, 0x00, 0xff, 0xff, 0x07, 0x35, 0xf5, 0x95, 0x69,
	0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ContentKamClient is the client API for ContentKam service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContentKamClient interface {
	// Create
	Create(ctx context.Context, in *Content, opts ...grpc.CallOption) (*ContentHandles, error)
	// Update
	Update(ctx context.Context, in *Content, opts ...grpc.CallOption) (*ContentHandles, error)
	// Get
	Get(ctx context.Context, in *ContentHandles, opts ...grpc.CallOption) (*Content, error)
	// Delete
	Delete(ctx context.Context, in *ContentHandles, opts ...grpc.CallOption) (*ContentHandles, error)
}

type contentKamClient struct {
	cc *grpc.ClientConn
}

func NewContentKamClient(cc *grpc.ClientConn) ContentKamClient {
	return &contentKamClient{cc}
}

func (c *contentKamClient) Create(ctx context.Context, in *Content, opts ...grpc.CallOption) (*ContentHandles, error) {
	out := new(ContentHandles)
	err := c.cc.Invoke(ctx, "/grpc.ContentKam/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentKamClient) Update(ctx context.Context, in *Content, opts ...grpc.CallOption) (*ContentHandles, error) {
	out := new(ContentHandles)
	err := c.cc.Invoke(ctx, "/grpc.ContentKam/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentKamClient) Get(ctx context.Context, in *ContentHandles, opts ...grpc.CallOption) (*Content, error) {
	out := new(Content)
	err := c.cc.Invoke(ctx, "/grpc.ContentKam/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentKamClient) Delete(ctx context.Context, in *ContentHandles, opts ...grpc.CallOption) (*ContentHandles, error) {
	out := new(ContentHandles)
	err := c.cc.Invoke(ctx, "/grpc.ContentKam/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentKamServer is the server API for ContentKam service.
type ContentKamServer interface {
	// Create
	Create(context.Context, *Content) (*ContentHandles, error)
	// Update
	Update(context.Context, *Content) (*ContentHandles, error)
	// Get
	Get(context.Context, *ContentHandles) (*Content, error)
	// Delete
	Delete(context.Context, *ContentHandles) (*ContentHandles, error)
}

func RegisterContentKamServer(s *grpc.Server, srv ContentKamServer) {
	s.RegisterService(&_ContentKam_serviceDesc, srv)
}

func _ContentKam_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Content)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentKamServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ContentKam/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentKamServer).Create(ctx, req.(*Content))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentKam_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Content)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentKamServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ContentKam/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentKamServer).Update(ctx, req.(*Content))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentKam_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContentHandles)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentKamServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ContentKam/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentKamServer).Get(ctx, req.(*ContentHandles))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentKam_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContentHandles)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentKamServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ContentKam/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentKamServer).Delete(ctx, req.(*ContentHandles))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContentKam_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.ContentKam",
	HandlerType: (*ContentKamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ContentKam_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ContentKam_Update_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ContentKam_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ContentKam_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contentkm.proto",
}