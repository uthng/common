package k8s

import (
    "strconv"
    //"errors"

    //"k8s.io/client-go/kubernetes"
    "k8s.io/client-go/rest"
)

// Config is an embedded type of Kubernetes' rest client Config
type Config struct {
    *rest.Config
}

// NewConfigFromRestTls returns a new config with TLS certificate files
func NewConfigFromRestTlsFile(serverName, host string, port int, apiPath, caFile, certFile, keyFile string) (*Config) {
    config := &Config {
        Config: &rest.Config {
            Host: host + ":" + strconv.Itoa(port),
            APIPath: apiPath,
            TLSClientConfig: rest.TLSClientConfig {
                Insecure: false,
                ServerName: serverName,
                CertFile: certFile,
                KeyFile: keyFile,
                CAFile: caFile,
            },
        },
    }

    return config
}

// NewConfigFromRestTls returns a new config with TLS certificate data
func NewConfigFromRestTlsData(serverName, host string, port int, apiPath string, caData, certData, keyData []byte) (*Config) {
    config := &Config {
        Config: &rest.Config {
            Host: host + ":" + strconv.Itoa(port),
            APIPath: apiPath,
            TLSClientConfig: rest.TLSClientConfig {
                Insecure: false,
                ServerName: serverName,
                CertData: certData,
                KeyData: keyData,
                CAData: caData,
            },
        },
    }

    return config
}

