FROM scratch
MAINTAINER Lanre Adelowo<yo@lanre.wtf>
ADD mapped mapped
EXPOSE 8090 1800

ENTRYPOINT ["/mapped"]
