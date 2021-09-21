package client

type Client struct {
	*HenrikdevClient
	*UnofficialValorantAPIClient
}

func New() *Client {
	hClient := NewHenrikdevClient()
	uClient := NewUnofficialValorantAPIClient()
	return &Client{
		HenrikdevClient:             hClient,
		UnofficialValorantAPIClient: uClient,
	}
}
