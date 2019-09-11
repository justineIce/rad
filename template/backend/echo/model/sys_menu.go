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

//sys_menu - 菜单表
type SysMenu struct {
	ID          string      `gorm:"column:id" json:"ID" form:"id" query:"id" valid:"Required;MaxSize(64)" validAlias:"编号"`                                            // 编号
	ParentID    string      `gorm:"column:parent_id" json:"ParentID" form:"parent_id" query:"parent_id" valid:"Required;MaxSize(64)" validAlias:"父级编号"`               // 父级编号
	RelationIds string      `gorm:"column:relation_ids" json:"RelationId" form:"relation_ids" query:"relation_ids" valid:"Required;MaxSize(100)" validAlias:"所有父级编号"` // 所有父级编号
	Name        string      `gorm:"column:name" json:"Name" form:"name" query:"name" valid:"Required;MaxSize(100)" validAlias:"名称"`                                   // 名称
	Sort        null.Float  `gorm:"column:sort" json:"Sort" form:"sort" query:"sort" validAlias:"排序"`                                                                 // 排序
	Href        null.String `gorm:"column:href" json:"Href" form:"href" query:"href" valid:"MaxSize(100)" validAlias:"链接"`                                            // 链接
	Target      null.String `gorm:"column:target" json:"Target" form:"target" query:"target" valid:"MaxSize(100)" validAlias:"目标"`                                    // 目标
	Icon        null.String `gorm:"column:icon" json:"Icon" form:"icon" query:"icon" valid:"MaxSize(100)" validAlias:"图标"`                                            // 图标
	RouteName   string      `gorm:"column:route_name" json:"RouteName" form:"route_name" query:"route_name" valid:"Required;MaxSize(50)" validAlias:"路由名称"`           // 权限标识
	Component   string      `gorm:"column:component" json:"Component" form:"component" query:"component" valid:"Required;MaxSize(255)" validAlias:"路由名称"`             // 权限标识
	CreatedBy   string      `gorm:"column:created_by" json:"CreatedBy" form:"created_by" query:"created_by" valid:"Required;MaxSize(64)" validAlias:"创建者"`            // 创建者
	CreatedAt   time.Time   `gorm:"column:created_at" json:"CreatedAt" form:"created_at" query:"created_at" validAlias:"创建时间"`                                        // 创建时间
	UpdatedBy   string      `gorm:"column:updated_by" json:"UpdatedBy" form:"updated_by" query:"updated_by" valid:"Required;MaxSize(64)" validAlias:"更新者"`            // 更新者
	UpdatedAt   time.Time   `gorm:"column:updated_at" json:"UpdatedAt" form:"updated_at" query:"updated_at" validAlias:"更新时间"`                                        // 更新时间
	Remarks     null.String `gorm:"column:remarks" json:"Remarks" form:"remarks" query:"remarks" valid:"MaxSize(100)" validAlias:"备注信息"`                              // 备注信息
	DeletedAt   null.Time   `gorm:"column:deleted_at" json:"DeletedAt" form:"deleted_at" query:"deleted_at" validAlias:"删除时间"`                                        // 删除时间
}

// TableName sets the insert table name for this struct type
func (s *SysMenu) TableName() string {
	return "sys_menu"
}
