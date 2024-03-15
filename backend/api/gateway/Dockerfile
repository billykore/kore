FROM devopsfaith/krakend
COPY . /etc/krakend
EXPOSE 3000
ENTRYPOINT ["/usr/bin/krakend", "run", "-d", "-c", "/etc/krakend/krakend.json" ]