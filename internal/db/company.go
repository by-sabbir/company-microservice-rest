package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/by-sabbir/company-microservice-rest/internal/company"
	"github.com/google/uuid"
)

type CompanyRow struct {
	ID             string
	Name           string
	Description    sql.NullString
	TotalEmployees int
	IsRegistered   bool
	Type           string
}

func convertCompany(c CompanyRow) company.Company {
	return company.Company{
		ID:             c.ID,
		Name:           c.Name,
		Description:    c.Description.String,
		TotalEmployees: c.TotalEmployees,
		IsRegistered:   c.IsRegistered,
		Type:           c.Type,
	}
}

func (d *DataBase) GetCompany(ctx context.Context, uuid string) (company.Company, error) {
	var cmpRow CompanyRow
	row := d.Client.QueryRowContext(
		ctx,
		`select id, name, description, total_employees, is_registered, type
		from company
		where id::text=$1`,
		uuid,
	)
	err := row.Scan(&cmpRow.ID, &cmpRow.Name, &cmpRow.Description,
		&cmpRow.TotalEmployees, &cmpRow.IsRegistered, &cmpRow.Type)
	if err != nil {
		return company.Company{}, fmt.Errorf("error fetching company from uuid: %+v", err)
	}
	return convertCompany(cmpRow), nil
}

func (d *DataBase) PostCompany(ctx context.Context, cmp company.Company) (company.Company, error) {
	cmp.ID = uuid.NewString()
	postRow := CompanyRow{
		ID:             cmp.ID,
		Name:           cmp.Name,
		Description:    sql.NullString{String: cmp.Description, Valid: true},
		TotalEmployees: cmp.TotalEmployees,
		IsRegistered:   cmp.IsRegistered,
		Type:           cmp.Type,
	}
	row, err := d.Client.NamedQueryContext(
		ctx,
		`insert into company
		(id, name, description, total_employees, is_registered, type)
		values
		(:id, :name, :description, :is_registered, :type)`,
		postRow,
	)
	if err != nil {
		return company.Company{}, fmt.Errorf("error posting Company: %+v", err)
	}
	if err := row.Close(); err != nil {
		return company.Company{}, fmt.Errorf("could not close the row, %+v", err)
	}
	defer row.Close()

	return convertCompany(postRow), nil
}

func (d *DataBase) DeleteCompany(ctx context.Context, uuid string) error {
	row, err := d.Client.ExecContext(
		ctx,
		`delete from company
		where id::text=$1`,
		uuid,
	)
	log.Println("deleted: ", row)
	if err != nil {
		return err
	}
	return nil
}

func (d *DataBase) PartialUpdateCompany(
	ctx context.Context, id string, cmp company.Company,
) (company.Company, error) {
	cmpRow := CompanyRow{
		ID:             cmp.ID,
		Name:           cmp.Name,
		Description:    sql.NullString{String: cmp.Description, Valid: true},
		TotalEmployees: cmp.TotalEmployees,
		IsRegistered:   cmp.IsRegistered,
		Type:           cmp.Type,
	}
	row, err := d.Client.NamedQueryContext(
		ctx,
		`update company set
		body=:body, slug=:slug, author=:author
		where id=:id`,
		cmpRow,
	)
	if err != nil {
		return company.Company{}, fmt.Errorf("error updating Company: %w", err)
	}
	if err := row.Close(); err != nil {
		return company.Company{}, fmt.Errorf("error updating row: %w", err)
	}
	defer row.Close()

	return convertCompany(cmpRow), nil
}
