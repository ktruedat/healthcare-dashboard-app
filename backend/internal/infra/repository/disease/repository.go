package disease

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ktruedat/healthy/internal/infra/repository/disease/model"
)

const (
	diseaseTableName         = "disease"
	diseaseCategoryTableName = "diseasecategory"
	idColumn                 = "id"
	nameColumn               = "name"
	categoryIDColumn         = "category_id"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepo(db *pgxpool.Pool) *repo {
	return &repo{db: db}
}

func (r *repo) GetDiseaseWithCategory(ctx context.Context, diseaseID int) (*model.Disease, error) {
	query, args, err := sq.Select(
		"d."+idColumn, "d."+nameColumn, "d."+categoryIDColumn,
		"c."+idColumn, "c."+nameColumn,
	).
		PlaceholderFormat(sq.Dollar).
		From(diseaseTableName + " d").
		Join(diseaseCategoryTableName + " c ON d." + categoryIDColumn + " = c." + idColumn).
		Where(sq.Eq{"d." + idColumn: diseaseID}).
		ToSql()
	if err != nil {
		return nil, err
	}

	row := r.db.QueryRow(ctx, query, args...)

	var disease model.Disease
	var category model.DiseaseCategory

	if err := row.Scan(&disease.ID, &disease.Name, &disease.CategoryID, &category.ID, &category.Name); err != nil {
		return nil, err
	}
	disease.Category = category

	return &disease, nil
}
