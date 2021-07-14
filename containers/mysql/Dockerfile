FROM mysql:5.7

# Japanese
RUN apt-get update \
    && apt-get install -y locales \
    && locale-gen ja_JP.UTF-8

ENV LANG="ja_JP.UTF-8" \
    TZ="Asia/Tokyo"

CMD ["mysqld", "--character-set-server=utf8mb4", "--collation-server=utf8mb4_unicode_ci"]