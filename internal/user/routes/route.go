package routes

import (
	"github.com/dwiprastyoisworo/go-restapi-plg-stack-temp/internal/user/handlers"
	repositories2 "github.com/dwiprastyoisworo/go-restapi-plg-stack-temp/internal/user/repositories"
	"github.com/dwiprastyoisworo/go-restapi-plg-stack-temp/internal/user/usecases"
	"github.com/dwiprastyoisworo/go-restapi-plg-stack-temp/lib/configs"
	"github.com/dwiprastyoisworo/go-restapi-plg-stack-temp/lib/models"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gorm.io/gorm"
)

type Route struct {
	db         *gorm.DB
	i18nBundle *i18n.Bundle
	ctx        *gin.Engine
	cfg        *configs.AppConfig
}

func NewUserRoute(db *gorm.DB, i18nBundle *i18n.Bundle, ctx *gin.Engine, cfg *configs.AppConfig) *Route {
	return &Route{db: db, i18nBundle: i18nBundle, ctx: ctx, cfg: cfg}
}

func (r *Route) RouteInit() {
	r.UserRouteInit()
}

func (r *Route) UserRouteInit() {
	group := r.ctx.Group("/user")
	genericUserRepo := repositories2.NewRepository[models.Users]()
	userRepo := repositories2.NewUserRepository()
	userUsecase := usecases.NewUserUsecase(genericUserRepo, userRepo, r.db)
	userHandler := handlers.NewUserHandler(userUsecase, r.i18nBundle, r.cfg)
	group.POST("/register", userHandler.Register)
}
