FROM google/debian:wheezy
WORKDIR /app
ADD api.tar /app/
COPY api /app/
EXPOSE 8888
ENTRYPOINT ["./api"]
