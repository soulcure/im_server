// Code generated by protoc-gen-go.
// source: IM.Msg.proto
// DO NOT EDIT!

package youmai

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 消息类型
type IM_CONTENT_TYPE int32

const (
	IM_CONTENT_TYPE_IM_CONTENT_TYPE_TEXT          IM_CONTENT_TYPE = 0
	IM_CONTENT_TYPE_IM_CONTENT_TYPE_SHORT_MESSAGE IM_CONTENT_TYPE = 1
	IM_CONTENT_TYPE_IM_CONTENT_TYPE_CONTACTS      IM_CONTENT_TYPE = 2
	IM_CONTENT_TYPE_IM_CONTENT_TYPE_RECOMMEND_APP IM_CONTENT_TYPE = 3
	IM_CONTENT_TYPE_IM_CONTENT_TYPE_NO_DISTURB    IM_CONTENT_TYPE = 4
	IM_CONTENT_TYPE_IM_CONTENT_TYPE_LOCATIONSHARE IM_CONTENT_TYPE = 5
	IM_CONTENT_TYPE_IM_CONTENT_TYPE_IMAGE         IM_CONTENT_TYPE = 6
	IM_CONTENT_TYPE_IM_CONTENT_TYPE_AT            IM_CONTENT_TYPE = 7
	IM_CONTENT_TYPE_IM_CONTENT_TYPE_URL           IM_CONTENT_TYPE = 8
	IM_CONTENT_TYPE_IM_CONTENT_TYPE_AUDIO         IM_CONTENT_TYPE = 9
	IM_CONTENT_TYPE_IM_CONTENT_TYPE_VIDEO         IM_CONTENT_TYPE = 10
	IM_CONTENT_TYPE_IM_CONTENT_TYPE_LOCATION      IM_CONTENT_TYPE = 11
	IM_CONTENT_TYPE_IM_CONTENT_TYPE_FILE          IM_CONTENT_TYPE = 12
	IM_CONTENT_TYPE_IM_CONTENT_TYPE_BIZCARD       IM_CONTENT_TYPE = 13
	IM_CONTENT_TYPE_IM_CONTENT_TYPE_KAQUAN        IM_CONTENT_TYPE = 14
)

var IM_CONTENT_TYPE_name = map[int32]string{
	0:  "IM_CONTENT_TYPE_TEXT",
	1:  "IM_CONTENT_TYPE_SHORT_MESSAGE",
	2:  "IM_CONTENT_TYPE_CONTACTS",
	3:  "IM_CONTENT_TYPE_RECOMMEND_APP",
	4:  "IM_CONTENT_TYPE_NO_DISTURB",
	5:  "IM_CONTENT_TYPE_LOCATIONSHARE",
	6:  "IM_CONTENT_TYPE_IMAGE",
	7:  "IM_CONTENT_TYPE_AT",
	8:  "IM_CONTENT_TYPE_URL",
	9:  "IM_CONTENT_TYPE_AUDIO",
	10: "IM_CONTENT_TYPE_VIDEO",
	11: "IM_CONTENT_TYPE_LOCATION",
	12: "IM_CONTENT_TYPE_FILE",
	13: "IM_CONTENT_TYPE_BIZCARD",
	14: "IM_CONTENT_TYPE_KAQUAN",
}
var IM_CONTENT_TYPE_value = map[string]int32{
	"IM_CONTENT_TYPE_TEXT":          0,
	"IM_CONTENT_TYPE_SHORT_MESSAGE": 1,
	"IM_CONTENT_TYPE_CONTACTS":      2,
	"IM_CONTENT_TYPE_RECOMMEND_APP": 3,
	"IM_CONTENT_TYPE_NO_DISTURB":    4,
	"IM_CONTENT_TYPE_LOCATIONSHARE": 5,
	"IM_CONTENT_TYPE_IMAGE":         6,
	"IM_CONTENT_TYPE_AT":            7,
	"IM_CONTENT_TYPE_URL":           8,
	"IM_CONTENT_TYPE_AUDIO":         9,
	"IM_CONTENT_TYPE_VIDEO":         10,
	"IM_CONTENT_TYPE_LOCATION":      11,
	"IM_CONTENT_TYPE_FILE":          12,
	"IM_CONTENT_TYPE_BIZCARD":       13,
	"IM_CONTENT_TYPE_KAQUAN":        14,
}

func (x IM_CONTENT_TYPE) Enum() *IM_CONTENT_TYPE {
	p := new(IM_CONTENT_TYPE)
	*p = x
	return p
}
func (x IM_CONTENT_TYPE) String() string {
	return proto.EnumName(IM_CONTENT_TYPE_name, int32(x))
}
func (x *IM_CONTENT_TYPE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(IM_CONTENT_TYPE_value, data, "IM_CONTENT_TYPE")
	if err != nil {
		return err
	}
	*x = IM_CONTENT_TYPE(value)
	return nil
}
func (IM_CONTENT_TYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

// 消息属性
type MsgProperty int32

const (
	MsgProperty_MSG_PTOPERTY_NONE    MsgProperty = 1
	MsgProperty_MSG_PROPERTY_RECEIPT MsgProperty = 2
	MsgProperty_MSG_PROPERTY_URGENT  MsgProperty = 3
	MsgProperty_MSG_PROPERTY_QUOTE   MsgProperty = 4
	MsgProperty_MSG_PROPERTY_RECALL  MsgProperty = 5
)

var MsgProperty_name = map[int32]string{
	1: "MSG_PTOPERTY_NONE",
	2: "MSG_PROPERTY_RECEIPT",
	3: "MSG_PROPERTY_URGENT",
	4: "MSG_PROPERTY_QUOTE",
	5: "MSG_PROPERTY_RECALL",
}
var MsgProperty_value = map[string]int32{
	"MSG_PTOPERTY_NONE":    1,
	"MSG_PROPERTY_RECEIPT": 2,
	"MSG_PROPERTY_URGENT":  3,
	"MSG_PROPERTY_QUOTE":   4,
	"MSG_PROPERTY_RECALL":  5,
}

func (x MsgProperty) Enum() *MsgProperty {
	p := new(MsgProperty)
	*p = x
	return p
}
func (x MsgProperty) String() string {
	return proto.EnumName(MsgProperty_name, int32(x))
}
func (x *MsgProperty) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MsgProperty_value, data, "MsgProperty")
	if err != nil {
		return err
	}
	*x = MsgProperty(value)
	return nil
}
func (MsgProperty) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

// 会话类型
type SessionType int32

const (
	SessionType_SESSION_TYPE_SINGLE    SessionType = 1
	SessionType_SESSION_TYPE_MULTICHAT SessionType = 2
	SessionType_SESSION_TYPE_ORGGROUP  SessionType = 3
)

var SessionType_name = map[int32]string{
	1: "SESSION_TYPE_SINGLE",
	2: "SESSION_TYPE_MULTICHAT",
	3: "SESSION_TYPE_ORGGROUP",
}
var SessionType_value = map[string]int32{
	"SESSION_TYPE_SINGLE":    1,
	"SESSION_TYPE_MULTICHAT": 2,
	"SESSION_TYPE_ORGGROUP":  3,
}

func (x SessionType) Enum() *SessionType {
	p := new(SessionType)
	*p = x
	return p
}
func (x SessionType) String() string {
	return proto.EnumName(SessionType_name, int32(x))
}
func (x *SessionType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SessionType_value, data, "SessionType")
	if err != nil {
		return err
	}
	*x = SessionType(value)
	return nil
}
func (SessionType) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

// ==========================================================
// 消息信息子结构
// ==========================================================
// 消息字体信息
// font_style |= _bBold ? 0x01 : 0;
// font_style |= _bItalic ? 0x02 : 0;
// font_style |= _bUnderLine ? 0x03 : 0;
//
// _bBold = font_style & 0x01 ? true : false;
// _bItalic = font_style & 0x02 ? true ； false;
// _bUnderLine = font_style & 0x03 ? true : false;
type MsgFont struct {
	FontName         *string `protobuf:"bytes,1,opt,name=font_name,json=fontName" json:"font_name,omitempty"`
	FontSize         *uint32 `protobuf:"varint,2,opt,name=font_size,json=fontSize" json:"font_size,omitempty"`
	FontColor        *uint32 `protobuf:"varint,3,opt,name=font_color,json=fontColor" json:"font_color,omitempty"`
	FontStyle        *uint32 `protobuf:"varint,4,opt,name=font_style,json=fontStyle" json:"font_style,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MsgFont) Reset()                    { *m = MsgFont{} }
func (m *MsgFont) String() string            { return proto.CompactTextString(m) }
func (*MsgFont) ProtoMessage()               {}
func (*MsgFont) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *MsgFont) GetFontName() string {
	if m != nil && m.FontName != nil {
		return *m.FontName
	}
	return ""
}

func (m *MsgFont) GetFontSize() uint32 {
	if m != nil && m.FontSize != nil {
		return *m.FontSize
	}
	return 0
}

func (m *MsgFont) GetFontColor() uint32 {
	if m != nil && m.FontColor != nil {
		return *m.FontColor
	}
	return 0
}

func (m *MsgFont) GetFontStyle() uint32 {
	if m != nil && m.FontStyle != nil {
		return *m.FontStyle
	}
	return 0
}

// Msg Data
type MsgData struct {
	MsgId            *int64           `protobuf:"varint,1,opt,name=msg_id,json=msgId" json:"msg_id,omitempty"`
	SrcUserId        *string          `protobuf:"bytes,2,opt,name=src_user_id,json=srcUserId" json:"src_user_id,omitempty"`
	DestUserId       *string          `protobuf:"bytes,3,opt,name=dest_user_id,json=destUserId" json:"dest_user_id,omitempty"`
	GroupId          *uint32          `protobuf:"varint,4,opt,name=group_id,json=groupId" json:"group_id,omitempty"`
	SessionType      *SessionType     `protobuf:"varint,5,opt,name=session_type,json=sessionType,enum=youmai.SessionType" json:"session_type,omitempty"`
	CreateTime       *uint64          `protobuf:"varint,6,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	MsgContent       *string          `protobuf:"bytes,7,opt,name=msg_content,json=msgContent" json:"msg_content,omitempty"`
	ContentType      *IM_CONTENT_TYPE `protobuf:"varint,8,opt,name=content_type,json=contentType,enum=youmai.IM_CONTENT_TYPE" json:"content_type,omitempty"`
	MsgFont          *MsgFont         `protobuf:"bytes,9,opt,name=msg_font,json=msgFont" json:"msg_font,omitempty"`
	MsgStatus        *uint32          `protobuf:"varint,10,opt,name=msg_status,json=msgStatus" json:"msg_status,omitempty"`
	MsgProperty      *MsgProperty     `protobuf:"varint,11,opt,name=msg_property,json=msgProperty,enum=youmai.MsgProperty" json:"msg_property,omitempty"`
	ForcePushIdsList []string         `protobuf:"bytes,12,rep,name=force_push_ids_list,json=forcePushIdsList" json:"force_push_ids_list,omitempty"`
	SrcRealname      *string          `protobuf:"bytes,13,opt,name=src_realname,json=srcRealname" json:"src_realname,omitempty"`
	SrcAvatar        *string          `protobuf:"bytes,14,opt,name=src_avatar,json=srcAvatar" json:"src_avatar,omitempty"`
	SrcMobile        *string          `protobuf:"bytes,15,opt,name=src_mobile,json=srcMobile" json:"src_mobile,omitempty"`
	SrcSex           *string          `protobuf:"bytes,16,opt,name=src_sex,json=srcSex" json:"src_sex,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *MsgData) Reset()                    { *m = MsgData{} }
func (m *MsgData) String() string            { return proto.CompactTextString(m) }
func (*MsgData) ProtoMessage()               {}
func (*MsgData) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *MsgData) GetMsgId() int64 {
	if m != nil && m.MsgId != nil {
		return *m.MsgId
	}
	return 0
}

func (m *MsgData) GetSrcUserId() string {
	if m != nil && m.SrcUserId != nil {
		return *m.SrcUserId
	}
	return ""
}

func (m *MsgData) GetDestUserId() string {
	if m != nil && m.DestUserId != nil {
		return *m.DestUserId
	}
	return ""
}

func (m *MsgData) GetGroupId() uint32 {
	if m != nil && m.GroupId != nil {
		return *m.GroupId
	}
	return 0
}

func (m *MsgData) GetSessionType() SessionType {
	if m != nil && m.SessionType != nil {
		return *m.SessionType
	}
	return SessionType_SESSION_TYPE_SINGLE
}

func (m *MsgData) GetCreateTime() uint64 {
	if m != nil && m.CreateTime != nil {
		return *m.CreateTime
	}
	return 0
}

func (m *MsgData) GetMsgContent() string {
	if m != nil && m.MsgContent != nil {
		return *m.MsgContent
	}
	return ""
}

func (m *MsgData) GetContentType() IM_CONTENT_TYPE {
	if m != nil && m.ContentType != nil {
		return *m.ContentType
	}
	return IM_CONTENT_TYPE_IM_CONTENT_TYPE_TEXT
}

func (m *MsgData) GetMsgFont() *MsgFont {
	if m != nil {
		return m.MsgFont
	}
	return nil
}

func (m *MsgData) GetMsgStatus() uint32 {
	if m != nil && m.MsgStatus != nil {
		return *m.MsgStatus
	}
	return 0
}

func (m *MsgData) GetMsgProperty() MsgProperty {
	if m != nil && m.MsgProperty != nil {
		return *m.MsgProperty
	}
	return MsgProperty_MSG_PTOPERTY_NONE
}

func (m *MsgData) GetForcePushIdsList() []string {
	if m != nil {
		return m.ForcePushIdsList
	}
	return nil
}

func (m *MsgData) GetSrcRealname() string {
	if m != nil && m.SrcRealname != nil {
		return *m.SrcRealname
	}
	return ""
}

func (m *MsgData) GetSrcAvatar() string {
	if m != nil && m.SrcAvatar != nil {
		return *m.SrcAvatar
	}
	return ""
}

func (m *MsgData) GetSrcMobile() string {
	if m != nil && m.SrcMobile != nil {
		return *m.SrcMobile
	}
	return ""
}

func (m *MsgData) GetSrcSex() string {
	if m != nil && m.SrcSex != nil {
		return *m.SrcSex
	}
	return ""
}

// 群消息
type ChatMsg struct {
	Data             *MsgData `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ChatMsg) Reset()                    { *m = ChatMsg{} }
func (m *ChatMsg) String() string            { return proto.CompactTextString(m) }
func (*ChatMsg) ProtoMessage()               {}
func (*ChatMsg) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *ChatMsg) GetData() *MsgData {
	if m != nil {
		return m.Data
	}
	return nil
}

type OfflineMsgNotify struct {
	OfflineMsgList   []*MsgData `protobuf:"bytes,1,rep,name=offline_msg_list,json=offlineMsgList" json:"offline_msg_list,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *OfflineMsgNotify) Reset()                    { *m = OfflineMsgNotify{} }
func (m *OfflineMsgNotify) String() string            { return proto.CompactTextString(m) }
func (*OfflineMsgNotify) ProtoMessage()               {}
func (*OfflineMsgNotify) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *OfflineMsgNotify) GetOfflineMsgList() []*MsgData {
	if m != nil {
		return m.OfflineMsgList
	}
	return nil
}

// 1 server收到sender消息 给sender ack
// 2 rever收到server消息 给server ack
type ChatMsg_Ack struct {
	UserId           *string     `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	MsgId            *int64      `protobuf:"varint,2,opt,name=msg_id,json=msgId" json:"msg_id,omitempty"`
	ErrerNo          *ERRNO_CODE `protobuf:"varint,3,opt,name=errer_no,json=errerNo,enum=youmai.ERRNO_CODE" json:"errer_no,omitempty"`
	IsTargetOnline   *bool       `protobuf:"varint,20,opt,name=is_target_online,json=isTargetOnline" json:"is_target_online,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *ChatMsg_Ack) Reset()                    { *m = ChatMsg_Ack{} }
func (m *ChatMsg_Ack) String() string            { return proto.CompactTextString(m) }
func (*ChatMsg_Ack) ProtoMessage()               {}
func (*ChatMsg_Ack) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *ChatMsg_Ack) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

func (m *ChatMsg_Ack) GetMsgId() int64 {
	if m != nil && m.MsgId != nil {
		return *m.MsgId
	}
	return 0
}

func (m *ChatMsg_Ack) GetErrerNo() ERRNO_CODE {
	if m != nil && m.ErrerNo != nil {
		return *m.ErrerNo
	}
	return ERRNO_CODE_ERRNO_CODE_OK
}

func (m *ChatMsg_Ack) GetIsTargetOnline() bool {
	if m != nil && m.IsTargetOnline != nil {
		return *m.IsTargetOnline
	}
	return false
}

type OfflineMsgAck struct {
	MsgIdList        []int64 `protobuf:"varint,1,rep,name=msg_id_list,json=msgIdList" json:"msg_id_list,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *OfflineMsgAck) Reset()                    { *m = OfflineMsgAck{} }
func (m *OfflineMsgAck) String() string            { return proto.CompactTextString(m) }
func (*OfflineMsgAck) ProtoMessage()               {}
func (*OfflineMsgAck) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func (m *OfflineMsgAck) GetMsgIdList() []int64 {
	if m != nil {
		return m.MsgIdList
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgFont)(nil), "youmai.MsgFont")
	proto.RegisterType((*MsgData)(nil), "youmai.MsgData")
	proto.RegisterType((*ChatMsg)(nil), "youmai.ChatMsg")
	proto.RegisterType((*OfflineMsgNotify)(nil), "youmai.OfflineMsgNotify")
	proto.RegisterType((*ChatMsg_Ack)(nil), "youmai.ChatMsg_Ack")
	proto.RegisterType((*OfflineMsgAck)(nil), "youmai.OfflineMsgAck")
	proto.RegisterEnum("youmai.IM_CONTENT_TYPE", IM_CONTENT_TYPE_name, IM_CONTENT_TYPE_value)
	proto.RegisterEnum("youmai.MsgProperty", MsgProperty_name, MsgProperty_value)
	proto.RegisterEnum("youmai.SessionType", SessionType_name, SessionType_value)
}

func init() { proto.RegisterFile("IM.Msg.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 951 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x94, 0xdd, 0x6e, 0xe2, 0x46,
	0x14, 0xc7, 0xeb, 0x40, 0xf8, 0x38, 0x10, 0x76, 0x3a, 0xf9, 0x72, 0xb2, 0x4d, 0xca, 0xd2, 0x1b,
	0x14, 0x69, 0x53, 0x29, 0x17, 0x95, 0xda, 0x3b, 0xc7, 0x78, 0x89, 0x55, 0x6c, 0xb3, 0x63, 0x53,
	0x35, 0x95, 0xaa, 0x91, 0x0b, 0x13, 0xc7, 0x2a, 0xc6, 0xc8, 0x33, 0x54, 0xcb, 0xde, 0xf5, 0xae,
	0x4f, 0xd0, 0x37, 0xea, 0x73, 0xf4, 0x55, 0xaa, 0x19, 0x1b, 0x12, 0x5c, 0x2e, 0xe7, 0xff, 0x3b,
	0x33, 0xe7, 0x7f, 0x3e, 0x6c, 0x40, 0x8f, 0xe9, 0xca, 0x09, 0xe3, 0x5b, 0x87, 0x47, 0xb7, 0xcb,
	0x2c, 0x15, 0x29, 0xae, 0xad, 0xd3, 0x55, 0x12, 0xc6, 0x97, 0xb8, 0x20, 0xf7, 0x21, 0x8f, 0xa7,
	0x39, 0xeb, 0xfd, 0xa9, 0x41, 0xdd, 0xe1, 0xd1, 0x87, 0x74, 0x21, 0xf0, 0x5b, 0x68, 0x3e, 0xa5,
	0x0b, 0x41, 0x17, 0x61, 0xc2, 0x74, 0xad, 0xab, 0xf5, 0x9b, 0xa4, 0x21, 0x05, 0x37, 0x4c, 0xd8,
	0x16, 0xf2, 0xf8, 0x33, 0xd3, 0x0f, 0xba, 0x5a, 0xff, 0x28, 0x87, 0x7e, 0xfc, 0x99, 0xe1, 0x2b,
	0x00, 0x05, 0xa7, 0xe9, 0x3c, 0xcd, 0xf4, 0x8a, 0xa2, 0x2a, 0xdc, 0x94, 0xc2, 0x16, 0x73, 0xb1,
	0x9e, 0x33, 0xbd, 0xfa, 0x82, 0x7d, 0x29, 0xf4, 0xfe, 0xad, 0x2a, 0x0f, 0x83, 0x50, 0x84, 0xf8,
	0x14, 0x6a, 0x09, 0x8f, 0x68, 0x3c, 0x53, 0x06, 0x2a, 0xe4, 0x30, 0xe1, 0x91, 0x3d, 0xc3, 0xd7,
	0xd0, 0xe2, 0xd9, 0x94, 0xae, 0x38, 0xcb, 0x24, 0x3b, 0x50, 0xe6, 0x9a, 0x3c, 0x9b, 0x4e, 0x38,
	0xcb, 0xec, 0x19, 0xee, 0x42, 0x7b, 0xc6, 0xb8, 0xd8, 0x06, 0x54, 0x54, 0x00, 0x48, 0xad, 0x88,
	0xb8, 0x80, 0x46, 0x94, 0xa5, 0xab, 0xa5, 0xa4, 0xb9, 0x83, 0xba, 0x3a, 0xdb, 0x33, 0xfc, 0x1d,
	0xb4, 0x39, 0xe3, 0x3c, 0x4e, 0x17, 0x54, 0xac, 0x97, 0x4c, 0x3f, 0xec, 0x6a, 0xfd, 0xce, 0xdd,
	0xf1, 0x6d, 0xde, 0xb6, 0x5b, 0x3f, 0x67, 0xc1, 0x7a, 0xc9, 0x48, 0x8b, 0xbf, 0x1c, 0xf0, 0xd7,
	0xd0, 0x9a, 0x66, 0x2c, 0x14, 0x8c, 0x8a, 0x38, 0x61, 0x7a, 0xad, 0xab, 0xf5, 0xab, 0x04, 0x72,
	0x29, 0x88, 0x13, 0x15, 0x20, 0x8b, 0x99, 0xa6, 0x0b, 0xc1, 0x16, 0x42, 0xaf, 0xe7, 0xa6, 0x12,
	0x1e, 0x99, 0xb9, 0x82, 0x7f, 0x80, 0x76, 0x01, 0xf3, 0xcc, 0x0d, 0x95, 0xf9, 0x7c, 0x93, 0xd9,
	0x76, 0xa8, 0xe9, 0xb9, 0x81, 0xe5, 0x06, 0x34, 0x78, 0x1c, 0x5b, 0xa4, 0x55, 0x04, 0xab, 0xec,
	0x37, 0xd0, 0x90, 0x8f, 0xcb, 0x36, 0xea, 0xcd, 0xae, 0xd6, 0x6f, 0xdd, 0xbd, 0xd9, 0xdc, 0x2b,
	0x06, 0x4a, 0xea, 0x49, 0x31, 0xd9, 0x2b, 0x90, 0x59, 0x29, 0x17, 0xa1, 0x58, 0x71, 0x1d, 0xf2,
	0x01, 0x24, 0x3c, 0xf2, 0x95, 0x20, 0x1b, 0x20, 0xf1, 0x32, 0x4b, 0x97, 0x2c, 0x13, 0x6b, 0xbd,
	0xb5, 0xdb, 0x00, 0x87, 0x47, 0xe3, 0x02, 0x11, 0x59, 0xd0, 0xe6, 0x80, 0xdf, 0xc3, 0xf1, 0x53,
	0x9a, 0x4d, 0x19, 0x5d, 0xae, 0xf8, 0x33, 0x8d, 0x67, 0x9c, 0xce, 0x63, 0x2e, 0xf4, 0x76, 0xb7,
	0xd2, 0x6f, 0x12, 0xa4, 0xd0, 0x78, 0xc5, 0x9f, 0xed, 0x19, 0x1f, 0xc5, 0x5c, 0xe0, 0x77, 0xd0,
	0x96, 0x43, 0xcc, 0x58, 0x38, 0x57, 0x2b, 0x76, 0xa4, 0xfa, 0x21, 0x07, 0x4b, 0x0a, 0x49, 0x1a,
	0x95, 0x21, 0xe1, 0x1f, 0xa1, 0x08, 0x33, 0xbd, 0xb3, 0x1d, 0xb3, 0xa1, 0x84, 0x0d, 0x4e, 0xd2,
	0xdf, 0xe2, 0x39, 0xd3, 0xdf, 0x6c, 0xb1, 0xa3, 0x04, 0x7c, 0x0e, 0x75, 0x89, 0x39, 0xfb, 0xa4,
	0x23, 0xc5, 0x6a, 0x3c, 0x9b, 0xfa, 0xec, 0x53, 0xef, 0x16, 0xea, 0xe6, 0x73, 0x28, 0x1c, 0x1e,
	0xe1, 0x6f, 0xa0, 0x3a, 0x0b, 0x45, 0xa8, 0xd6, 0x6b, 0xb7, 0x65, 0x72, 0xff, 0x88, 0x82, 0x3d,
	0x07, 0x90, 0xf7, 0xf4, 0x34, 0x8f, 0x17, 0xcc, 0xe1, 0x91, 0x9b, 0x8a, 0xf8, 0x69, 0x8d, 0xbf,
	0x07, 0x94, 0xe6, 0x1a, 0x95, 0xcd, 0x52, 0x95, 0x6a, 0xdd, 0xca, 0xbe, 0x47, 0x3a, 0xe9, 0xf6,
	0xb2, 0x2c, 0xbc, 0xf7, 0xb7, 0x06, 0xad, 0x22, 0x3f, 0x35, 0xa6, 0xbf, 0x4b, 0x9f, 0x9b, 0x45,
	0xcd, 0x3f, 0xb3, 0xda, 0x2a, 0x5f, 0xd2, 0x97, 0xed, 0x3f, 0x78, 0xbd, 0xfd, 0xef, 0xa1, 0xc1,
	0xb2, 0x8c, 0x65, 0x74, 0x91, 0xaa, 0xcd, 0xee, 0xdc, 0xe1, 0x4d, 0x4a, 0x8b, 0x10, 0xd7, 0xa3,
	0xa6, 0x37, 0xb0, 0x48, 0x5d, 0xc5, 0xb8, 0x29, 0xee, 0x03, 0x8a, 0x39, 0x15, 0x61, 0x16, 0x31,
	0x41, 0xd3, 0x85, 0x74, 0xa2, 0x9f, 0x74, 0xb5, 0x7e, 0x83, 0x74, 0x62, 0x1e, 0x28, 0xd9, 0x53,
	0x6a, 0xef, 0x5b, 0x38, 0x7a, 0xa9, 0x53, 0x3a, 0xbb, 0xce, 0x37, 0x36, 0x9e, 0xbd, 0xd4, 0x57,
	0x51, 0x9b, 0x62, 0xcf, 0x64, 0x25, 0x37, 0xff, 0x54, 0xe0, 0x4d, 0x69, 0x2b, 0xb1, 0x0e, 0x27,
	0x25, 0x89, 0x06, 0xd6, 0xcf, 0x01, 0xfa, 0x02, 0xbf, 0x83, 0xab, 0x32, 0xf1, 0x1f, 0x3c, 0x12,
	0x50, 0xc7, 0xf2, 0x7d, 0x63, 0x68, 0x21, 0x0d, 0x7f, 0x05, 0x7a, 0x39, 0x44, 0x1e, 0x0c, 0x33,
	0xf0, 0xd1, 0xc1, 0xbe, 0x07, 0x88, 0x65, 0x7a, 0x8e, 0x63, 0xb9, 0x03, 0x6a, 0x8c, 0xc7, 0xa8,
	0x82, 0xaf, 0xe1, 0xb2, 0x1c, 0xe2, 0x7a, 0x74, 0x60, 0xfb, 0xc1, 0x84, 0xdc, 0xa3, 0xea, 0xbe,
	0x27, 0x46, 0x9e, 0x69, 0x04, 0xb6, 0xe7, 0xfa, 0x0f, 0x06, 0xb1, 0xd0, 0x21, 0xbe, 0x80, 0xd3,
	0x72, 0x88, 0xed, 0x48, 0x7b, 0x35, 0x7c, 0x06, 0xb8, 0x8c, 0x8c, 0x00, 0xd5, 0xf1, 0x39, 0x1c,
	0x97, 0xf5, 0x09, 0x19, 0xa1, 0xc6, 0xbe, 0xb7, 0x8c, 0xc9, 0xc0, 0xf6, 0x50, 0x73, 0x1f, 0xfa,
	0xc9, 0x1e, 0x58, 0x1e, 0x82, 0x7d, 0x5d, 0xd8, 0x98, 0x44, 0xad, 0x7d, 0x0d, 0xfe, 0x60, 0x8f,
	0x2c, 0xd4, 0xc6, 0x6f, 0xe1, 0xbc, 0x4c, 0xee, 0xed, 0x5f, 0x4c, 0x83, 0x0c, 0xd0, 0x11, 0xbe,
	0x84, 0xb3, 0x32, 0xfc, 0xd1, 0xf8, 0x38, 0x31, 0x5c, 0xd4, 0xb9, 0xf9, 0x4b, 0x83, 0xd6, 0xab,
	0xcf, 0x1a, 0x9f, 0xc2, 0x97, 0x8e, 0x3f, 0xa4, 0xe3, 0xc0, 0x1b, 0x5b, 0x24, 0x78, 0xa4, 0xae,
	0xe7, 0xca, 0xe9, 0xe8, 0x70, 0xa2, 0x64, 0x52, 0xc8, 0xc4, 0x32, 0x2d, 0x7b, 0x1c, 0xa0, 0x03,
	0xd9, 0x80, 0x1d, 0x32, 0x21, 0x43, 0xcb, 0x0d, 0x50, 0x45, 0x76, 0x6c, 0x07, 0x7c, 0x9c, 0x78,
	0x81, 0x85, 0xaa, 0xff, 0xbb, 0x40, 0x2c, 0xd3, 0x18, 0x8d, 0xd0, 0xe1, 0xcd, 0xaf, 0xd0, 0x7a,
	0xf5, 0x87, 0x95, 0x71, 0xbe, 0xe5, 0xfb, 0xb6, 0xe7, 0x16, 0x0b, 0x63, 0xbb, 0xc3, 0x91, 0xf4,
	0x72, 0x09, 0x67, 0x3b, 0xc0, 0x99, 0x8c, 0x02, 0xdb, 0x7c, 0x30, 0xa4, 0x9b, 0x0b, 0x38, 0xdd,
	0x61, 0x1e, 0x19, 0x0e, 0x89, 0x37, 0x19, 0xa3, 0xca, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x20,
	0xeb, 0x0f, 0x13, 0x0f, 0x07, 0x00, 0x00,
}
