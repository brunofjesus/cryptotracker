frontend_npm_install:
	cd frontend && npm install

frontend_npm_build: frontend_npm_install
	cd frontend && npm run build

bundle_frontend: frontend_npm_build
	cp -r frontend/dist/* backend/rest/handler/static

build_backend: bundle_frontend
	cd backend && go mod download && cd cmd && go build -o cryptotracker
	mv backend/cmd/cryptotracker .

build: build_backend
	rm -rf backend/rest/handler/static/*
	echo "<h1>Your frontend will be here after compilation</h1>" > backend/rest/handler/static/index.html