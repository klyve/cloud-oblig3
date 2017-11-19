FROM golang

ENV PORT=$PORT

ENV MONGODB_HOST="ds157833.mlab.com"
ENV MONGODB_PORT="57833"
ENV MONGODB_DATABASE="cloud3"
ENV MONGODB_USERNAME="root"
ENV MONGODB_PASSWORD="HelloWorld123"
# mongodb://<dbuser>:<dbpassword>@ds157833.mlab.com:57833/cloud3


# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/klyve/cloud-oblig3

# Download the dependencies
RUN go get github.com/caarlos0/env
RUN go get github.com/gorilla/mux
RUN go get github.com/robfig/cron
RUN go get gopkg.in/mgo.v2

# Install the package
RUN go install github.com/klyve/cloud-oblig3



# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/cloud-oblig3

# Document that the service listens on port 8080.
# Uncomment this for other uses than heroku
# EXPOSE $PORT