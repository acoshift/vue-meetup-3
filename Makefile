NODE_BIN = node_modules/.bin
minify = $(NODE_BIN)/minify
browserify = $(NODE_BIN)/browserify
watchify = $(NODE_BIN)/watchify

.PHONY: component
component:
	NODE_ENV=production $(browserify) -t vueify -e component | $(minify) --js > script.js

dev-component:
	$(watchify) -p browserify-hmr -t vueify component -o script.js
