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

//v_sys_user - VIEW
type VSysUser struct {
	ID             string    `gorm:"column:id" json:"ID" form:"id" query:"id"`                                                       // 编号
	CompanyID      string    `gorm:"column:company_id" json:"CompanyID" form:"company_id" query:"company_id"`                        // 归属公司
	OfficeID       string    `gorm:"column:office_id" json:"OfficeID" form:"office_id" query:"office_id"`                            // 归属部门
	Username       string    `gorm:"column:username" json:"Username" form:"username" query:"username"`                               // 登录名，Table[search]，Valid[AlphaNumeric]
	Password       string    `gorm:"column:password" json:"Password" form:"password" query:"password"`                               // 密码，Table[none]，Form[none]，Valid[AlphaDash]
	Salt           string    `gorm:"column:salt" json:"Salt" form:"salt" query:"salt"`                                               // 盐，Table[none]，Form[none]
	Name           string    `gorm:"column:name" json:"Name" form:"name" query:"name"`                                               // 姓名，Table[search]
	Email          string    `gorm:"column:email" json:"Email" form:"email" query:"email"`                                           // 邮箱，Table[search]，Valid[Email]
	Phone          string    `gorm:"column:phone" json:"Phone" form:"phone" query:"phone"`                                           // 电话，Table[search]，Valid[Phone]
	Mobile         string    `gorm:"column:mobile" json:"Mobile" form:"mobile" query:"mobile"`                                       // 手机，Table[search]，Valid[Mobile]
	Photo          string    `gorm:"column:photo" json:"Photo" form:"photo" query:"photo"`                                           // 用户头像
	LoginIP        string    `gorm:"column:login_ip" json:"LoginIP" form:"login_ip" query:"login_ip"`                                // 最后登陆IP，Form[none]
	LoginDate      time.Time `gorm:"column:login_date" json:"LoginDate" form:"login_date" query:"login_date"`                        // 最后登陆时间，Form[none]
	LoginFlag      string    `gorm:"column:login_flag" json:"LoginFlag" form:"login_flag" query:"login_flag"`                        // 登录标记，Enum[0、不可登录，1、可登录]
	CreatedBy      string    `gorm:"column:created_by" json:"CreatedBy" form:"created_by" query:"created_by"`                        // 创建者
	CreatedAt      time.Time `gorm:"column:created_at" json:"CreatedAt" form:"created_at" query:"created_at"`                        // 创建时间
	UpdatedBy      string    `gorm:"column:updated_by" json:"UpdatedBy" form:"updated_by" query:"updated_by"`                        // 更新者
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"UpdatedAt" form:"updated_at" query:"updated_at"`                        // 更新时间
	Remarks        string    `gorm:"column:remarks" json:"Remark" form:"remarks" query:"remarks"`                                    // 备注信息
	DeletedAt      time.Time `gorm:"column:deleted_at" json:"DeletedAt" form:"deleted_at" query:"deleted_at"`                        // 删除时间
	RelationIds    string    `gorm:"column:relation_ids" json:"RelationId" form:"relation_ids" query:"relation_ids"`                 // 所有父级编号
	OfficeParentID string    `gorm:"column:office_parent_id" json:"OfficeParentID" form:"office_parent_id" query:"office_parent_id"` // 父级编号
	RoleIds        string    `gorm:"column:role_ids" json:"RoleIds" form:"role_ids" query:"role_ids"`
	RoleNames      string    `gorm:"column:role_names" json:"RoleName" form:"role_names" query:"role_names"`
}

// TableName sets the insert table name for this struct type
func (v *VSysUser) TableName() string {
	return "v_sys_user"
}
