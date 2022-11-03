# You are given a large JSON array of NFT collection stats as input (this is just an example)
# Program needs to run in the most optimal way from time and space complexity perspective (please add comments on your choice of implementation).
nft_collection_stats = [
{"collection_address": "0xabc", "collection_keywords": ["bored", "ape"], "metadata_link": "https://abc.xyz", "owners": ["0xbcd"]  },
{"collection_address": "0xdef", "collection_keywords": ["yellow", "ape"], "metadata_link": "https://def.xyz", "owners": ["0xbcd", "0xbde"]  },
]

# Create a method/function to return all NFT collections with a given keyword 
# Hint: it needs to work as fast as possible
def getNFTCollectionByKeyWord(keyword) :
  pass

# Add new NFT to the array. Method needs to check for duplicates - in case an NFT exists, 
# it needs to return an error. Customers want to add NFT-s fast, so it has to work as fast as possible.
# NFT is a duplicate if all the fields are the same.
# JSONObject has the same data as in the items in the nft_collection_stats
def addNFT(JSONObject) :
  pass

# We want to know which are the most popular keywords describing the NFT-s 
# and you have to write a method to return the top 5 most popular keywords, 
# with the most popular one first
# Hint: the number of keywords is large and remember the time and space complexity
def returnTop5Keywords() :
  pass