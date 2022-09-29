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

type SpeciesImageController struct {
	SpeciesImageService *services.SpeciesImageService
}

func NewSpeciesImageController(speciesImageService *services.SpeciesImageService) *SpeciesImageController {
	return &SpeciesImageController{
		SpeciesImageService: speciesImageService,
	}
}

func (speciesImageController *SpeciesImageController) UploadImage(c *fiber.Ctx) error {
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

	file, err := c.FormFile("attachment")
	if err != nil {
		return &fiber.Error{
			Code:    400,
			Message: "Wrong form key",
		}
	}
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
	if strings.Split(file.Header["Content-Type"][0], "/")[0] != "image" {
		return &fiber.Error{
			Code:    400,
			Message: "Invalid file type",
		}
	}

	speciesImageId := uuid.New()
	imageSource := fmt.Sprintf("%s/%s", constants.SPECIES_IMAGE_BASE_PATH, speciesImageId.String())

	speciesImage := types.SpeciesImageDto{
		SpeciesImageID:  speciesImageId,
		SpeciesID:       speciesId,
		FieldLocationID: fieldLocationId,
		ImageSource:     imageSource,
	}
	err = c.SaveFile(file, imageSource)

	if err != nil {
		return &fiber.Error{
			Code:    500,
			Message: err.Error(),
		}
	}

	_, err = speciesImageController.SpeciesImageService.UploadNewFile(speciesImage)
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

func (speciesImageController *SpeciesImageController) GetSpeciesImage(c *fiber.Ctx) error {
	speciesImageId, err := uuid.Parse(c.Params("speciesImageId"))
	if err != nil {
		return fiber.ErrInternalServerError
	}
	imageSource := fmt.Sprintf("%s/%s", constants.SPECIES_IMAGE_BASE_PATH, speciesImageId.String())
	c.Attachment(imageSource)
	return c.SendFile(imageSource, false)
}
