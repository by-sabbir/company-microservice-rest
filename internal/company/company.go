package company

import (
	"context"
	"errors"
	"log"
)

var (
	ErrTypeNotFound = errors.New("provided company type is not defined")
)

var CompanyType = []string{
	"Corporations", "NonProfit", "Cooperative", "Sole Proprietorship",
}

// Company - representation of a company structure
type Company struct {
	ID             string
	Name           string
	Description    string
	TotalEmployees string
	IsRegistered   bool
	Type           string
}

// Service - is the struct containing business logics
type Service struct {
	Store Store
}

// NewService - returns to a pointer to a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

// Store - this interface defines all the methods to operate
type Store interface {
	GetCompany(context.Context, string) (Company, error)
	PostCompany(context.Context, Company) (Company, error)
	DeleteCompany(context.Context, string) error
	PartialUpdateCompany(context.Context, string, Company) (Company, error)
}

func (c *Company) ScanType(t string) error {
	for _, v := range CompanyType {
		if c.Type == v {
			return nil
		}
	}
	return ErrTypeNotFound
}

func (s *Service) GetCompany(ctx context.Context, id string) (Company, error) {
	log.Println("Retreiving the Company")

	cmt, err := s.Store.GetCompany(ctx, id)
	if err != nil {
		return Company{}, err
	}
	return cmt, nil
}

func (s *Service) PostCompany(ctx context.Context, cmt Company) (Company, error) {
	log.Println("Creating Company")

	cmt, err := s.Store.PostCompany(ctx, cmt)
	if err != nil {
		return Company{}, err
	}
	return cmt, nil
}

func (s *Service) DeleteCompany(ctx context.Context, id string) error {
	log.Println("Deleting Company")

	err := s.Store.DeleteCompany(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) PartialUpdateCompany(
	ctx context.Context, id string, cmt Company,
) (Company, error) {
	log.Println("Updating Company...")
	cmt, err := s.Store.PartialUpdateCompany(ctx, id, cmt)
	if err != nil {
		return Company{}, err
	}
	return cmt, nil
}
