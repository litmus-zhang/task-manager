// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type PriorityLevel string

const (
	PriorityLevelLow    PriorityLevel = "low"
	PriorityLevelMedium PriorityLevel = "medium"
	PriorityLevelHigh   PriorityLevel = "high"
	PriorityLevelUrgent PriorityLevel = "urgent"
)

func (e *PriorityLevel) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = PriorityLevel(s)
	case string:
		*e = PriorityLevel(s)
	default:
		return fmt.Errorf("unsupported scan type for PriorityLevel: %T", src)
	}
	return nil
}

type NullPriorityLevel struct {
	PriorityLevel PriorityLevel `json:"priority_level"`
	Valid         bool          `json:"valid"` // Valid is true if PriorityLevel is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullPriorityLevel) Scan(value interface{}) error {
	if value == nil {
		ns.PriorityLevel, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.PriorityLevel.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullPriorityLevel) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.PriorityLevel), nil
}

type TimelineType string

const (
	TimelineTypeDaily   TimelineType = "daily"
	TimelineTypeWeekly  TimelineType = "weekly"
	TimelineTypeMonthly TimelineType = "monthly"
)

func (e *TimelineType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = TimelineType(s)
	case string:
		*e = TimelineType(s)
	default:
		return fmt.Errorf("unsupported scan type for TimelineType: %T", src)
	}
	return nil
}

type NullTimelineType struct {
	TimelineType TimelineType `json:"timeline_type"`
	Valid        bool         `json:"valid"` // Valid is true if TimelineType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullTimelineType) Scan(value interface{}) error {
	if value == nil {
		ns.TimelineType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.TimelineType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullTimelineType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.TimelineType), nil
}

type Category struct {
	CategoryID   int32          `json:"category_id"`
	UserID       int32          `json:"user_id"`
	CategoryName string         `json:"category_name"`
	Description  sql.NullString `json:"description"`
	ColorHex     string         `json:"color_hex"`
	CreatedAt    sql.NullTime   `json:"created_at"`
}

type SubscriptionPlan struct {
	PlanID      int32          `json:"plan_id"`
	PlanName    string         `json:"plan_name"`
	Description sql.NullString `json:"description"`
	Price       sql.NullString `json:"price"`
	CreatedAt   sql.NullTime   `json:"created_at"`
}

type Task struct {
	TaskID       int32          `json:"task_id"`
	CategoryID   int32          `json:"category_id"`
	TaskName     string         `json:"task_name"`
	Description  sql.NullString `json:"description"`
	TimelineType TimelineType   `json:"timeline_type"`
	Priority     PriorityLevel  `json:"priority"`
	CreatedAt    sql.NullTime   `json:"created_at"`
	DueDate      time.Time      `json:"due_date"`
	IsCompleted  sql.NullBool   `json:"is_completed"`
	UpdatedAt    sql.NullTime   `json:"updated_at"`
}

type TaskDeadline struct {
	TaskID             int32         `json:"task_id"`
	TaskName           string        `json:"task_name"`
	CategoryName       string        `json:"category_name"`
	ColorHex           string        `json:"color_hex"`
	TimelineType       TimelineType  `json:"timeline_type"`
	Priority           PriorityLevel `json:"priority"`
	DueDate            time.Time     `json:"due_date"`
	IsCompleted        sql.NullBool  `json:"is_completed"`
	DueDateDescription interface{}   `json:"due_date_description"`
}

type User struct {
	UserID                int32        `json:"user_id"`
	Username              string       `json:"username"`
	Email                 string       `json:"email"`
	PasswordHash          string       `json:"password_hash"`
	SubscriptionPlanID    int32        `json:"subscription_plan_id"`
	SubscriptionStartDate time.Time    `json:"subscription_start_date"`
	SubscriptionEndDate   sql.NullTime `json:"subscription_end_date"`
	CreatedAt             sql.NullTime `json:"created_at"`
}
