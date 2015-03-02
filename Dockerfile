# AUTHOR:         Joshua Mervine <joshua@mervine.net>
# DESCRIPTION:    Docker image for github.com/jmervine/gitfish
FROM progrium/busybox:latest
COPY gitfish /usr/bin/

CMD ( test -f /gitfishrc || \
      ( echo "standard run requires gitfishrc mounted"; exit 1 ) \
    ) && \
    ( source /gitfishrc; gitfish )
