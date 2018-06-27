FROM alpine
MAINTAINER wzy redis-stat
ENV TIME_ZONE=Asia/Shanghai
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2 && \
    ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone
COPY bin/linux/redis-stat /redis-stat
EXPOSE 10012
ENTRYPOINT ["/redis-stat", "--server=8080"]

