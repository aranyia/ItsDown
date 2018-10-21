# It's Down

Application for performing recurring health checks of specific web services. Integrated with Slack via a web hook, to send status messages.

## Configuration

### Environment variables
The application can be configured using environment variables set.

|Key              |Value purpose                                   |Value example                   |
|-----------------|------------------------------------------------|--------------------------------|
|interval         |Period to fire status check, in seconds         |300                             |
|services         |The path of the services configuration JSON file|services.json                   |
|slack-webhook-url|Webhook URL to use for posting Slack messages   |https://hooks.slack.com/services|

### Service endpoints
Services to be checked should be defined in a JSON file as an array of objects. The service has to have a _name_ to be displayed and a _statusCheck_ definition, which is an object that requires the _url_ and _httpMethod_ fields to be defined.

Example configuration:
```json
[
  {
    "name": "integration SIT",
    "statusCheck": {
      "url": "https://127.0.0.1:8081/integration/health",
      "httpMethod": "GET"
     }
  },
  {
    "name": "integration PreProd",
    "statusCheck": {
      "url": "https://127.0.0.1:9080/integration/testpost",
      "httpMethod": "POST"
     }
  }
]
```