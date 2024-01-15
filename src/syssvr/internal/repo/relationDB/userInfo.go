package relationDB

import (
	"context"
	"github.com/i-Things/things/shared/def"
	"github.com/i-Things/things/shared/stores"
	"gorm.io/gorm"
)

type UserInfoRepo struct {
	db *gorm.DB
}

func NewUserInfoRepo(in any) *UserInfoRepo {
	return &UserInfoRepo{db: stores.GetCommonConn(in)}
}

type UserInfoFilter struct {
	UserIDs        []int64
	TenantCode     string
	UserNames      []string
	UserName       string
	Phone          string
	Phones         []string
	Email          string
	Emails         []string
	WechatOpenIDs  []string
	Accounts       []string //账号查询 非模糊查询
	WechatUnionID  string
	WechatOpenID   string
	DingTalkUserID string
	WithRoles      bool
	WithTenant     bool
}

func (p UserInfoRepo) accountsFilter(db *gorm.DB, accounts []string) *gorm.DB {
	db = db.Where(db.Or("user_name in ?", accounts).
		Or("email in ?", accounts).
		Or("phone in ?", accounts))
	return db
}

func (p UserInfoRepo) fmtFilter(ctx context.Context, f UserInfoFilter) *gorm.DB {
	db := p.db.WithContext(ctx)
	if f.DingTalkUserID != "" {
		db = db.Where("ding_talk_user_id = ?", f.DingTalkUserID)
	}
	if f.WithRoles {
		db = db.Preload("Roles")
	}
	if f.WithTenant {
		db = db.Preload("Tenant")
	}
	if len(f.UserIDs) != 0 {
		db = db.Where("user_id in?", f.UserIDs)
	}
	if len(f.UserNames) != 0 {
		db = db.Where("user_name in ?", f.UserNames)
	}
	if len(f.Accounts) != 0 {
		db = p.accountsFilter(db, f.Accounts)
	}
	if f.UserName != "" {
		db = db.Where("user_name like ?", "%"+f.UserName+"%")
	}
	if f.Phone != "" {
		db = db.Where("phone like ?", "%"+f.Phone+"%")
	}
	if len(f.Phones) != 0 {
		db = db.Where("phone in ?", f.Phones)
	}
	if f.Email != "" {
		db = db.Where("email like ?", "%"+f.Email+"%")
	}
	if len(f.Emails) != 0 {
		db = db.Where("email in ?", f.Emails)
	}
	wechatOr := db
	var isWechat bool
	if f.WechatUnionID != "" {
		isWechat = true
		wechatOr = wechatOr.Or("wechat_union_id = ?", f.WechatUnionID)
	}
	if f.WechatOpenID != "" {
		isWechat = true
		wechatOr = wechatOr.Or("wechat_open_id = ?", f.WechatOpenID)
	}
	if isWechat {
		db = db.Where(wechatOr)
	}
	if f.TenantCode != "" {
		db = db.Where("tenant_code =?", f.TenantCode)
	}
	return db
}

func (p UserInfoRepo) Insert(ctx context.Context, data *SysTenantUserInfo) error {
	result := p.db.WithContext(ctx).Create(data)
	return stores.ErrFmt(result.Error)
}

func (p UserInfoRepo) FindOneByFilter(ctx context.Context, f UserInfoFilter) (*SysTenantUserInfo, error) {
	var result SysTenantUserInfo
	db := p.fmtFilter(ctx, f)
	err := db.First(&result).Error
	if err != nil {
		return nil, stores.ErrFmt(err)
	}
	return &result, nil
}
func (p UserInfoRepo) FindByFilter(ctx context.Context, f UserInfoFilter, page *def.PageInfo) ([]*SysTenantUserInfo, error) {
	var results []*SysTenantUserInfo
	db := p.fmtFilter(ctx, f).Model(&SysTenantUserInfo{})
	db = page.ToGorm(db)
	err := db.Find(&results).Error
	if err != nil {
		return nil, stores.ErrFmt(err)
	}
	return results, nil
}

func (p UserInfoRepo) CountByFilter(ctx context.Context, f UserInfoFilter) (size int64, err error) {
	db := p.fmtFilter(ctx, f).Model(&SysTenantUserInfo{})
	err = db.Count(&size).Error
	return size, stores.ErrFmt(err)
}

func (p UserInfoRepo) Update(ctx context.Context, data *SysTenantUserInfo) error {
	err := p.db.WithContext(ctx).Where("user_id = ?", data.UserID).Save(data).Error
	return stores.ErrFmt(err)
}

func (p UserInfoRepo) DeleteByFilter(ctx context.Context, f UserInfoFilter) error {
	db := p.fmtFilter(ctx, f)
	err := db.Delete(&SysTenantUserInfo{}).Error
	return stores.ErrFmt(err)
}

func (p UserInfoRepo) Delete(ctx context.Context, userID int64) error {
	err := p.db.WithContext(ctx).Where("user_id = ?", userID).Delete(&SysTenantUserInfo{}).Error
	return stores.ErrFmt(err)
}
func (p UserInfoRepo) FindOne(ctx context.Context, userID int64) (*SysTenantUserInfo, error) {
	var result SysTenantUserInfo
	err := p.db.WithContext(ctx).Where("user_id = ?", userID).First(&result).Error
	return &result, stores.ErrFmt(err)
}
