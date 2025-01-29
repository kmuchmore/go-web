once: templ tailwind esbuild

templ:
	templ generate -v

tailwind:
	npx @tailwindcss/cli -i ./view/index.css -o ./view/static/css/styles.min.css

esbuild:
	npx esbuild ./view/index.js --bundle --outdir=./view/static/js/

templ-live:
	templ generate --watch --proxybind="0.0.0.0" --proxy="http://0.0.0.0:3000" --open-browser=false -v

watch: 
	wgo -file .go go run cmd/main.go \
	:: wgo -dir view -file .go templ generate --proxy="http://0.0.0.0:3000" --notify-proxy \
	:: wgo -xdir view/static -file .css -file .templ npx @tailwindcss/cli -i ./view/index.css -o ./view/static/css/styles.min.css --minify \
	:: wgo -file view/index.js npx npx esbuild ./view/index.js --bundle --outdir=./view/static/js/ --minify \

live:
	make -j2 templ-live watch