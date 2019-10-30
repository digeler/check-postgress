FROM ubuntu:16.04
WORKDIR /ps
COPY ./postgress /ps
RUN mkdir -p /var/logs/pslogs
CMD [ "/bin/bash", "-c" ,"/ps/postgress > /var/logs/pslogs/$HOSTNAME.txt 2>&1" ]
