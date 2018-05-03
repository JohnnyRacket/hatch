# hatchery
Backend service to power hatch!

## Using Docker Compose

To run the app, run `docker-compose up` - this will spin up the server, a Postgre DB and a redis server. 

To rebuild the app, run `docker-compose up --build`

## Managing Secrets

In the root of the app, add a `secrets` directory with a `postgres_password` and a `pgadmin_password` document (no extension). Enter the password with no whitespaces around it. ** NEVER COMMIT THIS JOHN **