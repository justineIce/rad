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

//v_sys_role_menu_btn - VIEW
type VSysRoleMenuBtn struct {
	ID         string    `gorm:"column:id" json:"ID" form:"id" query:"id"`                                   // 编号
	SysMenuID  string    `gorm:"column:sys_menu_id" json:"SysMenuID" form:"sys_menu_id" query:"sys_menu_id"` // 菜单id
	Name       string    `gorm:"column:name" json:"Name" form:"name" query:"name"`                           // 名称
	Permission string    `gorm:"column:permission" json:"Permission" form:"permission" query:"permission"`   // 权限标识
	Remarks    string    `gorm:"column:remarks" json:"Remark" form:"remarks" query:"remarks"`                // 备注信息
	DeletedAt  time.Time `gorm:"column:deleted_at" json:"DeletedAt" form:"deleted_at" query:"deleted_at"`    // 删除时间
	SysRoleID  string    `gorm:"column:sys_role_id" json:"SysRoleID" form:"sys_role_id" query:"sys_role_id"` // 角色编号
	Method     string    `gorm:"column:method" json:"Method" form:"method" query:"method"`                   // 请求方式
	Path       string    `gorm:"column:path" json:"Path" form:"path" query:"path"`                           // 请求路径
	UpdatedBy  string    `gorm:"column:updated_by" json:"UpdatedBy" form:"updated_by" query:"updated_by"`    // 更新者
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"UpdatedAt" form:"updated_at" query:"updated_at"`    // 更新时间
	CreatedAt  time.Time `gorm:"column:created_at" json:"CreatedAt" form:"created_at" query:"created_at"`    // 创建时间
	CreatedBy  string    `gorm:"column:created_by" json:"CreatedBy" form:"created_by" query:"created_by"`    // 创建者

}

// TableName sets the insert table name for this struct type
func (v *VSysRoleMenuBtn) TableName() string {
	return "v_sys_role_menu_btn"
}
