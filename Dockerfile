# See compatilibity chart at
# https://gitlab.com/gitlab-org/gitlab-ci-multi-runner/
FROM gitlab/gitlab-runner:alpine-v1.11.4

RUN apk add --update postgresql
COPY ./docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod a+x /docker-entrypoint.sh

ENTRYPOINT ["/usr/bin/dumb-init", "/docker-entrypoint.sh"]
CMD ["gitlab-runner", "run"]
