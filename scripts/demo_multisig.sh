#!/usr/bin/env bash

wait() {
  echo "Waiting for chain to start..."
  while :; do
    RET=$(furyd status 2>&1)
    if [[ ($RET == Error*) || ($RET == *'"latest_block_height":"0"'*) ]]; then
      sleep 1
    else
      echo "A few more seconds..."
      sleep 6
      break
    fi
  done
}

RET=$(furyd status 2>&1)
if [[ ($RET == Error*) || ($RET == *'"latest_block_height":"0"'*) ]]; then
  wait
fi

GAS_PRICES="0.025ufury"
PASSWORD="12345678"
CHAIN_ID="pandora-4"

furyd_tx() {
  # Helper function to broadcast a transaction and supply the necessary args

  # Get module ($1) and specific tx ($1), which forms the tx command
  cmd="$1 $2"
  shift
  shift

  # Broadcast the transaction
  furyd tx $cmd \
    --gas-prices="$GAS_PRICES" \
    --chain-id="$CHAIN_ID" \
    --broadcast-mode block \
    -y \
    "$@" | jq .
    # The $@ adds any extra arguments to the end
}

furyd_q() {
  furyd q "$@" --output=json | jq .
}

# Create multi-sig key using Miguel, Francesco, and Shaun, but with a 2 key
# threshold, meaning that only two people have to sign any transactions.
furyd keys delete mfs-multisig -y  # remove just in case it already exists
furyd keys add mfs-multisig \
  --multisig miguel,francesco,shaun \
  --multisig-threshold 2

MULTISIG_ADDR=$(yes $PASSWORD | furyd keys show mfs-multisig -a)
MIGUEL_ADDR=$(yes $PASSWORD | furyd keys show miguel -a)
SHAUN_ADDR=$(yes $PASSWORD | furyd keys show shaun -a)

# Send tokens to the multi-sig address
furyd_tx bank send miguel "$MULTISIG_ADDR" 1000000ufury
# Check balance of the multi-sig address
furyd_q bank balances "$MULTISIG_ADDR"

# Send some tokens back to Miguel
#
# ...part 1: generate transaction
furyd tx bank send "$MULTISIG_ADDR" "$MIGUEL_ADDR" 123ufury \
  --generate-only \
  --gas-prices="$GAS_PRICES" \
  --chain-id="$CHAIN_ID" \
  > tx.json
#
# ...part 2: miguel signs tx.json and produces tx-signed-miguel.json
furyd tx sign tx.json \
  --from "$MIGUEL_ADDR" \
  --multisig "$MULTISIG_ADDR" \
  --sign-mode amino-json \
  --chain-id="$CHAIN_ID" \
  >> tx-signed-miguel.json
#
# ...part 2: shaun signs tx.json and produces tx-signed-shaun.json
furyd tx sign tx.json \
  --from "$SHAUN_ADDR" \
  --multisig "$MULTISIG_ADDR" \
  --sign-mode amino-json \
  --chain-id="$CHAIN_ID" \
  >> tx-signed-shaun.json
#
# ...part 3: join signatures
furyd tx multisign tx.json mfs-multisig tx-signed-miguel.json tx-signed-shaun.json \
  --from mfs-multisig \
  --chain-id="$CHAIN_ID" \
  > tx_ms.json
#
# ...part 4: broadcast
furyd_tx broadcast ./tx_ms.json

# Check Miguel's balance
furyd_q bank balances "$MIGUEL_ADDR"
# Check balance of the multi-sig address
furyd_q bank balances "$MULTISIG_ADDR"

# Cleanup
furyd keys delete mfs-multisig -y
rm tx.json tx-signed-miguel.json tx-signed-shaun.json tx_msg.json