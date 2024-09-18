# saturn-bot-protocol

This repository defines the protocol used by plugins of [saturn-bot](https://github.com/wndhydrnt/saturn-bot).

The protocol is defined in gRPC.

Plugin libraries that implement the protocol can use the integration test suite to verify the implementation. The directory [features](./features/) contains the test suite. Every [release](https://github.com/wndhydrnt/saturn-bot-protocol/releases) distributes binaries of the integration test suite. Download the binary and execute it:

```shell
integration-test -path <path to plugin implementation>
```
