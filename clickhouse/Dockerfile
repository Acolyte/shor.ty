FROM yandex/clickhouse-server:20.1.8.41

ADD . /etc/clickhouse-server
COPY ./dicts.tar.gz /etc/clickhouse-server/dicts.tar.gz
ADD ./init-dict.sh /docker-entrypoint-initdb.d/00-init-dict.sh
ADD ./init-dict.sql /docker-entrypoint-initdb.d/10-init-dict.sql
