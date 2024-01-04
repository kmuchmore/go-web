prep:
	@templ generate
	@npx postcss-cli ./view/base.css -o ./view/static/css/base.min.css --minify
	# @npx tailwind -i ./view/base.css -o ./view/static/css/base.css --minify
	# @esbuild ./view/base.js --bundle --minify --outfile=./view/static/js/dist.min.js
	@esbuild ./view/base.js --bundle --minify --outbase=view --outdir=./view/static/js/ --analyze

run: prep
	@go run cmd/main.go