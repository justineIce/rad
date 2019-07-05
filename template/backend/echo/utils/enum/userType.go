package enum

/*
角色类型 - 关于数据权限
角色：
	系统管理员
		所有数据管理
	管理员
		旗下组织架构的数据管理
	普通人员
		管理自己创建的数据
*/
const (
	UserPostSuperAdmin = "系统管理员" //系统管理员
	UserPostAdmin      = "记录管理员" //记录管理员
	UserPostGroup      = "记录统计员" //记录统计员
)
