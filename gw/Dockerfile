FROM scratch:latest

WORKDIR /app

COPY ./bin/server /app/server
COPY ./conf/ /app/conf/

EXPOSE 8030

ENTRYPOINT [ "/app/server" ]
