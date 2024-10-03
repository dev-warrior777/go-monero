################################################################################
# Development stagenet startup
################################################################################

# Use this After seting up the wallets with xmrswap.sh. It uses the --wallet-file
# parameter (instead of --wallet-dir) so that monero-wallet-rpc knws where to
# find the wallets files previously created.

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
ALICE_WALLETFILE="${WALLETS_DIR}/alice"
BOB_WALLETFILE="${WALLETS_DIR}/bob"
EXTRAXMR_WALLETFILE="${WALLETS_DIR}/extraxmr"

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
	--wallet-file ${ALICE_WALLETFILE} \
	--password "" \
	--disable-rpc-login \
	--log-file=alice.log \
	--log-level 1

monero-wallet-rpc \
	--detach \
	--stagenet \
	--rpc-bind-ip=127.0.0.1 \
	--rpc-bind-port ${BOB_BIND_PORT} \
	--wallet-file ${BOB_WALLETFILE} \
	--password "" \
	--disable-rpc-login \
	--log-file=bob.log \
	--log-level 1

monero-wallet-rpc \
	--detach \
	--stagenet \
	--rpc-bind-ip=127.0.0.1 \
	--rpc-bind-port ${EXTRAXMR_BIND_PORT} \
	--wallet-dir ${WALLETS_DIR} \
	--password "" \
	--disable-rpc-login \
	--log-file=extraxmr.log \
	--log-level 1

# monero-wallet-rpc \
# 	--detach \
# 	--stagenet \
# 	--rpc-bind-ip=127.0.0.1 \
# 	--rpc-bind-port ${EXTRAXMR_BIND_PORT} \
# 	--wallet-file ${EXTRAXMR_WALLETFILE} \
# 	--password "" \
# 	--disable-rpc-login \
# 	--log-file=extraxmr.log \
# 	--log-level 1
