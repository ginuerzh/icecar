// error
package errors

import (
	"fmt"
)

var (
	NoError             = Error{0, "success"}
	AuthError           = Error{1001, "auth error"}
	UserExistError      = Error{1002, "user exists"}
	AccessError         = Error{1003, "access error"}
	DbError             = Error{1004, "database error"}
	JsonError           = Error{1006, "json data invalid"}
	UserNotFoundError   = Error{1007, "user not found"}
	PasswordError       = Error{1008, "password invalid"}
	InvalidFileError    = Error{1009, "file invalid"}
	HttpError           = Error{1010, "http error"}
	FileNotFoundError   = Error{1011, "file not found"}
	InvalidCarError     = Error{1013, "invalid car"}
	InvalidAddrError    = Error{1014, "address invalid"}
	InvalidMsgError     = Error{1015, "message invalid"}
	DeviceTokenError    = Error{1016, "device token invalid"}
	ReviewNotFoundError = Error{1017, "review not found"}
	InviteCodeError     = Error{1018, "invite code invalid"}
	FileTooLargeError   = Error{1019, "file too large"}
)

type Error struct {
	id  int
	msg string
}

func (e Error) Error() string {
	return fmt.Sprintf("Err: %d, %s", e.id, e.msg)
}

func (e Error) ErrNo() int {
	return e.id
}

func (e Error) ErrMsg() string {
	return e.msg
}

func (e Error) ErrObj() map[string]interface{} {
	return map[string]interface{}{"error_id": e.ErrNo(), "error_desc": e.ErrMsg()}
}
