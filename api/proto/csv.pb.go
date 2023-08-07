// Code generated by protoc-gen-go. DO NOT EDIT.
// source: csv.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	proto "www.velocidex.com/golang/velociraptor/artifacts/proto"
	_ "www.velocidex.com/golang/velociraptor/proto"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetTableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Paging parameters.
	Rows     uint64 `protobuf:"varint,2,opt,name=rows,proto3" json:"rows,omitempty"`                         // Number of rows to fetch
	StartRow uint64 `protobuf:"varint,3,opt,name=start_row,json=startRow,proto3" json:"start_row,omitempty"` // First row to fetch.
	// For event artifacts we can get events after this time. Event
	// artifacts should specify the artifact name with client_id being
	// either "server" for server events or the client id for the
	// client events. Number of seconds since epoch.
	StartTime int64 `protobuf:"varint,13,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   int64 `protobuf:"varint,14,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// For collected artifacts tables.
	ClientId string `protobuf:"bytes,4,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	FlowId   string `protobuf:"bytes,5,opt,name=flow_id,json=flowId,proto3" json:"flow_id,omitempty"`
	Artifact string `protobuf:"bytes,6,opt,name=artifact,proto3" json:"artifact,omitempty"`
	// Can be log, uploads for collection additional tables.
	Type string `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	// For collected hunts. With hunts, type can be clients, hunt_status.
	HuntId string `protobuf:"bytes,8,opt,name=hunt_id,json=huntId,proto3" json:"hunt_id,omitempty"`
	// For notebook tables.
	NotebookId string `protobuf:"bytes,9,opt,name=notebook_id,json=notebookId,proto3" json:"notebook_id,omitempty"`
	CellId     string `protobuf:"bytes,10,opt,name=cell_id,json=cellId,proto3" json:"cell_id,omitempty"`
	TableId    int64  `protobuf:"varint,11,opt,name=table_id,json=tableId,proto3" json:"table_id,omitempty"`
	// For timelines
	Timeline string `protobuf:"bytes,16,opt,name=timeline,proto3" json:"timeline,omitempty"`
	// Skip these timeline components.
	SkipComponents []string `protobuf:"bytes,17,rep,name=skip_components,json=skipComponents,proto3" json:"skip_components,omitempty"`
	// For download handler when creating an export file - control
	// output format. Can be "csv", "jsonl"
	DownloadFormat string `protobuf:"bytes,12,opt,name=download_format,json=downloadFormat,proto3" json:"download_format,omitempty"`
	// Optionally for downloads, the caller may specify the filename.
	DownloadFilename string `protobuf:"bytes,18,opt,name=download_filename,json=downloadFilename,proto3" json:"download_filename,omitempty"`
	// If specified only emit these columns.
	Columns []string `protobuf:"bytes,15,rep,name=columns,proto3" json:"columns,omitempty"`
	// If specified, transform the table first.
	SortColumn    string `protobuf:"bytes,19,opt,name=sort_column,json=sortColumn,proto3" json:"sort_column,omitempty"`
	SortDirection bool   `protobuf:"varint,20,opt,name=sort_direction,json=sortDirection,proto3" json:"sort_direction,omitempty"`
	FilterColumn  string `protobuf:"bytes,21,opt,name=filter_column,json=filterColumn,proto3" json:"filter_column,omitempty"`
	FilterRegex   string `protobuf:"bytes,22,opt,name=filter_regex,json=filterRegex,proto3" json:"filter_regex,omitempty"`
	// This transformation takes a range from the larger result set
	// and pages within that range.
	StartIdx uint64 `protobuf:"varint,27,opt,name=start_idx,json=startIdx,proto3" json:"start_idx,omitempty"`
	EndIdx   uint64 `protobuf:"varint,28,opt,name=end_idx,json=endIdx,proto3" json:"end_idx,omitempty"`
	// The org id may be specified in the query string - The protobuf
	// is normally parsed from the query string directly.
	OrgId string `protobuf:"bytes,23,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// The required timezone to export in.
	Timezone string `protobuf:"bytes,24,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Version  uint64 `protobuf:"varint,25,opt,name=version,proto3" json:"version,omitempty"`
	// Used for VFS components
	VfsComponents []string `protobuf:"bytes,26,rep,name=vfs_components,json=vfsComponents,proto3" json:"vfs_components,omitempty"`
}

func (x *GetTableRequest) Reset() {
	*x = GetTableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_csv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTableRequest) ProtoMessage() {}

func (x *GetTableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_csv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTableRequest.ProtoReflect.Descriptor instead.
func (*GetTableRequest) Descriptor() ([]byte, []int) {
	return file_csv_proto_rawDescGZIP(), []int{0}
}

func (x *GetTableRequest) GetRows() uint64 {
	if x != nil {
		return x.Rows
	}
	return 0
}

func (x *GetTableRequest) GetStartRow() uint64 {
	if x != nil {
		return x.StartRow
	}
	return 0
}

func (x *GetTableRequest) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *GetTableRequest) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *GetTableRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *GetTableRequest) GetFlowId() string {
	if x != nil {
		return x.FlowId
	}
	return ""
}

func (x *GetTableRequest) GetArtifact() string {
	if x != nil {
		return x.Artifact
	}
	return ""
}

func (x *GetTableRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GetTableRequest) GetHuntId() string {
	if x != nil {
		return x.HuntId
	}
	return ""
}

func (x *GetTableRequest) GetNotebookId() string {
	if x != nil {
		return x.NotebookId
	}
	return ""
}

func (x *GetTableRequest) GetCellId() string {
	if x != nil {
		return x.CellId
	}
	return ""
}

func (x *GetTableRequest) GetTableId() int64 {
	if x != nil {
		return x.TableId
	}
	return 0
}

func (x *GetTableRequest) GetTimeline() string {
	if x != nil {
		return x.Timeline
	}
	return ""
}

func (x *GetTableRequest) GetSkipComponents() []string {
	if x != nil {
		return x.SkipComponents
	}
	return nil
}

func (x *GetTableRequest) GetDownloadFormat() string {
	if x != nil {
		return x.DownloadFormat
	}
	return ""
}

func (x *GetTableRequest) GetDownloadFilename() string {
	if x != nil {
		return x.DownloadFilename
	}
	return ""
}

func (x *GetTableRequest) GetColumns() []string {
	if x != nil {
		return x.Columns
	}
	return nil
}

func (x *GetTableRequest) GetSortColumn() string {
	if x != nil {
		return x.SortColumn
	}
	return ""
}

func (x *GetTableRequest) GetSortDirection() bool {
	if x != nil {
		return x.SortDirection
	}
	return false
}

func (x *GetTableRequest) GetFilterColumn() string {
	if x != nil {
		return x.FilterColumn
	}
	return ""
}

func (x *GetTableRequest) GetFilterRegex() string {
	if x != nil {
		return x.FilterRegex
	}
	return ""
}

func (x *GetTableRequest) GetStartIdx() uint64 {
	if x != nil {
		return x.StartIdx
	}
	return 0
}

func (x *GetTableRequest) GetEndIdx() uint64 {
	if x != nil {
		return x.EndIdx
	}
	return 0
}

func (x *GetTableRequest) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *GetTableRequest) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *GetTableRequest) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *GetTableRequest) GetVfsComponents() []string {
	if x != nil {
		return x.VfsComponents
	}
	return nil
}

type Row struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cell []string `protobuf:"bytes,1,rep,name=cell,proto3" json:"cell,omitempty"`
}

func (x *Row) Reset() {
	*x = Row{}
	if protoimpl.UnsafeEnabled {
		mi := &file_csv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Row) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Row) ProtoMessage() {}

func (x *Row) ProtoReflect() protoreflect.Message {
	mi := &file_csv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Row.ProtoReflect.Descriptor instead.
func (*Row) Descriptor() ([]byte, []int) {
	return file_csv_proto_rawDescGZIP(), []int{1}
}

func (x *Row) GetCell() []string {
	if x != nil {
		return x.Cell
	}
	return nil
}

type GetTableResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Columns     []string            `protobuf:"bytes,1,rep,name=columns,proto3" json:"columns,omitempty"`
	Rows        []*Row              `protobuf:"bytes,2,rep,name=rows,proto3" json:"rows,omitempty"`
	TotalRows   int64               `protobuf:"varint,3,opt,name=total_rows,json=totalRows,proto3" json:"total_rows,omitempty"`
	ColumnTypes []*proto.ColumnType `protobuf:"bytes,4,rep,name=column_types,json=columnTypes,proto3" json:"column_types,omitempty"`
	StartTime   int64               `protobuf:"varint,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime     int64               `protobuf:"varint,6,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *GetTableResponse) Reset() {
	*x = GetTableResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_csv_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTableResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTableResponse) ProtoMessage() {}

func (x *GetTableResponse) ProtoReflect() protoreflect.Message {
	mi := &file_csv_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTableResponse.ProtoReflect.Descriptor instead.
func (*GetTableResponse) Descriptor() ([]byte, []int) {
	return file_csv_proto_rawDescGZIP(), []int{2}
}

func (x *GetTableResponse) GetColumns() []string {
	if x != nil {
		return x.Columns
	}
	return nil
}

func (x *GetTableResponse) GetRows() []*Row {
	if x != nil {
		return x.Rows
	}
	return nil
}

func (x *GetTableResponse) GetTotalRows() int64 {
	if x != nil {
		return x.TotalRows
	}
	return 0
}

func (x *GetTableResponse) GetColumnTypes() []*proto.ColumnType {
	if x != nil {
		return x.ColumnTypes
	}
	return nil
}

func (x *GetTableResponse) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *GetTableResponse) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

var File_csv_proto protoreflect.FileDescriptor

var file_csv_proto_rawDesc = []byte{
	0x0a, 0x09, 0x63, 0x73, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74,
	0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x06, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x6f, 0x77, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x72, 0x6f, 0x77, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x52, 0x6f, 0x77, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x68, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x68, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x6f, 0x74, 0x65, 0x62, 0x6f,
	0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x6f, 0x74,
	0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x65, 0x6c, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x65, 0x6c, 0x6c, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74,
	0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74,
	0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x6b, 0x69, 0x70, 0x5f,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x11, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0e, 0x73, 0x6b, 0x69, 0x70, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x27, 0x0a, 0x0f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x12,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69,
	0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x73, 0x6f, 0x72, 0x74, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x21, 0x0a,
	0x0c, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18, 0x16, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x67, 0x65, 0x78,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x78, 0x18, 0x1b, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x49, 0x64, 0x78, 0x12, 0x17, 0x0a,
	0x07, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x78, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x65, 0x6e, 0x64, 0x49, 0x64, 0x78, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64,
	0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x19, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x76, 0x66, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x1a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x76, 0x66, 0x73,
	0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x19, 0x0a, 0x03, 0x52, 0x6f,
	0x77, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x65, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x65, 0x6c, 0x6c, 0x22, 0xf0, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x63, 0x6f,
	0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x13, 0xe2, 0xfc, 0xe3,
	0xc4, 0x01, 0x0d, 0x12, 0x0b, 0x54, 0x68, 0x65, 0x20, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73,
	0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x12, 0x1e, 0x0a, 0x04, 0x72, 0x6f, 0x77,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x6f, 0x77, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x52, 0x6f, 0x77, 0x73, 0x12, 0x34, 0x0a, 0x0c, 0x63, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x31, 0x5a, 0x2f, 0x77, 0x77, 0x77, 0x2e,
	0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x72, 0x61, 0x70, 0x74, 0x6f,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_csv_proto_rawDescOnce sync.Once
	file_csv_proto_rawDescData = file_csv_proto_rawDesc
)

func file_csv_proto_rawDescGZIP() []byte {
	file_csv_proto_rawDescOnce.Do(func() {
		file_csv_proto_rawDescData = protoimpl.X.CompressGZIP(file_csv_proto_rawDescData)
	})
	return file_csv_proto_rawDescData
}

var file_csv_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_csv_proto_goTypes = []interface{}{
	(*GetTableRequest)(nil),  // 0: proto.GetTableRequest
	(*Row)(nil),              // 1: proto.Row
	(*GetTableResponse)(nil), // 2: proto.GetTableResponse
	(*proto.ColumnType)(nil), // 3: proto.ColumnType
}
var file_csv_proto_depIdxs = []int32{
	1, // 0: proto.GetTableResponse.rows:type_name -> proto.Row
	3, // 1: proto.GetTableResponse.column_types:type_name -> proto.ColumnType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_csv_proto_init() }
func file_csv_proto_init() {
	if File_csv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_csv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTableRequest); i {
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
		file_csv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Row); i {
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
		file_csv_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTableResponse); i {
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
			RawDescriptor: file_csv_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_csv_proto_goTypes,
		DependencyIndexes: file_csv_proto_depIdxs,
		MessageInfos:      file_csv_proto_msgTypes,
	}.Build()
	File_csv_proto = out.File
	file_csv_proto_rawDesc = nil
	file_csv_proto_goTypes = nil
	file_csv_proto_depIdxs = nil
}
