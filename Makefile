ifeq ($(OS),Windows_NT)
SHELL := powershell.exe
.SHELLFLAGS := -NoProfile -Command
endif
  
ifeq ($(OS),Windows_NT)

dir_backend = .\backend
else

dir_backend = ./backend
endif

backend-migration-re:
	(cd ${dir_backend} && make migration-re)

backend-run:
	(cd ${dir_backend} && make run)

backend-update-go-dd:
	(cd ${dir_backend} && make update-go-dd tag=v${tag})