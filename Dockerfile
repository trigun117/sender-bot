FROM alpine

COPY sender-bot .

RUN apk update && apk add --no-cache ca-certificates

ENTRYPOINT [ "./sender-bot" ]