# v1.4.9 (2022-07-05)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.4.8 (2022-06-29)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.4.7 (2022-06-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.4.6 (2022-05-19)

* No change notes available for this release.

# v1.4.5 (2022-05-17)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.4.4 (2022-04-25)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.4.3 (2022-03-30)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.4.2 (2022-03-24)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.4.1 (2022-03-23)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.4.0 (2022-03-08)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.3.0 (2022-02-24)

* **Feature**: API client updated
* **Feature**: Adds RetryMaxAttempts and RetryMod to API client Options. This allows the API clients' default Retryer to be configured from the shared configuration files or environment variables. Adding a new Retry mode of `Adaptive`. `Adaptive` retry mode is an experimental mode, adding client rate limiting when throttles reponses are received from an API. See [retry.AdaptiveMode](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/aws/retry#AdaptiveMode) for more details, and configuration options.
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.2.1 (2022-01-28)

* **Bug Fix**: Updates SDK API client deserialization to pre-allocate byte slice and string response payloads, [#1565](https://github.com/aws/aws-sdk-go-v2/pull/1565). Thanks to [Tyson Mote](https://github.com/tysonmote) for submitting this PR.

# v1.2.0 (2022-01-14)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.1.0 (2022-01-07)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.1 (2021-12-02)

* **Bug Fix**: Fixes a bug that prevented aws.EndpointResolverWithOptions from being used by the service client. ([#1514](https://github.com/aws/aws-sdk-go-v2/pull/1514))
* **Dependency Update**: Updated to the latest SDK module versions

# v1.0.0 (2021-11-19)

* **Release**: New AWS service client module
* **Dependency Update**: Updated to the latest SDK module versions

