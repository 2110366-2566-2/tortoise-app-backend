package services

import (
	"log"
	"net/http"
	"strconv"

	"github.com/2110366-2566-2/tortoise-app-backend/internal/database"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"github.com/2110366-2566-2/tortoise-app-backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type ReportHandler struct {
	handler *database.Handler
}

func NewReportHandler(handler *database.Handler) *ReportHandler {
	return &ReportHandler{handler: handler}
}

func (h *ReportHandler) CreatePartyReport(c *gin.Context) {
	var report models.PartyReport
	if err := c.BindJSON(&report); err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "failed to bind Report"})
		return
	}

	report.Description = utils.SanitizeString(report.Description)

	_, err := h.handler.CreatePartyReport(c, &report)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": &report})
}

func (h *ReportHandler) CreateSystemReport(c *gin.Context) {
	var report models.SystemReport
	if err := c.BindJSON(&report); err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "failed to bind partyReport"})
		return
	}

	report.Description = utils.SanitizeString(report.Description)

	_, err := h.handler.CreateSystemReport(c, &report)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": &report})
}

func (h *ReportHandler) GetReport(c *gin.Context) {
	category := c.Query("category")
	is_solved_str := c.Query("is_solved")

	category = utils.SanitizeString(category)

	var is_solved *bool
	if is_solved_str != "" {
		b, err := strconv.ParseBool(is_solved_str)
		if err != nil {
			log.Println("Error: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "incorrect is_solved"})
			return
		}
		is_solved = &b
	}

	partyReports, systemReports, err := h.handler.GetReport(c, category, is_solved)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "incorrect category"})
		return
	}

	if category == "party" {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": bson.M{"party_reports_count": len(*partyReports), "reports_category_party": &partyReports}})
	}
	if category == "system" {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": bson.M{"system_reports_count": len(*systemReports), "reports_category_system": &systemReports}})
	}
	if category == "all" || category == "" {
		c.JSON(http.StatusOK, gin.H{"success": true, "data": bson.M{"party_reports_count": len(*partyReports), "reports_category_party": &partyReports,
			"system_reports_count": len(*systemReports), "reports_category_system": &systemReports}})
	}
}
