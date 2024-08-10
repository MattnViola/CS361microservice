# CS361 Microservice A
## Author: Matthew Norman

# Summary

This is an OAuth Microservice that takes a lot of the work out of setting up an OAuth account. As of this writing, it's currently deployed and reachable at https://cs361micro-4qdz6le7kq-uc.a.run.app/.

# Requesting Data

### This is done through a redirect of the user: 

#### Redirect => https://cs361micro-4qdz6le7kq-uc.a.run.app/login?user_callback=**<Parameter 1>**&data_callback=**<Parameter 2>** 

##### Parameter 1: user_callback: The URL you want the user to be sent to after returning to your server.

##### Parameter 2: data_callback: The URL you want the user's data to be posted to. You must be listening for this request (see Retrieving data).

# Recieving Data 

### The user will be redirected back to the user_callback with a parameter, key, which will equal their key. 
### You must be listening for post requests to the data_callback parameter seen when requesting data. The payload will be a JSON object with various information about the user, including their key and other information.

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
