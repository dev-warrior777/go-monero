################################################################################
# Development stagenet startup
################################################################################

# Use this plus the golang gen_* tools to generate wallets in WALLETS_DIR

# change as needed
export PATH=$PATH:~/monero-x86_64-linux-gnu-v0.18.3.4

NODE_RPC_PORT=38081

ALICE_BIND_PORT=38284
BOB_BIND_PORT=38184
EXTRAXMR_BIND_PORT=38484

set -evx

# change as needed
# ROOT=/home/dev/dextest/xmr
ROOT=/home/dev/dextest/xmr-testnet
# datadir
WALLETS_DIR=${ROOT}/wallets
if [ ! -d $WALLETS_DIR ]
then
	mkdir -p $WALLETS_DIR
fi

monerod \
	--detach \
	--stagenet \
	--rpc-bind-ip 127.0.0.1 \
	--rpc-bind-port ${NODE_RPC_PORT} \
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

monero-wallet-rpc \
	--detach \
	--stagenet \
	--rpc-bind-ip=127.0.0.1 \
	--rpc-bind-port ${BOB_BIND_PORT} \
	--wallet-dir ${WALLETS_DIR} \
	--password "" \
	--disable-rpc-login

monero-wallet-rpc \
	--detach \
	--stagenet \
	--rpc-bind-ip=127.0.0.1 \
	--rpc-bind-port ${EXTRAXMR_BIND_PORT} \
	--wallet-dir ${WALLETS_DIR} \
	--password "" \
	--disable-rpc-login
