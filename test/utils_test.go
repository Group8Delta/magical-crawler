package test

import (
	"fmt"
	"magical-crwler/models"
	"magical-crwler/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessageGenerator(t *testing.T) {
	ad := []models.Ad{models.Ad{}}
	result := utils.GenerateFilterMessage(ad)
	fmt.Println(result)
	assert.NotEmpty(t, result)
}
