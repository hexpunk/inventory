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

## different logging library?

I'm enjoying [zerolog](https://github.com/rs/zerolog), but I'm a little concerned that since its API differs from the built-in log library, that's going to make integration with other libraries and middlewares more painful than it should be. Maybe look for alternatives?

## support for no javascript (?)

HTMX has the `hx-boost` directive to provide itself as a progressive enhancement and fall back to full server-side rendering. This might be nice to have, although I don't know if it's nice enough to justify the logic that would have to be built into the template system to distinguish between boosted and non-boosted views. I don't want to need two templates per view. I don't want to have to be mindful every time I add a new view. If a pattern can support this use case without any cost to adding new features, it might be worth the geek cred to be able to use the site on a potato.