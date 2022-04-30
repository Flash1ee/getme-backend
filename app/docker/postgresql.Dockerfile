FROM postgres:14

RUN apt-get update && apt-get install -y --no-install-recommends apt-utils
RUN apt-get update \
      && apt-cache showpkg postgresql-$PG_MAJOR-rum \
      && apt-get install -y --no-install-recommends \
           postgresql-$PG_MAJOR-rum  postgresql-$PG_MAJOR-rum-dbgsym\
      && rm -rf /var/lib/apt/lists/*

RUN apt-get update && apt-get install -y git
RUN apt-get update && apt-get install -y make
RUN git clone https://github.com/postgrespro/hunspell_dicts
WORKDIR /hunspell_dicts/hunspell_ru_ru
RUN make USE_PGXS=1 install
RUN rm -rf ./hunspell_dicts

WORKDIR /


RUN mkdir -p /docker-entrypoint-initdb.d