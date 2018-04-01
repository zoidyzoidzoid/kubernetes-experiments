{
  global: {
    // User-defined global parameters; accessible to all component and environments, Ex:
    // replicas: 4,
  },
  components: {
    // Component-level parameters, defined initially from 'ks prototype use ...'
    // Each object below should correspond to a component in the components/ directory
    "docker-counter-deployment": {
      image: "zoidbergwill/docker-counter",
      name: "docker-counter-deployment",
      port: 80,
      replicas: 1,
    },
  },
}
