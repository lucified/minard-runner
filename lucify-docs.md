
# Internal notes

## Deployment

### Locally

Local deployments require that:
- Valid AWS credentials are in place.
- [ecs-updater](https://github.com/lucified/ecs-updater) is installed globally
- [lucify-notifier](https://github.com/lucified/lucify-notifier) is installed globally

Install ecs-updater and lucify-notifier globally with
```bash
npm install -g ecs-updater lucify-notifier
```

Deploy to staging with
```bash
LUCIFY_ENV=staging lucify-notifier ecs-updater
```

Deploy to production with
```bash
LUCIFY_ENV=production lucify-notifier ecs-updater
```

### Continous integration

Once pull requests are merged to the `lucified` branch
in GitHub, on the condition that the tests pass, CircleCI
will deploy to staging automatically.

### Chatops

Chatops deployments are supported.

## Restart

Restart `minard-runner` in staging with
```bash
ecs-updater -s restart-terraform
```
Restart `minard-runner` in production with
```bash
LUCIFY_ENV=production ecs-updater -s restart-terraform
```
