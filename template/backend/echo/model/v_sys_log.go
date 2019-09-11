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
type VSysLog struct {
	ID           string    `gorm:"column:id" json:"ID" form:"id" query:"id" valid:"Required;MaxSize(64)" validAlias:"编号"`                                                 // 编号
	Title        string    `gorm:"column:title" json:"Title" form:"title" query:"title" valid:"Required;MaxSize(100)" validAlias:"日志标题"`                                  // 日志标题
	CreatedBy    string    `gorm:"column:created_by" json:"CreatedBy" form:"created_by" query:"created_by" valid:"Required;MaxSize(64)" validAlias:"创建者"`                 // 创建者
	CreatedAt    time.Time `gorm:"column:created_at" json:"CreatedAt" form:"created_at" query:"created_at" validAlias:"创建时间"`                                             // 创建时间
	RemoteAddr   string    `gorm:"column:remote_addr" json:"RemoteAddr" form:"remote_addr" query:"remote_addr" valid:"Required;MaxSize(100)" validAlias:"操作IP地址"`         // 操作IP地址
	UserAgent    string    `gorm:"column:user_agent" json:"UserAgent" form:"user_agent" query:"user_agent" valid:"Required;MaxSize(500)" validAlias:"用户代理"`               // 用户代理
	RequestURL   string    `gorm:"column:request_url" json:"RequestURL" form:"request_url" query:"request_url" valid:"Required;MaxSize(100)" validAlias:"请求URI"`          // 请求URI
	Method       string    `gorm:"column:method" json:"Method" form:"method" query:"method" valid:"Required;MaxSize(5)" validAlias:"操作方式"`                                // 操作方式
	Params       string    `gorm:"column:params" json:"Params" form:"params" query:"params" valid:"Required;MaxSize(100)" validAlias:"操作提交的数据"`                           // 操作提交的数据
	Exception    string    `gorm:"column:exception" json:"Exception" form:"exception" query:"exception" valid:"MaxSize(100)" validAlias:"异常信息"`                           // 异常信息
	SysCompanyID string    `gorm:"column:sys_company_id" json:"SysCompanyID" form:"sys_company_id" query:"sys_company_id" valid:"Required;MaxSize(64)" validAlias:"公司id"` // 公司id
	UpdatedBy    string    `gorm:"column:updated_by" json:"UpdatedBy" form:"updated_by" query:"updated_by" valid:"Required;MaxSize(64)" validAlias:"更新者"`                 // 更新者
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"UpdatedAt" form:"updated_at" query:"updated_at" validAlias:"更新时间"`                                             // 更新时间
	DeletedAt    null.Time `gorm:"column:deleted_at" json:"DeletedAt" form:"deleted_at" query:"deleted_at" validAlias:"删除时间"`                                             // 删除时间
	ResultStatus int64     `gorm:"column:result_status" form:"result_status"`                                                                                             //请求结果状态码
	ResultMsg    string    `gorm:"column:result_msg" form:"result_msg"`
	Name         string    `gorm:"column:name" json:"Name" form:"name" query:"name" valid:"Required;MaxSize(50)" validAlias:"姓名"` // 姓名，Table[search]
}

// TableName sets the insert table name for this struct type
func (v *VSysLog) TableName() string {
	return "v_sys_log"
}
