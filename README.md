# NBC Proof of Reserve

NBC multisig signers:
```
Signer 1: 02bf0fd86a31568497c7635d9b48d48194cd12a3083ba04e599c0ccdb1b0ba955b
Signer 2: 024237a4f5fe8057ebf5ee890a892be3958ed3691ab7c0af84d72d197ec961bf98
Signer 3: 02b0305abe6d6bae7ee95e4a7fede1281b6d3df7f0d841d4edb155b1978b0835f7
Signer 4: 029097a2513dff395905fa3d7b4d1dce258608fd936d4552add96c6e7a4d2d5c3a
Signer 5: 02d6b28c3e9ca7cc870afc95192c3a6fa6de6a4acc9fb0a7dc8c8ca7fb46c4cfd4
Signer 6: 021bf77c146362f5a99f208b3512b52de22b0c5ba5e1eb4d5b7eeb610c013d032e
Signer 7: 03479cd4022fa8c4ebc16c7cd4fc9396c4cf6ebc37daed1ac09803390026785e1a
```

To verify that the deposit addresses belong to our Bitcoin Multisig wallet, you can verify that each of these wallet addresses is generated from a combination of the 7 signers’ public keys and the user’s NBC address as follows:

- Pull code from the repository: https://github.com/TrustlessComputer/proof-of-reserve


- Run `go run main.go <NBC-address>` command where <NBC-address> is the NBC address you’d like to verify.


- The command will return a corresponding Bitcoin deposit address for the <NBC-address> above.


For example, with the NBC address: `0xA73795E3caaED8F37c92530Fb939175054927175`, its corresponding Bitcoin deposit address is `bc1qjpn4qvqhxuh3hj0sndxqylzpc3wnqu84zwew29hm7zyjnqw9j5sqnmygf0`, which can be found on the TVL page’s Bitcoin deposit addresses on the app.

