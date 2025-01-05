package disease

import (
	"context"
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
)

func TestGetDiseaseWithCategory(t *testing.T) {
	// Set up a connection pool to the test database
	dbpool, err := pgxpool.New(context.Background(), "postgres://admin:admin@localhost:5432/healthcare_db")
	if err != nil {
		t.Fatalf("Unable to connect to database: %v", err)
	}
	defer dbpool.Close()

	// Create a new repository instance
	repo := NewRepo(dbpool)

	// Insert test data using Squirrel
	insertCategoryQuery, insertCategoryArgs, err := sq.Insert(diseaseCategoryTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(idColumn, nameColumn).
		Values(999, "Category1").
		ToSql()
	if err != nil {
		t.Fatalf("Unable to build insert category query: %v", err)
	}
	if _, err = dbpool.Exec(context.Background(), insertCategoryQuery, insertCategoryArgs...); err != nil {
		t.Fatalf("Unable to insert test category data: %v", err)
	}

	insertDiseaseQuery, insertDiseaseArgs, err := sq.Insert(diseaseTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(idColumn, nameColumn, categoryIDColumn).
		Values(999, "Disease1", 999).
		ToSql()
	if err != nil {
		t.Fatalf("Unable to build insert disease query: %v", err)
	}
	if _, err = dbpool.Exec(context.Background(), insertDiseaseQuery, insertDiseaseArgs...); err != nil {
		t.Fatalf("Unable to insert test disease data: %v", err)
	}

	// Call the function to test
	disease, err := repo.GetDiseaseWithCategory(context.Background(), 999)
	if err != nil {
		t.Fatalf("Error getting disease with category: %v", err)
	}

	// Assert the results
	assert.NotNil(t, disease)
	assert.Equal(t, 999, disease.ID)
	assert.Equal(t, "Disease1", disease.Name)
	assert.Equal(t, 999, disease.CategoryID)
	assert.Equal(t, 999, disease.Category.ID)
	assert.Equal(t, "Category1", disease.Category.Name)

	// Clean up test data using Squirrel
	deleteDiseaseQuery, deleteDiseaseArgs, err := sq.Delete(diseaseTableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: 999}).
		ToSql()
	if err != nil {
		t.Fatalf("Unable to build delete disease query: %v", err)
	}
	_, err = dbpool.Exec(context.Background(), deleteDiseaseQuery, deleteDiseaseArgs...)
	if err != nil {
		t.Fatalf("Unable to delete test disease data: %v", err)
	}

	deleteCategoryQuery, deleteCategoryArgs, err := sq.Delete(diseaseCategoryTableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: 999}).
		ToSql()
	if err != nil {
		t.Fatalf("Unable to build delete category query: %v", err)
	}
	_, err = dbpool.Exec(context.Background(), deleteCategoryQuery, deleteCategoryArgs...)
	if err != nil {
		t.Fatalf("Unable to delete test category data: %v", err)
	}
}
