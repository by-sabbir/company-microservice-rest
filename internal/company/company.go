package company

import (
	"context"
	"errors"
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

// Store - this interface defines all the methods to operate
type Store interface {
	GetCompany(context.Context, string) (Company, error)
	PostCompany(context.Context, Company) (Company, error)
	DeleteCompany(context.Context, string) error
	UpdateCompany(context.Context, string, Company) (Company, error)
}

func (c *Company) ScanType(t string) error {
	for _, v := range CompanyType {
		if c.Type == v {
			return nil
		}
	}
	return ErrTypeNotFound
}
