package main

type Marketplace struct {
	Name     string
	ID       string
	Endpoint string
	Country  string
}

var marketplaces = map[string]Marketplace{
	"CA": Marketplace{
		Name: "Amazon Canada",
		ID: "A2EUQ1WTGCTBG2",
		Endpoint: "https://mws.amazonservices.com",
		Country: "CA",
	},
	"A2EUQ1WTGCTBG2": Marketplace{
		Name: "Amazon Canada",
		ID: "A2EUQ1WTGCTBG2",
		Endpoint: "https://mws.amazonservices.com",
		Country: "CA",
	},

	"US": Marketplace{
		Name: "Amazon United States",
		ID: "ATVPDKIKX0DER",
		Endpoint: "https://mws.amazonservices.com",
		Country: "US",
	},
	"ATVPDKIKX0DER": Marketplace{
		Name: "Amazon United States",
		ID: "ATVPDKIKX0DER",
		Endpoint: "https://mws.amazonservices.com",
		Country: "US",
	},

	"MX": Marketplace{
		Name: "Amazon Mexico",
		ID: "A1AM78C64UM0Y8",
		Endpoint: "https://mws.amazonservices.com",
		Country: "MX",
	},
	"A1AM78C64UM0Y8": Marketplace{
		Name: "Amazon Mexico",
		ID: "A1AM78C64UM0Y8",
		Endpoint: "https://mws.amazonservices.com",
		Country: "MX",
	},

	"ES": Marketplace{
		Name: "Amazon Spain",
		ID: "A1RKKUPIHCS9HS",
		Endpoint: "https://mws-eu.amazonservices.com",
		Country: "ES",
	},
	"A1RKKUPIHCS9HS": Marketplace{
		Name: "Amazon Spain",
		ID: "A1RKKUPIHCS9HS",
		Endpoint: "https://mws-eu.amazonservices.com",
		Country: "ES",
	},

	"UK": Marketplace{ // alias of GB
		Name: "Amazon United Kingdom",
		ID: "A1F83G8C2ARO7P",
		Endpoint: "https://mws-eu.amazonservices.com",
		Country: "GB",
	},
	"GB": Marketplace{
		Name: "Amazon United Kingdom",
		ID: "A1F83G8C2ARO7P",
		Endpoint: "https://mws-eu.amazonservices.com",
		Country: "GB",
	},
	"A1F83G8C2ARO7P": Marketplace{
		Name: "Amazon United Kingdom",
		ID: "A1F83G8C2ARO7P",
		Endpoint: "https://mws-eu.amazonservices.com",
		Country: "GB",
	},

	"FR": Marketplace{
		Name: "Amazon France",
		ID: "A13V1IB3VIYZZH",
		Endpoint: "https://mws-eu.amazonservices.com",
		Country: "FR",
	},
	"A13V1IB3VIYZZH": Marketplace{
		Name: "Amazon France",
		ID: "A13V1IB3VIYZZH",
		Endpoint: "https://mws-eu.amazonservices.com",
		Country: "FR",
	},

	"DE": Marketplace{
		Name: "Amazon Germany",
		ID: "A1PA6795UKMFR9",
		Endpoint: "https://mws-eu.amazonservices.com",
		Country: "DE",
	},
	"A1PA6795UKMFR9": Marketplace{
		Name: "Amazon Germany",
		ID: "A1PA6795UKMFR9",
		Endpoint: "https://mws-eu.amazonservices.com",
		Country: "DE",
	},

	"IT": Marketplace{
		Name: "Amazon Italy",
		ID: "APJ6JRA9NG5V4",
		Endpoint: "https://mws-eu.amazonservices.com",
		Country: "IT",
	},
	"APJ6JRA9NG5V4": Marketplace{
		Name: "Amazon Italy",
		ID: "APJ6JRA9NG5V4",
		Endpoint: "https://mws-eu.amazonservices.com",
		Country: "IT",
	},

	"BR": Marketplace{
		Name: "Amazon Brazil",
		ID: "A2Q3Y263D00KWC",
		Endpoint: "https://mws.amazonservices.com",
		Country: "BR",
	},
	"A2Q3Y263D00KWC": Marketplace{
		Name: "Amazon Brazil",
		ID: "A2Q3Y263D00KWC",
		Endpoint: "https://mws.amazonservices.com",
		Country: "BR",
	},

	"IN": Marketplace{
		Name: "Amazon India",
		ID: "A21TJRUUN4KGV",
		Endpoint: "https://mws.amazonservices.in",
		Country: "IN",
	},
	"A21TJRUUN4KGV": Marketplace{
		Name: "Amazon India",
		ID: "A21TJRUUN4KGV",
		Endpoint: "https://mws.amazonservices.in",
		Country: "IN",
	},

	"CN": Marketplace{
		Name: "Amazon China",
		ID: "AAHKV2X7AFYLW",
		Endpoint: "https://mws.amazonservices.com.cn",
		Country: "CN",
	},
	"AAHKV2X7AFYLW": Marketplace{
		Name: "Amazon China",
		ID: "AAHKV2X7AFYLW",
		Endpoint: "https://mws.amazonservices.com.cn",
		Country: "CN",
	},

	"JP": Marketplace{
		Name: "Amazon Japan",
		ID: "A1VC38T7YXB528",
		Endpoint: "https://mws.amazonservices.jp",
		Country: "JP",
	},
	"A1VC38T7YXB528": Marketplace{
		Name: "Amazon Japan",
		ID: "A1VC38T7YXB528",
		Endpoint: "https://mws.amazonservices.jp",
		Country: "JP",
	},

	"AU": Marketplace{
		Name: "Amazon Australia",
		ID: "A39IBJ37TRP1C6",
		Endpoint: "https://mws.amazonservices.com.au",
		Country: "AU",
	},
	"A39IBJ37TRP1C6": Marketplace{
		Name: "Amazon Australia",
		ID: "A39IBJ37TRP1C6",
		Endpoint: "https://mws.amazonservices.com.au",
		Country: "AU",
	},
}

func GetAllMarketplaces() map[string]Marketplace {
	return marketplaces
}

func GetMarketplace(key string) Marketplace {
	return marketplaces[key]
}

func GetMarketplaceEndpoint(key string) string {
	marketplace := GetMarketplace(key)
	return marketplace.Endpoint
}

func GetMarketplaceID(key string) string {
	marketplace := GetMarketplace(key)
	return marketplace.ID
}
