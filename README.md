# Basic AF vanity address generator

This is a vanity address generator for Bitcoin (non-segwit, non-multisig, legacy addresses only). A vanity address starts with a specific sequence of characters that have meaning (or not). For example, Satoshi Dice addresses all start with `1dice`.

I created this as a quick'n'dirty™ solution for generating Bitcoin addresses on OSX. You don't have to use it for vanity addresses; just pass `1` as your search term to generate random addresses.

**Disclaimer: This generator is _extremely_ slow! If you are generating a large number of addresses, I suggest you search for a short sequence and enable case-insensitivity.**

## How to use

`./bafvanitygen [search term] [count] [case-insensitive]`

`[search term]` - What you want your addresses to all start with. Must start with the number 1, followed by your string.

`[count]` - How many matching addresses to find

`[case-insensitive]` - If you are okay with the sequence using random capitalization, enter `true` here, otherwise ignore.

Matching addresses will be printed out to the screen as they are found. A `.csv` file, named after the string you searched for, will be generated in the same directory containing the addresses and corresponding private keys.

## Credits

None, I just pooped this out in a few hours on a Wednesday morning.