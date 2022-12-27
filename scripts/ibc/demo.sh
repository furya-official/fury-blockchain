#!/usr/bin/env bash

# This script is not meant to be run automatically, but one line at a time
#
# Also, in order for this script to work with Starport, cmd/furyd needs to be
# renamed to cmd/fury-blockchaind. Not sure how we can avoid having to do this.
#
# The steps should be run from the project folder (fury-blockchain) not ibc/

HOME_1="./data_1"
HOME_2="./data_2"
CONFIG_1="./scripts/ibc/config_1.yml"
CONFIG_2="./scripts/ibc/config_2.yml"
GAS_PRICES_1="0.025ufury"
GAS_PRICES_2="0.025uatom"
CHAIN_ID_1="pandora-4.1"
CHAIN_ID_2="pandora-4.2"
RPC_1_HTTP="http://localhost:26659"
RPC_2_HTTP="http://localhost:26661"
RPC_1_TCP="tcp://localhost:26659"
RPC_2_TCP="tcp://localhost:26661"
PREFIX="fury"
FAUCET_1="http://localhost:4500"
FAUCET_2="http://localhost:4502"

furyd1_tx() {
  # Helper function to broadcast a transaction and supply the necessary args

  # Get module ($1) and specific tx ($1), which forms the tx command
  cmd="$1 $2"
  shift
  shift

  # Broadcast the transaction
  furyd1 tx $cmd \
    --gas-prices="$GAS_PRICES_1" \
    --chain-id="$CHAIN_ID_1" \
    --node="$RPC_1_TCP" \
    --home "$HOME_1" \
    --keyring-backend="test" \
    --keyring-dir "$HOME_1" \
    --broadcast-mode block \
    -y \
    "$@" | jq .
    # The $@ adds any extra arguments to the end
}

furyd1_q() {
  furyd1 q "$@" --node="$RPC_1_TCP" --output=json | jq .
}

furyd2_tx() {
  # Helper function to broadcast a transaction and supply the necessary args

  # Get module ($1) and specific tx ($1), which forms the tx command
  cmd="$1 $2"
  shift
  shift

  # Broadcast the transaction
  furyd2 tx $cmd \
    --gas-prices="$GAS_PRICES_2" \
    --chain-id="$CHAIN_ID_2" \
    --node="$RPC_2_TCP" \
    --home "$HOME_2" \
    --keyring-backend="test" \
    --keyring-dir "$HOME_2" \
    --broadcast-mode block \
    -y \
    "$@" | jq .
    # The $@ adds any extra arguments to the end
}

furyd2_q() {
  furyd2 q "$@" --node="$RPC_2_TCP" --output=json | jq .
}

# Start up two chains in separate terminals, one at a time
starport serve --config "$CONFIG_1" --home "$HOME_1" --reset-once
starport serve --config "$CONFIG_2" --home "$HOME_2" --reset-once

# Check that keys were created
furyd1 keys list --keyring-backend "test" --keyring-dir "$HOME_1"
furyd2 keys list --keyring-backend "test" --keyring-dir "$HOME_2"

# Configure relayer
rm ~/.starport/relayer/config.yml
starport relayer configure \
  --source-rpc="$RPC_1_HTTP" \
  --source-gasprice="$GAS_PRICES_1" \
  --source-prefix="$PREFIX" \
  --source-faucet="$FAUCET_1" \
  --target-rpc="$RPC_2_HTTP" \
  --target-gasprice="$GAS_PRICES_2" \
  --target-prefix="$PREFIX" \
  --target-faucet="$FAUCET_2"

# Send tokens to relayer on fury
# [[update address if need be]]
# [[can be skipped if faucets working]]
furyd1_tx bank send alice fury1q65z2ky63ks52mcztgjr7dm62lwpm46gkg5422 1000000ufury

# Send tokens to relayer on cosmos
# [[update address if need be]]
# [[can be skipped if faucets working]]
furyd2_tx bank send charlie fury1q65z2ky63ks52mcztgjr7dm62lwpm46gkg5422 1000000uatom

# Connect the two chains
starport relayer connect

# Send tokens from pandora-4.1 to pandora-4.2
# [[update channel ID if need be]]
# [[receiver address is arbitrary]]
furyd1_tx ibc-transfer transfer transfer channel-0 fury16qeg5rzwhamydtlarc9v6e3ld46x9lxv5tkh4u 123ufury --from=alice

# Send tokens from pandora-4.2 to pandora-4.1
# [[update channel ID if need be]]
# [[receiver address is arbitrary]]
furyd2_tx ibc-transfer transfer transfer channel-0 fury1fe3v2dwp6mr25hflwljdddp8vh3cseymp3kmpv 123uatom --from=charlie

# Query balance on pandora-4.1
furyd1_q bank balances fury1fe3v2dwp6mr25hflwljdddp8vh3cseymp3kmpv

# Query balance on pandora-4.2
furyd2_q bank balances fury16qeg5rzwhamydtlarc9v6e3ld46x9lxv5tkh4u

# Query denom traces on pandora-4.1
furyd1_q ibc-transfer denom-traces

# Query denom traces on pandora-4.2
furyd2_q ibc-transfer denom-traces

# Clean up when you're done!
rm -r ./data_1
rm -r ./data_2
rm ~/go/bin/furyd1
rm ~/go/bin/furyd2
