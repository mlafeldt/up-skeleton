# Up Skeleton

A skeleton project showing how to deploy a Go service as an AWS Lambda function via [Up](https://apex.github.io/up/).

Note: While I haven't used Up in production yet, it definitely is one of the easiest ways to prototype microservices in AWS.

## Usage

```bash
# Set AWS credentials
$ export AWS_ACCESS_KEY_ID=...
$ export AWS_SECRET_ACCESS_KEY=...
$ export AWS_REGION=...

# Deploy the service to staging and print its endpoint once done
$ make stage

# Test the service endpoint
$ curl `up url staging`
Hello World!

# Deploy the service to production
$ make prod

# Delete all AWS resources again
$ make destroy

# Test service locally
$ make local PORT=10000
```
