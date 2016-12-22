
# Minard runner

Dockerized version [gitlab-ci-multi-runner](https://gitlab.com/gitlab-org/gitlab-ci-multi-runner)
for use with [minard-backend](https://github.com/lucified/minard-backend).

## About

This wraps [gitlab/gitlab-runner:alpine](https://hub.docker.com/r/gitlab/gitlab-runner/)
with an entrypoint that supports automatic registration
of the runner instance on startup. This makes it
easier to start fresh development environments or deploy new
runners in production environments.

## License

[MIT Expat license](LICENSE)

## Lucify-specific documentation

[Documentation](lucify-docs.md) for Lucify's internal use.
