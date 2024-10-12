FROM mysql:8.4

RUN set -eux \
    && ln -sf /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
    && echo "USE mysql;" > /docker-entrypoint-initdb.d/timezones.sql \
    && mysql_tzinfo_to_sql /usr/share/zoneinfo >> /docker-entrypoint-initdb.d/timezones.sql \
    && mkdir /var/log/mysql/ \
    && chown -R mysql:mysql /var/log/mysql/

COPY ./conf/my.cnf /etc/mysql/my.cnf