package configuration

func ApiUserConfig() string {
	var baseUrl = "https://test-account-service.azurewebsites.net"
	var apiPath = "/api/v1/jguard/profile"
	var token = "?code=rEMjj3NwvpK2XizxqeQEd5UvGpHDUGfXpLVAyAjNFl8CAzFuel-Idw=="
	res := baseUrl + apiPath + token
	return res
}
