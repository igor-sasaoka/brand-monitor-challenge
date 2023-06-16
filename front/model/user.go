package model

import "gorm.io/gorm"

type User struct {
    gorm.Model   
    Name string
    Email string 
    Password string
    Session string
}

type UserRegisterDTO struct {
    UserLoginDTO
    Name string `form:"name" binding:"required"`
    PasswordConfirm string `form:"passwordConfirm"`
}

type UserLoginDTO struct {
    Email string `form:"email" binding:"required"`
    Password string `form:"password" binding:"required"`
}
