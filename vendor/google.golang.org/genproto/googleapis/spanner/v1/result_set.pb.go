// Code generated by protoc-gen-go.
// source: google/spanner/v1/result_set.proto
// DO NOT EDIT!

package spanner

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/struct"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Results from [Read][google.spanner.v1.Spanner.Read] or
// [ExecuteSql][google.spanner.v1.Spanner.ExecuteSql].
type ResultSet struct {
	// Metadata about the result set, such as row type information.
	Metadata *ResultSetMetadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// Each element in `rows` is a row whose format is defined by
	// [metadata.row_type][google.spanner.v1.ResultSetMetadata.row_type]. The ith element
	// in each row matches the ith field in
	// [metadata.row_type][google.spanner.v1.ResultSetMetadata.row_type]. Elements are
	// encoded based on type as described
	// [here][google.spanner.v1.TypeCode].
	Rows []*google_protobuf1.ListValue `protobuf:"bytes,2,rep,name=rows" json:"rows,omitempty"`
	// Query plan and execution statistics for the query that produced this
	// result set. These can be requested by setting
	// [ExecuteSqlRequest.query_mode][google.spanner.v1.ExecuteSqlRequest.query_mode].
	Stats *ResultSetStats `protobuf:"bytes,3,opt,name=stats" json:"stats,omitempty"`
}

func (m *ResultSet) Reset()                    { *m = ResultSet{} }
func (m *ResultSet) String() string            { return proto.CompactTextString(m) }
func (*ResultSet) ProtoMessage()               {}
func (*ResultSet) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *ResultSet) GetMetadata() *ResultSetMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ResultSet) GetRows() []*google_protobuf1.ListValue {
	if m != nil {
		return m.Rows
	}
	return nil
}

func (m *ResultSet) GetStats() *ResultSetStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

// Partial results from a streaming read or SQL query. Streaming reads and
// SQL queries better tolerate large result sets, large rows, and large
// values, but are a little trickier to consume.
type PartialResultSet struct {
	// Metadata about the result set, such as row type information.
	// Only present in the first response.
	Metadata *ResultSetMetadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// A streamed result set consists of a stream of values, which might
	// be split into many `PartialResultSet` messages to accommodate
	// large rows and/or large values. Every N complete values defines a
	// row, where N is equal to the number of entries in
	// [metadata.row_type.fields][google.spanner.v1.StructType.fields].
	//
	// Most values are encoded based on type as described
	// [here][google.spanner.v1.TypeCode].
	//
	// It is possible that the last value in values is "chunked",
	// meaning that the rest of the value is sent in subsequent
	// `PartialResultSet`(s). This is denoted by the [chunked_value][google.spanner.v1.PartialResultSet.chunked_value]
	// field. Two or more chunked values can be merged to form a
	// complete value as follows:
	//
	//   * `bool/number/null`: cannot be chunked
	//   * `string`: concatenate the strings
	//   * `list`: concatenate the lists. If the last element in a list is a
	//     `string`, `list`, or `object`, merge it with the first element in
	//     the next list by applying these rules recursively.
	//   * `object`: concatenate the (field name, field value) pairs. If a
	//     field name is duplicated, then apply these rules recursively
	//     to merge the field values.
	//
	// Some examples of merging:
	//
	//     # Strings are concatenated.
	//     "foo", "bar" => "foobar"
	//
	//     # Lists of non-strings are concatenated.
	//     [2, 3], [4] => [2, 3, 4]
	//
	//     # Lists are concatenated, but the last and first elements are merged
	//     # because they are strings.
	//     ["a", "b"], ["c", "d"] => ["a", "bc", "d"]
	//
	//     # Lists are concatenated, but the last and first elements are merged
	//     # because they are lists. Recursively, the last and first elements
	//     # of the inner lists are merged because they are strings.
	//     ["a", ["b", "c"]], [["d"], "e"] => ["a", ["b", "cd"], "e"]
	//
	//     # Non-overlapping object fields are combined.
	//     {"a": "1"}, {"b": "2"} => {"a": "1", "b": 2"}
	//
	//     # Overlapping object fields are merged.
	//     {"a": "1"}, {"a": "2"} => {"a": "12"}
	//
	//     # Examples of merging objects containing lists of strings.
	//     {"a": ["1"]}, {"a": ["2"]} => {"a": ["12"]}
	//
	// For a more complete example, suppose a streaming SQL query is
	// yielding a result set whose rows contain a single string
	// field. The following `PartialResultSet`s might be yielded:
	//
	//     {
	//       "metadata": { ... }
	//       "values": ["Hello", "W"]
	//       "chunked_value": true
	//       "resume_token": "Af65..."
	//     }
	//     {
	//       "values": ["orl"]
	//       "chunked_value": true
	//       "resume_token": "Bqp2..."
	//     }
	//     {
	//       "values": ["d"]
	//       "resume_token": "Zx1B..."
	//     }
	//
	// This sequence of `PartialResultSet`s encodes two rows, one
	// containing the field value `"Hello"`, and a second containing the
	// field value `"World" = "W" + "orl" + "d"`.
	Values []*google_protobuf1.Value `protobuf:"bytes,2,rep,name=values" json:"values,omitempty"`
	// If true, then the final value in [values][google.spanner.v1.PartialResultSet.values] is chunked, and must
	// be combined with more values from subsequent `PartialResultSet`s
	// to obtain a complete field value.
	ChunkedValue bool `protobuf:"varint,3,opt,name=chunked_value,json=chunkedValue" json:"chunked_value,omitempty"`
	// Streaming calls might be interrupted for a variety of reasons, such
	// as TCP connection loss. If this occurs, the stream of results can
	// be resumed by re-sending the original request and including
	// `resume_token`. Note that executing any other transaction in the
	// same session invalidates the token.
	ResumeToken []byte `protobuf:"bytes,4,opt,name=resume_token,json=resumeToken,proto3" json:"resume_token,omitempty"`
	// Query plan and execution statistics for the query that produced this
	// streaming result set. These can be requested by setting
	// [ExecuteSqlRequest.query_mode][google.spanner.v1.ExecuteSqlRequest.query_mode] and are sent
	// only once with the last response in the stream.
	Stats *ResultSetStats `protobuf:"bytes,5,opt,name=stats" json:"stats,omitempty"`
}

func (m *PartialResultSet) Reset()                    { *m = PartialResultSet{} }
func (m *PartialResultSet) String() string            { return proto.CompactTextString(m) }
func (*PartialResultSet) ProtoMessage()               {}
func (*PartialResultSet) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *PartialResultSet) GetMetadata() *ResultSetMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *PartialResultSet) GetValues() []*google_protobuf1.Value {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *PartialResultSet) GetChunkedValue() bool {
	if m != nil {
		return m.ChunkedValue
	}
	return false
}

func (m *PartialResultSet) GetResumeToken() []byte {
	if m != nil {
		return m.ResumeToken
	}
	return nil
}

func (m *PartialResultSet) GetStats() *ResultSetStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

// Metadata about a [ResultSet][google.spanner.v1.ResultSet] or [PartialResultSet][google.spanner.v1.PartialResultSet].
type ResultSetMetadata struct {
	// Indicates the field names and types for the rows in the result
	// set.  For example, a SQL query like `"SELECT UserId, UserName FROM
	// Users"` could return a `row_type` value like:
	//
	//     "fields": [
	//       { "name": "UserId", "type": { "code": "INT64" } },
	//       { "name": "UserName", "type": { "code": "STRING" } },
	//     ]
	RowType *StructType `protobuf:"bytes,1,opt,name=row_type,json=rowType" json:"row_type,omitempty"`
	// If the read or SQL query began a transaction as a side-effect, the
	// information about the new transaction is yielded here.
	Transaction *Transaction `protobuf:"bytes,2,opt,name=transaction" json:"transaction,omitempty"`
}

func (m *ResultSetMetadata) Reset()                    { *m = ResultSetMetadata{} }
func (m *ResultSetMetadata) String() string            { return proto.CompactTextString(m) }
func (*ResultSetMetadata) ProtoMessage()               {}
func (*ResultSetMetadata) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *ResultSetMetadata) GetRowType() *StructType {
	if m != nil {
		return m.RowType
	}
	return nil
}

func (m *ResultSetMetadata) GetTransaction() *Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

// Additional statistics about a [ResultSet][google.spanner.v1.ResultSet] or [PartialResultSet][google.spanner.v1.PartialResultSet].
type ResultSetStats struct {
	// [QueryPlan][google.spanner.v1.QueryPlan] for the query associated with this result.
	QueryPlan *QueryPlan `protobuf:"bytes,1,opt,name=query_plan,json=queryPlan" json:"query_plan,omitempty"`
	// Aggregated statistics from the execution of the query. Only present when
	// the query is profiled. For example, a query could return the statistics as
	// follows:
	//
	//     {
	//       "rows_returned": "3",
	//       "elapsed_time": "1.22 secs",
	//       "cpu_time": "1.19 secs"
	//     }
	QueryStats *google_protobuf1.Struct `protobuf:"bytes,2,opt,name=query_stats,json=queryStats" json:"query_stats,omitempty"`
}

func (m *ResultSetStats) Reset()                    { *m = ResultSetStats{} }
func (m *ResultSetStats) String() string            { return proto.CompactTextString(m) }
func (*ResultSetStats) ProtoMessage()               {}
func (*ResultSetStats) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *ResultSetStats) GetQueryPlan() *QueryPlan {
	if m != nil {
		return m.QueryPlan
	}
	return nil
}

func (m *ResultSetStats) GetQueryStats() *google_protobuf1.Struct {
	if m != nil {
		return m.QueryStats
	}
	return nil
}

func init() {
	proto.RegisterType((*ResultSet)(nil), "google.spanner.v1.ResultSet")
	proto.RegisterType((*PartialResultSet)(nil), "google.spanner.v1.PartialResultSet")
	proto.RegisterType((*ResultSetMetadata)(nil), "google.spanner.v1.ResultSetMetadata")
	proto.RegisterType((*ResultSetStats)(nil), "google.spanner.v1.ResultSetStats")
}

func init() { proto.RegisterFile("google/spanner/v1/result_set.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcb, 0x6e, 0x13, 0x31,
	0x14, 0x86, 0x35, 0xe9, 0x85, 0xd4, 0x13, 0x10, 0xb5, 0x04, 0x1d, 0x45, 0x05, 0xa5, 0x29, 0x8b,
	0xac, 0x3c, 0x4a, 0x59, 0x10, 0xa9, 0x9b, 0xaa, 0x2c, 0xd8, 0x80, 0x14, 0x9c, 0xa8, 0x0b, 0x36,
	0xa3, 0xd3, 0xc4, 0x0c, 0xa3, 0x3a, 0xf6, 0xd4, 0xf6, 0x24, 0xca, 0x82, 0x25, 0x62, 0xc9, 0x7b,
	0xf0, 0x00, 0x3c, 0x1f, 0xf2, 0x25, 0x17, 0x98, 0x08, 0x09, 0xa9, 0x3b, 0xc7, 0xfe, 0xfe, 0xf3,
	0x9f, 0xff, 0xcc, 0x09, 0xea, 0xe6, 0x52, 0xe6, 0x9c, 0xa5, 0xba, 0x04, 0x21, 0x98, 0x4a, 0xe7,
	0xfd, 0x54, 0x31, 0x5d, 0x71, 0x93, 0x69, 0x66, 0x48, 0xa9, 0xa4, 0x91, 0xf8, 0xd8, 0x33, 0x24,
	0x30, 0x64, 0xde, 0x6f, 0x9f, 0x06, 0x19, 0x94, 0x45, 0x0a, 0x42, 0x48, 0x03, 0xa6, 0x90, 0x42,
	0x7b, 0xc1, 0xfa, 0xd5, 0xfd, 0xba, 0xad, 0x3e, 0xa7, 0xda, 0xa8, 0x6a, 0x12, 0xca, 0xb5, 0x77,
	0x58, 0xde, 0x57, 0x4c, 0x2d, 0xb3, 0x92, 0x83, 0x08, 0xcc, 0x79, 0x9d, 0x31, 0x0a, 0x84, 0x86,
	0x89, 0xf5, 0xf9, 0xcb, 0x66, 0x1b, 0x5a, 0x96, 0xcc, 0xbf, 0x76, 0x7f, 0x45, 0xe8, 0x88, 0xba,
	0x28, 0x23, 0x66, 0xf0, 0x15, 0x6a, 0xce, 0x98, 0x81, 0x29, 0x18, 0x48, 0xa2, 0x4e, 0xd4, 0x8b,
	0x2f, 0x5e, 0x91, 0x5a, 0x2c, 0xb2, 0xe6, 0x3f, 0x04, 0x96, 0xae, 0x55, 0x98, 0xa0, 0x7d, 0x25,
	0x17, 0x3a, 0x69, 0x74, 0xf6, 0x7a, 0xf1, 0x45, 0x7b, 0xa5, 0x5e, 0x65, 0x24, 0xef, 0x0b, 0x6d,
	0x6e, 0x80, 0x57, 0x8c, 0x3a, 0x0e, 0xbf, 0x41, 0x07, 0xda, 0x80, 0xd1, 0xc9, 0x9e, 0xb3, 0x3b,
	0xfb, 0x97, 0xdd, 0xc8, 0x82, 0xd4, 0xf3, 0xdd, 0x6f, 0x0d, 0xf4, 0x74, 0x08, 0xca, 0x14, 0xc0,
	0x1f, 0xb6, 0xff, 0xc3, 0xb9, 0x6d, 0x6f, 0x95, 0xe0, 0x79, 0x2d, 0x81, 0xef, 0x3e, 0x50, 0xf8,
	0x1c, 0x3d, 0x9e, 0x7c, 0xa9, 0xc4, 0x1d, 0x9b, 0x66, 0xee, 0xc6, 0xe5, 0x68, 0xd2, 0x56, 0xb8,
	0x74, 0x30, 0x3e, 0x43, 0x2d, 0xbb, 0x2e, 0x33, 0x96, 0x19, 0x79, 0xc7, 0x44, 0xb2, 0xdf, 0x89,
	0x7a, 0x2d, 0x1a, 0xfb, 0xbb, 0xb1, 0xbd, 0xda, 0xcc, 0xe1, 0xe0, 0x3f, 0xe7, 0xf0, 0x23, 0x42,
	0xc7, 0xb5, 0x40, 0x78, 0x80, 0x9a, 0x4a, 0x2e, 0x32, 0xfb, 0xa1, 0xc3, 0x20, 0x5e, 0xec, 0xa8,
	0x38, 0x72, 0x0b, 0x37, 0x5e, 0x96, 0x8c, 0x3e, 0x52, 0x72, 0x61, 0x0f, 0xf8, 0x0a, 0xc5, 0x5b,
	0x3b, 0x94, 0x34, 0x9c, 0xf8, 0xe5, 0x0e, 0xf1, 0x78, 0x43, 0xd1, 0x6d, 0x49, 0xf7, 0x7b, 0x84,
	0x9e, 0xfc, 0xd9, 0x2b, 0xbe, 0x44, 0x68, 0xb3, 0xbc, 0xa1, 0xa1, 0xd3, 0x1d, 0x35, 0x3f, 0x5a,
	0x68, 0xc8, 0x41, 0xd0, 0xa3, 0xfb, 0xd5, 0x11, 0x0f, 0x50, 0xec, 0xc5, 0x7e, 0x40, 0xbe, 0xa3,
	0x93, 0xda, 0x77, 0xf1, 0x61, 0xa8, 0x37, 0x72, 0xb6, 0xd7, 0x5f, 0xd1, 0xb3, 0x89, 0x9c, 0xd5,
	0x7d, 0xae, 0x37, 0xfd, 0x0d, 0xad, 0x7c, 0x18, 0x7d, 0x1a, 0x04, 0x28, 0x97, 0x1c, 0x44, 0x4e,
	0xa4, 0xca, 0xd3, 0x9c, 0x09, 0x57, 0x3c, 0xf5, 0x4f, 0x50, 0x16, 0x7a, 0xeb, 0x4f, 0x74, 0x19,
	0x8e, 0x3f, 0x1b, 0x27, 0xef, 0xbc, 0xf4, 0x2d, 0x97, 0xd5, 0x94, 0x8c, 0x82, 0xcb, 0x4d, 0xff,
	0xf6, 0xd0, 0xc9, 0x5f, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x15, 0x1f, 0xa6, 0x3e, 0x04,
	0x00, 0x00,
}
