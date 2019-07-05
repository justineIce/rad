package model

type VSysUser struct {
	SysUser
	SysCompanyId      string `gorm:"column:sys_company_id" json:"sys_company_id"`
	OfficeRelationIds string `gorm:"column:office_relation_ids" json:"office_relation_ids"`
	OfficeParentId    string `gorm:"column:office_parent_id" json:"office_parent_id"`
}

func (s *VSysUser) TableName() string {
	return "v_sys_user"
}
