package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

func main() {

	// Check all of our arguments
	if len(os.Args) < 3 {
		fmt.Println(`Must pass arguments:

		(1) string to search for
		(2) how many matches to find
		(3 - optional) 'true' if you want to allow case insensitive matches

		A csv file will be generated with the list of private keys and addresses.
		`)

		return
	}
	str, cstr := os.Args[1], os.Args[2]

	if str[:1] != "1" {
		fmt.Println("The string to search for must start with a 1.")
		return
	}

	target, err := strconv.Atoi(cstr)
	if err != nil {
		fmt.Println("Second arg must be a number.")
		return
	}

	var ignore bool
	if len(os.Args) > 3 && os.Args[3] == "true" {
		ignore = true
	}

	// Open our csv file
	output, err := os.Create(str + ".csv")
	if err != nil {
		fmt.Printf("Couldn't create csv file: %s\n", err)
		return
	}

	_, err = output.WriteString("number,address,private key\n")
	if err != nil {
		fmt.Printf("Couldn't write to csv file: %s\n", err)
		return
	}

	var count, hits int

	for hits < target {
		// Create private key...
		key, err := btcec.NewPrivateKey(btcec.S256())
		if err != nil {
			fmt.Printf("Error creating new private key: %s\n", err)
			return
		}

		// ...convert that to WIF...
		wif, err := btcutil.NewWIF(key, &chaincfg.MainNetParams, true)
		if err != nil {
			fmt.Printf("Error creating WIF from private key: %s\n", err)
			return
		}

		// ...and get the address to compare to our search parameter
		pk, err := btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeCompressed(), &chaincfg.MainNetParams)
		if err != nil {
			fmt.Printf("Error getting address from WIF: %s\n", err)
			return
		}

		addr := pk.EncodeAddress()

		// If they are willing to ignore case, compare lower case. Otherwise, compare the string directly
		if ignore && strings.ToLower(addr[:len(str)]) == strings.ToLower(str) || !ignore && addr[:len(str)] == str {
			hits++
			fmt.Printf("%d: %s\n", hits, addr)

			// Update our file so the user can cancel at any time and not lose everything
			_, err = output.WriteString(fmt.Sprintf("%d,%s,%s\n", hits, addr, wif.String()))
			if err != nil {
				fmt.Printf("Couldn't write to csv file: %s\n", err)
				return
			}
		}

		// Make a basic af activity indicator
		count++
		fmt.Printf("Searching: %d\r", count)
	}

	// End
	fmt.Printf("Done! Found %d matching keys after trying %d possibilies.\n", hits, count)
}
