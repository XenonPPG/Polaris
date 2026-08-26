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
	results := make([]gin.H, 0, len(retrieved))
	for _, r := range retrieved {
		results = append(results, gin.H{
			"source_id":  r.SourceID,
			"data":       r.Data,
			"type":       r.Type,
			"similarity": r.GetSimilarity(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"results": results,
	})
}
