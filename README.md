#About this project - simple blockchain implementation 

This project shows a simple Go-based end-to-end blockchain workflow. Each wallet is given a distinct private key, public key, and blockchain address encoded in Base58. The process starts with creating a wallet address using ECDSA keys. 
These keys allow the wallet to sign transactions locally, guaranteeing authenticity and guarding against manipulation. 
The wallet server creates the request when the user submits a transaction, signs it with the private key, and then transmits the signed transaction to the blockchain server.

By examining the signature, confirming the sender's public key, and making sure the sender has sufficient funds, the blockchain server validates the transaction. 
The transaction is put into a transaction pool, which serves as a holding area for ongoing operations, after it has been validated. 
These pending transactions just wait to be added to the subsequent block; they are not yet a part of the blockchain.

The blockchain server gathers transactions from the pool, compiles them, calculates the proof-of-work, and creates a new block when mining is initiated. 
After that, this freshly mined block is added to the real blockchain, permanently documenting every transaction it contains. The transaction pool is cleared after mining, balances are updated appropriately, and the wallet user interface can refresh to display the updated wallet amount. 
The system's entire operation is shown in the pictures below, from wallet creation to transaction signing and, at the end, the mining procedure that adds them to the blockchain.

<img width="1914" height="850" alt="chain" src="https://github.com/user-attachments/assets/90436811-7ca1-4200-a21b-df8867fb8c2d" />
<img width="1315" height="631" alt="wallet" src="https://github.com/user-attachments/assets/69f906f6-a462-4e71-99c4-c6da109c7c90" />

