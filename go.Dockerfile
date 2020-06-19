FROM golang:1.14-alpine as builder

ENV APP_USER app
ENV APP_HOME /app
WORKDIR $APP_HOME

# avoid using the root user in the docker image
# @see: https://semaphoreci.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker
# @see: https://www.thegeekdiary.com/run-docker-as-a-non-root-user/
RUN adduser --disabled-password $APP_USER
RUN chown -R $APP_USER $APP_HOME
USER $APP_USER

COPY home_go .

RUN go mod verify
RUN go build -o main

FROM alpine

ENV APP_USER app
ENV APP_HOME /app
WORKDIR $APP_HOME

RUN adduser --disabled-password $APP_USER
RUN chown -R $APP_USER $APP_HOME
USER $APP_USER

COPY certs certs
COPY static static
COPY --from=builder $APP_HOME/main .

CMD ["./main"]
