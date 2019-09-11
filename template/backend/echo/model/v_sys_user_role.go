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

//v_sys_user_role - VIEW
type VSysUserRole struct {
	ID          string    `gorm:"column:id" json:"ID" form:"id" query:"id"`                                           // 编号
	SysOfficeID string    `gorm:"column:sys_office_id" json:"SysOfficeID" form:"sys_office_id" query:"sys_office_id"` // 归属机构
	Name        string    `gorm:"column:name" json:"Name" form:"name" query:"name"`                                   // 角色名称
	RoleType    string    `gorm:"column:role_type" json:"RoleType" form:"role_type" query:"role_type"`                // 角色类型，用于数据权限区分，2、超级管理员，1、管理员，0、普通账户
	CreatedBy   string    `gorm:"column:created_by" json:"CreatedBy" form:"created_by" query:"created_by"`            // 创建者
	CreatedAt   time.Time `gorm:"column:created_at" json:"CreatedAt" form:"created_at" query:"created_at"`            // 创建时间
	UpdatedBy   string    `gorm:"column:updated_by" json:"UpdatedBy" form:"updated_by" query:"updated_by"`            // 更新者
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"UpdatedAt" form:"updated_at" query:"updated_at"`            // 更新时间
	Remarks     string    `gorm:"column:remarks" json:"Remark" form:"remarks" query:"remarks"`                        // 备注信息
	DeletedAt   time.Time `gorm:"column:deleted_at" json:"DeletedAt" form:"deleted_at" query:"deleted_at"`            // 删除时间
	SysUserID   string    `gorm:"column:sys_user_id" json:"SysUserID" form:"sys_user_id" query:"sys_user_id"`         // 用户编号

}

// TableName sets the insert table name for this struct type
func (v *VSysUserRole) TableName() string {
	return "v_sys_user_role"
}
