package service

import (
	"errors"
	"fmt"
	"time"

	"study/StockCalculator/ledger/model"

	"github.com/mssola/user_agent"
	"github.com/wxw9868/util"
	"github.com/wxw9868/util/randomname"
	"gorm.io/gorm"
)

type UserService struct{}

func (us *UserService) Register(username, email, password string) error {
	if !errors.Is(db.Where("email = ?", email).First(&model.User{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("邮箱已存在！")
	}
	if !errors.Is(db.Where("username = ?", username).First(&model.User{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户已存在！")
	}

	password, err := util.NewPassword("ledger").Encrypt(password)
	if err != nil {
		return err
	}

	if err := db.Create(&model.User{Username: username, Nickname: randomname.GenerateName(1), Password: password, Email: email, Avatar: "assets/image/avatar/avatar.png"}).Error; err != nil {
		return fmt.Errorf("注册失败: %s", err)
	}
	return nil
}

type APIUser struct {
	ID       uint
	Email    string `gorm:"column:email;type:string;comment:邮箱"`
	Password string `gorm:"column:password;type:string;comment:账户密码"`
}

func (us *UserService) Login(email, password string) (*model.User, error) {
	password, err := util.NewPassword("ledger").Encrypt(password)
	if err != nil {
		return nil, err
	}
	var user model.User
	result := db.Model(&model.User{}).Where("email = ?", email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户不存在！")
	}
	if password != user.Password {
		return nil, errors.New("用户密码错误！")
	}

	// 报存登录成功的日志信息
	us.LoginLog("userAgent", &model.UserLoginLog{
		InfoId:    user.ID,
		LoginName: user.Username,
		Ipaddr:    "",
		Status:    1,
		Msg:       "登录成功",
		Module:    "系统后台",
	})

	return &user, nil
}

func (us *UserService) ForgotPassword(email, newPassword string) error {
	var user model.User
	if errors.Is(db.Where("email = ?", email).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("邮箱错误！")
	}

	newPassword, _ = util.NewPassword("ledger").Encrypt(newPassword)
	if err := db.Model(&user).Updates(model.User{Password: newPassword}).Error; err != nil {
		return errors.New("修改密码失败！")
	}
	return nil
}

func (us *UserService) GetUserInfo(id uint) (*model.User, error) {
	var user model.User
	if err := db.Where("id =?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (us *UserService) UpdateUser(user model.User) error {
	if err := db.Model(&user).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (us *UserService) UpdateAvatar(id uint, column string, value any) error {
	if err := db.Model(&model.User{}).Where("id = ?", id).Update(column, value).Error; err != nil {
		return err
	}
	return nil
}

func (us *UserService) ChangePassword(id uint, oldPassword, newPassword string) error {
	var user model.User
	if err := db.First(&user, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("用户不存在！")
	}

	password := util.NewPassword("ledger")

	oldPassword, _ = password.Encrypt(oldPassword)
	if oldPassword != user.Password {
		return errors.New("原密码输入错误！")
	}

	newPassword, _ = password.Encrypt(newPassword)
	if err := db.Model(&user).Updates(model.User{Password: newPassword}).Error; err != nil {
		return errors.New("修改密码失败！")
	}
	return nil
}

// LoginLog 记录登录日志
func (as *UserService) LoginLog(userAgent string, loginData *model.UserLoginLog) {
	ua := user_agent.New(userAgent)
	browser, _ := ua.Browser()
	loginData.LoginLocation = GetCityByIp(loginData.Ipaddr)
	loginData.Browser = browser
	loginData.Os = ua.OS()
	loginData.LoginTime = time.Now()

	if err := db.Create(loginData).Error; err != nil {
		fmt.Printf("记录登录日志: %s", err)
	}
}
