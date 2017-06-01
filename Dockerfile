FROM alpine
MAINTAINER Say.li <120011676@qq.com>

# 原utc时间，修改成cst（中国标准东八区时间）
ENV TZ Asia/Shanghai
RUN apk update && apk add ca-certificates && \
    apk add tzdata && \
    ln -sf /usr/share/zoneinfo/${TZ} /etc/localtime && \
    echo ${TZ} > /etc/timezone


ENV WORK_DIR /opt/

COPY go_practice ${WORK_DIR}

EXPOSE 8080
WORKDIR ${WORK_DIR}
CMD ./go_practice
