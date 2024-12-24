-- name: GetCampaignFromCIDs :many
SELECT * FROM campaign 
WHERE cid = ANY(sqlc.arg(cids)::UUID[]); 