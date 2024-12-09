FROM golang:alpine AS build
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /bin/yung-gpt-cli ./main.go

FROM alpine:3.20
RUN apk --no-cache add ca-certificates tzdata && \
    update-ca-certificates

ARG UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    appuser
USER appuser

COPY --from=build /bin/yung-gpt-cli /bin/

ENTRYPOINT [ "/bin/yung-gpt-cli" ]