## Cases

```
-- name: GetAccountForUpdate :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;
```

If we remove `NO KEY`, we will get the `deadlock` error.

When perform multiple tx concurrently, the `SELECT FROM ACCOUNT` in a `Account` table tx to get unlock from `INSERT INTO` in the `Transfer` table.

```
ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");
```

Because of the `foreign key` constrain, if we update the `account_foregin_key` on `Transfer` table, it will make the `Account` table `SELECT` query to wait until the `INSERT` query on `Transfer` table finish.
