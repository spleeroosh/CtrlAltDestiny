package user

import (
	"net/http"
	"strconv"
	"time"

	"CtrlAltDestiny/internal/api/http/user/request"
	"CtrlAltDestiny/internal/config"
	"CtrlAltDestiny/internal/entity"

	"CtrlAltDestiny/internal/pkg/log"
	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
)

const defaultCacheTTL = 500 * time.Millisecond

type Routes struct {
	manager     Manager
	log         log.Logger
	middlewares []gin.HandlerFunc
}

func NewRoutes(conf config.Config, manager Manager, log log.Logger) *Routes {
	return &Routes{
		manager: manager,
		log:     log,
		//middlewares: []gin.HandlerFunc{
		//	jwt.New(
		//		jwt.HMACSecret([]byte(conf.App.AuthSecret)),
		//		jwt.Logger(logger),
		//	),
		//	ratelimit.New(
		//		ratelimit.NonAuthMaxRPS(conf.App.RateLimitRPS),
		//		ratelimit.NonAuthBurst(conf.App.RateLimitBurst),
		//	),
		//},
		//handle: ginjson.NewErrorHandler(logger),
	}
}

func (r *Routes) Apply(e *gin.Engine) {
	g := e.Group("/api/v1")
	g.Use(r.middlewares...)

	cacheStore := persist.NewMemoryStore(defaultCacheTTL)

	g.GET("/", func(c *gin.Context) { c.Status(http.StatusNoContent) })
	g.GET("/users/:id", cache.CacheByRequestURI(cacheStore, 0), r.getUser)
	g.POST("/users", r.createUser)
	g.PUT("/users/:id", r.updateUser)
	g.DELETE("/users/:id", r.deleteUser)
}

func (r *Routes) getUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		r.log.Err(err).Msg("invalid user id")
		return
	}

	user, err := r.manager.GetUser(c.Request.Context(), userID)
	if err != nil {
		r.log.Err(err).Msg("could not get user")
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (r *Routes) createUser(c *gin.Context) {
	var model request.User
	if err := c.ShouldBindJSON(&model); err != nil {
		r.log.Err(err).Msg("invalid user model")
		return
	}

	err := r.manager.CreateUser(c.Request.Context(), entity.User{
		Age:    model.Age,
		Name:   model.Name,
		Social: model.Social,
	})
	if err != nil {
		r.log.Err(err).Msg("could not create new user")
		return
	}

	c.Status(http.StatusCreated)
}

func (r *Routes) updateUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		r.log.Err(err).Msg("invalid user ID")
		return
	}

	var model request.User
	if err = c.ShouldBindJSON(&model); err != nil {
		r.log.Err(err).Msg("invalid user model")
		return
	}

	err = r.manager.UpdateUser(c.Request.Context(), entity.User{
		ID:     userID,
		Age:    model.Age,
		Name:   model.Name,
		Social: model.Social,
	})
	if err != nil {
		r.log.Err(err).Msg("could not update user")
		return
	}

	c.Status(http.StatusNoContent)
}

func (r *Routes) deleteUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		r.log.Err(err).Msg("invalid user ID")
		return
	}

	if err := r.manager.DeleteUser(c.Request.Context(), userID); err != nil {
		r.log.Err(err).Msg("could not delete user")
		return
	}

	c.Status(http.StatusNoContent)
}
