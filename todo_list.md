1. Dynamic error codes and messages with a struct variable. Normally, actual errors should be log into an external service such as Promethus, Graylog, Elasticsearch. In this app, I respond the actual error as JSON response to client and it's not best practice.

2. Should we return HTTP STATUS OK (200) for all responses or HTTP STATUS according to the event?

3. Validator package message can be override and I can type custom message accordingly.

4. Test cases doesn't suffice and they can be various and test more scenarios.