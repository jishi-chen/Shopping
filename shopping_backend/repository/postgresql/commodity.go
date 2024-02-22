package postgresql

import (
	"context"
	"database/sql"

	"shopping_backend/domain"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

type postgresqlCommodityRepository struct {
	db *sql.DB
}

// NewPostgresqlCommodityRepository ...
func NewPostgresqlCommodityRepository(db *sql.DB) domain.CommodityRepository {
	return &postgresqlCommodityRepository{db}
}

func (p *postgresqlCommodityRepository) GetByID(ctx context.Context, id string) (*domain.Commodity, error) {
	row := p.db.QueryRow("SELECT id FROM commoditys WHERE id = $1", id)
	d := &domain.Commodity{}
	if err := row.Scan(&d.ID, &d.UserID, &d.Name); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return d, nil
}

func (p *postgresqlCommodityRepository) Store(ctx context.Context, d *domain.Commodity) error {
	if d.ID == "" {
		d.ID = uuid.Must(uuid.NewV4()).String()
	}
	_, err := p.db.Exec(
		"INSERT INTO commoditys (id, user_id, name) VALUES ($1, $2, $3)",
		d.ID, d.UserID, d.Name,
	)
	if err != nil {
		logrus.Error(err, d)
		return err
	}
	return nil
}
