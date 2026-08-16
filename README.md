# Self Hosted Link Shortener

![Go Version](https://img.shields.io/badge/Go-1.25-00ADD8?logo=go&logoColor=white)
![Fiber](https://img.shields.io/badge/Fiber-v3.4.0-00ADD8)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Docker](https://img.shields.io/badge/Docker-ready-2496ED?logo=docker&logoColor=white)

This is a Link Shortener application that is completely self-hosted using a Raspberry Pi and [Dokploy](https://dokploy.com/). The logic behind this link shortener is very simple and not different from others; the only reason for its development was that some link shorteners are now charging for their services, and, come on! There is a POST method on the "/" route that is protected by basic authentication via headers (`Authorization: "Basic ${Base64(login+password)}"`); you can see it at: https://link.vanin.dev/docs/

#### **Example**: 
```bash
> curl -X 'POST' \
    'http://localhost:8080/' \
    -H 'accept: application/json' \
    -H 'Auth: Basic Base64(loginpassword)' \
    -H 'Content-Type: application/json' \
    -d '{
      "description": "string",
      "name": "string",
      "url": "string"
    }'
```

> [!NOTE]
> Stack:
> - Golang ~ v1.25
>   - Fiber ~ v3.4.0
> - Postgres
> - Docker


> [!TIP]
> The Swagger/OpenAPI document served at `/docs` is generated with [openapi-builder](https://github.com/LuigiVanin/openapi-builder), a Go library of mine for building OpenAPI documents programmatically.

## Environment Variables

The application is configured via the following environment variables (see [.env.example](.env.example)):

| Variable        | Description                                                        | Required |
| --------------- | ------------------------------------------------------------------ | -------- |
| `USER_LOGIN`     | Login used for the Basic authentication credentials on the POST route | Yes      |
| `USER_PASSWORD`  | Password used for the Basic authentication credentials on the POST route | Yes      |
| `DATABASE_URL`   | Postgres connection string used to store the shortened links       | Yes      |

## Installation & Development

1. Clone the repository and copy [.env.example](.env.example) to `.env`, filling in the variables described above.
2. Have a Postgres instance available and pointed to by `DATABASE_URL`.
3. Run the server using `make`, which automatically injects the `.env` file into the environment:

```bash
> make run   # runs the server once
> make dev   # runs the server with air, reloading on changes
```

Alternatively, without make, the `.env` variables need to be exported manually:

```bash
> export $(cat .env | xargs) && go run ./cmd

# Or using air for development hot reloading
> export $(cat .env | xargs) && air
```

## Docker

To run the project using docker it is necessary to build the image first: 

```bash
> docker build . -t link-shortener
```

Afterwards, the only thing necessary is to run the image, specifying the port to be mapped on the host machine and configuring the .env file you will be using for the deploy (the specification of the file can be seen in [.env.example](.env.example)):

```bash
> docker run -d --name link-shortener --env-file .env -p HOST:8080 link-shortener
```

## License

[MIT](./LICENSE.md)