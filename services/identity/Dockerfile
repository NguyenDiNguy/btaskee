FROM base

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GO111MODULE=on

RUN mkdir -p /app
ENV XDIR=/app
ENV XKEY=CI_SECRET

RUN apk add --no-cache git
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache bash

COPY . /app
WORKDIR /app/services/identity
RUN go build -o identity

ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh ./wait-for-it.sh
RUN ["chmod", "+x", "./wait-for-it.sh"]

EXPOSE 8080 9080
ENTRYPOINT ["./identity"] 
