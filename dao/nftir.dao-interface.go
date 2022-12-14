/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
	@Purpose: env.utils.go provides function that helps load environemnt variables
*/

// @package
package dao

// @import
import "NFTir/server/models"

// @notice DAO interface
type NftirDao interface {

	// @notice Gets all the collections from dynamoDB table
	// 
	// @return *[]models.Collection
	// 
	// @return error
	GetAll() (*[]models.Collection, error)

	// @notice Gets the total number of collections stored in the NFTier dynamoDB table
	// 
	// @return *models.HttpStatusMessage
	// 
	// @return error
	GetStatus() (*models.HttpStatusMessage, error)

	// @notice Gets a subset of collections based on the params passed in
	// 
	// @param volumeUsd *float64
	// 
	// @param marketCapUsd *float64
	// 
	// @return *[]models.Colelction
	// 
	// @return error
	GetCollectionsGreaterThan(volumeUsd *float32) (*[]models.Collection, error)
}