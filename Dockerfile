FROM golang

ENV DIAG_PORT 8081
ENV SERVICE_PORT 8080

EXPOSE $SERVICE_PORT
EXPOSE $DIAG_PORT

COPY ./bin/linux-amd64/k8s-gopherconeu /
CMD ["/k8s-gopherconeu"]
