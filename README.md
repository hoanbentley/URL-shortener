### Overview
This is a service which named [URL-shortener] for a shorten link, this service base-on clean architecture, right now this service can handle login/list/create/search/redirect functions.
To make it run:
- `go run main.go`
- Import Postman collection from `doc` to check example

### List API
- Login
- Redirect
- Create
- List
- Search
- Delete

![alt text](https://github.com/hoanbentley/URL-shortener/tree/main/doc/structure.png?raw=true)

### DB Schema
```sql
-- users definition

CREATE TABLE users (
	user_id TEXT NOT NULL,
	password TEXT NOT NULL
);

INSERT INTO users (user_id, password) VALUES('admin', 'admin');

-- urls definition

CREATE TABLE urls (
	id INTEGER PRIMARY KEY,
	short_code TEXT NOT NULL,
	full_url TEXT NOT NULL,
    created_date INTEGER,
	expiry INTEGER,
	number_of_hits INTEGER
);
```