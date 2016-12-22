
const env = process.env.LUCIFY_ENV === 'production' ? 'production' : 'staging';

module.exports = {
  REGION: 'eu-west-1',
  CLUSTER: 'minard',
  SERVICE: `minard-runner-${env}`,
  CONTAINER: 'runner',
  IMAGE: 'runner',
  BUCKET: 'lucify-configuration',
  KEY: `ecs_services/runner_${env}`,
  DOCKERFILE: 'Dockerfile',
};
