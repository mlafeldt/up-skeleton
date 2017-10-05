# Up Skeleton

A skeleton Go service deployed via [Up](https://apex.github.io/up/).

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

# Delete all resources again
$ make destroy
```
