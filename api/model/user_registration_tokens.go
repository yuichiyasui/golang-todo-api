// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// UserRegistrationToken is an object representing the database table.
type UserRegistrationToken struct {
	ID        uint64    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Token     string    `boil:"token" json:"token" toml:"token" yaml:"token"`
	ExpiresAt time.Time `boil:"expires_at" json:"expires_at" toml:"expires_at" yaml:"expires_at"`

	R *userRegistrationTokenR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userRegistrationTokenL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserRegistrationTokenColumns = struct {
	ID        string
	Token     string
	ExpiresAt string
}{
	ID:        "id",
	Token:     "token",
	ExpiresAt: "expires_at",
}

var UserRegistrationTokenTableColumns = struct {
	ID        string
	Token     string
	ExpiresAt string
}{
	ID:        "user_registration_tokens.id",
	Token:     "user_registration_tokens.token",
	ExpiresAt: "user_registration_tokens.expires_at",
}

// Generated where

var UserRegistrationTokenWhere = struct {
	ID        whereHelperuint64
	Token     whereHelperstring
	ExpiresAt whereHelpertime_Time
}{
	ID:        whereHelperuint64{field: "`user_registration_tokens`.`id`"},
	Token:     whereHelperstring{field: "`user_registration_tokens`.`token`"},
	ExpiresAt: whereHelpertime_Time{field: "`user_registration_tokens`.`expires_at`"},
}

// UserRegistrationTokenRels is where relationship names are stored.
var UserRegistrationTokenRels = struct {
}{}

// userRegistrationTokenR is where relationships are stored.
type userRegistrationTokenR struct {
}

// NewStruct creates a new relationship struct
func (*userRegistrationTokenR) NewStruct() *userRegistrationTokenR {
	return &userRegistrationTokenR{}
}

// userRegistrationTokenL is where Load methods for each relationship are stored.
type userRegistrationTokenL struct{}

var (
	userRegistrationTokenAllColumns            = []string{"id", "token", "expires_at"}
	userRegistrationTokenColumnsWithoutDefault = []string{"token", "expires_at"}
	userRegistrationTokenColumnsWithDefault    = []string{"id"}
	userRegistrationTokenPrimaryKeyColumns     = []string{"id"}
	userRegistrationTokenGeneratedColumns      = []string{}
)

type (
	// UserRegistrationTokenSlice is an alias for a slice of pointers to UserRegistrationToken.
	// This should almost always be used instead of []UserRegistrationToken.
	UserRegistrationTokenSlice []*UserRegistrationToken
	// UserRegistrationTokenHook is the signature for custom UserRegistrationToken hook methods
	UserRegistrationTokenHook func(context.Context, boil.ContextExecutor, *UserRegistrationToken) error

	userRegistrationTokenQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userRegistrationTokenType                 = reflect.TypeOf(&UserRegistrationToken{})
	userRegistrationTokenMapping              = queries.MakeStructMapping(userRegistrationTokenType)
	userRegistrationTokenPrimaryKeyMapping, _ = queries.BindMapping(userRegistrationTokenType, userRegistrationTokenMapping, userRegistrationTokenPrimaryKeyColumns)
	userRegistrationTokenInsertCacheMut       sync.RWMutex
	userRegistrationTokenInsertCache          = make(map[string]insertCache)
	userRegistrationTokenUpdateCacheMut       sync.RWMutex
	userRegistrationTokenUpdateCache          = make(map[string]updateCache)
	userRegistrationTokenUpsertCacheMut       sync.RWMutex
	userRegistrationTokenUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var userRegistrationTokenAfterSelectMu sync.Mutex
var userRegistrationTokenAfterSelectHooks []UserRegistrationTokenHook

var userRegistrationTokenBeforeInsertMu sync.Mutex
var userRegistrationTokenBeforeInsertHooks []UserRegistrationTokenHook
var userRegistrationTokenAfterInsertMu sync.Mutex
var userRegistrationTokenAfterInsertHooks []UserRegistrationTokenHook

var userRegistrationTokenBeforeUpdateMu sync.Mutex
var userRegistrationTokenBeforeUpdateHooks []UserRegistrationTokenHook
var userRegistrationTokenAfterUpdateMu sync.Mutex
var userRegistrationTokenAfterUpdateHooks []UserRegistrationTokenHook

var userRegistrationTokenBeforeDeleteMu sync.Mutex
var userRegistrationTokenBeforeDeleteHooks []UserRegistrationTokenHook
var userRegistrationTokenAfterDeleteMu sync.Mutex
var userRegistrationTokenAfterDeleteHooks []UserRegistrationTokenHook

var userRegistrationTokenBeforeUpsertMu sync.Mutex
var userRegistrationTokenBeforeUpsertHooks []UserRegistrationTokenHook
var userRegistrationTokenAfterUpsertMu sync.Mutex
var userRegistrationTokenAfterUpsertHooks []UserRegistrationTokenHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *UserRegistrationToken) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userRegistrationTokenAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *UserRegistrationToken) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userRegistrationTokenBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *UserRegistrationToken) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userRegistrationTokenAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *UserRegistrationToken) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userRegistrationTokenBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *UserRegistrationToken) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userRegistrationTokenAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *UserRegistrationToken) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userRegistrationTokenBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *UserRegistrationToken) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userRegistrationTokenAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *UserRegistrationToken) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userRegistrationTokenBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *UserRegistrationToken) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userRegistrationTokenAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUserRegistrationTokenHook registers your hook function for all future operations.
func AddUserRegistrationTokenHook(hookPoint boil.HookPoint, userRegistrationTokenHook UserRegistrationTokenHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		userRegistrationTokenAfterSelectMu.Lock()
		userRegistrationTokenAfterSelectHooks = append(userRegistrationTokenAfterSelectHooks, userRegistrationTokenHook)
		userRegistrationTokenAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		userRegistrationTokenBeforeInsertMu.Lock()
		userRegistrationTokenBeforeInsertHooks = append(userRegistrationTokenBeforeInsertHooks, userRegistrationTokenHook)
		userRegistrationTokenBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		userRegistrationTokenAfterInsertMu.Lock()
		userRegistrationTokenAfterInsertHooks = append(userRegistrationTokenAfterInsertHooks, userRegistrationTokenHook)
		userRegistrationTokenAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		userRegistrationTokenBeforeUpdateMu.Lock()
		userRegistrationTokenBeforeUpdateHooks = append(userRegistrationTokenBeforeUpdateHooks, userRegistrationTokenHook)
		userRegistrationTokenBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		userRegistrationTokenAfterUpdateMu.Lock()
		userRegistrationTokenAfterUpdateHooks = append(userRegistrationTokenAfterUpdateHooks, userRegistrationTokenHook)
		userRegistrationTokenAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		userRegistrationTokenBeforeDeleteMu.Lock()
		userRegistrationTokenBeforeDeleteHooks = append(userRegistrationTokenBeforeDeleteHooks, userRegistrationTokenHook)
		userRegistrationTokenBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		userRegistrationTokenAfterDeleteMu.Lock()
		userRegistrationTokenAfterDeleteHooks = append(userRegistrationTokenAfterDeleteHooks, userRegistrationTokenHook)
		userRegistrationTokenAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		userRegistrationTokenBeforeUpsertMu.Lock()
		userRegistrationTokenBeforeUpsertHooks = append(userRegistrationTokenBeforeUpsertHooks, userRegistrationTokenHook)
		userRegistrationTokenBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		userRegistrationTokenAfterUpsertMu.Lock()
		userRegistrationTokenAfterUpsertHooks = append(userRegistrationTokenAfterUpsertHooks, userRegistrationTokenHook)
		userRegistrationTokenAfterUpsertMu.Unlock()
	}
}

// One returns a single userRegistrationToken record from the query.
func (q userRegistrationTokenQuery) One(ctx context.Context, exec boil.ContextExecutor) (*UserRegistrationToken, error) {
	o := &UserRegistrationToken{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for user_registration_tokens")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all UserRegistrationToken records from the query.
func (q userRegistrationTokenQuery) All(ctx context.Context, exec boil.ContextExecutor) (UserRegistrationTokenSlice, error) {
	var o []*UserRegistrationToken

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to UserRegistrationToken slice")
	}

	if len(userRegistrationTokenAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all UserRegistrationToken records in the query.
func (q userRegistrationTokenQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count user_registration_tokens rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q userRegistrationTokenQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if user_registration_tokens exists")
	}

	return count > 0, nil
}

// UserRegistrationTokens retrieves all the records using an executor.
func UserRegistrationTokens(mods ...qm.QueryMod) userRegistrationTokenQuery {
	mods = append(mods, qm.From("`user_registration_tokens`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`user_registration_tokens`.*"})
	}

	return userRegistrationTokenQuery{q}
}

// FindUserRegistrationToken retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUserRegistrationToken(ctx context.Context, exec boil.ContextExecutor, iD uint64, selectCols ...string) (*UserRegistrationToken, error) {
	userRegistrationTokenObj := &UserRegistrationToken{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `user_registration_tokens` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, userRegistrationTokenObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from user_registration_tokens")
	}

	if err = userRegistrationTokenObj.doAfterSelectHooks(ctx, exec); err != nil {
		return userRegistrationTokenObj, err
	}

	return userRegistrationTokenObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UserRegistrationToken) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("model: no user_registration_tokens provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userRegistrationTokenColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userRegistrationTokenInsertCacheMut.RLock()
	cache, cached := userRegistrationTokenInsertCache[key]
	userRegistrationTokenInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userRegistrationTokenAllColumns,
			userRegistrationTokenColumnsWithDefault,
			userRegistrationTokenColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userRegistrationTokenType, userRegistrationTokenMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userRegistrationTokenType, userRegistrationTokenMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `user_registration_tokens` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `user_registration_tokens` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `user_registration_tokens` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, userRegistrationTokenPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "model: unable to insert into user_registration_tokens")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == userRegistrationTokenMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for user_registration_tokens")
	}

CacheNoHooks:
	if !cached {
		userRegistrationTokenInsertCacheMut.Lock()
		userRegistrationTokenInsertCache[key] = cache
		userRegistrationTokenInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the UserRegistrationToken.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UserRegistrationToken) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	userRegistrationTokenUpdateCacheMut.RLock()
	cache, cached := userRegistrationTokenUpdateCache[key]
	userRegistrationTokenUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userRegistrationTokenAllColumns,
			userRegistrationTokenPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("model: unable to update user_registration_tokens, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `user_registration_tokens` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, userRegistrationTokenPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userRegistrationTokenType, userRegistrationTokenMapping, append(wl, userRegistrationTokenPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update user_registration_tokens row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by update for user_registration_tokens")
	}

	if !cached {
		userRegistrationTokenUpdateCacheMut.Lock()
		userRegistrationTokenUpdateCache[key] = cache
		userRegistrationTokenUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q userRegistrationTokenQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all for user_registration_tokens")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected for user_registration_tokens")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserRegistrationTokenSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("model: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userRegistrationTokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `user_registration_tokens` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userRegistrationTokenPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all in userRegistrationToken slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected all in update all userRegistrationToken")
	}
	return rowsAff, nil
}

var mySQLUserRegistrationTokenUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UserRegistrationToken) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("model: no user_registration_tokens provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userRegistrationTokenColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLUserRegistrationTokenUniqueColumns, o)

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

	userRegistrationTokenUpsertCacheMut.RLock()
	cache, cached := userRegistrationTokenUpsertCache[key]
	userRegistrationTokenUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			userRegistrationTokenAllColumns,
			userRegistrationTokenColumnsWithDefault,
			userRegistrationTokenColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			userRegistrationTokenAllColumns,
			userRegistrationTokenPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("model: unable to upsert user_registration_tokens, could not build update column list")
		}

		ret := strmangle.SetComplement(userRegistrationTokenAllColumns, strmangle.SetIntersect(insert, update))

		cache.query = buildUpsertQueryMySQL(dialect, "`user_registration_tokens`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `user_registration_tokens` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(userRegistrationTokenType, userRegistrationTokenMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userRegistrationTokenType, userRegistrationTokenMapping, ret)
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

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "model: unable to upsert for user_registration_tokens")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == userRegistrationTokenMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(userRegistrationTokenType, userRegistrationTokenMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "model: unable to retrieve unique values for user_registration_tokens")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for user_registration_tokens")
	}

CacheNoHooks:
	if !cached {
		userRegistrationTokenUpsertCacheMut.Lock()
		userRegistrationTokenUpsertCache[key] = cache
		userRegistrationTokenUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single UserRegistrationToken record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UserRegistrationToken) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no UserRegistrationToken provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userRegistrationTokenPrimaryKeyMapping)
	sql := "DELETE FROM `user_registration_tokens` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete from user_registration_tokens")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by delete for user_registration_tokens")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q userRegistrationTokenQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("model: no userRegistrationTokenQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from user_registration_tokens")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for user_registration_tokens")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserRegistrationTokenSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(userRegistrationTokenBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userRegistrationTokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `user_registration_tokens` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userRegistrationTokenPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from userRegistrationToken slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for user_registration_tokens")
	}

	if len(userRegistrationTokenAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *UserRegistrationToken) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUserRegistrationToken(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserRegistrationTokenSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserRegistrationTokenSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userRegistrationTokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `user_registration_tokens`.* FROM `user_registration_tokens` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userRegistrationTokenPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in UserRegistrationTokenSlice")
	}

	*o = slice

	return nil
}

// UserRegistrationTokenExists checks if the UserRegistrationToken row exists.
func UserRegistrationTokenExists(ctx context.Context, exec boil.ContextExecutor, iD uint64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `user_registration_tokens` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if user_registration_tokens exists")
	}

	return exists, nil
}

// Exists checks if the UserRegistrationToken row exists.
func (o *UserRegistrationToken) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return UserRegistrationTokenExists(ctx, exec, o.ID)
}