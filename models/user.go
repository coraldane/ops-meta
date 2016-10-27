package models

import (
	"github.com/coraldane/ops-meta/db"
	"github.com/coraldane/ops-meta/utils"
	"github.com/toolkits/logger"
)

type User struct {
	Id            int64
	UserName      string `form:"userName"`
	LoginPwd      string `form:"loginPwd"`
	RealName      string `form:"realName"`
	PhoneNo       string `form:"phoneNo"`
	Email         string `form:"email"`
	RoleName      string `form:"roleName" orm:"default(NORMAL)"`
	AccountStatus int8   `orm:"default(1)"`
}

func (u *User) TableUnique() [][]string {
	return [][]string{
		[]string{"UserName"},
	}
}

func (this *User) Insert() (int64, error) {
	this.LoginPwd = utils.Md5Hex(this.LoginPwd)
	this.RoleName = "NORMAL"
	this.AccountStatus = 1
	return db.NewOrm().Insert(this)
}

func (this *User) CheckExists() bool {
	var rowCount int
	strSql := `select count(*) from t_user where user_name=? `
	db.NewOrm().Raw(strSql, this.UserName).QueryRow(&rowCount)
	return rowCount > 0
}

func (this *User) Update() (int64, error) {
	strSql := `update t_user set real_name=?, phone_no=?, email=?, role_name=? where id=?`
	result, err := db.NewOrm().Raw(strSql, this.RealName, this.PhoneNo, this.Email, this.RoleName, this.Id).Exec()
	if nil != err {
		logger.Errorln("update error", err)
		return 0, err
	}
	return result.RowsAffected()
}

func (this *User) ChangeLoginPasswd() (int64, error) {
	result, err := db.NewOrm().Raw(`update t_user set login_pwd=? where user_name=?`, utils.Md5Hex(this.LoginPwd), this.UserName).Exec()
	if nil != err {
		return 0, err
	}
	return result.RowsAffected()
}

func CheckLogin(userName, loginPwd string) (*User, error) {
	var u User
	strSql := `select id, user_name, real_name, phone_no, email, role_name, account_status from t_user where user_name=? and login_pwd=? and account_status=1`
	err := db.NewOrm().Raw(strSql, userName, utils.Md5Hex(loginPwd)).QueryRow(&u)
	return &u, err
}

func GetUserById(userId int64) *User {
	var u User
	strSql := `select id, user_name, real_name, phone_no, email, role_name, account_status from t_user where id=?`
	err := db.NewOrm().Raw(strSql, userId).QueryRow(&u)
	if nil != err {
		logger.Errorln("query error", err)
	}
	return &u
}

func QueryUserList(queryDto QueryUserDto, pageInfo *PageInfo) ([]User, *PageInfo) {
	var rows []User
	query := db.NewOrm().QueryTable(User{})
	if "" != queryDto.UserName {
		query = query.Filter("user_name__contains", queryDto.UserName)
	}
	if "" != queryDto.RealName {
		query = query.Filter("RealName", queryDto.RealName)
	}
	if "" != queryDto.RoleName {
		query = query.Filter("RoleName", queryDto.RoleName)
	}

	rowCount, err := query.Count()
	if nil != err {
		logger.Errorln("queryCount error", err)
		pageInfo.SetRowCount(0)
		return nil, pageInfo
	}
	pageInfo.SetRowCount(rowCount)

	_, err = query.OrderBy("Id").Offset(pageInfo.GetStartIndex()).Limit(pageInfo.PageSize).All(&rows,
		"UserName", "RealName", "PhoneNo", "Email", "RoleName", "AccountStatus")
	if nil != err {
		logger.Errorln("QueryUserList error", err)
	}
	return rows, pageInfo
}
