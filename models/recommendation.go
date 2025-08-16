package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Recommendation struct {
	ID           uint        `gorm:"primaryKey" json:"id"`
	Title        string      `json:"title"`
	Description  string      `json:"description"`
	Tags         Tags        `gorm:"type:json" json:"tags"`
	Duration     int64       `json:"duration"`
}

type Tags []string

func (tag *Tags) Scan(value any) error {
	bytes, ok := value.([]byte)

	if !ok {
		return errors.New("failed to scan Tags")
	}

	return json.Unmarshal(bytes, tag)
}

func (tag Tags) Value() (driver.Value, error) {
	if tag == nil {
		return nil, nil
	}

	return json.Marshal(tag)
}
