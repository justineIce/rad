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

//v_sys_role_menu - VIEW
type VSysRoleMenu struct {
	ID          string    `gorm:"column:id" json:"ID" form:"id" query:"id"`                                       // 编号
	ParentID    string    `gorm:"column:parent_id" json:"ParentID" form:"parent_id" query:"parent_id"`            // 父级编号
	RelationIds string    `gorm:"column:relation_ids" json:"RelationId" form:"relation_ids" query:"relation_ids"` // 所有父级编号
	Name        string    `gorm:"column:name" json:"Name" form:"name" query:"name"`                               // 名称
	Sort        float64   `gorm:"column:sort" json:"Sort" form:"sort" query:"sort"`                               // 排序
	Href        string    `gorm:"column:href" json:"Href" form:"href" query:"href"`                               // 链接
	Target      string    `gorm:"column:target" json:"Target" form:"target" query:"target"`                       // 目标
	Icon        string    `gorm:"column:icon" json:"Icon" form:"icon" query:"icon"`                               // 图标
	IsShow      string    `gorm:"column:is_show" json:"IsShow" form:"is_show" query:"is_show"`                    // 是否在菜单中显示
	RouteName   string    `gorm:"column:route_name" json:"RouteName" form:"route_name" query:"route_name"`        // 路由名称
	Component   string    `gorm:"column:component" json:"Component" form:"component" query:"component"`           // 权限标识
	CreatedBy   string    `gorm:"column:created_by" json:"CreatedBy" form:"created_by" query:"created_by"`        // 创建者
	CreatedAt   time.Time `gorm:"column:created_at" json:"CreatedAt" form:"created_at" query:"created_at"`        // 创建时间
	UpdatedBy   string    `gorm:"column:updated_by" json:"UpdatedBy" form:"updated_by" query:"updated_by"`        // 更新者
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"UpdatedAt" form:"updated_at" query:"updated_at"`        // 更新时间
	Remarks     string    `gorm:"column:remarks" json:"Remark" form:"remarks" query:"remarks"`                    // 备注信息
	DeletedAt   time.Time `gorm:"column:deleted_at" json:"DeletedAt" form:"deleted_at" query:"deleted_at"`        // 删除时间
	SysRoleID   string    `gorm:"column:sys_role_id" json:"SysRoleID" form:"sys_role_id" query:"sys_role_id"`     // 角色编号
}

// TableName sets the insert table name for this struct type
func (v *VSysRoleMenu) TableName() string {
	return "v_sys_role_menu"
}
