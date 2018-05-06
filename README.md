# hatchery
Backend service to power hatch!

## Using Docker Compose

To run the app, run `docker-compose up` - this will spin up the server, a Postgre DB and a redis server. 

To rebuild the app, run `docker-compose up --build`

## GO Tools

### Twirp
As a dev dependency you will need to install retool and protobuf locally.
retool:
`go get github.com/twitchtv/retool`
`retool sync`

protobuf:
`homebrew install protobuf` for mac

for other users here is the link: https://developers.google.com/protocol-buffers/docs/gotutorial 

## Managing Secrets

In the root of the app, add a `secrets` directory with a `postgres_password` document (no extension). Enter the password with no whitespaces around it. 

--

### Notes on the UI-Routing

UI routing is strict - for static files to be accessible from the web, place them in src/assets exclusively.
