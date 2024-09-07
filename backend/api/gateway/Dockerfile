FROM devopsfaith/krakend
COPY ./api/gateway /etc/krakend
EXPOSE 3000
ENTRYPOINT ["/usr/bin/krakend", "run", "-d", "-c", "/etc/krakend/gateway.json" ]