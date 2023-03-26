package service

import (
	"context"
	"github.com/venomuz/alif-task/internal/models"
	"github.com/venomuz/alif-task/internal/storage/psqlrepo"
)

func NewSettingsService(settingsRepo psqlrepo.Settings) *SettingsService {
	return &SettingsService{settingsRepo: settingsRepo}
}

type SettingsService struct {
	settingsRepo psqlrepo.Settings
}

func (s *SettingsService) Create(ctx context.Context, input models.CreateSettingInput) (models.SettingOut, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SettingsService) Update(ctx context.Context, input models.UpdateSettingInput) (models.SettingOut, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SettingsService) GetByID(ctx context.Context, ID uint32) (models.SettingOut, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SettingsService) GetAll(ctx context.Context) ([]models.SettingOut, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SettingsService) GetByKey(ctx context.Context, key string) (models.SettingOut, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SettingsService) DeleteByID(ctx context.Context, ID uint32) error {
	//TODO implement me
	panic("implement me")
}
