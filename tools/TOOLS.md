# tools

## create_migration.sh

This shell script is a quick way to generate a new SQL migration file. It expects a brief description of the migration as an argument. It creates a file in the project path with the following pattern:

`migrations/{seconds since unix epoch}-{slugified migration description}.sql`

It prepends the number of seconds since the unix epoch so that, when sorted, the migrations files are run in order of their creation.

### example

```bash
$ ./tools/create_migration.sh Add new column!
+ /workspaces/inventory/migrations/1700955979-add-new-column.sql
```
