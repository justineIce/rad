package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	UserPostSuperAdminCode = "2" //超级管理员
	UserPostAdminCode      = "1" //管理员
	UserPostGroupCode      = "0" //普通账号
)

//后台系统登陆
type SysUserLoginInfo struct {
	ID                string               //ID
	Name              string               //账号名称
	Sex               string               //性别
	CompanyId         string               //公司ID
	OfficeId          string               //组织ID
	OfficeParentId    string               //组织关系
	OfficeRelationIds string               //组织关系
	UserType          string               //账号类型
	SysRoles          []VSysUserRole       //角色，支持多角色
	SysMenu           []VSysRoleMenu       //菜单
	SysMenuBtns       []VSysRoleMenuBtn    //菜单按钮权限
	IsSuperAdmin      bool                 //是否为超级管理员
	IsAdmin           bool                 //是否管理员
	PathAuth          map[string]time.Time //url已通过的鉴权
}

// login token struct
type SysUserLoginToken struct {
	ID                 string //ID
	Unix               int64  //时间戳
	jwt.StandardClaims        //jwt
}

// token struct
type SysToken struct {
	Token       string //token
	TokenExpire int64  //token过期时间戳
}

//菜单树
type SysMenuTree struct {
	VSysRoleMenu
	Children []SysMenuTree
}

//登录token struct
func (info SysUserLoginInfo) GetLoginToken() *SysUserLoginToken {
	return &SysUserLoginToken{
		ID:   info.ID,
		Unix: time.Now().Unix(),
	}
}

//是否为超级管理员
func IsSuperAdmin(roles []VSysUserRole) bool {
	for i := 0; i < len(roles); i++ {
		if roles[i].RoleType == UserPostSuperAdminCode {
			return true
		}
	}
	return false
}

//是否为超级管理员
func IsAdmin(roles []VSysUserRole) bool {
	for i := 0; i < len(roles); i++ {
		if roles[i].RoleType == UserPostAdminCode {
			return true
		}
	}
	return false
}

//登录信息
func GetLoginInfo(vSysUser VSysUser, roles []VSysUserRole, menus []VSysRoleMenu, menuBtns []VSysRoleMenuBtn) SysUserLoginInfo {
	return SysUserLoginInfo{
		ID:                vSysUser.ID,
		Name:              vSysUser.Name,
		CompanyId:         vSysUser.CompanyID,
		OfficeId:          vSysUser.OfficeID,
		OfficeParentId:    vSysUser.OfficeParentID,
		OfficeRelationIds: vSysUser.RelationIds,
		SysRoles:          roles,
		//SysMenuThree:      GetSysMenuThree(menus, "0"),
		SysMenu:      menus,
		SysMenuBtns:  menuBtns,
		IsAdmin:      IsAdmin(roles),
		IsSuperAdmin: IsSuperAdmin(roles),
		PathAuth:     make(map[string]time.Time, 0),
	}
}

//获取树形菜单
func GetSysMenuThree(menus []VSysRoleMenu, parentId string) []SysMenuTree {
	var smuts = make([]SysMenuTree, 0)
	var smut SysMenuTree
	for i := 0; i < len(menus); i++ {
		if menus[i].ParentID != parentId {
			continue
		}
		smut = SysMenuTree{}
		smut.ID = menus[i].ID
		smut.ParentID = menus[i].ParentID
		smut.Name = menus[i].Name
		smut.Href = menus[i].Href
		smut.Icon = menus[i].Icon
		smut.IsShow = menus[i].IsShow
		smut.Permission = menus[i].Permission
		smut.RelationIds = menus[i].RelationIds
		smut.Remarks = menus[i].Remarks
		smut.Target = menus[i].Target
		smut.Children = GetSysMenuThree(menus, menus[i].ID)
		smuts = append(smuts, smut)
	}
	return smuts
}
