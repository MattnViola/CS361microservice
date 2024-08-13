# CS361 Microservice A

## Author: Matthew Norman

# Summary

This is an OAuth Microservice that takes a lot of the work out of setting up an OAuth account. As of this writing, it's currently deployed and reachable at https://cs361micro-4qdz6le7kq-uc.a.run.app/.

# Requesting Data

### This is done through a redirect of the user, and then a post request to the microservice :

#### Redirect => https://cs361micro-4qdz6le7kq-uc.a.run.app/login?user_callback=**<Parameter 1>**

##### Parameter 1: user_callback: The URL you want the user to be sent to after returning to your server.

#### Post => https://cs361microhttps://cs361micro-4qdz6le7kq-uc.a.run.app/userdata

##### In the post body, have a json with { key : <user's key>} The key will be received when the user returns.

# Recieving Data

### The user will be redirected back to the user_callback with a parameter, key, which will equal their key. It can then be used as the value in the post json.

### An example post request can be seen in test.mjs in the /tester folder.

## Example Payload:

{
"aud": "...",
"exp": 1723024604,
"family_name": "Norman",
"given_name": "Matthew",
"iat": 1722988604,
"iss": "https://secure361login.us.auth0.com/",
"key": "your key here",
"name": "Matthew Norman",
"nickname": "matt",
"picture": "your picture here",
"sid": "other gibberish",
"updated_at": "2024-08-06T03:02:31.139Z",
}

# UML Sequence Diagram

![image](https://github.com/user-attachments/assets/d3ba4eb1-3ce4-4af5-af53-0144c127f170)
