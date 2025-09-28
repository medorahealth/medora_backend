package repo

import (
	"context"
	"database/sql"
	"github.com/medorahealth/medora_backend/internal/model"
)

// LabRepository is the interface for lab database operations.
type LabRepository interface {
	CreateLab(ctx context.Context, l *model.Lab) error
}

// labRepo implements the LabRepository interface.
type labRepo struct {
	db *sql.DB
}

// NewLabRepo creates a new lab repository.
func NewLabRepo(db *sql.DB) LabRepository {
	return &labRepo{db: db}
}

// CreateLab inserts a new lab record into the database.
func (r *labRepo) CreateLab(ctx context.Context, l *model.Lab) error {
	query := `
		INSERT INTO Labs (
			operating_dr_id, legal_name, description, accreditation, license_number,
			license_pdf, pollution_policy_pdf, contact_email, contact_phone,
			line1, line2, city, state, pincode, latitude, longitude,
			opens_at, closes_at, emergency_service
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13,
			$14, $15, $16, $17, $18, $19
		) RETURNING labID, is_active, rating, ratings_count, created_at, updated_at;
	`
	return r.db.QueryRowContext(ctx, query,
		l.OperatingDrID, l.LegalName, l.Description, l.Accreditation, l.LicenseNumber,
		l.LicensePDF, l.PollutionPolicyPDF, l.ContactEmail, l.ContactPhone,
		l.Line1, l.Line2, l.City, l.State, l.Pincode, l.Latitude, l.Longitude,
		l.OpensAt, l.ClosesAt, l.EmergencyService,
	).Scan(
		&l.ID, &l.IsActive, &l.Rating, &l.RatingsCount, &l.CreatedAt, &l.UpdatedAt,
	)
}
