package main

import (
	"github.com/SeedlngTeam/Core/migrations"
)

// func catchAll(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "<h1>Welcome to Seedlng!!")
// }

// Plaid Example

// configuration := plaid.NewConfiguration()

// configuration.AddDefaultHeader("PLAID-CLIENT-ID", os.Getenv("CLIENT_ID"))
// configuration.AddDefaultHeader("PLAID-SECRET", os.Getnev("SECRET"))
// configuration.UseEnvironment(plaid.Sandbox)
// client := plaid.NewAPIClient(configuration)

// func getAccessToken(c *gin.Context) {
// 	ctx := context.Background()
// 	publicToken := c.PostForm("public_token")

// 	// exchange the public_token for an access_token
// 	exchangePublicTokenReq := plaid.NewItemPublicTokenExchangeRequest(sandboxPublicTokenResp.GetPublicToken())
// 		exchangePublicTokenResp, _, err := client.PlaidApi.ItemPublicTokenExchange(ctx).ItemPublicTokenExchangeRequest(
// 			*exchangePublicTokenReq,
// 		).Execute()

// 	accessToken := exchangePublicTokenResp.GetAccessToken()
// 	itemID := exchangePublicTokenResp.GetItemId()

// 	fmt.Println("public token: " + publicToken)
// 	fmt.Println("access token: " + accessToken)
// 	fmt.Println("item ID: " + itemID)

// 	c.JSON(http.StatusOK, gin.H{
// 		"access_token": accessToken,
// 		"item_id": itemID,
// 	})
// }

func main() {
	// r := gin.Default()
	// r.POST("/exchange_public_token", getAccessToken)

	// err := r.Run(":8000")
	// if err != nil {
	// 	panic("unable to start server")
	// }

	migrations.Migrate()

	// http.HandleFunc("/", catchAll)
	// fmt.Print("Server Staring on :3000 ...")
	// http.ListenAndServe("https://localhost:3000", nil)
}
