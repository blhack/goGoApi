package config

var HttpPort = "80"
var Domain = "localhost"
var CertFile = "server.crt"
var KeyFile = "server.key"
var CookieDomain = "localhost" //the domain to provide in cookies
var DbUser = ""
var DbPass = ""
var DbName = "goGoApi"

var WhiteListOrigins = map[string]bool{
	"localhost":      true,
	"localhost:3000": true,
}
