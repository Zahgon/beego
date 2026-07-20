package alils

import (
	"fmt"
	"math"

	github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
)

var _ = github_com_gogo_protobuf_proto.Marshal

var _ = fmt.Errorf

var _ = math.Inf

var (
	ErrInvalidLengthLog = fmt.Errorf("proto: negative length found during unmarshaling")

	ErrIntOverflowLog = fmt.Errorf("proto: integer overflow")
)

type Log struct {
	Time            *uint32       `protobuf:"varint,1,req,name=Time" json:"Time,omitempty"`
	Contents        []*LogContent `protobuf:"bytes,2,rep,name=Contents" json:"Contents,omitempty"`
	XXXUnrecognized []byte        `json:"-"`
}

func (m *Log) Reset() { _ = "STUB: not implemented"; return }

func (m *Log) String() string { _ = "STUB: not implemented"; return "" }

func (*Log) ProtoMessage() { _ = "STUB: not implemented"; return }

func (m *Log) GetTime() uint32 { _ = "STUB: not implemented"; return 0 }

func (m *Log) GetContents() []*LogContent { _ = "STUB: not implemented"; return nil }

type LogContent struct {
	Key             *string `protobuf:"bytes,1,req,name=Key" json:"Key,omitempty"`
	Value           *string `protobuf:"bytes,2,req,name=Value" json:"Value,omitempty"`
	XXXUnrecognized []byte  `json:"-"`
}

func (m *LogContent) Reset() { _ = "STUB: not implemented"; return }

func (m *LogContent) String() string { _ = "STUB: not implemented"; return "" }

func (*LogContent) ProtoMessage() { _ = "STUB: not implemented"; return }

func (m *LogContent) GetKey() string { _ = "STUB: not implemented"; return "" }

func (m *LogContent) GetValue() string { _ = "STUB: not implemented"; return "" }

type LogGroup struct {
	Logs            []*Log  `protobuf:"bytes,1,rep,name=Logs" json:"Logs,omitempty"`
	Reserved        *string `protobuf:"bytes,2,opt,name=Reserved" json:"Reserved,omitempty"`
	Topic           *string `protobuf:"bytes,3,opt,name=Topic" json:"Topic,omitempty"`
	Source          *string `protobuf:"bytes,4,opt,name=Source" json:"Source,omitempty"`
	XXXUnrecognized []byte  `json:"-"`
}

func (m *LogGroup) Reset() { _ = "STUB: not implemented"; return }

func (m *LogGroup) String() string { _ = "STUB: not implemented"; return "" }

func (*LogGroup) ProtoMessage() { _ = "STUB: not implemented"; return }

func (m *LogGroup) GetLogs() []*Log { _ = "STUB: not implemented"; return nil }

func (m *LogGroup) GetReserved() string { _ = "STUB: not implemented"; return "" }

func (m *LogGroup) GetTopic() string { _ = "STUB: not implemented"; return "" }

func (m *LogGroup) GetSource() string { _ = "STUB: not implemented"; return "" }

type LogGroupList struct {
	LogGroups       []*LogGroup `protobuf:"bytes,1,rep,name=logGroups" json:"logGroups,omitempty"`
	XXXUnrecognized []byte      `json:"-"`
}

func (m *LogGroupList) Reset() { _ = "STUB: not implemented"; return }

func (m *LogGroupList) String() string { _ = "STUB: not implemented"; return "" }

func (*LogGroupList) ProtoMessage() { _ = "STUB: not implemented"; return }

func (m *LogGroupList) GetLogGroups() []*LogGroup { _ = "STUB: not implemented"; return nil }

func (m *Log) Marshal() (data []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (m *Log) MarshalTo(data []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (m *LogContent) Marshal() (data []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (m *LogContent) MarshalTo(data []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (m *LogGroup) Marshal() (data []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (m *LogGroup) MarshalTo(data []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (m *LogGroupList) Marshal() (data []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *LogGroupList) MarshalTo(data []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func encodeFixed64Log(data []byte, offset int, v uint64) int { _ = "STUB: not implemented"; return 0 }

func encodeFixed32Log(data []byte, offset int, v uint32) int { _ = "STUB: not implemented"; return 0 }

func encodeVarintLog(data []byte, offset int, v uint64) int { _ = "STUB: not implemented"; return 0 }

func (m *Log) Size() (n int) { _ = "STUB: not implemented"; return 0 }

func (m *LogContent) Size() (n int) { _ = "STUB: not implemented"; return 0 }

func (m *LogGroup) Size() (n int) { _ = "STUB: not implemented"; return 0 }

func (m *LogGroupList) Size() (n int) { _ = "STUB: not implemented"; return 0 }

func sovLog(x uint64) (n int) { _ = "STUB: not implemented"; return 0 }

func sozLog(x uint64) (n int) { _ = "STUB: not implemented"; return 0 }

func (m *Log) Unmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }

func (m *LogContent) Unmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }

func (m *LogGroup) Unmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }

func (m *LogGroupList) Unmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }

func skipLog(data []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }
