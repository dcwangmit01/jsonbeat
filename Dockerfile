FROM ubuntu:14.04

# stop warning messages from apt configuration on install
ENV TERM=linux
ENV DEBIAN_FRONTEND noninteractive

# install program deps
RUN apt-get -y install curl

# install the program
COPY ./bin/jsonbeat /root/jsonbeat
COPY ./etc/jsonbeat.yaml.private /root/jsonbeat.yaml

# create a run script that will load the elasticsearch template before starting the app
RUN echo "#!/bin/bash" > /root/run.sh && \
    echo "/root/jsonbeat -c /root/jsonbeat.yaml -e -v" >> /root/run.sh && \
    chmod a+x /root/run.sh

# set the startup command
CMD ["/root/run.sh"]
