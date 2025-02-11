package errors

import "gin-blog/pkg/globals"

var (
	ErrTokenAuth       = NewAppeError(globals.CodeClientErr, "token验证不合法")
	ErrEmailOrPassword = NewAppeError(globals.CodeClientErr, "账号或密码错误")
	ErrUserHasExist    = NewAppeError(globals.CodeClientErr, "用户已存在")
	ErrPermissions     = NewAppeError(globals.CodeServerErr, "权限认证错误")
)
