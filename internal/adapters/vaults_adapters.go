package adapters

import (
	ports_services "github.com/Arthur-Conti/keepass-project/internal/ports/services"
	"github.com/Arthur-Conti/keepass-project/internal/repositories"
	"github.com/Arthur-Conti/keepass-project/internal/services"
)

func NewVaultsServiceAdapter() ports_services.VaultsServiceInterface {
	vaultRepository := repositories.NewVaultRepository()
	vaultService := services.NewVaultsService(vaultRepository)
	return vaultService
}