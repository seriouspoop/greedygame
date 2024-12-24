-- name: GetAllTargetingRules :many
SELECT * FROM targeting_rule;

-- name: GetTargetRuleByID :one
SELECT * FROM targeting_rule
WHERE cid = $1;