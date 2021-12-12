const gateway = process.env.SERVER;

if (!gateway) throw Error('SERVER environment variable is not set');

module.exports = {
  reactStrictMode: true,
  env: {
    GW: gateway
  }
};
