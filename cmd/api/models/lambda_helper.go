// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "p7": Model Helpers
//
// Command:
// $ goagen
// --design=github.com/markusklems/p7/cmd/api/design
// --out=$(GOPATH)/src/github.com/markusklems/p7/cmd/api
// --version=v1.2.0-dirty

package models

import (
	"context"
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/markusklems/p7/cmd/api/app"
	"time"
)

// MediaType Retrieval Functions

// ListLambdaCode returns an array of view: code.
func (m *LambdaDB) ListLambdaCode(ctx context.Context) []*app.LambdaCode {
	defer goa.MeasureSince([]string{"goa", "db", "lambda", "listlambdacode"}, time.Now())

	var native []*Lambda
	var objs []*app.LambdaCode
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Lambda", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.LambdaToLambdaCode())
	}

	return objs
}

// LambdaToLambdaCode loads a Lambda and builds the code view of media type Lambda.
func (m *Lambda) LambdaToLambdaCode() *app.LambdaCode {
	lambda := &app.LambdaCode{}
	lambda.Code = m.Code

	return lambda
}

// OneLambdaCode loads a Lambda and builds the code view of media type Lambda.
func (m *LambdaDB) OneLambdaCode(ctx context.Context, id int) (*app.LambdaCode, error) {
	defer goa.MeasureSince([]string{"goa", "db", "lambda", "onelambdacode"}, time.Now())

	var native Lambda
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Lambda", "error", err.Error())
		return nil, err
	}

	view := *native.LambdaToLambdaCode()
	return &view, err
}

// MediaType Retrieval Functions

// ListLambda returns an array of view: default.
func (m *LambdaDB) ListLambda(ctx context.Context) []*app.Lambda {
	defer goa.MeasureSince([]string{"goa", "db", "lambda", "listlambda"}, time.Now())

	var native []*Lambda
	var objs []*app.Lambda
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Lambda", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.LambdaToLambda())
	}

	return objs
}

// LambdaToLambda loads a Lambda and builds the default view of media type Lambda.
func (m *Lambda) LambdaToLambda() *app.Lambda {
	lambda := &app.Lambda{}
	lambda.Environment = &m.Environment
	lambda.ID = m.ID
	lambda.Method = m.Method
	lambda.Name = m.Name

	return lambda
}

// OneLambda loads a Lambda and builds the default view of media type Lambda.
func (m *LambdaDB) OneLambda(ctx context.Context, id int) (*app.Lambda, error) {
	defer goa.MeasureSince([]string{"goa", "db", "lambda", "onelambda"}, time.Now())

	var native Lambda
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Lambda", "error", err.Error())
		return nil, err
	}

	view := *native.LambdaToLambda()
	return &view, err
}

// MediaType Retrieval Functions

// ListLambdaFull returns an array of view: full.
func (m *LambdaDB) ListLambdaFull(ctx context.Context) []*app.LambdaFull {
	defer goa.MeasureSince([]string{"goa", "db", "lambda", "listlambdafull"}, time.Now())

	var native []*Lambda
	var objs []*app.LambdaFull
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Lambda", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.LambdaToLambdaFull())
	}

	return objs
}

// LambdaToLambdaFull loads a Lambda and builds the full view of media type Lambda.
func (m *Lambda) LambdaToLambdaFull() *app.LambdaFull {
	lambda := &app.LambdaFull{}
	lambda.Code = m.Code
	lambda.CreatedAt = &m.CreatedAt
	lambda.Environment = &m.Environment
	lambda.ID = m.ID
	lambda.Method = m.Method
	lambda.Name = m.Name
	lambda.UpdatedAt = &m.UpdatedAt

	return lambda
}

// OneLambdaFull loads a Lambda and builds the full view of media type Lambda.
func (m *LambdaDB) OneLambdaFull(ctx context.Context, id int) (*app.LambdaFull, error) {
	defer goa.MeasureSince([]string{"goa", "db", "lambda", "onelambdafull"}, time.Now())

	var native Lambda
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Lambda", "error", err.Error())
		return nil, err
	}

	view := *native.LambdaToLambdaFull()
	return &view, err
}

// MediaType Retrieval Functions

// ListLambdaTiny returns an array of view: tiny.
func (m *LambdaDB) ListLambdaTiny(ctx context.Context) []*app.LambdaTiny {
	defer goa.MeasureSince([]string{"goa", "db", "lambda", "listlambdatiny"}, time.Now())

	var native []*Lambda
	var objs []*app.LambdaTiny
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Lambda", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.LambdaToLambdaTiny())
	}

	return objs
}

// LambdaToLambdaTiny loads a Lambda and builds the tiny view of media type Lambda.
func (m *Lambda) LambdaToLambdaTiny() *app.LambdaTiny {
	lambda := &app.LambdaTiny{}
	lambda.ID = m.ID
	lambda.Method = m.Method
	lambda.Name = m.Name

	return lambda
}

// OneLambdaTiny loads a Lambda and builds the tiny view of media type Lambda.
func (m *LambdaDB) OneLambdaTiny(ctx context.Context, id int) (*app.LambdaTiny, error) {
	defer goa.MeasureSince([]string{"goa", "db", "lambda", "onelambdatiny"}, time.Now())

	var native Lambda
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Lambda", "error", err.Error())
		return nil, err
	}

	view := *native.LambdaToLambdaTiny()
	return &view, err
}
