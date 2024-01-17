// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package menumanage

import (
	"context"

	"github.com/i-Things/things/src/syssvr/internal/svc"
	"github.com/i-Things/things/src/syssvr/pb/sys"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ApiDeleteReq              = sys.ApiDeleteReq
	ApiInfo                   = sys.ApiInfo
	ApiInfoIndexReq           = sys.ApiInfoIndexReq
	ApiInfoIndexResp          = sys.ApiInfoIndexResp
	AppInfo                   = sys.AppInfo
	AppInfoIndexReq           = sys.AppInfoIndexReq
	AppInfoIndexResp          = sys.AppInfoIndexResp
	AppModuleIndexReq         = sys.AppModuleIndexReq
	AppModuleIndexResp        = sys.AppModuleIndexResp
	AppModuleMultiUpdateReq   = sys.AppModuleMultiUpdateReq
	AreaInfo                  = sys.AreaInfo
	AreaInfoIndexReq          = sys.AreaInfoIndexReq
	AreaInfoIndexResp         = sys.AreaInfoIndexResp
	AreaInfoReadReq           = sys.AreaInfoReadReq
	AreaWithID                = sys.AreaWithID
	AuthApiInfo               = sys.AuthApiInfo
	ConfigResp                = sys.ConfigResp
	DateRange                 = sys.DateRange
	JwtToken                  = sys.JwtToken
	LoginLogCreateReq         = sys.LoginLogCreateReq
	LoginLogIndexReq          = sys.LoginLogIndexReq
	LoginLogIndexResp         = sys.LoginLogIndexResp
	LoginLogInfo              = sys.LoginLogInfo
	Map                       = sys.Map
	MenuInfo                  = sys.MenuInfo
	MenuInfoIndexReq          = sys.MenuInfoIndexReq
	MenuInfoIndexResp         = sys.MenuInfoIndexResp
	ModuleInfo                = sys.ModuleInfo
	ModuleInfoIndexReq        = sys.ModuleInfoIndexReq
	ModuleInfoIndexResp       = sys.ModuleInfoIndexResp
	OperLogCreateReq          = sys.OperLogCreateReq
	OperLogIndexReq           = sys.OperLogIndexReq
	OperLogIndexResp          = sys.OperLogIndexResp
	OperLogInfo               = sys.OperLogInfo
	PageInfo                  = sys.PageInfo
	PageInfo_OrderBy          = sys.PageInfo_OrderBy
	Point                     = sys.Point
	ProjectInfo               = sys.ProjectInfo
	ProjectInfoIndexReq       = sys.ProjectInfoIndexReq
	ProjectInfoIndexResp      = sys.ProjectInfoIndexResp
	ProjectWithID             = sys.ProjectWithID
	Response                  = sys.Response
	RoleApiAuthReq            = sys.RoleApiAuthReq
	RoleApiIndexReq           = sys.RoleApiIndexReq
	RoleApiIndexResp          = sys.RoleApiIndexResp
	RoleApiMultiUpdateReq     = sys.RoleApiMultiUpdateReq
	RoleAppIndexReq           = sys.RoleAppIndexReq
	RoleAppIndexResp          = sys.RoleAppIndexResp
	RoleAppMultiUpdateReq     = sys.RoleAppMultiUpdateReq
	RoleAppUpdateReq          = sys.RoleAppUpdateReq
	RoleInfo                  = sys.RoleInfo
	RoleInfoIndexReq          = sys.RoleInfoIndexReq
	RoleInfoIndexResp         = sys.RoleInfoIndexResp
	RoleMenuIndexReq          = sys.RoleMenuIndexReq
	RoleMenuIndexResp         = sys.RoleMenuIndexResp
	RoleMenuMultiUpdateReq    = sys.RoleMenuMultiUpdateReq
	RoleModuleIndexReq        = sys.RoleModuleIndexReq
	RoleModuleIndexResp       = sys.RoleModuleIndexResp
	RoleModuleMultiUpdateReq  = sys.RoleModuleMultiUpdateReq
	TenantApiInfo             = sys.TenantApiInfo
	TenantAppApiIndexReq      = sys.TenantAppApiIndexReq
	TenantAppApiIndexResp     = sys.TenantAppApiIndexResp
	TenantAppCreateReq        = sys.TenantAppCreateReq
	TenantAppIndexReq         = sys.TenantAppIndexReq
	TenantAppIndexResp        = sys.TenantAppIndexResp
	TenantAppMenu             = sys.TenantAppMenu
	TenantAppMenuIndexReq     = sys.TenantAppMenuIndexReq
	TenantAppMenuIndexResp    = sys.TenantAppMenuIndexResp
	TenantAppModule           = sys.TenantAppModule
	TenantAppMultiUpdateReq   = sys.TenantAppMultiUpdateReq
	TenantAppWithIDOrCode     = sys.TenantAppWithIDOrCode
	TenantInfo                = sys.TenantInfo
	TenantInfoCreateReq       = sys.TenantInfoCreateReq
	TenantInfoIndexReq        = sys.TenantInfoIndexReq
	TenantInfoIndexResp       = sys.TenantInfoIndexResp
	TenantModuleCreateReq     = sys.TenantModuleCreateReq
	TenantModuleIndexReq      = sys.TenantModuleIndexReq
	TenantModuleIndexResp     = sys.TenantModuleIndexResp
	TenantModuleWithIDOrCode  = sys.TenantModuleWithIDOrCode
	UserArea                  = sys.UserArea
	UserAreaIndexReq          = sys.UserAreaIndexReq
	UserAreaIndexResp         = sys.UserAreaIndexResp
	UserAreaMultiUpdateReq    = sys.UserAreaMultiUpdateReq
	UserCaptchaReq            = sys.UserCaptchaReq
	UserCaptchaResp           = sys.UserCaptchaResp
	UserChangePwdReq          = sys.UserChangePwdReq
	UserCheckTokenReq         = sys.UserCheckTokenReq
	UserCheckTokenResp        = sys.UserCheckTokenResp
	UserCreateResp            = sys.UserCreateResp
	UserForgetPwdReq          = sys.UserForgetPwdReq
	UserInfo                  = sys.UserInfo
	UserInfoCreateReq         = sys.UserInfoCreateReq
	UserInfoDeleteReq         = sys.UserInfoDeleteReq
	UserInfoIndexReq          = sys.UserInfoIndexReq
	UserInfoIndexResp         = sys.UserInfoIndexResp
	UserInfoReadReq           = sys.UserInfoReadReq
	UserLoginReq              = sys.UserLoginReq
	UserLoginResp             = sys.UserLoginResp
	UserProject               = sys.UserProject
	UserProjectIndexReq       = sys.UserProjectIndexReq
	UserProjectIndexResp      = sys.UserProjectIndexResp
	UserProjectMultiUpdateReq = sys.UserProjectMultiUpdateReq
	UserRegisterReq           = sys.UserRegisterReq
	UserRegisterResp          = sys.UserRegisterResp
	UserRoleIndexReq          = sys.UserRoleIndexReq
	UserRoleIndexResp         = sys.UserRoleIndexResp
	UserRoleMultiUpdateReq    = sys.UserRoleMultiUpdateReq
	WithAppCodeID             = sys.WithAppCodeID
	WithID                    = sys.WithID
	WithIDCode                = sys.WithIDCode

	MenuManage interface {
		MenuInfoCreate(ctx context.Context, in *MenuInfo, opts ...grpc.CallOption) (*WithID, error)
		MenuInfoIndex(ctx context.Context, in *MenuInfoIndexReq, opts ...grpc.CallOption) (*MenuInfoIndexResp, error)
		MenuInfoUpdate(ctx context.Context, in *MenuInfo, opts ...grpc.CallOption) (*Response, error)
		MenuInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Response, error)
	}

	defaultMenuManage struct {
		cli zrpc.Client
	}

	directMenuManage struct {
		svcCtx *svc.ServiceContext
		svr    sys.MenuManageServer
	}
)

func NewMenuManage(cli zrpc.Client) MenuManage {
	return &defaultMenuManage{
		cli: cli,
	}
}

func NewDirectMenuManage(svcCtx *svc.ServiceContext, svr sys.MenuManageServer) MenuManage {
	return &directMenuManage{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

func (m *defaultMenuManage) MenuInfoCreate(ctx context.Context, in *MenuInfo, opts ...grpc.CallOption) (*WithID, error) {
	client := sys.NewMenuManageClient(m.cli.Conn())
	return client.MenuInfoCreate(ctx, in, opts...)
}

func (d *directMenuManage) MenuInfoCreate(ctx context.Context, in *MenuInfo, opts ...grpc.CallOption) (*WithID, error) {
	return d.svr.MenuInfoCreate(ctx, in)
}

func (m *defaultMenuManage) MenuInfoIndex(ctx context.Context, in *MenuInfoIndexReq, opts ...grpc.CallOption) (*MenuInfoIndexResp, error) {
	client := sys.NewMenuManageClient(m.cli.Conn())
	return client.MenuInfoIndex(ctx, in, opts...)
}

func (d *directMenuManage) MenuInfoIndex(ctx context.Context, in *MenuInfoIndexReq, opts ...grpc.CallOption) (*MenuInfoIndexResp, error) {
	return d.svr.MenuInfoIndex(ctx, in)
}

func (m *defaultMenuManage) MenuInfoUpdate(ctx context.Context, in *MenuInfo, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewMenuManageClient(m.cli.Conn())
	return client.MenuInfoUpdate(ctx, in, opts...)
}

func (d *directMenuManage) MenuInfoUpdate(ctx context.Context, in *MenuInfo, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.MenuInfoUpdate(ctx, in)
}

func (m *defaultMenuManage) MenuInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewMenuManageClient(m.cli.Conn())
	return client.MenuInfoDelete(ctx, in, opts...)
}

func (d *directMenuManage) MenuInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.MenuInfoDelete(ctx, in)
}
