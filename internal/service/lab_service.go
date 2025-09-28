package service

import (
	"context"

	// NOTE: Replace 'yourproject' with your Go module name
	"github.com/medorahealth/medora_backend/internal/repo"
	"github.com/medorahealth/medora_backend/internal/model"
)

// LabService defines the business logic operations for labs.
type LabService interface {
	Register(ctx context.Context, req *model.RegisterLabRequest) (*model.Lab, error)
}

// labService implements the LabService interface.
type labService struct {
	labRepo repo.LabRepository
}

// NewLabService creates a new lab service.
func NewLabService(labRepo repo.LabRepository) LabService {
	return &labService{labRepo: labRepo}
}

// Register handles the business logic for creating a new lab.
func (s *labService) Register(ctx context.Context, req *model.RegisterLabRequest) (*model.Lab, error) {
	newLab := &model.Lab{
		OperatingDrID:      req.OperatingDrID,
		LegalName:          req.LegalName,
		Description:        req.Description,
		Accreditation:      req.Accreditation,
		LicenseNumber:      req.LicenseNumber,
		LicensePDF:         req.LicensePDF,
		PollutionPolicyPDF: req.PollutionPolicyPDF,
		ContactEmail:       req.ContactEmail,
		ContactPhone:       req.ContactPhone,
		Line1:              req.Line1,
		Line2:              req.Line2,
		City:               req.City,
		State:              req.State,
		Pincode:            req.Pincode,
		Latitude:           req.Latitude,
		Longitude:          req.Longitude,
		OpensAt:            req.OpensAt,
		ClosesAt:           req.ClosesAt,
		EmergencyService:   req.EmergencyService,
	}

	err := s.labRepo.CreateLab(ctx, newLab)
	if err != nil {
		return nil, err
	}

	return newLab, nil
}
