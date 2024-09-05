package repositories

import (
	"github.com/Arthur-Conti/keepass-project/internal/config"
	"github.com/Arthur-Conti/keepass-project/internal/models"
)

type VaultRepository struct {}

func NewVaultRepository() *VaultRepository {
	return &VaultRepository{}
}

func (repository *VaultRepository) ListAllVaultsByUserID(userID string) ([]models.Vault, error) {
	postgres := config.NewPostgresConnection()
	err := postgres.Connect()
	if err != nil {
		logger.Error("error connecting to postgres: " + err.Error())
		return nil, err
	}
	defer postgres.CloseConnection()
	query := `SELECT id, user_id, name, description, created_time FROM vaults WHERE user_id = $1`
	rows, err := postgres.RunQuery(query, userID) 
	if err != nil {
		logger.Error("error listing all vaults: " + err.Error())
		return nil, err
	}
	var vaults []models.Vault
	for rows.Next() {
		var vault models.Vault
		err := rows.Scan(&vault.ID, &vault.UserID, &vault.Name, &vault.Description, &vault.CreatedTime)
		if err != nil {
			logger.Error("error scannig rows: " + err.Error())
			return nil, err
		}
		vaults = append(vaults, vault)
	}
	return vaults, nil
}

func (repository *VaultRepository) GetVaultByID(id string) (models.Vault, error) {
	var vault models.Vault
	postgres := config.NewPostgresConnection()
	err := postgres.Connect()
	if err != nil {
		logger.Error("error connecting to postgres: " + err.Error())
		return vault, err
	}
	defer postgres.CloseConnection()
	query := `SELECT id, user_id, name, description, created_time FROM vaults WHERE id = $1`
	postgres.RunQueryRow(query, id).Scan(&vault.ID, &vault.UserID, &vault.Name, &vault.Description, &vault.CreatedTime)
	return vault, nil
}

func (repository *VaultRepository) CreateVault(vault models.CreateVault) error {
	postgres := config.NewPostgresConnection()
	err := postgres.Connect()
	if err != nil {
		logger.Error("error connecting to postgres: " + err.Error())
		return err
	}
	defer postgres.CloseConnection() 
	createVaultQuery := `
	INSERT INTO vaults (id, user_id, name, description, password)
	VALUES ($1, $2, $3, $4, $5)
	`
	postgres.RunQueryRow(createVaultQuery, vault.ID, vault.UserID, vault.Name, vault.Description, vault.Password)
	return nil
}