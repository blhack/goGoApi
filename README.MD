# GoGoApi

## Installation

## Config

See config/config.go.  Server.crt and Server.key are required and can be obtained either via openssl, letsencrypt, or a commercial vendor.  Indicate the path to these files.

Modify DbUser and DbPass to reflect your local parameters.  By default the API will look for a DB called "goGoApi", but you should create a database with whatever name you want and relect it in this file.

One option for loading the db schema would be: mysql -u $user -p goGoApi < schema.sql

Cookiedomain should also be set if you are working elsewhere than localhost.  Similarly, if you are serving your frontend from a different sub/domain than your api (for instance: www.example.com = frontend and api.example.com = api.), add these frontend domains to the list under "WhitelistOrigins"

Rebuild the binary after modifying this config file.

## Build

     cd goGoApi
     go get ./...
     go build -o goGoApi goGoApi.go

# Start
     ./goGoApi

(Will fail if your user does not have rights to open ports 80 and 443.  You could solve this with sudo ./goGoApi, however running as root as not always/usually not advisable)

## API Conventions:	
### Binary requests will return JSON strings which adhere to the following convention:

     {"success":true/false,"status":"relevant additional (optional) information related to the request"}


### THe API is organized into a branching tree structure:

	/auth/
	/items/
    /etc/etc/ (this will change)

Relevant endpoints will exist below these structures.  A trailing forward slash '/' implies a container which contains additional API resources.  A lack of a trailing slash '/' implies an endpoint.  This is modelled after a unix file system.


## API Resources

### Who am I?

    /auth/whoami

If you are authenticated will return:

{"success":true,"status":"ryan"}

If you are NOT authenticated will return:

{"success":false,"status":""}

### Login

     /auth/login?username=$foo&password=bar

If you provide a incorrect username and password combination will return:

{"success":false,"status":"invalid username/password combination"}

If you provide a correct/registered username and password combination will return:

{"success":true,"status":"user authenticated"}

### Register

    /auth/register?username=$foo&pass1=$bar&pass2=$baz

All three of these are required.  Will return TRUE if a user is registered.  Will return false if user is NOT registered for any of the following reasons:

* Invalid query parameters
* User exists

Additional information is provided in the "status" field of the response.
