## Generating Code From SQL

If you want to modify schema, change in `db/migrations/001_init.up.sql`

If you want to add queries, change in `db/queries.sql`

Once changed, run `make gen`

## Update schema

Run `make drop` and then `make up`

## Entering Database

Run `make enterdb`