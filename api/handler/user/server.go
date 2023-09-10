package user

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"net/http"
	"os"

	"github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"
	"github.com/Seiya-Tagami/Recollect-Service/api/usecase/user"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	LoginUser(c *gin.Context)
	LogoutUser(c *gin.Context)
}

type handler struct {
	userInteractor user.Interactor
}

func New(userInteractor user.Interactor) Handler {
	return &handler{userInteractor}
}

func (h *handler) GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := h.userInteractor.GetUser(id)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h *handler) CreateUser(c *gin.Context) {
	userReq := entity.User{}
	if err := c.BindJSON(&userReq); err != nil {
		panic(err)
	}

	user, err := h.userInteractor.CreateUser(userReq)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h *handler) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	userReq := entity.User{}
	if err := c.BindJSON(&userReq); err != nil {
		panic(err)
	}

	user, err := h.userInteractor.UpdateUser(userReq, id)
	if err != nil {
		panic(err)
	}

	session := sessions.Default(c)
	fmt.Printf("Session: %s\n", session.Get("user_id"))

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h *handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	err := h.userInteractor.DeleteUser(id)
	if err != nil {
		panic(err)
	}

	c.Status(http.StatusNoContent)
}

func (h *handler) LoginUser(c *gin.Context) {
	userReq := entity.User{}
	if err := c.BindJSON(&userReq); err != nil {
		panic(err)
	}

	tokenString, err := h.userInteractor.LoginUser(userReq.UserID, userReq.Password)
	if err != nil {
		panic(err)
	}

	c.SetCookie("reco_cookie", tokenString, 3600, "/", os.Getenv("API_DOMAIN"), false, true)

	c.Status(http.StatusNoContent)
}

func (h *handler) LogoutUser(c *gin.Context) {
	c.SetCookie("reco_cookie", "", 0, "/", os.Getenv("API_DOMAIN"), false, true)

	c.Status(http.StatusNoContent)
}
