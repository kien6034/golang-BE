#

## Accounts

```
Table accounts as A {
  id bigserial [pk]
  owner varchar [not null]
  balance bigint [not null]
  currency varchar [not null]
  created_at timestamptz [not null, default: `now()`]
}
```

## Entries

The next table is entries. This table will record all changes to the account balance.

```

Table entries {
    id bigserial [pk]
    account_id bigint [ref: > A.id]
    amount bigint
    created_at timestamptz [default: `now()`]
}
```

## Transfers

It records all the money transfers between 2 accounts.

```
Table transfers {
  id bigserial [pk]
  from_account_id bigint [ref: > A.id]
  to_account_id bigint [ref: > A.id]
  amount bigint
  created_at timestamptz [default: `now()`]
}
```
