package ssh

import (
    "fmt"
    "io/ioutil"
    "os"
    "bufio"
    "strings"
    "strconv"
    "path/filepath"

    "golang.org/x/crypto/ssh"
)

type Config struct {
    Host string
    Port int
    ClientConfig *ssh.ClientConfig
}

func NewClientConfigWithKeyFile(username string, sshKey string, host string, port int, checkHostKey bool) (*Config, error) {
    var hostKey ssh.PublicKey

    c := &Config {
        Host: host,
        Port: port,
    }

    // Read private key
    key, err := ioutil.ReadFile(sshKey)
    if err != nil {
        return nil, err
    }

    // Create the Signer for this private key.
    signer, err := ssh.ParsePrivateKey(key)
    if err != nil {
        return nil, err
    }

    if checkHostKey {
        //arr := strings.Split(host, ":")
        hostKey, err = getHostKey(host, strconv.Itoa(port))
        if err != nil {
            return nil, err
        }
    }

    c.ClientConfig = &ssh.ClientConfig {
        User: username,
        Auth: []ssh.AuthMethod {
            //ssh.Password("chrYsal1s-adm1n"),
            //ssh.PublicKeyFile("/home/uthng/.ssh/ssh_servers"),
            ssh.PublicKeys(signer),
        },
        //HostKeyCallback: ssh.FixedHostKey(hostKey),
        //HostKeyCallback: ssh.InsecureIgnoreHostKey(),
        //HostKeyCallback: nil,
    }

    if checkHostKey {
        c.ClientConfig.HostKeyCallback = ssh.FixedHostKey(hostKey)
    } else {
        c.ClientConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()
    }

    return c, nil
}


/////////// PRIVATE FUNCTIONS ////////////////////////////
func getHostKey(host, port string) (ssh.PublicKey, error) {
    // $HOME/.ssh/known_hosts
    file, err := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var hostport string
    if port == "22" {
        // standard port assumes 22
        // 192.168.10.53 ssh-rsa AAAAB3Nza...vguvx+81N1xaw==
        hostport = host
    } else {
        // non-standard port(s)
        // [ssh.example.com]:1999,[93.184.216.34]:1999 ssh-rsa AAAAB3Nza...vguvx+81N1xaw==
        hostport = "[" + host + "]:" + port
    }

    scanner := bufio.NewScanner(file)
    var hostKey ssh.PublicKey
    for scanner.Scan() {
        fields := strings.Split(scanner.Text(), " ")
        if len(fields) != 3 {
            continue
        }
        if strings.Contains(fields[0], hostport) {
            var err error
            hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
            if err != nil {
                return nil, err
            }
            break // scanning line by line, first occurrence will be returned
        }
    }

    if hostKey == nil {
        return nil, fmt.Errorf("No hostkey for %s", host+":"+port)
    }

    return hostKey, nil
}

