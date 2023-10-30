# NBC Proof of Reserve

To verify that the deposit addresses belong to our Bitcoin Multisig wallet, you can verify that each of these wallet addresses is generated from a combination of the 7 following signers’ public keys and the user’s NBC address as follows:

- Pull code from the repository: https://github.com/TrustlessComputer/proof-of-reserve


- Run `go run main.go <NBC-address>` command where <NBC-address> is the NBC address you’d like to verify.


- The command will return a corresponding Bitcoin deposit address for the <NBC-address> above.


For example, with the NBC address: `0xA73795E3caaED8F37c92530Fb939175054927175`, its corresponding Bitcoin deposit address is `bc1qjpn4qvqhxuh3hj0sndxqylzpc3wnqu84zwew29hm7zyjnqw9j5sqnmygf0`, which can be found on the TVL page’s Bitcoin deposit addresses on the app.

