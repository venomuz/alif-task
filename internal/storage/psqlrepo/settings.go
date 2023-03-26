package psqlrepo

import (
	"context"
	"github.com/venomuz/alif-task/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewSettingsRepo(db *gorm.DB) *SettingsRepo {
	return &SettingsRepo{
		db: db,
	}
}

type SettingsRepo struct {
	db *gorm.DB
}

func (s *SettingsRepo) Create(ctx context.Context, setting *models.SettingOut) error {
	err := s.db.WithContext(ctx).Select(
		"title",
		"key",
		"value",
		"created_at",
	).Create(setting).Error

	return err
}

func (s *SettingsRepo) Update(ctx context.Context, setting *models.SettingOut) error {
	columns := map[string]interface{}{
		"title":      setting.Title,
		"key":        setting.Key,
		"value":      setting.Value,
		"updated_at": setting.UpdatedAt,
	}

	err := s.db.WithContext(ctx).Clauses(clause.Returning{}).Model(setting).Updates(columns).Error

	return err
}

func (s *SettingsRepo) GetByID(ctx context.Context, ID uint32) (models.SettingOut, error) {
	var setting models.SettingOut

	err := s.db.WithContext(ctx).First(&setting, "id = ?", ID).Error

	return setting, err
}

func (s *SettingsRepo) GetAll(ctx context.Context) ([]models.SettingOut, error) {
	var settings []models.SettingOut

	err := s.db.WithContext(ctx).Order("id desc").Find(&settings).Error

	return settings, err
}

func (s *SettingsRepo) GetByKey(ctx context.Context, key string) (models.SettingOut, error) {
	var setting models.SettingOut

	err := s.db.WithContext(ctx).Where("key = ?", key).First(&setting).Error

	return setting, err
}

func (s *SettingsRepo) DeleteByID(ctx context.Context, ID uint32) error {
	err := s.db.WithContext(ctx).Delete(models.Settings{ID: ID}).Error

	return err
}
