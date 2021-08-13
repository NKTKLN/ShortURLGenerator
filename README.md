# Short URL Generator (shurl.ml)

This is a simple project for understanding how to work with GIN, PostgresSQL and Redis in golang. There is a Russian version of this project on YouTube for more clarification.

Playlist: [click](https://youtube.com/)

## Running this on the server

### Prerequisites

- [Docker CE](https://docs.docker.com/engine/install/)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Edit this in the env/app.env

```
VIRTUAL_HOST=yourdomain.com
LETSENCRYPT_HOST=yourdomain.com
LETSENCRYPT_EMAIL=your@mail.com
```

### Edit this in the env/email.env

```
SENDER_EMAIL=your@mail.com
EMAIL_PASSWORD=mysecretpassword
SMTP_ADDRESS=smtp.google.ru
```

### Edit this in the env/bot.env

You can get the token for the bot [here](https://t.me/BotFather).

```
API_TOKEN=0123456789:abcdefghijklmnopqrstuvwxyz123456789
BOT_USERNAME=mybotname_bot
```

### And then run it with this command

```
docker-compose up -d --build
```

## Change the PG administrator's data

### Edit this in the env/pg.env

```
POSTGRES_USER=postgres
POSTGRES_PASSWORD=mysecretpassword
POSTGRES_DB=shorturldb
```

And recreate the container if it has already been created, otherwise run it.

#

### Copyright © 2021 [NKTKLN](https://github.com/NKTKLN).
