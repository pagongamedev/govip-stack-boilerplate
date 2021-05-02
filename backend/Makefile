ifeq ($(OS),Windows_NT)
SHELL := powershell.exe
.SHELLFLAGS := -NoProfile -Command
endif
  
ifeq ($(OS),Windows_NT)
build:
	$$env:GO111MODULE="on"; $$env:API_ENV="localhost";

compile:
	$$env:GO111MODULE="on"; $$env:API_ENV="localhost";
	go build -o app.exe .\main.go; .\app.exe

run:
	$$env:GO111MODULE="on"; $$env:API_ENV="localhost";
	go run .\main.go

dir_migration_pre = .\domain\
dir_migration_post = "\migration\postgres"
else

build:
	env GO111MODULE=on go build -o app main.go;

compile:
	env GO111MODULE=on go build -o app main.go;
	env API_ENV=localhost ./app

run:
	env GO111MODULE=on API_ENV=localhost go run main.go

dir_migration_pre = ./domain/
dir_migration_post = /migration/postgres
endif

migration-re:
	(cd ${dir_migration_pre}auth${dir_migration_post} && sql-migrate down -limit=0 -env=localhost && sql-migrate down -limit=0 -env=seed);
	(cd ${dir_migration_pre}auth${dir_migration_post} && sql-migrate up -env=localhost && sql-migrate up -env=seed)

migration-down:
	(cd ${dir_migration_pre}auth${dir_migration_post}  && sql-migrate down -limit=0 -env=localhost && sql-migrate down -limit=0 -env=seed);

migration-up:
	(cd ${dir_migration_pre}auth${dir_migration_post}  && sql-migrate up -env=localhost && sql-migrate up -env=seed)

migration-re-integration:
	(cd ${dir_migration_pre}auth${dir_migration_post}  && sql-migrate down -limit=0 -env=localhost && sql-migrate down -limit=0 -env=seed-integration);
	(cd ${dir_migration_pre}auth${dir_migration_post}  && sql-migrate up -env=localhost && sql-migrate up -env=seed-integration)
