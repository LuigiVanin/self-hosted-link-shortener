## Docker

```bash
> docker build . -t link-shortener
```

```bash
> docker run -d --name link-shortener --env-file .env -p HOST:CONTAINER link-shortener
```