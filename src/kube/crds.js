const conditionsSchema = {
  description: 'Conditions is an array of the current observed status of an API Resource.',
  type: 'array',
  items: {
    description: 'Condition contains details for one aspect of the current state of this API Resource.',
    type: 'object',
    required: [
      'type',
      'status',
    ],
    properties: {
      lastTransitionTime: {
        description: 'lastTransitionTime is the last time the condition transitioned from one status to another. This should be when the underlying condition changed.  If that is not known, then using the time when the API field changed is acceptable.',
        type: 'string',
        format: 'date-time',
      },
      message: {
        description: 'message is a human readable message indicating details about the transition. This may be an empty string.',
        type: 'string'
      },
      observedGeneration: {
        description: 'observedGeneration represents the .metadata.generation that the condition was set based upon. For instance, if .metadata.generation is currently 12, but the .status.conditions[x].observedGeneration is 9, the condition is out of date with respect to the current state of the instance.',
        type: 'integer',
        format: 'int64'
      },
      reason: {
        description: 'reason contains a programmatic identifier indicating the reason for the condition\'s last transition. Producers of specific condition types may define expected values and meanings for this field, and whether the values are considered a guaranteed API. The value should be a CamelCase string. This field may not be empty.',
        type: 'string'
      },
      status: {
        description: 'status of the condition, one of True, False, Unknown.',
        type: 'string',
        enum: ['True', 'False', 'Unknown']
      },
      type: {
        description: 'type of condition in CamelCase or in foo.example.com/CamelCase.',
        type: 'string'
      },
    },
  },
};

const battlesCrd = {
  metadata: {
    name: 'battles.spirits.dev'
  },
  spec: {
    group: 'spirits.dev',
    names: {
      categories: ['spiritsworld'],
      kind: 'Battle',
      listKind: 'BattleList',
      plural: 'battles',
      singular: 'battle',
    },
    scope: 'Namespaced',
    versions: [
      {
        name: 'v1alpha1',
        served: true,
        storage: true,
        subresources: {
          status: {},
        },
        additionalPrinterColumns: [
          {
            jsonPath: '.status.phase',
            name: 'Phase',
            type: 'string',
          }
        ],
        schema: {
          openAPIV3Schema: {
            description: 'Battle is a skirmish amongst spirits.',
            type: 'object',
            required: [
              'spec',
            ],
            properties: {
              apiVersion: {
                description: 'APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources',
                type: 'string',
              },
              kind: {
                description: 'Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds',
                type: 'string',
              },
              metadata: {
                type: 'object',
              },
              spec: {
                description: 'Spec for configuring the Battle.',
                type: 'object',
                required: [
                  'spirits',
                ],
                properties: {
                  spirits: {
                    description: 'The spirits involved in this Battle.',
                    type: 'array',
                    items: {
                      description: 'A name of a spirit object involved in this Battle.',
                      type: 'string',
                    },
                  },
                },
              },
              status: {
                description: 'Status of the Battle.',
                type: 'object',
                properties: {
                  conditions: conditionsSchema,
                  phase: {
                    description: 'Human-readable description of the current status of the Battle.',
                    type: 'string',
                  },
                },
              },
            },
          },
        },
      },
    ],
  },
};

const spiritsCrd = {
  metadata: {
    name: 'spirits.spirits.dev'
  },
  spec: {
    group: 'spirits.dev',
    names: {
      categories: ['spiritsworld'],
      kind: 'Spirit',
      listKind: 'SpiritList',
      plural: 'spirits',
      singular: 'spirit',
    },
    scope: 'Namespaced',
    versions: [
      {
        name: 'v1alpha1',
        served: true,
        storage: true,
        subresources: {
          status: {},
        },
        additionalPrinterColumns: [
          {
            jsonPath: '.spec.stats.health',
            name: 'Health',
            type: 'string',
          },
        ],
        schema: {
          openAPIV3Schema: {
            description: 'Spirit is a single actor in a Battle.',
            type: 'object',
            required: [
              'spec',
            ],
            properties: {
              apiVersion: {
                description: 'APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources',
                type: 'string',
              },
              kind: {
                description: 'Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds',
                type: 'string',
              },
              metadata: {
                type: 'object',
              },
              spec: {
                description: 'Spec for configuring the Spirit.',
                type: 'object',
                required: [
                  'stats',
                ],
                properties: {
                  stats: {
                    description: 'Quantitative properties of the Spirit. These are utilized and manipulated throughout the course of a Battle.',
                    type: 'object',
                    required: [
                      'health',
                    ],
                    properties: {
                      health: {
                        description: 'A quantitative representation of the energy of the Spirit. When this drops to 0, the Spirit is no longer able to participate in a Battle.',
                        type: 'integer',
                        format: 'int64',
                        minimum: 1,
                      },
                      power: {
                        description: 'A quantitative representation of the might of the Spirit.',
                        type: 'integer',
                        format: 'int64',
                        minimum: 0,
                        default: 0,
                      },
                      armor: {
                        description: 'A quantitative representation of the defense of the Spirit.',
                        type: 'integer',
                        format: 'int64',
                        minimum: 0,
                        default: 0,
                      },
                      agility: {
                        description: 'A quantitative representation of the speed of the Spirit.',
                        type: 'integer',
                        format: 'int64',
                        minimum: 0,
                        default: 0,
                      },
                    },
                  },
                },
              },
              status: {
                description: 'Status of the Spirit.',
                type: 'object',
                properties: {
                  conditions: conditionsSchema,
                },
              },
            },
          },
        },
      },
    ],
  },
};

const crds = [
  spiritsCrd,
  battlesCrd,
];

const createOrUpdateCRD = (client, crd) => {
  return client.readCustomResourceDefinition(crd.metadata.name).then((rsp) => {
    // CRD exitsts - update it.
    rsp.body.spec = crd.spec;
    return client.replaceCustomResourceDefinition(crd.metadata.name, rsp.body);
  }).catch((error) => {
    if (error.statusCode === 404) {
      // CRD does not exist - create it.
      return client.createCustomResourceDefinition(crd);
    }

    // Unexpected error - propagate.
    return Promise.reject(error);
  });
}

module.exports = {
  upsert: (client) => {
    return Promise.all(crds.map((crd) => createOrUpdateCRD(client, crd)));
  },
};
