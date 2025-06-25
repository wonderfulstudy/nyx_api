package user

import "errors"

func (r *UserListRequest) Validate() error {
	// 两个都为空：合法
	if r.Page == "" && r.Limit == "" {
		return nil
	}

	// 两个都不为空：合法
	if r.Page != "" && r.Limit != "" {
		return nil
	}

	// 其他情况非法（只传一个、或传了别的字段）
	return errors.New("page 和 limit 参数必须同时存在或同时为空")
}
