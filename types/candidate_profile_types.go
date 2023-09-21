package types

import "time"

type Education struct {
	ID                string    `bson:"_id,omitempty"`
	UserID            string    `bson:"user_id" validate:"required"`
	QualificationName string    `bson:"qualification_name" validate:"required"`
	Major             string    `bson:"major" validate:"required"`
	InstituteName     string    `bson:"institute_name" validate:"required"`
	StartDate         time.Time `bson:"start_date"`
	CompletionDate    time.Time `bson:"completion_date"`
	GpaScore          float64   `bson:"cgpa"`
	CreatedAt         time.Time `bson:"created_at"`
	UpdatedAt         time.Time `bson:"updated_at"`
}

type Experience struct {
	ID                 string    `bson:"_id,omitempty"`
	UserID             string    `bson:"user_id" validate:"required"`
	JobTitle           string    `bson:"job_title" validate:"required"`
	CompanyName        string    `bson:"company_name" validate:"required"`
	JobLocationCity    string    `bson:"job_location_city,omitempty"`
	JobLocationState   string    `bson:"job_location_state,omitempty"`
	JobLocationCountry string    `bson:"job_location_country,omitempty"`
	JobDescription     string    `bson:"job_description,omitempty"`
	StartDate          time.Time `bson:"start_date" validate:"required"`
	EndDate            time.Time `bson:"end_date" validate:"required"`
	IsCurrentJob       bool      `bson:"is_current_job" validate:"required"`
	CreatedAt          time.Time `bson:"created_at,omitempty"`
	UpdatedAt          time.Time `bson:"updated_at,omitempty"`
}

type Skill struct {
	ID        string    `bson:"_id,omitempty"`
	UserID    string    `bson:"user_id" validate:"required"`
	SkillName string    `bson:"name" validate:"required"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

type Hobby struct {
	HobbyName string `bson:"name" validate:"required"`
}
