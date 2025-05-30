-- name: GetMemberships :many
SELECT * FROM memberships
WHERE created_by = $1
ORDER BY id
LIMIT $2 OFFSET $3;

-- name: CountMemberships :one
SELECT COUNT(*) as total_memberships
FROM memberships;

-- name: GetMembershipByID :one
SELECT * FROM memberships
WHERE id = $1 AND created_by = $2;

-- name: CreateMembership :one
INSERT INTO memberships (
    membership_name, membership_length, cost, created_by
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateMembership :one
UPDATE memberships
    set membership_name = $4,
    membership_length = $5,
    cost = $6,
    version = version + 1
WHERE id = $1 AND created_by = $2 AND version = $3
RETURNING version;

-- name: DeleteMembership :execresult
DELETE FROM memberships
WHERE id = $1 AND created_by = $2;