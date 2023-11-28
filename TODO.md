# TODO

## unified build tools

I'm using [Bun](https://bun.sh/) right now for javascript dependency management. It would be nice if I didn't have to manually run `bun install` before running the various Go commands. I've seen people use Make. I could also make a bunch of custom commands in `package.json` and run everything from Bun.

It would be nicer to host lifecycle scripts from the `go` command so getting deps would presumably happen before the Go language server starts complaining about missing embedded asset files.

## static asset caching

I've added `Cache-Control: max-age=604800, stale-while-revalidate=86400` to static assets that serves as a best first try. It says stuff is cached for a week and the stale assets are okay to use for up to a day after being declared stale while the asset is fetched in the background. This is better than nothing since a fouc only happens on the first load, and revalidation can happen in the background, avoiding another fuoc when the cache expires.

I think a better solution would be to add a query param to each static asset URL with a value that's unique to the server build, like a hash or something. Then change cache-control header to be immutable. That's what MDN recommends as a common practice. It should altogether avoid the "hard refresh to see the update" problem. I'll need to get a page template system stood up to support this.

## automatic migrations

This requires any database support at all first.

This is probably going to be necessary. I need to look at other people's patterns for this. In HTTP mode, this makes perfect sense. But is this going to be a problem in CGI mode when this runs on every last request?

How can this be made as fast as possible so it doesn't matter if it runs per request? Make sure this logs when it happens.

## add stacktraces to errors

Zerolog has the ability to print out stacktraces of errors, but I guess errors don't have stacktraces by default in Go. Weird. I'll have to pick an error helper library to add stacktraces.

## dark mode

Bootstrap has a built-in dark mode now! Unfortunately, it doesn't automatically use the `prefers-color-scheme` media query.

Current options include:

1. Use javascript to detect `prefers-color-scheme` and set `data-bs-theme` in the DOM. 😞
2. Use Bootstrap's sass options and mixins to use the media query.

Neither is ideal. Both will require figuring out build pipeline stuff for frontend assets. I guess punt for now. This is more of a "nice to have".

## separate debugging symbols?

I read a thing recently about separating debugging symbols from your binary to make a smaller executable while still being able to debug production binaries. I don't know if Go supports this. C does. Rust does. Maybe Go does? I feel like this can wait at least until I have a Makefile or whatever to streamline build steps.

## make default environment type "production"

Once it makes any actual difference at all, default to production mode instead of development mode for security purposes.

Or possibly remove environment names altogether if it doesn't really matter.

## always turn on foreign_keys pragma

In order to enforce foreign key constraints, this pragma must be turned on for every db session.
