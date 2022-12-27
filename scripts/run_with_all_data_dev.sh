#!/usr/bin/env bash

PASSWORD="12345678"

furyd init local --chain-id pandora-4

yes 'y' | furyd keys delete miguel --force
yes 'y' | furyd keys delete francesco --force
yes 'y' | furyd keys delete shaun --force
yes 'y' | furyd keys delete fee --force
yes 'y' | furyd keys delete fee2 --force
yes 'y' | furyd keys delete fee3 --force
yes 'y' | furyd keys delete fee4 --force
yes 'y' | furyd keys delete fee5 --force
yes 'y' | furyd keys delete reserveOut --force

yes $PASSWORD | furyd keys add miguel
yes $PASSWORD | furyd keys add francesco
yes $PASSWORD | furyd keys add shaun
yes $PASSWORD | furyd keys add fee
yes $PASSWORD | furyd keys add fee2
yes $PASSWORD | furyd keys add fee3
yes $PASSWORD | furyd keys add fee4
yes $PASSWORD | furyd keys add fee5
yes $PASSWORD | furyd keys add reserveOut

# Note: important to add 'miguel' as a genesis-account since this is the chain's validator
yes $PASSWORD | furyd add-genesis-account "$(furyd keys show miguel -a)" 1000000000000ufury,1000000000000res,1000000000000rez,1000000000000uxgbp
yes $PASSWORD | furyd add-genesis-account "$(furyd keys show francesco -a)" 1000000000000ufury,1000000000000res,1000000000000rez
yes $PASSWORD | furyd add-genesis-account "$(furyd keys show shaun -a)" 1000000000000ufury,1000000000000res,1000000000000rez

# Add pubkey-based genesis accounts
MIGUEL_ADDR="fury107pmtx9wyndup8f9lgj6d7dnfq5kuf3sapg0vx"    # address from did:fury:4XJLBfGtWSGKSz4BeRxdun's pubkey
FRANCESCO_ADDR="fury1cpa6w2wnqyxpxm4rryfjwjnx75kn4xt372dp3y" # address from did:fury:UKzkhVSHc3qEFva5EY2XHt's pubkey
SHAUN_ADDR="fury1d5u5ta7np7vefxa7ttpuy5aurg7q5regm0t2un"     # address from did:fury:U4tSpzzv91HHqWW1YmFkHJ's pubkey
yes $PASSWORD | furyd add-genesis-account "$MIGUEL_ADDR" 1000000000000ufury,1000000000000res,1000000000000rez
yes $PASSWORD | furyd add-genesis-account "$FRANCESCO_ADDR" 1000000000000ufury,1000000000000res,1000000000000rez
yes $PASSWORD | furyd add-genesis-account "$SHAUN_ADDR" 1000000000000ufury,1000000000000res,1000000000000rez

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

# Set max deposit period to 30s for faster governance
MAX_DEPOSIT_PERIOD="30s"  # example: "172800s"
FROM="\"max_deposit_period\": \"172800s\""
TO="\"max_deposit_period\": \"$MAX_DEPOSIT_PERIOD\""
sed -i "s/$FROM/$TO/" "$HOME"/.furyd/config/genesis.json

# Set voting period to 30s for faster governance
MAX_VOTING_PERIOD="30s"  # example: "172800s"
FROM="\"voting_period\": \"172800s\""
TO="\"voting_period\": \"$MAX_VOTING_PERIOD\""
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

FROM="laddr = \"tcp:\/\/127.0.0.1:26657\""
TO="laddr = \"tcp:\/\/0.0.0.0:26657\""
sed -i "s/$FROM/$TO/" "$HOME"/.furyd/config/config.toml

# Uncomment the below to set timeouts to 1s for shorter block times
#sed -i 's/timeout_commit = "5s"/timeout_commit = "1s"/g' "$HOME"/.furyd/config/config.toml
#sed -i 's/timeout_propose = "3s"/timeout_propose = "1s"/g' "$HOME"/.furyd/config/config.toml

furyd start --pruning "nothing" --inv-check-period 1
