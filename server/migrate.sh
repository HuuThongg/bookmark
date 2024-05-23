#?/bin/bash
source .env
cd db/migrations && goose postgres "$dbString" up
