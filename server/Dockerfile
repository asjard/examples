FROM scratch:latest

WORKDIR /app

COPY bin/server_linux /app
COPY conf/ /app/conf/

EXPOSE 7030
EXPOSE 7031
EXPOSE 7032

ENTRYPOINT [ "/app/server_linux" ]
