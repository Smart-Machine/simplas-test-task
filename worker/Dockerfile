FROM --platform=linux/amd64 alpine

RUN apk add libc6-compat

WORKDIR /worker

COPY ./worker/build/bin ./
COPY ./data/data.json ./

CMD ["/worker/bin"]