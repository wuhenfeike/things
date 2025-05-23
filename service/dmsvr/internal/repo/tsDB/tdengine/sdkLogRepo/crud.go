package sdkLogRepo

import (
	"context"
	"database/sql"
	"fmt"
	"gitee.com/unitedrhino/share/stores"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/domain/deviceLog"

	"gitee.com/unitedrhino/share/def"
	sq "gitee.com/unitedrhino/squirrel"
)

func (d SDKLogRepo) GetDeviceSDKLog(ctx context.Context,
	filter deviceLog.SDKFilter, page def.PageInfo2) ([]*deviceLog.SDK, error) {
	sqSql := sq.Select("*").From(d.GetSDKLogStableName())
	sqSql = d.fillFilter(sqSql, filter)
	sqSql = page.FmtSql(sqSql)
	sqlStr, value, err := sqSql.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := d.t.QueryContext(ctx, sqlStr, value...)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		} else {
			return []*deviceLog.SDK{}, nil
		}
	}
	var datas []map[string]any
	stores.Scan(rows, &datas)
	retLogs := make([]*deviceLog.SDK, 0, len(datas))
	for _, v := range datas {
		retLogs = append(retLogs, ToDeviceSDKLog(filter.ProductID, v))
	}
	return retLogs, nil
}

func (d SDKLogRepo) Insert(ctx context.Context, data *deviceLog.SDK) error {
	sql := fmt.Sprintf(
		" %s using %s tags('%s','%s')(`ts`, `content`,`log_level`) values (?,?,?)",
		d.GetSDKLogTableName(data.ProductID, data.DeviceName), d.GetSDKLogStableName(), data.ProductID, data.DeviceName)
	//if _, err := d.t.ExecContext(ctx, sql, data.Timestamp, data.Content, data.LogLevel); err != nil {
	//	logx.WithContext(ctx).Errorf(
	//		sql+"%s.EventTable product_id:%v device_name:%v err:%v",
	//		utils.FuncName(), data.ProductID, data.DeviceName, err)
	//	return err
	//}
	d.t.AsyncInsert(sql, data.Timestamp, data.Content, data.LogLevel)
	return nil
}
func (d SDKLogRepo) fillFilter(sql sq.SelectBuilder, filter deviceLog.SDKFilter) sq.SelectBuilder {
	if len(filter.ProductID) != 0 {
		sql = sql.Where("`product_id`=?", filter.ProductID)
	}
	if len(filter.DeviceName) != 0 {
		sql = sql.Where("`device_name`=?", filter.DeviceName)
	}
	if filter.LogLevel != 0 {
		sql = sql.Where("`log_level`=?", filter.LogLevel)
	}
	return sql
}
func (d SDKLogRepo) GetCountLog(ctx context.Context, filter deviceLog.SDKFilter, page def.PageInfo2) (int64, error) {
	sqSql := sq.Select("Count(1)").From(d.GetSDKLogStableName())
	sqSql = d.fillFilter(sqSql, filter)
	sqSql = page.FmtWhere(sqSql)
	sqlStr, value, err := sqSql.ToSql()
	if err != nil {
		return 0, err
	}
	row := d.t.QueryRowContext(ctx, sqlStr, value...)
	if err != nil {
		return 0, err
	}
	var (
		total int64
	)

	err = row.Scan(&total)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	return total, nil
}
