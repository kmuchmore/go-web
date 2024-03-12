dev:
	@templ generate --watch --proxy="http://localhost:3000" --cmd="make devrun"

build:
	@go build -o cmd/server cmd/main.go

templ:
	@templ generate

postcss:
	@npx postcss-cli ./view/base.css -o ./view/static/css/base.min.css --minify

esbuild:
	@esbuild ./view/base.js --bundle --minify --outbase=view --outdir=./view/static/js/ --analyze

devrun: postcss esbuild
	@go run cmd/main.go

run: templ postcss esbuild
	@go run cmd/main.go