FROM gcr.io/distroless/static:nonroot

ARG COMMIT=latest
ENV COMMIT=${COMMIT}

COPY service /bin/app

USER nonroot:nonroot

ENTRYPOINT ["app"]
