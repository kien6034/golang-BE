## Installation

https://github.com/kyleconroy/sqlc

## Stepts

```
sqlc generate
```

### DBTX breakdown

- ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error): This method executes a query without returning any rows. It's typically used for queries that don't return rows, such as INSERT, UPDATE, or DELETE.

- PrepareContext(ctx context.Context, query string) (\*sql.Stmt, error): This method creates a prepared (or parameterized) statement for later execution. This can provide performance improvements because the database can optimize the execution plan for this query.

- QueryContext(ctx context.Context, query string, args ...interface{}) (\*sql.Rows, error): This method executes a query that returns rows, typically a SELECT.

- QueryRowContext(ctx context.Context, query string, args ...interface{}) \*sql.Row: This method is a convenience wrapper over QueryContext(). It executes a query that is expected to return at least one row, but it only returns the first row from the result set.
