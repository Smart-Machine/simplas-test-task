FROM --platform=linux/amd64 alpine

RUN apk add libc6-compat

WORKDIR /service

COPY ./service/build/bin ./

EXPOSE 8000

CMD ["/service/bin"]