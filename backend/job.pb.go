// Code generated by protoc-gen-gogo.
// source: job.proto
// DO NOT EDIT!

package backend

import proto "code.google.com/p/gogoprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type FrequencyMsg_DAYS int32

const (
	FrequencyMsg_NODAY FrequencyMsg_DAYS = -1
	FrequencyMsg_SUN   FrequencyMsg_DAYS = 0
	FrequencyMsg_MON   FrequencyMsg_DAYS = 1
	FrequencyMsg_TUE   FrequencyMsg_DAYS = 2
	FrequencyMsg_WED   FrequencyMsg_DAYS = 3
	FrequencyMsg_THU   FrequencyMsg_DAYS = 4
	FrequencyMsg_FRI   FrequencyMsg_DAYS = 5
	FrequencyMsg_SAT   FrequencyMsg_DAYS = 6
)

var FrequencyMsg_DAYS_name = map[int32]string{
	-1: "NODAY",
	0:  "SUN",
	1:  "MON",
	2:  "TUE",
	3:  "WED",
	4:  "THU",
	5:  "FRI",
	6:  "SAT",
}
var FrequencyMsg_DAYS_value = map[string]int32{
	"NODAY": -1,
	"SUN":   0,
	"MON":   1,
	"TUE":   2,
	"WED":   3,
	"THU":   4,
	"FRI":   5,
	"SAT":   6,
}

func (x FrequencyMsg_DAYS) Enum() *FrequencyMsg_DAYS {
	p := new(FrequencyMsg_DAYS)
	*p = x
	return p
}
func (x FrequencyMsg_DAYS) String() string {
	return proto.EnumName(FrequencyMsg_DAYS_name, int32(x))
}
func (x FrequencyMsg_DAYS) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *FrequencyMsg_DAYS) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FrequencyMsg_DAYS_value, data, "FrequencyMsg_DAYS")
	if err != nil {
		return err
	}
	*x = FrequencyMsg_DAYS(value)
	return nil
}

type JobMsg struct {
	Name             *string       `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Frequency        *FrequencyMsg `protobuf:"bytes,2,req,name=frequency" json:"frequency,omitempty"`
	ExecPath         *string       `protobuf:"bytes,3,req,name=exec_path" json:"exec_path,omitempty"`
	Args             *string       `protobuf:"bytes,4,opt,name=args" json:"args,omitempty"`
	Suspend          *bool         `protobuf:"varint,5,opt,name=suspend" json:"suspend,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *JobMsg) Reset()         { *m = JobMsg{} }
func (m *JobMsg) String() string { return proto.CompactTextString(m) }
func (*JobMsg) ProtoMessage()    {}

func (m *JobMsg) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *JobMsg) GetFrequency() *FrequencyMsg {
	if m != nil {
		return m.Frequency
	}
	return nil
}

func (m *JobMsg) GetExecPath() string {
	if m != nil && m.ExecPath != nil {
		return *m.ExecPath
	}
	return ""
}

func (m *JobMsg) GetArgs() string {
	if m != nil && m.Args != nil {
		return *m.Args
	}
	return ""
}

func (m *JobMsg) GetSuspend() bool {
	if m != nil && m.Suspend != nil {
		return *m.Suspend
	}
	return false
}

type FrequencyMsg struct {
	Month            *int32             `protobuf:"varint,1,req,name=month,def=-1" json:"month,omitempty"`
	Day              *int32             `protobuf:"varint,2,req,name=day,def=-1" json:"day,omitempty"`
	Weekday          *FrequencyMsg_DAYS `protobuf:"varint,3,req,name=weekday,enum=backend.FrequencyMsg_DAYS,def=-1" json:"weekday,omitempty"`
	Hour             *int32             `protobuf:"varint,4,req,name=hour,def=-1" json:"hour,omitempty"`
	Minute           *int32             `protobuf:"varint,5,req,name=minute,def=-1" json:"minute,omitempty"`
	Second           *int32             `protobuf:"varint,6,req,name=second,def=-1" json:"second,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *FrequencyMsg) Reset()         { *m = FrequencyMsg{} }
func (m *FrequencyMsg) String() string { return proto.CompactTextString(m) }
func (*FrequencyMsg) ProtoMessage()    {}

const Default_FrequencyMsg_Month int32 = -1
const Default_FrequencyMsg_Day int32 = -1
const Default_FrequencyMsg_Weekday FrequencyMsg_DAYS = FrequencyMsg_NODAY
const Default_FrequencyMsg_Hour int32 = -1
const Default_FrequencyMsg_Minute int32 = -1
const Default_FrequencyMsg_Second int32 = -1

func (m *FrequencyMsg) GetMonth() int32 {
	if m != nil && m.Month != nil {
		return *m.Month
	}
	return Default_FrequencyMsg_Month
}

func (m *FrequencyMsg) GetDay() int32 {
	if m != nil && m.Day != nil {
		return *m.Day
	}
	return Default_FrequencyMsg_Day
}

func (m *FrequencyMsg) GetWeekday() FrequencyMsg_DAYS {
	if m != nil && m.Weekday != nil {
		return *m.Weekday
	}
	return Default_FrequencyMsg_Weekday
}

func (m *FrequencyMsg) GetHour() int32 {
	if m != nil && m.Hour != nil {
		return *m.Hour
	}
	return Default_FrequencyMsg_Hour
}

func (m *FrequencyMsg) GetMinute() int32 {
	if m != nil && m.Minute != nil {
		return *m.Minute
	}
	return Default_FrequencyMsg_Minute
}

func (m *FrequencyMsg) GetSecond() int32 {
	if m != nil && m.Second != nil {
		return *m.Second
	}
	return Default_FrequencyMsg_Second
}

func init() {
	proto.RegisterEnum("backend.FrequencyMsg_DAYS", FrequencyMsg_DAYS_name, FrequencyMsg_DAYS_value)
}