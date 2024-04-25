FROM --platform=linux/amd64 alpine

RUN apk add libc6-compat

WORKDIR /httpProxy

COPY ./httpProxy/build/bin ./

EXPOSE 8001

CMD ["/httpProxy/bin"]