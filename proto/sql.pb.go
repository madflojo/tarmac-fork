// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v5.27.1
// source: sql.proto

package proto

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// SQLQuery is a structure used to create SQL queries to a SQL Database.
type SQLQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Query is the SQL Query to be executed. This field should be a byte slice
	// to avoid conflicts with JSON encoding.
	Query []byte `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *SQLQuery) Reset() {
	*x = SQLQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sql_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SQLQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SQLQuery) ProtoMessage() {}

func (x *SQLQuery) ProtoReflect() protoreflect.Message {
	mi := &file_sql_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SQLQuery.ProtoReflect.Descriptor instead.
func (*SQLQuery) Descriptor() ([]byte, []int) {
	return file_sql_proto_rawDescGZIP(), []int{0}
}

func (x *SQLQuery) GetQuery() []byte {
	if x != nil {
		return x.Query
	}
	return nil
}

// SQLQueryResponse is a structure supplied as a response message to a SQL
// Database Query.
type SQLQueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Status is the human readable error message or success message for function
	// execution.
	Status *Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// LastInsertID is the ID of the last inserted row. This field is only
	// populated if the query was an insert query and the database supports
	// returning the last inserted ID.
	LastInsertId int64 `protobuf:"varint,2,opt,name=last_insert_id,json=lastInsertId,proto3" json:"last_insert_id,omitempty"`
	// RowsAffected is the number of rows affected by the query. This field is
	// only populated if the query was an insert, update, or delete query.
	RowsAffected int64 `protobuf:"varint,3,opt,name=rows_affected,json=rowsAffected,proto3" json:"rows_affected,omitempty"`
	// Columns is a list of column names returned by the query. This field is
	// only populated if the query was a select query.
	Columns []string `protobuf:"bytes,4,rep,name=columns,proto3" json:"columns,omitempty"`
	// Rows is a list of rows returned by the query. This field is only populated
	// if the query was a select query.
	Rows []*Row `protobuf:"bytes,5,rep,name=rows,proto3" json:"rows,omitempty"`
}

func (x *SQLQueryResponse) Reset() {
	*x = SQLQueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sql_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SQLQueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SQLQueryResponse) ProtoMessage() {}

func (x *SQLQueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sql_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SQLQueryResponse.ProtoReflect.Descriptor instead.
func (*SQLQueryResponse) Descriptor() ([]byte, []int) {
	return file_sql_proto_rawDescGZIP(), []int{1}
}

func (x *SQLQueryResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *SQLQueryResponse) GetLastInsertId() int64 {
	if x != nil {
		return x.LastInsertId
	}
	return 0
}

func (x *SQLQueryResponse) GetRowsAffected() int64 {
	if x != nil {
		return x.RowsAffected
	}
	return 0
}

func (x *SQLQueryResponse) GetColumns() []string {
	if x != nil {
		return x.Columns
	}
	return nil
}

func (x *SQLQueryResponse) GetRows() []*Row {
	if x != nil {
		return x.Rows
	}
	return nil
}

// Row is a structure used to represent a row in a SQL query response.
// The data field is a map of column names to column values.
type Row struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data map[string][]byte `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Row) Reset() {
	*x = Row{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sql_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Row) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Row) ProtoMessage() {}

func (x *Row) ProtoReflect() protoreflect.Message {
	mi := &file_sql_proto_msgTypes[2]
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
	return file_sql_proto_rawDescGZIP(), []int{2}
}

func (x *Row) GetData() map[string][]byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_sql_proto protoreflect.FileDescriptor

var file_sql_proto_rawDesc = []byte{
	0x0a, 0x09, 0x73, 0x71, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x61, 0x72,
	0x6d, 0x61, 0x63, 0x2e, 0x73, 0x71, 0x6c, 0x1a, 0x0c, 0x74, 0x61, 0x72, 0x6d, 0x61, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x08, 0x53, 0x51, 0x4c, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0xc4, 0x01, 0x0a, 0x10, 0x53, 0x51, 0x4c, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74,
	0x61, 0x72, 0x6d, 0x61, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x73,
	0x65, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6c, 0x61,
	0x73, 0x74, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x6f,
	0x77, 0x73, 0x5f, 0x61, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x72, 0x6f, 0x77, 0x73, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x12, 0x23, 0x0a, 0x04, 0x72, 0x6f, 0x77,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x61, 0x72, 0x6d, 0x61, 0x63,
	0x2e, 0x73, 0x71, 0x6c, 0x2e, 0x52, 0x6f, 0x77, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x22, 0x6d,
	0x0a, 0x03, 0x52, 0x6f, 0x77, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x61, 0x72, 0x6d, 0x61, 0x63, 0x2e, 0x73, 0x71, 0x6c,
	0x2e, 0x52, 0x6f, 0x77, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x28, 0x5a,
	0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x61, 0x72, 0x6d,
	0x61, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x74, 0x61, 0x72, 0x6d, 0x61,
	0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sql_proto_rawDescOnce sync.Once
	file_sql_proto_rawDescData = file_sql_proto_rawDesc
)

func file_sql_proto_rawDescGZIP() []byte {
	file_sql_proto_rawDescOnce.Do(func() {
		file_sql_proto_rawDescData = protoimpl.X.CompressGZIP(file_sql_proto_rawDescData)
	})
	return file_sql_proto_rawDescData
}

var file_sql_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_sql_proto_goTypes = []interface{}{
	(*SQLQuery)(nil),         // 0: tarmac.sql.SQLQuery
	(*SQLQueryResponse)(nil), // 1: tarmac.sql.SQLQueryResponse
	(*Row)(nil),              // 2: tarmac.sql.Row
	nil,                      // 3: tarmac.sql.Row.DataEntry
	(*Status)(nil),           // 4: tarmac.Status
}
var file_sql_proto_depIdxs = []int32{
	4, // 0: tarmac.sql.SQLQueryResponse.status:type_name -> tarmac.Status
	2, // 1: tarmac.sql.SQLQueryResponse.rows:type_name -> tarmac.sql.Row
	3, // 2: tarmac.sql.Row.data:type_name -> tarmac.sql.Row.DataEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_sql_proto_init() }
func file_sql_proto_init() {
	if File_sql_proto != nil {
		return
	}
	file_tarmac_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sql_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SQLQuery); i {
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
		file_sql_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SQLQueryResponse); i {
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
		file_sql_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sql_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sql_proto_goTypes,
		DependencyIndexes: file_sql_proto_depIdxs,
		MessageInfos:      file_sql_proto_msgTypes,
	}.Build()
	File_sql_proto = out.File
	file_sql_proto_rawDesc = nil
	file_sql_proto_goTypes = nil
	file_sql_proto_depIdxs = nil
}