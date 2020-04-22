# Goober

## About

Goober can be used as standalone service or as a part of larger system. The primary idea behind this project is to provide extensive configuration options in order to meet wide range of use cases.

## Features

### Frontend

Goober comes with small and lightweight web frontend. Enabled by default.

### Security

By default Goober does not restrict access to it's apis but it can be configured to use Basic Auth (https highly advised), JWT and role based authorization.

If `jwt.type` and `jwt.key` are set in configuration file Goober is going to look for valid JWTs and deny unauthorized requests. Alternatively you can specyfy `jwt.jwk_url` In which case Goober is going to fetch public keys on startup and use them to validate requests. Set `jwt.issuer` to validate issuer.

In order to enable Basic Auth you just need to define local users in config file like so `admin: username:password:role`. Admin field can be an array. role field is optional.

If you want to use role based authorization you just need to define allowed roles in config file. If you do so Goober is going to look for `x-goober-role` claim in jwt payload (or check last value in `username:password:role` triplet). Role format looks like this: `rolename:permissions`. Valid permissions values are r - read, w - write, d - delete.

##### Note

If Basic Auth or JWT auth is enabled web panel and file serving endpoint access is going to be restricded. If this behaviour is not desirable set `skip_frontend_auth` or `skip_serving_auth` to true.

### Backend agnosticysm

Goober can be plugged into postgres, mysql, mssql or use sqlite3 by default. Set `db.type`, `db.host` `db.port`, `db.user`, `db.name`, `db.pass` as you need.

## Example configs

```yaml
host: 'localhost' # default: localhost
port: '9000' # default: 3000
domain: 'localhost:9000' # default: localhost:3000
upload_dir: ./data # default: ./data
frontend: true # default: true
skip_serving_auth: true # default: false
admin:
  - 'admin:admin'
  - 'user:user'
```

```yaml
host: '0.0.0.0'
port: '9000'
domain: 'localhost:9000'
upload_dir: ./data
frontend: false
jwt:
  type: 'HS256'
  key: 'secret'
  # jwk_url: 'http://localhost:4444/.well-known/jwks.json'
  # issuer: ''
roles:
  - 'admin:rwd'
  - 'user:rw'
db:
  type: 'postgres'
  host: 'localhost'
  name: 'goober'
  pass: 'postgres'
  port: '5432'
  user: 'postgres'
```
