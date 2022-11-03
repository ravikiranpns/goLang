// To execute Go code, please declare a func main() in a package "main"

package main

import (
	"encoding/json"
	"fmt"
)

type NFTCollections struct {
	collection_address  string
	collection_keywords []string
	metadata_link       string
	owners              []string
}

var nft_collection_stats = []byte(`[{"collection_address": "0xabc", "collection_keywords": ["bored", "ape"], "metadata_link": "https://abc.xyz", "owners": ["0xbcd"]  },
{"collection_address": "0xdef", "collection_keywords": ["yellow", "ape"], "metadata_link": "https://def.xyz", "owners": ["0xbcd", "0xbde"]  }
]`)

/*Create a method/function to return all NFT collections with a given keyword
Hint: it needs to work as fast as possible
*/
func getNFTCollectionByKeyWord(keyword string) []NFTCollections {
	var NFTColl []NFTCollections

	//r := &NFTCollections{}
	buffer := []byte(nft_collection_stats)

	err := json.Unmarshal(buffer, &NFTColl)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	//fmt.Println(buffer)
	fmt.Printf("%+v", NFTColl)

	return NFTColl

}

func main() {

	NftColl := getNFTCollectionByKeyWord("yellow")

	fmt.Println(NftColl)

}
