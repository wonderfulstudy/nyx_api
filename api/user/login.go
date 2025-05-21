package user

func CheckCredentials(username, password string) (bool, error) {
	var storedPassword string
	err := configs.MYSQL_DB.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil // 用户不存在
		}
		return false, err // 数据库查询错误
	}

	// 校验密码（假设密码是明文存储，实际应使用哈希校验）
	if storedPassword == password {
		return true, nil
	}
	return false, nil
}

func handleLogin(c *gin.Context) {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// 记录请求体
	configs.Logger.Infof("Login request: username=%s, password=%s", loginRequest.Username, loginRequest.Password)

	valid, err := CheckCredentials(loginRequest.Username, loginRequest.Password)
	if err != nil {
		// 记录错误信息
		configs.Logger.Infof("CheckCredentials error: %v", err)
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	if valid {
		c.JSON(200, gin.H{"message": "Login successful"})
	} else {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
	}
}