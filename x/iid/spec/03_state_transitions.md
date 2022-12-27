# State Transitions

This document describes the state transitions pertaining a [IidDocument](02_state.md#IidDocument) according to the [iid operations](https://www.w3.org/TR/iid-core/#method-operations):

1. [Create](03_state_transitions.md#Create)
2. [Resolve](03_state_transitions.md#Resolve)
3. [Update](03_state_transitions.md#Update)
4. [Deactivate](03_state_transitions.md#Deactivate)

A [IidMetadata](02_state.md#iidmetadata) lifecycle follows the lifecycle of a  [IidDocument](02_state.md#iiddocument) 

### Create

[IidDocument](02_state.md#IidDocument) are created via the rpc method [CreateIidDocument](https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/proto/iid/tx.proto#L64) that accepts a [MsgCreateIidDocument](./04_messages.md#MsgCreateIidDocument) messages as parameter.

The operation will fail if:
- the signer account has insufficient funds 
- the iid is malformed 
- a iid document with the same iid exists
- verifications 
  - the verification method is invalid (according to the verification method specifications) 
  - there is more than one verification method with the same id
  - relationships are empty
  - relationships contain unsupported values (according to the iid method specifications)
- services are invalid (according to the services specifications)
- Linked Resources are invalid (according to the Linked Resources specifications)
- Accorded Rights are invalid (according to the Accorded Rights specifications)
- Linked Entities are invalid (according to the Linked Entities specifications)

Example: 

<!-- 

cosmos-cashd tx iid create-iid \
 900d82bc-2bfe-45a7-ab22-a8d11773568e \
 --from vasp --node https://cosmos-cash.app.beta.starport.cloud:443 --chain-id cosmoscash-testnet
-->

```javascript
/* gRPC message */
CreateIidDocument(
    MsgCreateIidDocument(
        "fury:impacthub-3:900d82bc2bfe45a7ab22a8d11773568e",
        [], // controller
        [   // verifications
            {
                "relationships": ["authentication"],
                {
                    "controller": "fury:impacthub-3:900d82bc2bfe45a7ab22a8d11773568e",
                    "id": "fury:impacthub-3:900d82bc2bfe45a7ab22a8d11773568e#fury1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0",
                    "publicKeyHex": "0248a5178d7a90ec187b3c3d533a4385db905f6fcdaac5026859ca5ef7b0b1c3b5",
                    "type": "EcdsaSecp256k1VerificationKey2019"
                },
                [],
            },
        ],
        [], // services
        [], // Accorded rights
        [], // linked resources
        [], // linked entities       
        "fury1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0" // signer
    )
)

/* Resolved iid document */
{
  "IidDocument": {
    "context": [
      "https://www.w3.org/ns/iid/v1"
    ],
    "id": "fury:impacthub-3:900d82bc2bfe45a7ab22a8d11773568e",
    "controller": [],
    "verificationMethod": [
      {
        "controller": "fury:impacthub-3:900d82bc2bfe45a7ab22a8d11773568e",
        "id": "fury:impacthub-3:900d82bc2bfe45a7ab22a8d11773568e#fury1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0",
        "publicKeyHex": "0248a5178d7a90ec187b3c3d533a4385db905f6fcdaac5026859ca5ef7b0b1c3b5",
        "type": "EcdsaSecp256k1VerificationKey2019"
      }
    ],
    "service": [], 
    "AccordedRight": [],
    "LinkedResource": [],
    "LinkedEntity": [],      
    "authentication": [
      "fury:impacthub-3:900d82bc2bfe45a7ab22a8d11773568e#fury1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0"
    ],
    "assertionMethod": [],
    "keyAgreement": [],
    "capabilityInvocation": [],
    "capabilityDelegation": []
  },
  "iidMetadata": {
    "versionId": "571615b8146082deaac90fa01afc8ff88e5a71b4c9c29bcaffef2d11b39a0437",
    "created": "2021-08-23T08:24:26.972761898Z",
    "updated": "2021-08-23T08:24:26.972761898Z",
    "deactivated": false
  }
}

```

##### Implementation Source

- server: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/keeper/msg_server.go#L28
- client: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/client/cli/tx.go#L68

### Resolve

[IidDocument](02_state.md#iiddocument) are resolved via the rpc method [QueryIidDocument](https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/proto/iid/query.proto#L31) that accepts a [QueryIidDocumentRequest](./04_messages.md#QueryIidDocumentRequest) messages as parameter.


The operation will fail if:
- the iid does not exists


```javascript
/* gRPC message */
QueryIidDocument(
    QueryIidDocumentRequest(
        "fury:impacthub-3:900d82bc2bfe45a7ab22a8d11773568e"
    )
)

```

##### Implementation Source

- server: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/keeper/grpc_query.go#L28
- client: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/client/cli/query.go#L69

### Update

[IidDocument](02_state.md#iiddocument) are updated via the rpc methods:

- [UpdateIidDocument](https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/proto/iid/tx.proto#L82)
- [AddVerification](https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/proto/iid/tx.proto#L96)
- [RevokeVerification](https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/proto/iid/tx.proto#L119)
- [SetVerificationRelationships](https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/proto/iid/tx.proto#L107)
- [AddService](https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/proto/iid/tx.proto#L134)
- [DeleteService](https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/proto/iid/tx.proto#L145)
- [AddController](https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/proto/iid/tx.proto#L161)
- [DeleteController](https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/proto/iid/tx.proto#L172)
- [AddLinkedResource](https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/proto/iid/tx.proto#L134)
- [DeleteLinkedResource](https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/proto/iid/tx.proto#L145)
- [AddLinkedEntity](https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/proto/iid/tx.proto#L134)
- [DeleteLinkedEntity](https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/proto/iid/tx.proto#L145)
- [AddAccordedRight](https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/proto/iid/tx.proto#L134)
- [DeleteAccordedRight](https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/proto/iid/tx.proto#L145)


All the operations will fail if:

- the signer account has insufficient funds
- the signer account address doesn't match the verification method listed in the `Authorization` verification relationships
- the target iid does not exists

The following sections provide specific details for each method invocation.

#### UpdateIidDocument 

The  `UpdateIidDocument` method is used to **overwrite** the  [IidDocument](02_state.md#iiddocument) controllers. It accepts a [MsgUpdateIidDocument](./04_messages.md#MsgUpdateIidDocument) as a parameter.

The operation will fail if:

- any of the provided controllers is not a valid iid

```javascript
/* gRPC message */
UpdateIidDocument(
    MsgUpdateIidDocument(
        "fury:impacthub-3:900d82bc2bfe45a7ab22a8d11773568e",
        ["fury:key:fury1sl48sj2jjed7enrv3lzzplr9wc2f5js5tzjph8"],
        "fury1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0"
    )
)
```

##### Implementation Source

- server: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/keeper/msg_server.go#L65
- client: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/client/cli/tx.go#L277

#### AddVerification

The `AddVerification` method is used to add new [verification methods](https://w3c.github.io/iid-core/#verification-methods) and [verification relationships](https://w3c.github.io/iid-core/#verification-relationships) to a [IidDocument](02_state.md#IidDocument). It accepts a [MsgAddVerification](./04_messages.md#MsgAddVerification) as a parameter.

The operation will fail if:

- the verification method is invalid (according to the verification method specifications) 
- the verification method id already exists for the iid document
- the verification relationships are empty
- the verification relationships contain unsupported values (according to the iid method specification)

```javascript
/* gRPC message */
AddVerification(
    MsgAddVerification(
        "fury:impacthub-3:900d82bc2bfe45a7ab22a8d11773568e",
        {
            "relationships": ["authentication"],
            {
                "controller": "fury:impacthub-3:900d82bc2bfe45a7ab22a8d11773568e",
                "id": "fury:impacthub-3:900d82bc2bfe45a7ab22a8d11773568e#cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2",
                "publicKeyHex": "03786095e15eb228f4e15692eda6e0607a313cc081ad54d69aadd15d515e304590",
                "type": "EcdsaSecp256k1VerificationKey2019"
            },
            [],
        },
        "fury1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0" // signer
    )
)

```

##### Implementation Source

- server: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/keeper/msg_server.go#L98
- client: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/client/cli/tx.go#L131

#### RevokeVerification

The `RevokeVerification` method is used to remove existing [verification methods](https://w3c.github.io/iid-core/#verification-methods) and [verification relationships](https://w3c.github.io/iid-core/#verification-relationships) from a [IidDocument](02_state.md#IidDocument). It accepts a [MsgRevokeVerification](./04_messages.md#MsgRevokeVerification) as a parameter.

The operation will fail if:

- the verification method id is not found


```javascript
/* gRPC message */
RevokeVerification(
    MsgRevokeVerification(
        "fury:impacthub-3:900d82bc2bfe45a7ab22a8d11773568e",
        "fury:impacthub-3:900d82bc2bfe45a7ab22a8d11773568e#cosmos1lvl2s8x4pta5f96appxrwn3mypsvumukvk7ck2",
        "fury1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0" // signer
    )
)

```

##### Implementation source

- server: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/keeper/msg_server.go#L304
- client: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/client/cli/tx.go#L233


#### SetVerificationRelationships


The `SetVerificationRelationships` method is used to **overwrite** existing [verification relationships](https://w3c.github.io/iid-core/#verification-relationships) for a [verification methods](https://w3c.github.io/iid-core/#verification-methods) in a [IidDocument](02_state.md#IidDocument). It accepts a [MsgSetVerificationRelationships](./04_messages.md#MsgSetVerificationRelationships) as a parameter.

The operation will fail if:

- the verification method id is not found for the target iid document
- the verification relationships are empty 
- the verification relationships contain unsupported values (according to the iid method specification)


```javascript
/* gRPC message */
SetVerificationRelationships(
    MsgSetVerificationRelationships(
        "fury:impacthub-3:900d82bc2bfe45a7ab22a8d11773568e",
        "fury:impacthub-3:900d82bc2bfe45a7ab22a8d11773568e#fury1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0",
        ["authentication", "capabilityInvocation"]
        "fury1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0" // signer
    )
)

```

##### Implementation source

- server: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/keeper/msg_server.go#L348
- client: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/client/cli/tx.go#L390


#### AddService


The `AddService` method is used to add a [service](https://w3c.github.io/iid-core/#services) in a [IidDocument](02_state.md#IidDocument). It accepts a [MsgAddService](./04_messages.md#MsgAddService) as a parameter.

The operation will fail if:

- a service with the same id already present in the iid document
- the service definition is invalid (according to the iid services specification)

```javascript
/* gRPC message */
AddService(
    MsgAddService(
        "fury:impacthub-3:900d82bc2bfe45a7ab22a8d11773568e",
        {
            "agent:xyz",
            "iidCommMessaging",
            "https://agent.xyz/1234",
        }
        "fury1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0" // signer
    )
)

```

##### Implementation source

- server: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/keeper/msg_server.go#L116
- client: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/client/cli/tx.go#L193

#### DeleteService


The `DeleteService` method is used to remove a [service](https://w3c.github.io/iid-core/#services) from a [IidDocument](02_state.md#IidDocument). It accepts a [MsgDeleteService](./04_messages.md#MsgDeleteService) as a parameter.

The operation will fail if:

- the service id does not match any service in the iid document

```javascript
/* gRPC message */
DeleteService(
    MsgDeleteService(
        "fury:impacthub-3:900d82bc2bfe45a7ab22a8d11773568e",
        "agent:xyz",
        "fury1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0" // signer
    )
)

```

##### Implementation source

- server: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/keeper/msg_server.go#L323
- client: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/client/cli/tx.go#L274

#### AddLinkedResource


The `AddLinkedResource` method is used to add a [LinkedResource](#) in a [IidDocument](02_state.md#IidDocument). It accepts a [MsgAddLinkedResource](./04_messages.md#MsgAddLinkedResource) as a parameter.

The operation will fail if:

- a LinkedResource with the same id already present in the iid document
- the LinkedResource definition is invalid (according to the LinkedResource specification)

```javascript
/* gRPC message */
AddLinkedResource(
    MsgAddLinkedResource(
        "fury:impacthub-3:900d82bc2bfe45a7ab22a8d11773568e",
            {
            "fury:entity:abc123#****",
            "entityProfile",
            "Test Clean Cooking Collection",
            "application/json",
            "#cellnode-pandora/public/****", 
            "****", 
            "false",
            "right",
            },
        "fury1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0" // signer
    )
)

```

##### Implementation source

- server: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/keeper/msg_server.go#L134
- client: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/client/cli/tx.go#L445

#### DeleteLinkedResource


The `DeleteLinkedResource` method is used to remove a [LinkedResource](#) from a [IidDocument](02_state.md#iiddocument). It accepts a [MsgDeleteService](./04_messages.md#MsgDeleteService) as a parameter.

The operation will fail if:

- the LinkedResource id does not match any LinkedResource in the iid document

```javascript
/* gRPC message */
DeleteLinkedResource(
    MsgDeleteLinkedResource(
        "fury:impacthub-3:900d82bc2bfe45a7ab22a8d11773568e",
        "resource id",
        "fury1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0" // signer
    )
)

```

##### Implementation source

- server: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/keeper/msg_server.go#L152
- client: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/client/cli/tx.go#L490

#### AddAccordedRight


The `AddAccordedRight` method is used to add a [AccordedRight](#) in a [IidDocument](02_state.md#IidDocument). It accepts a [MsgAddAccordedRight](./04_messages.md#MsgAddAccordedRight) as a parameter.

The operation will fail if:

- a AccordedRight with the same id already present in the iid document
- the AccordedRight definition is invalid (according to the LinkedResource specification)

```javascript
/* gRPC message */
AddAccordedRight(
    MsgAddAccordedRight(
        "fury:impacthub-3:900d82bc2bfe45a7ab22a8d11773568e",
            {
                "fury:entity:abc123#mintNFT",
                "mint",
                "cw721",
                "msgMintNFT",
                "#fury"
            },
        "fury1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0" // signer
    )
)

```

##### Implementation source

- server: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/keeper/msg_server.go#L218
- client: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/client/cli/tx.go#L527

#### DeleteAccordedRight


The `DeleteAccordedRight` method is used to remove a [AccordedRight](#) from a [IidDocument](02_state.md#iiddocument). It accepts a [MsgDeleteAccordedRight](./04_messages.md#MsgDeleteAccordedRight) as a parameter.

The operation will fail if:

- the AccordedRight id does not match any AccordedRight in the iid document

```javascript
/* gRPC message */
DeleteAccordedRight(
    MsgDeleteAccordedRight(
        "fury:impacthub-3:900d82bc2bfe45a7ab22a8d11773568e",
        "right id",
        "fury1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0" // signer
    )
)

```

##### Implementation source

- server: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/keeper/msg_server.go#L236
- client: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/client/cli/tx.go#L569


#### AddLinkedEntity


The `AddLinkedEntity` method is used to add a [LinkedEntity](#) in a [IidDocument](02_state.md#IidDocument). It accepts a [MsgAddLinkedEntity](./04_messages.md#MsgAddLinkedEntity) as a parameter.

The operation will fail if:

- a LinkedEntity with the same id already present in the iid document
- the LinkedEntity definition is invalid (according to the LinkedResource specification)

```javascript
/* gRPC message */
AddLinkedEntity(
    MsgAddLinkedEntity(
        "fury:impacthub-3:900d82bc2bfe45a7ab22a8d11773568e",
            {
                "fury:entity:abc123#123",
                "relationship",
            },
        "fury1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0" // signer
    )
)

```

##### Implementation source

- server: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/keeper/msg_server.go#L176
- client: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/client/cli/tx.go#L726

#### DeleteLinkedEntity


The `DeleteLinkedEntity` method is used to remove a [LinkedEntity](#) from a [IidDocument](02_state.md#iiddocument). It accepts a [MsgDeleteLinkedEntity](./04_messages.md#MsgDeleteLinkedEntitiy) as a parameter.

The operation will fail if:

- the LinkedEntity id does not match any LinkedEntity in the iid document

```javascript
/* gRPC message */
DeleteLinkedEntity(
    MsgDeleteLinkedEntity(
        "fury:impacthub-3:900d82bc2bfe45a7ab22a8d11773568e",
        "enitiy id",
        "fury1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0" // signer
    )
)

```

##### Implementation source

- server: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/keeper/msg_server.go#L194
- client: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/client/cli/tx.go#L765

#### AddController


The `AddController` method is used to add a [Controller](#) in a [IidDocument](02_state.md#IidDocument). It accepts a [MsgAddController](./04_messages.md#MsgAddController) as a parameter.

The operation will fail if:

- a LinkedEntity with the same id already present in the iid document
- the LinkedEntity definition is invalid (according to the LinkedResource specification)

```javascript
/* gRPC message */
AddController(
    MsgAddController(
        "fury:impacthub-3:900d82bc2bfe45a7ab22a8d11773568e",
            {
                "fury:entity:abc123#123",
            },
        "fury1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0" // signer
    )
)

```

##### Implementation source

- server: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/keeper/msg_server.go#L367
- client: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/client/cli/tx.go#L312

#### DeleteDeleteController


The `DeleteController` method is used to remove a [Controller](#) from a [IidDocument](02_state.md#iiddocument). It accepts a [MsgDeleteController](./04_messages.md#MsgDeleteController) as a parameter.

The operation will fail if:

- the Controller id does not match any controller in the iid document

```javascript
/* gRPC message */
DeleteController(
    MsgDeleteController(
        "fury:impacthub-3:900d82bc2bfe45a7ab22a8d11773568e",
        "Controller id",
        "fury1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0" // signer
    )
)

```

##### Implementation source

- server: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/keeper/msg_server.go#L385
- client: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/client/cli/tx.go#L351

#### AddController


The `AddIidContext` method is used to add a [IidContext](#) in a [IidDocument](02_state.md#IidDocument). It accepts a [MsgAddIidContext](./04_messages.md#MsgAddIidContext) as a parameter.

The operation will fail if:

- a IidContext with the same key already present in the iid document
- the IidContext definition is invalid (according to the IidContext specification)

```javascript
/* gRPC message */
AddIidContext(
    MsgAddIidContext(
        "fury:impacthub-3:900d82bc2bfe45a7ab22a8d11773568e",
            {
                "fury",
                "context string"
            },
        "fury1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0" // signer
    )
)

```

##### Implementation source

- server: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/keeper/msg_server.go#L267
- client: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/client/cli/tx.go#L606

#### DeleteDeleteController


The `DeleteIidContext` method is used to remove a [IidContext](#) from a [IidDocument](02_state.md#iiddocument). It accepts a [MsgDeleteIidContext](./04_messages.md#MsgDeleteIidContext) as a parameter.

The operation will fail if:

- the IID context key does not match any context in the iid document

```javascript
/* gRPC message */
DeleteIidContext(
    MsgDeleteIidContext(
        "fury:impacthub-3:900d82bc2bfe45a7ab22a8d11773568e",
        "Context Key",
        "fury1x5hrv0hngmg8gls5cft7nphqs83njj25pwxpt0" // signer
    )
)

```

##### Implementation source

- server: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/keeper/msg_server.go#L279
- client: https://github.com/furya-official/fury-blockchain/blob/devel/iid-module/x/iid/client/cli/tx.go#L647
