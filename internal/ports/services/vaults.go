package ports_services

import "github.com/Arthur-Conti/keepass-project/internal/models"

type VaultsServiceInterface interface {
	ListAllVaultsByUserID(string) ([]models.Vault, error)
	GetVaultByID(string) (models.Vault, error)
	CreateVault(models.CreateVault) error
}