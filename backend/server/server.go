package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/wichadak/eDNA/config"
	"github.com/wichadak/eDNA/controllers"
	"github.com/wichadak/eDNA/database"
	"github.com/wichadak/eDNA/middlewares"
	"github.com/wichadak/eDNA/repositories"
	"github.com/wichadak/eDNA/services"
	"github.com/wichadak/eDNA/types"
	"github.com/wichadak/eDNA/utils"
)

func StartServer() {
	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(logger.New())

	config := config.ReadConfig()
	dbConnection := database.NewPgConnection(&config.Postgres)
	redisConnection := database.NewRedisConnection(&config.Redis)
	defer dbConnection.Close()
	defer redisConnection.Close()

	mockDB := database.NewMockDB()

	authMiddleware := middlewares.NewAuthMiddleware(redisConnection)

	userRepository := repositories.NewUserRepository(dbConnection)
	userService := services.NewUserService(userRepository, redisConnection)
	userController := controllers.NewUserController(userService)
	authService := services.NewAuthService(userRepository, redisConnection)
	authController := controllers.NewAuthController(userService, authService)

	kingdomRepository := repositories.NewKingdomRepository(dbConnection)
	kingdomService := services.NewKingdomService(kingdomRepository)
	kingdomController := controllers.NewKingdomController(kingdomService)

	phylumRepository := repositories.NewPhylumRepository(dbConnection)
	phylumService := services.NewPhylumService(phylumRepository)
	phylumController := controllers.NewPhylumController(phylumService)

	orderRepository := repositories.NewOrderRepository(dbConnection)
	orderService := services.NewOrderService(orderRepository)
	orderController := controllers.NewOrderController(orderService)

	classRepository := repositories.NewClassRepository(dbConnection)
	classService := services.NewClassService(classRepository)
	classController := controllers.NewClassController(classService)

	familyRepository := repositories.NewFamilyRepository(dbConnection)
	familyService := services.NewFamilyService(familyRepository)
	familyController := controllers.NewFamilyController(familyService)

	genusRepository := repositories.NewGenusRepository(dbConnection)
	genusService := services.NewGenusService(genusRepository)
	genusController := controllers.NewGenusController(genusService)

	speciesRepository := repositories.NewSpeciesRepository(dbConnection)
	speciesService := services.NewSpeciesService(speciesRepository)
	speciesController := controllers.NewSpeciesController(speciesService)

	summaryRepository := repositories.NewSummaryRepository(dbConnection)
	summaryService := services.NewSummaryService(summaryRepository)
	summaryController := controllers.NewSummaryController(summaryService)

	evennessRepository := repositories.NewEvennessRepository(dbConnection)
	evennessService := services.NewEvennessService(evennessRepository)
	evennessController := controllers.NewEvennessController(evennessService)

	diversityRepository := repositories.NewDiversityRepository(dbConnection)
	diversityService := services.NewDiversityService(diversityRepository)
	diversityController := controllers.NewDiversityController(diversityService)

	densityRepository := repositories.NewDensityRepository(dbConnection)
	densityService := services.NewDensityService(densityRepository)
	densityController := controllers.NewDensityController(densityService)

	// densityRepository := repositories.NewDensityRepository(dbConnection)
	// densityService := services.NewDensityService(densityRepository)
	// densityController := controllers.NewDensityController(densityService)

	// evennessRepository := repositories.NewEvennessRepository(dbConnection)
	// evennessService := services.NewEvennessService(evennessRepository)
	// evennessController := controllers.NewEvennessController(evennessService)

	// diversityRepository := repositories.NewDiversityRepository(dbConnection)
	// diversityService := services.NewDiversityService(diversityRepository)
	// diversityController := controllers.NewDiversityController(diversityService)

	// speciesImageRepository := repositories.NewSpeciesImageRepository(dbConnection)
	// speciesImageService := services.NewSpeciesImageService(speciesImageRepository)
	// speciesImageController := controllers.NewSpeciesImageController(speciesImageService)

	// speciesVideoRepository := repositories.NewSpeciesVideoRepository(dbConnection)
	// speciesVideoService := services.NewSpeciesVideoService(speciesVideoRepository)
	// speciesVideoController := controllers.NewSpeciesVideoController(speciesVideoService)

	// siteRepository := repositories.NewSiteRepository(dbConnection)
	// siteService := services.NewSiteService(siteRepository)
	// siteController := controllers.NewSiteController(siteService)

	// fieldLocationRepository := repositories.NewFieldLocationRepository(dbConnection)
	// fieldLocationService := services.NewFieldLocationService(fieldLocationRepository)
	// fieldLocationController := controllers.NewFieldLocationController(fieldLocationService)

	// mainStationRepository := repositories.NewMainStationRepository(dbConnection)
	// mainStationService := services.NewMainStationService(mainStationRepository)
	// mainStationController := controllers.NewMainStationController(mainStationService)

	speciesTypeController := controllers.NewSpeciesTypeController()

	occurrenceRepository := repositories.NewOccurrenceRepository(mockDB)
	occurrenceService := services.NewOccurrenceService(occurrenceRepository)
	occurrenceController := controllers.NewOccurrenceController(occurrenceService)

	majorGroupRepository := repositories.NewMajorGroupRepository(dbConnection)
	majorGroupService := services.NewMajorGroupService(majorGroupRepository)
	majorGroupController := controllers.NewMajorGroupController(majorGroupService)

	assetRepository := repositories.NewAssetRepository(dbConnection)
	assetService := services.NewAssetService(assetRepository)
	assetController := controllers.NewAssetController(assetService)

	platformRepository := repositories.NewPlatformRepository(dbConnection)
	platformService := services.NewPlatformService(platformRepository)
	platformController := controllers.NewPlatformController(platformService)

	stationRepository := repositories.NewStationRepository(dbConnection)
	stationService := services.NewStationService(stationRepository)
	stationController := controllers.NewStationController(stationService)

	locationRepository := repositories.NewLocationRepository(dbConnection)
	locationService := services.NewLocationService(locationRepository)
	locationController := controllers.NewLocationController(locationService)

	// projectRepository := repositories.NewProjectRepository(dbConnection)
	// projectService := services.NewProjectService(projectRepository)
	// projectController := controllers.NewProjectController(projectService)

	// platformRepository := repositories.NewPlatformRepository(dbConnection)
	// platformService := services.NewPlatformService(platformRepository)
	// platformController := controllers.NewPlatformController(platformService)

	app.Get("/user/:userID", userController.GetUserByID)
	app.Post("/user", userController.CreateUser)

	app.Post("/register", authController.Register)
	app.Post("/login", authController.Login)
	app.Get("/me", authMiddleware.AuthenticationRequired, authController.GetMe)

	healthCheckAPI := app.Group("/healthz")
	healthCheckAPI.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(types.BaseResponse{
			Ok:      true,
			Message: "Healthy",
		})
	})

	kingdomAPI := app.Group("/kingdoms")
	kingdomAPI.Get("/", kingdomController.ListKingdom)

	phylumAPI := app.Group("/phylums")
	phylumAPI.Get("/", phylumController.ListPhylum)

	orderAPI := app.Group("/orders")
	orderAPI.Get("/", orderController.ListOrder)

	classAPI := app.Group("/classes")
	classAPI.Get("/", classController.ListClass)

	familyAPI := app.Group("/families")
	familyAPI.Get("/", familyController.ListFamily)

	genusAPI := app.Group("/genus")
	genusAPI.Get("/", genusController.ListGenus)

	speciesAPI := app.Group("/species") //pls fix
	speciesAPI.Get("/", speciesController.ListSpecies)
	// speciesAPI.Get("/count", speciesController.GetSpeciesCountGroupBySpeciesType)
	speciesAPI.Get("/:speciesId/details", speciesController.GetSpeciesDetails)

	// densityAPI := app.Group("/densities")
	// densityAPI.Get("/", densityController.ListDensity)

	// evennessAPI := app.Group("/evenness")
	// evennessAPI.Get("/", evennessController.ListEvenness)
	// // evennessAPI.Get("/year", evennessController.ListEvennessYear)
	// // evennessAPI.Get("/project", evennessController.ListEvennessProject)

	// diversityAPI := app.Group("/diversity")
	// diversityAPI.Get("/", diversityController.ListDiversity)

	// speciesImageAPI := app.Group("/species-image") //no
	// speciesImageAPI.Get("/:speciesImageId", speciesImageController.GetSpeciesImage)
	// speciesImageAPI.Put("/upload/:speciesId/:fieldLocationId", speciesImageController.UploadImage)

	// speciesVideoAPI := app.Group("/species-video") //no
	// speciesVideoAPI.Get("/:speciesVideoId", speciesVideoController.GetSpeciesVideo)
	// speciesVideoAPI.Put("/upload/:speciesId/:fieldLocationId", speciesVideoController.UploadVideo)

	// siteAPI := app.Group("/sites") //project
	// siteAPI.Get("/", siteController.ListSite)

	// fieldLocationAPI := app.Group("/field-locations") //platform
	// fieldLocationAPI.Get("/", fieldLocationController.ListFieldLocation)

	// mainStationAPI := app.Group("/main-stations") //site
	// mainStationAPI.Get("/", mainStationController.ListMainStation)

	speciesTypeAPI := app.Group("/species-types") //mockup
	speciesTypeAPI.Get("/", speciesTypeController.ListSpeciesType)

	occurrenceAPI := app.Group("/occurrences") //mockup
	occurrenceAPI.Get("/", occurrenceController.ListOccurrence)

	majorGroupAPI := app.Group("/major-groups")
	majorGroupAPI.Get("/", majorGroupController.ListMajorGroup)

	assetAPI := app.Group("/asset")
	assetAPI.Get("/", assetController.ListAsset)

	platformAPI := app.Group("/platform")
	platformAPI.Get("/", platformController.ListPlatform)

	stationAPI := app.Group("/station")
	stationAPI.Get("/", stationController.ListStation)

	locationAPI := app.Group("/location")
	locationAPI.Get("/", locationController.ListLocation)

	summaryAPI := app.Group("/summary")
	summaryAPI.Get("/", summaryController.ListSummary)

	evennessAPI := app.Group("/evenness")
	evennessAPI.Get("/", evennessController.ListEvenness)

	diversityAPI := app.Group("/diversity")
	diversityAPI.Get("/", diversityController.ListDiversity)

	densityAPI := app.Group("/density")
	densityAPI.Get("/", densityController.ListDensity)

	// projectAPI := app.Group("/project")
	// projectAPI.Get("/", projectController.ListProject)

	// platformAPI := app.Group("/platform")
	// platformAPI.Get("/", platformController.ListPlatform)

	app.Listen(fmt.Sprintf(":%d", config.Application.Port))
}
