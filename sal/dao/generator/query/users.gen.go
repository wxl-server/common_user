// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"common_user/sal/dao/generator/model"
)

func newUserPO(db *gorm.DB, opts ...gen.DOOption) userPO {
	_userPO := userPO{}

	_userPO.userPODo.UseDB(db, opts...)
	_userPO.userPODo.UseModel(&model.UserPO{})

	tableName := _userPO.userPODo.TableName()
	_userPO.ALL = field.NewAsterisk(tableName)
	_userPO.ID = field.NewInt64(tableName, "id")
	_userPO.Email = field.NewString(tableName, "email")
	_userPO.Password = field.NewString(tableName, "password")
	_userPO.Extra = field.NewString(tableName, "extra")
	_userPO.CreatedAt = field.NewTime(tableName, "created_at")
	_userPO.UpdatedAt = field.NewTime(tableName, "updated_at")
	_userPO.DeletedAt = field.NewField(tableName, "deleted_at")

	_userPO.fillFieldMap()

	return _userPO
}

type userPO struct {
	userPODo

	ALL       field.Asterisk
	ID        field.Int64
	Email     field.String
	Password  field.String
	Extra     field.String
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (u userPO) Table(newTableName string) *userPO {
	u.userPODo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userPO) As(alias string) *userPO {
	u.userPODo.DO = *(u.userPODo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userPO) updateTableName(table string) *userPO {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.Email = field.NewString(table, "email")
	u.Password = field.NewString(table, "password")
	u.Extra = field.NewString(table, "extra")
	u.CreatedAt = field.NewTime(table, "created_at")
	u.UpdatedAt = field.NewTime(table, "updated_at")
	u.DeletedAt = field.NewField(table, "deleted_at")

	u.fillFieldMap()

	return u
}

func (u *userPO) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userPO) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 7)
	u.fieldMap["id"] = u.ID
	u.fieldMap["email"] = u.Email
	u.fieldMap["password"] = u.Password
	u.fieldMap["extra"] = u.Extra
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["deleted_at"] = u.DeletedAt
}

func (u userPO) clone(db *gorm.DB) userPO {
	u.userPODo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userPO) replaceDB(db *gorm.DB) userPO {
	u.userPODo.ReplaceDB(db)
	return u
}

type userPODo struct{ gen.DO }

type IUserPODo interface {
	gen.SubQuery
	Debug() IUserPODo
	WithContext(ctx context.Context) IUserPODo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserPODo
	WriteDB() IUserPODo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserPODo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserPODo
	Not(conds ...gen.Condition) IUserPODo
	Or(conds ...gen.Condition) IUserPODo
	Select(conds ...field.Expr) IUserPODo
	Where(conds ...gen.Condition) IUserPODo
	Order(conds ...field.Expr) IUserPODo
	Distinct(cols ...field.Expr) IUserPODo
	Omit(cols ...field.Expr) IUserPODo
	Join(table schema.Tabler, on ...field.Expr) IUserPODo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserPODo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserPODo
	Group(cols ...field.Expr) IUserPODo
	Having(conds ...gen.Condition) IUserPODo
	Limit(limit int) IUserPODo
	Offset(offset int) IUserPODo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserPODo
	Unscoped() IUserPODo
	Create(values ...*model.UserPO) error
	CreateInBatches(values []*model.UserPO, batchSize int) error
	Save(values ...*model.UserPO) error
	First() (*model.UserPO, error)
	Take() (*model.UserPO, error)
	Last() (*model.UserPO, error)
	Find() ([]*model.UserPO, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserPO, err error)
	FindInBatches(result *[]*model.UserPO, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UserPO) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserPODo
	Assign(attrs ...field.AssignExpr) IUserPODo
	Joins(fields ...field.RelationField) IUserPODo
	Preload(fields ...field.RelationField) IUserPODo
	FirstOrInit() (*model.UserPO, error)
	FirstOrCreate() (*model.UserPO, error)
	FindByPage(offset int, limit int) (result []*model.UserPO, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserPODo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userPODo) Debug() IUserPODo {
	return u.withDO(u.DO.Debug())
}

func (u userPODo) WithContext(ctx context.Context) IUserPODo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userPODo) ReadDB() IUserPODo {
	return u.Clauses(dbresolver.Read)
}

func (u userPODo) WriteDB() IUserPODo {
	return u.Clauses(dbresolver.Write)
}

func (u userPODo) Session(config *gorm.Session) IUserPODo {
	return u.withDO(u.DO.Session(config))
}

func (u userPODo) Clauses(conds ...clause.Expression) IUserPODo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userPODo) Returning(value interface{}, columns ...string) IUserPODo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userPODo) Not(conds ...gen.Condition) IUserPODo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userPODo) Or(conds ...gen.Condition) IUserPODo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userPODo) Select(conds ...field.Expr) IUserPODo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userPODo) Where(conds ...gen.Condition) IUserPODo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userPODo) Order(conds ...field.Expr) IUserPODo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userPODo) Distinct(cols ...field.Expr) IUserPODo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userPODo) Omit(cols ...field.Expr) IUserPODo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userPODo) Join(table schema.Tabler, on ...field.Expr) IUserPODo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userPODo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserPODo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userPODo) RightJoin(table schema.Tabler, on ...field.Expr) IUserPODo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userPODo) Group(cols ...field.Expr) IUserPODo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userPODo) Having(conds ...gen.Condition) IUserPODo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userPODo) Limit(limit int) IUserPODo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userPODo) Offset(offset int) IUserPODo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userPODo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserPODo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userPODo) Unscoped() IUserPODo {
	return u.withDO(u.DO.Unscoped())
}

func (u userPODo) Create(values ...*model.UserPO) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userPODo) CreateInBatches(values []*model.UserPO, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userPODo) Save(values ...*model.UserPO) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userPODo) First() (*model.UserPO, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserPO), nil
	}
}

func (u userPODo) Take() (*model.UserPO, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserPO), nil
	}
}

func (u userPODo) Last() (*model.UserPO, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserPO), nil
	}
}

func (u userPODo) Find() ([]*model.UserPO, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserPO), err
}

func (u userPODo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserPO, err error) {
	buf := make([]*model.UserPO, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userPODo) FindInBatches(result *[]*model.UserPO, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userPODo) Attrs(attrs ...field.AssignExpr) IUserPODo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userPODo) Assign(attrs ...field.AssignExpr) IUserPODo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userPODo) Joins(fields ...field.RelationField) IUserPODo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userPODo) Preload(fields ...field.RelationField) IUserPODo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userPODo) FirstOrInit() (*model.UserPO, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserPO), nil
	}
}

func (u userPODo) FirstOrCreate() (*model.UserPO, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserPO), nil
	}
}

func (u userPODo) FindByPage(offset int, limit int) (result []*model.UserPO, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userPODo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userPODo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userPODo) Delete(models ...*model.UserPO) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userPODo) withDO(do gen.Dao) *userPODo {
	u.DO = *do.(*gen.DO)
	return u
}
