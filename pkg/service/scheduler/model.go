// Copyright (C) 2017 ScyllaDB

package scheduler

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/scylla-manager/pkg/scheduler"
	"github.com/scylladb/scylla-manager/pkg/scheduler/trigger"
	"github.com/scylladb/scylla-manager/pkg/service"
	"github.com/scylladb/scylla-manager/pkg/store"
	"github.com/scylladb/scylla-manager/pkg/util/duration"
	"github.com/scylladb/scylla-manager/pkg/util/uuid"
	"go.uber.org/multierr"
)

// TaskType specifies the type of a Task.
type TaskType string

// TaskType enumeration.
const (
	UnknownTask               TaskType = "unknown"
	BackupTask                TaskType = "backup"
	HealthCheckAlternatorTask TaskType = "healthcheck_alternator"
	HealthCheckCQLTask        TaskType = "healthcheck"
	HealthCheckRESTTask       TaskType = "healthcheck_rest"
	RepairTask                TaskType = "repair"
	ValidateBackupTask        TaskType = "validate_backup"

	mockTask TaskType = "mock"
)

func (t TaskType) String() string {
	return string(t)
}

func (t TaskType) MarshalText() (text []byte, err error) {
	return []byte(t.String()), nil
}

func (t *TaskType) UnmarshalText(text []byte) error {
	switch TaskType(text) {
	case UnknownTask:
		*t = UnknownTask
	case BackupTask:
		*t = BackupTask
	case HealthCheckAlternatorTask:
		*t = HealthCheckAlternatorTask
	case HealthCheckCQLTask:
		*t = HealthCheckCQLTask
	case HealthCheckRESTTask:
		*t = HealthCheckRESTTask
	case RepairTask:
		*t = RepairTask
	case ValidateBackupTask:
		*t = ValidateBackupTask
	case mockTask:
		*t = mockTask
	default:
		return fmt.Errorf("unrecognized TaskType %q", text)
	}
	return nil
}

func (t TaskType) isHealthCheck() bool {
	switch t {
	case HealthCheckAlternatorTask, HealthCheckCQLTask, HealthCheckRESTTask:
		return true
	}
	return false
}

// Schedule specify task schedule.
type Schedule struct {
	gocqlx.UDT

	StartDate            time.Time         `json:"start_date"`
	Interval             duration.Duration `json:"interval" db:"interval_seconds"`
	NumRetries           int               `json:"num_retries"`
	RetryInitialInterval duration.Duration `json:"retry_initial_interval"`
}

func (s Schedule) trigger() scheduler.Trigger {
	return trigger.NewLegacy(s.StartDate, s.Interval.Duration())
}

func (s Schedule) Validate() error {
	// TO-DO add validation
	return nil
}

// Task specify task type, properties and schedule.
type Task struct {
	ClusterID  uuid.UUID       `json:"cluster_id"`
	Type       TaskType        `json:"type"`
	ID         uuid.UUID       `json:"id"`
	Name       string          `json:"name"`
	Tags       []string        `json:"tags,omitempty"`
	Enabled    bool            `json:"enabled,omitempty"`
	Sched      Schedule        `json:"schedule,omitempty"`
	Properties json.RawMessage `json:"properties,omitempty"`
}

func (t *Task) String() string {
	return fmt.Sprintf("%s/%s", t.Type, t.ID)
}

func (t *Task) Validate() error {
	if t == nil {
		return service.ErrNilPtr
	}

	var errs error
	if t.ID == uuid.Nil {
		errs = multierr.Append(errs, errors.New("missing ID"))
	}
	if t.ClusterID == uuid.Nil {
		errs = multierr.Append(errs, errors.New("missing ClusterID"))
	}
	if _, e := uuid.Parse(t.Name); e == nil {
		errs = multierr.Append(errs, errors.New("name cannot be an UUID"))
	}
	switch t.Type {
	case "", UnknownTask:
		errs = multierr.Append(errs, errors.New("no TaskType specified"))
	default:
		var tp TaskType
		errs = multierr.Append(errs, tp.UnmarshalText([]byte(t.Type)))
	}
	errs = multierr.Append(errs, t.Sched.Validate())

	return service.ErrValidate(errors.Wrap(errs, "invalid task"))
}

// Status specifies the status of a Task.
type Status string

// Status enumeration.
const (
	StatusNew      Status = "NEW"
	StatusRunning  Status = "RUNNING"
	StatusStopping Status = "STOPPING"
	StatusStopped  Status = "STOPPED"
	StatusWaiting  Status = "WAITING"
	StatusDone     Status = "DONE"
	StatusError    Status = "ERROR"
	StatusAborted  Status = "ABORTED"
)

var allStatuses = []Status{
	StatusNew,
	StatusRunning,
	StatusStopping,
	StatusStopped,
	StatusWaiting,
	StatusDone,
	StatusError,
	StatusAborted,
}

func (s Status) String() string {
	return string(s)
}

func (s Status) MarshalText() (text []byte, err error) {
	return []byte(s.String()), nil
}

func (s *Status) UnmarshalText(text []byte) error {
	switch Status(text) {
	case StatusNew:
		*s = StatusNew
	case StatusRunning:
		*s = StatusRunning
	case StatusStopping:
		*s = StatusStopping
	case StatusStopped:
		*s = StatusStopped
	case StatusWaiting:
		*s = StatusWaiting
	case StatusDone:
		*s = StatusDone
	case StatusError:
		*s = StatusError
	case StatusAborted:
		*s = StatusAborted
	default:
		return fmt.Errorf("unrecognized Status %q", text)
	}
	return nil
}

// Run describes a running instance of a Task.
type Run struct {
	ClusterID uuid.UUID  `json:"cluster_id"`
	Type      TaskType   `json:"type"`
	TaskID    uuid.UUID  `json:"task_id"`
	ID        uuid.UUID  `json:"id"`
	Status    Status     `json:"status"`
	Cause     string     `json:"cause,omitempty"`
	Owner     string     `json:"owner"`
	StartTime time.Time  `json:"start_time"`
	EndTime   *time.Time `json:"end_time,omitempty"`
}

func newRunFromTaskInfo(ti taskInfo) *Run {
	return &Run{
		ClusterID: ti.ClusterID,
		Type:      ti.TaskType,
		TaskID:    ti.TaskID,
		ID:        uuid.NewTime(),
		StartTime: now(),
		Status:    StatusRunning,
	}
}

type suspendInfo struct {
	ClusterID    uuid.UUID   `json:"-"`
	StartedAt    time.Time   `json:"started_at"`
	PendingTasks []uuid.UUID `json:"pending_tasks"`
	RunningTask  []uuid.UUID `json:"running_tasks"`
}

var _ store.Entry = &suspendInfo{}

func (v *suspendInfo) Key() (clusterID uuid.UUID, key string) {
	return v.ClusterID, "scheduler_suspended"
}

func (v *suspendInfo) MarshalBinary() (data []byte, err error) {
	return json.Marshal(v)
}

func (v *suspendInfo) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, v)
}
