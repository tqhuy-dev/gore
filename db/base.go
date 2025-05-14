package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"strings"
)

func parseSQLProjection(opt *OptionFilter) (projections string) {
	projections = "*"
	if opt == nil {
		return
	}
	if len(opt.Projection) == 0 {
		return
	}
	arr := make([]string, 0, len(opt.Projection))
	for _, ele := range opt.Projection {
		arr = append(arr, ele)
	}
	projections = " " + strings.Join(arr, ",") + " "
	return
}

func parseSQLCondition(condition map[string]interface{}) (conditionStr string, values []interface{}) {
	conditionStr = ""
	if condition == nil {
		return
	}
	if len(condition) == 0 {
		return
	}
	arrCondition := make([]string, 0, len(condition))
	arrValue := make([]interface{}, 0, len(condition))
	for key, value := range condition {
		arrCondition = append(arrCondition, fmt.Sprintf("%s = $%d", key, len(arrValue)+1))
		arrValue = append(arrValue, value)
	}
	conditionStr = strings.Join(arrCondition, " AND ")
	return
}

type SortType int

const (
	Asc  SortType = 1
	Desc SortType = 2
)

type OptionFilter struct {
	Projection []string
	Limit      int
	Offset     int
	SortBy     map[string]SortType
}

type ModifyResult struct {
	ModifiedCount int64
	NewCount      int64
}
type BaseRepo[T any, IdEntity any] interface {
	FindOneByID(ctx context.Context, id IdEntity, opt *OptionFilter) (*T, error)
	FindOneByCondition(ctx context.Context, condition map[string]interface{}, opt *OptionFilter) (*T, error)
	FindByCondition(ctx context.Context, condition map[string]interface{}, opt *OptionFilter) ([]*T, error)
}

type baseSQLImpl[T any, IdEntity any] struct {
	NameTable string
	cnn       pgx.Conn
}

func (b *baseSQLImpl[T, IdEntity]) FindOneByID(ctx context.Context, id IdEntity, opt *OptionFilter) (*T, error) {
	var data T

	sql := `SELECT ` + parseSQLProjection(opt) + ` FROM ` + b.NameTable + ` WHERE id = ? limit 1`
	row, err := b.cnn.Query(ctx, sql, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	data, err = pgx.CollectOneRow(row, pgx.RowToStructByPos[T])
	return &data, nil
}

func (b *baseSQLImpl[T, IdEntity]) FindOneByCondition(ctx context.Context, condition map[string]interface{}, opt *OptionFilter) (*T, error) {
	var data T
	sql := `SELECT ` + parseSQLProjection(opt) + ` FROM ` + b.NameTable
	conditionStr, values := parseSQLCondition(condition)
	if values != nil && len(values) > 0 && len(conditionStr) > 0 {
		conditionStr = " WHERE " + conditionStr
	}
	sql += conditionStr
	row, err := b.cnn.Query(ctx, sql, values...)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	data, err = pgx.CollectOneRow(row, pgx.RowToStructByPos[T])
	return &data, nil
}

func (b *baseSQLImpl[T, IdEntity]) FindByCondition(ctx context.Context, condition map[string]interface{}, opt *OptionFilter) ([]*T, error) {
	var data []*T
	sql := `SELECT ` + parseSQLProjection(opt) + ` FROM ` + b.NameTable
	conditionStr, values := parseSQLCondition(condition)
	if values != nil && len(values) > 0 && len(conditionStr) > 0 {
		conditionStr = " WHERE " + conditionStr
	}
	sql += conditionStr
	row, err := b.cnn.Query(ctx, sql, values...)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		var element T
		element, err = pgx.CollectOneRow(row, pgx.RowToStructByPos[T])
		data = append(data, &element)
	}
	return data, nil
}

func NewBaseRepo[T any, IdEntity any](nameTable string, cnn pgx.Conn) BaseRepo[T, IdEntity] {
	return &baseSQLImpl[T, IdEntity]{
		NameTable: nameTable,
		cnn:       cnn,
	}
}
