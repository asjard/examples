FROM scratch:latest

WORKDIR /app

COPY ./bin/readme /app/server
COPY ./conf/ /app/conf/

EXPOSE 6030

ENTRYPOINT [ "/app/server" ]
