# Simple Go server for demo purposes

1. `/` = `GET` for a health json response.
2. `/quote` = `GET` for a random quote from a smart person. `POST` with json like `{newquote: <string>}` to add a quote.
3. `/search` = `POST` with a json like `{word: <string>}`

A couple of unit tests are available as well but very limited.
