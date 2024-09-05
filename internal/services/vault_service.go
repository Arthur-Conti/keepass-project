package services

import (
	"github.com/Arthur-Conti/keepass-project/internal/config"
	"github.com/Arthur-Conti/keepass-project/internal/models"
	ports "github.com/Arthur-Conti/keepass-project/internal/ports/repositories"
)

type vaultsService struct {
	repository ports.VaultsRepositoryInterface
}

func NewVaultsService(repository ports.VaultsRepositoryInterface) *vaultsService {
	return &vaultsService{
		repository: repository,
	}
}

func (service *vaultsService) ListAllVaultsByUserID(userID string) ([]models.Vault, error) {
	return service.repository.ListAllVaultsByUserID(userID)
}

func (service *vaultsService) GetVaultByID(id string) (models.Vault, error) {
	return service.repository.GetVaultByID(id)
}

func (service *vaultsService) CreateVault(vault models.CreateVault) error {
	vaultID, err := config.CreateID()
	if err != nil {
		logger.Error("error creating vault ID: " + err.Error())
		return err
	}
	vault.ID = vaultID
	return service.repository.CreateVault(vault)
}
