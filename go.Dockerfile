FROM golang:1.14-alpine as builder

# @see: https://semaphoreci.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker
# @see: https://www.thegeekdiary.com/run-docker-as-a-non-root-user/
ENV APP_USER app
ENV APP_HOME /app
WORKDIR $APP_HOME

RUN adduser --disabled-password $APP_USER
RUN chown -R $APP_USER $APP_HOME
USER $APP_USER

COPY home_go .

RUN go mod verify
RUN go build -o hello

FROM alpine

ENV APP_USER app
ENV APP_HOME /app
WORKDIR $APP_HOME

RUN adduser --disabled-password $APP_USER
RUN chown -R $APP_USER $APP_HOME
USER $APP_USER

COPY --from=builder $APP_HOME/hello .

CMD ["./hello"]
