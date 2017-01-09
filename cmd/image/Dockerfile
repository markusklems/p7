FROM alpine:3.5
DOCKER_API_VERSION=1.24
RUN groupadd -r image && useradd -r -g image image
USER image
WORKDIR /app
ADD image.tar /app/
COPY image /app/
EXPOSE 8890
ENTRYPOINT ["./image"]
