package types

import "time"

type Company struct {
	ID             string    `bson:"_id,omitempty"`
	Name           string    `bson:"name" validate:"required"`
	Description    string    `bson:"description"`
	EstablishedAt  time.Time `bson:"establishedAt"`
	CompanyWebsite string    `bson:"companyWebsite"`
	CompanyEmail   string    `bson:"companyEmail"`
	AdminId        string    `bson:"adminId"`
	CreatedAt      time.Time `bson:"createdAt"`
	UpdatedAt      time.Time `bson:"updatedAt"`
}

type CompanyImage struct {
	ID           string    `bson:"_id,omitempty"`
	CompanyID    string    `bson:"companyId" validate:"required"`
	CompanyImage string    `bson:"companyImage" validate:"required"`
	CreatedAt    time.Time `bson:"createdAt"`
	UpdatedAt    time.Time `bson:"updatedAt"`
}
