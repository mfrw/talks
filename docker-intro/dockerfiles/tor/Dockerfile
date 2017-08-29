FROM alpine:latest
LABEL maintainer "Muhammad Falak R Wani <falakreyaz@gmail.com>"
RUN apk --update add tor && adduser -D anon
COPY torrc /etc/tor/torrc
EXPOSE 9050 9051
USER anon
CMD [ "tor" ]
