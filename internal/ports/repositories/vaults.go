package ports_repositories

import "github.com/Arthur-Conti/keepass-project/internal/models"

type VaultsRepositoryInterface interface {
	ListAllVaultsByUserID(string) ([]models.Vault, error)
	GetVaultByID(string) (models.Vault, error)
	CreateVault(models.CreateVault) error
}
