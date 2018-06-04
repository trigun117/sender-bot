FROM alpine

COPY sender-bot .

ENTRYPOINT [ "./sender-bot" ]