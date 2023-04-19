// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: chat_model.sql

package sqlc_queries

import (
	"context"
)

const chatModelByID = `-- name: ChatModelByID :one
SELECT id, name, label, is_default, url, api_auth_header, api_auth_key FROM chat_model WHERE id = $1
`

func (q *Queries) ChatModelByID(ctx context.Context, id int32) (ChatModel, error) {
	row := q.db.QueryRowContext(ctx, chatModelByID, id)
	var i ChatModel
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Label,
		&i.IsDefault,
		&i.Url,
		&i.ApiAuthHeader,
		&i.ApiAuthKey,
	)
	return i, err
}

const chatModelByName = `-- name: ChatModelByName :one
SELECT id, name, label, is_default, url, api_auth_header, api_auth_key FROM chat_model WHERE name = $1
`

func (q *Queries) ChatModelByName(ctx context.Context, name string) (ChatModel, error) {
	row := q.db.QueryRowContext(ctx, chatModelByName, name)
	var i ChatModel
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Label,
		&i.IsDefault,
		&i.Url,
		&i.ApiAuthHeader,
		&i.ApiAuthKey,
	)
	return i, err
}

const createChatModel = `-- name: CreateChatModel :one
INSERT INTO chat_model (name, label, is_default, url, api_auth_header, api_auth_key)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, name, label, is_default, url, api_auth_header, api_auth_key
`

type CreateChatModelParams struct {
	Name          string
	Label         string
	IsDefault     bool
	Url           string
	ApiAuthHeader string
	ApiAuthKey    string
}

func (q *Queries) CreateChatModel(ctx context.Context, arg CreateChatModelParams) (ChatModel, error) {
	row := q.db.QueryRowContext(ctx, createChatModel,
		arg.Name,
		arg.Label,
		arg.IsDefault,
		arg.Url,
		arg.ApiAuthHeader,
		arg.ApiAuthKey,
	)
	var i ChatModel
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Label,
		&i.IsDefault,
		&i.Url,
		&i.ApiAuthHeader,
		&i.ApiAuthKey,
	)
	return i, err
}

const deleteChatModel = `-- name: DeleteChatModel :exec
DELETE FROM chat_model WHERE id = $1
`

func (q *Queries) DeleteChatModel(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteChatModel, id)
	return err
}

const getDefaultChatModel = `-- name: GetDefaultChatModel :one
SELECT id, name, label, is_default, url, api_auth_header, api_auth_key FROM chat_model WHERE is_default = true
`

func (q *Queries) GetDefaultChatModel(ctx context.Context) (ChatModel, error) {
	row := q.db.QueryRowContext(ctx, getDefaultChatModel)
	var i ChatModel
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Label,
		&i.IsDefault,
		&i.Url,
		&i.ApiAuthHeader,
		&i.ApiAuthKey,
	)
	return i, err
}

const listChatModels = `-- name: ListChatModels :many
SELECT id, name, label, is_default, url, api_auth_header, api_auth_key FROM chat_model ORDER BY id
`

func (q *Queries) ListChatModels(ctx context.Context) ([]ChatModel, error) {
	rows, err := q.db.QueryContext(ctx, listChatModels)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ChatModel
	for rows.Next() {
		var i ChatModel
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Label,
			&i.IsDefault,
			&i.Url,
			&i.ApiAuthHeader,
			&i.ApiAuthKey,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateChatModel = `-- name: UpdateChatModel :one
UPDATE chat_model SET name = $2, label = $3, is_default = $4, url = $5, api_auth_header = $6, api_auth_key = $7
WHERE id = $1
RETURNING id, name, label, is_default, url, api_auth_header, api_auth_key
`

type UpdateChatModelParams struct {
	ID            int32
	Name          string
	Label         string
	IsDefault     bool
	Url           string
	ApiAuthHeader string
	ApiAuthKey    string
}

func (q *Queries) UpdateChatModel(ctx context.Context, arg UpdateChatModelParams) (ChatModel, error) {
	row := q.db.QueryRowContext(ctx, updateChatModel,
		arg.ID,
		arg.Name,
		arg.Label,
		arg.IsDefault,
		arg.Url,
		arg.ApiAuthHeader,
		arg.ApiAuthKey,
	)
	var i ChatModel
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Label,
		&i.IsDefault,
		&i.Url,
		&i.ApiAuthHeader,
		&i.ApiAuthKey,
	)
	return i, err
}

const updateChatModelKey = `-- name: UpdateChatModelKey :one
UPDATE chat_model SET api_auth_key = $2
WHERE id = $1
RETURNING id, name, label, is_default, url, api_auth_header, api_auth_key
`

type UpdateChatModelKeyParams struct {
	ID         int32
	ApiAuthKey string
}

func (q *Queries) UpdateChatModelKey(ctx context.Context, arg UpdateChatModelKeyParams) (ChatModel, error) {
	row := q.db.QueryRowContext(ctx, updateChatModelKey, arg.ID, arg.ApiAuthKey)
	var i ChatModel
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Label,
		&i.IsDefault,
		&i.Url,
		&i.ApiAuthHeader,
		&i.ApiAuthKey,
	)
	return i, err
}
