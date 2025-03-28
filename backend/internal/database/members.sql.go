// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: members.sql

package database

import (
	"context"
	"time"
)

const createMember = `-- name: CreateMember :one
INSERT INTO members (
    member_name, member_contact, membership, created_by, membership_start, membership_end
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING id, member_name, member_contact, membership, created_by, membership_start, membership_end
`

type CreateMemberParams struct {
	MemberName      string     `json:"member_name"`
	MemberContact   string     `json:"member_contact"`
	Membership      int64      `json:"membership"`
	CreatedBy       int64      `json:"created_by"`
	MembershipStart *time.Time `json:"membership_start"`
	MembershipEnd   *time.Time `json:"membership_end"`
}

func (q *Queries) CreateMember(ctx context.Context, arg CreateMemberParams) (Member, error) {
	row := q.db.QueryRow(ctx, createMember,
		arg.MemberName,
		arg.MemberContact,
		arg.Membership,
		arg.CreatedBy,
		arg.MembershipStart,
		arg.MembershipEnd,
	)
	var i Member
	err := row.Scan(
		&i.ID,
		&i.MemberName,
		&i.MemberContact,
		&i.Membership,
		&i.CreatedBy,
		&i.MembershipStart,
		&i.MembershipEnd,
	)
	return i, err
}

const deleteMember = `-- name: DeleteMember :exec
DELETE FROM members
WHERE id = $1
`

func (q *Queries) DeleteMember(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteMember, id)
	return err
}

const getMemberByID = `-- name: GetMemberByID :one
SELECT id, member_name, member_contact, membership, created_by, membership_start, membership_end FROM members
WHERE id = $1
`

func (q *Queries) GetMemberByID(ctx context.Context, id int64) (Member, error) {
	row := q.db.QueryRow(ctx, getMemberByID, id)
	var i Member
	err := row.Scan(
		&i.ID,
		&i.MemberName,
		&i.MemberContact,
		&i.Membership,
		&i.CreatedBy,
		&i.MembershipStart,
		&i.MembershipEnd,
	)
	return i, err
}

const updateMember = `-- name: UpdateMember :one
UPDATE members
    set member_name = $2,
    member_contact = $3,
    membership = $4,
    membership_start = $5,
    membership_end = $6
WHERE id = $1
RETURNING membership_start, membership_end
`

type UpdateMemberParams struct {
	ID              int64      `json:"id"`
	MemberName      string     `json:"member_name"`
	MemberContact   string     `json:"member_contact"`
	Membership      int64      `json:"membership"`
	MembershipStart *time.Time `json:"membership_start"`
	MembershipEnd   *time.Time `json:"membership_end"`
}

type UpdateMemberRow struct {
	MembershipStart *time.Time `json:"membership_start"`
	MembershipEnd   *time.Time `json:"membership_end"`
}

func (q *Queries) UpdateMember(ctx context.Context, arg UpdateMemberParams) (UpdateMemberRow, error) {
	row := q.db.QueryRow(ctx, updateMember,
		arg.ID,
		arg.MemberName,
		arg.MemberContact,
		arg.Membership,
		arg.MembershipStart,
		arg.MembershipEnd,
	)
	var i UpdateMemberRow
	err := row.Scan(&i.MembershipStart, &i.MembershipEnd)
	return i, err
}
