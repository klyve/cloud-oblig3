FROM golang:1.8

ENV PORT=$PORT

ENV MONGODB_HOST="ds157833.mlab.com"
ENV MONGODB_PORT="57833"
ENV MONGODB_DATABASE="cloud3"
ENV MONGODB_USERNAME="root"
ENV MONGODB_PASSWORD="HelloWorld123"


WORKDIR "/opt"

ADD .docker_build/cloud3 /opt/bin/cloud3
ADD ./recipes /opt/recipes

CMD ["/opt/bin/cloud3"]
