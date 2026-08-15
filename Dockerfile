FROM golang:1.25.3-alpine AS builder

# Folder where the content of the docker project will be stored when container is created
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Build the binary
# CGO_ENABLED=0 creates a statically linked binary, which is better for alpine/scratch
RUN CGO_ENABLED=0 GOOS=linux go build -o link-shortener ./cmd/main.go

# Setting up multi stage build so the final container is smaller
FROM alpine:latest AS runner

WORKDIR /app

COPY --from=builder /app/link-shortener .

EXPOSE 8080

CMD [ "./link-shortener" ]