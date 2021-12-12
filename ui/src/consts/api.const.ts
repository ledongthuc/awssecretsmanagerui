const GW = process.env.GW;

export const ENDPOINTS = {
  getRegions: `${GW}/api/aws/regions`,
  getSecretList: `${GW}/api/secrets`,
  getSecretDetail: `${GW}/api/secrets/detail`,
  getSecretValue: `${GW}/api/secrets/value`,
  updateSecretValue: (region: string) =>
    `${GW}/api/secrets/value/update?region=${region}`,
  updateSecretBinary: (arn: string, region: string) =>
    `${GW}/api/secrets/value/upload?arn=${arn}&region=${region}`
};
