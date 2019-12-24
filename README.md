# SASL/OAUTHBEARER AccessTokenProvider

## Overview 
This pkg is intended to be used as a complement to [Shopify's sarama library](https://github.com/Shopify/sarama). It provides an implementation of the sarama.AccessTokenProvider interface to be employed by clients using the SASL/OAUTHBEARER security protocol for Apache Kafka.

This package leverages Golang's oauth2 and clientcredentials packages to perform the 2 legged client credentials flow to obtain an Access Token outside the context of a user. 

## Usage

```
go get github.com/damiannolan/sasl
```

Import as:

```go
import (
    "github.com/damiannolan/sasl/oauthbearer"
)
```

The appropriate configuration settings must be set on the `sarama.Config`. It is recommended to use this protocol over TLS.

```go
cfg := sarama.NewConfig()

cfg.Net.SASL.Enable = true
cfg.Net.SASL.Mechanism = sarama.SASLTypeOAuth
cfg.Net.SASL.AccessTokenProvider = NewTokenProvider(clientID, clientSecret, tokenURL)
```

## References

- [Configuring SASL/OAUTHBEARER](https://docs.confluent.io/current/kafka/authentication_sasl/authentication_sasl_oauth.html)
- [KIP-255 OAuth Authentication via SASL/OAUTHBEARER](https://cwiki.apache.org/confluence/pages/viewpage.action?pageId=75968876)
- [sarama](https://github.com/Shopify/sarama)
- [oauth2](https://godoc.org/golang.org/x/oauth2)
- [oauth2/clientcredentials](https://godoc.org/golang.org/x/oauth2)