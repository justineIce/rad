package model

import (
	"database/sql"
	"github.com/guregu/null"
	"time"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

//sys_user - 用户表
type SysUser struct {
	ID        string      `gorm:"column:id" json:"ID" form:"id" query:"id" valid:"Required;MaxSize(64)" validAlias:"编号"`                                       // 编号
	CompanyID string      `gorm:"column:company_id" json:"CompanyID" form:"company_id" query:"company_id" valid:"Required;MaxSize(64)" validAlias:"归属公司"`      // 归属公司
	OfficeID  string      `gorm:"column:office_id" json:"OfficeID" form:"office_id" query:"office_id" valid:"Required;MaxSize(64)" validAlias:"归属部门"`          // 归属部门
	Username  string      `gorm:"column:username" json:"Username" form:"username" query:"username" valid:"Required;MaxSize(50);AlphaNumeric" validAlias:"登录名"` // 登录名，Table[search]，Valid[AlphaNumeric]
	Password  string      `gorm:"column:password" json:"Password" form:"password" query:"password" valid:"Required;MaxSize(50);AlphaDash" validAlias:"密码"`     // 密码，Table[none]，Form[none]，Valid[AlphaDash]
	Salt      string      `gorm:"column:salt" json:"Salt" form:"salt" query:"salt" valid:"Required;MaxSize(10)" validAlias:"盐"`                                // 盐，Table[none]，Form[none]
	Name      string      `gorm:"column:name" json:"Name" form:"name" query:"name" valid:"Required;MaxSize(50)" validAlias:"姓名"`                               // 姓名，Table[search]
	Email     null.String `gorm:"column:email" json:"Email" form:"email" query:"email" valid:"MaxSize(100);Email" validAlias:"邮箱"`                             // 邮箱，Table[search]，Valid[Email]
	Phone     null.String `gorm:"column:phone" json:"Phone" form:"phone" query:"phone" valid:"MaxSize(100);Phone" validAlias:"电话"`                             // 电话，Table[search]，Valid[Phone]
	Mobile    null.String `gorm:"column:mobile" json:"Mobile" form:"mobile" query:"mobile" valid:"MaxSize(100);Mobile" validAlias:"手机"`                        // 手机，Table[search]，Valid[Mobile]
	Photo     null.String `gorm:"column:photo" json:"Photo" form:"photo" query:"photo" valid:"MaxSize(100)" validAlias:"用户头像"`                                 // 用户头像
	LoginIP   null.String `gorm:"column:login_ip" json:"LoginIP" form:"login_ip" query:"login_ip" valid:"MaxSize(100)" validAlias:"最后登陆IP"`                    // 最后登陆IP，Form[none]
	LoginDate null.Time   `gorm:"column:login_date" json:"LoginDate" form:"login_date" query:"login_date" validAlias:"最后登陆时间"`                                 // 最后登陆时间，Form[none]
	LoginFlag string      `gorm:"column:login_flag" json:"LoginFlag" form:"login_flag" query:"login_flag" valid:"Required;MaxSize(1)" validAlias:"登录标记"`       // 登录标记，Enum[0、不可登录，1、可登录]
	CreatedBy string      `gorm:"column:created_by" json:"CreatedBy" form:"created_by" query:"created_by" valid:"Required;MaxSize(64)" validAlias:"创建者"`       // 创建者
	CreatedAt time.Time   `gorm:"column:created_at" json:"CreatedAt" form:"created_at" query:"created_at" validAlias:"创建时间"`                                   // 创建时间
	UpdatedBy string      `gorm:"column:updated_by" json:"UpdatedBy" form:"updated_by" query:"updated_by" valid:"Required;MaxSize(64)" validAlias:"更新者"`       // 更新者
	UpdatedAt time.Time   `gorm:"column:updated_at" json:"UpdatedAt" form:"updated_at" query:"updated_at" validAlias:"更新时间"`                                   // 更新时间
	Remarks   null.String `gorm:"column:remarks" json:"Remark" form:"remarks" query:"remarks" valid:"MaxSize(100)" validAlias:"备注信息"`                          // 备注信息
	DeletedAt null.Time   `gorm:"column:deleted_at" json:"DeletedAt" form:"deleted_at" query:"deleted_at" validAlias:"删除时间"`                                   // 删除时间

}

// TableName sets the insert table name for this struct type
func (s *SysUser) TableName() string {
	return "sys_user"
}
