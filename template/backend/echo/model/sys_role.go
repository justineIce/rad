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

//sys_role - 角色表
type SysRole struct {
	ID          string      `gorm:"column:id" json:"ID" form:"id" query:"id" valid:"Required;MaxSize(64)" validAlias:"编号"`                                             // 编号
	SysOfficeID string      `gorm:"column:sys_office_id" json:"SysOfficeID" form:"sys_office_id" query:"sys_office_id" valid:"Required;MaxSize(64)" validAlias:"归属机构"` // 归属机构
	Name        string      `gorm:"column:name" json:"Name" form:"name" query:"name" valid:"Required;MaxSize(100)" validAlias:"角色名称"`                                  // 角色名称
	RoleType    string      `gorm:"column:role_type" json:"RoleType" form:"role_type" query:"role_type" valid:"Required;MaxSize(1)" validAlias:"角色类型"`                 // 角色类型，用于数据权限区分，2、超级管理员，1、管理员，0、普通账户
	CreatedBy   string      `gorm:"column:created_by" json:"CreatedBy" form:"created_by" query:"created_by" valid:"Required;MaxSize(64)" validAlias:"创建者"`             // 创建者
	CreatedAt   time.Time   `gorm:"column:created_at" json:"CreatedAt" form:"created_at" query:"created_at" validAlias:"创建时间"`                                         // 创建时间
	UpdatedBy   string      `gorm:"column:updated_by" json:"UpdatedBy" form:"updated_by" query:"updated_by" valid:"Required;MaxSize(64)" validAlias:"更新者"`             // 更新者
	UpdatedAt   time.Time   `gorm:"column:updated_at" json:"UpdatedAt" form:"updated_at" query:"updated_at" validAlias:"更新时间"`                                         // 更新时间
	Remarks     null.String `gorm:"column:remarks" json:"Remarks" form:"remarks" query:"remarks" valid:"MaxSize(100)" validAlias:"备注信息"`                               // 备注信息
	DeletedAt   null.Time   `gorm:"column:deleted_at" json:"DeletedAt" form:"deleted_at" query:"deleted_at" validAlias:"删除时间"`                                         // 删除时间

}

// TableName sets the insert table name for this struct type
func (s *SysRole) TableName() string {
	return "sys_role"
}
