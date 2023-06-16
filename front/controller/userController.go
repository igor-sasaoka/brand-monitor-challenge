package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	util "github.com/igor-sasaoka/brand-monitor-front"
	"github.com/igor-sasaoka/brand-monitor-front/database"
	"github.com/igor-sasaoka/brand-monitor-front/model"
    "github.com/google/uuid"
)

func errorPage(c *gin.Context, code int, err string) {
    c.HTML(http.StatusBadRequest, "error.html", gin.H{
        "error": err,
    })
}

func Register(c *gin.Context) {
    var formData model.UserRegisterDTO
    //validate form fields
    if err := c.ShouldBind(&formData); err != nil {
        errorPage(c, http.StatusBadRequest, err.Error())
        return
    }
    //validate password confirmation
    if formData.Password != formData.PasswordConfirm {
       errorPage(c, http.StatusBadRequest, "Password confirmation missmatch") 
        return
    }

    //check if email already being used
    var u model.User
    database.DB().First(&u, "email = ?", formData.Email)
    if u.Email != "" {
        errorPage(c, http.StatusConflict, "Email already being used")
        return
    }

    //hash password to store
    h, err := util.HashPassword(formData.Password) 
    if err != nil {
        errorPage(c, http.StatusBadRequest,"Error while storing password, try again later")
        return
    }

    session := uuid.New().String()
    database.DB().Create(&model.User{
        Name:     formData.Name,
        Email:    formData.Email,
        Password: h,
        Session: session,
    })

    c.SetCookie("session", session, 3600, "/", "localhost", false, true)
    c.Redirect(http.StatusFound, "/app/")
}

func Login(c *gin.Context) {
    var formData model.UserLoginDTO
    //validate form fields
    if err := c.ShouldBind(&formData); err != nil {
        errorPage(c, http.StatusBadRequest, err.Error())
        return
    }

    //get user from db
    var u model.User
    database.DB().First(&u, "email = ?", formData.Email)
    if u.Email == "" {
        errorPage(c, http.StatusNotFound, "User not found")
        return
    }
   
    //check password
    if !util.CheckPassword(formData.Password, u.Password) {
        errorPage(c, http.StatusUnauthorized, "Wrong credentials")
        return
    } 

    session := uuid.New().String()
    database.DB().Model(&u).Update("Session", session)
    c.SetCookie("session", session, 3600, "/", "localhost", false, true)
    c.Redirect(http.StatusFound, "/app/")
}
