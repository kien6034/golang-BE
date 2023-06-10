## What is a Db transaction?

Provide a reliable and consistent unit of work, even in case of system failure.

## ACID properties

1. Atomicity: all or nothing
2. Consistency: only valid data is written
3. Isolation: concurrent transactions do not interfere with each other
4. Durability: once committed, data is committed forever

## A DB transaction in the simple bank example

A user create a transaction to transfer money from `account A` to `account B`.
Here is what tx should look like

1. an ix to create the `transfer` from A to B
2. an ix to create `entry` for `A`
3. an ix to create `entry` for `B`
4. an ix to update `account A` balance
5. an ix to update `account B` balance
