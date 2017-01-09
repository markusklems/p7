FROM google/debian:wheezy
ADD p7 /
EXPOSE 8080
CMD ["/p7"]
