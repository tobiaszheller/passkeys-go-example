# passkeys go example

This is demo of how to add passkeys instead of password for webapp.

It does few simplifications:
- user data is stored in memory
- session data is stored in memory

For production ready support, those material provides good suggestions:
- https://developers.google.com/identity/passkeys/developer-guides/server-introduction
- https://developers.yubico.com/Passkeys/Passkey_relying_party_implementation_guidance/Initialize_a_passkey_relying_party.html

## How to run

`go run main.go`