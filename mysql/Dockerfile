FROM mysql:8.0.23

ENV MYSQL_ROOT_PASSWORD rootpass

ENV MYSQL_USER souhub

ENV MYSQL_PASSWORD devpass

ENV MYSQL_DATABASE zeus

COPY actresses.sql /docker-entrypoint-initdb.d

COPY wemen.sql /docker-entrypoint-initdb.d

COPY zeus.cnf /etc/mysql/conf.d

EXPOSE 3306
