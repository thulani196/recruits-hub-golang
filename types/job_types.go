package types

import "time"

type Job struct {
	ID              string    `bson:"_id,omitempty"`
	Title           string    `bson:"title" validate:"required"`
	Description     string    `bson:"description" validate:"required"`
	IsActive        bool      `bson:"isActive" validate:"required"`
	CompanyId       string    `bson:"companyId"`
	Industry        string    `bson:"industry"`
	TotalApplicants int64     `bson:"totalApplicants"`
	CreatedAt       time.Time `bson:"createdAt"`
	UpdatedAt       time.Time `bson:"updatedAt"`
}

type Industry struct {
	ID           string    `bson:"_id,omitempty"`
	IndustryName string    `bson:"categoryName,omitempty"`
	Jobs         []string  `bson:"jobs"`
	CreatedAt    time.Time `bson:"createdAt"`
	UpdatedAt    time.Time `bson:"updatedAt"`
}
