FROM scratch
LABEL version="1.1"
LABEL maintainer="mario-leander.reimer@qaware.de"

EXPOSE 9090
ENTRYPOINT ["/app"]

COPY app /app