FROM base

RUN mkdir -p /app
ENV XDIR=/app
ENV XKEY=CI_SECRET

COPY . /app
WORKDIR /app/swagger
RUN go build -o swagger

ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh ./wait-for-it.sh
RUN ["chmod", "+x", "./wait-for-it.sh"]

EXPOSE 8080 9080
ENTRYPOINT ["./swagger"] 
