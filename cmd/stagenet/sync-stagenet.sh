################################################################################
# Development stagenet daemon startup sync
################################################################################

export PATH=$PATH:~/monero-x86_64-linux-gnu-v0.18.3.4

# CURRENT HEIGHT: height": 1,668,9xx -- AUG 18

# Let Beta use the default datadir at ~/.bitmonero/stagenet

echo "==================================================================================================="
echo "monerod listens on stdin so you can type any valid monerod command such as \"status\" into this tty"
echo "==================================================================================================="

monerod --stagenet
