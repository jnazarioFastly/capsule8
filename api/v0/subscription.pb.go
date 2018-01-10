// Code generated by protoc-gen-go. DO NOT EDIT.
// source: capsule8/api/v0/subscription.proto

package capsule8_api_v0

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The ContainerEventView specifies the level of detail to include for
// ContainerEvents.
type ContainerEventView int32

const (
	// Default view of a ContainerEvent includes just basic information
	ContainerEventView_BASIC ContainerEventView = 0
	// Full view of a ContainerEvent includes raw Docker and OCI config JSON
	// payloads
	ContainerEventView_FULL ContainerEventView = 1
)

var ContainerEventView_name = map[int32]string{
	0: "BASIC",
	1: "FULL",
}
var ContainerEventView_value = map[string]int32{
	"BASIC": 0,
	"FULL":  1,
}

func (x ContainerEventView) String() string {
	return proto.EnumName(ContainerEventView_name, int32(x))
}
func (ContainerEventView) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

// Possible interval types
type ThrottleModifier_IntervalType int32

const (
	// milliseconds
	ThrottleModifier_MILLISECOND ThrottleModifier_IntervalType = 0
	// seconds
	ThrottleModifier_SECOND ThrottleModifier_IntervalType = 1
	// minutes
	ThrottleModifier_MINUTE ThrottleModifier_IntervalType = 2
	// hours
	ThrottleModifier_HOUR ThrottleModifier_IntervalType = 3
)

var ThrottleModifier_IntervalType_name = map[int32]string{
	0: "MILLISECOND",
	1: "SECOND",
	2: "MINUTE",
	3: "HOUR",
}
var ThrottleModifier_IntervalType_value = map[string]int32{
	"MILLISECOND": 0,
	"SECOND":      1,
	"MINUTE":      2,
	"HOUR":        3,
}

func (x ThrottleModifier_IntervalType) String() string {
	return proto.EnumName(ThrottleModifier_IntervalType_name, int32(x))
}
func (ThrottleModifier_IntervalType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor3, []int{12, 0}
}

//
// The Subscription message identifies a subscriber's interest in
// telemetry events.
//
type Subscription struct {
	// Return events matching one or more of the specified event
	// filters. If no event filters are specified, then no events
	// will be returned.
	EventFilter *EventFilter `protobuf:"bytes,1,opt,name=event_filter,json=eventFilter" json:"event_filter,omitempty"`
	// If not empty, then only return events from containers matched
	// by one or more of the specified container filters.
	ContainerFilter *ContainerFilter `protobuf:"bytes,2,opt,name=container_filter,json=containerFilter" json:"container_filter,omitempty"`
	// If not empty, then only return events that occurred after
	// the specified relative duration subtracted from the current
	// time (recorder time). If the resulting time is in the past, then the
	// subscription will search for historic events before streaming
	// live ones.
	SinceDuration *google_protobuf1.Int64Value `protobuf:"bytes,10,opt,name=since_duration,json=sinceDuration" json:"since_duration,omitempty"`
	// If not empty, then only return events that occurred before
	// the specified relative duration added to `since_duration`.
	// If `since_duration` is not supplied, return events from now and until
	// the specified relative duration is hit.
	ForDuration *google_protobuf1.Int64Value `protobuf:"bytes,11,opt,name=for_duration,json=forDuration" json:"for_duration,omitempty"`
	// If not empty, apply the specified modifier to the subscription.
	Modifier *Modifier `protobuf:"bytes,20,opt,name=modifier" json:"modifier,omitempty"`
}

func (m *Subscription) Reset()                    { *m = Subscription{} }
func (m *Subscription) String() string            { return proto.CompactTextString(m) }
func (*Subscription) ProtoMessage()               {}
func (*Subscription) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Subscription) GetEventFilter() *EventFilter {
	if m != nil {
		return m.EventFilter
	}
	return nil
}

func (m *Subscription) GetContainerFilter() *ContainerFilter {
	if m != nil {
		return m.ContainerFilter
	}
	return nil
}

func (m *Subscription) GetSinceDuration() *google_protobuf1.Int64Value {
	if m != nil {
		return m.SinceDuration
	}
	return nil
}

func (m *Subscription) GetForDuration() *google_protobuf1.Int64Value {
	if m != nil {
		return m.ForDuration
	}
	return nil
}

func (m *Subscription) GetModifier() *Modifier {
	if m != nil {
		return m.Modifier
	}
	return nil
}

// The ContainerFilter restricts events in the Subscription to the
// running containers indicated. All of the fields in this message are
// effectively "ORed" together to create the list of containers to
// monitor for the subscription.
type ContainerFilter struct {
	// Zero or more container IDs (e.g.
	// 254dd98a7bf1581560ddace9f98b7933bfb3c2f5fc0504ec1b8dcc9614bc7062)
	Ids []string `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
	// Zero or more container names (e.g. /ecstatic_darwin)
	Names []string `protobuf:"bytes,2,rep,name=names" json:"names,omitempty"`
	// Zero or more container image IDs (e.g.
	// d462265d362c919b7dd37f8ba80caa822d13704695f47c8fc42a1c2266ecd164)
	ImageIds []string `protobuf:"bytes,3,rep,name=image_ids,json=imageIds" json:"image_ids,omitempty"`
	// Container image name (shell-style globs are supported). May be of the
	// form "busybox", "foo/bar" or
	// "sha256:d462265d362c919b7dd37f8ba80caa822d13704695f47c8fc42a1c2266ecd164"
	ImageNames []string `protobuf:"bytes,4,rep,name=image_names,json=imageNames" json:"image_names,omitempty"`
}

func (m *ContainerFilter) Reset()                    { *m = ContainerFilter{} }
func (m *ContainerFilter) String() string            { return proto.CompactTextString(m) }
func (*ContainerFilter) ProtoMessage()               {}
func (*ContainerFilter) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *ContainerFilter) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *ContainerFilter) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *ContainerFilter) GetImageIds() []string {
	if m != nil {
		return m.ImageIds
	}
	return nil
}

func (m *ContainerFilter) GetImageNames() []string {
	if m != nil {
		return m.ImageNames
	}
	return nil
}

// The EventFilter specifies events to include. All of the specified
// fields are effectively "ORed" together to create the list of events
// included in the Subscription.
type EventFilter struct {
	// Zero or more filters specifying which system calls to include
	SyscallEvents []*SyscallEventFilter `protobuf:"bytes,1,rep,name=syscall_events,json=syscallEvents" json:"syscall_events,omitempty"`
	// Zero or more filters specifying which process events to include
	ProcessEvents []*ProcessEventFilter `protobuf:"bytes,2,rep,name=process_events,json=processEvents" json:"process_events,omitempty"`
	// Zero or more filters specifying which file events to include
	FileEvents []*FileEventFilter `protobuf:"bytes,3,rep,name=file_events,json=fileEvents" json:"file_events,omitempty"`
	// Zero or more kernel functional calls to include
	KernelEvents []*KernelFunctionCallFilter `protobuf:"bytes,4,rep,name=kernel_events,json=kernelEvents" json:"kernel_events,omitempty"`
	// Zero or more network events to include
	NetworkEvents []*NetworkEventFilter `protobuf:"bytes,5,rep,name=network_events,json=networkEvents" json:"network_events,omitempty"`
	// Zero or more container events to include
	ContainerEvents []*ContainerEventFilter `protobuf:"bytes,10,rep,name=container_events,json=containerEvents" json:"container_events,omitempty"`
	// Zero or more character generators to configure and return events from
	// (for debugging)
	ChargenEvents []*ChargenEventFilter `protobuf:"bytes,100,rep,name=chargen_events,json=chargenEvents" json:"chargen_events,omitempty"`
	// Zero or more ticker generators to configure and return events from
	// (for debugging)
	TickerEvents []*TickerEventFilter `protobuf:"bytes,101,rep,name=ticker_events,json=tickerEvents" json:"ticker_events,omitempty"`
}

func (m *EventFilter) Reset()                    { *m = EventFilter{} }
func (m *EventFilter) String() string            { return proto.CompactTextString(m) }
func (*EventFilter) ProtoMessage()               {}
func (*EventFilter) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *EventFilter) GetSyscallEvents() []*SyscallEventFilter {
	if m != nil {
		return m.SyscallEvents
	}
	return nil
}

func (m *EventFilter) GetProcessEvents() []*ProcessEventFilter {
	if m != nil {
		return m.ProcessEvents
	}
	return nil
}

func (m *EventFilter) GetFileEvents() []*FileEventFilter {
	if m != nil {
		return m.FileEvents
	}
	return nil
}

func (m *EventFilter) GetKernelEvents() []*KernelFunctionCallFilter {
	if m != nil {
		return m.KernelEvents
	}
	return nil
}

func (m *EventFilter) GetNetworkEvents() []*NetworkEventFilter {
	if m != nil {
		return m.NetworkEvents
	}
	return nil
}

func (m *EventFilter) GetContainerEvents() []*ContainerEventFilter {
	if m != nil {
		return m.ContainerEvents
	}
	return nil
}

func (m *EventFilter) GetChargenEvents() []*ChargenEventFilter {
	if m != nil {
		return m.ChargenEvents
	}
	return nil
}

func (m *EventFilter) GetTickerEvents() []*TickerEventFilter {
	if m != nil {
		return m.TickerEvents
	}
	return nil
}

// The SyscallEventFilter specifies which system call events to
// include in the Subscription. The specified fields are effectively
// "ANDed" to specify a matching event.
type SyscallEventFilter struct {
	// Required; type of system call event (entry or exit)
	Type             SyscallEventType `protobuf:"varint,1,opt,name=type,enum=capsule8.api.v0.SyscallEventType" json:"type,omitempty"`
	FilterExpression *Expression      `protobuf:"bytes,100,opt,name=filter_expression,json=filterExpression" json:"filter_expression,omitempty"`
	// Required; system call number from
	// arch/x86/entry/syscalls/syscall_64.tbl
	Id *google_protobuf1.Int64Value `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	// Optional; precise value of a particular system call argument
	Arg0 *google_protobuf1.UInt64Value `protobuf:"bytes,10,opt,name=arg0" json:"arg0,omitempty"`
	Arg1 *google_protobuf1.UInt64Value `protobuf:"bytes,11,opt,name=arg1" json:"arg1,omitempty"`
	Arg2 *google_protobuf1.UInt64Value `protobuf:"bytes,12,opt,name=arg2" json:"arg2,omitempty"`
	Arg3 *google_protobuf1.UInt64Value `protobuf:"bytes,13,opt,name=arg3" json:"arg3,omitempty"`
	Arg4 *google_protobuf1.UInt64Value `protobuf:"bytes,14,opt,name=arg4" json:"arg4,omitempty"`
	Arg5 *google_protobuf1.UInt64Value `protobuf:"bytes,15,opt,name=arg5" json:"arg5,omitempty"`
	// Optional; return value of the system call (if type indicates exit).
	Ret *google_protobuf1.Int64Value `protobuf:"bytes,20,opt,name=ret" json:"ret,omitempty"`
}

func (m *SyscallEventFilter) Reset()                    { *m = SyscallEventFilter{} }
func (m *SyscallEventFilter) String() string            { return proto.CompactTextString(m) }
func (*SyscallEventFilter) ProtoMessage()               {}
func (*SyscallEventFilter) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *SyscallEventFilter) GetType() SyscallEventType {
	if m != nil {
		return m.Type
	}
	return SyscallEventType_SYSCALL_EVENT_TYPE_UNKNOWN
}

func (m *SyscallEventFilter) GetFilterExpression() *Expression {
	if m != nil {
		return m.FilterExpression
	}
	return nil
}

func (m *SyscallEventFilter) GetId() *google_protobuf1.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *SyscallEventFilter) GetArg0() *google_protobuf1.UInt64Value {
	if m != nil {
		return m.Arg0
	}
	return nil
}

func (m *SyscallEventFilter) GetArg1() *google_protobuf1.UInt64Value {
	if m != nil {
		return m.Arg1
	}
	return nil
}

func (m *SyscallEventFilter) GetArg2() *google_protobuf1.UInt64Value {
	if m != nil {
		return m.Arg2
	}
	return nil
}

func (m *SyscallEventFilter) GetArg3() *google_protobuf1.UInt64Value {
	if m != nil {
		return m.Arg3
	}
	return nil
}

func (m *SyscallEventFilter) GetArg4() *google_protobuf1.UInt64Value {
	if m != nil {
		return m.Arg4
	}
	return nil
}

func (m *SyscallEventFilter) GetArg5() *google_protobuf1.UInt64Value {
	if m != nil {
		return m.Arg5
	}
	return nil
}

func (m *SyscallEventFilter) GetRet() *google_protobuf1.Int64Value {
	if m != nil {
		return m.Ret
	}
	return nil
}

// The ProcessEventFilter specifies which process events to include in
// the Subscription. The specified fields are effectively "ANDed" to
// specify a matching event.
type ProcessEventFilter struct {
	// Required; the process event type to match
	Type             ProcessEventType `protobuf:"varint,1,opt,name=type,enum=capsule8.api.v0.ProcessEventType" json:"type,omitempty"`
	FilterExpression *Expression      `protobuf:"bytes,100,opt,name=filter_expression,json=filterExpression" json:"filter_expression,omitempty"`
	// Optional; require exact match on the filename passed to execve(2)
	ExecFilename *google_protobuf1.StringValue `protobuf:"bytes,12,opt,name=exec_filename,json=execFilename" json:"exec_filename,omitempty"`
	// Optional; require pattern match on the filename passed to execve(2)
	ExecFilenamePattern *google_protobuf1.StringValue `protobuf:"bytes,13,opt,name=exec_filename_pattern,json=execFilenamePattern" json:"exec_filename_pattern,omitempty"`
	// Optional; require exact match on exit code
	ExitCode *google_protobuf1.Int32Value `protobuf:"bytes,14,opt,name=exit_code,json=exitCode" json:"exit_code,omitempty"`
}

func (m *ProcessEventFilter) Reset()                    { *m = ProcessEventFilter{} }
func (m *ProcessEventFilter) String() string            { return proto.CompactTextString(m) }
func (*ProcessEventFilter) ProtoMessage()               {}
func (*ProcessEventFilter) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *ProcessEventFilter) GetType() ProcessEventType {
	if m != nil {
		return m.Type
	}
	return ProcessEventType_PROCESS_EVENT_TYPE_UNKNOWN
}

func (m *ProcessEventFilter) GetFilterExpression() *Expression {
	if m != nil {
		return m.FilterExpression
	}
	return nil
}

func (m *ProcessEventFilter) GetExecFilename() *google_protobuf1.StringValue {
	if m != nil {
		return m.ExecFilename
	}
	return nil
}

func (m *ProcessEventFilter) GetExecFilenamePattern() *google_protobuf1.StringValue {
	if m != nil {
		return m.ExecFilenamePattern
	}
	return nil
}

func (m *ProcessEventFilter) GetExitCode() *google_protobuf1.Int32Value {
	if m != nil {
		return m.ExitCode
	}
	return nil
}

// The FileEventFilter specifies which file events to include in the
// Subscription. The specified fields are effectively "ANDed" to
// specify a matching event.
type FileEventFilter struct {
	// Required; the file event type to match
	Type             FileEventType `protobuf:"varint,1,opt,name=type,enum=capsule8.api.v0.FileEventType" json:"type,omitempty"`
	FilterExpression *Expression   `protobuf:"bytes,100,opt,name=filter_expression,json=filterExpression" json:"filter_expression,omitempty"`
	// Optional; require exact match on the filename being acted upon
	Filename *google_protobuf1.StringValue `protobuf:"bytes,10,opt,name=filename" json:"filename,omitempty"`
	// Optional; require pattern match on the filename being acted upon
	FilenamePattern *google_protobuf1.StringValue `protobuf:"bytes,11,opt,name=filename_pattern,json=filenamePattern" json:"filename_pattern,omitempty"`
	// Optional; for file open events, require a match of the bits set
	// for the open(2) flags argument
	OpenFlagsMask *google_protobuf1.Int32Value `protobuf:"bytes,12,opt,name=open_flags_mask,json=openFlagsMask" json:"open_flags_mask,omitempty"`
	// Optional; for file open events, require a match of the bits set
	// for the open(2) or creat(2) mode argument
	CreateModeMask *google_protobuf1.Int32Value `protobuf:"bytes,13,opt,name=create_mode_mask,json=createModeMask" json:"create_mode_mask,omitempty"`
}

func (m *FileEventFilter) Reset()                    { *m = FileEventFilter{} }
func (m *FileEventFilter) String() string            { return proto.CompactTextString(m) }
func (*FileEventFilter) ProtoMessage()               {}
func (*FileEventFilter) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *FileEventFilter) GetType() FileEventType {
	if m != nil {
		return m.Type
	}
	return FileEventType_FILE_EVENT_TYPE_UNKNOWN
}

func (m *FileEventFilter) GetFilterExpression() *Expression {
	if m != nil {
		return m.FilterExpression
	}
	return nil
}

func (m *FileEventFilter) GetFilename() *google_protobuf1.StringValue {
	if m != nil {
		return m.Filename
	}
	return nil
}

func (m *FileEventFilter) GetFilenamePattern() *google_protobuf1.StringValue {
	if m != nil {
		return m.FilenamePattern
	}
	return nil
}

func (m *FileEventFilter) GetOpenFlagsMask() *google_protobuf1.Int32Value {
	if m != nil {
		return m.OpenFlagsMask
	}
	return nil
}

func (m *FileEventFilter) GetCreateModeMask() *google_protobuf1.Int32Value {
	if m != nil {
		return m.CreateModeMask
	}
	return nil
}

// The KernelFunctionCallFilter specifies which kernel function call
// events to include in the Subscription. The arguments map defines
// values that will be fetched at each call and returned along with
// the event. In order to minimize event volume, a filter may be
// included that filters the kernel function calls based on the
// observed values of the specified arguments at the time of the
// kernel function call.
type KernelFunctionCallFilter struct {
	// Required; the kernel function call event type to match
	Type KernelFunctionCallEventType `protobuf:"varint,1,opt,name=type,enum=capsule8.api.v0.KernelFunctionCallEventType" json:"type,omitempty"`
	// Required; the kernel symbol to match on
	Symbol string `protobuf:"bytes,10,opt,name=symbol" json:"symbol,omitempty"`
	// Optional; the field names and data to be returned by the kernel
	// when the event triggers. Note that this is a map. The keys are the
	// names to assign to the returned fields, and the values are a string
	// describing the data to return, usually an expression involving the
	// register containing the desired data and a suffix indicating the
	// type of the data (e.g., "s32", "string", "u64", etc.). This map is
	// used to construct the "fetchargs" passed to the kernel when creating
	// the kernel probe.
	Arguments map[string]string `protobuf:"bytes,11,rep,name=arguments" json:"arguments,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Optional; a filter to apply to kernel probe.
	FilterExpression *Expression `protobuf:"bytes,100,opt,name=filter_expression,json=filterExpression" json:"filter_expression,omitempty"`
}

func (m *KernelFunctionCallFilter) Reset()                    { *m = KernelFunctionCallFilter{} }
func (m *KernelFunctionCallFilter) String() string            { return proto.CompactTextString(m) }
func (*KernelFunctionCallFilter) ProtoMessage()               {}
func (*KernelFunctionCallFilter) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func (m *KernelFunctionCallFilter) GetType() KernelFunctionCallEventType {
	if m != nil {
		return m.Type
	}
	return KernelFunctionCallEventType_KERNEL_FUNCTION_CALL_EVENT_TYPE_UNKNOWN
}

func (m *KernelFunctionCallFilter) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *KernelFunctionCallFilter) GetArguments() map[string]string {
	if m != nil {
		return m.Arguments
	}
	return nil
}

func (m *KernelFunctionCallFilter) GetFilterExpression() *Expression {
	if m != nil {
		return m.FilterExpression
	}
	return nil
}

// The NetworkEventFilter specifies which network events to include in
// the Subscription. The included filter can be used to specify
// precisely which network events should be included.
type NetworkEventFilter struct {
	// Required; the network event type to match
	Type NetworkEventType `protobuf:"varint,1,opt,name=type,enum=capsule8.api.v0.NetworkEventType" json:"type,omitempty"`
	// Optional; a filter to apply to events. Only events for which the
	// evaluation of the filter expression is true will be returned.
	FilterExpression *Expression `protobuf:"bytes,100,opt,name=filter_expression,json=filterExpression" json:"filter_expression,omitempty"`
}

func (m *NetworkEventFilter) Reset()                    { *m = NetworkEventFilter{} }
func (m *NetworkEventFilter) String() string            { return proto.CompactTextString(m) }
func (*NetworkEventFilter) ProtoMessage()               {}
func (*NetworkEventFilter) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

func (m *NetworkEventFilter) GetType() NetworkEventType {
	if m != nil {
		return m.Type
	}
	return NetworkEventType_NETWORK_EVENT_TYPE_UNKNOWN
}

func (m *NetworkEventFilter) GetFilterExpression() *Expression {
	if m != nil {
		return m.FilterExpression
	}
	return nil
}

// The ContainerEventFilter specifies which container lifecycle events
// to include in the Subscription. In order to restrict them to
// specific containers, use the ContainerFilter.
type ContainerEventFilter struct {
	// Required, specify the particular type of event type to match
	Type ContainerEventType `protobuf:"varint,1,opt,name=type,enum=capsule8.api.v0.ContainerEventType" json:"type,omitempty"`
	// Optional, specifies how much detail to include in container events
	View ContainerEventView `protobuf:"varint,2,opt,name=view,enum=capsule8.api.v0.ContainerEventView" json:"view,omitempty"`
	// Optional; a filter to apply to events. Only events for which the
	// evaluation of the filter expression is true will be returned.
	FilterExpression *Expression `protobuf:"bytes,100,opt,name=filter_expression,json=filterExpression" json:"filter_expression,omitempty"`
}

func (m *ContainerEventFilter) Reset()                    { *m = ContainerEventFilter{} }
func (m *ContainerEventFilter) String() string            { return proto.CompactTextString(m) }
func (*ContainerEventFilter) ProtoMessage()               {}
func (*ContainerEventFilter) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{8} }

func (m *ContainerEventFilter) GetType() ContainerEventType {
	if m != nil {
		return m.Type
	}
	return ContainerEventType_CONTAINER_EVENT_TYPE_UNKNOWN
}

func (m *ContainerEventFilter) GetView() ContainerEventView {
	if m != nil {
		return m.View
	}
	return ContainerEventView_BASIC
}

func (m *ContainerEventFilter) GetFilterExpression() *Expression {
	if m != nil {
		return m.FilterExpression
	}
	return nil
}

// The ChargenEventFilter configures a character stream generator and
// includes events from it in the Subscription.
type ChargenEventFilter struct {
	// Required; the length of character sequence strings to generate
	Length uint64 `protobuf:"varint,1,opt,name=length" json:"length,omitempty"`
}

func (m *ChargenEventFilter) Reset()                    { *m = ChargenEventFilter{} }
func (m *ChargenEventFilter) String() string            { return proto.CompactTextString(m) }
func (*ChargenEventFilter) ProtoMessage()               {}
func (*ChargenEventFilter) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{9} }

func (m *ChargenEventFilter) GetLength() uint64 {
	if m != nil {
		return m.Length
	}
	return 0
}

// The TickerEventFilter configures a ticker stream generator and
// includes events from it in the Subscription.
type TickerEventFilter struct {
	// Required; the interval at which ticker events are generated
	Interval int64 `protobuf:"varint,1,opt,name=interval" json:"interval,omitempty"`
}

func (m *TickerEventFilter) Reset()                    { *m = TickerEventFilter{} }
func (m *TickerEventFilter) String() string            { return proto.CompactTextString(m) }
func (*TickerEventFilter) ProtoMessage()               {}
func (*TickerEventFilter) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{10} }

func (m *TickerEventFilter) GetInterval() int64 {
	if m != nil {
		return m.Interval
	}
	return 0
}

// Modifier specifies which stream modifiers to apply if any. For a given
// stream, a modifier can apply a throttle or limit etc. Modifiers can be
// used together.
type Modifier struct {
	Throttle *ThrottleModifier `protobuf:"bytes,1,opt,name=throttle" json:"throttle,omitempty"`
	Limit    *LimitModifier    `protobuf:"bytes,2,opt,name=limit" json:"limit,omitempty"`
}

func (m *Modifier) Reset()                    { *m = Modifier{} }
func (m *Modifier) String() string            { return proto.CompactTextString(m) }
func (*Modifier) ProtoMessage()               {}
func (*Modifier) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{11} }

func (m *Modifier) GetThrottle() *ThrottleModifier {
	if m != nil {
		return m.Throttle
	}
	return nil
}

func (m *Modifier) GetLimit() *LimitModifier {
	if m != nil {
		return m.Limit
	}
	return nil
}

// The ThrottleModifier modulates events sent by the Sensor to one per
// time interval specified.
type ThrottleModifier struct {
	// Required; the interval to use
	Interval int64 `protobuf:"varint,1,opt,name=interval" json:"interval,omitempty"`
	// Required; the intreval type (milliseconds, seconds, etc.)
	IntervalType ThrottleModifier_IntervalType `protobuf:"varint,2,opt,name=interval_type,json=intervalType,enum=capsule8.api.v0.ThrottleModifier_IntervalType" json:"interval_type,omitempty"`
}

func (m *ThrottleModifier) Reset()                    { *m = ThrottleModifier{} }
func (m *ThrottleModifier) String() string            { return proto.CompactTextString(m) }
func (*ThrottleModifier) ProtoMessage()               {}
func (*ThrottleModifier) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{12} }

func (m *ThrottleModifier) GetInterval() int64 {
	if m != nil {
		return m.Interval
	}
	return 0
}

func (m *ThrottleModifier) GetIntervalType() ThrottleModifier_IntervalType {
	if m != nil {
		return m.IntervalType
	}
	return ThrottleModifier_MILLISECOND
}

// The LimitModifier cancels the subscription on each Sensor after the
// specified number of events. The entire Subscription may return more
// events that this depending on how many active Sensors there are.
type LimitModifier struct {
	// Limit the number of events
	Limit int64 `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
}

func (m *LimitModifier) Reset()                    { *m = LimitModifier{} }
func (m *LimitModifier) String() string            { return proto.CompactTextString(m) }
func (*LimitModifier) ProtoMessage()               {}
func (*LimitModifier) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{13} }

func (m *LimitModifier) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func init() {
	proto.RegisterType((*Subscription)(nil), "capsule8.api.v0.Subscription")
	proto.RegisterType((*ContainerFilter)(nil), "capsule8.api.v0.ContainerFilter")
	proto.RegisterType((*EventFilter)(nil), "capsule8.api.v0.EventFilter")
	proto.RegisterType((*SyscallEventFilter)(nil), "capsule8.api.v0.SyscallEventFilter")
	proto.RegisterType((*ProcessEventFilter)(nil), "capsule8.api.v0.ProcessEventFilter")
	proto.RegisterType((*FileEventFilter)(nil), "capsule8.api.v0.FileEventFilter")
	proto.RegisterType((*KernelFunctionCallFilter)(nil), "capsule8.api.v0.KernelFunctionCallFilter")
	proto.RegisterType((*NetworkEventFilter)(nil), "capsule8.api.v0.NetworkEventFilter")
	proto.RegisterType((*ContainerEventFilter)(nil), "capsule8.api.v0.ContainerEventFilter")
	proto.RegisterType((*ChargenEventFilter)(nil), "capsule8.api.v0.ChargenEventFilter")
	proto.RegisterType((*TickerEventFilter)(nil), "capsule8.api.v0.TickerEventFilter")
	proto.RegisterType((*Modifier)(nil), "capsule8.api.v0.Modifier")
	proto.RegisterType((*ThrottleModifier)(nil), "capsule8.api.v0.ThrottleModifier")
	proto.RegisterType((*LimitModifier)(nil), "capsule8.api.v0.LimitModifier")
	proto.RegisterEnum("capsule8.api.v0.ContainerEventView", ContainerEventView_name, ContainerEventView_value)
	proto.RegisterEnum("capsule8.api.v0.ThrottleModifier_IntervalType", ThrottleModifier_IntervalType_name, ThrottleModifier_IntervalType_value)
}

func init() { proto.RegisterFile("capsule8/api/v0/subscription.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 1194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x57, 0xdb, 0x6e, 0x1a, 0x47,
	0x18, 0x0e, 0x07, 0x47, 0xf0, 0x73, 0xcc, 0xd4, 0xad, 0xa8, 0x13, 0xa5, 0xee, 0x46, 0x96, 0x92,
	0x36, 0x05, 0x07, 0xdb, 0x8d, 0x55, 0xf5, 0xe4, 0x10, 0x93, 0xd0, 0xd8, 0x8e, 0xb5, 0xd8, 0xbe,
	0x45, 0xeb, 0xdd, 0x01, 0x8f, 0x58, 0x76, 0x57, 0x33, 0x03, 0x36, 0x57, 0x7d, 0x8a, 0x5e, 0xb6,
	0x8f, 0xd3, 0x07, 0xa8, 0x2a, 0xf5, 0x05, 0x7a, 0xd7, 0x97, 0xa8, 0xe6, 0xb0, 0xb0, 0xb0, 0xc1,
	0x70, 0x11, 0xdf, 0xed, 0xfc, 0xf3, 0x7d, 0x1f, 0xff, 0x69, 0xfe, 0x19, 0xc0, 0xb0, 0xad, 0x80,
	0x0d, 0x5d, 0xbc, 0x5f, 0xb3, 0x02, 0x52, 0x1b, 0x6d, 0xd7, 0xd8, 0xf0, 0x92, 0xd9, 0x94, 0x04,
	0x9c, 0xf8, 0x5e, 0x35, 0xa0, 0x3e, 0xf7, 0x51, 0x29, 0xc4, 0x54, 0xad, 0x80, 0x54, 0x47, 0xdb,
	0x1b, 0x5b, 0xf3, 0x24, 0x8e, 0x5d, 0x3c, 0xc0, 0x9c, 0x8e, 0x3b, 0x78, 0x84, 0x3d, 0xae, 0x78,
	0x1b, 0x9b, 0xf3, 0x30, 0x7c, 0x13, 0x50, 0xcc, 0xd8, 0x44, 0x79, 0xe3, 0x71, 0xcf, 0xf7, 0x7b,
	0x2e, 0xae, 0xc9, 0xd5, 0xe5, 0xb0, 0x5b, 0xbb, 0xa6, 0x56, 0x10, 0x60, 0xca, 0xd4, 0xbe, 0xf1,
	0x4f, 0x12, 0xf2, 0xed, 0x88, 0x43, 0xe8, 0x27, 0xc8, 0xcb, 0x5f, 0xe8, 0x74, 0x89, 0xcb, 0x31,
	0xad, 0x24, 0x36, 0x13, 0x4f, 0x73, 0xf5, 0x47, 0xd5, 0x39, 0x0f, 0xab, 0x87, 0x02, 0xd4, 0x94,
	0x18, 0x33, 0x87, 0xa7, 0x0b, 0xf4, 0x0e, 0xca, 0xb6, 0xef, 0x71, 0x8b, 0x78, 0x98, 0x86, 0x22,
	0x49, 0x29, 0xb2, 0x19, 0x13, 0x69, 0x84, 0x40, 0x2d, 0x54, 0xb2, 0x67, 0x0d, 0xe8, 0x15, 0x14,
	0x19, 0xf1, 0x6c, 0xdc, 0x71, 0x86, 0xd4, 0x12, 0xfe, 0x55, 0x40, 0x4a, 0x3d, 0xac, 0xaa, 0xb8,
	0xaa, 0x61, 0x5c, 0xd5, 0x96, 0xc7, 0xbf, 0xdd, 0xbd, 0xb0, 0xdc, 0x21, 0x36, 0x0b, 0x92, 0xf2,
	0x5a, 0x33, 0xd0, 0x8f, 0x90, 0xef, 0xfa, 0x74, 0xaa, 0x90, 0x5b, 0xae, 0x90, 0xeb, 0xfa, 0x74,
	0xc2, 0xdf, 0x83, 0xcc, 0xc0, 0x77, 0x48, 0x97, 0x60, 0x5a, 0x59, 0x97, 0xdc, 0xcf, 0x63, 0x81,
	0x1c, 0x6b, 0x80, 0x39, 0x81, 0x1a, 0xd7, 0x50, 0x9a, 0x0b, 0x0f, 0x95, 0x21, 0x45, 0x1c, 0x56,
	0x49, 0x6c, 0xa6, 0x9e, 0x66, 0x4d, 0xf1, 0x89, 0xd6, 0x61, 0xcd, 0xb3, 0x06, 0x98, 0x55, 0x92,
	0xd2, 0xa6, 0x16, 0xe8, 0x21, 0x64, 0xc9, 0xc0, 0xea, 0xe1, 0x8e, 0x40, 0xa7, 0xe4, 0x4e, 0x46,
	0x1a, 0x5a, 0x0e, 0x43, 0x5f, 0x40, 0x4e, 0x6d, 0x2a, 0x62, 0x5a, 0x6e, 0x83, 0x34, 0x9d, 0x08,
	0x8b, 0xf1, 0x5f, 0x1a, 0x72, 0x91, 0xea, 0xa0, 0x5f, 0xa0, 0xc8, 0xc6, 0xcc, 0xb6, 0x5c, 0x57,
	0xf5, 0x8e, 0x72, 0x20, 0x57, 0x7f, 0x12, 0x8b, 0xa2, 0xad, 0x60, 0xd1, 0xd2, 0x16, 0x58, 0xc4,
	0xc6, 0x84, 0x56, 0x40, 0x7d, 0x1b, 0x33, 0x16, 0x6a, 0x25, 0x17, 0x68, 0x9d, 0x2a, 0xd8, 0x8c,
	0x56, 0x10, 0xb1, 0x31, 0x74, 0x00, 0xb9, 0x2e, 0x71, 0x71, 0x28, 0x94, 0x92, 0x42, 0xf1, 0x1e,
	0x69, 0x12, 0x17, 0x47, 0x55, 0xa0, 0x1b, 0x1a, 0x18, 0x3a, 0x81, 0x42, 0x1f, 0x53, 0x0f, 0x4f,
	0x22, 0x4b, 0x4b, 0x91, 0x67, 0x31, 0x91, 0x77, 0x12, 0xd5, 0x1c, 0x7a, 0xb6, 0x28, 0x69, 0xc3,
	0x72, 0x5d, 0xad, 0x96, 0x57, 0xfc, 0x69, 0x78, 0x1e, 0xe6, 0xd7, 0x3e, 0xed, 0x87, 0x82, 0x6b,
	0x0b, 0xc2, 0x3b, 0x51, 0xb0, 0x99, 0xf0, 0xbc, 0x88, 0x8d, 0xa1, 0xd3, 0xe8, 0x39, 0xd0, 0x6a,
	0x20, 0xd5, 0xb6, 0x16, 0x9f, 0x83, 0xa8, 0xde, 0xf4, 0x30, 0x4c, 0xbd, 0xb3, 0xaf, 0x2c, 0xda,
	0xc3, 0x5e, 0xa8, 0xe7, 0x2c, 0xf0, 0xae, 0xa1, 0x60, 0x33, 0xde, 0xd9, 0x11, 0x1b, 0x43, 0x6f,
	0xa0, 0xc0, 0x89, 0xdd, 0x9f, 0xba, 0x86, 0xa5, 0x94, 0x11, 0x93, 0x3a, 0x93, 0xa8, 0xa8, 0x52,
	0x9e, 0x4f, 0x4d, 0xcc, 0xf8, 0x3d, 0x0d, 0x28, 0xde, 0x37, 0x68, 0x0f, 0xd2, 0x7c, 0x1c, 0x60,
	0x39, 0x3e, 0x8a, 0xf5, 0x2f, 0x6f, 0x6d, 0xb5, 0xb3, 0x71, 0x80, 0x4d, 0x09, 0x47, 0x6f, 0xe1,
	0x81, 0x1a, 0x19, 0x9d, 0xe9, 0x24, 0xab, 0x38, 0xfa, 0xc0, 0xc6, 0x46, 0xd0, 0x04, 0x62, 0x96,
	0x15, 0x6b, 0x6a, 0x41, 0x5f, 0x43, 0x92, 0x38, 0x7a, 0xf0, 0xdc, 0x7a, 0xd6, 0x93, 0xc4, 0x41,
	0xdb, 0x90, 0xb6, 0x68, 0x6f, 0x5b, 0x0f, 0x97, 0x47, 0x31, 0xf8, 0x79, 0x04, 0x2f, 0x91, 0x9a,
	0xf1, 0x42, 0x0f, 0x93, 0xe5, 0x8c, 0x17, 0x9a, 0x51, 0xaf, 0xe4, 0x57, 0x64, 0xd4, 0x35, 0x63,
	0xa7, 0x52, 0x58, 0x91, 0xb1, 0xa3, 0x19, 0xbb, 0x95, 0xe2, 0x8a, 0x8c, 0x5d, 0xcd, 0xd8, 0xab,
	0x94, 0x56, 0x64, 0xec, 0xa1, 0x6f, 0x20, 0x45, 0x31, 0xd7, 0x93, 0xf0, 0xd6, 0xcc, 0x0a, 0x9c,
	0xf1, 0x6f, 0x12, 0x50, 0x7c, 0x16, 0x2c, 0xed, 0x8f, 0x28, 0xe5, 0x4e, 0xfa, 0xe3, 0x00, 0x0a,
	0xf8, 0x06, 0xdb, 0xe2, 0x86, 0xc2, 0x62, 0x92, 0x2e, 0xac, 0x4b, 0x9b, 0x53, 0xe2, 0xf5, 0x54,
	0x44, 0x79, 0x41, 0x69, 0x6a, 0x06, 0x3a, 0x85, 0x4f, 0x67, 0x24, 0x3a, 0x81, 0xc5, 0x39, 0xa6,
	0xde, 0xc2, 0x82, 0x45, 0xa5, 0x3e, 0x89, 0x4a, 0x9d, 0x2a, 0x22, 0xda, 0x87, 0x2c, 0xbe, 0x21,
	0xbc, 0x63, 0xfb, 0x0e, 0xd6, 0x45, 0xfc, 0x60, 0x86, 0x77, 0xea, 0x4a, 0x24, 0x23, 0xd0, 0x0d,
	0xdf, 0xc1, 0xc6, 0x1f, 0x29, 0x28, 0xcd, 0x4d, 0x4a, 0x54, 0x9f, 0xc9, 0xf1, 0xe3, 0xc5, 0x93,
	0xf5, 0x4e, 0x12, 0xbc, 0x0f, 0x99, 0x49, 0x6e, 0x61, 0x85, 0x84, 0x4c, 0xd0, 0xe8, 0x0d, 0x94,
	0x63, 0x29, 0xcd, 0xad, 0xa0, 0x50, 0xea, 0xce, 0xa5, 0xb3, 0x01, 0x25, 0x3f, 0xc0, 0x5e, 0xa7,
	0xeb, 0x5a, 0x3d, 0xd6, 0x19, 0x58, 0xac, 0xaf, 0xab, 0x7c, 0x6b, 0x52, 0x0b, 0x82, 0xd3, 0x14,
	0x94, 0x63, 0x8b, 0xf5, 0xd1, 0x21, 0x94, 0x6d, 0x8a, 0x2d, 0x8e, 0x3b, 0x03, 0xdf, 0xc1, 0x4a,
	0xa5, 0xb0, 0x5c, 0xa5, 0xa8, 0x48, 0xc7, 0xbe, 0x83, 0x85, 0x8c, 0xf1, 0x77, 0x12, 0x2a, 0x8b,
	0x6e, 0x21, 0xf4, 0xf3, 0x4c, 0xa5, 0x9e, 0xaf, 0x70, 0x7d, 0xcd, 0xd7, 0xed, 0x33, 0xb8, 0xcf,
	0xc6, 0x83, 0x4b, 0xdf, 0x95, 0xb9, 0xce, 0x9a, 0x7a, 0x85, 0x2e, 0x20, 0x6b, 0xd1, 0xde, 0x70,
	0x20, 0x67, 0x7c, 0x4e, 0xce, 0xf8, 0xfd, 0x95, 0x6f, 0xc7, 0xea, 0x41, 0x48, 0x3d, 0xf4, 0x38,
	0x1d, 0x9b, 0x53, 0xa9, 0x8f, 0xd7, 0x27, 0x1b, 0xdf, 0x43, 0x71, 0xf6, 0x67, 0xc4, 0x33, 0xa9,
	0x8f, 0xc7, 0x32, 0x19, 0x59, 0x53, 0x7c, 0x8a, 0x67, 0xd2, 0x48, 0x64, 0x55, 0xce, 0xf3, 0xac,
	0xa9, 0x16, 0xdf, 0x25, 0xf7, 0x13, 0xc6, 0x6f, 0x09, 0x40, 0xf1, 0xbb, 0x78, 0xe9, 0x78, 0x89,
	0x52, 0xee, 0xa2, 0xfb, 0x8d, 0xbf, 0x12, 0xb0, 0xfe, 0xa1, 0x5b, 0x1d, 0xbd, 0x9c, 0xf1, 0xec,
	0xc9, 0x92, 0xa7, 0x40, 0xc4, 0xb7, 0x97, 0x90, 0x1e, 0x11, 0x7c, 0x2d, 0x53, 0xb0, 0x9c, 0x78,
	0x41, 0xf0, 0xb5, 0x29, 0x09, 0x1f, 0x31, 0xa8, 0xe7, 0x80, 0xe2, 0x2f, 0x0b, 0xd1, 0x7a, 0x2e,
	0xf6, 0x7a, 0xfc, 0x4a, 0xc6, 0x94, 0x36, 0xf5, 0xca, 0xa8, 0xc1, 0x83, 0xd8, 0xe3, 0x01, 0x6d,
	0x40, 0x86, 0x78, 0x1c, 0xd3, 0x91, 0xe5, 0x4a, 0x78, 0xca, 0x9c, 0xac, 0x8d, 0x5f, 0x21, 0x13,
	0xbe, 0xa3, 0xd1, 0x0f, 0x90, 0xe1, 0x57, 0xd4, 0xe7, 0xdc, 0xc5, 0xfa, 0x2f, 0x48, 0xbc, 0x88,
	0x67, 0x1a, 0x30, 0x7d, 0x7c, 0x87, 0x14, 0xb4, 0x0b, 0x6b, 0x2e, 0x19, 0x10, 0xae, 0x1f, 0x00,
	0xf1, 0xd9, 0x77, 0x24, 0x76, 0x27, 0x44, 0x05, 0x36, 0xfe, 0x4c, 0x40, 0x79, 0x5e, 0xf4, 0x36,
	0x8f, 0x51, 0x1b, 0x0a, 0xe1, 0x77, 0x47, 0x56, 0x55, 0x15, 0xa7, 0xba, 0xd4, 0x55, 0x31, 0x29,
	0x24, 0x4d, 0x16, 0x38, 0x4f, 0x22, 0x2b, 0xe3, 0x00, 0xf2, 0xd1, 0x5d, 0x54, 0x82, 0xdc, 0x71,
	0xeb, 0xe8, 0xa8, 0xd5, 0x3e, 0x6c, 0xbc, 0x3f, 0x79, 0x5d, 0xbe, 0x87, 0x00, 0xee, 0xeb, 0xef,
	0x84, 0xf8, 0x3e, 0x6e, 0x9d, 0x9c, 0x9f, 0x1d, 0x96, 0x93, 0x28, 0x03, 0xe9, 0xb7, 0xef, 0xcf,
	0xcd, 0x72, 0xca, 0xd8, 0x82, 0xc2, 0x4c, 0x80, 0xe2, 0x00, 0xa9, 0x7c, 0xa8, 0x08, 0xd4, 0xe2,
	0xab, 0x67, 0x80, 0xe2, 0x5d, 0x83, 0xb2, 0xb0, 0xf6, 0xea, 0xa0, 0xdd, 0x6a, 0x94, 0xef, 0x09,
	0xc5, 0xe6, 0xf9, 0xd1, 0x51, 0x39, 0x71, 0x79, 0x5f, 0xce, 0xb8, 0x9d, 0xff, 0x03, 0x00, 0x00,
	0xff, 0xff, 0xf7, 0x49, 0x8a, 0x4a, 0xce, 0x0e, 0x00, 0x00,
}