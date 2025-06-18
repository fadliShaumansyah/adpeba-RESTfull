package controllers

import (
	"go-adpeba/config"
	"go-adpeba/helper"
	"go-adpeba/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//` GetLeaders retrieves all leaders`
func GetLeaders(c *gin.Context) {
    var leaders []models.Leader
    config.DB.Find(&leaders)
    c.JSON(http.StatusOK, leaders)
}


// GetLeaderByID retrieves a leader by ID
func GetLeaderByID(c *gin.Context) {
    var leader models.Leader
    id := c.Param("id")

    if err := config.DB.First(&leader, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
        return
    }
    c.JSON(http.StatusOK, leader)
}


// CreateLeader creates a new leader
func CreateLeader(c *gin.Context) {
    var input models.Leader

    if err := c.ShouldBindJSON(&input); err != nil {
        // Format error validasi jadi array
        errorList := helper.FormatValidationError(err)

        // Gunakan helper response error (versi sudah di-upgrade)
        helper.RespondBadRequest(c, "Validasi gagal", errorList)
        return
    }

    // Simpan ke database
    if err := config.DB.Create(&input).Error; err != nil {
        helper.RespondInternalError(c, "Gagal menyimpan ke database")
        return
    }

    // Response sukses
    helper.RespondCreated(c, "Data leader berhasil dibuat", input)
}

// UpdateLeader updates a leader by ID
func UpdateLeader(c *gin.Context) {
    var leader models.Leader
    id := c.Param("id")

    if err := config.DB.First(&leader, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
        return
    }

    var input models.Leader
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    config.DB.Model(&leader).Updates(input)
    c.JSON(http.StatusOK, leader)
}

// DeleteLeader deletes a leader by ID
func DeleteLeader(c *gin.Context) {
    var leader models.Leader
    id := c.Param("id")

    if err := config.DB.Delete(&leader, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
