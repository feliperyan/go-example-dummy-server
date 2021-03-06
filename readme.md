# Simple Go server for demo purposes

## Backend mode

1. `/` = `GET` for a health json response.
2. `/quote` = `GET` for a random quote from a smart person. `POST` with json like `{quote: <string>}` to add a quote.
3. `/search` = `POST` with a json like `{word: <string>}`

---

## Frontend mode

If the environment variable `QUOTEAPIENDPOINT` is available then run the server in _frontend_ mode.

In _frontend_ mode the server will `HTTP GET` on `QUOTEAPIENDPOINT` to get a Quote and return it. Remember to set the full path to the server such as:

`export QUOTEAPIENDPOINT=http://localhost:8081/quote`

Point your browser to `/fetch_quote`

If the `AUDIENCE` env var is set, then the frontend is expected to be running on something like Cloud Run and to have access to the underlying service account in order to acquire a [Google-signed OpenID Connect token](https://cloud.google.com/run/docs/authenticating/service-to-service#set-up-sa). 

`export AUDIENCE=https://somebackend-abcdefgh-ts.a.run.app/` notice how it does not contain the path.

---

## CI

A couple of mostly dummy unit tests are available as well but very limited.

--- 

## OpenAPI

`openapi.yaml` can be used with API Gateway or CloudEndpoints, etc
