FROM --platform=$BUILDPLATFORM golang AS builder

WORKDIR /usr/src/app

ENV CGO_ENABLED=0

COPY . .

ARG TARGETOS TARGETARCH
RUN --mount=target=. \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg \
    GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /usr/local/bin/greedygame ./cmd/server/main.go

CMD ["greedygame"]

#stage 2
FROM alpine AS runner

RUN apk --no-cache add ca-certificates

#copy local file so we can also run the locally built container directly
COPY --from=builder /usr/src/app/etc /usr/src/app/etc
COPY --from=builder /usr/local/bin/greedygame /

CMD ["/greedygame"]