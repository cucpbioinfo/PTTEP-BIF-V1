package controllers

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/wichadak/eDNA/constants"
	"github.com/wichadak/eDNA/services"
	"github.com/wichadak/eDNA/types"
)

type SpeciesVideoController struct {
	SpeciesVideoService *services.SpeciesVideoService
}

func NewSpeciesVideoController(speciesVideoService *services.SpeciesVideoService) *SpeciesVideoController {
	return &SpeciesVideoController{
		SpeciesVideoService: speciesVideoService,
	}
}

func (speciesVideoController *SpeciesVideoController) UploadVideo(c *fiber.Ctx) error {
	//  PUT /species-image/:speciesId
	speciesId, err := uuid.Parse(c.Params("speciesId"))
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid speciesId",
		}
	}
	fieldLocationId, err := uuid.Parse(c.Params("fieldLocationId"))
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid fieldLocationId",
		}
	}

	file, _ := c.FormFile("attachment")
	if len(file.Header["Content-Type"]) != 1 {
		return &fiber.Error{
			Code:    400,
			Message: "Upload fail",
		}
	}
	if len(strings.Split(file.Header["Content-Type"][0], "/")) < 1 {
		return &fiber.Error{
			Code:    400,
			Message: "Upload fail",
		}
	}
	if strings.Split(file.Header["Content-Type"][0], "/")[0] != "video" {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid file type",
		}
	}

	speciesVideoId := uuid.New()
	videoSource := fmt.Sprintf("%s/%s", constants.SPECIES_VIDEO_BASE_PATH, speciesVideoId.String())

	speciesVideo := types.SpeciesVideoDto{
		SpeciesVideoID:  speciesVideoId,
		SpeciesID:       speciesId,
		FieldLocationID: fieldLocationId,
		VideoSource:     videoSource,
	}
	err = c.SaveFile(file, videoSource)
	if err != nil {
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	_, err = speciesVideoController.SpeciesVideoService.UploadNewFile(speciesVideo)
	if err != nil {
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	return c.JSON(types.BaseResponse{
		Ok:      true,
		Message: "success",
	})
}

func (speciesVideoController *SpeciesVideoController) GetSpeciesVideo(c *fiber.Ctx) error {
	speciesVideoId, err := uuid.Parse(c.Params("speciesVideoId"))
	if err != nil {
		return fiber.ErrInternalServerError
	}
	videoSource := fmt.Sprintf("%s/%s", constants.SPECIES_VIDEO_BASE_PATH, speciesVideoId.String())
	c.Attachment(videoSource)
	return c.Response().SendFile(videoSource)
}
