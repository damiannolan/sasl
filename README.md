# SASL/OAUTHBEARER Access Token Provider

## Overview 

This package is intended to be used as a complement to [Shopify/sarama](https://github.com/Shopify/sarama). It provides an implementation of the `sarama.AccessTokenProvider` interface to be employed by clients using the SASL/OAUTHBEARER mechanism for Apache Kafka.

The very popular `golang/oauth2` and `golang/oauth2/clientcredentials` are leveraged to perform the 2 legged client credentials flow to obtain an Access Token outside the context of a user. 

## Installation
```
go get github.com/damiannolan/sasl/oauthbearer
```

## Usage

Configure `sarama.Config` as desired for producer/consumer clients and enable SASL/OAUTHBEARER with the appropriate settings. For production setups it is recommended to use this authentication mechanism over a secure connection. This can be achieved by setting `Net.TLS.Enable` to `true` and providing a `*tls.Config` through `Net.TLS.Config`.

```go
import (
    "github.com/damiannolan/sasl/oauthbearer"
)

func main() {
    cfg := sarama.NewConfig()

    cfg.Net.SASL.Enable = true
    cfg.Net.SASL.Mechanism = sarama.SASLTypeOAuth
    cfg.Net.SASL.AccessTokenProvider = oauthbearer.NewTokenProvider(clientID, clientSecret, tokenURL)
    ...
}
```

## References

- [Configuring SASL/OAUTHBEARER](https://docs.confluent.io/current/kafka/authentication_sasl/authentication_sasl_oauth.html)
- [KIP-255 OAuth Authentication via SASL/OAUTHBEARER](https://cwiki.apache.org/confluence/pages/viewpage.action?pageId=75968876)
- [sarama](https://github.com/Shopify/sarama)
- [oauth2](https://godoc.org/golang.org/x/oauth2)
- [oauth2/clientcredentials](https://godoc.org/golang.org/x/oauth2)