################################################################################
# Development stagenet startup
################################################################################

export PATH=$PATH:~/monero-x86_64-linux-gnu-v0.18.3.3

ALICE_BIND_PORT=38083
BETA_NODE_RPC_PORT=38081

# Let Beta use the default datadir at ~/.bitmonero/stagenet

ROOT=./xmr
WALLETS_DIR=${ROOT}/wallets

set -evx

monerod \
	--detach \
	--stagenet \
	--rpc-bind-ip 127.0.0.1 \
	--rpc-bind-port ${BETA_NODE_RPC_PORT} \
	--fixed-difficulty 1 \
	--log-level 1

monero-wallet-rpc \
	--detach \
	--stagenet \
	--rpc-bind-ip=127.0.0.1 \
	--rpc-bind-port ${ALICE_BIND_PORT} \
	--wallet-dir ${WALLETS_DIR} \
	--password "" \
	--disable-rpc-login
