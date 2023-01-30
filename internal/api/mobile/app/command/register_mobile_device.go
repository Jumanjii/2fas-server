package command

import (
	"github.com/2fas/api/internal/api/mobile/domain"
	"github.com/google/uuid"
)

type RegisterMobileDevice struct {
	Id       uuid.UUID
	Name     string `json:"name" validate:"required,max=128"`
	Platform string `json:"platform" validate:"required,oneof=ios android huawei"`
	FcmToken string `json:"fcm_token" validate:"required,max=256"`
}

type RegisterMobileDeviceHandler struct {
	Repository domain.MobileDeviceRepository
}

func (h *RegisterMobileDeviceHandler) Handle(cmd *RegisterMobileDevice) error {
	mobileDevice := domain.NewMobileDevice(cmd.Id, cmd.Name, cmd.Platform, cmd.FcmToken)

	return h.Repository.Save(mobileDevice)
}