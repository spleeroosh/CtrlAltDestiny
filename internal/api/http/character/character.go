package character

import (
	"net/http"
	"strconv"
	"time"

	"CtrlAltDestiny/internal/api/http/character/request"
	"CtrlAltDestiny/internal/config"
	"CtrlAltDestiny/internal/entity"

	"CtrlAltDestiny/internal/pkg/log"
	"github.com/gin-gonic/gin"
)

const defaultCacheTTL = 500 * time.Millisecond

type Routes struct {
	uc          Usecases
	log         log.Logger
	middlewares []gin.HandlerFunc
}

func NewRoutes(conf config.Config, uc Usecases, log log.Logger) *Routes {
	return &Routes{
		uc:  uc,
		log: log,
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

	g.GET("/", func(c *gin.Context) { c.Status(http.StatusNoContent) })
	g.GET("/character/:id", r.getCharacter)
	g.POST("/character", r.createCharacter)
	//g.PUT("/character/:id", r.updateCharacter)
	//g.DELETE("/character/:id", r.deleteCharacter)
}

func (r *Routes) getCharacter(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		r.log.Err(err).Msg("invalid character id")
		return
	}

	character, err := r.uc.GetCharacter(c.Request.Context(), id)
	if err != nil {
		r.log.Err(err).Msg("could not get character")
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": character})
}

func (r *Routes) createCharacter(c *gin.Context) {
	var model request.Character
	if err := c.ShouldBindJSON(&model); err != nil {
		r.log.Err(err).Msg("invalid character model")
		return
	}

	err := r.uc.CreateCharacter(c.Request.Context(), entity.Character{
		Age:             model.Age,
		Name:            model.Name,
		UserID:          model.UserID,
		Profession:      model.Profession,
		BurnoutLevel:    0,
		MotivationLevel: 99,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	})
	if err != nil {
		r.log.Err(err).Msg("could not create new character")
		return
	}

	c.Status(http.StatusCreated)
}

//func (r *Routes) updateCharacter(c *gin.Context) {
//	userID, err := strconv.Atoi(c.Param("id"))
//	if err != nil {
//		r.log.Err(err).Msg("invalid user ID")
//		return
//	}
//
//	var model request.User
//	if err = c.ShouldBindJSON(&model); err != nil {
//		r.log.Err(err).Msg("invalid user model")
//		return
//	}
//
//	err = r.manager.UpdateUser(c.Request.Context(), entity.User{
//		ID:     userID,
//		Age:    model.Age,
//		Name:   model.Name,
//		Social: model.Social,
//	})
//	if err != nil {
//		r.log.Err(err).Msg("could not update user")
//		return
//	}
//
//	c.Status(http.StatusNoContent)
//}

//func (r *Routes) deleteCharacter(c *gin.Context) {
//	userID, err := strconv.Atoi(c.Param("id"))
//	if err != nil {
//		r.log.Err(err).Msg("invalid user ID")
//		return
//	}
//
//	if err := r.manager.DeleteUser(c.Request.Context(), userID); err != nil {
//		r.log.Err(err).Msg("could not delete user")
//		return
//	}
//
//	c.Status(http.StatusNoContent)
//}
