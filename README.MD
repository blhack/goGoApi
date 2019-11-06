# Flex IMS -- The flexible Inventory Mangement System.

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
