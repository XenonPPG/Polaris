package controllers

import (
	"net/http"
	"polaris/internal/db"
	"polaris/internal/gateway/utils"
	"polaris/internal/pipeline"

	"github.com/gin-gonic/gin"
)

type IngestRequest struct {
	Name string `json:"name,omitempty"`
	Data []byte `json:"data,omitempty"`
	Type string `json:"type,omitempty"`
}

type RetrieveResultItem struct {
	SourceID   uint    `json:"source_id"`
	Data       string  `json:"data"`
	Type       string  `json:"type"`
	Similarity float32 `json:"similarity"`
}

type RetrieveResponse struct {
	Results []RetrieveResultItem `json:"results"`
}

// Ingest godoc
// @Summary      Ingest content
// @Description  Accepts content (name, data, type) and passes it to the processing pipeline for indexing
// @Tags         content
// @Accept       json
// @Produce      json
// @Param        request body IngestRequest true "Content to ingest"
// @Success      200 {object} map[string]string "content ingested successfully"
// @Failure      400 {object} map[string]interface{} "Invalid request body"
// @Failure      500 {object} map[string]interface{} "Failed to ingest content"
// @Router       /ingest [post]
func (s *Service) Ingest(c *gin.Context) {
	var req IngestRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err, "body")
	}

	err := s.mainService.Ingest(c.Request.Context(), pipeline.Content{
		Name: req.Name,
		Data: req.Data,
		Type: db.ContentType(req.Type),
	})
	if err != nil {
		utils.InternalServerError(c, err, "failed to ingest content")
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"msg": "content ingested successfully",
		},
	)
}

// Retrieve godoc
// @Summary      Retrieve relevant content
// @Description  Searches previously ingested content based on a text query
// @Tags         content
// @Accept       json
// @Produce      json
// @Param        query query string true "Search query"
// @Success      200 {object} RetrieveResponse
// @Failure      400 {object} map[string]interface{} "Empty query"
// @Failure      500 {object} map[string]interface{} "Failed to retrieve content"
// @Router       /retrieve [get]
func (s *Service) Retrieve(c *gin.Context) {
	query := c.Query("query")
	if len(query) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "empty query",
		})
	}

	retrieved, err := s.mainService.Retrieve(c.Request.Context(), query)
	if err != nil {
		utils.InternalServerError(c, err, "failed to retrieve content")
		return
	}
	results := make([]RetrieveResultItem, 0, len(retrieved))
	for _, r := range retrieved {
		results = append(results, RetrieveResultItem{
			SourceID:   r.SourceID,
			Data:       r.Data,
			Type:       string(r.Type),
			Similarity: r.GetSimilarity(),
		})
	}

	c.JSON(http.StatusOK, RetrieveResponse{
		Results: results,
	})
}
