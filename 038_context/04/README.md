Context makes it possible to manage a chain of calls within the same call path by signaling context’s Done channel.

source: https://rakyll.org/leakingctx/

#### My Notes

- Context carries the context of the requst throughout the backend layer e.g. from recieving the request to accessing the database

- Therefore it can be used to store values that are pertinent to the request e.g. session id,user id

- It can also be used to timeout or cancel asynchronous processes