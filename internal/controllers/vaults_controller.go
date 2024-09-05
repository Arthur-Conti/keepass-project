package controllers

import (
	"net/http"
	"github.com/Arthur-Conti/keepass-project/internal/adapters"
	"github.com/Arthur-Conti/keepass-project/internal/models"
	"github.com/gin-gonic/gin"
)

type vaultsController struct {}

func NewVaultsController() *vaultsController {
	return &vaultsController{}
}

func (controller *vaultsController) ListAllVaultsByUserIDHandler(ginContext *gin.Context) {
	vaultsService := adapters.NewVaultsServiceAdapter()
	userID, _ := ginContext.Params.Get("user_id")
	vaults, err := vaultsService.ListAllVaultsByUserID(userID)
	if err != nil {
		sendError(ginContext, http.StatusInternalServerError, "error listing vaults: " + err.Error())
		return
	}
	sendSuccess(ginContext, http.StatusOK, vaults)
}

func (controller *vaultsController) GetVaultByIDHandler(ginContext *gin.Context) {
	vaultsService := adapters.NewVaultsServiceAdapter()
	id, _ := ginContext.Params.Get("id")
	vault, err := vaultsService.GetVaultByID(id)
	if err != nil {
		sendError(ginContext, http.StatusInternalServerError, "error getting vault by id: " + err.Error())
		return
	}
	sendSuccess(ginContext, http.StatusOK, vault)
}

func (controller *vaultsController) CreateVaultHandler(ginContext *gin.Context) {
	vaultsService := adapters.NewVaultsServiceAdapter()
	requestContent := VaultInformationRequest{}
	err := ginContext.BindJSON(&requestContent)
	if err != nil {
		sendError(ginContext, http.StatusInternalServerError, "error binding json: " + err.Error())
		return
	}
	err = vaultsService.CreateVault(models.CreateVault{
		UserID: requestContent.UserID,
		Name: requestContent.Name,
		Description: requestContent.Description,
		Password: requestContent.Password,
	})
	if err != nil {
		sendError(ginContext, http.StatusInternalServerError, "error creating vault: " + err.Error())
		return
	}
	sendSuccess(ginContext, http.StatusCreated, "Created")
}