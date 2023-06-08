frontend_npm_install:
	cd frontend && npm install

frontend_npm_build: frontend_npm_install
	cd frontend && npm run build

bundle_frontend: frontend_npm_build
	cp -r frontend/dist/* backend/rest/handler/static

backend_dependency:
	cd backend && go mod download

build_backend: backend_dependency
	cd backend && GOOS=$(system) GOARCH=$(arch) go build -o out
	mv backend/out bin/cryptotracker_$(system)_$(arch)$(extension)
	zip bin/cryptotracker_$(system)_$(arch)_v$(version).zip bin/cryptotracker_$(system)_$(arch)$(extension)

build:
	make bundle_frontend
	mkdir bin
	make build_backend system=windows arch=amd64 extension=.exe
	make build_backend system=windows arch=386 extension=.exe
	make build_backend system=darwin arch=amd64
	make build_backend system=darwin arch=arm64
	make build_backend system=linux arch=386
	make build_backend system=linux arch=amd64
	make build_backend system=linux arch=arm64

	rm -rf backend/rest/handler/static/*
	echo "<h1>Your frontend will be here after compilation</h1>" > backend/rest/handler/static/index.html