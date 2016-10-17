FROM golang

RUN apt-get update && apt-get install -y libyara-dev git

RUN mkdir -p /go/src/yarascanner

ADD *.go /go/src/yarascanner/
RUN go get yarascanner
RUN go install yarascanner

RUN git clone https://github.com/Yara-Rules/rules.git /rules
RUN mkdir /uploads

ENV IPADDR 0.0.0.0
ENV PORT 9999
ENV UPLOADS /uploads
ENV INDEXES -i /rules/malware_index.yar

EXPOSE ${PORT}
CMD /go/bin/yarascanner -address ${IPADDR} -port ${PORT} -uploads ${UPLOADS} ${INDEXES}
