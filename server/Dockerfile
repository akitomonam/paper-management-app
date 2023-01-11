FROM golang:1.19

# Create user & add to sudoers
ARG UID
ARG GID
ARG UNAME

RUN groupadd -g ${GID} ${UNAME} && \
    useradd -u ${UID} -g ${UNAME} -G sudo  -m ${UNAME} && \
    echo "${UNAME}:${UNAME}" | chpasswd && \
    echo "${UNAME} ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers

ENV UID ${UID}
ENV GID ${GID}
ENV UNAME ${UNAME}


# Change Permission
COPY ./set_volume_permission.sh /home/${UNAME}/
RUN chmod +x /home/${UNAME}/set_volume_permission.sh

RUN /home/${UNAME}/set_volume_permission.sh

# コンテナ内に作業ディレクトリを作成
RUN mkdir -p /akitomonam/src/github.com//paper-management-app
# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/github.com/akitomonam/paper-management-app/server

COPY ./server/main.go ./

CMD ["go", "run", "main.go"]
