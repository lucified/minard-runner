FROM gitlab/gitlab-runner:alpine

RUN apk add --update postgresql
COPY ./docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod a+x /docker-entrypoint.sh

ENTRYPOINT ["/usr/bin/dumb-init", "/docker-entrypoint.sh"]
CMD ["gitlab-runner", "run"]
