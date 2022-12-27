#!/usr/bin/env bash

PASSWORD="12345678"

furyd init local --chain-id pandora-4

yes 'y' | furyd keys delete miguel --force
yes $PASSWORD | furyd keys add miguel

# Note: important to add 'miguel' as a genesis-account since this is the chain's validator
yes $PASSWORD | furyd add-genesis-account "$(furyd keys show miguel -a)" 1000000000000ufury,1000000000000res,1000000000000rez,1000000000000uxgbp

# Add pubkey-based genesis accounts
MIGUEL_ADDR="fury107pmtx9wyndup8f9lgj6d7dnfq5kuf3sapg0vx"    # address from did:fury:4XJLBfGtWSGKSz4BeRxdun's pubkey
yes $PASSWORD | furyd add-genesis-account "$MIGUEL_ADDR" 1000000000000ufury,1000000000000res,1000000000000rez

# Add fury did
FURY_DID="did:fury:U4tSpzzv91HHqWW1YmFkHJ"
FROM="\"fury_did\": \"\""
TO="\"fury_did\": \"$FURY_DID\""
sed -i "s/$FROM/$TO/" "$HOME"/.furyd/config/genesis.json

# Set staking token (both bond_denom and mint_denom)
STAKING_TOKEN="ufury"
FROM="\"bond_denom\": \"stake\""
TO="\"bond_denom\": \"$STAKING_TOKEN\""
sed -i "s/$FROM/$TO/" "$HOME"/.furyd/config/genesis.json
FROM="\"mint_denom\": \"stake\""
TO="\"mint_denom\": \"$STAKING_TOKEN\""
sed -i "s/$FROM/$TO/" "$HOME"/.furyd/config/genesis.json

# Set fee token (both for gov min deposit and crisis constant fee)
FEE_TOKEN="ufury"
FROM="\"stake\""
TO="\"$FEE_TOKEN\""
sed -i "s/$FROM/$TO/" "$HOME"/.furyd/config/genesis.json

# Set reserved bond tokens
RESERVED_BOND_TOKENS=""  # example: " \"abc\", \"def\", \"ghi\" "
FROM="\"reserved_bond_tokens\": \[\]"
TO="\"reserved_bond_tokens\": \[$RESERVED_BOND_TOKENS\]"
sed -i "s/$FROM/$TO/" "$HOME"/.furyd/config/genesis.json

# Set min-gas-prices (using fee token)
FROM="minimum-gas-prices = \"\""
TO="minimum-gas-prices = \"0.025$FEE_TOKEN\""
sed -i "s/$FROM/$TO/" "$HOME"/.furyd/config/app.toml

# TODO: config missing from new version (REF: https://github.com/cosmos/cosmos-sdk/issues/8529)
#furyd config chain-id pandora-4
#furyd config output json
#furyd config indent true
#furyd config trust-node true

furyd gentx miguel 1000000ufury --chain-id pandora-4

furyd collect-gentxs
furyd validate-genesis

# Enable REST API (assumed to be at line 104 of app.toml)
FROM="enable = false"
TO="enable = true"
sed -i "104s/$FROM/$TO/" "$HOME"/.furyd/config/app.toml

# Enable Swagger docs (assumed to be at line 107 of app.toml)
FROM="swagger = false"
TO="swagger = true"
sed -i "107s/$FROM/$TO/" "$HOME"/.furyd/config/app.toml

# Uncomment the below to broadcast node RPC endpoint
#FROM="laddr = \"tcp:\/\/127.0.0.1:26657\""
#TO="laddr = \"tcp:\/\/0.0.0.0:26657\""
#sed -i "s/$FROM/$TO/" "$HOME"/.furyd/config/config.toml

# Uncomment the below to set timeouts to 1s for shorter block times
#sed -i 's/timeout_commit = "5s"/timeout_commit = "1s"/g' "$HOME"/.furyd/config/config.toml
#sed -i 's/timeout_propose = "3s"/timeout_propose = "1s"/g' "$HOME"/.furyd/config/config.toml

furyd start --pruning "nothing" --log_level "trace" --trace
