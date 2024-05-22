alias authorityd=./simapp/build/simd

for arg in "$@"
do
    case $arg in
        -r|--reset)
        rm -rf .authority
        shift
        ;;
    esac
done

if ! [ -f .authority/data/priv_validator_state.json ]; then
  authorityd init validator --chain-id "authority-1" --home .authority &> /dev/null

  authorityd keys add validator --home .authority --keyring-backend test &> /dev/null
  authorityd genesis add-genesis-account validator 1000000ustake --home .authority --keyring-backend test
  OWNER=$(authorityd keys add owner --home .authority --keyring-backend test --output json | jq .address)
  authorityd genesis add-genesis-account owner 2500000uusdc --home .authority --keyring-backend test
  PENDING_OWNER=$(authorityd keys add pending-owner --home .authority --keyring-backend test --output json | jq .address)
  authorityd genesis add-genesis-account pending-owner 2500000uusdc --home .authority --keyring-backend test

  TEMP=.authority/genesis.json
  touch $TEMP && jq '.app_state.authority.owner = '$OWNER'' .authority/config/genesis.json > $TEMP && mv $TEMP .authority/config/genesis.json
  touch $TEMP && jq '.app_state.staking.params.bond_denom = "ustake"' .authority/config/genesis.json > $TEMP && mv $TEMP .authority/config/genesis.json

  authorityd genesis gentx validator 1000000ustake --chain-id "authority-1" --home .authority --keyring-backend test &> /dev/null
  authorityd genesis collect-gentxs --home .authority &> /dev/null

  sed -i '' 's/timeout_commit = "5s"/timeout_commit = "1s"/g' .authority/config/config.toml
fi

authorityd start --home .authority
