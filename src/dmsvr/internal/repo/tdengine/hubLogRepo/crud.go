package hubLogRepo

import (
	"context"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/i-Things/things/shared/def"
	"github.com/i-Things/things/shared/store"
	"github.com/i-Things/things/src/dmsvr/internal/domain/device"
)

func (d HubLogRepo) GetCountLog(ctx context.Context, productID, deviceName string, page def.PageInfo2) (int64, error) {
	sqSql := sq.Select("Count(1)").From(getLogStableName(productID)).
		Where("`device_name`=?", deviceName)
	sqSql = page.FmtWhere(sqSql)
	sqlStr, value, err := sqSql.ToSql()
	if err != nil {
		return 0, err
	}
	row := d.t.QueryRow(sqlStr, value...)
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

func (d HubLogRepo) GetDeviceLog(ctx context.Context, productID, deviceName string, page def.PageInfo2) ([]*device.HubLog, error) {
	sql := sq.Select("*").From(getLogStableName(productID)).
		Where("`device_name`=?", deviceName).OrderBy("`ts` desc")
	sql = page.FmtSql(sql)
	sqlStr, value, err := sql.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := d.t.Query(sqlStr, value...)
	if err != nil {
		return nil, err
	}
	var datas []map[string]any
	store.Scan(rows, &datas)
	retLogs := make([]*device.HubLog, 0, len(datas))
	for _, v := range datas {
		retLogs = append(retLogs, ToDeviceLog(productID, v))
	}
	return retLogs, nil
}

func (d HubLogRepo) Insert(ctx context.Context, data *device.HubLog) error {
	sql := fmt.Sprintf("insert into %s using %s tags('%s')(`ts`, `content`, `topic`, `action`,"+
		" `request_id`, `trance_id`, `result_type`) values (?,?,?,?,?,?,?);",
		getLogTableName(data.ProductID, data.DeviceName), getLogStableName(data.ProductID), data.DeviceName)
	if _, err := d.t.Exec(sql, data.Timestamp, data.Content, data.Topic, data.Action,
		data.RequestID, data.TranceID, data.ResultType); err != nil {
		return err
	}
	return nil
}
