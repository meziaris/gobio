# GoBio API

GoBio: Link in bio tool built on GoLang, inspired by Linktr.ee

## Framework
- Web: Echo
- Middleware: JWT
- Configuration : GoDotEnv
- Database: PostgreSQL

### Run in Docker

First, install [Docker](https://docs.docker.com/get-docker/) and
[Docker Compose](https://docs.docker.com/compose/install/).

Then run the following command.
```bash
./.maintain/docker/create-image.sh
./.maintain/docker/start-docker-compose.sh
```
