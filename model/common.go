package model

import (
	"fmt"
	"regexp"
)

func VerifyMobile(mobile string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobile)
}
func VerifyVfcode(vfcode string) error {
	num := `[0-9]{4}`
	if b, err := regexp.MatchString(num, vfcode); !b || err != nil {
		return fmt.Errorf("[验证码必须纯数字 0_9]!")
	}
	return nil
}
func VerifyPwd(pwd string) error {
	if len(pwd) < 6 || len(pwd) > 16 {
		return fmt.Errorf("[密码长度错误]!")
	}
	//num := `[0-9]{1}`
	//a_z := `[a-z]{1}`
	//A_Z := `[A-Z]{1}`
	//symbol := `[!@#~$%^&*()+|_]{1}`
	//if b, err := regexp.MatchString(num, pwd); !b || err != nil {
	//	return fmt.Errorf("[密码必须包含数字 0_9]!")
	//}
	//if b, err := regexp.MatchString(a_z, pwd); !b || err != nil {
	//	return fmt.Errorf("[密码必须包含小写字母 a_z]!")
	//}
	//if b, err := regexp.MatchString(A_Z, pwd); !b || err != nil {
	//	return fmt.Errorf("[密码必须包含大写字母 A_Z]!")
	//}
	//if b, err := regexp.MatchString(symbol, pwd); !b || err != nil {
	//	return fmt.Errorf("[密码必须包含特殊字符]!")
	//}
	return nil
}
