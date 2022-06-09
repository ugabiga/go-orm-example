// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

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

// Project is an object representing the database table.
type Project struct {
	ID          int64         `boil:"id" json:"id" toml:"id" yaml:"id"`
	Title       string        `boil:"title" json:"title" toml:"title" yaml:"title"`
	Description string        `boil:"description" json:"description" toml:"description" yaml:"description"`
	Status      ProjectStatus `boil:"status" json:"status" toml:"status" yaml:"status"`
	UpdatedAt   time.Time     `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	CreatedAt   time.Time     `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *projectR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L projectL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ProjectColumns = struct {
	ID          string
	Title       string
	Description string
	Status      string
	UpdatedAt   string
	CreatedAt   string
}{
	ID:          "id",
	Title:       "title",
	Description: "description",
	Status:      "status",
	UpdatedAt:   "updated_at",
	CreatedAt:   "created_at",
}

var ProjectTableColumns = struct {
	ID          string
	Title       string
	Description string
	Status      string
	UpdatedAt   string
	CreatedAt   string
}{
	ID:          "projects.id",
	Title:       "projects.title",
	Description: "projects.description",
	Status:      "projects.status",
	UpdatedAt:   "projects.updated_at",
	CreatedAt:   "projects.created_at",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperProjectStatus struct{ field string }

func (w whereHelperProjectStatus) EQ(x ProjectStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelperProjectStatus) NEQ(x ProjectStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperProjectStatus) LT(x ProjectStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelperProjectStatus) LTE(x ProjectStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperProjectStatus) GT(x ProjectStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelperProjectStatus) GTE(x ProjectStatus) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var ProjectWhere = struct {
	ID          whereHelperint64
	Title       whereHelperstring
	Description whereHelperstring
	Status      whereHelperProjectStatus
	UpdatedAt   whereHelpertime_Time
	CreatedAt   whereHelpertime_Time
}{
	ID:          whereHelperint64{field: "\"projects\".\"id\""},
	Title:       whereHelperstring{field: "\"projects\".\"title\""},
	Description: whereHelperstring{field: "\"projects\".\"description\""},
	Status:      whereHelperProjectStatus{field: "\"projects\".\"status\""},
	UpdatedAt:   whereHelpertime_Time{field: "\"projects\".\"updated_at\""},
	CreatedAt:   whereHelpertime_Time{field: "\"projects\".\"created_at\""},
}

// ProjectRels is where relationship names are stored.
var ProjectRels = struct {
	ProjectTasks string
}{
	ProjectTasks: "ProjectTasks",
}

// projectR is where relationships are stored.
type projectR struct {
	ProjectTasks ProjectTaskSlice `boil:"ProjectTasks" json:"ProjectTasks" toml:"ProjectTasks" yaml:"ProjectTasks"`
}

// NewStruct creates a new relationship struct
func (*projectR) NewStruct() *projectR {
	return &projectR{}
}

func (r *projectR) GetProjectTasks() ProjectTaskSlice {
	if r == nil {
		return nil
	}
	return r.ProjectTasks
}

// projectL is where Load methods for each relationship are stored.
type projectL struct{}

var (
	projectAllColumns            = []string{"id", "title", "description", "status", "updated_at", "created_at"}
	projectColumnsWithoutDefault = []string{"title", "description", "status"}
	projectColumnsWithDefault    = []string{"id", "updated_at", "created_at"}
	projectPrimaryKeyColumns     = []string{"id"}
	projectGeneratedColumns      = []string{}
)

type (
	// ProjectSlice is an alias for a slice of pointers to Project.
	// This should almost always be used instead of []Project.
	ProjectSlice []*Project
	// ProjectHook is the signature for custom Project hook methods
	ProjectHook func(context.Context, boil.ContextExecutor, *Project) error

	projectQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	projectType                 = reflect.TypeOf(&Project{})
	projectMapping              = queries.MakeStructMapping(projectType)
	projectPrimaryKeyMapping, _ = queries.BindMapping(projectType, projectMapping, projectPrimaryKeyColumns)
	projectInsertCacheMut       sync.RWMutex
	projectInsertCache          = make(map[string]insertCache)
	projectUpdateCacheMut       sync.RWMutex
	projectUpdateCache          = make(map[string]updateCache)
	projectUpsertCacheMut       sync.RWMutex
	projectUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var projectAfterSelectHooks []ProjectHook

var projectBeforeInsertHooks []ProjectHook
var projectAfterInsertHooks []ProjectHook

var projectBeforeUpdateHooks []ProjectHook
var projectAfterUpdateHooks []ProjectHook

var projectBeforeDeleteHooks []ProjectHook
var projectAfterDeleteHooks []ProjectHook

var projectBeforeUpsertHooks []ProjectHook
var projectAfterUpsertHooks []ProjectHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Project) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Project) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Project) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Project) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Project) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Project) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Project) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Project) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Project) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddProjectHook registers your hook function for all future operations.
func AddProjectHook(hookPoint boil.HookPoint, projectHook ProjectHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		projectAfterSelectHooks = append(projectAfterSelectHooks, projectHook)
	case boil.BeforeInsertHook:
		projectBeforeInsertHooks = append(projectBeforeInsertHooks, projectHook)
	case boil.AfterInsertHook:
		projectAfterInsertHooks = append(projectAfterInsertHooks, projectHook)
	case boil.BeforeUpdateHook:
		projectBeforeUpdateHooks = append(projectBeforeUpdateHooks, projectHook)
	case boil.AfterUpdateHook:
		projectAfterUpdateHooks = append(projectAfterUpdateHooks, projectHook)
	case boil.BeforeDeleteHook:
		projectBeforeDeleteHooks = append(projectBeforeDeleteHooks, projectHook)
	case boil.AfterDeleteHook:
		projectAfterDeleteHooks = append(projectAfterDeleteHooks, projectHook)
	case boil.BeforeUpsertHook:
		projectBeforeUpsertHooks = append(projectBeforeUpsertHooks, projectHook)
	case boil.AfterUpsertHook:
		projectAfterUpsertHooks = append(projectAfterUpsertHooks, projectHook)
	}
}

// One returns a single project record from the query.
func (q projectQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Project, error) {
	o := &Project{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for projects")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Project records from the query.
func (q projectQuery) All(ctx context.Context, exec boil.ContextExecutor) (ProjectSlice, error) {
	var o []*Project

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Project slice")
	}

	if len(projectAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Project records in the query.
func (q projectQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count projects rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q projectQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if projects exists")
	}

	return count > 0, nil
}

// ProjectTasks retrieves all the project_task's ProjectTasks with an executor.
func (o *Project) ProjectTasks(mods ...qm.QueryMod) projectTaskQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"project_tasks\".\"project_id\"=?", o.ID),
	)

	return ProjectTasks(queryMods...)
}

// LoadProjectTasks allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (projectL) LoadProjectTasks(ctx context.Context, e boil.ContextExecutor, singular bool, maybeProject interface{}, mods queries.Applicator) error {
	var slice []*Project
	var object *Project

	if singular {
		object = maybeProject.(*Project)
	} else {
		slice = *maybeProject.(*[]*Project)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &projectR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &projectR{}
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
		qm.From(`project_tasks`),
		qm.WhereIn(`project_tasks.project_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load project_tasks")
	}

	var resultSlice []*ProjectTask
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice project_tasks")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on project_tasks")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for project_tasks")
	}

	if len(projectTaskAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.ProjectTasks = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &projectTaskR{}
			}
			foreign.R.Project = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.ProjectID {
				local.R.ProjectTasks = append(local.R.ProjectTasks, foreign)
				if foreign.R == nil {
					foreign.R = &projectTaskR{}
				}
				foreign.R.Project = local
				break
			}
		}
	}

	return nil
}

// AddProjectTasks adds the given related objects to the existing relationships
// of the project, optionally inserting them as new records.
// Appends related to o.R.ProjectTasks.
// Sets related.R.Project appropriately.
func (o *Project) AddProjectTasks(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*ProjectTask) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ProjectID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"project_tasks\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"project_id"}),
				strmangle.WhereClause("\"", "\"", 2, projectTaskPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ProjectID, rel.TaskID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ProjectID = o.ID
		}
	}

	if o.R == nil {
		o.R = &projectR{
			ProjectTasks: related,
		}
	} else {
		o.R.ProjectTasks = append(o.R.ProjectTasks, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &projectTaskR{
				Project: o,
			}
		} else {
			rel.R.Project = o
		}
	}
	return nil
}

// Projects retrieves all the records using an executor.
func Projects(mods ...qm.QueryMod) projectQuery {
	mods = append(mods, qm.From("\"projects\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"projects\".*"})
	}

	return projectQuery{q}
}

// FindProject retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindProject(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Project, error) {
	projectObj := &Project{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"projects\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, projectObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from projects")
	}

	if err = projectObj.doAfterSelectHooks(ctx, exec); err != nil {
		return projectObj, err
	}

	return projectObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Project) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no projects provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(projectColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	projectInsertCacheMut.RLock()
	cache, cached := projectInsertCache[key]
	projectInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			projectAllColumns,
			projectColumnsWithDefault,
			projectColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(projectType, projectMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(projectType, projectMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"projects\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"projects\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
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

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into projects")
	}

	if !cached {
		projectInsertCacheMut.Lock()
		projectInsertCache[key] = cache
		projectInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Project.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Project) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	projectUpdateCacheMut.RLock()
	cache, cached := projectUpdateCache[key]
	projectUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			projectAllColumns,
			projectPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update projects, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"projects\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, projectPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(projectType, projectMapping, append(wl, projectPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update projects row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for projects")
	}

	if !cached {
		projectUpdateCacheMut.Lock()
		projectUpdateCache[key] = cache
		projectUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q projectQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for projects")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for projects")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ProjectSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), projectPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"projects\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, projectPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in project slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all project")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Project) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no projects provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(projectColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
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
	key := buf.String()
	strmangle.PutBuffer(buf)

	projectUpsertCacheMut.RLock()
	cache, cached := projectUpsertCache[key]
	projectUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			projectAllColumns,
			projectColumnsWithDefault,
			projectColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			projectAllColumns,
			projectPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert projects, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(projectPrimaryKeyColumns))
			copy(conflict, projectPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"projects\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(projectType, projectMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(projectType, projectMapping, ret)
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
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert projects")
	}

	if !cached {
		projectUpsertCacheMut.Lock()
		projectUpsertCache[key] = cache
		projectUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Project record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Project) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Project provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), projectPrimaryKeyMapping)
	sql := "DELETE FROM \"projects\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from projects")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for projects")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q projectQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no projectQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from projects")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for projects")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ProjectSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(projectBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), projectPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"projects\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, projectPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from project slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for projects")
	}

	if len(projectAfterDeleteHooks) != 0 {
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
func (o *Project) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindProject(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProjectSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ProjectSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), projectPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"projects\".* FROM \"projects\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, projectPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ProjectSlice")
	}

	*o = slice

	return nil
}

// ProjectExists checks if the Project row exists.
func ProjectExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"projects\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if projects exists")
	}

	return exists, nil
}