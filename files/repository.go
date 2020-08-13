package files

import (
	"mime/multipart"
	"strings"
	"time"

	"github.com/hichuyamichu-me/goober/errors"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

type Repository struct {
	fs         afero.Fs
	db         *gorm.DB
	blockedExt []string
}

func NewRepository(db *gorm.DB) *Repository {
	exts := viper.GetStringSlice("blocked_extentions")
	fs := afero.NewBasePathFs(afero.NewOsFs(), viper.GetString("upload_dir"))
	return &Repository{fs: fs, db: db, blockedExt: exts}
}

func (r *Repository) Get(id uuid.UUID) (*File, error) {
	const op errors.Op = "upload/repository.Find"

	f := &File{}
	if err := r.db.Where(&File{ID: id}).First(&f).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.E(errors.Errorf("file not found"), errors.NotFound, op)
		}
		return nil, errors.E(err, errors.Internal, op)
	}

	inner, err := r.fs.Open(id.String())
	if err != nil {
		return nil, err
	}
	f.Inner = inner

	return f, nil
}

func (r *Repository) Save(file *multipart.FileHeader) (string, error) {
	const op errors.Op = "upload/service.Save"

	for _, sufix := range r.blockedExt {
		if strings.HasSuffix(file.Filename, sufix) {
			return "", errors.E(errors.Errorf("illegal file extention"), errors.Invalid, op)
		}
	}

	src, err := file.Open()
	if err != nil {
		return "", errors.E(err, errors.Internal, op)
	}
	defer src.Close()

	f := &File{Name: file.Filename, Size: file.Size, CreatedAt: time.Now().Unix()}
	if r.db.Create(f).Error != nil {
		return "", errors.E(err, errors.Internal, op)
	}

	err = afero.SafeWriteReader(r.fs, f.ID.String(), src)
	if err != nil {
		return "", errors.E(err, errors.Internal, op)
	}
	return f.ID.String(), nil
}

func (r *Repository) List(page int) ([]*File, int, error) {
	const op errors.Op = "upload/service.Save"
	const pageSize = 50

	skip := page * pageSize
	files := []*File{}
	total := 0
	if err := r.db.Limit(25).Offset(skip).Order("created_at DESC").Find(&files).Count(&total).Error; err != nil {
		return nil, 0, errors.E(err, errors.Internal, op)
	}

	return files, total, nil
}

func (r *Repository) Delete(id uuid.UUID) error {
	const op errors.Op = "upload/service.DeleteFile"

	file := &File{ID: id}
	if err := r.db.Delete(file).Error; err != nil {
		return errors.E(err, errors.Internal, op)
	}

	err := r.fs.Remove(id.String())
	if err != nil {
		return errors.E(err, errors.IO, op)
	}

	return nil
}
