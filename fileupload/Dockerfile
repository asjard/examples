FROM scratch:latest

WORKDIR /app

COPY ./bin/fileupload /app/server
COPY ./conf/ /app/conf/

EXPOSE 7030

ENTRYPOINT [ "/app/server" ]
