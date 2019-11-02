package repositories

import "github.com/echoloveyou/micro/bos_web_proto/models"

type AdminUserRepository struct {
	Repository
}

type IAdminUserRepository interface {
	Add(adminUser models.AdminUser) bool
}

func NewAdminUserRepository() IAdminUserRepository {
	var (
		adminUserRepository AdminUserRepository
	)
	adminUserRepository.GDB()
	return &adminUserRepository
}

func (repository *AdminUserRepository) Add(adminUser models.AdminUser) bool {
	err := repository.Model(&adminUser).Create(&adminUser).Error
	if err != nil {
		return false
	}
	return true
}
