# Build stage
FROM golang:1.15.10-alpine3.12 AS builder


COPY ${PWD} /app
WORKDIR /app


# Toggle CGO based on your app requeirment. CGO_ENABLED=1 for enabling CGO

RUN CGO_ENABLED=0 go build -ldflags '-s -w -extldflags "-static"' -o /app/appbin -x *.go
# RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o bin/boolang  *.go





# Final stage
FROM alpine:3.12
LABEL MAINTAINER Author lorezi

# Following commands are for installing CA certs (for proper functioning of HTTPS and other TLS)
RUN apk --update add ca-certificates && \
    rm -rf /var/cache/apk/*

# Add new user 'appuser'
RUN adduser -D appuser
USER appuser

COPY --from=builder /app /home/appuser/app

WORKDIR /home/appuser/app

# Since running as non-root user, port bindings < 1024 in not possible
# 8000 for HTTP; 8443 for HTTPS;
EXPOSE 8080
EXPOSE 8443

CMD [ "./appbin" ]

