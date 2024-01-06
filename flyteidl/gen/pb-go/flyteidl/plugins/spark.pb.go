// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/plugins/spark.proto

package plugins

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type SparkApplication_Type int32

const (
	SparkApplication_PYTHON SparkApplication_Type = 0
	SparkApplication_JAVA   SparkApplication_Type = 1
	SparkApplication_SCALA  SparkApplication_Type = 2
	SparkApplication_R      SparkApplication_Type = 3
)

var SparkApplication_Type_name = map[int32]string{
	0: "PYTHON",
	1: "JAVA",
	2: "SCALA",
	3: "R",
}

var SparkApplication_Type_value = map[string]int32{
	"PYTHON": 0,
	"JAVA":   1,
	"SCALA":  2,
	"R":      3,
}

func (x SparkApplication_Type) String() string {
	return proto.EnumName(SparkApplication_Type_name, int32(x))
}

func (SparkApplication_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ca8a069b9820144a, []int{0, 0}
}

type SparkApplication struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SparkApplication) Reset()         { *m = SparkApplication{} }
func (m *SparkApplication) String() string { return proto.CompactTextString(m) }
func (*SparkApplication) ProtoMessage()    {}
func (*SparkApplication) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca8a069b9820144a, []int{0}
}

func (m *SparkApplication) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SparkApplication.Unmarshal(m, b)
}
func (m *SparkApplication) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SparkApplication.Marshal(b, m, deterministic)
}
func (m *SparkApplication) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SparkApplication.Merge(m, src)
}
func (m *SparkApplication) XXX_Size() int {
	return xxx_messageInfo_SparkApplication.Size(m)
}
func (m *SparkApplication) XXX_DiscardUnknown() {
	xxx_messageInfo_SparkApplication.DiscardUnknown(m)
}

var xxx_messageInfo_SparkApplication proto.InternalMessageInfo

// Custom Proto for Spark Plugin.
type SparkJob struct {
	ApplicationType     SparkApplication_Type `protobuf:"varint,1,opt,name=applicationType,proto3,enum=flyteidl.plugins.SparkApplication_Type" json:"applicationType,omitempty"`
	MainApplicationFile string                `protobuf:"bytes,2,opt,name=mainApplicationFile,proto3" json:"mainApplicationFile,omitempty"`
	MainClass           string                `protobuf:"bytes,3,opt,name=mainClass,proto3" json:"mainClass,omitempty"`
	SparkConf           map[string]string     `protobuf:"bytes,4,rep,name=sparkConf,proto3" json:"sparkConf,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	HadoopConf          map[string]string     `protobuf:"bytes,5,rep,name=hadoopConf,proto3" json:"hadoopConf,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ExecutorPath        string                `protobuf:"bytes,6,opt,name=executorPath,proto3" json:"executorPath,omitempty"`
	// Databricks job configuration.
	// Config structure can be found here. https://docs.databricks.com/dev-tools/api/2.0/jobs.html#request-structure.
	DatabricksConf *_struct.Struct `protobuf:"bytes,7,opt,name=databricksConf,proto3" json:"databricksConf,omitempty"`
	// Databricks access token. https://docs.databricks.com/dev-tools/api/latest/authentication.html
	// This token can be set in either flytepropeller or flytekit.
	DatabricksToken string `protobuf:"bytes,8,opt,name=databricksToken,proto3" json:"databricksToken,omitempty"`
	// Domain name of your deployment. Use the form <account>.cloud.databricks.com.
	// This instance name can be set in either flytepropeller or flytekit.
	DatabricksInstance   string   `protobuf:"bytes,9,opt,name=databricksInstance,proto3" json:"databricksInstance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SparkJob) Reset()         { *m = SparkJob{} }
func (m *SparkJob) String() string { return proto.CompactTextString(m) }
func (*SparkJob) ProtoMessage()    {}
func (*SparkJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca8a069b9820144a, []int{1}
}

func (m *SparkJob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SparkJob.Unmarshal(m, b)
}
func (m *SparkJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SparkJob.Marshal(b, m, deterministic)
}
func (m *SparkJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SparkJob.Merge(m, src)
}
func (m *SparkJob) XXX_Size() int {
	return xxx_messageInfo_SparkJob.Size(m)
}
func (m *SparkJob) XXX_DiscardUnknown() {
	xxx_messageInfo_SparkJob.DiscardUnknown(m)
}

var xxx_messageInfo_SparkJob proto.InternalMessageInfo

func (m *SparkJob) GetApplicationType() SparkApplication_Type {
	if m != nil {
		return m.ApplicationType
	}
	return SparkApplication_PYTHON
}

func (m *SparkJob) GetMainApplicationFile() string {
	if m != nil {
		return m.MainApplicationFile
	}
	return ""
}

func (m *SparkJob) GetMainClass() string {
	if m != nil {
		return m.MainClass
	}
	return ""
}

func (m *SparkJob) GetSparkConf() map[string]string {
	if m != nil {
		return m.SparkConf
	}
	return nil
}

func (m *SparkJob) GetHadoopConf() map[string]string {
	if m != nil {
		return m.HadoopConf
	}
	return nil
}

func (m *SparkJob) GetExecutorPath() string {
	if m != nil {
		return m.ExecutorPath
	}
	return ""
}

func (m *SparkJob) GetDatabricksConf() *_struct.Struct {
	if m != nil {
		return m.DatabricksConf
	}
	return nil
}

func (m *SparkJob) GetDatabricksToken() string {
	if m != nil {
		return m.DatabricksToken
	}
	return ""
}

func (m *SparkJob) GetDatabricksInstance() string {
	if m != nil {
		return m.DatabricksInstance
	}
	return ""
}

func init() {
	proto.RegisterEnum("flyteidl.plugins.SparkApplication_Type", SparkApplication_Type_name, SparkApplication_Type_value)
	proto.RegisterType((*SparkApplication)(nil), "flyteidl.plugins.SparkApplication")
	proto.RegisterType((*SparkJob)(nil), "flyteidl.plugins.SparkJob")
	proto.RegisterMapType((map[string]string)(nil), "flyteidl.plugins.SparkJob.HadoopConfEntry")
	proto.RegisterMapType((map[string]string)(nil), "flyteidl.plugins.SparkJob.SparkConfEntry")
}

func init() { proto.RegisterFile("flyteidl/plugins/spark.proto", fileDescriptor_ca8a069b9820144a) }

var fileDescriptor_ca8a069b9820144a = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xdf, 0x8b, 0xd3, 0x40,
	0x10, 0x36, 0xfd, 0x65, 0x33, 0x27, 0x6d, 0x18, 0x05, 0x43, 0xe9, 0x43, 0xe9, 0x8b, 0x51, 0x70,
	0x23, 0xf5, 0x45, 0xc4, 0xe3, 0xc8, 0x15, 0xf5, 0x2c, 0xa2, 0x67, 0x5a, 0x04, 0x7d, 0xdb, 0xa4,
	0xdb, 0x34, 0x74, 0x6f, 0x77, 0x49, 0x36, 0x62, 0xfe, 0x79, 0x91, 0x6c, 0xec, 0xe5, 0x2e, 0x9c,
	0x82, 0x6f, 0xb3, 0xdf, 0x7c, 0xf3, 0xcd, 0xf0, 0x7d, 0x2c, 0x4c, 0x77, 0xbc, 0xd4, 0x2c, 0xdd,
	0x72, 0x5f, 0xf1, 0x22, 0x49, 0x45, 0xee, 0xe7, 0x8a, 0x66, 0x07, 0xa2, 0x32, 0xa9, 0x25, 0x3a,
	0xc7, 0x2e, 0xf9, 0xd3, 0x9d, 0x4c, 0x13, 0x29, 0x13, 0xce, 0x7c, 0xd3, 0x8f, 0x8a, 0x9d, 0x9f,
	0xeb, 0xac, 0x88, 0x75, 0xcd, 0x9f, 0x9f, 0x83, 0xb3, 0xae, 0xc6, 0x03, 0xa5, 0x78, 0x1a, 0x53,
	0x9d, 0x4a, 0x31, 0x27, 0xd0, 0xdb, 0x94, 0x8a, 0x21, 0xc0, 0xe0, 0xf2, 0xdb, 0xe6, 0xe2, 0xf3,
	0x27, 0xe7, 0x1e, 0x0e, 0xa1, 0xb7, 0x0a, 0xbe, 0x06, 0x8e, 0x85, 0x36, 0xf4, 0xd7, 0xcb, 0xe0,
	0x63, 0xe0, 0x74, 0xb0, 0x0f, 0x56, 0xe8, 0x74, 0xe7, 0xbf, 0x7a, 0x30, 0x34, 0x22, 0x2b, 0x19,
	0xe1, 0x17, 0x18, 0xd3, 0x46, 0xab, 0xd2, 0x71, 0xad, 0x99, 0xe5, 0x8d, 0x16, 0x4f, 0x48, 0xfb,
	0x34, 0xd2, 0xde, 0x4c, 0x2a, 0x7a, 0xd8, 0x9e, 0xc7, 0x17, 0xf0, 0xf0, 0x8a, 0xa6, 0xe2, 0x06,
	0xf1, 0x5d, 0xca, 0x99, 0xdb, 0x99, 0x59, 0x9e, 0x1d, 0xde, 0xd5, 0xc2, 0x29, 0xd8, 0x15, 0xbc,
	0xe4, 0x34, 0xcf, 0xdd, 0xae, 0xe1, 0x35, 0x00, 0xbe, 0x07, 0xdb, 0x58, 0xb6, 0x94, 0x62, 0xe7,
	0xf6, 0x66, 0x5d, 0xef, 0x64, 0xf1, 0xf4, 0x2f, 0xc7, 0xad, 0x64, 0x54, 0x17, 0x15, 0xf7, 0xad,
	0xd0, 0x59, 0x19, 0x36, 0xb3, 0xb8, 0x02, 0xd8, 0xd3, 0xad, 0x94, 0xca, 0x28, 0xf5, 0x8d, 0xd2,
	0xb3, 0x7f, 0x28, 0x5d, 0x5c, 0x93, 0x6b, 0xa9, 0x1b, 0xd3, 0x38, 0x87, 0x07, 0xec, 0x27, 0x8b,
	0x0b, 0x2d, 0xb3, 0x4b, 0xaa, 0xf7, 0xee, 0xc0, 0x5c, 0x7d, 0x0b, 0xc3, 0x33, 0x18, 0x6d, 0xa9,
	0xa6, 0x51, 0x96, 0xc6, 0x87, 0xdc, 0xec, 0xbc, 0x3f, 0xb3, 0xbc, 0x93, 0xc5, 0x63, 0x52, 0x67,
	0x4c, 0x8e, 0x19, 0x93, 0xb5, 0xc9, 0x38, 0x6c, 0xd1, 0xd1, 0x83, 0x71, 0x83, 0x6c, 0xe4, 0x81,
	0x09, 0x77, 0x68, 0xf6, 0xb4, 0x61, 0x24, 0x80, 0x0d, 0xf4, 0x41, 0xe4, 0x9a, 0x8a, 0x98, 0xb9,
	0xb6, 0x21, 0xdf, 0xd1, 0x99, 0xbc, 0x81, 0xd1, 0x6d, 0x9f, 0xd0, 0x81, 0xee, 0x81, 0x95, 0x26,
	0x7c, 0x3b, 0xac, 0x4a, 0x7c, 0x04, 0xfd, 0x1f, 0x94, 0x17, 0xc7, 0xe4, 0xea, 0xc7, 0xeb, 0xce,
	0x2b, 0x6b, 0x72, 0x0a, 0xe3, 0x96, 0x37, 0xff, 0x33, 0x7e, 0x7e, 0xf6, 0xfd, 0x34, 0x49, 0xf5,
	0xbe, 0x88, 0x48, 0x2c, 0xaf, 0x7c, 0xe3, 0xbf, 0xcc, 0x92, 0xba, 0xf0, 0xaf, 0xbf, 0x4b, 0xc2,
	0x84, 0xaf, 0xa2, 0xe7, 0x89, 0xf4, 0xdb, 0x3f, 0x28, 0x1a, 0x18, 0xe3, 0x5e, 0xfe, 0x0e, 0x00,
	0x00, 0xff, 0xff, 0x07, 0x7f, 0xfa, 0xf2, 0x5c, 0x03, 0x00, 0x00,
}
