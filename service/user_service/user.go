package user

import (
	"errors"
	"fmt"
	"nyx_api/models"
	"nyx_api/pkg/aes"
	"nyx_api/pkg/redis"
	"nyx_api/pkg/setting"
	"nyx_api/util"
	"strconv"
	"time"

	wallet "nyx_api/service/wallet_service"

	"github.com/gin-gonic/gin"
)

func CreateUserService(req CreateRequest) error {
	password, err := aes.AesEncryptCBCBase64(setting.DefaultPassword)
	if err != nil {
		return errors.New("加密默认密码失败")
	}

	user := models.User{
		Uuid:     util.GenerateStringUUID(),
		Username: req.Username,
		Name:     req.Name,
		Password: password,
		RoleId:   2,
	}

	if err := models.AddUser(&user); err != nil {
		return errors.New("创建用户失败：" + err.Error())
	}

	return nil
}

func UpdateUserService(req UpdateRequest) error {
	user := models.User{
		Uuid:         req.Uuid,
		Name:         req.Name,
		Avatar:       req.Avatar,
		Introduction: req.Introduction,
		RoleId:       req.RoleId,
		Phone:        req.Phone,
		Address:      req.Address,
	}

	if err := models.UpdateUser(&user); err != nil {
		return errors.New("更新用户信息失败：" + err.Error())
	}

	return nil
}

func DeleteUserService(req DeleteRequest) error {
	if err := models.DeleteUser(req.Uuid); err != nil {
		return errors.New("更新用户信息失败：" + err.Error())
	}

	return nil
}

func LoginService(req LoginRequest) (string, error) {
	// 在数据库中查询用户
	user, err := models.GetByUsername(req.Username)
	if err != nil {
		return "", errors.New("数据库查询用户失败：" + err.Error())
	}

	// 将数据库中的密码密文和请求体中的密码密文分别解析
	dbPwd, err := aes.AesDecryptCBCBase64(user.Password)
	if err != nil {
		return "", errors.New("数据库中用户密码解密失败" + err.Error())
	}

	// 比较两个密码是否一致
	if dbPwd != req.Password {
		return "", errors.New("用户名或密码不匹配")
	}

	// 密码校验通过后用用户名和密码生成一个用户token，并存进redis中
	token, err := util.GenerateToken(user.Uuid, user.Username, user.Phone)
	if err != nil {
		return "", errors.New("生成用户token失败" + err.Error())
	}
	result := redis.RDB.Set(redis.CTX, "user:"+req.Username, token, 3*time.Hour)
	if result.Err() != nil {
		return "", errors.New("输出用户token到redis失败" + result.Err().Error())
	}
	return token, nil
}

func ListService(req UserListRequest) (UserListResponse, error) {
	var response UserListResponse
	page, err := strconv.Atoi(req.Page)
	if err != nil {
		return UserListResponse{}, errors.New("请求体中的分页参数page不是数值类型")
	}
	limit, err := strconv.Atoi(req.Limit)
	if err != nil {
		return UserListResponse{}, errors.New("请求体中的分页参数limit不是数值类型")
	}
	users, err := models.ListUsers(page, limit)
	if err != nil {
		return UserListResponse{}, errors.New("从数据库获取所有用户数据失败" + err.Error())
	}

	for _, user := range users {
		role, err := models.GetRoleById(user.RoleId)
		if err != nil {
			return UserListResponse{}, errors.New("从数据库中获取用户角色数据失败" + err.Error())
		}
		walletInfo, _ := wallet.WalletInfoService(user.Uuid)
		response.Users = append(response.Users, InfoResponse{
			Uuid:         user.Uuid,
			Username:     user.Username,
			Name:         user.Name,
			Avatar:       user.Avatar,
			Introduction: user.Introduction,
			Roles:        []string{role.KeyName},
			Phone:        user.Phone,
			Address:      user.Address,
			WorkersCount: models.GetWorkerCountByOnweruser(user.Id),
			WalletInfo:   walletInfo,
		})
	}
	response.Total = len(users)

	return response, nil
}

func InfoService(uuid string) (InfoResponse, error) {
	var response InfoResponse
	user, err := models.GetByUuid(uuid)
	if err != nil {
		return InfoResponse{}, errors.New("从数据库中获取用户信息失败" + err.Error())
	}
	role, err := models.GetRoleById(user.RoleId)
	if err != nil {
		return InfoResponse{}, errors.New("从数据库中获取用户角色信息失败" + err.Error())
	}

	response = InfoResponse{
		Uuid:     user.Uuid,
		Username: user.Username,
		Name:     user.Name,
		Phone:    user.Phone,
		Address:  user.Address,
		Avatar:   user.Avatar,
		Roles:    []string{role.KeyName},
	}

	return response, nil
}

func LogoutService(c *gin.Context) error {
	var (
		username             any
		phone                any
		contextUsernameExist bool
		contextPhoneExist    bool
		redisUsernameExist   int64
		redisPhoneExist      int64
		redisUsernameErr     error
		redisPhoneErr        error
	)
	username, contextUsernameExist = c.Get("username")
	phone, contextPhoneExist = c.Get("phone")
	if !contextUsernameExist && !contextPhoneExist {
		return errors.New("获取token中的用户信息失败")
	}

	if contextUsernameExist {
		redisUsernameExist, redisUsernameErr = redis.RDB.Exists(redis.CTX, "user:"+username.(string)).Result()
		if redisUsernameErr != nil {
			res := fmt.Sprintf("检查redis中用户名缓存失败, 错误信息：%v", redisUsernameErr)
			return errors.New(res)
		}

		if redisUsernameExist == 1 {
			_, err := redis.RDB.Del(redis.CTX, "user:"+username.(string)).Result()
			if err != nil {
				return errors.New("删除用户名缓存失败, 错误信息：" + err.Error())
			}
		}
	}

	if contextPhoneExist {
		redisPhoneExist, redisPhoneErr = redis.RDB.Exists(redis.CTX, "user:"+phone.(string)).Result()
		if redisPhoneErr != nil {
			res := fmt.Sprintf("检查redis中电话号码缓存失败, 错误信息：%v", redisPhoneErr)
			return errors.New(res)
		}

		if redisPhoneExist == 1 {
			_, err := redis.RDB.Del(redis.CTX, "user:"+username.(string)).Result()
			if err != nil {
				return errors.New("删除用户电话号码缓存失败, 错误信息：" + err.Error())
			}
		}
	}

	return nil
}
