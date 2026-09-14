package etcdserverpb
import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	authpb "go.etcd.io/etcd/api/v3/authpb"
	mvccpb "go.etcd.io/etcd/api/v3/mvccpb"
	_ "go.etcd.io/etcd/api/v3/versionpb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)
const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)
type AlarmType int32
const (
	AlarmType_NONE    AlarmType = 0 
	AlarmType_NOSPACE AlarmType = 1 
	AlarmType_CORRUPT AlarmType = 2 
)
var (
	AlarmType_name = map[int32]string{
		0: "NONE",
		1: "NOSPACE",
		2: "CORRUPT",
	}
	AlarmType_value = map[string]int32{
		"NONE":    0,
		"NOSPACE": 1,
		"CORRUPT": 2,
	}
)
func (x AlarmType) Enum() *AlarmType {
	p := new(AlarmType)
	*p = x
	return p
}
func (x AlarmType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}
func (AlarmType) Descriptor() protoreflect.EnumDescriptor {
	return file_rpc_proto_enumTypes[0].Descriptor()
}
func (AlarmType) Type() protoreflect.EnumType {
	return &file_rpc_proto_enumTypes[0]
}
func (x AlarmType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}
func (AlarmType) EnumDescriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{0}
}
type RangeRequest_SortOrder int32
const (
	RangeRequest_NONE    RangeRequest_SortOrder = 0 
	RangeRequest_ASCEND  RangeRequest_SortOrder = 1 
	RangeRequest_DESCEND RangeRequest_SortOrder = 2 
)
var (
	RangeRequest_SortOrder_name = map[int32]string{
		0: "NONE",
		1: "ASCEND",
		2: "DESCEND",
	}
	RangeRequest_SortOrder_value = map[string]int32{
		"NONE":    0,
		"ASCEND":  1,
		"DESCEND": 2,
	}
)
func (x RangeRequest_SortOrder) Enum() *RangeRequest_SortOrder {
	p := new(RangeRequest_SortOrder)
	*p = x
	return p
}
func (x RangeRequest_SortOrder) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}
func (RangeRequest_SortOrder) Descriptor() protoreflect.EnumDescriptor {
	return file_rpc_proto_enumTypes[1].Descriptor()
}
func (RangeRequest_SortOrder) Type() protoreflect.EnumType {
	return &file_rpc_proto_enumTypes[1]
}
func (x RangeRequest_SortOrder) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}
func (RangeRequest_SortOrder) EnumDescriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{1, 0}
}
type RangeRequest_SortTarget int32
const (
	RangeRequest_KEY     RangeRequest_SortTarget = 0
	RangeRequest_VERSION RangeRequest_SortTarget = 1
	RangeRequest_CREATE  RangeRequest_SortTarget = 2
	RangeRequest_MOD     RangeRequest_SortTarget = 3
	RangeRequest_VALUE   RangeRequest_SortTarget = 4
)
var (
	RangeRequest_SortTarget_name = map[int32]string{
		0: "KEY",
		1: "VERSION",
		2: "CREATE",
		3: "MOD",
		4: "VALUE",
	}
	RangeRequest_SortTarget_value = map[string]int32{
		"KEY":     0,
		"VERSION": 1,
		"CREATE":  2,
		"MOD":     3,
		"VALUE":   4,
	}
)
func (x RangeRequest_SortTarget) Enum() *RangeRequest_SortTarget {
	p := new(RangeRequest_SortTarget)
	*p = x
	return p
}
func (x RangeRequest_SortTarget) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}
func (RangeRequest_SortTarget) Descriptor() protoreflect.EnumDescriptor {
	return file_rpc_proto_enumTypes[2].Descriptor()
}
func (RangeRequest_SortTarget) Type() protoreflect.EnumType {
	return &file_rpc_proto_enumTypes[2]
}
func (x RangeRequest_SortTarget) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}
func (RangeRequest_SortTarget) EnumDescriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{1, 1}
}
type Compare_CompareResult int32
const (
	Compare_EQUAL     Compare_CompareResult = 0
	Compare_GREATER   Compare_CompareResult = 1
	Compare_LESS      Compare_CompareResult = 2
	Compare_NOT_EQUAL Compare_CompareResult = 3
)
var (
	Compare_CompareResult_name = map[int32]string{
		0: "EQUAL",
		1: "GREATER",
		2: "LESS",
		3: "NOT_EQUAL",
	}
	Compare_CompareResult_value = map[string]int32{
		"EQUAL":     0,
		"GREATER":   1,
		"LESS":      2,
		"NOT_EQUAL": 3,
	}
)
func (x Compare_CompareResult) Enum() *Compare_CompareResult {
	p := new(Compare_CompareResult)
	*p = x
	return p
}
func (x Compare_CompareResult) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}
func (Compare_CompareResult) Descriptor() protoreflect.EnumDescriptor {
	return file_rpc_proto_enumTypes[3].Descriptor()
}
func (Compare_CompareResult) Type() protoreflect.EnumType {
	return &file_rpc_proto_enumTypes[3]
}
func (x Compare_CompareResult) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}
func (Compare_CompareResult) EnumDescriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{9, 0}
}
type Compare_CompareTarget int32
const (
	Compare_VERSION Compare_CompareTarget = 0
	Compare_CREATE  Compare_CompareTarget = 1
	Compare_MOD     Compare_CompareTarget = 2
	Compare_VALUE   Compare_CompareTarget = 3
	Compare_LEASE   Compare_CompareTarget = 4
)
var (
	Compare_CompareTarget_name = map[int32]string{
		0: "VERSION",
		1: "CREATE",
		2: "MOD",
		3: "VALUE",
		4: "LEASE",
	}
	Compare_CompareTarget_value = map[string]int32{
		"VERSION": 0,
		"CREATE":  1,
		"MOD":     2,
		"VALUE":   3,
		"LEASE":   4,
	}
)
func (x Compare_CompareTarget) Enum() *Compare_CompareTarget {
	p := new(Compare_CompareTarget)
	*p = x
	return p
}
func (x Compare_CompareTarget) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}
func (Compare_CompareTarget) Descriptor() protoreflect.EnumDescriptor {
	return file_rpc_proto_enumTypes[4].Descriptor()
}
func (Compare_CompareTarget) Type() protoreflect.EnumType {
	return &file_rpc_proto_enumTypes[4]
}
func (x Compare_CompareTarget) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}
func (Compare_CompareTarget) EnumDescriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{9, 1}
}
type WatchCreateRequest_FilterType int32
const (
	WatchCreateRequest_NOPUT WatchCreateRequest_FilterType = 0
	WatchCreateRequest_NODELETE WatchCreateRequest_FilterType = 1
)
var (
	WatchCreateRequest_FilterType_name = map[int32]string{
		0: "NOPUT",
		1: "NODELETE",
	}
	WatchCreateRequest_FilterType_value = map[string]int32{
		"NOPUT":    0,
		"NODELETE": 1,
	}
)
func (x WatchCreateRequest_FilterType) Enum() *WatchCreateRequest_FilterType {
	p := new(WatchCreateRequest_FilterType)
	*p = x
	return p
}
func (x WatchCreateRequest_FilterType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}
func (WatchCreateRequest_FilterType) Descriptor() protoreflect.EnumDescriptor {
	return file_rpc_proto_enumTypes[5].Descriptor()
}
func (WatchCreateRequest_FilterType) Type() protoreflect.EnumType {
	return &file_rpc_proto_enumTypes[5]
}
func (x WatchCreateRequest_FilterType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}
func (WatchCreateRequest_FilterType) EnumDescriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{21, 0}
}
type AlarmRequest_AlarmAction int32
const (
	AlarmRequest_GET        AlarmRequest_AlarmAction = 0
	AlarmRequest_ACTIVATE   AlarmRequest_AlarmAction = 1
	AlarmRequest_DEACTIVATE AlarmRequest_AlarmAction = 2
)
var (
	AlarmRequest_AlarmAction_name = map[int32]string{
		0: "GET",
		1: "ACTIVATE",
		2: "DEACTIVATE",
	}
	AlarmRequest_AlarmAction_value = map[string]int32{
		"GET":        0,
		"ACTIVATE":   1,
		"DEACTIVATE": 2,
	}
)
func (x AlarmRequest_AlarmAction) Enum() *AlarmRequest_AlarmAction {
	p := new(AlarmRequest_AlarmAction)
	*p = x
	return p
}
func (x AlarmRequest_AlarmAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}
func (AlarmRequest_AlarmAction) Descriptor() protoreflect.EnumDescriptor {
	return file_rpc_proto_enumTypes[6].Descriptor()
}
func (AlarmRequest_AlarmAction) Type() protoreflect.EnumType {
	return &file_rpc_proto_enumTypes[6]
}
func (x AlarmRequest_AlarmAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}
func (AlarmRequest_AlarmAction) EnumDescriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{54, 0}
}
type DowngradeRequest_DowngradeAction int32
const (
	DowngradeRequest_VALIDATE DowngradeRequest_DowngradeAction = 0
	DowngradeRequest_ENABLE   DowngradeRequest_DowngradeAction = 1
	DowngradeRequest_CANCEL   DowngradeRequest_DowngradeAction = 2
)
var (
	DowngradeRequest_DowngradeAction_name = map[int32]string{
		0: "VALIDATE",
		1: "ENABLE",
		2: "CANCEL",
	}
	DowngradeRequest_DowngradeAction_value = map[string]int32{
		"VALIDATE": 0,
		"ENABLE":   1,
		"CANCEL":   2,
	}
)
func (x DowngradeRequest_DowngradeAction) Enum() *DowngradeRequest_DowngradeAction {
	p := new(DowngradeRequest_DowngradeAction)
	*p = x
	return p
}
func (x DowngradeRequest_DowngradeAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}
func (DowngradeRequest_DowngradeAction) Descriptor() protoreflect.EnumDescriptor {
	return file_rpc_proto_enumTypes[7].Descriptor()
}
func (DowngradeRequest_DowngradeAction) Type() protoreflect.EnumType {
	return &file_rpc_proto_enumTypes[7]
}
func (x DowngradeRequest_DowngradeAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}
func (DowngradeRequest_DowngradeAction) EnumDescriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{57, 0}
}
type ResponseHeader struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	ClusterId uint64 `protobuf:"varint,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	MemberId uint64 `protobuf:"varint,2,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	Revision int64 `protobuf:"varint,3,opt,name=revision,proto3" json:"revision,omitempty"`
	RaftTerm      uint64 `protobuf:"varint,4,opt,name=raft_term,json=raftTerm,proto3" json:"raft_term,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *ResponseHeader) Reset() {
	*x = ResponseHeader{}
	mi := &file_rpc_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *ResponseHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*ResponseHeader) ProtoMessage() {}
func (x *ResponseHeader) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*ResponseHeader) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{0}
}
func (x *ResponseHeader) GetClusterId() uint64 {
	if x != nil {
		return x.ClusterId
	}
	return 0
}
func (x *ResponseHeader) GetMemberId() uint64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}
func (x *ResponseHeader) GetRevision() int64 {
	if x != nil {
		return x.Revision
	}
	return 0
}
func (x *ResponseHeader) GetRaftTerm() uint64 {
	if x != nil {
		return x.RaftTerm
	}
	return 0
}
type RangeRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	RangeEnd []byte `protobuf:"bytes,2,opt,name=range_end,json=rangeEnd,proto3" json:"range_end,omitempty"`
	Limit int64 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Revision int64 `protobuf:"varint,4,opt,name=revision,proto3" json:"revision,omitempty"`
	SortOrder RangeRequest_SortOrder `protobuf:"varint,5,opt,name=sort_order,json=sortOrder,proto3,enum=etcdserverpb.RangeRequest_SortOrder" json:"sort_order,omitempty"`
	SortTarget RangeRequest_SortTarget `protobuf:"varint,6,opt,name=sort_target,json=sortTarget,proto3,enum=etcdserverpb.RangeRequest_SortTarget" json:"sort_target,omitempty"`
	Serializable bool `protobuf:"varint,7,opt,name=serializable,proto3" json:"serializable,omitempty"`
	KeysOnly bool `protobuf:"varint,8,opt,name=keys_only,json=keysOnly,proto3" json:"keys_only,omitempty"`
	CountOnly bool `protobuf:"varint,9,opt,name=count_only,json=countOnly,proto3" json:"count_only,omitempty"`
	MinModRevision int64 `protobuf:"varint,10,opt,name=min_mod_revision,json=minModRevision,proto3" json:"min_mod_revision,omitempty"`
	MaxModRevision int64 `protobuf:"varint,11,opt,name=max_mod_revision,json=maxModRevision,proto3" json:"max_mod_revision,omitempty"`
	MinCreateRevision int64 `protobuf:"varint,12,opt,name=min_create_revision,json=minCreateRevision,proto3" json:"min_create_revision,omitempty"`
	MaxCreateRevision int64 `protobuf:"varint,13,opt,name=max_create_revision,json=maxCreateRevision,proto3" json:"max_create_revision,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}
func (x *RangeRequest) Reset() {
	*x = RangeRequest{}
	mi := &file_rpc_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *RangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*RangeRequest) ProtoMessage() {}
func (x *RangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*RangeRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{1}
}
func (x *RangeRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}
func (x *RangeRequest) GetRangeEnd() []byte {
	if x != nil {
		return x.RangeEnd
	}
	return nil
}
func (x *RangeRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}
func (x *RangeRequest) GetRevision() int64 {
	if x != nil {
		return x.Revision
	}
	return 0
}
func (x *RangeRequest) GetSortOrder() RangeRequest_SortOrder {
	if x != nil {
		return x.SortOrder
	}
	return RangeRequest_NONE
}
func (x *RangeRequest) GetSortTarget() RangeRequest_SortTarget {
	if x != nil {
		return x.SortTarget
	}
	return RangeRequest_KEY
}
func (x *RangeRequest) GetSerializable() bool {
	if x != nil {
		return x.Serializable
	}
	return false
}
func (x *RangeRequest) GetKeysOnly() bool {
	if x != nil {
		return x.KeysOnly
	}
	return false
}
func (x *RangeRequest) GetCountOnly() bool {
	if x != nil {
		return x.CountOnly
	}
	return false
}
func (x *RangeRequest) GetMinModRevision() int64 {
	if x != nil {
		return x.MinModRevision
	}
	return 0
}
func (x *RangeRequest) GetMaxModRevision() int64 {
	if x != nil {
		return x.MaxModRevision
	}
	return 0
}
func (x *RangeRequest) GetMinCreateRevision() int64 {
	if x != nil {
		return x.MinCreateRevision
	}
	return 0
}
func (x *RangeRequest) GetMaxCreateRevision() int64 {
	if x != nil {
		return x.MaxCreateRevision
	}
	return 0
}
type RangeResponse struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	Header *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Kvs []*mvccpb.KeyValue `protobuf:"bytes,2,rep,name=kvs,proto3" json:"kvs,omitempty"`
	More bool `protobuf:"varint,3,opt,name=more,proto3" json:"more,omitempty"`
	Count         int64 `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *RangeResponse) Reset() {
	*x = RangeResponse{}
	mi := &file_rpc_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *RangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*RangeResponse) ProtoMessage() {}
func (x *RangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*RangeResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{2}
}
func (x *RangeResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
func (x *RangeResponse) GetKvs() []*mvccpb.KeyValue {
	if x != nil {
		return x.Kvs
	}
	return nil
}
func (x *RangeResponse) GetMore() bool {
	if x != nil {
		return x.More
	}
	return false
}
func (x *RangeResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}
type PutRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Lease int64 `protobuf:"varint,3,opt,name=lease,proto3" json:"lease,omitempty"`
	PrevKv bool `protobuf:"varint,4,opt,name=prev_kv,json=prevKv,proto3" json:"prev_kv,omitempty"`
	IgnoreValue bool `protobuf:"varint,5,opt,name=ignore_value,json=ignoreValue,proto3" json:"ignore_value,omitempty"`
	IgnoreLease   bool `protobuf:"varint,6,opt,name=ignore_lease,json=ignoreLease,proto3" json:"ignore_lease,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *PutRequest) Reset() {
	*x = PutRequest{}
	mi := &file_rpc_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *PutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*PutRequest) ProtoMessage() {}
func (x *PutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*PutRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{3}
}
func (x *PutRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}
func (x *PutRequest) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}
func (x *PutRequest) GetLease() int64 {
	if x != nil {
		return x.Lease
	}
	return 0
}
func (x *PutRequest) GetPrevKv() bool {
	if x != nil {
		return x.PrevKv
	}
	return false
}
func (x *PutRequest) GetIgnoreValue() bool {
	if x != nil {
		return x.IgnoreValue
	}
	return false
}
func (x *PutRequest) GetIgnoreLease() bool {
	if x != nil {
		return x.IgnoreLease
	}
	return false
}
type PutResponse struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	Header *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	PrevKv        *mvccpb.KeyValue `protobuf:"bytes,2,opt,name=prev_kv,json=prevKv,proto3" json:"prev_kv,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *PutResponse) Reset() {
	*x = PutResponse{}
	mi := &file_rpc_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *PutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*PutResponse) ProtoMessage() {}
func (x *PutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*PutResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{4}
}
func (x *PutResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
func (x *PutResponse) GetPrevKv() *mvccpb.KeyValue {
	if x != nil {
		return x.PrevKv
	}
	return nil
}
type DeleteRangeRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	RangeEnd []byte `protobuf:"bytes,2,opt,name=range_end,json=rangeEnd,proto3" json:"range_end,omitempty"`
	PrevKv        bool `protobuf:"varint,3,opt,name=prev_kv,json=prevKv,proto3" json:"prev_kv,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *DeleteRangeRequest) Reset() {
	*x = DeleteRangeRequest{}
	mi := &file_rpc_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *DeleteRangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*DeleteRangeRequest) ProtoMessage() {}
func (x *DeleteRangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*DeleteRangeRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{5}
}
func (x *DeleteRangeRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}
func (x *DeleteRangeRequest) GetRangeEnd() []byte {
	if x != nil {
		return x.RangeEnd
	}
	return nil
}
func (x *DeleteRangeRequest) GetPrevKv() bool {
	if x != nil {
		return x.PrevKv
	}
	return false
}
type DeleteRangeResponse struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	Header *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Deleted int64 `protobuf:"varint,2,opt,name=deleted,proto3" json:"deleted,omitempty"`
	PrevKvs       []*mvccpb.KeyValue `protobuf:"bytes,3,rep,name=prev_kvs,json=prevKvs,proto3" json:"prev_kvs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *DeleteRangeResponse) Reset() {
	*x = DeleteRangeResponse{}
	mi := &file_rpc_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *DeleteRangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*DeleteRangeResponse) ProtoMessage() {}
func (x *DeleteRangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*DeleteRangeResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{6}
}
func (x *DeleteRangeResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
func (x *DeleteRangeResponse) GetDeleted() int64 {
	if x != nil {
		return x.Deleted
	}
	return 0
}
func (x *DeleteRangeResponse) GetPrevKvs() []*mvccpb.KeyValue {
	if x != nil {
		return x.PrevKvs
	}
	return nil
}
type RequestOp struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Request       isRequestOp_Request `protobuf_oneof:"request"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *RequestOp) Reset() {
	*x = RequestOp{}
	mi := &file_rpc_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *RequestOp) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*RequestOp) ProtoMessage() {}
func (x *RequestOp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*RequestOp) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{7}
}
func (x *RequestOp) GetRequest() isRequestOp_Request {
	if x != nil {
		return x.Request
	}
	return nil
}
func (x *RequestOp) GetRequestRange() *RangeRequest {
	if x != nil {
		if x, ok := x.Request.(*RequestOp_RequestRange); ok {
			return x.RequestRange
		}
	}
	return nil
}
func (x *RequestOp) GetRequestPut() *PutRequest {
	if x != nil {
		if x, ok := x.Request.(*RequestOp_RequestPut); ok {
			return x.RequestPut
		}
	}
	return nil
}
func (x *RequestOp) GetRequestDeleteRange() *DeleteRangeRequest {
	if x != nil {
		if x, ok := x.Request.(*RequestOp_RequestDeleteRange); ok {
			return x.RequestDeleteRange
		}
	}
	return nil
}
func (x *RequestOp) GetRequestTxn() *TxnRequest {
	if x != nil {
		if x, ok := x.Request.(*RequestOp_RequestTxn); ok {
			return x.RequestTxn
		}
	}
	return nil
}
type isRequestOp_Request interface {
	isRequestOp_Request()
}
type RequestOp_RequestRange struct {
	RequestRange *RangeRequest `protobuf:"bytes,1,opt,name=request_range,json=requestRange,proto3,oneof"`
}
type RequestOp_RequestPut struct {
	RequestPut *PutRequest `protobuf:"bytes,2,opt,name=request_put,json=requestPut,proto3,oneof"`
}
type RequestOp_RequestDeleteRange struct {
	RequestDeleteRange *DeleteRangeRequest `protobuf:"bytes,3,opt,name=request_delete_range,json=requestDeleteRange,proto3,oneof"`
}
type RequestOp_RequestTxn struct {
	RequestTxn *TxnRequest `protobuf:"bytes,4,opt,name=request_txn,json=requestTxn,proto3,oneof"`
}
func (*RequestOp_RequestRange) isRequestOp_Request() {}
func (*RequestOp_RequestPut) isRequestOp_Request() {}
func (*RequestOp_RequestDeleteRange) isRequestOp_Request() {}
func (*RequestOp_RequestTxn) isRequestOp_Request() {}
type ResponseOp struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Response      isResponseOp_Response `protobuf_oneof:"response"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *ResponseOp) Reset() {
	*x = ResponseOp{}
	mi := &file_rpc_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *ResponseOp) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*ResponseOp) ProtoMessage() {}
func (x *ResponseOp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*ResponseOp) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{8}
}
func (x *ResponseOp) GetResponse() isResponseOp_Response {
	if x != nil {
		return x.Response
	}
	return nil
}
func (x *ResponseOp) GetResponseRange() *RangeResponse {
	if x != nil {
		if x, ok := x.Response.(*ResponseOp_ResponseRange); ok {
			return x.ResponseRange
		}
	}
	return nil
}
func (x *ResponseOp) GetResponsePut() *PutResponse {
	if x != nil {
		if x, ok := x.Response.(*ResponseOp_ResponsePut); ok {
			return x.ResponsePut
		}
	}
	return nil
}
func (x *ResponseOp) GetResponseDeleteRange() *DeleteRangeResponse {
	if x != nil {
		if x, ok := x.Response.(*ResponseOp_ResponseDeleteRange); ok {
			return x.ResponseDeleteRange
		}
	}
	return nil
}
func (x *ResponseOp) GetResponseTxn() *TxnResponse {
	if x != nil {
		if x, ok := x.Response.(*ResponseOp_ResponseTxn); ok {
			return x.ResponseTxn
		}
	}
	return nil
}
type isResponseOp_Response interface {
	isResponseOp_Response()
}
type ResponseOp_ResponseRange struct {
	ResponseRange *RangeResponse `protobuf:"bytes,1,opt,name=response_range,json=responseRange,proto3,oneof"`
}
type ResponseOp_ResponsePut struct {
	ResponsePut *PutResponse `protobuf:"bytes,2,opt,name=response_put,json=responsePut,proto3,oneof"`
}
type ResponseOp_ResponseDeleteRange struct {
	ResponseDeleteRange *DeleteRangeResponse `protobuf:"bytes,3,opt,name=response_delete_range,json=responseDeleteRange,proto3,oneof"`
}
type ResponseOp_ResponseTxn struct {
	ResponseTxn *TxnResponse `protobuf:"bytes,4,opt,name=response_txn,json=responseTxn,proto3,oneof"`
}
func (*ResponseOp_ResponseRange) isResponseOp_Response() {}
func (*ResponseOp_ResponsePut) isResponseOp_Response() {}
func (*ResponseOp_ResponseDeleteRange) isResponseOp_Response() {}
func (*ResponseOp_ResponseTxn) isResponseOp_Response() {}
type Compare struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Result Compare_CompareResult `protobuf:"varint,1,opt,name=result,proto3,enum=etcdserverpb.Compare_CompareResult" json:"result,omitempty"`
	Target Compare_CompareTarget `protobuf:"varint,2,opt,name=target,proto3,enum=etcdserverpb.Compare_CompareTarget" json:"target,omitempty"`
	Key []byte `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	TargetUnion isCompare_TargetUnion `protobuf_oneof:"target_union"`
	RangeEnd      []byte `protobuf:"bytes,64,opt,name=range_end,json=rangeEnd,proto3" json:"range_end,omitempty"` 
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *Compare) Reset() {
	*x = Compare{}
	mi := &file_rpc_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *Compare) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*Compare) ProtoMessage() {}
func (x *Compare) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*Compare) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{9}
}
func (x *Compare) GetResult() Compare_CompareResult {
	if x != nil {
		return x.Result
	}
	return Compare_EQUAL
}
func (x *Compare) GetTarget() Compare_CompareTarget {
	if x != nil {
		return x.Target
	}
	return Compare_VERSION
}
func (x *Compare) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}
func (x *Compare) GetTargetUnion() isCompare_TargetUnion {
	if x != nil {
		return x.TargetUnion
	}
	return nil
}
func (x *Compare) GetVersion() int64 {
	if x != nil {
		if x, ok := x.TargetUnion.(*Compare_Version); ok {
			return x.Version
		}
	}
	return 0
}
func (x *Compare) GetCreateRevision() int64 {
	if x != nil {
		if x, ok := x.TargetUnion.(*Compare_CreateRevision); ok {
			return x.CreateRevision
		}
	}
	return 0
}
func (x *Compare) GetModRevision() int64 {
	if x != nil {
		if x, ok := x.TargetUnion.(*Compare_ModRevision); ok {
			return x.ModRevision
		}
	}
	return 0
}
func (x *Compare) GetValue() []byte {
	if x != nil {
		if x, ok := x.TargetUnion.(*Compare_Value); ok {
			return x.Value
		}
	}
	return nil
}
func (x *Compare) GetLease() int64 {
	if x != nil {
		if x, ok := x.TargetUnion.(*Compare_Lease); ok {
			return x.Lease
		}
	}
	return 0
}
func (x *Compare) GetRangeEnd() []byte {
	if x != nil {
		return x.RangeEnd
	}
	return nil
}
type isCompare_TargetUnion interface {
	isCompare_TargetUnion()
}
type Compare_Version struct {
	Version int64 `protobuf:"varint,4,opt,name=version,proto3,oneof"`
}
type Compare_CreateRevision struct {
	CreateRevision int64 `protobuf:"varint,5,opt,name=create_revision,json=createRevision,proto3,oneof"`
}
type Compare_ModRevision struct {
	ModRevision int64 `protobuf:"varint,6,opt,name=mod_revision,json=modRevision,proto3,oneof"`
}
type Compare_Value struct {
	Value []byte `protobuf:"bytes,7,opt,name=value,proto3,oneof"`
}
type Compare_Lease struct {
	Lease int64 `protobuf:"varint,8,opt,name=lease,proto3,oneof"` 
}
func (*Compare_Version) isCompare_TargetUnion() {}
func (*Compare_CreateRevision) isCompare_TargetUnion() {}
func (*Compare_ModRevision) isCompare_TargetUnion() {}
func (*Compare_Value) isCompare_TargetUnion() {}
func (*Compare_Lease) isCompare_TargetUnion() {}
type TxnRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Compare []*Compare `protobuf:"bytes,1,rep,name=compare,proto3" json:"compare,omitempty"`
	Success []*RequestOp `protobuf:"bytes,2,rep,name=success,proto3" json:"success,omitempty"`
	Failure       []*RequestOp `protobuf:"bytes,3,rep,name=failure,proto3" json:"failure,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *TxnRequest) Reset() {
	*x = TxnRequest{}
	mi := &file_rpc_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *TxnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*TxnRequest) ProtoMessage() {}
func (x *TxnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*TxnRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{10}
}
func (x *TxnRequest) GetCompare() []*Compare {
	if x != nil {
		return x.Compare
	}
	return nil
}
func (x *TxnRequest) GetSuccess() []*RequestOp {
	if x != nil {
		return x.Success
	}
	return nil
}
func (x *TxnRequest) GetFailure() []*RequestOp {
	if x != nil {
		return x.Failure
	}
	return nil
}
type TxnResponse struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	Header *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Succeeded bool `protobuf:"varint,2,opt,name=succeeded,proto3" json:"succeeded,omitempty"`
	Responses     []*ResponseOp `protobuf:"bytes,3,rep,name=responses,proto3" json:"responses,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *TxnResponse) Reset() {
	*x = TxnResponse{}
	mi := &file_rpc_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *TxnResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*TxnResponse) ProtoMessage() {}
func (x *TxnResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*TxnResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{11}
}
func (x *TxnResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
func (x *TxnResponse) GetSucceeded() bool {
	if x != nil {
		return x.Succeeded
	}
	return false
}
func (x *TxnResponse) GetResponses() []*ResponseOp {
	if x != nil {
		return x.Responses
	}
	return nil
}
type CompactionRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Revision int64 `protobuf:"varint,1,opt,name=revision,proto3" json:"revision,omitempty"`
	Physical      bool `protobuf:"varint,2,opt,name=physical,proto3" json:"physical,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *CompactionRequest) Reset() {
	*x = CompactionRequest{}
	mi := &file_rpc_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *CompactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*CompactionRequest) ProtoMessage() {}
func (x *CompactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*CompactionRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{12}
}
func (x *CompactionRequest) GetRevision() int64 {
	if x != nil {
		return x.Revision
	}
	return 0
}
func (x *CompactionRequest) GetPhysical() bool {
	if x != nil {
		return x.Physical
	}
	return false
}
type CompactionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Header        *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *CompactionResponse) Reset() {
	*x = CompactionResponse{}
	mi := &file_rpc_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *CompactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*CompactionResponse) ProtoMessage() {}
func (x *CompactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*CompactionResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{13}
}
func (x *CompactionResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
type HashRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *HashRequest) Reset() {
	*x = HashRequest{}
	mi := &file_rpc_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *HashRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*HashRequest) ProtoMessage() {}
func (x *HashRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*HashRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{14}
}
type HashKVRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Revision      int64 `protobuf:"varint,1,opt,name=revision,proto3" json:"revision,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *HashKVRequest) Reset() {
	*x = HashKVRequest{}
	mi := &file_rpc_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *HashKVRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*HashKVRequest) ProtoMessage() {}
func (x *HashKVRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*HashKVRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{15}
}
func (x *HashKVRequest) GetRevision() int64 {
	if x != nil {
		return x.Revision
	}
	return 0
}
type HashKVResponse struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	Header *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Hash uint32 `protobuf:"varint,2,opt,name=hash,proto3" json:"hash,omitempty"`
	CompactRevision int64 `protobuf:"varint,3,opt,name=compact_revision,json=compactRevision,proto3" json:"compact_revision,omitempty"`
	HashRevision  int64 `protobuf:"varint,4,opt,name=hash_revision,json=hashRevision,proto3" json:"hash_revision,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *HashKVResponse) Reset() {
	*x = HashKVResponse{}
	mi := &file_rpc_proto_msgTypes[16]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *HashKVResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*HashKVResponse) ProtoMessage() {}
func (x *HashKVResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[16]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*HashKVResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{16}
}
func (x *HashKVResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
func (x *HashKVResponse) GetHash() uint32 {
	if x != nil {
		return x.Hash
	}
	return 0
}
func (x *HashKVResponse) GetCompactRevision() int64 {
	if x != nil {
		return x.CompactRevision
	}
	return 0
}
func (x *HashKVResponse) GetHashRevision() int64 {
	if x != nil {
		return x.HashRevision
	}
	return 0
}
type HashResponse struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	Header *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Hash          uint32 `protobuf:"varint,2,opt,name=hash,proto3" json:"hash,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *HashResponse) Reset() {
	*x = HashResponse{}
	mi := &file_rpc_proto_msgTypes[17]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *HashResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*HashResponse) ProtoMessage() {}
func (x *HashResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[17]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*HashResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{17}
}
func (x *HashResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
func (x *HashResponse) GetHash() uint32 {
	if x != nil {
		return x.Hash
	}
	return 0
}
type SnapshotRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *SnapshotRequest) Reset() {
	*x = SnapshotRequest{}
	mi := &file_rpc_proto_msgTypes[18]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *SnapshotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*SnapshotRequest) ProtoMessage() {}
func (x *SnapshotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[18]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*SnapshotRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{18}
}
type SnapshotResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	RemainingBytes uint64 `protobuf:"varint,2,opt,name=remaining_bytes,json=remainingBytes,proto3" json:"remaining_bytes,omitempty"`
	Blob []byte `protobuf:"bytes,3,opt,name=blob,proto3" json:"blob,omitempty"`
	Version       string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *SnapshotResponse) Reset() {
	*x = SnapshotResponse{}
	mi := &file_rpc_proto_msgTypes[19]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *SnapshotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*SnapshotResponse) ProtoMessage() {}
func (x *SnapshotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[19]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*SnapshotResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{19}
}
func (x *SnapshotResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
func (x *SnapshotResponse) GetRemainingBytes() uint64 {
	if x != nil {
		return x.RemainingBytes
	}
	return 0
}
func (x *SnapshotResponse) GetBlob() []byte {
	if x != nil {
		return x.Blob
	}
	return nil
}
func (x *SnapshotResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}
type WatchRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	RequestUnion  isWatchRequest_RequestUnion `protobuf_oneof:"request_union"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *WatchRequest) Reset() {
	*x = WatchRequest{}
	mi := &file_rpc_proto_msgTypes[20]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *WatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*WatchRequest) ProtoMessage() {}
func (x *WatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[20]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*WatchRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{20}
}
func (x *WatchRequest) GetRequestUnion() isWatchRequest_RequestUnion {
	if x != nil {
		return x.RequestUnion
	}
	return nil
}
func (x *WatchRequest) GetCreateRequest() *WatchCreateRequest {
	if x != nil {
		if x, ok := x.RequestUnion.(*WatchRequest_CreateRequest); ok {
			return x.CreateRequest
		}
	}
	return nil
}
func (x *WatchRequest) GetCancelRequest() *WatchCancelRequest {
	if x != nil {
		if x, ok := x.RequestUnion.(*WatchRequest_CancelRequest); ok {
			return x.CancelRequest
		}
	}
	return nil
}
func (x *WatchRequest) GetProgressRequest() *WatchProgressRequest {
	if x != nil {
		if x, ok := x.RequestUnion.(*WatchRequest_ProgressRequest); ok {
			return x.ProgressRequest
		}
	}
	return nil
}
type isWatchRequest_RequestUnion interface {
	isWatchRequest_RequestUnion()
}
type WatchRequest_CreateRequest struct {
	CreateRequest *WatchCreateRequest `protobuf:"bytes,1,opt,name=create_request,json=createRequest,proto3,oneof"`
}
type WatchRequest_CancelRequest struct {
	CancelRequest *WatchCancelRequest `protobuf:"bytes,2,opt,name=cancel_request,json=cancelRequest,proto3,oneof"`
}
type WatchRequest_ProgressRequest struct {
	ProgressRequest *WatchProgressRequest `protobuf:"bytes,3,opt,name=progress_request,json=progressRequest,proto3,oneof"`
}
func (*WatchRequest_CreateRequest) isWatchRequest_RequestUnion() {}
func (*WatchRequest_CancelRequest) isWatchRequest_RequestUnion() {}
func (*WatchRequest_ProgressRequest) isWatchRequest_RequestUnion() {}
type WatchCreateRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	RangeEnd []byte `protobuf:"bytes,2,opt,name=range_end,json=rangeEnd,proto3" json:"range_end,omitempty"`
	StartRevision int64 `protobuf:"varint,3,opt,name=start_revision,json=startRevision,proto3" json:"start_revision,omitempty"`
	ProgressNotify bool `protobuf:"varint,4,opt,name=progress_notify,json=progressNotify,proto3" json:"progress_notify,omitempty"`
	Filters []WatchCreateRequest_FilterType `protobuf:"varint,5,rep,packed,name=filters,proto3,enum=etcdserverpb.WatchCreateRequest_FilterType" json:"filters,omitempty"`
	PrevKv bool `protobuf:"varint,6,opt,name=prev_kv,json=prevKv,proto3" json:"prev_kv,omitempty"`
	WatchId int64 `protobuf:"varint,7,opt,name=watch_id,json=watchId,proto3" json:"watch_id,omitempty"`
	Fragment      bool `protobuf:"varint,8,opt,name=fragment,proto3" json:"fragment,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *WatchCreateRequest) Reset() {
	*x = WatchCreateRequest{}
	mi := &file_rpc_proto_msgTypes[21]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *WatchCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*WatchCreateRequest) ProtoMessage() {}
func (x *WatchCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[21]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*WatchCreateRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{21}
}
func (x *WatchCreateRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}
func (x *WatchCreateRequest) GetRangeEnd() []byte {
	if x != nil {
		return x.RangeEnd
	}
	return nil
}
func (x *WatchCreateRequest) GetStartRevision() int64 {
	if x != nil {
		return x.StartRevision
	}
	return 0
}
func (x *WatchCreateRequest) GetProgressNotify() bool {
	if x != nil {
		return x.ProgressNotify
	}
	return false
}
func (x *WatchCreateRequest) GetFilters() []WatchCreateRequest_FilterType {
	if x != nil {
		return x.Filters
	}
	return nil
}
func (x *WatchCreateRequest) GetPrevKv() bool {
	if x != nil {
		return x.PrevKv
	}
	return false
}
func (x *WatchCreateRequest) GetWatchId() int64 {
	if x != nil {
		return x.WatchId
	}
	return 0
}
func (x *WatchCreateRequest) GetFragment() bool {
	if x != nil {
		return x.Fragment
	}
	return false
}
type WatchCancelRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	WatchId       int64 `protobuf:"varint,1,opt,name=watch_id,json=watchId,proto3" json:"watch_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *WatchCancelRequest) Reset() {
	*x = WatchCancelRequest{}
	mi := &file_rpc_proto_msgTypes[22]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *WatchCancelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*WatchCancelRequest) ProtoMessage() {}
func (x *WatchCancelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[22]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*WatchCancelRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{22}
}
func (x *WatchCancelRequest) GetWatchId() int64 {
	if x != nil {
		return x.WatchId
	}
	return 0
}
type WatchProgressRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *WatchProgressRequest) Reset() {
	*x = WatchProgressRequest{}
	mi := &file_rpc_proto_msgTypes[23]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *WatchProgressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*WatchProgressRequest) ProtoMessage() {}
func (x *WatchProgressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[23]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*WatchProgressRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{23}
}
type WatchResponse struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	Header *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	WatchId int64 `protobuf:"varint,2,opt,name=watch_id,json=watchId,proto3" json:"watch_id,omitempty"`
	Created bool `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Canceled bool `protobuf:"varint,4,opt,name=canceled,proto3" json:"canceled,omitempty"`
	CompactRevision int64 `protobuf:"varint,5,opt,name=compact_revision,json=compactRevision,proto3" json:"compact_revision,omitempty"`
	CancelReason string `protobuf:"bytes,6,opt,name=cancel_reason,json=cancelReason,proto3" json:"cancel_reason,omitempty"`
	Fragment      bool            `protobuf:"varint,7,opt,name=fragment,proto3" json:"fragment,omitempty"`
	Events        []*mvccpb.Event `protobuf:"bytes,11,rep,name=events,proto3" json:"events,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *WatchResponse) Reset() {
	*x = WatchResponse{}
	mi := &file_rpc_proto_msgTypes[24]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *WatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*WatchResponse) ProtoMessage() {}
func (x *WatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[24]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*WatchResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{24}
}
func (x *WatchResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
func (x *WatchResponse) GetWatchId() int64 {
	if x != nil {
		return x.WatchId
	}
	return 0
}
func (x *WatchResponse) GetCreated() bool {
	if x != nil {
		return x.Created
	}
	return false
}
func (x *WatchResponse) GetCanceled() bool {
	if x != nil {
		return x.Canceled
	}
	return false
}
func (x *WatchResponse) GetCompactRevision() int64 {
	if x != nil {
		return x.CompactRevision
	}
	return 0
}
func (x *WatchResponse) GetCancelReason() string {
	if x != nil {
		return x.CancelReason
	}
	return ""
}
func (x *WatchResponse) GetFragment() bool {
	if x != nil {
		return x.Fragment
	}
	return false
}
func (x *WatchResponse) GetEvents() []*mvccpb.Event {
	if x != nil {
		return x.Events
	}
	return nil
}
type LeaseGrantRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	TTL int64 `protobuf:"varint,1,opt,name=TTL,proto3" json:"TTL,omitempty"`
	ID            int64 `protobuf:"varint,2,opt,name=ID,proto3" json:"ID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *LeaseGrantRequest) Reset() {
	*x = LeaseGrantRequest{}
	mi := &file_rpc_proto_msgTypes[25]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *LeaseGrantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*LeaseGrantRequest) ProtoMessage() {}
func (x *LeaseGrantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[25]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*LeaseGrantRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{25}
}
func (x *LeaseGrantRequest) GetTTL() int64 {
	if x != nil {
		return x.TTL
	}
	return 0
}
func (x *LeaseGrantRequest) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}
type LeaseGrantResponse struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	Header *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	ID int64 `protobuf:"varint,2,opt,name=ID,proto3" json:"ID,omitempty"`
	TTL           int64  `protobuf:"varint,3,opt,name=TTL,proto3" json:"TTL,omitempty"`
	Error         string `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *LeaseGrantResponse) Reset() {
	*x = LeaseGrantResponse{}
	mi := &file_rpc_proto_msgTypes[26]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *LeaseGrantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*LeaseGrantResponse) ProtoMessage() {}
func (x *LeaseGrantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[26]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*LeaseGrantResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{26}
}
func (x *LeaseGrantResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
func (x *LeaseGrantResponse) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}
func (x *LeaseGrantResponse) GetTTL() int64 {
	if x != nil {
		return x.TTL
	}
	return 0
}
func (x *LeaseGrantResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}
type LeaseRevokeRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	ID            int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *LeaseRevokeRequest) Reset() {
	*x = LeaseRevokeRequest{}
	mi := &file_rpc_proto_msgTypes[27]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *LeaseRevokeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*LeaseRevokeRequest) ProtoMessage() {}
func (x *LeaseRevokeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[27]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*LeaseRevokeRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{27}
}
func (x *LeaseRevokeRequest) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}
type LeaseRevokeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Header        *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *LeaseRevokeResponse) Reset() {
	*x = LeaseRevokeResponse{}
	mi := &file_rpc_proto_msgTypes[28]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *LeaseRevokeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*LeaseRevokeResponse) ProtoMessage() {}
func (x *LeaseRevokeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[28]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*LeaseRevokeResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{28}
}
func (x *LeaseRevokeResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
type LeaseCheckpoint struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Remaining_TTL int64 `protobuf:"varint,2,opt,name=remaining_TTL,json=remainingTTL,proto3" json:"remaining_TTL,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *LeaseCheckpoint) Reset() {
	*x = LeaseCheckpoint{}
	mi := &file_rpc_proto_msgTypes[29]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *LeaseCheckpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*LeaseCheckpoint) ProtoMessage() {}
func (x *LeaseCheckpoint) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[29]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*LeaseCheckpoint) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{29}
}
func (x *LeaseCheckpoint) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}
func (x *LeaseCheckpoint) GetRemaining_TTL() int64 {
	if x != nil {
		return x.Remaining_TTL
	}
	return 0
}
type LeaseCheckpointRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Checkpoints   []*LeaseCheckpoint     `protobuf:"bytes,1,rep,name=checkpoints,proto3" json:"checkpoints,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *LeaseCheckpointRequest) Reset() {
	*x = LeaseCheckpointRequest{}
	mi := &file_rpc_proto_msgTypes[30]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *LeaseCheckpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*LeaseCheckpointRequest) ProtoMessage() {}
func (x *LeaseCheckpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[30]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*LeaseCheckpointRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{30}
}
func (x *LeaseCheckpointRequest) GetCheckpoints() []*LeaseCheckpoint {
	if x != nil {
		return x.Checkpoints
	}
	return nil
}
type LeaseCheckpointResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Header        *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *LeaseCheckpointResponse) Reset() {
	*x = LeaseCheckpointResponse{}
	mi := &file_rpc_proto_msgTypes[31]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *LeaseCheckpointResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*LeaseCheckpointResponse) ProtoMessage() {}
func (x *LeaseCheckpointResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[31]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*LeaseCheckpointResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{31}
}
func (x *LeaseCheckpointResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
type LeaseKeepAliveRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	ID            int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *LeaseKeepAliveRequest) Reset() {
	*x = LeaseKeepAliveRequest{}
	mi := &file_rpc_proto_msgTypes[32]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *LeaseKeepAliveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*LeaseKeepAliveRequest) ProtoMessage() {}
func (x *LeaseKeepAliveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[32]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*LeaseKeepAliveRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{32}
}
func (x *LeaseKeepAliveRequest) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}
type LeaseKeepAliveResponse struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	Header *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	ID int64 `protobuf:"varint,2,opt,name=ID,proto3" json:"ID,omitempty"`
	TTL           int64 `protobuf:"varint,3,opt,name=TTL,proto3" json:"TTL,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *LeaseKeepAliveResponse) Reset() {
	*x = LeaseKeepAliveResponse{}
	mi := &file_rpc_proto_msgTypes[33]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *LeaseKeepAliveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*LeaseKeepAliveResponse) ProtoMessage() {}
func (x *LeaseKeepAliveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[33]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*LeaseKeepAliveResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{33}
}
func (x *LeaseKeepAliveResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
func (x *LeaseKeepAliveResponse) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}
func (x *LeaseKeepAliveResponse) GetTTL() int64 {
	if x != nil {
		return x.TTL
	}
	return 0
}
type LeaseTimeToLiveRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Keys          bool `protobuf:"varint,2,opt,name=keys,proto3" json:"keys,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *LeaseTimeToLiveRequest) Reset() {
	*x = LeaseTimeToLiveRequest{}
	mi := &file_rpc_proto_msgTypes[34]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *LeaseTimeToLiveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*LeaseTimeToLiveRequest) ProtoMessage() {}
func (x *LeaseTimeToLiveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[34]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*LeaseTimeToLiveRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{34}
}
func (x *LeaseTimeToLiveRequest) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}
func (x *LeaseTimeToLiveRequest) GetKeys() bool {
	if x != nil {
		return x.Keys
	}
	return false
}
type LeaseTimeToLiveResponse struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	Header *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	ID int64 `protobuf:"varint,2,opt,name=ID,proto3" json:"ID,omitempty"`
	TTL int64 `protobuf:"varint,3,opt,name=TTL,proto3" json:"TTL,omitempty"`
	GrantedTTL int64 `protobuf:"varint,4,opt,name=grantedTTL,proto3" json:"grantedTTL,omitempty"`
	Keys          [][]byte `protobuf:"bytes,5,rep,name=keys,proto3" json:"keys,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *LeaseTimeToLiveResponse) Reset() {
	*x = LeaseTimeToLiveResponse{}
	mi := &file_rpc_proto_msgTypes[35]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *LeaseTimeToLiveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*LeaseTimeToLiveResponse) ProtoMessage() {}
func (x *LeaseTimeToLiveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[35]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*LeaseTimeToLiveResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{35}
}
func (x *LeaseTimeToLiveResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
func (x *LeaseTimeToLiveResponse) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}
func (x *LeaseTimeToLiveResponse) GetTTL() int64 {
	if x != nil {
		return x.TTL
	}
	return 0
}
func (x *LeaseTimeToLiveResponse) GetGrantedTTL() int64 {
	if x != nil {
		return x.GrantedTTL
	}
	return 0
}
func (x *LeaseTimeToLiveResponse) GetKeys() [][]byte {
	if x != nil {
		return x.Keys
	}
	return nil
}
type LeaseLeasesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *LeaseLeasesRequest) Reset() {
	*x = LeaseLeasesRequest{}
	mi := &file_rpc_proto_msgTypes[36]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *LeaseLeasesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*LeaseLeasesRequest) ProtoMessage() {}
func (x *LeaseLeasesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[36]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*LeaseLeasesRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{36}
}
type LeaseStatus struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ID            int64                  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"` 
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *LeaseStatus) Reset() {
	*x = LeaseStatus{}
	mi := &file_rpc_proto_msgTypes[37]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *LeaseStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*LeaseStatus) ProtoMessage() {}
func (x *LeaseStatus) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[37]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*LeaseStatus) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{37}
}
func (x *LeaseStatus) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}
type LeaseLeasesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Header        *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Leases        []*LeaseStatus         `protobuf:"bytes,2,rep,name=leases,proto3" json:"leases,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *LeaseLeasesResponse) Reset() {
	*x = LeaseLeasesResponse{}
	mi := &file_rpc_proto_msgTypes[38]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *LeaseLeasesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*LeaseLeasesResponse) ProtoMessage() {}
func (x *LeaseLeasesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[38]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*LeaseLeasesResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{38}
}
func (x *LeaseLeasesResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
func (x *LeaseLeasesResponse) GetLeases() []*LeaseStatus {
	if x != nil {
		return x.Leases
	}
	return nil
}
type Member struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PeerURLs []string `protobuf:"bytes,3,rep,name=peerURLs,proto3" json:"peerURLs,omitempty"`
	ClientURLs []string `protobuf:"bytes,4,rep,name=clientURLs,proto3" json:"clientURLs,omitempty"`
	IsLearner     bool `protobuf:"varint,5,opt,name=isLearner,proto3" json:"isLearner,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *Member) Reset() {
	*x = Member{}
	mi := &file_rpc_proto_msgTypes[39]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*Member) ProtoMessage() {}
func (x *Member) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[39]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*Member) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{39}
}
func (x *Member) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}
func (x *Member) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}
func (x *Member) GetPeerURLs() []string {
	if x != nil {
		return x.PeerURLs
	}
	return nil
}
func (x *Member) GetClientURLs() []string {
	if x != nil {
		return x.ClientURLs
	}
	return nil
}
func (x *Member) GetIsLearner() bool {
	if x != nil {
		return x.IsLearner
	}
	return false
}
type MemberAddRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	PeerURLs []string `protobuf:"bytes,1,rep,name=peerURLs,proto3" json:"peerURLs,omitempty"`
	IsLearner     bool `protobuf:"varint,2,opt,name=isLearner,proto3" json:"isLearner,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *MemberAddRequest) Reset() {
	*x = MemberAddRequest{}
	mi := &file_rpc_proto_msgTypes[40]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *MemberAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*MemberAddRequest) ProtoMessage() {}
func (x *MemberAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[40]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*MemberAddRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{40}
}
func (x *MemberAddRequest) GetPeerURLs() []string {
	if x != nil {
		return x.PeerURLs
	}
	return nil
}
func (x *MemberAddRequest) GetIsLearner() bool {
	if x != nil {
		return x.IsLearner
	}
	return false
}
type MemberAddResponse struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	Header *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Member *Member `protobuf:"bytes,2,opt,name=member,proto3" json:"member,omitempty"`
	Members       []*Member `protobuf:"bytes,3,rep,name=members,proto3" json:"members,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *MemberAddResponse) Reset() {
	*x = MemberAddResponse{}
	mi := &file_rpc_proto_msgTypes[41]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *MemberAddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*MemberAddResponse) ProtoMessage() {}
func (x *MemberAddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[41]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*MemberAddResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{41}
}
func (x *MemberAddResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
func (x *MemberAddResponse) GetMember() *Member {
	if x != nil {
		return x.Member
	}
	return nil
}
func (x *MemberAddResponse) GetMembers() []*Member {
	if x != nil {
		return x.Members
	}
	return nil
}
type MemberRemoveRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	ID            uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *MemberRemoveRequest) Reset() {
	*x = MemberRemoveRequest{}
	mi := &file_rpc_proto_msgTypes[42]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *MemberRemoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*MemberRemoveRequest) ProtoMessage() {}
func (x *MemberRemoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[42]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*MemberRemoveRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{42}
}
func (x *MemberRemoveRequest) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}
type MemberRemoveResponse struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	Header *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Members       []*Member `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *MemberRemoveResponse) Reset() {
	*x = MemberRemoveResponse{}
	mi := &file_rpc_proto_msgTypes[43]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *MemberRemoveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*MemberRemoveResponse) ProtoMessage() {}
func (x *MemberRemoveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[43]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*MemberRemoveResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{43}
}
func (x *MemberRemoveResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
func (x *MemberRemoveResponse) GetMembers() []*Member {
	if x != nil {
		return x.Members
	}
	return nil
}
type MemberUpdateRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	PeerURLs      []string `protobuf:"bytes,2,rep,name=peerURLs,proto3" json:"peerURLs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *MemberUpdateRequest) Reset() {
	*x = MemberUpdateRequest{}
	mi := &file_rpc_proto_msgTypes[44]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *MemberUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*MemberUpdateRequest) ProtoMessage() {}
func (x *MemberUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[44]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*MemberUpdateRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{44}
}
func (x *MemberUpdateRequest) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}
func (x *MemberUpdateRequest) GetPeerURLs() []string {
	if x != nil {
		return x.PeerURLs
	}
	return nil
}
type MemberUpdateResponse struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	Header *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Members       []*Member `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *MemberUpdateResponse) Reset() {
	*x = MemberUpdateResponse{}
	mi := &file_rpc_proto_msgTypes[45]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *MemberUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*MemberUpdateResponse) ProtoMessage() {}
func (x *MemberUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[45]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*MemberUpdateResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{45}
}
func (x *MemberUpdateResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
func (x *MemberUpdateResponse) GetMembers() []*Member {
	if x != nil {
		return x.Members
	}
	return nil
}
type MemberListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Linearizable  bool                   `protobuf:"varint,1,opt,name=linearizable,proto3" json:"linearizable,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *MemberListRequest) Reset() {
	*x = MemberListRequest{}
	mi := &file_rpc_proto_msgTypes[46]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *MemberListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*MemberListRequest) ProtoMessage() {}
func (x *MemberListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[46]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*MemberListRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{46}
}
func (x *MemberListRequest) GetLinearizable() bool {
	if x != nil {
		return x.Linearizable
	}
	return false
}
type MemberListResponse struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	Header *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Members       []*Member `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *MemberListResponse) Reset() {
	*x = MemberListResponse{}
	mi := &file_rpc_proto_msgTypes[47]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *MemberListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*MemberListResponse) ProtoMessage() {}
func (x *MemberListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[47]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*MemberListResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{47}
}
func (x *MemberListResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
func (x *MemberListResponse) GetMembers() []*Member {
	if x != nil {
		return x.Members
	}
	return nil
}
type MemberPromoteRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	ID            uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *MemberPromoteRequest) Reset() {
	*x = MemberPromoteRequest{}
	mi := &file_rpc_proto_msgTypes[48]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *MemberPromoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*MemberPromoteRequest) ProtoMessage() {}
func (x *MemberPromoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[48]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*MemberPromoteRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{48}
}
func (x *MemberPromoteRequest) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}
type MemberPromoteResponse struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	Header *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Members       []*Member `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *MemberPromoteResponse) Reset() {
	*x = MemberPromoteResponse{}
	mi := &file_rpc_proto_msgTypes[49]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *MemberPromoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*MemberPromoteResponse) ProtoMessage() {}
func (x *MemberPromoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[49]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*MemberPromoteResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{49}
}
func (x *MemberPromoteResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
func (x *MemberPromoteResponse) GetMembers() []*Member {
	if x != nil {
		return x.Members
	}
	return nil
}
type DefragmentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *DefragmentRequest) Reset() {
	*x = DefragmentRequest{}
	mi := &file_rpc_proto_msgTypes[50]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *DefragmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*DefragmentRequest) ProtoMessage() {}
func (x *DefragmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[50]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*DefragmentRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{50}
}
type DefragmentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Header        *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *DefragmentResponse) Reset() {
	*x = DefragmentResponse{}
	mi := &file_rpc_proto_msgTypes[51]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *DefragmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*DefragmentResponse) ProtoMessage() {}
func (x *DefragmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[51]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*DefragmentResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{51}
}
func (x *DefragmentResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
type MoveLeaderRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	TargetID      uint64 `protobuf:"varint,1,opt,name=targetID,proto3" json:"targetID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *MoveLeaderRequest) Reset() {
	*x = MoveLeaderRequest{}
	mi := &file_rpc_proto_msgTypes[52]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *MoveLeaderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*MoveLeaderRequest) ProtoMessage() {}
func (x *MoveLeaderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[52]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*MoveLeaderRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{52}
}
func (x *MoveLeaderRequest) GetTargetID() uint64 {
	if x != nil {
		return x.TargetID
	}
	return 0
}
type MoveLeaderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Header        *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *MoveLeaderResponse) Reset() {
	*x = MoveLeaderResponse{}
	mi := &file_rpc_proto_msgTypes[53]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *MoveLeaderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*MoveLeaderResponse) ProtoMessage() {}
func (x *MoveLeaderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[53]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*MoveLeaderResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{53}
}
func (x *MoveLeaderResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
type AlarmRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Action AlarmRequest_AlarmAction `protobuf:"varint,1,opt,name=action,proto3,enum=etcdserverpb.AlarmRequest_AlarmAction" json:"action,omitempty"`
	MemberID uint64 `protobuf:"varint,2,opt,name=memberID,proto3" json:"memberID,omitempty"`
	Alarm         AlarmType `protobuf:"varint,3,opt,name=alarm,proto3,enum=etcdserverpb.AlarmType" json:"alarm,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AlarmRequest) Reset() {
	*x = AlarmRequest{}
	mi := &file_rpc_proto_msgTypes[54]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AlarmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AlarmRequest) ProtoMessage() {}
func (x *AlarmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[54]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AlarmRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{54}
}
func (x *AlarmRequest) GetAction() AlarmRequest_AlarmAction {
	if x != nil {
		return x.Action
	}
	return AlarmRequest_GET
}
func (x *AlarmRequest) GetMemberID() uint64 {
	if x != nil {
		return x.MemberID
	}
	return 0
}
func (x *AlarmRequest) GetAlarm() AlarmType {
	if x != nil {
		return x.Alarm
	}
	return AlarmType_NONE
}
type AlarmMember struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	MemberID uint64 `protobuf:"varint,1,opt,name=memberID,proto3" json:"memberID,omitempty"`
	Alarm         AlarmType `protobuf:"varint,2,opt,name=alarm,proto3,enum=etcdserverpb.AlarmType" json:"alarm,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AlarmMember) Reset() {
	*x = AlarmMember{}
	mi := &file_rpc_proto_msgTypes[55]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AlarmMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AlarmMember) ProtoMessage() {}
func (x *AlarmMember) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[55]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AlarmMember) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{55}
}
func (x *AlarmMember) GetMemberID() uint64 {
	if x != nil {
		return x.MemberID
	}
	return 0
}
func (x *AlarmMember) GetAlarm() AlarmType {
	if x != nil {
		return x.Alarm
	}
	return AlarmType_NONE
}
type AlarmResponse struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	Header *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Alarms        []*AlarmMember `protobuf:"bytes,2,rep,name=alarms,proto3" json:"alarms,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AlarmResponse) Reset() {
	*x = AlarmResponse{}
	mi := &file_rpc_proto_msgTypes[56]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AlarmResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AlarmResponse) ProtoMessage() {}
func (x *AlarmResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[56]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AlarmResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{56}
}
func (x *AlarmResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
func (x *AlarmResponse) GetAlarms() []*AlarmMember {
	if x != nil {
		return x.Alarms
	}
	return nil
}
type DowngradeRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Action DowngradeRequest_DowngradeAction `protobuf:"varint,1,opt,name=action,proto3,enum=etcdserverpb.DowngradeRequest_DowngradeAction" json:"action,omitempty"`
	Version       string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *DowngradeRequest) Reset() {
	*x = DowngradeRequest{}
	mi := &file_rpc_proto_msgTypes[57]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *DowngradeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*DowngradeRequest) ProtoMessage() {}
func (x *DowngradeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[57]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*DowngradeRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{57}
}
func (x *DowngradeRequest) GetAction() DowngradeRequest_DowngradeAction {
	if x != nil {
		return x.Action
	}
	return DowngradeRequest_VALIDATE
}
func (x *DowngradeRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}
type DowngradeResponse struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	Header *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Version       string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *DowngradeResponse) Reset() {
	*x = DowngradeResponse{}
	mi := &file_rpc_proto_msgTypes[58]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *DowngradeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*DowngradeResponse) ProtoMessage() {}
func (x *DowngradeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[58]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*DowngradeResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{58}
}
func (x *DowngradeResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
func (x *DowngradeResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}
type DowngradeVersionTestRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ver           string                 `protobuf:"bytes,1,opt,name=ver,proto3" json:"ver,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *DowngradeVersionTestRequest) Reset() {
	*x = DowngradeVersionTestRequest{}
	mi := &file_rpc_proto_msgTypes[59]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *DowngradeVersionTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*DowngradeVersionTestRequest) ProtoMessage() {}
func (x *DowngradeVersionTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[59]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*DowngradeVersionTestRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{59}
}
func (x *DowngradeVersionTestRequest) GetVer() string {
	if x != nil {
		return x.Ver
	}
	return ""
}
type StatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *StatusRequest) Reset() {
	*x = StatusRequest{}
	mi := &file_rpc_proto_msgTypes[60]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *StatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*StatusRequest) ProtoMessage() {}
func (x *StatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[60]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{60}
}
type StatusResponse struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	Header *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	DbSize int64 `protobuf:"varint,3,opt,name=dbSize,proto3" json:"dbSize,omitempty"`
	Leader uint64 `protobuf:"varint,4,opt,name=leader,proto3" json:"leader,omitempty"`
	RaftIndex uint64 `protobuf:"varint,5,opt,name=raftIndex,proto3" json:"raftIndex,omitempty"`
	RaftTerm uint64 `protobuf:"varint,6,opt,name=raftTerm,proto3" json:"raftTerm,omitempty"`
	RaftAppliedIndex uint64 `protobuf:"varint,7,opt,name=raftAppliedIndex,proto3" json:"raftAppliedIndex,omitempty"`
	Errors []string `protobuf:"bytes,8,rep,name=errors,proto3" json:"errors,omitempty"`
	DbSizeInUse int64 `protobuf:"varint,9,opt,name=dbSizeInUse,proto3" json:"dbSizeInUse,omitempty"`
	IsLearner bool `protobuf:"varint,10,opt,name=isLearner,proto3" json:"isLearner,omitempty"`
	StorageVersion string `protobuf:"bytes,11,opt,name=storageVersion,proto3" json:"storageVersion,omitempty"`
	DbSizeQuota int64 `protobuf:"varint,12,opt,name=dbSizeQuota,proto3" json:"dbSizeQuota,omitempty"`
	DowngradeInfo *DowngradeInfo `protobuf:"bytes,13,opt,name=downgradeInfo,proto3" json:"downgradeInfo,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	mi := &file_rpc_proto_msgTypes[61]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*StatusResponse) ProtoMessage() {}
func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[61]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{61}
}
func (x *StatusResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
func (x *StatusResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}
func (x *StatusResponse) GetDbSize() int64 {
	if x != nil {
		return x.DbSize
	}
	return 0
}
func (x *StatusResponse) GetLeader() uint64 {
	if x != nil {
		return x.Leader
	}
	return 0
}
func (x *StatusResponse) GetRaftIndex() uint64 {
	if x != nil {
		return x.RaftIndex
	}
	return 0
}
func (x *StatusResponse) GetRaftTerm() uint64 {
	if x != nil {
		return x.RaftTerm
	}
	return 0
}
func (x *StatusResponse) GetRaftAppliedIndex() uint64 {
	if x != nil {
		return x.RaftAppliedIndex
	}
	return 0
}
func (x *StatusResponse) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}
func (x *StatusResponse) GetDbSizeInUse() int64 {
	if x != nil {
		return x.DbSizeInUse
	}
	return 0
}
func (x *StatusResponse) GetIsLearner() bool {
	if x != nil {
		return x.IsLearner
	}
	return false
}
func (x *StatusResponse) GetStorageVersion() string {
	if x != nil {
		return x.StorageVersion
	}
	return ""
}
func (x *StatusResponse) GetDbSizeQuota() int64 {
	if x != nil {
		return x.DbSizeQuota
	}
	return 0
}
func (x *StatusResponse) GetDowngradeInfo() *DowngradeInfo {
	if x != nil {
		return x.DowngradeInfo
	}
	return nil
}
type DowngradeInfo struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	TargetVersion string `protobuf:"bytes,2,opt,name=targetVersion,proto3" json:"targetVersion,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *DowngradeInfo) Reset() {
	*x = DowngradeInfo{}
	mi := &file_rpc_proto_msgTypes[62]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *DowngradeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*DowngradeInfo) ProtoMessage() {}
func (x *DowngradeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[62]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*DowngradeInfo) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{62}
}
func (x *DowngradeInfo) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}
func (x *DowngradeInfo) GetTargetVersion() string {
	if x != nil {
		return x.TargetVersion
	}
	return ""
}
type AuthEnableRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AuthEnableRequest) Reset() {
	*x = AuthEnableRequest{}
	mi := &file_rpc_proto_msgTypes[63]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AuthEnableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AuthEnableRequest) ProtoMessage() {}
func (x *AuthEnableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[63]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AuthEnableRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{63}
}
type AuthDisableRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AuthDisableRequest) Reset() {
	*x = AuthDisableRequest{}
	mi := &file_rpc_proto_msgTypes[64]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AuthDisableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AuthDisableRequest) ProtoMessage() {}
func (x *AuthDisableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[64]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AuthDisableRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{64}
}
type AuthStatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AuthStatusRequest) Reset() {
	*x = AuthStatusRequest{}
	mi := &file_rpc_proto_msgTypes[65]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AuthStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AuthStatusRequest) ProtoMessage() {}
func (x *AuthStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[65]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AuthStatusRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{65}
}
type AuthenticateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AuthenticateRequest) Reset() {
	*x = AuthenticateRequest{}
	mi := &file_rpc_proto_msgTypes[66]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AuthenticateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AuthenticateRequest) ProtoMessage() {}
func (x *AuthenticateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[66]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AuthenticateRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{66}
}
func (x *AuthenticateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}
func (x *AuthenticateRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}
type AuthUserAddRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Name           string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password       string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Options        *authpb.UserAddOptions `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
	HashedPassword string                 `protobuf:"bytes,4,opt,name=hashedPassword,proto3" json:"hashedPassword,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}
func (x *AuthUserAddRequest) Reset() {
	*x = AuthUserAddRequest{}
	mi := &file_rpc_proto_msgTypes[67]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AuthUserAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AuthUserAddRequest) ProtoMessage() {}
func (x *AuthUserAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[67]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AuthUserAddRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{67}
}
func (x *AuthUserAddRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}
func (x *AuthUserAddRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}
func (x *AuthUserAddRequest) GetOptions() *authpb.UserAddOptions {
	if x != nil {
		return x.Options
	}
	return nil
}
func (x *AuthUserAddRequest) GetHashedPassword() string {
	if x != nil {
		return x.HashedPassword
	}
	return ""
}
type AuthUserGetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AuthUserGetRequest) Reset() {
	*x = AuthUserGetRequest{}
	mi := &file_rpc_proto_msgTypes[68]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AuthUserGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AuthUserGetRequest) ProtoMessage() {}
func (x *AuthUserGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[68]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AuthUserGetRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{68}
}
func (x *AuthUserGetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}
type AuthUserDeleteRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AuthUserDeleteRequest) Reset() {
	*x = AuthUserDeleteRequest{}
	mi := &file_rpc_proto_msgTypes[69]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AuthUserDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AuthUserDeleteRequest) ProtoMessage() {}
func (x *AuthUserDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[69]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AuthUserDeleteRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{69}
}
func (x *AuthUserDeleteRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}
type AuthUserChangePasswordRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	HashedPassword string `protobuf:"bytes,3,opt,name=hashedPassword,proto3" json:"hashedPassword,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}
func (x *AuthUserChangePasswordRequest) Reset() {
	*x = AuthUserChangePasswordRequest{}
	mi := &file_rpc_proto_msgTypes[70]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AuthUserChangePasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AuthUserChangePasswordRequest) ProtoMessage() {}
func (x *AuthUserChangePasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[70]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AuthUserChangePasswordRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{70}
}
func (x *AuthUserChangePasswordRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}
func (x *AuthUserChangePasswordRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}
func (x *AuthUserChangePasswordRequest) GetHashedPassword() string {
	if x != nil {
		return x.HashedPassword
	}
	return ""
}
type AuthUserGrantRoleRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Role          string `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AuthUserGrantRoleRequest) Reset() {
	*x = AuthUserGrantRoleRequest{}
	mi := &file_rpc_proto_msgTypes[71]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AuthUserGrantRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AuthUserGrantRoleRequest) ProtoMessage() {}
func (x *AuthUserGrantRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[71]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AuthUserGrantRoleRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{71}
}
func (x *AuthUserGrantRoleRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}
func (x *AuthUserGrantRoleRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}
type AuthUserRevokeRoleRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Role          string                 `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AuthUserRevokeRoleRequest) Reset() {
	*x = AuthUserRevokeRoleRequest{}
	mi := &file_rpc_proto_msgTypes[72]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AuthUserRevokeRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AuthUserRevokeRoleRequest) ProtoMessage() {}
func (x *AuthUserRevokeRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[72]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AuthUserRevokeRoleRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{72}
}
func (x *AuthUserRevokeRoleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}
func (x *AuthUserRevokeRoleRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}
type AuthRoleAddRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AuthRoleAddRequest) Reset() {
	*x = AuthRoleAddRequest{}
	mi := &file_rpc_proto_msgTypes[73]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AuthRoleAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AuthRoleAddRequest) ProtoMessage() {}
func (x *AuthRoleAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[73]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AuthRoleAddRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{73}
}
func (x *AuthRoleAddRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}
type AuthRoleGetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Role          string                 `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AuthRoleGetRequest) Reset() {
	*x = AuthRoleGetRequest{}
	mi := &file_rpc_proto_msgTypes[74]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AuthRoleGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AuthRoleGetRequest) ProtoMessage() {}
func (x *AuthRoleGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[74]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AuthRoleGetRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{74}
}
func (x *AuthRoleGetRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}
type AuthUserListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AuthUserListRequest) Reset() {
	*x = AuthUserListRequest{}
	mi := &file_rpc_proto_msgTypes[75]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AuthUserListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AuthUserListRequest) ProtoMessage() {}
func (x *AuthUserListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[75]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AuthUserListRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{75}
}
type AuthRoleListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AuthRoleListRequest) Reset() {
	*x = AuthRoleListRequest{}
	mi := &file_rpc_proto_msgTypes[76]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AuthRoleListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AuthRoleListRequest) ProtoMessage() {}
func (x *AuthRoleListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[76]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AuthRoleListRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{76}
}
type AuthRoleDeleteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Role          string                 `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AuthRoleDeleteRequest) Reset() {
	*x = AuthRoleDeleteRequest{}
	mi := &file_rpc_proto_msgTypes[77]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AuthRoleDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AuthRoleDeleteRequest) ProtoMessage() {}
func (x *AuthRoleDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[77]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AuthRoleDeleteRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{77}
}
func (x *AuthRoleDeleteRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}
type AuthRoleGrantPermissionRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Perm          *authpb.Permission `protobuf:"bytes,2,opt,name=perm,proto3" json:"perm,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AuthRoleGrantPermissionRequest) Reset() {
	*x = AuthRoleGrantPermissionRequest{}
	mi := &file_rpc_proto_msgTypes[78]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AuthRoleGrantPermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AuthRoleGrantPermissionRequest) ProtoMessage() {}
func (x *AuthRoleGrantPermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[78]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AuthRoleGrantPermissionRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{78}
}
func (x *AuthRoleGrantPermissionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}
func (x *AuthRoleGrantPermissionRequest) GetPerm() *authpb.Permission {
	if x != nil {
		return x.Perm
	}
	return nil
}
type AuthRoleRevokePermissionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Role          string                 `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	Key           []byte                 `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	RangeEnd      []byte                 `protobuf:"bytes,3,opt,name=range_end,json=rangeEnd,proto3" json:"range_end,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AuthRoleRevokePermissionRequest) Reset() {
	*x = AuthRoleRevokePermissionRequest{}
	mi := &file_rpc_proto_msgTypes[79]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AuthRoleRevokePermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AuthRoleRevokePermissionRequest) ProtoMessage() {}
func (x *AuthRoleRevokePermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[79]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AuthRoleRevokePermissionRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{79}
}
func (x *AuthRoleRevokePermissionRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}
func (x *AuthRoleRevokePermissionRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}
func (x *AuthRoleRevokePermissionRequest) GetRangeEnd() []byte {
	if x != nil {
		return x.RangeEnd
	}
	return nil
}
type AuthEnableResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Header        *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AuthEnableResponse) Reset() {
	*x = AuthEnableResponse{}
	mi := &file_rpc_proto_msgTypes[80]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AuthEnableResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AuthEnableResponse) ProtoMessage() {}
func (x *AuthEnableResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[80]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AuthEnableResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{80}
}
func (x *AuthEnableResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
type AuthDisableResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Header        *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AuthDisableResponse) Reset() {
	*x = AuthDisableResponse{}
	mi := &file_rpc_proto_msgTypes[81]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AuthDisableResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AuthDisableResponse) ProtoMessage() {}
func (x *AuthDisableResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[81]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AuthDisableResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{81}
}
func (x *AuthDisableResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
type AuthStatusResponse struct {
	state   protoimpl.MessageState `protogen:"open.v1"`
	Header  *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Enabled bool                   `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	AuthRevision  uint64 `protobuf:"varint,3,opt,name=authRevision,proto3" json:"authRevision,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AuthStatusResponse) Reset() {
	*x = AuthStatusResponse{}
	mi := &file_rpc_proto_msgTypes[82]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AuthStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AuthStatusResponse) ProtoMessage() {}
func (x *AuthStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[82]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AuthStatusResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{82}
}
func (x *AuthStatusResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
func (x *AuthStatusResponse) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}
func (x *AuthStatusResponse) GetAuthRevision() uint64 {
	if x != nil {
		return x.AuthRevision
	}
	return 0
}
type AuthenticateResponse struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	Header *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Token         string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AuthenticateResponse) Reset() {
	*x = AuthenticateResponse{}
	mi := &file_rpc_proto_msgTypes[83]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AuthenticateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AuthenticateResponse) ProtoMessage() {}
func (x *AuthenticateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[83]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AuthenticateResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{83}
}
func (x *AuthenticateResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
func (x *AuthenticateResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}
type AuthUserAddResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Header        *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AuthUserAddResponse) Reset() {
	*x = AuthUserAddResponse{}
	mi := &file_rpc_proto_msgTypes[84]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AuthUserAddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AuthUserAddResponse) ProtoMessage() {}
func (x *AuthUserAddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[84]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AuthUserAddResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{84}
}
func (x *AuthUserAddResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
type AuthUserGetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Header        *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Roles         []string               `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AuthUserGetResponse) Reset() {
	*x = AuthUserGetResponse{}
	mi := &file_rpc_proto_msgTypes[85]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AuthUserGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AuthUserGetResponse) ProtoMessage() {}
func (x *AuthUserGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[85]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AuthUserGetResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{85}
}
func (x *AuthUserGetResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
func (x *AuthUserGetResponse) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}
type AuthUserDeleteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Header        *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AuthUserDeleteResponse) Reset() {
	*x = AuthUserDeleteResponse{}
	mi := &file_rpc_proto_msgTypes[86]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AuthUserDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AuthUserDeleteResponse) ProtoMessage() {}
func (x *AuthUserDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[86]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AuthUserDeleteResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{86}
}
func (x *AuthUserDeleteResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
type AuthUserChangePasswordResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Header        *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AuthUserChangePasswordResponse) Reset() {
	*x = AuthUserChangePasswordResponse{}
	mi := &file_rpc_proto_msgTypes[87]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AuthUserChangePasswordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AuthUserChangePasswordResponse) ProtoMessage() {}
func (x *AuthUserChangePasswordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[87]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AuthUserChangePasswordResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{87}
}
func (x *AuthUserChangePasswordResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
type AuthUserGrantRoleResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Header        *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AuthUserGrantRoleResponse) Reset() {
	*x = AuthUserGrantRoleResponse{}
	mi := &file_rpc_proto_msgTypes[88]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AuthUserGrantRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AuthUserGrantRoleResponse) ProtoMessage() {}
func (x *AuthUserGrantRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[88]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AuthUserGrantRoleResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{88}
}
func (x *AuthUserGrantRoleResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
type AuthUserRevokeRoleResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Header        *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AuthUserRevokeRoleResponse) Reset() {
	*x = AuthUserRevokeRoleResponse{}
	mi := &file_rpc_proto_msgTypes[89]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AuthUserRevokeRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AuthUserRevokeRoleResponse) ProtoMessage() {}
func (x *AuthUserRevokeRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[89]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AuthUserRevokeRoleResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{89}
}
func (x *AuthUserRevokeRoleResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
type AuthRoleAddResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Header        *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AuthRoleAddResponse) Reset() {
	*x = AuthRoleAddResponse{}
	mi := &file_rpc_proto_msgTypes[90]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AuthRoleAddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AuthRoleAddResponse) ProtoMessage() {}
func (x *AuthRoleAddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[90]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AuthRoleAddResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{90}
}
func (x *AuthRoleAddResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
type AuthRoleGetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Header        *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Perm          []*authpb.Permission   `protobuf:"bytes,2,rep,name=perm,proto3" json:"perm,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AuthRoleGetResponse) Reset() {
	*x = AuthRoleGetResponse{}
	mi := &file_rpc_proto_msgTypes[91]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AuthRoleGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AuthRoleGetResponse) ProtoMessage() {}
func (x *AuthRoleGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[91]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AuthRoleGetResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{91}
}
func (x *AuthRoleGetResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
func (x *AuthRoleGetResponse) GetPerm() []*authpb.Permission {
	if x != nil {
		return x.Perm
	}
	return nil
}
type AuthRoleListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Header        *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Roles         []string               `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AuthRoleListResponse) Reset() {
	*x = AuthRoleListResponse{}
	mi := &file_rpc_proto_msgTypes[92]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AuthRoleListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AuthRoleListResponse) ProtoMessage() {}
func (x *AuthRoleListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[92]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AuthRoleListResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{92}
}
func (x *AuthRoleListResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
func (x *AuthRoleListResponse) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}
type AuthUserListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Header        *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Users         []string               `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AuthUserListResponse) Reset() {
	*x = AuthUserListResponse{}
	mi := &file_rpc_proto_msgTypes[93]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AuthUserListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AuthUserListResponse) ProtoMessage() {}
func (x *AuthUserListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[93]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AuthUserListResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{93}
}
func (x *AuthUserListResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
func (x *AuthUserListResponse) GetUsers() []string {
	if x != nil {
		return x.Users
	}
	return nil
}
type AuthRoleDeleteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Header        *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AuthRoleDeleteResponse) Reset() {
	*x = AuthRoleDeleteResponse{}
	mi := &file_rpc_proto_msgTypes[94]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AuthRoleDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AuthRoleDeleteResponse) ProtoMessage() {}
func (x *AuthRoleDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[94]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AuthRoleDeleteResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{94}
}
func (x *AuthRoleDeleteResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
type AuthRoleGrantPermissionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Header        *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AuthRoleGrantPermissionResponse) Reset() {
	*x = AuthRoleGrantPermissionResponse{}
	mi := &file_rpc_proto_msgTypes[95]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AuthRoleGrantPermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AuthRoleGrantPermissionResponse) ProtoMessage() {}
func (x *AuthRoleGrantPermissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[95]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AuthRoleGrantPermissionResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{95}
}
func (x *AuthRoleGrantPermissionResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
type AuthRoleRevokePermissionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Header        *ResponseHeader        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *AuthRoleRevokePermissionResponse) Reset() {
	*x = AuthRoleRevokePermissionResponse{}
	mi := &file_rpc_proto_msgTypes[96]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *AuthRoleRevokePermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*AuthRoleRevokePermissionResponse) ProtoMessage() {}
func (x *AuthRoleRevokePermissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[96]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*AuthRoleRevokePermissionResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{96}
}
func (x *AuthRoleRevokePermissionResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}
type RangeStreamResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	RangeResponse *RangeResponse `protobuf:"bytes,1,opt,name=range_response,json=rangeResponse,proto3" json:"range_response,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *RangeStreamResponse) Reset() {
	*x = RangeStreamResponse{}
	mi := &file_rpc_proto_msgTypes[97]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *RangeStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*RangeStreamResponse) ProtoMessage() {}
func (x *RangeStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[97]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*RangeStreamResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{97}
}
func (x *RangeStreamResponse) GetRangeResponse() *RangeResponse {
	if x != nil {
		return x.RangeResponse
	}
	return nil
}
var File_rpc_proto protoreflect.FileDescriptor
const file_rpc_proto_rawDesc = "" +
	"\n" +
	"\trpc.proto\x12\fetcdserverpb\x1a\x18etcd/api/mvccpb/kv.proto\x1a\x1aetcd/api/authpb/auth.proto\x1a etcd/api/versionpb/version.proto\x1a\x1cgoogle/api/annotations.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\x8e\x01\n" +
	"\x0eResponseHeader\x12\x1d\n" +
	"\n" +
	"cluster_id\x18\x01 \x01(\x04R\tclusterId\x12\x1b\n" +
	"\tmember_id\x18\x02 \x01(\x04R\bmemberId\x12\x1a\n" +
	"\brevision\x18\x03 \x01(\x03R\brevision\x12\x1b\n" +
	"\traft_term\x18\x04 \x01(\x04R\braftTerm:\a\x82\xb5\x18\x033.0\"\xc3\x05\n" +
	"\fRangeRequest\x12\x10\n" +
	"\x03key\x18\x01 \x01(\fR\x03key\x12\x1b\n" +
	"\trange_end\x18\x02 \x01(\fR\brangeEnd\x12\x14\n" +
	"\x05limit\x18\x03 \x01(\x03R\x05limit\x12\x1a\n" +
	"\brevision\x18\x04 \x01(\x03R\brevision\x12C\n" +
	"\n" +
	"sort_order\x18\x05 \x01(\x0e2$.etcdserverpb.RangeRequest.SortOrderR\tsortOrder\x12F\n" +
	"\vsort_target\x18\x06 \x01(\x0e2%.etcdserverpb.RangeRequest.SortTargetR\n" +
	"sortTarget\x12\"\n" +
	"\fserializable\x18\a \x01(\bR\fserializable\x12\x1b\n" +
	"\tkeys_only\x18\b \x01(\bR\bkeysOnly\x12\x1d\n" +
	"\n" +
	"count_only\x18\t \x01(\bR\tcountOnly\x121\n" +
	"\x10min_mod_revision\x18\n" +
	" \x01(\x03B\a\x8a\xb5\x18\x033.1R\x0eminModRevision\x121\n" +
	"\x10max_mod_revision\x18\v \x01(\x03B\a\x8a\xb5\x18\x033.1R\x0emaxModRevision\x127\n" +
	"\x13min_create_revision\x18\f \x01(\x03B\a\x8a\xb5\x18\x033.1R\x11minCreateRevision\x127\n" +
	"\x13max_create_revision\x18\r \x01(\x03B\a\x8a\xb5\x18\x033.1R\x11maxCreateRevision\"7\n" +
	"\tSortOrder\x12\b\n" +
	"\x04NONE\x10\x00\x12\n" +
	"\n" +
	"\x06ASCEND\x10\x01\x12\v\n" +
	"\aDESCEND\x10\x02\x1a\a\x92\xb5\x18\x033.0\"K\n" +
	"\n" +
	"SortTarget\x12\a\n" +
	"\x03KEY\x10\x00\x12\v\n" +
	"\aVERSION\x10\x01\x12\n" +
	"\n" +
	"\x06CREATE\x10\x02\x12\a\n" +
	"\x03MOD\x10\x03\x12\t\n" +
	"\x05VALUE\x10\x04\x1a\a\x92\xb5\x18\x033.0:\a\x82\xb5\x18\x033.0\"\x9c\x01\n" +
	"\rRangeResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header\x12\"\n" +
	"\x03kvs\x18\x02 \x03(\v2\x10.mvccpb.KeyValueR\x03kvs\x12\x12\n" +
	"\x04more\x18\x03 \x01(\bR\x04more\x12\x14\n" +
	"\x05count\x18\x04 \x01(\x03R\x05count:\a\x82\xb5\x18\x033.0\"\xcd\x01\n" +
	"\n" +
	"PutRequest\x12\x10\n" +
	"\x03key\x18\x01 \x01(\fR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\fR\x05value\x12\x14\n" +
	"\x05lease\x18\x03 \x01(\x03R\x05lease\x12 \n" +
	"\aprev_kv\x18\x04 \x01(\bB\a\x8a\xb5\x18\x033.1R\x06prevKv\x12*\n" +
	"\fignore_value\x18\x05 \x01(\bB\a\x8a\xb5\x18\x033.2R\vignoreValue\x12*\n" +
	"\fignore_lease\x18\x06 \x01(\bB\a\x8a\xb5\x18\x033.2R\vignoreLease:\a\x82\xb5\x18\x033.0\"\x80\x01\n" +
	"\vPutResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header\x122\n" +
	"\aprev_kv\x18\x02 \x01(\v2\x10.mvccpb.KeyValueB\a\x8a\xb5\x18\x033.1R\x06prevKv:\a\x82\xb5\x18\x033.0\"n\n" +
	"\x12DeleteRangeRequest\x12\x10\n" +
	"\x03key\x18\x01 \x01(\fR\x03key\x12\x1b\n" +
	"\trange_end\x18\x02 \x01(\fR\brangeEnd\x12 \n" +
	"\aprev_kv\x18\x03 \x01(\bB\a\x8a\xb5\x18\x033.1R\x06prevKv:\a\x82\xb5\x18\x033.0\"\xa4\x01\n" +
	"\x13DeleteRangeResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header\x12\x18\n" +
	"\adeleted\x18\x02 \x01(\x03R\adeleted\x124\n" +
	"\bprev_kvs\x18\x03 \x03(\v2\x10.mvccpb.KeyValueB\a\x8a\xb5\x18\x033.1R\aprevKvs:\a\x82\xb5\x18\x033.0\"\xbb\x02\n" +
	"\tRequestOp\x12A\n" +
	"\rrequest_range\x18\x01 \x01(\v2\x1a.etcdserverpb.RangeRequestH\x00R\frequestRange\x12;\n" +
	"\vrequest_put\x18\x02 \x01(\v2\x18.etcdserverpb.PutRequestH\x00R\n" +
	"requestPut\x12T\n" +
	"\x14request_delete_range\x18\x03 \x01(\v2 .etcdserverpb.DeleteRangeRequestH\x00R\x12requestDeleteRange\x12D\n" +
	"\vrequest_txn\x18\x04 \x01(\v2\x18.etcdserverpb.TxnRequestB\a\x8a\xb5\x18\x033.3H\x00R\n" +
	"requestTxn:\a\x82\xb5\x18\x033.0B\t\n" +
	"\arequest\"\xc9\x02\n" +
	"\n" +
	"ResponseOp\x12D\n" +
	"\x0eresponse_range\x18\x01 \x01(\v2\x1b.etcdserverpb.RangeResponseH\x00R\rresponseRange\x12>\n" +
	"\fresponse_put\x18\x02 \x01(\v2\x19.etcdserverpb.PutResponseH\x00R\vresponsePut\x12W\n" +
	"\x15response_delete_range\x18\x03 \x01(\v2!.etcdserverpb.DeleteRangeResponseH\x00R\x13responseDeleteRange\x12G\n" +
	"\fresponse_txn\x18\x04 \x01(\v2\x19.etcdserverpb.TxnResponseB\a\x8a\xb5\x18\x033.3H\x00R\vresponseTxn:\a\x82\xb5\x18\x033.0B\n" +
	"\n" +
	"\bresponse\"\xa8\x04\n" +
	"\aCompare\x12;\n" +
	"\x06result\x18\x01 \x01(\x0e2#.etcdserverpb.Compare.CompareResultR\x06result\x12;\n" +
	"\x06target\x18\x02 \x01(\x0e2#.etcdserverpb.Compare.CompareTargetR\x06target\x12\x10\n" +
	"\x03key\x18\x03 \x01(\fR\x03key\x12\x1a\n" +
	"\aversion\x18\x04 \x01(\x03H\x00R\aversion\x12)\n" +
	"\x0fcreate_revision\x18\x05 \x01(\x03H\x00R\x0ecreateRevision\x12#\n" +
	"\fmod_revision\x18\x06 \x01(\x03H\x00R\vmodRevision\x12\x16\n" +
	"\x05value\x18\a \x01(\fH\x00R\x05value\x12\x1f\n" +
	"\x05lease\x18\b \x01(\x03B\a\x8a\xb5\x18\x033.3H\x00R\x05lease\x12$\n" +
	"\trange_end\x18@ \x01(\fB\a\x8a\xb5\x18\x033.3R\brangeEnd\"R\n" +
	"\rCompareResult\x12\t\n" +
	"\x05EQUAL\x10\x00\x12\v\n" +
	"\aGREATER\x10\x01\x12\b\n" +
	"\x04LESS\x10\x02\x12\x16\n" +
	"\tNOT_EQUAL\x10\x03\x1a\a\x9a\xb5\x18\x033.1\x1a\a\x92\xb5\x18\x033.0\"Y\n" +
	"\rCompareTarget\x12\v\n" +
	"\aVERSION\x10\x00\x12\n" +
	"\n" +
	"\x06CREATE\x10\x01\x12\a\n" +
	"\x03MOD\x10\x02\x12\t\n" +
	"\x05VALUE\x10\x03\x12\x12\n" +
	"\x05LEASE\x10\x04\x1a\a\x9a\xb5\x18\x033.3\x1a\a\x92\xb5\x18\x033.0:\a\x82\xb5\x18\x033.0B\x0e\n" +
	"\ftarget_union\"\xac\x01\n" +
	"\n" +
	"TxnRequest\x12/\n" +
	"\acompare\x18\x01 \x03(\v2\x15.etcdserverpb.CompareR\acompare\x121\n" +
	"\asuccess\x18\x02 \x03(\v2\x17.etcdserverpb.RequestOpR\asuccess\x121\n" +
	"\afailure\x18\x03 \x03(\v2\x17.etcdserverpb.RequestOpR\afailure:\a\x82\xb5\x18\x033.0\"\xa2\x01\n" +
	"\vTxnResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header\x12\x1c\n" +
	"\tsucceeded\x18\x02 \x01(\bR\tsucceeded\x126\n" +
	"\tresponses\x18\x03 \x03(\v2\x18.etcdserverpb.ResponseOpR\tresponses:\a\x82\xb5\x18\x033.0\"T\n" +
	"\x11CompactionRequest\x12\x1a\n" +
	"\brevision\x18\x01 \x01(\x03R\brevision\x12\x1a\n" +
	"\bphysical\x18\x02 \x01(\bR\bphysical:\a\x82\xb5\x18\x033.0\"S\n" +
	"\x12CompactionResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header:\a\x82\xb5\x18\x033.0\"\x16\n" +
	"\vHashRequest:\a\x82\xb5\x18\x033.0\"4\n" +
	"\rHashKVRequest\x12\x1a\n" +
	"\brevision\x18\x01 \x01(\x03R\brevision:\a\x82\xb5\x18\x033.3\"\xbc\x01\n" +
	"\x0eHashKVResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header\x12\x12\n" +
	"\x04hash\x18\x02 \x01(\rR\x04hash\x12)\n" +
	"\x10compact_revision\x18\x03 \x01(\x03R\x0fcompactRevision\x12,\n" +
	"\rhash_revision\x18\x04 \x01(\x03B\a\x8a\xb5\x18\x033.6R\fhashRevision:\a\x82\xb5\x18\x033.3\"a\n" +
	"\fHashResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header\x12\x12\n" +
	"\x04hash\x18\x02 \x01(\rR\x04hash:\a\x82\xb5\x18\x033.0\"\x1a\n" +
	"\x0fSnapshotRequest:\a\x82\xb5\x18\x033.3\"\xb1\x01\n" +
	"\x10SnapshotResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header\x12'\n" +
	"\x0fremaining_bytes\x18\x02 \x01(\x04R\x0eremainingBytes\x12\x12\n" +
	"\x04blob\x18\x03 \x01(\fR\x04blob\x12!\n" +
	"\aversion\x18\x04 \x01(\tB\a\x8a\xb5\x18\x033.6R\aversion:\a\x82\xb5\x18\x033.3\"\x98\x02\n" +
	"\fWatchRequest\x12I\n" +
	"\x0ecreate_request\x18\x01 \x01(\v2 .etcdserverpb.WatchCreateRequestH\x00R\rcreateRequest\x12I\n" +
	"\x0ecancel_request\x18\x02 \x01(\v2 .etcdserverpb.WatchCancelRequestH\x00R\rcancelRequest\x12X\n" +
	"\x10progress_request\x18\x03 \x01(\v2\".etcdserverpb.WatchProgressRequestB\a\x8a\xb5\x18\x033.4H\x00R\x0fprogressRequest:\a\x82\xb5\x18\x033.0B\x0f\n" +
	"\rrequest_union\"\x87\x03\n" +
	"\x12WatchCreateRequest\x12\x10\n" +
	"\x03key\x18\x01 \x01(\fR\x03key\x12\x1b\n" +
	"\trange_end\x18\x02 \x01(\fR\brangeEnd\x12%\n" +
	"\x0estart_revision\x18\x03 \x01(\x03R\rstartRevision\x12'\n" +
	"\x0fprogress_notify\x18\x04 \x01(\bR\x0eprogressNotify\x12N\n" +
	"\afilters\x18\x05 \x03(\x0e2+.etcdserverpb.WatchCreateRequest.FilterTypeB\a\x8a\xb5\x18\x033.1R\afilters\x12 \n" +
	"\aprev_kv\x18\x06 \x01(\bB\a\x8a\xb5\x18\x033.1R\x06prevKv\x12\"\n" +
	"\bwatch_id\x18\a \x01(\x03B\a\x8a\xb5\x18\x033.4R\awatchId\x12#\n" +
	"\bfragment\x18\b \x01(\bB\a\x8a\xb5\x18\x033.4R\bfragment\".\n" +
	"\n" +
	"FilterType\x12\t\n" +
	"\x05NOPUT\x10\x00\x12\f\n" +
	"\bNODELETE\x10\x01\x1a\a\x92\xb5\x18\x033.1:\a\x82\xb5\x18\x033.0\"A\n" +
	"\x12WatchCancelRequest\x12\"\n" +
	"\bwatch_id\x18\x01 \x01(\x03B\a\x8a\xb5\x18\x033.1R\awatchId:\a\x82\xb5\x18\x033.1\"\x1f\n" +
	"\x14WatchProgressRequest:\a\x82\xb5\x18\x033.4\"\xc4\x02\n" +
	"\rWatchResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header\x12\x19\n" +
	"\bwatch_id\x18\x02 \x01(\x03R\awatchId\x12\x18\n" +
	"\acreated\x18\x03 \x01(\bR\acreated\x12\x1a\n" +
	"\bcanceled\x18\x04 \x01(\bR\bcanceled\x12)\n" +
	"\x10compact_revision\x18\x05 \x01(\x03R\x0fcompactRevision\x12,\n" +
	"\rcancel_reason\x18\x06 \x01(\tB\a\x8a\xb5\x18\x033.4R\fcancelReason\x12#\n" +
	"\bfragment\x18\a \x01(\bB\a\x8a\xb5\x18\x033.4R\bfragment\x12%\n" +
	"\x06events\x18\v \x03(\v2\r.mvccpb.EventR\x06events:\a\x82\xb5\x18\x033.0\">\n" +
	"\x11LeaseGrantRequest\x12\x10\n" +
	"\x03TTL\x18\x01 \x01(\x03R\x03TTL\x12\x0e\n" +
	"\x02ID\x18\x02 \x01(\x03R\x02ID:\a\x82\xb5\x18\x033.0\"\x8b\x01\n" +
	"\x12LeaseGrantResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header\x12\x0e\n" +
	"\x02ID\x18\x02 \x01(\x03R\x02ID\x12\x10\n" +
	"\x03TTL\x18\x03 \x01(\x03R\x03TTL\x12\x14\n" +
	"\x05error\x18\x04 \x01(\tR\x05error:\a\x82\xb5\x18\x033.0\"-\n" +
	"\x12LeaseRevokeRequest\x12\x0e\n" +
	"\x02ID\x18\x01 \x01(\x03R\x02ID:\a\x82\xb5\x18\x033.0\"T\n" +
	"\x13LeaseRevokeResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header:\a\x82\xb5\x18\x033.0\"O\n" +
	"\x0fLeaseCheckpoint\x12\x0e\n" +
	"\x02ID\x18\x01 \x01(\x03R\x02ID\x12#\n" +
	"\rremaining_TTL\x18\x02 \x01(\x03R\fremainingTTL:\a\x82\xb5\x18\x033.4\"b\n" +
	"\x16LeaseCheckpointRequest\x12?\n" +
	"\vcheckpoints\x18\x01 \x03(\v2\x1d.etcdserverpb.LeaseCheckpointR\vcheckpoints:\a\x82\xb5\x18\x033.4\"X\n" +
	"\x17LeaseCheckpointResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header:\a\x82\xb5\x18\x033.4\"0\n" +
	"\x15LeaseKeepAliveRequest\x12\x0e\n" +
	"\x02ID\x18\x01 \x01(\x03R\x02ID:\a\x82\xb5\x18\x033.0\"y\n" +
	"\x16LeaseKeepAliveResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header\x12\x0e\n" +
	"\x02ID\x18\x02 \x01(\x03R\x02ID\x12\x10\n" +
	"\x03TTL\x18\x03 \x01(\x03R\x03TTL:\a\x82\xb5\x18\x033.0\"E\n" +
	"\x16LeaseTimeToLiveRequest\x12\x0e\n" +
	"\x02ID\x18\x01 \x01(\x03R\x02ID\x12\x12\n" +
	"\x04keys\x18\x02 \x01(\bR\x04keys:\a\x82\xb5\x18\x033.1\"\xae\x01\n" +
	"\x17LeaseTimeToLiveResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header\x12\x0e\n" +
	"\x02ID\x18\x02 \x01(\x03R\x02ID\x12\x10\n" +
	"\x03TTL\x18\x03 \x01(\x03R\x03TTL\x12\x1e\n" +
	"\n" +
	"grantedTTL\x18\x04 \x01(\x03R\n" +
	"grantedTTL\x12\x12\n" +
	"\x04keys\x18\x05 \x03(\fR\x04keys:\a\x82\xb5\x18\x033.1\"\x1d\n" +
	"\x12LeaseLeasesRequest:\a\x82\xb5\x18\x033.3\"&\n" +
	"\vLeaseStatus\x12\x0e\n" +
	"\x02ID\x18\x01 \x01(\x03R\x02ID:\a\x82\xb5\x18\x033.3\"\x87\x01\n" +
	"\x13LeaseLeasesResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header\x121\n" +
	"\x06leases\x18\x02 \x03(\v2\x19.etcdserverpb.LeaseStatusR\x06leases:\a\x82\xb5\x18\x033.3\"\x98\x01\n" +
	"\x06Member\x12\x0e\n" +
	"\x02ID\x18\x01 \x01(\x04R\x02ID\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x1a\n" +
	"\bpeerURLs\x18\x03 \x03(\tR\bpeerURLs\x12\x1e\n" +
	"\n" +
	"clientURLs\x18\x04 \x03(\tR\n" +
	"clientURLs\x12%\n" +
	"\tisLearner\x18\x05 \x01(\bB\a\x8a\xb5\x18\x033.4R\tisLearner:\a\x82\xb5\x18\x033.0\"^\n" +
	"\x10MemberAddRequest\x12\x1a\n" +
	"\bpeerURLs\x18\x01 \x03(\tR\bpeerURLs\x12%\n" +
	"\tisLearner\x18\x02 \x01(\bB\a\x8a\xb5\x18\x033.4R\tisLearner:\a\x82\xb5\x18\x033.0\"\xb0\x01\n" +
	"\x11MemberAddResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header\x12,\n" +
	"\x06member\x18\x02 \x01(\v2\x14.etcdserverpb.MemberR\x06member\x12.\n" +
	"\amembers\x18\x03 \x03(\v2\x14.etcdserverpb.MemberR\amembers:\a\x82\xb5\x18\x033.0\".\n" +
	"\x13MemberRemoveRequest\x12\x0e\n" +
	"\x02ID\x18\x01 \x01(\x04R\x02ID:\a\x82\xb5\x18\x033.0\"\x85\x01\n" +
	"\x14MemberRemoveResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header\x12.\n" +
	"\amembers\x18\x02 \x03(\v2\x14.etcdserverpb.MemberR\amembers:\a\x82\xb5\x18\x033.0\"J\n" +
	"\x13MemberUpdateRequest\x12\x0e\n" +
	"\x02ID\x18\x01 \x01(\x04R\x02ID\x12\x1a\n" +
	"\bpeerURLs\x18\x02 \x03(\tR\bpeerURLs:\a\x82\xb5\x18\x033.0\"\x8e\x01\n" +
	"\x14MemberUpdateResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header\x127\n" +
	"\amembers\x18\x02 \x03(\v2\x14.etcdserverpb.MemberB\a\x8a\xb5\x18\x033.1R\amembers:\a\x82\xb5\x18\x033.0\"I\n" +
	"\x11MemberListRequest\x12+\n" +
	"\flinearizable\x18\x01 \x01(\bB\a\x8a\xb5\x18\x033.5R\flinearizable:\a\x82\xb5\x18\x033.0\"\x83\x01\n" +
	"\x12MemberListResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header\x12.\n" +
	"\amembers\x18\x02 \x03(\v2\x14.etcdserverpb.MemberR\amembers:\a\x82\xb5\x18\x033.0\"/\n" +
	"\x14MemberPromoteRequest\x12\x0e\n" +
	"\x02ID\x18\x01 \x01(\x04R\x02ID:\a\x82\xb5\x18\x033.4\"\x86\x01\n" +
	"\x15MemberPromoteResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header\x12.\n" +
	"\amembers\x18\x02 \x03(\v2\x14.etcdserverpb.MemberR\amembers:\a\x82\xb5\x18\x033.4\"\x1c\n" +
	"\x11DefragmentRequest:\a\x82\xb5\x18\x033.0\"S\n" +
	"\x12DefragmentResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header:\a\x82\xb5\x18\x033.0\"8\n" +
	"\x11MoveLeaderRequest\x12\x1a\n" +
	"\btargetID\x18\x01 \x01(\x04R\btargetID:\a\x82\xb5\x18\x033.3\"S\n" +
	"\x12MoveLeaderResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header:\a\x82\xb5\x18\x033.3\"\xe1\x01\n" +
	"\fAlarmRequest\x12>\n" +
	"\x06action\x18\x01 \x01(\x0e2&.etcdserverpb.AlarmRequest.AlarmActionR\x06action\x12\x1a\n" +
	"\bmemberID\x18\x02 \x01(\x04R\bmemberID\x12-\n" +
	"\x05alarm\x18\x03 \x01(\x0e2\x17.etcdserverpb.AlarmTypeR\x05alarm\"=\n" +
	"\vAlarmAction\x12\a\n" +
	"\x03GET\x10\x00\x12\f\n" +
	"\bACTIVATE\x10\x01\x12\x0e\n" +
	"\n" +
	"DEACTIVATE\x10\x02\x1a\a\x92\xb5\x18\x033.0:\a\x82\xb5\x18\x033.0\"a\n" +
	"\vAlarmMember\x12\x1a\n" +
	"\bmemberID\x18\x01 \x01(\x04R\bmemberID\x12-\n" +
	"\x05alarm\x18\x02 \x01(\x0e2\x17.etcdserverpb.AlarmTypeR\x05alarm:\a\x82\xb5\x18\x033.0\"\x81\x01\n" +
	"\rAlarmResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header\x121\n" +
	"\x06alarms\x18\x02 \x03(\v2\x19.etcdserverpb.AlarmMemberR\x06alarms:\a\x82\xb5\x18\x033.0\"\xbf\x01\n" +
	"\x10DowngradeRequest\x12F\n" +
	"\x06action\x18\x01 \x01(\x0e2..etcdserverpb.DowngradeRequest.DowngradeActionR\x06action\x12\x18\n" +
	"\aversion\x18\x02 \x01(\tR\aversion\"@\n" +
	"\x0fDowngradeAction\x12\f\n" +
	"\bVALIDATE\x10\x00\x12\n" +
	"\n" +
	"\x06ENABLE\x10\x01\x12\n" +
	"\n" +
	"\x06CANCEL\x10\x02\x1a\a\x92\xb5\x18\x033.5:\a\x82\xb5\x18\x033.5\"l\n" +
	"\x11DowngradeResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header\x12\x18\n" +
	"\aversion\x18\x02 \x01(\tR\aversion:\a\x82\xb5\x18\x033.5\"8\n" +
	"\x1bDowngradeVersionTestRequest\x12\x10\n" +
	"\x03ver\x18\x01 \x01(\tR\x03ver:\a\x82\xb5\x18\x033.6\"\x18\n" +
	"\rStatusRequest:\a\x82\xb5\x18\x033.0\"\xa3\x04\n" +
	"\x0eStatusResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header\x12\x18\n" +
	"\aversion\x18\x02 \x01(\tR\aversion\x12\x16\n" +
	"\x06dbSize\x18\x03 \x01(\x03R\x06dbSize\x12\x16\n" +
	"\x06leader\x18\x04 \x01(\x04R\x06leader\x12\x1c\n" +
	"\traftIndex\x18\x05 \x01(\x04R\traftIndex\x12\x1a\n" +
	"\braftTerm\x18\x06 \x01(\x04R\braftTerm\x123\n" +
	"\x10raftAppliedIndex\x18\a \x01(\x04B\a\x8a\xb5\x18\x033.4R\x10raftAppliedIndex\x12\x1f\n" +
	"\x06errors\x18\b \x03(\tB\a\x8a\xb5\x18\x033.4R\x06errors\x12)\n" +
	"\vdbSizeInUse\x18\t \x01(\x03B\a\x8a\xb5\x18\x033.4R\vdbSizeInUse\x12%\n" +
	"\tisLearner\x18\n" +
	" \x01(\bB\a\x8a\xb5\x18\x033.4R\tisLearner\x12/\n" +
	"\x0estorageVersion\x18\v \x01(\tB\a\x8a\xb5\x18\x033.6R\x0estorageVersion\x12)\n" +
	"\vdbSizeQuota\x18\f \x01(\x03B\a\x8a\xb5\x18\x033.6R\vdbSizeQuota\x12J\n" +
	"\rdowngradeInfo\x18\r \x01(\v2\x1b.etcdserverpb.DowngradeInfoB\a\x8a\xb5\x18\x033.6R\rdowngradeInfo:\a\x82\xb5\x18\x033.0\"O\n" +
	"\rDowngradeInfo\x12\x18\n" +
	"\aenabled\x18\x01 \x01(\bR\aenabled\x12$\n" +
	"\rtargetVersion\x18\x02 \x01(\tR\rtargetVersion\"\x1c\n" +
	"\x11AuthEnableRequest:\a\x82\xb5\x18\x033.0\"\x1d\n" +
	"\x12AuthDisableRequest:\a\x82\xb5\x18\x033.0\"\x1c\n" +
	"\x11AuthStatusRequest:\a\x82\xb5\x18\x033.5\"N\n" +
	"\x13AuthenticateRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword:\a\x82\xb5\x18\x033.0\"\xb9\x01\n" +
	"\x12AuthUserAddRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\x129\n" +
	"\aoptions\x18\x03 \x01(\v2\x16.authpb.UserAddOptionsB\a\x8a\xb5\x18\x033.4R\aoptions\x12/\n" +
	"\x0ehashedPassword\x18\x04 \x01(\tB\a\x8a\xb5\x18\x033.5R\x0ehashedPassword:\a\x82\xb5\x18\x033.0\"1\n" +
	"\x12AuthUserGetRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name:\a\x82\xb5\x18\x033.0\"4\n" +
	"\x15AuthUserDeleteRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name:\a\x82\xb5\x18\x033.0\"\x89\x01\n" +
	"\x1dAuthUserChangePasswordRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\x12/\n" +
	"\x0ehashedPassword\x18\x03 \x01(\tB\a\x8a\xb5\x18\x033.5R\x0ehashedPassword:\a\x82\xb5\x18\x033.0\"K\n" +
	"\x18AuthUserGrantRoleRequest\x12\x12\n" +
	"\x04user\x18\x01 \x01(\tR\x04user\x12\x12\n" +
	"\x04role\x18\x02 \x01(\tR\x04role:\a\x82\xb5\x18\x033.0\"L\n" +
	"\x19AuthUserRevokeRoleRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x12\n" +
	"\x04role\x18\x02 \x01(\tR\x04role:\a\x82\xb5\x18\x033.0\"1\n" +
	"\x12AuthRoleAddRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name:\a\x82\xb5\x18\x033.0\"1\n" +
	"\x12AuthRoleGetRequest\x12\x12\n" +
	"\x04role\x18\x01 \x01(\tR\x04role:\a\x82\xb5\x18\x033.0\"\x1e\n" +
	"\x13AuthUserListRequest:\a\x82\xb5\x18\x033.0\"\x1e\n" +
	"\x13AuthRoleListRequest:\a\x82\xb5\x18\x033.0\"4\n" +
	"\x15AuthRoleDeleteRequest\x12\x12\n" +
	"\x04role\x18\x01 \x01(\tR\x04role:\a\x82\xb5\x18\x033.0\"e\n" +
	"\x1eAuthRoleGrantPermissionRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12&\n" +
	"\x04perm\x18\x02 \x01(\v2\x12.authpb.PermissionR\x04perm:\a\x82\xb5\x18\x033.0\"m\n" +
	"\x1fAuthRoleRevokePermissionRequest\x12\x12\n" +
	"\x04role\x18\x01 \x01(\tR\x04role\x12\x10\n" +
	"\x03key\x18\x02 \x01(\fR\x03key\x12\x1b\n" +
	"\trange_end\x18\x03 \x01(\fR\brangeEnd:\a\x82\xb5\x18\x033.0\"S\n" +
	"\x12AuthEnableResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header:\a\x82\xb5\x18\x033.0\"T\n" +
	"\x13AuthDisableResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header:\a\x82\xb5\x18\x033.0\"\x91\x01\n" +
	"\x12AuthStatusResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header\x12\x18\n" +
	"\aenabled\x18\x02 \x01(\bR\aenabled\x12\"\n" +
	"\fauthRevision\x18\x03 \x01(\x04R\fauthRevision:\a\x82\xb5\x18\x033.5\"k\n" +
	"\x14AuthenticateResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header\x12\x14\n" +
	"\x05token\x18\x02 \x01(\tR\x05token:\a\x82\xb5\x18\x033.0\"T\n" +
	"\x13AuthUserAddResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header:\a\x82\xb5\x18\x033.0\"j\n" +
	"\x13AuthUserGetResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header\x12\x14\n" +
	"\x05roles\x18\x02 \x03(\tR\x05roles:\a\x82\xb5\x18\x033.0\"W\n" +
	"\x16AuthUserDeleteResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header:\a\x82\xb5\x18\x033.0\"_\n" +
	"\x1eAuthUserChangePasswordResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header:\a\x82\xb5\x18\x033.0\"Z\n" +
	"\x19AuthUserGrantRoleResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header:\a\x82\xb5\x18\x033.0\"[\n" +
	"\x1aAuthUserRevokeRoleResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header:\a\x82\xb5\x18\x033.0\"T\n" +
	"\x13AuthRoleAddResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header:\a\x82\xb5\x18\x033.0\"\x85\x01\n" +
	"\x13AuthRoleGetResponse\x12=\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderB\a\x8a\xb5\x18\x033.0R\x06header\x12/\n" +
	"\x04perm\x18\x02 \x03(\v2\x12.authpb.PermissionB\a\x8a\xb5\x18\x033.0R\x04perm\"k\n" +
	"\x14AuthRoleListResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header\x12\x14\n" +
	"\x05roles\x18\x02 \x03(\tR\x05roles:\a\x82\xb5\x18\x033.0\"k\n" +
	"\x14AuthUserListResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header\x12\x14\n" +
	"\x05users\x18\x02 \x03(\tR\x05users:\a\x82\xb5\x18\x033.0\"W\n" +
	"\x16AuthRoleDeleteResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header:\a\x82\xb5\x18\x033.0\"`\n" +
	"\x1fAuthRoleGrantPermissionResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header:\a\x82\xb5\x18\x033.0\"a\n" +
	" AuthRoleRevokePermissionResponse\x124\n" +
	"\x06header\x18\x01 \x01(\v2\x1c.etcdserverpb.ResponseHeaderR\x06header:\a\x82\xb5\x18\x033.0\"b\n" +
	"\x13RangeStreamResponse\x12B\n" +
	"\x0erange_response\x18\x01 \x01(\v2\x1b.etcdserverpb.RangeResponseR\rrangeResponse:\a\x82\xb5\x18\x033.7*A\n" +
	"\tAlarmType\x12\b\n" +
	"\x04NONE\x10\x00\x12\v\n" +
	"\aNOSPACE\x10\x01\x12\x14\n" +
	"\aCORRUPT\x10\x02\x1a\a\x9a\xb5\x18\x033.3\x1a\a\x92\xb5\x18\x033.02\xb6\x04\n" +
	"\x02KV\x12Y\n" +
	"\x05Range\x12\x1a.etcdserverpb.RangeRequest\x1a\x1b.etcdserverpb.RangeResponse\"\x17\x82\xd3\xe4\x93\x02\x11:\x01*\"\f/v3/kv/range\x12P\n" +
	"\vRangeStream\x12\x1a.etcdserverpb.RangeRequest\x1a!.etcdserverpb.RangeStreamResponse\"\x000\x01\x12Q\n" +
	"\x03Put\x12\x18.etcdserverpb.PutRequest\x1a\x19.etcdserverpb.PutResponse\"\x15\x82\xd3\xe4\x93\x02\x0f:\x01*\"\n" +
	"/v3/kv/put\x12q\n" +
	"\vDeleteRange\x12 .etcdserverpb.DeleteRangeRequest\x1a!.etcdserverpb.DeleteRangeResponse\"\x1d\x82\xd3\xe4\x93\x02\x17:\x01*\"\x12/v3/kv/deleterange\x12Q\n" +
	"\x03Txn\x12\x18.etcdserverpb.TxnRequest\x1a\x19.etcdserverpb.TxnResponse\"\x15\x82\xd3\xe4\x93\x02\x0f:\x01*\"\n" +
	"/v3/kv/txn\x12j\n" +
	"\aCompact\x12\x1f.etcdserverpb.CompactionRequest\x1a .etcdserverpb.CompactionResponse\"\x1c\x82\xd3\xe4\x93\x02\x16:\x01*\"\x11/v3/kv/compaction2c\n" +
	"\x05Watch\x12Z\n" +
	"\x05Watch\x12\x1a.etcdserverpb.WatchRequest\x1a\x1b.etcdserverpb.WatchResponse\"\x14\x82\xd3\xe4\x93\x02\x0e:\x01*\"\t/v3/watch(\x010\x012\xad\x05\n" +
	"\x05Lease\x12k\n" +
	"\n" +
	"LeaseGrant\x12\x1f.etcdserverpb.LeaseGrantRequest\x1a .etcdserverpb.LeaseGrantResponse\"\x1a\x82\xd3\xe4\x93\x02\x14:\x01*\"\x0f/v3/lease/grant\x12\x89\x01\n" +
	"\vLeaseRevoke\x12 .etcdserverpb.LeaseRevokeRequest\x1a!.etcdserverpb.LeaseRevokeResponse\"5\x82\xd3\xe4\x93\x02/:\x01*Z\x18:\x01*\"\x13/v3/kv/lease/revoke\"\x10/v3/lease/revoke\x12\x7f\n" +
	"\x0eLeaseKeepAlive\x12#.etcdserverpb.LeaseKeepAliveRequest\x1a$.etcdserverpb.LeaseKeepAliveResponse\"\x1e\x82\xd3\xe4\x93\x02\x18:\x01*\"\x13/v3/lease/keepalive(\x010\x01\x12\x9d\x01\n" +
	"\x0fLeaseTimeToLive\x12$.etcdserverpb.LeaseTimeToLiveRequest\x1a%.etcdserverpb.LeaseTimeToLiveResponse\"=\x82\xd3\xe4\x93\x027:\x01*Z\x1c:\x01*\"\x17/v3/kv/lease/timetolive\"\x14/v3/lease/timetolive\x12\x89\x01\n" +
	"\vLeaseLeases\x12 .etcdserverpb.LeaseLeasesRequest\x1a!.etcdserverpb.LeaseLeasesResponse\"5\x82\xd3\xe4\x93\x02/:\x01*Z\x18:\x01*\"\x13/v3/kv/lease/leases\"\x10/v3/lease/leases2\xea\x04\n" +
	"\aCluster\x12o\n" +
	"\tMemberAdd\x12\x1e.etcdserverpb.MemberAddRequest\x1a\x1f.etcdserverpb.MemberAddResponse\"!\x82\xd3\xe4\x93\x02\x1b:\x01*\"\x16/v3/cluster/member/add\x12{\n" +
	"\fMemberRemove\x12!.etcdserverpb.MemberRemoveRequest\x1a\".etcdserverpb.MemberRemoveResponse\"$\x82\xd3\xe4\x93\x02\x1e:\x01*\"\x19/v3/cluster/member/remove\x12{\n" +
	"\fMemberUpdate\x12!.etcdserverpb.MemberUpdateRequest\x1a\".etcdserverpb.MemberUpdateResponse\"$\x82\xd3\xe4\x93\x02\x1e:\x01*\"\x19/v3/cluster/member/update\x12s\n" +
	"\n" +
	"MemberList\x12\x1f.etcdserverpb.MemberListRequest\x1a .etcdserverpb.MemberListResponse\"\"\x82\xd3\xe4\x93\x02\x1c:\x01*\"\x17/v3/cluster/member/list\x12\x7f\n" +
	"\rMemberPromote\x12\".etcdserverpb.MemberPromoteRequest\x1a#.etcdserverpb.MemberPromoteResponse\"%\x82\xd3\xe4\x93\x02\x1f:\x01*\"\x1a/v3/cluster/member/promote2\x80\a\n" +
	"\vMaintenance\x12b\n" +
	"\x05Alarm\x12\x1a.etcdserverpb.AlarmRequest\x1a\x1b.etcdserverpb.AlarmResponse\" \x82\xd3\xe4\x93\x02\x1a:\x01*\"\x15/v3/maintenance/alarm\x12f\n" +
	"\x06Status\x12\x1b.etcdserverpb.StatusRequest\x1a\x1c.etcdserverpb.StatusResponse\"!\x82\xd3\xe4\x93\x02\x1b:\x01*\"\x16/v3/maintenance/status\x12v\n" +
	"\n" +
	"Defragment\x12\x1f.etcdserverpb.DefragmentRequest\x1a .etcdserverpb.DefragmentResponse\"%\x82\xd3\xe4\x93\x02\x1f:\x01*\"\x1a/v3/maintenance/defragment\x12^\n" +
	"\x04Hash\x12\x19.etcdserverpb.HashRequest\x1a\x1a.etcdserverpb.HashResponse\"\x1f\x82\xd3\xe4\x93\x02\x19:\x01*\"\x14/v3/maintenance/hash\x12f\n" +
	"\x06HashKV\x12\x1b.etcdserverpb.HashKVRequest\x1a\x1c.etcdserverpb.HashKVResponse\"!\x82\xd3\xe4\x93\x02\x1b:\x01*\"\x16/v3/maintenance/hashkv\x12p\n" +
	"\bSnapshot\x12\x1d.etcdserverpb.SnapshotRequest\x1a\x1e.etcdserverpb.SnapshotResponse\"#\x82\xd3\xe4\x93\x02\x1d:\x01*\"\x18/v3/maintenance/snapshot0\x01\x12\x7f\n" +
	"\n" +
	"MoveLeader\x12\x1f.etcdserverpb.MoveLeaderRequest\x1a .etcdserverpb.MoveLeaderResponse\".\x82\xd3\xe4\x93\x02(:\x01*\"#/v3/maintenance/transfer-leadership\x12r\n" +
	"\tDowngrade\x12\x1e.etcdserverpb.DowngradeRequest\x1a\x1f.etcdserverpb.DowngradeResponse\"$\x82\xd3\xe4\x93\x02\x1e:\x01*\"\x19/v3/maintenance/downgrade2\xa7\x10\n" +
	"\x04Auth\x12k\n" +
	"\n" +
	"AuthEnable\x12\x1f.etcdserverpb.AuthEnableRequest\x1a .etcdserverpb.AuthEnableResponse\"\x1a\x82\xd3\xe4\x93\x02\x14:\x01*\"\x0f/v3/auth/enable\x12o\n" +
	"\vAuthDisable\x12 .etcdserverpb.AuthDisableRequest\x1a!.etcdserverpb.AuthDisableResponse\"\x1b\x82\xd3\xe4\x93\x02\x15:\x01*\"\x10/v3/auth/disable\x12k\n" +
	"\n" +
	"AuthStatus\x12\x1f.etcdserverpb.AuthStatusRequest\x1a .etcdserverpb.AuthStatusResponse\"\x1a\x82\xd3\xe4\x93\x02\x14:\x01*\"\x0f/v3/auth/status\x12w\n" +
	"\fAuthenticate\x12!.etcdserverpb.AuthenticateRequest\x1a\".etcdserverpb.AuthenticateResponse\" \x82\xd3\xe4\x93\x02\x1a:\x01*\"\x15/v3/auth/authenticate\x12l\n" +
	"\aUserAdd\x12 .etcdserverpb.AuthUserAddRequest\x1a!.etcdserverpb.AuthUserAddResponse\"\x1c\x82\xd3\xe4\x93\x02\x16:\x01*\"\x11/v3/auth/user/add\x12l\n" +
	"\aUserGet\x12 .etcdserverpb.AuthUserGetRequest\x1a!.etcdserverpb.AuthUserGetResponse\"\x1c\x82\xd3\xe4\x93\x02\x16:\x01*\"\x11/v3/auth/user/get\x12p\n" +
	"\bUserList\x12!.etcdserverpb.AuthUserListRequest\x1a\".etcdserverpb.AuthUserListResponse\"\x1d\x82\xd3\xe4\x93\x02\x17:\x01*\"\x12/v3/auth/user/list\x12x\n" +
	"\n" +
	"UserDelete\x12#.etcdserverpb.AuthUserDeleteRequest\x1a$.etcdserverpb.AuthUserDeleteResponse\"\x1f\x82\xd3\xe4\x93\x02\x19:\x01*\"\x14/v3/auth/user/delete\x12\x92\x01\n" +
	"\x12UserChangePassword\x12+.etcdserverpb.AuthUserChangePasswordRequest\x1a,.etcdserverpb.AuthUserChangePasswordResponse\"!\x82\xd3\xe4\x93\x02\x1b:\x01*\"\x16/v3/auth/user/changepw\x12\x80\x01\n" +
	"\rUserGrantRole\x12&.etcdserverpb.AuthUserGrantRoleRequest\x1a'.etcdserverpb.AuthUserGrantRoleResponse\"\x1e\x82\xd3\xe4\x93\x02\x18:\x01*\"\x13/v3/auth/user/grant\x12\x84\x01\n" +
	"\x0eUserRevokeRole\x12'.etcdserverpb.AuthUserRevokeRoleRequest\x1a(.etcdserverpb.AuthUserRevokeRoleResponse\"\x1f\x82\xd3\xe4\x93\x02\x19:\x01*\"\x14/v3/auth/user/revoke\x12l\n" +
	"\aRoleAdd\x12 .etcdserverpb.AuthRoleAddRequest\x1a!.etcdserverpb.AuthRoleAddResponse\"\x1c\x82\xd3\xe4\x93\x02\x16:\x01*\"\x11/v3/auth/role/add\x12l\n" +
	"\aRoleGet\x12 .etcdserverpb.AuthRoleGetRequest\x1a!.etcdserverpb.AuthRoleGetResponse\"\x1c\x82\xd3\xe4\x93\x02\x16:\x01*\"\x11/v3/auth/role/get\x12p\n" +
	"\bRoleList\x12!.etcdserverpb.AuthRoleListRequest\x1a\".etcdserverpb.AuthRoleListResponse\"\x1d\x82\xd3\xe4\x93\x02\x17:\x01*\"\x12/v3/auth/role/list\x12x\n" +
	"\n" +
	"RoleDelete\x12#.etcdserverpb.AuthRoleDeleteRequest\x1a$.etcdserverpb.AuthRoleDeleteResponse\"\x1f\x82\xd3\xe4\x93\x02\x19:\x01*\"\x14/v3/auth/role/delete\x12\x92\x01\n" +
	"\x13RoleGrantPermission\x12,.etcdserverpb.AuthRoleGrantPermissionRequest\x1a-.etcdserverpb.AuthRoleGrantPermissionResponse\"\x1e\x82\xd3\xe4\x93\x02\x18:\x01*\"\x13/v3/auth/role/grant\x12\x96\x01\n" +
	"\x14RoleRevokePermission\x12-.etcdserverpb.AuthRoleRevokePermissionRequest\x1a..etcdserverpb.AuthRoleRevokePermissionResponse\"\x1f\x82\xd3\xe4\x93\x02\x19:\x01*\"\x14/v3/auth/role/revokeBW\x92A/Z\x1f\n" +
	"\x1d\n" +
	"\x06ApiKey\x12\x13\b\x02\x1a\rAuthorization \x02b\f\n" +
	"\n" +
	"\n" +
	"\x06ApiKey\x12\x00Z#go.etcd.io/etcd/api/v3/etcdserverpbb\x06proto3"
var (
	file_rpc_proto_rawDescOnce sync.Once
	file_rpc_proto_rawDescData []byte
)
func file_rpc_proto_rawDescGZIP() []byte {
	file_rpc_proto_rawDescOnce.Do(func() {
		file_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_rpc_proto_rawDesc), len(file_rpc_proto_rawDesc)))
	})
	return file_rpc_proto_rawDescData
}
var file_rpc_proto_enumTypes = make([]protoimpl.EnumInfo, 8)
var file_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 98)
var file_rpc_proto_goTypes = []any{
	(AlarmType)(0),                           
	(RangeRequest_SortOrder)(0),              
	(RangeRequest_SortTarget)(0),             
	(Compare_CompareResult)(0),               
	(Compare_CompareTarget)(0),               
	(WatchCreateRequest_FilterType)(0),       
	(AlarmRequest_AlarmAction)(0),            
	(DowngradeRequest_DowngradeAction)(0),    
	(*ResponseHeader)(nil),                   
	(*RangeRequest)(nil),                     
	(*RangeResponse)(nil),                    
	(*PutRequest)(nil),                       
	(*PutResponse)(nil),                      
	(*DeleteRangeRequest)(nil),               
	(*DeleteRangeResponse)(nil),              
	(*RequestOp)(nil),                        
	(*ResponseOp)(nil),                       
	(*Compare)(nil),                          
	(*TxnRequest)(nil),                       
	(*TxnResponse)(nil),                      
	(*CompactionRequest)(nil),                
	(*CompactionResponse)(nil),               
	(*HashRequest)(nil),                      
	(*HashKVRequest)(nil),                    
	(*HashKVResponse)(nil),                   
	(*HashResponse)(nil),                     
	(*SnapshotRequest)(nil),                  
	(*SnapshotResponse)(nil),                 
	(*WatchRequest)(nil),                     
	(*WatchCreateRequest)(nil),               
	(*WatchCancelRequest)(nil),               
	(*WatchProgressRequest)(nil),             
	(*WatchResponse)(nil),                    
	(*LeaseGrantRequest)(nil),                
	(*LeaseGrantResponse)(nil),               
	(*LeaseRevokeRequest)(nil),               
	(*LeaseRevokeResponse)(nil),              
	(*LeaseCheckpoint)(nil),                  
	(*LeaseCheckpointRequest)(nil),           
	(*LeaseCheckpointResponse)(nil),          
	(*LeaseKeepAliveRequest)(nil),            
	(*LeaseKeepAliveResponse)(nil),           
	(*LeaseTimeToLiveRequest)(nil),           
	(*LeaseTimeToLiveResponse)(nil),          
	(*LeaseLeasesRequest)(nil),               
	(*LeaseStatus)(nil),                      
	(*LeaseLeasesResponse)(nil),              
	(*Member)(nil),                           
	(*MemberAddRequest)(nil),                 
	(*MemberAddResponse)(nil),                
	(*MemberRemoveRequest)(nil),              
	(*MemberRemoveResponse)(nil),             
	(*MemberUpdateRequest)(nil),              
	(*MemberUpdateResponse)(nil),             
	(*MemberListRequest)(nil),                
	(*MemberListResponse)(nil),               
	(*MemberPromoteRequest)(nil),             
	(*MemberPromoteResponse)(nil),            
	(*DefragmentRequest)(nil),                
	(*DefragmentResponse)(nil),               
	(*MoveLeaderRequest)(nil),                
	(*MoveLeaderResponse)(nil),               
	(*AlarmRequest)(nil),                     
	(*AlarmMember)(nil),                      
	(*AlarmResponse)(nil),                    
	(*DowngradeRequest)(nil),                 
	(*DowngradeResponse)(nil),                
	(*DowngradeVersionTestRequest)(nil),      
	(*StatusRequest)(nil),                    
	(*StatusResponse)(nil),                   
	(*DowngradeInfo)(nil),                    
	(*AuthEnableRequest)(nil),                
	(*AuthDisableRequest)(nil),               
	(*AuthStatusRequest)(nil),                
	(*AuthenticateRequest)(nil),              
	(*AuthUserAddRequest)(nil),               
	(*AuthUserGetRequest)(nil),               
	(*AuthUserDeleteRequest)(nil),            
	(*AuthUserChangePasswordRequest)(nil),    
	(*AuthUserGrantRoleRequest)(nil),         
	(*AuthUserRevokeRoleRequest)(nil),        
	(*AuthRoleAddRequest)(nil),               
	(*AuthRoleGetRequest)(nil),               
	(*AuthUserListRequest)(nil),              
	(*AuthRoleListRequest)(nil),              
	(*AuthRoleDeleteRequest)(nil),            
	(*AuthRoleGrantPermissionRequest)(nil),   
	(*AuthRoleRevokePermissionRequest)(nil),  
	(*AuthEnableResponse)(nil),               
	(*AuthDisableResponse)(nil),              
	(*AuthStatusResponse)(nil),               
	(*AuthenticateResponse)(nil),             
	(*AuthUserAddResponse)(nil),              
	(*AuthUserGetResponse)(nil),              
	(*AuthUserDeleteResponse)(nil),           
	(*AuthUserChangePasswordResponse)(nil),   
	(*AuthUserGrantRoleResponse)(nil),        
	(*AuthUserRevokeRoleResponse)(nil),       
	(*AuthRoleAddResponse)(nil),              
	(*AuthRoleGetResponse)(nil),              
	(*AuthRoleListResponse)(nil),             
	(*AuthUserListResponse)(nil),             
	(*AuthRoleDeleteResponse)(nil),           
	(*AuthRoleGrantPermissionResponse)(nil),  
	(*AuthRoleRevokePermissionResponse)(nil), 
	(*RangeStreamResponse)(nil),              
	(*mvccpb.KeyValue)(nil),                  
	(*mvccpb.Event)(nil),                     
	(*authpb.UserAddOptions)(nil),            
	(*authpb.Permission)(nil),                
}
var file_rpc_proto_depIdxs = []int32{
	1,   
	2,   
	8,   
	106, 
	8,   
	106, 
	8,   
	106, 
	9,   
	11,  
	13,  
	18,  
	10,  
	12,  
	14,  
	19,  
	3,   
	4,   
	17,  
	15,  
	15,  
	8,   
	16,  
	8,   
	8,   
	8,   
	8,   
	29,  
	30,  
	31,  
	5,   
	8,   
	107, 
	8,   
	8,   
	37,  
	8,   
	8,   
	8,   
	8,   
	45,  
	8,   
	47,  
	47,  
	8,   
	47,  
	8,   
	47,  
	8,   
	47,  
	8,   
	47,  
	8,   
	8,   
	6,   
	0,   
	0,   
	8,   
	63,  
	7,   
	8,   
	8,   
	70,  
	108, 
	109, 
	8,   
	8,   
	8,   
	8,   
	8,   
	8,   
	8,   
	8,   
	8,   
	8,   
	8,   
	8,   
	109, 
	8,   
	8,   
	8,   
	8,   
	8,   
	10,  
	9,   
	9,   
	11,  
	13,  
	18,  
	20,  
	28,  
	33,  
	35,  
	40,  
	42,  
	44,  
	48,  
	50,  
	52,  
	54,  
	56,  
	62,  
	68,  
	58,  
	22,  
	23,  
	26,  
	60,  
	65,  
	71,  
	72,  
	73,  
	74,  
	75,  
	76,  
	83,  
	77,  
	78,  
	79,  
	80,  
	81,  
	82,  
	84,  
	85,  
	86,  
	87,  
	10,  
	105, 
	12,  
	14,  
	19,  
	21,  
	32,  
	34,  
	36,  
	41,  
	43,  
	46,  
	49,  
	51,  
	53,  
	55,  
	57,  
	64,  
	69,  
	59,  
	25,  
	24,  
	27,  
	61,  
	66,  
	88,  
	89,  
	90,  
	91,  
	92,  
	93,  
	101, 
	94,  
	95,  
	96,  
	97,  
	98,  
	99,  
	100, 
	102, 
	103, 
	104, 
	126, 
	84,  
	84,  
	84,  
	0,   
}
func init() { file_rpc_proto_init() }
func file_rpc_proto_init() {
	if File_rpc_proto != nil {
		return
	}
	file_rpc_proto_msgTypes[7].OneofWrappers = []any{
		(*RequestOp_RequestRange)(nil),
		(*RequestOp_RequestPut)(nil),
		(*RequestOp_RequestDeleteRange)(nil),
		(*RequestOp_RequestTxn)(nil),
	}
	file_rpc_proto_msgTypes[8].OneofWrappers = []any{
		(*ResponseOp_ResponseRange)(nil),
		(*ResponseOp_ResponsePut)(nil),
		(*ResponseOp_ResponseDeleteRange)(nil),
		(*ResponseOp_ResponseTxn)(nil),
	}
	file_rpc_proto_msgTypes[9].OneofWrappers = []any{
		(*Compare_Version)(nil),
		(*Compare_CreateRevision)(nil),
		(*Compare_ModRevision)(nil),
		(*Compare_Value)(nil),
		(*Compare_Lease)(nil),
	}
	file_rpc_proto_msgTypes[20].OneofWrappers = []any{
		(*WatchRequest_CreateRequest)(nil),
		(*WatchRequest_CancelRequest)(nil),
		(*WatchRequest_ProgressRequest)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_rpc_proto_rawDesc), len(file_rpc_proto_rawDesc)),
			NumEnums:      8,
			NumMessages:   98,
			NumExtensions: 0,
			NumServices:   6,
		},
		GoTypes:           file_rpc_proto_goTypes,
		DependencyIndexes: file_rpc_proto_depIdxs,
		EnumInfos:         file_rpc_proto_enumTypes,
		MessageInfos:      file_rpc_proto_msgTypes,
	}.Build()
	File_rpc_proto = out.File
	file_rpc_proto_goTypes = nil
	file_rpc_proto_depIdxs = nil
}
