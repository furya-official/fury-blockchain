#!/usr/bin/env bash

# IT IS RECOMMENDED TO RUN THE BLOCKCHAIN USING run_with_all_data_dev.sh SINCE
# THIS SETS GOVERNANCE PERIODS TO 30 seconds FOR FASTER GOVERNANCE

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

echo "Query transfer params before param change"
furyd_q ibc-transfer params

echo "Submitting param change proposal"
furyd_tx gov submit-proposal param-change demo_gov_param_change_proposal.json --from=miguel

echo "Query proposal 1"
furyd_q gov proposal 1

echo "Depositing 10000000ufury to reach minimum deposit"
furyd_tx gov deposit 1 10000000ufury --from=miguel

echo "Query proposal 1 deposits"
furyd_q gov deposits 1

echo "Voting yes for proposal"
furyd_tx gov vote 1 yes --from=miguel

echo "Query proposal 1 tally"
furyd_q gov tally 1

echo "Waiting for proposal to pass..."
while :; do
  RET=$(furyd_q gov proposal 1 2>&1)
  if [[ ($RET == *'PROPOSAL_STATUS_VOTING_PERIOD'*) ]]; then
    sleep 1
  else
    echo "A few more seconds..."
    sleep 6
    break
  fi
done

echo "Query transfer params (expected to be true and false)"
furyd_q ibc-transfer params
