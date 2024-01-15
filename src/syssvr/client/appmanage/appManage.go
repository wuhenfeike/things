// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package appmanage

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
	AreaInfoTreeReq           = sys.AreaInfoTreeReq
	AreaInfoTreeResp          = sys.AreaInfoTreeResp
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

	AppManage interface {
		AppInfoCreate(ctx context.Context, in *AppInfo, opts ...grpc.CallOption) (*WithID, error)
		AppInfoIndex(ctx context.Context, in *AppInfoIndexReq, opts ...grpc.CallOption) (*AppInfoIndexResp, error)
		AppInfoUpdate(ctx context.Context, in *AppInfo, opts ...grpc.CallOption) (*Response, error)
		AppInfoDelete(ctx context.Context, in *WithIDCode, opts ...grpc.CallOption) (*Response, error)
		AppInfoRead(ctx context.Context, in *WithIDCode, opts ...grpc.CallOption) (*AppInfo, error)
		AppModuleIndex(ctx context.Context, in *AppModuleIndexReq, opts ...grpc.CallOption) (*AppModuleIndexResp, error)
		AppModuleMultiUpdate(ctx context.Context, in *AppModuleMultiUpdateReq, opts ...grpc.CallOption) (*Response, error)
	}

	defaultAppManage struct {
		cli zrpc.Client
	}

	directAppManage struct {
		svcCtx *svc.ServiceContext
		svr    sys.AppManageServer
	}
)

func NewAppManage(cli zrpc.Client) AppManage {
	return &defaultAppManage{
		cli: cli,
	}
}

func NewDirectAppManage(svcCtx *svc.ServiceContext, svr sys.AppManageServer) AppManage {
	return &directAppManage{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

func (m *defaultAppManage) AppInfoCreate(ctx context.Context, in *AppInfo, opts ...grpc.CallOption) (*WithID, error) {
	client := sys.NewAppManageClient(m.cli.Conn())
	return client.AppInfoCreate(ctx, in, opts...)
}

func (d *directAppManage) AppInfoCreate(ctx context.Context, in *AppInfo, opts ...grpc.CallOption) (*WithID, error) {
	return d.svr.AppInfoCreate(ctx, in)
}

func (m *defaultAppManage) AppInfoIndex(ctx context.Context, in *AppInfoIndexReq, opts ...grpc.CallOption) (*AppInfoIndexResp, error) {
	client := sys.NewAppManageClient(m.cli.Conn())
	return client.AppInfoIndex(ctx, in, opts...)
}

func (d *directAppManage) AppInfoIndex(ctx context.Context, in *AppInfoIndexReq, opts ...grpc.CallOption) (*AppInfoIndexResp, error) {
	return d.svr.AppInfoIndex(ctx, in)
}

func (m *defaultAppManage) AppInfoUpdate(ctx context.Context, in *AppInfo, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewAppManageClient(m.cli.Conn())
	return client.AppInfoUpdate(ctx, in, opts...)
}

func (d *directAppManage) AppInfoUpdate(ctx context.Context, in *AppInfo, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.AppInfoUpdate(ctx, in)
}

func (m *defaultAppManage) AppInfoDelete(ctx context.Context, in *WithIDCode, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewAppManageClient(m.cli.Conn())
	return client.AppInfoDelete(ctx, in, opts...)
}

func (d *directAppManage) AppInfoDelete(ctx context.Context, in *WithIDCode, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.AppInfoDelete(ctx, in)
}

func (m *defaultAppManage) AppInfoRead(ctx context.Context, in *WithIDCode, opts ...grpc.CallOption) (*AppInfo, error) {
	client := sys.NewAppManageClient(m.cli.Conn())
	return client.AppInfoRead(ctx, in, opts...)
}

func (d *directAppManage) AppInfoRead(ctx context.Context, in *WithIDCode, opts ...grpc.CallOption) (*AppInfo, error) {
	return d.svr.AppInfoRead(ctx, in)
}

func (m *defaultAppManage) AppModuleIndex(ctx context.Context, in *AppModuleIndexReq, opts ...grpc.CallOption) (*AppModuleIndexResp, error) {
	client := sys.NewAppManageClient(m.cli.Conn())
	return client.AppModuleIndex(ctx, in, opts...)
}

func (d *directAppManage) AppModuleIndex(ctx context.Context, in *AppModuleIndexReq, opts ...grpc.CallOption) (*AppModuleIndexResp, error) {
	return d.svr.AppModuleIndex(ctx, in)
}

func (m *defaultAppManage) AppModuleMultiUpdate(ctx context.Context, in *AppModuleMultiUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	client := sys.NewAppManageClient(m.cli.Conn())
	return client.AppModuleMultiUpdate(ctx, in, opts...)
}

func (d *directAppManage) AppModuleMultiUpdate(ctx context.Context, in *AppModuleMultiUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.AppModuleMultiUpdate(ctx, in)
}
