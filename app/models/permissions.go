// Code generated by SQLBoiler 4.8.6 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Permission is an object representing the database table.
type Permission struct {
	ID        uint64      `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	DeletedAt null.Time   `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	Name      null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`

	R *permissionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L permissionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PermissionColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	Name      string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Name:      "name",
}

var PermissionTableColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	Name      string
}{
	ID:        "permissions.id",
	CreatedAt: "permissions.created_at",
	UpdatedAt: "permissions.updated_at",
	DeletedAt: "permissions.deleted_at",
	Name:      "permissions.name",
}

// Generated where

type whereHelperuint64 struct{ field string }

func (w whereHelperuint64) EQ(x uint64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperuint64) NEQ(x uint64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperuint64) LT(x uint64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperuint64) LTE(x uint64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperuint64) GT(x uint64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperuint64) GTE(x uint64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperuint64) IN(slice []uint64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperuint64) NIN(slice []uint64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var PermissionWhere = struct {
	ID        whereHelperuint64
	CreatedAt whereHelpernull_Time
	UpdatedAt whereHelpernull_Time
	DeletedAt whereHelpernull_Time
	Name      whereHelpernull_String
}{
	ID:        whereHelperuint64{field: "`permissions`.`id`"},
	CreatedAt: whereHelpernull_Time{field: "`permissions`.`created_at`"},
	UpdatedAt: whereHelpernull_Time{field: "`permissions`.`updated_at`"},
	DeletedAt: whereHelpernull_Time{field: "`permissions`.`deleted_at`"},
	Name:      whereHelpernull_String{field: "`permissions`.`name`"},
}

// PermissionRels is where relationship names are stored.
var PermissionRels = struct {
	Users string
}{
	Users: "Users",
}

// permissionR is where relationships are stored.
type permissionR struct {
	Users UserSlice `boil:"Users" json:"Users" toml:"Users" yaml:"Users"`
}

// NewStruct creates a new relationship struct
func (*permissionR) NewStruct() *permissionR {
	return &permissionR{}
}

// permissionL is where Load methods for each relationship are stored.
type permissionL struct{}

var (
	permissionAllColumns            = []string{"id", "created_at", "updated_at", "deleted_at", "name"}
	permissionColumnsWithoutDefault = []string{"created_at", "updated_at", "deleted_at", "name"}
	permissionColumnsWithDefault    = []string{"id"}
	permissionPrimaryKeyColumns     = []string{"id"}
	permissionGeneratedColumns      = []string{}
)

type (
	// PermissionSlice is an alias for a slice of pointers to Permission.
	// This should almost always be used instead of []Permission.
	PermissionSlice []*Permission
	// PermissionHook is the signature for custom Permission hook methods
	PermissionHook func(boil.Executor, *Permission) error

	permissionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	permissionType                 = reflect.TypeOf(&Permission{})
	permissionMapping              = queries.MakeStructMapping(permissionType)
	permissionPrimaryKeyMapping, _ = queries.BindMapping(permissionType, permissionMapping, permissionPrimaryKeyColumns)
	permissionInsertCacheMut       sync.RWMutex
	permissionInsertCache          = make(map[string]insertCache)
	permissionUpdateCacheMut       sync.RWMutex
	permissionUpdateCache          = make(map[string]updateCache)
	permissionUpsertCacheMut       sync.RWMutex
	permissionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var permissionAfterSelectHooks []PermissionHook

var permissionBeforeInsertHooks []PermissionHook
var permissionAfterInsertHooks []PermissionHook

var permissionBeforeUpdateHooks []PermissionHook
var permissionAfterUpdateHooks []PermissionHook

var permissionBeforeDeleteHooks []PermissionHook
var permissionAfterDeleteHooks []PermissionHook

var permissionBeforeUpsertHooks []PermissionHook
var permissionAfterUpsertHooks []PermissionHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Permission) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range permissionAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Permission) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range permissionBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Permission) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range permissionAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Permission) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range permissionBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Permission) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range permissionAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Permission) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range permissionBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Permission) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range permissionAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Permission) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range permissionBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Permission) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range permissionAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPermissionHook registers your hook function for all future operations.
func AddPermissionHook(hookPoint boil.HookPoint, permissionHook PermissionHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		permissionAfterSelectHooks = append(permissionAfterSelectHooks, permissionHook)
	case boil.BeforeInsertHook:
		permissionBeforeInsertHooks = append(permissionBeforeInsertHooks, permissionHook)
	case boil.AfterInsertHook:
		permissionAfterInsertHooks = append(permissionAfterInsertHooks, permissionHook)
	case boil.BeforeUpdateHook:
		permissionBeforeUpdateHooks = append(permissionBeforeUpdateHooks, permissionHook)
	case boil.AfterUpdateHook:
		permissionAfterUpdateHooks = append(permissionAfterUpdateHooks, permissionHook)
	case boil.BeforeDeleteHook:
		permissionBeforeDeleteHooks = append(permissionBeforeDeleteHooks, permissionHook)
	case boil.AfterDeleteHook:
		permissionAfterDeleteHooks = append(permissionAfterDeleteHooks, permissionHook)
	case boil.BeforeUpsertHook:
		permissionBeforeUpsertHooks = append(permissionBeforeUpsertHooks, permissionHook)
	case boil.AfterUpsertHook:
		permissionAfterUpsertHooks = append(permissionAfterUpsertHooks, permissionHook)
	}
}

// One returns a single permission record from the query.
func (q permissionQuery) One(exec boil.Executor) (*Permission, error) {
	o := &Permission{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for permissions")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Permission records from the query.
func (q permissionQuery) All(exec boil.Executor) (PermissionSlice, error) {
	var o []*Permission

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Permission slice")
	}

	if len(permissionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Permission records in the query.
func (q permissionQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count permissions rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q permissionQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if permissions exists")
	}

	return count > 0, nil
}

// Users retrieves all the user's Users with an executor.
func (o *Permission) Users(mods ...qm.QueryMod) userQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.InnerJoin("`users_permissions` on `users`.`id` = `users_permissions`.`user_id`"),
		qm.Where("`users_permissions`.`permission_id`=?", o.ID),
	)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "`users`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`users`.*"})
	}

	return query
}

// LoadUsers allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (permissionL) LoadUsers(e boil.Executor, singular bool, maybePermission interface{}, mods queries.Applicator) error {
	var slice []*Permission
	var object *Permission

	if singular {
		object = maybePermission.(*Permission)
	} else {
		slice = *maybePermission.(*[]*Permission)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &permissionR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &permissionR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.Select("`users`.id, `users`.first_name, `users`.last_name, `users`.email, `users`.password, `users`.is_locked, `users`.is_expried, `users`.created_at, `users`.updated_at, `users`.deleted_at, `a`.`permission_id`"),
		qm.From("`users`"),
		qm.InnerJoin("`users_permissions` as `a` on `users`.`id` = `a`.`user_id`"),
		qm.WhereIn("`a`.`permission_id` in ?", args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load users")
	}

	var resultSlice []*User

	var localJoinCols []uint64
	for results.Next() {
		one := new(User)
		var localJoinCol uint64

		err = results.Scan(&one.ID, &one.FirstName, &one.LastName, &one.Email, &one.Password, &one.IsLocked, &one.IsExpried, &one.CreatedAt, &one.UpdatedAt, &one.DeletedAt, &localJoinCol)
		if err != nil {
			return errors.Wrap(err, "failed to scan eager loaded results for users")
		}
		if err = results.Err(); err != nil {
			return errors.Wrap(err, "failed to plebian-bind eager loaded slice users")
		}

		resultSlice = append(resultSlice, one)
		localJoinCols = append(localJoinCols, localJoinCol)
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(userAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Users = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &userR{}
			}
			foreign.R.Permissions = append(foreign.R.Permissions, object)
		}
		return nil
	}

	for i, foreign := range resultSlice {
		localJoinCol := localJoinCols[i]
		for _, local := range slice {
			if local.ID == localJoinCol {
				local.R.Users = append(local.R.Users, foreign)
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.Permissions = append(foreign.R.Permissions, local)
				break
			}
		}
	}

	return nil
}

// AddUsers adds the given related objects to the existing relationships
// of the permission, optionally inserting them as new records.
// Appends related to o.R.Users.
// Sets related.R.Permissions appropriately.
func (o *Permission) AddUsers(exec boil.Executor, insert bool, related ...*User) error {
	var err error
	for _, rel := range related {
		if insert {
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		}
	}

	for _, rel := range related {
		query := "insert into `users_permissions` (`permission_id`, `user_id`) values (?, ?)"
		values := []interface{}{o.ID, rel.ID}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, query)
			fmt.Fprintln(boil.DebugWriter, values)
		}
		_, err = exec.Exec(query, values...)
		if err != nil {
			return errors.Wrap(err, "failed to insert into join table")
		}
	}
	if o.R == nil {
		o.R = &permissionR{
			Users: related,
		}
	} else {
		o.R.Users = append(o.R.Users, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &userR{
				Permissions: PermissionSlice{o},
			}
		} else {
			rel.R.Permissions = append(rel.R.Permissions, o)
		}
	}
	return nil
}

// SetUsers removes all previously related items of the
// permission replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Permissions's Users accordingly.
// Replaces o.R.Users with related.
// Sets related.R.Permissions's Users accordingly.
func (o *Permission) SetUsers(exec boil.Executor, insert bool, related ...*User) error {
	query := "delete from `users_permissions` where `permission_id` = ?"
	values := []interface{}{o.ID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	_, err := exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	removeUsersFromPermissionsSlice(o, related)
	if o.R != nil {
		o.R.Users = nil
	}
	return o.AddUsers(exec, insert, related...)
}

// RemoveUsers relationships from objects passed in.
// Removes related items from R.Users (uses pointer comparison, removal does not keep order)
// Sets related.R.Permissions.
func (o *Permission) RemoveUsers(exec boil.Executor, related ...*User) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	query := fmt.Sprintf(
		"delete from `users_permissions` where `permission_id` = ? and `user_id` in (%s)",
		strmangle.Placeholders(dialect.UseIndexPlaceholders, len(related), 2, 1),
	)
	values := []interface{}{o.ID}
	for _, rel := range related {
		values = append(values, rel.ID)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	_, err = exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}
	removeUsersFromPermissionsSlice(o, related)
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Users {
			if rel != ri {
				continue
			}

			ln := len(o.R.Users)
			if ln > 1 && i < ln-1 {
				o.R.Users[i] = o.R.Users[ln-1]
			}
			o.R.Users = o.R.Users[:ln-1]
			break
		}
	}

	return nil
}

func removeUsersFromPermissionsSlice(o *Permission, related []*User) {
	for _, rel := range related {
		if rel.R == nil {
			continue
		}
		for i, ri := range rel.R.Permissions {
			if o.ID != ri.ID {
				continue
			}

			ln := len(rel.R.Permissions)
			if ln > 1 && i < ln-1 {
				rel.R.Permissions[i] = rel.R.Permissions[ln-1]
			}
			rel.R.Permissions = rel.R.Permissions[:ln-1]
			break
		}
	}
}

// Permissions retrieves all the records using an executor.
func Permissions(mods ...qm.QueryMod) permissionQuery {
	mods = append(mods, qm.From("`permissions`"))
	return permissionQuery{NewQuery(mods...)}
}

// FindPermission retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPermission(exec boil.Executor, iD uint64, selectCols ...string) (*Permission, error) {
	permissionObj := &Permission{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `permissions` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, permissionObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from permissions")
	}

	if err = permissionObj.doAfterSelectHooks(exec); err != nil {
		return permissionObj, err
	}

	return permissionObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Permission) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no permissions provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if queries.MustTime(o.CreatedAt).IsZero() {
		queries.SetScanner(&o.CreatedAt, currTime)
	}
	if queries.MustTime(o.UpdatedAt).IsZero() {
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(permissionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	permissionInsertCacheMut.RLock()
	cache, cached := permissionInsertCache[key]
	permissionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			permissionAllColumns,
			permissionColumnsWithDefault,
			permissionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(permissionType, permissionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(permissionType, permissionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `permissions` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `permissions` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `permissions` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, permissionPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}
	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into permissions")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == permissionMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}
	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for permissions")
	}

CacheNoHooks:
	if !cached {
		permissionInsertCacheMut.Lock()
		permissionInsertCache[key] = cache
		permissionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// Update uses an executor to update the Permission.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Permission) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	queries.SetScanner(&o.UpdatedAt, currTime)

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	permissionUpdateCacheMut.RLock()
	cache, cached := permissionUpdateCache[key]
	permissionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			permissionAllColumns,
			permissionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update permissions, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `permissions` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, permissionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(permissionType, permissionMapping, append(wl, permissionPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	var result sql.Result
	result, err = exec.Exec(cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update permissions row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for permissions")
	}

	if !cached {
		permissionUpdateCacheMut.Lock()
		permissionUpdateCache[key] = cache
		permissionUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAll updates all rows with the specified column values.
func (q permissionQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for permissions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for permissions")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PermissionSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), permissionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `permissions` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, permissionPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in permission slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all permission")
	}
	return rowsAff, nil
}

var mySQLPermissionUniqueColumns = []string{
	"id",
	"name",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Permission) Upsert(exec boil.Executor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no permissions provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if queries.MustTime(o.CreatedAt).IsZero() {
		queries.SetScanner(&o.CreatedAt, currTime)
	}
	queries.SetScanner(&o.UpdatedAt, currTime)

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(permissionColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLPermissionUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	permissionUpsertCacheMut.RLock()
	cache, cached := permissionUpsertCache[key]
	permissionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			permissionAllColumns,
			permissionColumnsWithDefault,
			permissionColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			permissionAllColumns,
			permissionPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert permissions, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`permissions`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `permissions` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(permissionType, permissionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(permissionType, permissionMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}
	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for permissions")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == permissionMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(permissionType, permissionMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for permissions")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}
	err = exec.QueryRow(cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for permissions")
	}

CacheNoHooks:
	if !cached {
		permissionUpsertCacheMut.Lock()
		permissionUpsertCache[key] = cache
		permissionUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// Delete deletes a single Permission record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Permission) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Permission provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), permissionPrimaryKeyMapping)
	sql := "DELETE FROM `permissions` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from permissions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for permissions")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q permissionQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no permissionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from permissions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for permissions")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PermissionSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(permissionBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), permissionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `permissions` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, permissionPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from permission slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for permissions")
	}

	if len(permissionAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Permission) Reload(exec boil.Executor) error {
	ret, err := FindPermission(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PermissionSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PermissionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), permissionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `permissions`.* FROM `permissions` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, permissionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PermissionSlice")
	}

	*o = slice

	return nil
}

// PermissionExists checks if the Permission row exists.
func PermissionExists(exec boil.Executor, iD uint64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `permissions` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if permissions exists")
	}

	return exists, nil
}
