package upload

import (
	"github.com/hichuyamichu-me/goober/errors"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Repository performs CRUD on files table
type Repository struct {
	db *gorm.DB
}

// NewRepository creates new Repository
func NewRepository(db *gorm.DB) *Repository {
	r := &Repository{db: db}
	return r
}

// Find finds files
func (r *Repository) Find(skip int) ([]*File, error) {
	const op errors.Op = "upload/repository.Find"

	files := []*File{}
	if err := r.db.Limit(25).Offset(skip).Order("created_at DESC").Find(&files).Error; err != nil {
		return nil, errors.E(err, errors.Internal, op)
	}

	return files, nil
}

// FindOne finds one file
func (r *Repository) FindOne(id uuid.UUID) (*File, error) {
	const op errors.Op = "upload/repository.FindOne"

	file := &File{}
	if err := r.db.Where(&File{ID: id}).First(&file).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.E(errors.Errorf("file not found"), errors.NotFound, op)
		}
		return nil, errors.E(err, errors.Internal, op)
	}

	return file, nil
}

// Create saves file info in DB
func (r *Repository) Create(f *File) error {
	const op errors.Op = "upload/repository.Create"

	if err := r.db.Create(f).Error; err != nil {
		return errors.E(err, errors.Internal, op)
	}

	return nil
}

func (r *Repository) Delete(file *File) error {
	const op errors.Op = "upload/repository.Delete"

	if err := r.db.Delete(file).Error; err != nil {
		return errors.E(err, errors.Internal, op)
	}

	return nil
}
