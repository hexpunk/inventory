package main

import "embed"

//go:embed node_modules/bootstrap/dist/js/bootstrap.js
//go:embed node_modules/bootstrap/dist/js/bootstrap.min.js
//go:embed node_modules/bootstrap/dist/css/bootstrap.css
//go:embed node_modules/bootstrap/dist/css/bootstrap.min.css
//go:embed node_modules/htmx.org/dist/htmx.js
//go:embed node_modules/htmx.org/dist/htmx.min.js
var nodeModules embed.FS
