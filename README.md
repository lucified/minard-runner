
# GitLab

Fork of [gitlab-ci-multi-runner](https://github.com/ayufan/gitlab-ci-multi-runner) for
use with [minard-backend](https://github.com/lucified/minard-backend)

## Changes

### Build configuration

We have made modification that allow the runner to
autoregister with GitLab on startup. This makes it
easier to start fresh development environments or deploy new
runners in production environments. This has been mostly implemented
in [docker-entrypoint.sh](docker-entrypoint.sh).

### Deployment

We have added some files that help deploy the project
into Lucify's internal environments.

## License

[MIT Expat license](LICENSE)

## Lucify-specific documentation

[Documentation](lucify-docs.md) for Lucify's internal use.
