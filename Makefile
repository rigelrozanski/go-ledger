all: get_ledgerjs install

install:
	go install

get_ledgerjs:
	wget "https://raw.githubusercontent.com/LedgerHQ/ledger-wallet-api/master/ledger.js" \
		-q -O ledger.js
	wget "https://raw.githubusercontent.com/LedgerHQ/ledger-wallet-api/develop/ledger-eth.js" \
		-q -O ledger-eth.js

.PHONY: all install get_ledgerjs
