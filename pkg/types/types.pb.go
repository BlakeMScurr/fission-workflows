// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/types/types.proto

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	pkg/types/types.proto

It has these top-level messages:
	Workflow
	WorkflowSpec
	WorkflowStatus
	WorkflowInvocation
	WorkflowInvocationSpec
	WorkflowInvocationStatus
	DependencyConfig
	Task
	TaskSpec
	TaskStatus
	TaskDependencyParameters
	TaskInvocation
	TaskInvocationSpec
	TaskInvocationStatus
	ObjectMetadata
	TypedValue
	Error
	FnRef
	TypedValueMap
	TypedValueList
*/
package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type WorkflowStatus_Status int32

const (
	WorkflowStatus_PENDING WorkflowStatus_Status = 0
	//        PARSING = 1; // During validation/parsing
	WorkflowStatus_READY   WorkflowStatus_Status = 2
	WorkflowStatus_FAILED  WorkflowStatus_Status = 3
	WorkflowStatus_DELETED WorkflowStatus_Status = 4
)

var WorkflowStatus_Status_name = map[int32]string{
	0: "PENDING",
	2: "READY",
	3: "FAILED",
	4: "DELETED",
}
var WorkflowStatus_Status_value = map[string]int32{
	"PENDING": 0,
	"READY":   2,
	"FAILED":  3,
	"DELETED": 4,
}

func (x WorkflowStatus_Status) String() string {
	return proto.EnumName(WorkflowStatus_Status_name, int32(x))
}
func (WorkflowStatus_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type WorkflowInvocationStatus_Status int32

const (
	WorkflowInvocationStatus_UNKNOWN     WorkflowInvocationStatus_Status = 0
	WorkflowInvocationStatus_SCHEDULED   WorkflowInvocationStatus_Status = 1
	WorkflowInvocationStatus_IN_PROGRESS WorkflowInvocationStatus_Status = 2
	WorkflowInvocationStatus_SUCCEEDED   WorkflowInvocationStatus_Status = 3
	WorkflowInvocationStatus_FAILED      WorkflowInvocationStatus_Status = 4
	WorkflowInvocationStatus_ABORTED     WorkflowInvocationStatus_Status = 5
)

var WorkflowInvocationStatus_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "SCHEDULED",
	2: "IN_PROGRESS",
	3: "SUCCEEDED",
	4: "FAILED",
	5: "ABORTED",
}
var WorkflowInvocationStatus_Status_value = map[string]int32{
	"UNKNOWN":     0,
	"SCHEDULED":   1,
	"IN_PROGRESS": 2,
	"SUCCEEDED":   3,
	"FAILED":      4,
	"ABORTED":     5,
}

func (x WorkflowInvocationStatus_Status) String() string {
	return proto.EnumName(WorkflowInvocationStatus_Status_name, int32(x))
}
func (WorkflowInvocationStatus_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5, 0}
}

type TaskStatus_Status int32

const (
	TaskStatus_STARTED TaskStatus_Status = 0
	TaskStatus_READY   TaskStatus_Status = 1
	TaskStatus_FAILED  TaskStatus_Status = 2
)

var TaskStatus_Status_name = map[int32]string{
	0: "STARTED",
	1: "READY",
	2: "FAILED",
}
var TaskStatus_Status_value = map[string]int32{
	"STARTED": 0,
	"READY":   1,
	"FAILED":  2,
}

func (x TaskStatus_Status) String() string {
	return proto.EnumName(TaskStatus_Status_name, int32(x))
}
func (TaskStatus_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{9, 0} }

type TaskDependencyParameters_DependencyType int32

const (
	TaskDependencyParameters_DATA           TaskDependencyParameters_DependencyType = 0
	TaskDependencyParameters_CONTROL        TaskDependencyParameters_DependencyType = 1
	TaskDependencyParameters_DYNAMIC_OUTPUT TaskDependencyParameters_DependencyType = 2
)

var TaskDependencyParameters_DependencyType_name = map[int32]string{
	0: "DATA",
	1: "CONTROL",
	2: "DYNAMIC_OUTPUT",
}
var TaskDependencyParameters_DependencyType_value = map[string]int32{
	"DATA":           0,
	"CONTROL":        1,
	"DYNAMIC_OUTPUT": 2,
}

func (x TaskDependencyParameters_DependencyType) String() string {
	return proto.EnumName(TaskDependencyParameters_DependencyType_name, int32(x))
}
func (TaskDependencyParameters_DependencyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{10, 0}
}

type TaskInvocationStatus_Status int32

const (
	TaskInvocationStatus_UNKNOWN     TaskInvocationStatus_Status = 0
	TaskInvocationStatus_SCHEDULED   TaskInvocationStatus_Status = 1
	TaskInvocationStatus_IN_PROGRESS TaskInvocationStatus_Status = 2
	TaskInvocationStatus_SUCCEEDED   TaskInvocationStatus_Status = 3
	TaskInvocationStatus_FAILED      TaskInvocationStatus_Status = 4
	TaskInvocationStatus_ABORTED     TaskInvocationStatus_Status = 5
	TaskInvocationStatus_SKIPPED     TaskInvocationStatus_Status = 6
)

var TaskInvocationStatus_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "SCHEDULED",
	2: "IN_PROGRESS",
	3: "SUCCEEDED",
	4: "FAILED",
	5: "ABORTED",
	6: "SKIPPED",
}
var TaskInvocationStatus_Status_value = map[string]int32{
	"UNKNOWN":     0,
	"SCHEDULED":   1,
	"IN_PROGRESS": 2,
	"SUCCEEDED":   3,
	"FAILED":      4,
	"ABORTED":     5,
	"SKIPPED":     6,
}

func (x TaskInvocationStatus_Status) String() string {
	return proto.EnumName(TaskInvocationStatus_Status_name, int32(x))
}
func (TaskInvocationStatus_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{13, 0}
}

//
// Workflow Model
//
type Workflow struct {
	Metadata *ObjectMetadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Spec     *WorkflowSpec   `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	Status   *WorkflowStatus `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *Workflow) Reset()                    { *m = Workflow{} }
func (m *Workflow) String() string            { return proto.CompactTextString(m) }
func (*Workflow) ProtoMessage()               {}
func (*Workflow) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Workflow) GetMetadata() *ObjectMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Workflow) GetSpec() *WorkflowSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Workflow) GetStatus() *WorkflowStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// Workflow Definition
//
// The workflowDefinition contains the definition of a workflow.
//
// Ideally the source code (json, yaml) can be converted directly to this message.
// Naming, triggers and versioning of the workflow itself is out of the scope of this data structure, which is delegated
// to the user/system upon the creation of a workflow.
type WorkflowSpec struct {
	// apiVersion describes what version is of the workflow definition.
	// By default the workflow engine will assume the latest version to be used.
	ApiVersion string `protobuf:"bytes,1,opt,name=apiVersion" json:"apiVersion,omitempty"`
	// Tasks contains the specs of the tasks, with the key being the task id.
	//
	// Note: Dependency graph is build into the tasks.
	Tasks map[string]*TaskSpec `protobuf:"bytes,2,rep,name=tasks" json:"tasks,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// From which task should the workflow return the output? Future: multiple? Implicit?
	OutputTask  string `protobuf:"bytes,3,opt,name=outputTask" json:"outputTask,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	// The UID that the workflow should have. Only use this in case you want to force a specific UID.
	ForceId string `protobuf:"bytes,5,opt,name=forceId" json:"forceId,omitempty"`
	// Name is solely for human-readablity
	Name string `protobuf:"bytes,6,opt,name=name" json:"name,omitempty"`
	// Internal indicates whether is a workflow should be visible to a human (default) or not.
	//
	Internal bool `protobuf:"varint,7,opt,name=internal" json:"internal,omitempty"`
}

func (m *WorkflowSpec) Reset()                    { *m = WorkflowSpec{} }
func (m *WorkflowSpec) String() string            { return proto.CompactTextString(m) }
func (*WorkflowSpec) ProtoMessage()               {}
func (*WorkflowSpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *WorkflowSpec) GetApiVersion() string {
	if m != nil {
		return m.ApiVersion
	}
	return ""
}

func (m *WorkflowSpec) GetTasks() map[string]*TaskSpec {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *WorkflowSpec) GetOutputTask() string {
	if m != nil {
		return m.OutputTask
	}
	return ""
}

func (m *WorkflowSpec) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *WorkflowSpec) GetForceId() string {
	if m != nil {
		return m.ForceId
	}
	return ""
}

func (m *WorkflowSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WorkflowSpec) GetInternal() bool {
	if m != nil {
		return m.Internal
	}
	return false
}

type WorkflowStatus struct {
	Status    WorkflowStatus_Status      `protobuf:"varint,1,opt,name=status,enum=WorkflowStatus_Status" json:"status,omitempty"`
	UpdatedAt *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=updatedAt" json:"updatedAt,omitempty"`
	// Tasks contains the status of the tasks, with the key being the task id.
	Tasks map[string]*TaskStatus `protobuf:"bytes,3,rep,name=tasks" json:"tasks,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Error *Error                 `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
}

func (m *WorkflowStatus) Reset()                    { *m = WorkflowStatus{} }
func (m *WorkflowStatus) String() string            { return proto.CompactTextString(m) }
func (*WorkflowStatus) ProtoMessage()               {}
func (*WorkflowStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *WorkflowStatus) GetStatus() WorkflowStatus_Status {
	if m != nil {
		return m.Status
	}
	return WorkflowStatus_PENDING
}

func (m *WorkflowStatus) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *WorkflowStatus) GetTasks() map[string]*TaskStatus {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *WorkflowStatus) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

//
// Workflow Invocation Model
//
type WorkflowInvocation struct {
	Metadata *ObjectMetadata           `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Spec     *WorkflowInvocationSpec   `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	Status   *WorkflowInvocationStatus `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *WorkflowInvocation) Reset()                    { *m = WorkflowInvocation{} }
func (m *WorkflowInvocation) String() string            { return proto.CompactTextString(m) }
func (*WorkflowInvocation) ProtoMessage()               {}
func (*WorkflowInvocation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *WorkflowInvocation) GetMetadata() *ObjectMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *WorkflowInvocation) GetSpec() *WorkflowInvocationSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *WorkflowInvocation) GetStatus() *WorkflowInvocationStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// Workflow Invocation Model
type WorkflowInvocationSpec struct {
	WorkflowId string                 `protobuf:"bytes,1,opt,name=workflowId" json:"workflowId,omitempty"`
	Inputs     map[string]*TypedValue `protobuf:"bytes,2,rep,name=inputs" json:"inputs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// ParentId contains the id of the encapsulating workflow invocation.
	//
	// This used within the workflow engine; for user-provided workflow invocations the parentId is ignored.
	ParentId string `protobuf:"bytes,3,opt,name=parentId" json:"parentId,omitempty"`
}

func (m *WorkflowInvocationSpec) Reset()                    { *m = WorkflowInvocationSpec{} }
func (m *WorkflowInvocationSpec) String() string            { return proto.CompactTextString(m) }
func (*WorkflowInvocationSpec) ProtoMessage()               {}
func (*WorkflowInvocationSpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *WorkflowInvocationSpec) GetWorkflowId() string {
	if m != nil {
		return m.WorkflowId
	}
	return ""
}

func (m *WorkflowInvocationSpec) GetInputs() map[string]*TypedValue {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *WorkflowInvocationSpec) GetParentId() string {
	if m != nil {
		return m.ParentId
	}
	return ""
}

type WorkflowInvocationStatus struct {
	Status    WorkflowInvocationStatus_Status `protobuf:"varint,1,opt,name=status,enum=WorkflowInvocationStatus_Status" json:"status,omitempty"`
	UpdatedAt *google_protobuf.Timestamp      `protobuf:"bytes,2,opt,name=updatedAt" json:"updatedAt,omitempty"`
	Tasks     map[string]*TaskInvocation      `protobuf:"bytes,3,rep,name=tasks" json:"tasks,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Output    *TypedValue                     `protobuf:"bytes,4,opt,name=output" json:"output,omitempty"`
	// In case the task ID also exists in the workflow spec, the dynamic task will be
	// used as an overlay over the static task.
	DynamicTasks map[string]*Task `protobuf:"bytes,5,rep,name=dynamicTasks" json:"dynamicTasks,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Error        *Error           `protobuf:"bytes,6,opt,name=error" json:"error,omitempty"`
}

func (m *WorkflowInvocationStatus) Reset()                    { *m = WorkflowInvocationStatus{} }
func (m *WorkflowInvocationStatus) String() string            { return proto.CompactTextString(m) }
func (*WorkflowInvocationStatus) ProtoMessage()               {}
func (*WorkflowInvocationStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *WorkflowInvocationStatus) GetStatus() WorkflowInvocationStatus_Status {
	if m != nil {
		return m.Status
	}
	return WorkflowInvocationStatus_UNKNOWN
}

func (m *WorkflowInvocationStatus) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *WorkflowInvocationStatus) GetTasks() map[string]*TaskInvocation {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *WorkflowInvocationStatus) GetOutput() *TypedValue {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *WorkflowInvocationStatus) GetDynamicTasks() map[string]*Task {
	if m != nil {
		return m.DynamicTasks
	}
	return nil
}

func (m *WorkflowInvocationStatus) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type DependencyConfig struct {
	// Dependencies for this task to execute
	Requires map[string]*TaskDependencyParameters `protobuf:"bytes,1,rep,name=requires" json:"requires,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Number of dependencies to wait for
	Await int32 `protobuf:"varint,2,opt,name=await" json:"await,omitempty"`
}

func (m *DependencyConfig) Reset()                    { *m = DependencyConfig{} }
func (m *DependencyConfig) String() string            { return proto.CompactTextString(m) }
func (*DependencyConfig) ProtoMessage()               {}
func (*DependencyConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DependencyConfig) GetRequires() map[string]*TaskDependencyParameters {
	if m != nil {
		return m.Requires
	}
	return nil
}

func (m *DependencyConfig) GetAwait() int32 {
	if m != nil {
		return m.Await
	}
	return 0
}

//
// Task Model
//
type Task struct {
	Metadata *ObjectMetadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Spec     *TaskSpec       `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	Status   *TaskStatus     `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Task) GetMetadata() *ObjectMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Task) GetSpec() *TaskSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Task) GetStatus() *TaskStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// A task is the primitive unit of a workflow, representing an action that needs to be performed in order to continue.
//
// A task as a number of inputs and exactly two outputs
// Id is specified outside of TaskSpec
type TaskSpec struct {
	// Name/identifier of the function
	FunctionRef string                 `protobuf:"bytes,1,opt,name=functionRef" json:"functionRef,omitempty"`
	Inputs      map[string]*TypedValue `protobuf:"bytes,2,rep,name=inputs" json:"inputs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Dependencies for this task to execute
	Requires map[string]*TaskDependencyParameters `protobuf:"bytes,3,rep,name=requires" json:"requires,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Number of dependencies to wait for
	Await int32 `protobuf:"varint,4,opt,name=await" json:"await,omitempty"`
	// Transform the output, or override the output with a literal
	Output *TypedValue `protobuf:"bytes,5,opt,name=output" json:"output,omitempty"`
}

func (m *TaskSpec) Reset()                    { *m = TaskSpec{} }
func (m *TaskSpec) String() string            { return proto.CompactTextString(m) }
func (*TaskSpec) ProtoMessage()               {}
func (*TaskSpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *TaskSpec) GetFunctionRef() string {
	if m != nil {
		return m.FunctionRef
	}
	return ""
}

func (m *TaskSpec) GetInputs() map[string]*TypedValue {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *TaskSpec) GetRequires() map[string]*TaskDependencyParameters {
	if m != nil {
		return m.Requires
	}
	return nil
}

func (m *TaskSpec) GetAwait() int32 {
	if m != nil {
		return m.Await
	}
	return 0
}

func (m *TaskSpec) GetOutput() *TypedValue {
	if m != nil {
		return m.Output
	}
	return nil
}

type TaskStatus struct {
	Status    TaskStatus_Status          `protobuf:"varint,1,opt,name=status,enum=TaskStatus_Status" json:"status,omitempty"`
	UpdatedAt *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=updatedAt" json:"updatedAt,omitempty"`
	FnRef     *FnRef                     `protobuf:"bytes,3,opt,name=fnRef" json:"fnRef,omitempty"`
	Error     *Error                     `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
}

func (m *TaskStatus) Reset()                    { *m = TaskStatus{} }
func (m *TaskStatus) String() string            { return proto.CompactTextString(m) }
func (*TaskStatus) ProtoMessage()               {}
func (*TaskStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *TaskStatus) GetStatus() TaskStatus_Status {
	if m != nil {
		return m.Status
	}
	return TaskStatus_STARTED
}

func (m *TaskStatus) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *TaskStatus) GetFnRef() *FnRef {
	if m != nil {
		return m.FnRef
	}
	return nil
}

func (m *TaskStatus) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type TaskDependencyParameters struct {
	Type  TaskDependencyParameters_DependencyType `protobuf:"varint,1,opt,name=type,enum=TaskDependencyParameters_DependencyType" json:"type,omitempty"`
	Alias string                                  `protobuf:"bytes,2,opt,name=alias" json:"alias,omitempty"`
}

func (m *TaskDependencyParameters) Reset()                    { *m = TaskDependencyParameters{} }
func (m *TaskDependencyParameters) String() string            { return proto.CompactTextString(m) }
func (*TaskDependencyParameters) ProtoMessage()               {}
func (*TaskDependencyParameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *TaskDependencyParameters) GetType() TaskDependencyParameters_DependencyType {
	if m != nil {
		return m.Type
	}
	return TaskDependencyParameters_DATA
}

func (m *TaskDependencyParameters) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

//
// Task Invocation Model
//
type TaskInvocation struct {
	Metadata *ObjectMetadata       `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Spec     *TaskInvocationSpec   `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	Status   *TaskInvocationStatus `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *TaskInvocation) Reset()                    { *m = TaskInvocation{} }
func (m *TaskInvocation) String() string            { return proto.CompactTextString(m) }
func (*TaskInvocation) ProtoMessage()               {}
func (*TaskInvocation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *TaskInvocation) GetMetadata() *ObjectMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *TaskInvocation) GetSpec() *TaskInvocationSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *TaskInvocation) GetStatus() *TaskInvocationStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type TaskInvocationSpec struct {
	// Id of the task to be invoked (no ambiguatity at this point
	FnRef *FnRef `protobuf:"bytes,1,opt,name=fnRef" json:"fnRef,omitempty"`
	// TaskId is the id of the task within the workflow
	TaskId string `protobuf:"bytes,2,opt,name=taskId" json:"taskId,omitempty"`
	// Inputs contain all inputs to the task invocation
	Inputs map[string]*TypedValue `protobuf:"bytes,3,rep,name=inputs" json:"inputs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	//
	InvocationId string `protobuf:"bytes,4,opt,name=invocationId" json:"invocationId,omitempty"`
}

func (m *TaskInvocationSpec) Reset()                    { *m = TaskInvocationSpec{} }
func (m *TaskInvocationSpec) String() string            { return proto.CompactTextString(m) }
func (*TaskInvocationSpec) ProtoMessage()               {}
func (*TaskInvocationSpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *TaskInvocationSpec) GetFnRef() *FnRef {
	if m != nil {
		return m.FnRef
	}
	return nil
}

func (m *TaskInvocationSpec) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *TaskInvocationSpec) GetInputs() map[string]*TypedValue {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *TaskInvocationSpec) GetInvocationId() string {
	if m != nil {
		return m.InvocationId
	}
	return ""
}

type TaskInvocationStatus struct {
	Status    TaskInvocationStatus_Status `protobuf:"varint,1,opt,name=status,enum=TaskInvocationStatus_Status" json:"status,omitempty"`
	UpdatedAt *google_protobuf.Timestamp  `protobuf:"bytes,2,opt,name=updatedAt" json:"updatedAt,omitempty"`
	Output    *TypedValue                 `protobuf:"bytes,3,opt,name=output" json:"output,omitempty"`
	Error     *Error                      `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
}

func (m *TaskInvocationStatus) Reset()                    { *m = TaskInvocationStatus{} }
func (m *TaskInvocationStatus) String() string            { return proto.CompactTextString(m) }
func (*TaskInvocationStatus) ProtoMessage()               {}
func (*TaskInvocationStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *TaskInvocationStatus) GetStatus() TaskInvocationStatus_Status {
	if m != nil {
		return m.Status
	}
	return TaskInvocationStatus_UNKNOWN
}

func (m *TaskInvocationStatus) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *TaskInvocationStatus) GetOutput() *TypedValue {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *TaskInvocationStatus) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

//
// Common
//
type ObjectMetadata struct {
	Id        string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=createdAt" json:"createdAt,omitempty"`
}

func (m *ObjectMetadata) Reset()                    { *m = ObjectMetadata{} }
func (m *ObjectMetadata) String() string            { return proto.CompactTextString(m) }
func (*ObjectMetadata) ProtoMessage()               {}
func (*ObjectMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *ObjectMetadata) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ObjectMetadata) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

// Copy of protobuf's Any, to avoid protobuf requirement of a protobuf-based type.
type TypedValue struct {
	Type   string            `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Value  []byte            `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *TypedValue) Reset()                    { *m = TypedValue{} }
func (m *TypedValue) String() string            { return proto.CompactTextString(m) }
func (*TypedValue) ProtoMessage()               {}
func (*TypedValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *TypedValue) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *TypedValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *TypedValue) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type Error struct {
	Code    string `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *Error) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// FnRef is an immutable, unique reference to a function on a specific function runtime environment.
//
// The string representation (via String or Format): runtime://runtimeId
type FnRef struct {
	// Runtime is the Function Runtime environment (fnenv) that was used to resolve the function.
	Runtime string `protobuf:"bytes,2,opt,name=runtime" json:"runtime,omitempty"`
	// ID is the runtime-specific identifier of the function.
	ID string `protobuf:"bytes,3,opt,name=ID" json:"ID,omitempty"`
}

func (m *FnRef) Reset()                    { *m = FnRef{} }
func (m *FnRef) String() string            { return proto.CompactTextString(m) }
func (*FnRef) ProtoMessage()               {}
func (*FnRef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *FnRef) GetRuntime() string {
	if m != nil {
		return m.Runtime
	}
	return ""
}

func (m *FnRef) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

// Utility wrapper for a TypedValue map
type TypedValueMap struct {
	Value map[string]*TypedValue `protobuf:"bytes,1,rep,name=Value" json:"Value,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *TypedValueMap) Reset()                    { *m = TypedValueMap{} }
func (m *TypedValueMap) String() string            { return proto.CompactTextString(m) }
func (*TypedValueMap) ProtoMessage()               {}
func (*TypedValueMap) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *TypedValueMap) GetValue() map[string]*TypedValue {
	if m != nil {
		return m.Value
	}
	return nil
}

// Utility wrapper for a TypedValue list
type TypedValueList struct {
	Value []*TypedValue `protobuf:"bytes,1,rep,name=Value" json:"Value,omitempty"`
}

func (m *TypedValueList) Reset()                    { *m = TypedValueList{} }
func (m *TypedValueList) String() string            { return proto.CompactTextString(m) }
func (*TypedValueList) ProtoMessage()               {}
func (*TypedValueList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *TypedValueList) GetValue() []*TypedValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Workflow)(nil), "Workflow")
	proto.RegisterType((*WorkflowSpec)(nil), "WorkflowSpec")
	proto.RegisterType((*WorkflowStatus)(nil), "WorkflowStatus")
	proto.RegisterType((*WorkflowInvocation)(nil), "WorkflowInvocation")
	proto.RegisterType((*WorkflowInvocationSpec)(nil), "WorkflowInvocationSpec")
	proto.RegisterType((*WorkflowInvocationStatus)(nil), "WorkflowInvocationStatus")
	proto.RegisterType((*DependencyConfig)(nil), "DependencyConfig")
	proto.RegisterType((*Task)(nil), "Task")
	proto.RegisterType((*TaskSpec)(nil), "TaskSpec")
	proto.RegisterType((*TaskStatus)(nil), "TaskStatus")
	proto.RegisterType((*TaskDependencyParameters)(nil), "TaskDependencyParameters")
	proto.RegisterType((*TaskInvocation)(nil), "TaskInvocation")
	proto.RegisterType((*TaskInvocationSpec)(nil), "TaskInvocationSpec")
	proto.RegisterType((*TaskInvocationStatus)(nil), "TaskInvocationStatus")
	proto.RegisterType((*ObjectMetadata)(nil), "ObjectMetadata")
	proto.RegisterType((*TypedValue)(nil), "TypedValue")
	proto.RegisterType((*Error)(nil), "Error")
	proto.RegisterType((*FnRef)(nil), "FnRef")
	proto.RegisterType((*TypedValueMap)(nil), "TypedValueMap")
	proto.RegisterType((*TypedValueList)(nil), "TypedValueList")
	proto.RegisterEnum("WorkflowStatus_Status", WorkflowStatus_Status_name, WorkflowStatus_Status_value)
	proto.RegisterEnum("WorkflowInvocationStatus_Status", WorkflowInvocationStatus_Status_name, WorkflowInvocationStatus_Status_value)
	proto.RegisterEnum("TaskStatus_Status", TaskStatus_Status_name, TaskStatus_Status_value)
	proto.RegisterEnum("TaskDependencyParameters_DependencyType", TaskDependencyParameters_DependencyType_name, TaskDependencyParameters_DependencyType_value)
	proto.RegisterEnum("TaskInvocationStatus_Status", TaskInvocationStatus_Status_name, TaskInvocationStatus_Status_value)
}

func init() { proto.RegisterFile("pkg/types/types.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x58, 0xcd, 0x6e, 0xdb, 0xc6,
	0x13, 0x0f, 0x29, 0x51, 0x96, 0x46, 0xb6, 0xa2, 0xff, 0xfe, 0xf3, 0xc1, 0xa8, 0x69, 0xe3, 0x30,
	0x2d, 0x62, 0x34, 0x0d, 0xdd, 0x38, 0x2d, 0x9a, 0x8f, 0x5e, 0x14, 0x91, 0x4e, 0x89, 0x38, 0x92,
	0x41, 0xcb, 0x09, 0x12, 0xa0, 0x08, 0x68, 0x71, 0x65, 0x30, 0x96, 0x48, 0x96, 0xa4, 0x62, 0xf8,
	0xdc, 0x43, 0x9f, 0xa0, 0x0f, 0xd0, 0x5e, 0x7b, 0xe8, 0xa9, 0xc7, 0xbe, 0x41, 0x0f, 0xbd, 0x16,
	0xe8, 0x1b, 0xf4, 0xd0, 0x57, 0x28, 0x76, 0xb9, 0x24, 0x77, 0x69, 0x29, 0xae, 0x0b, 0xe7, 0x62,
	0xef, 0xce, 0xce, 0xce, 0xec, 0xfc, 0x66, 0x7e, 0xb3, 0x4b, 0xc1, 0xc5, 0xf0, 0x60, 0x7f, 0x3d,
	0x39, 0x0a, 0x71, 0x9c, 0xfe, 0xd5, 0xc3, 0x28, 0x48, 0x82, 0xce, 0xb5, 0xfd, 0x20, 0xd8, 0x9f,
	0xe0, 0x75, 0x3a, 0xdb, 0x9b, 0x8d, 0xd7, 0x13, 0x6f, 0x8a, 0xe3, 0xc4, 0x99, 0x86, 0xa9, 0x82,
	0xf6, 0xad, 0x04, 0xf5, 0xe7, 0x41, 0x74, 0x30, 0x9e, 0x04, 0x87, 0xe8, 0x16, 0xd4, 0xa7, 0x38,
	0x71, 0x5c, 0x27, 0x71, 0x54, 0x69, 0x55, 0x5a, 0x6b, 0x6e, 0x9c, 0xd7, 0x07, 0x7b, 0xaf, 0xf1,
	0x28, 0x79, 0xca, 0xc4, 0x76, 0xae, 0x80, 0xae, 0x43, 0x35, 0x0e, 0xf1, 0x48, 0x95, 0xa9, 0xe2,
	0x8a, 0x9e, 0x59, 0xd9, 0x09, 0xf1, 0xc8, 0xa6, 0x4b, 0xe8, 0x26, 0xd4, 0xe2, 0xc4, 0x49, 0x66,
	0xb1, 0x5a, 0x61, 0xd6, 0x72, 0x25, 0x2a, 0xb6, 0xd9, 0xb2, 0xf6, 0xb3, 0x0c, 0xcb, 0xfc, 0x7e,
	0xf4, 0x01, 0x80, 0x13, 0x7a, 0xcf, 0x70, 0x14, 0x7b, 0x81, 0x4f, 0xcf, 0xd2, 0xb0, 0x39, 0x09,
	0xd2, 0x41, 0x49, 0x9c, 0xf8, 0x20, 0x56, 0xe5, 0xd5, 0xca, 0x5a, 0x73, 0x43, 0x15, 0xbc, 0xeb,
	0x43, 0xb2, 0x64, 0xfa, 0x49, 0x74, 0x64, 0xa7, 0x6a, 0xc4, 0x5e, 0x30, 0x4b, 0xc2, 0x59, 0x42,
	0x96, 0xe8, 0x69, 0x1a, 0x36, 0x27, 0x41, 0xab, 0xd0, 0x74, 0x71, 0x3c, 0x8a, 0xbc, 0x30, 0x21,
	0x0e, 0xab, 0x54, 0x81, 0x17, 0x21, 0x15, 0x96, 0xc6, 0x41, 0x34, 0xc2, 0x96, 0xab, 0x2a, 0x74,
	0x35, 0x9b, 0x22, 0x04, 0x55, 0xdf, 0x99, 0x62, 0xb5, 0x46, 0xc5, 0x74, 0x8c, 0x3a, 0x50, 0xf7,
	0xfc, 0x04, 0x47, 0xbe, 0x33, 0x51, 0x97, 0x56, 0xa5, 0xb5, 0xba, 0x9d, 0xcf, 0x3b, 0x3d, 0x80,
	0xe2, 0x80, 0xa8, 0x0d, 0x95, 0x03, 0x7c, 0xc4, 0x42, 0x24, 0x43, 0x74, 0x0d, 0x94, 0x37, 0xce,
	0x64, 0x86, 0x19, 0xb2, 0x0d, 0x1a, 0x0e, 0x45, 0x35, 0x95, 0x3f, 0x90, 0xef, 0x49, 0xda, 0x6f,
	0x32, 0xb4, 0x44, 0x30, 0x91, 0x9e, 0xa3, 0x4d, 0x8c, 0xb5, 0x36, 0x2e, 0x95, 0xd0, 0xd6, 0x45,
	0xd0, 0xd1, 0x3d, 0x68, 0xcc, 0x42, 0xd7, 0x49, 0xb0, 0xdb, 0x4d, 0x98, 0xaf, 0x8e, 0x9e, 0xd6,
	0x8b, 0x9e, 0xd5, 0x8b, 0x3e, 0xcc, 0xea, 0xc5, 0x2e, 0x94, 0xd1, 0xa7, 0x19, 0xfa, 0x15, 0x8a,
	0x7e, 0xa7, 0xec, 0xe8, 0x38, 0xfe, 0x57, 0x41, 0xc1, 0x51, 0x14, 0x44, 0x14, 0xd9, 0xe6, 0x46,
	0x4d, 0x37, 0xc9, 0xcc, 0x4e, 0x85, 0x1d, 0xf3, 0x04, 0x44, 0xae, 0x8b, 0x88, 0x34, 0x53, 0x44,
	0xd2, 0x68, 0x38, 0x4c, 0xee, 0x43, 0x8d, 0x41, 0xd1, 0x84, 0xa5, 0x6d, 0xb3, 0x6f, 0x58, 0xfd,
	0xc7, 0xed, 0x73, 0xa8, 0x01, 0x8a, 0x6d, 0x76, 0x8d, 0x17, 0x6d, 0x19, 0x01, 0xd4, 0x36, 0xbb,
	0xd6, 0x96, 0x69, 0xb4, 0x2b, 0x44, 0xc7, 0x30, 0xb7, 0xcc, 0xa1, 0x69, 0xb4, 0xab, 0xda, 0x0f,
	0x12, 0xa0, 0x2c, 0x08, 0xcb, 0x7f, 0x13, 0x8c, 0x1c, 0x9a, 0xf4, 0x53, 0x11, 0xe2, 0x96, 0x40,
	0x88, 0xcb, 0xfa, 0x71, 0x7b, 0x1c, 0x35, 0xee, 0x94, 0xa8, 0x71, 0x65, 0x9e, 0xba, 0x48, 0x92,
	0x3f, 0x24, 0xb8, 0x34, 0xdf, 0x26, 0x29, 0xef, 0xc3, 0x6c, 0xc5, 0xcd, 0xe8, 0x52, 0x48, 0xd0,
	0x43, 0xa8, 0x79, 0x7e, 0x38, 0x4b, 0x32, 0xbe, 0xdc, 0x58, 0x70, 0x38, 0xdd, 0xa2, 0x5a, 0x69,
	0xea, 0xd8, 0x16, 0x52, 0xcb, 0xa1, 0x13, 0x61, 0x3f, 0xb1, 0x5c, 0xc6, 0x9c, 0x7c, 0xde, 0xd9,
	0x84, 0x26, 0xb7, 0xe5, 0x5f, 0xa5, 0xee, 0x28, 0xc4, 0xee, 0x33, 0x22, 0xe2, 0x53, 0xf7, 0x7b,
	0x15, 0xd4, 0x45, 0x00, 0xa0, 0x7b, 0xa5, 0xc2, 0x5e, 0x5d, 0x88, 0xd5, 0xd9, 0x95, 0xf8, 0x03,
	0xb1, 0xc4, 0x3f, 0x5c, 0xec, 0xf2, 0x78, 0xb1, 0xdf, 0x80, 0x5a, 0xda, 0x5a, 0x58, 0xb5, 0x0b,
	0x41, 0xb3, 0x25, 0x34, 0x80, 0x65, 0xf7, 0xc8, 0x77, 0xa6, 0xde, 0x88, 0x1a, 0x50, 0x15, 0xea,
	0xe7, 0xd6, 0x62, 0x3f, 0x06, 0xa7, 0x9d, 0xba, 0x13, 0x0c, 0x14, 0x14, 0xab, 0xcd, 0xa3, 0x98,
	0x75, 0x02, 0xc5, 0x3e, 0x12, 0xf3, 0x74, 0x9e, 0x86, 0x55, 0x9c, 0x81, 0xcb, 0x55, 0x67, 0x13,
	0xfe, 0x77, 0xec, 0x2c, 0x73, 0x2c, 0xbe, 0x27, 0x5a, 0x54, 0xa8, 0x45, 0x3e, 0xe7, 0x5f, 0xf3,
	0x74, 0xdd, 0xed, 0x3f, 0xe9, 0x0f, 0x9e, 0xf7, 0xdb, 0xe7, 0xd0, 0x0a, 0x34, 0x76, 0x7a, 0x5f,
	0x99, 0xc6, 0x2e, 0xa1, 0xa9, 0x84, 0xce, 0x43, 0xd3, 0xea, 0xbf, 0xda, 0xb6, 0x07, 0x8f, 0x6d,
	0x73, 0x67, 0xa7, 0x2d, 0xd3, 0xf5, 0xdd, 0x5e, 0xcf, 0x34, 0x0d, 0x4a, 0xe3, 0x82, 0xd2, 0x55,
	0x62, 0xa7, 0xfb, 0x68, 0x60, 0x13, 0x4a, 0x2b, 0xda, 0xaf, 0x12, 0xb4, 0x0d, 0x1c, 0x62, 0xdf,
	0xc5, 0xfe, 0xe8, 0xa8, 0x17, 0xf8, 0x63, 0x6f, 0x1f, 0x3d, 0x84, 0x7a, 0x84, 0xbf, 0x99, 0x79,
	0x11, 0x26, 0xc5, 0x44, 0x10, 0xbf, 0xa6, 0x97, 0x95, 0x74, 0x9b, 0x69, 0xa4, 0x28, 0xe7, 0x1b,
	0xd0, 0x05, 0x50, 0x9c, 0x43, 0xc7, 0x4b, 0x2b, 0x49, 0xb1, 0xd3, 0x49, 0xe7, 0x19, 0xac, 0x08,
	0x1b, 0xe6, 0x40, 0xb1, 0x2e, 0x42, 0x71, 0x85, 0x42, 0x51, 0xb8, 0xdd, 0x76, 0x22, 0x67, 0x8a,
	0x13, 0x1c, 0x09, 0xdd, 0xec, 0x10, 0xaa, 0xf4, 0x6a, 0x3a, 0x55, 0x0f, 0x7a, 0x5f, 0xe8, 0x41,
	0xdc, 0xd5, 0x91, 0x76, 0x9d, 0x1b, 0xa5, 0xae, 0x23, 0x74, 0xd2, 0xac, 0xcf, 0xfc, 0x25, 0x43,
	0x3d, 0xdb, 0x47, 0x2e, 0xc6, 0xf1, 0xcc, 0x1f, 0xd1, 0x1a, 0xc0, 0x63, 0x16, 0x14, 0x2f, 0x42,
	0xb7, 0x4b, 0xbd, 0xe5, 0x62, 0xee, 0x74, 0x6e, 0x37, 0xb9, 0xcb, 0x65, 0x20, 0xe5, 0xd6, 0xe5,
	0x62, 0xc3, 0x89, 0xc8, 0x57, 0x39, 0xe4, 0x39, 0x9e, 0x29, 0x0b, 0x79, 0x76, 0x56, 0x1d, 0xea,
	0x9d, 0xa5, 0xf9, 0x4f, 0x29, 0x65, 0x26, 0xa3, 0xc2, 0xc7, 0xa5, 0x5e, 0x87, 0xb8, 0x0c, 0x9d,
	0x5d, 0x77, 0xbb, 0x0a, 0xca, 0x98, 0xe6, 0xb3, 0xc2, 0x7a, 0xc5, 0x26, 0x99, 0xd9, 0xa9, 0xf0,
	0xed, 0x97, 0xb5, 0xf6, 0x09, 0x4f, 0xdb, 0x9d, 0x61, 0x97, 0xd2, 0x8d, 0xbb, 0x65, 0x25, 0x8e,
	0x92, 0xb2, 0xf6, 0x93, 0x04, 0xea, 0x22, 0x18, 0xd0, 0x97, 0x50, 0x25, 0x8f, 0x55, 0x16, 0xea,
	0xda, 0x42, 0xbc, 0x38, 0x8a, 0x92, 0xe4, 0xd8, 0x74, 0x17, 0x2d, 0x8a, 0x89, 0xe7, 0xc4, 0x34,
	0xf4, 0x86, 0x9d, 0x4e, 0xb4, 0x87, 0xd0, 0x12, 0xb5, 0x51, 0x1d, 0xaa, 0x46, 0x77, 0xd8, 0x6d,
	0x9f, 0x23, 0x07, 0xee, 0x0d, 0xfa, 0x43, 0x7b, 0xb0, 0xd5, 0x96, 0x10, 0x82, 0x96, 0xf1, 0xa2,
	0xdf, 0x7d, 0x6a, 0xf5, 0x5e, 0x0d, 0x76, 0x87, 0xdb, 0xbb, 0xc3, 0xb6, 0xac, 0x7d, 0x2f, 0x41,
	0x4b, 0x6c, 0x7c, 0xa7, 0xa3, 0xdf, 0x4d, 0x81, 0x7e, 0xff, 0x2f, 0x35, 0x51, 0x8e, 0x88, 0xb7,
	0x4b, 0x44, 0xbc, 0x58, 0x56, 0x15, 0x29, 0xf9, 0xb7, 0x04, 0xe8, 0xb8, 0xad, 0x22, 0x8d, 0xd2,
	0xbc, 0x34, 0x5e, 0x82, 0x1a, 0xb9, 0x8f, 0x2c, 0x97, 0x01, 0xc4, 0x66, 0xe8, 0x8b, 0x9c, 0xb0,
	0x15, 0xd6, 0x01, 0x8f, 0x9b, 0x9e, 0x4b, 0x5d, 0x0d, 0x96, 0xbd, 0x5c, 0xcb, 0x72, 0xd9, 0x2b,
	0x59, 0x90, 0x9d, 0xd9, 0x83, 0xe0, 0x17, 0x19, 0x2e, 0xcc, 0x83, 0x04, 0x7d, 0x56, 0x22, 0xc8,
	0xd5, 0xb9, 0xc8, 0x9d, 0x1d, 0x55, 0x8a, 0x26, 0x53, 0x59, 0x7c, 0x99, 0xbf, 0x9d, 0x31, 0xaf,
	0xdf, 0xe9, 0x45, 0x47, 0x69, 0xf8, 0xc4, 0xda, 0xde, 0x36, 0x8d, 0x76, 0x4d, 0x7b, 0x09, 0x2d,
	0xb1, 0x3a, 0x51, 0x0b, 0x64, 0x2f, 0x7b, 0x13, 0xca, 0x9e, 0x4b, 0xa0, 0x18, 0x45, 0x98, 0x41,
	0x51, 0x39, 0x19, 0x8a, 0x5c, 0x59, 0xfb, 0x91, 0xb4, 0xaa, 0x3c, 0x78, 0xf2, 0xdd, 0x93, 0xb3,
	0xb7, 0x51, 0x70, 0xb2, 0xc8, 0xee, 0x32, 0x4b, 0x28, 0x5a, 0x87, 0xda, 0xc4, 0xd9, 0xc3, 0x13,
	0xae, 0xe3, 0xe7, 0x66, 0xf4, 0x2d, 0xba, 0xc2, 0x2a, 0x2d, 0x55, 0xeb, 0xdc, 0x87, 0x26, 0x27,
	0x9e, 0x53, 0x45, 0x82, 0x9f, 0x06, 0x5f, 0x38, 0x9f, 0x83, 0x42, 0xc1, 0x27, 0xc7, 0x1b, 0x05,
	0x6e, 0x7e, 0x3c, 0x32, 0x26, 0x1f, 0x71, 0x53, 0x1c, 0xc7, 0xce, 0x7e, 0xb6, 0x31, 0x9b, 0x6a,
	0x77, 0x40, 0xa1, 0xe4, 0x21, 0x2a, 0xd1, 0xcc, 0x27, 0x9f, 0xc9, 0x99, 0x0a, 0x9b, 0x12, 0x20,
	0x2d, 0x83, 0xbd, 0x80, 0x65, 0xcb, 0xd0, 0xbe, 0x93, 0x60, 0xa5, 0x88, 0xe3, 0xa9, 0x13, 0x92,
	0x0b, 0x80, 0x8e, 0xd9, 0xd3, 0xe2, 0x8a, 0x2e, 0x2c, 0xeb, 0x74, 0xc0, 0x5e, 0x8a, 0x74, 0x4c,
	0x3e, 0x7c, 0x0a, 0xe1, 0x7f, 0x27, 0xcb, 0x5d, 0x68, 0x15, 0x0b, 0x5b, 0x5e, 0x9c, 0x90, 0x8d,
	0xfc, 0x49, 0xc4, 0x8d, 0xf4, 0xdf, 0xa3, 0xa5, 0x97, 0x0a, 0xfd, 0xa5, 0x60, 0xaf, 0x46, 0xb3,
	0x7e, 0xf7, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x36, 0x7d, 0xf8, 0xf4, 0x43, 0x10, 0x00, 0x00,
}
