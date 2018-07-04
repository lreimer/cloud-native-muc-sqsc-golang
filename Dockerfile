FROM scratch
LABEL version="1.0"
LABEL maintainer="mario-leander.reimer@qaware.de"

EXPOSE 9090
ENTRYPOINT ["/app"]

COPY app /app