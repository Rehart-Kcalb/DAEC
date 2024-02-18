-- name: InsertTaken :exec
INSERT INTO taken_expressions (expression_id,agent,calculator)
VALUES ($1,$2,$3);

-- name: DeleteTaken :exec
DELETE FROM taken_expressions
WHERE expression_id = $1;

-- name: ClearTakenForAgent :exec
DELETE FROM taken_expressions
WHERE agent = $1;

-- name: GetAgentExpressions :many
SELECT *
FROM taken_expressions
WHERE agent = $1;

-- name: GetAgents :many
SELECT agent from taken_expressions;

-- name: InsertExpression :one
INSERT INTO expressions (expression, status, created_at)
VALUES ($1, $2, $3)
RETURNING id;
 
-- name: UpdateOperation :exec 
UPDATE operations 
SET cost = $2
WHERE operation = $1;

-- name: InsertSubExpression :exec
INSERT INTO sub_expressions(expression_id,operand1,operation,operand2,cost,exec_order)
VALUES ($1,$2,$3,$4,$5,$6);

-- name: InputResultSubExpression :exec
UPDATE sub_expressions
SET result = $2 , status = 'done'
WHERE id = $1;

-- name: InputResultExpression :exec 
UPDATE expressions
SET result = $2, status = 'done' , calculated_at = now()
WHERE id = $1;

-- name: GetSubexpressions :many
SELECT se.*
FROM sub_expressions se
INNER JOIN expressions e ON se.expression_id = e.id
WHERE e.id = $1;

-- name: GetExpressions :many
SELECT * from expressions;

-- name: GetExpression :one
SELECT * FROM expressions
WHERE id = $1;

-- name: GetOperations :many
SELECT * FROM operations;

-- name: GetTask :one
SELECT * FROM expressions
WHERE status = 'wait' LIMIT 1;

-- name: UpdateExpressionStatus :exec
UPDATE expressions SET status = $2 WHERE id = $1;

-- name: Monitor :many
SELECT t.agent,t.calculator,e.expression,e.status 
FROM taken_expressions t 
RIGHT JOIN expressions e ON expressions.id = t.expression_id;
