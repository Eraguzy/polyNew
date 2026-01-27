# PolyNew

This repository is the backend behind this [Telegram channel](https://t.me/isapolynew), which consists of a feed of notifications for new events having the selected tags on Polymarket.

PM [@eragyoza](https://t.me/eragoza) for requests.

## API
An HTTP server is listening on Koyeb [here](https://sophisticated-aloisia-isabellesdesk-6e0ebfef.koyeb.app/).

Endpoints :

`/fetch-store-tag` => Get a tag by its slug and store it. By default the tag will be set as tracked.

`/set-tracked-tag` => Set a stored tag as tracked/not tracked. If tracked, new events related to the tag will be outputted by `/compare-events`

`/notify-tracked-tags` => Triggers an alert displaying tracked tags.

`/compare-events` => Compares events to detect new items, and sends a Telegram alert if new events are detected. A goroutine constantly calls the function associated to this route to refresh events when the server is up.

## Services
Uses Supabase (Postgres) to host the DB.

Uses Koyeb to host the HTTP server.

Uses a Github Action to poll the koyeb server (and keep it awake via `/keepalive`).