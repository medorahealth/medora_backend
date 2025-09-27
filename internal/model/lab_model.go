package model

import (
	"time"
)

// Lab represents the Labs table in the database.
type Lab struct {
	ID                 int64     `json:"id" db:"labID"`
	OperatingDrID      int64     `json:"operating_dr_id" db:"operating_dr_id"`
	LegalName          string    `json:"legal_name" db:"legal_name"`
	Description        *string   `json:"description,omitempty" db:"description"`
	Accreditation      string    `json:"accreditation" db:"accreditation"`
	LicenseNumber      *string   `json:"license_number,omitempty" db:"license_number"`
	LicensePDF         string    `json:"license_pdf" db:"license_pdf"`
	PollutionPolicyPDF string    `json:"pollution_policy_pdf" db:"pollution_policy_pdf"`
	ContactEmail       *string   `json:"contact_email,omitempty" db:"contact_email"`
	ContactPhone       *string   `json:"contact_phone,omitempty" db:"contact_phone"`
	Line1              string    `json:"line1" db:"line1"`
	Line2              *string   `json:"line2,omitempty" db:"line2"`
	City               string    `json:"city" db:"city"`
	State              string    `json:"state" db:"state"`
	Pincode            string    `json:"pincode" db:"pincode"`
	Latitude           *float64  `json:"latitude,omitempty" db:"latitude"`
	Longitude          *float64  `json:"longitude,omitempty" db:"longitude"`
	OpensAt            *string   `json:"opens_at,omitempty" db:"opens_at"`
	ClosesAt           *string   `json:"closes_at,omitempty" db:"closes_at"`
	EmergencyService   bool      `json:"emergency_service" db:"emergency_service"`
	IsActive           bool      `json:"is_active" db:"is_active"`
	Rating             float64   `json:"rating" db:"rating"`
	RatingsCount       int32     `json:"ratings_count" db:"ratings_count"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" db:"updated_at"`
}

// RegisterLabRequest is the DTO for creating a new lab.
type RegisterLabRequest struct {
	OperatingDrID      int64    `json:"operating_dr_id" validate:"required"`
	LegalName          string   `json:"legal_name" validate:"required,min=3,max=255"`
	Description        *string  `json:"description,omitempty"`
	Accreditation      string   `json:"accreditation" validate:"required"`
	LicenseNumber      *string  `json:"license_number,omitempty"`
	LicensePDF         string   `json:"license_pdf" validate:"required,url"`
	PollutionPolicyPDF string   `json:"pollution_policy_pdf" validate:"required,url"`
	ContactEmail       *string  `json:"contact_email,omitempty" validate:"omitempty,email"`
	ContactPhone       *string  `json:"contact_phone,omitempty" validate:"omitempty,e164"`
	Line1              string   `json:"line1" validate:"required"`
	Line2              *string  `json:"line2,omitempty"`
	City               string   `json:"city" validate:"required"`
	State              string   `json:"state" validate:"required"`
	Pincode            string   `json:"pincode" validate:"required,len=6"`
	Latitude           *float64 `json:"latitude,omitempty" validate:"omitempty,latitude"`
	Longitude          *float64 `json:"longitude,omitempty" validate:"omitempty,longitude"`
	OpensAt            *string  `json:"opens_at,omitempty" validate:"omitempty,time=15:04:05"`
	ClosesAt           *string  `json:"closes_at,omitempty" validate:"omitempty,time=15:04:05"`
	EmergencyService   bool     `json:"emergency_service"`
}
