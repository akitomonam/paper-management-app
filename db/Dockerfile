FROM mysql:8

# ポートの開放
EXPOSE 3306

# Create user & add to sudoers
ARG UID
ARG GID
ARG UNAME

RUN groupadd -g ${GID} ${UNAME} && groupadd sudo  && \
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

ADD ./db/init.sql /docker-entrypoint-initdb.d/init.sql
ADD ./db/my.cnf /etc/mysql/conf.d/my.cnf
RUN chmod 644 /etc/mysql/conf.d/my.cnf

# 実行コマンド
CMD ["mysqld"]
