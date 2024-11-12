-- name: CreateBids :one
INSERT INTO bids ("product_id", "bidder_id", "bid_amount")
VALUES ($1, $2, $3)
RETURNING id;

-- name: GetBidsByProductId :many
SELECT * FROM bids WHERE product_id = $1
ORDER BY bid_amount DESC
LIMIT $2 OFFSET $3;

-- name: GetHighestBidByProductId :one
SELECT * FROM bids WHERE product_id = $1
ORDER BY bid_amount DESC
LIMIT 1;