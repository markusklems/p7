FROM alpine:3.5
RUN adduser -D -u 142 image
USER image
WORKDIR /app
ADD image.tar /app/
COPY image /app/
EXPOSE 8890
ENTRYPOINT ["./image"]
