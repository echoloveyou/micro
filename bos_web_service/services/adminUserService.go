package services

import (
	"context"
	"github.com/echoloveyou/micro/bos_web_proto/models"
	"github.com/echoloveyou/micro/bos_web_service/repositories"
)

type AdminUserService struct {
	adminUserRepository repositories.AdminUserRepository
}

func (service AdminUserService) AddAdminUser(ctx context.Context, in *models.AdminUser, out *models.AdminUser) error {
	if service.adminUserRepository.Add(*in) {
		return nil
	}
	return nil
}

//
//func (service AdminUserService) AddAdminUser(ctx context.Context, in *models.AdminUser, opts ...client.CallOption) (*models.AdminUser, error) {
//	if service.adminUserRepository.Add(*in) {
//		return in, nil
//	}
//	return in, nil
//}





