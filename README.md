# 🚇 sshtunnel

Ultra simple SSH tunnelling for Go programs.

Supports passphrase protected key files

## Installation

```bash
go get -u github.com/SteffenLoges/sshtunnel
```

Or better with `dep`:

```bash
dep ensure -add github.com/SteffenLoges/sshtunnel
```

## Example

```go
var auth ssh.AuthMethod
var err error

// Simple auth with password
auth = ssh.Password("YOUR_PASSWORD")

// ------------------------------------------

// Auth with plain text private key
// key := "-----BEGIN OPENSSH PRIVATE KEY-----\n" +
// 	"b3BlbnNzaC1rZXktdjEAAAAACmFlczI1Ni1jdHIAAAAGYmNyeXB0AAAAGAAAABAvYNehTJ\n" +
// 	"vg20muvaYEi5J+AAAAZAAAAAEAAAAzAAAAC3NzaC1lZDI1NTE5AAAAIMwxZkB2W4TG3unO\n" +
// 	"ysm35PApQssySCfxGJfq72WL02gzAAAAoP3vmBKeQ2elWudBsJRz6NoV14n4VVVQobogox\n" +
// 	"qoW6UF/aSw3Xf4fsTPUc48mxw7f9Tih+VEQ2MYrjXj4qwEBk9jd6/I/NPcqFVL1/fOjlcN\n" +
// 	"yIFlfWtXxzLmhaoTxfkfbZwZPLtiZn6HRGfNCztHhmCtLNGjM/Ey7UCoE3uejeyzvksfKZ\n" +
// 	"SPNuJmshHAKqAoHdw1FJJz4B9eS89u7ewCLs8=\n" +
// 	"-----END OPENSSH PRIVATE KEY-----"

// Use nil as 2nd parameter for unprotected key files
// auth, err = sshtunnel.ParsePrivateKey([]byte(key), []byte("YOUR_PASSPHRASE"))

// ------------------------------------------

// Auth with path to private key
// Use nil as 2nd parameter for unprotected key files
// auth, err = sshtunnel.ParsePrivateKeyFile("d:/id_ed25519", []byte("YOUR_PASSPHRASE"))

// ==========================================

if err != nil {
    panic(err)
}

// Setup the tunnel, but do not yet start it yet.
tunnel := sshtunnel.NewSSHTunnel(
    // User and host of tunnel server, it will default to port 22
    // if not specified.
    "ec2-user@jumpbox.us-east-1.mydomain.com",

    // AuthMethod specified above
    auth,

    // The destination host and port of the actual server.
    "dqrsdfdssdfx.us-east-1.redshift.amazonaws.com:5439",

    // The local port you want to bind the remote port to.
    // Specifying "0" will lead to a random port.
    "8443",
)

// You can provide a logger for debugging, or remove this line to
// make it silent.
tunnel.Log = log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds)

// Start the server in the background. You will need to wait a
// small amount of time for it to bind to the localhost port
// before you can start sending connections.
go tunnel.Start()
time.Sleep(100 * time.Millisecond)

// NewSSHTunnel will bind to a random port so that you can have
// multiple SSH tunnels available. The port is available through:
//   tunnel.Local.Port

// You can use any normal Go code to connect to the destination server
// through localhost. You may need to use 127.0.0.1 for some libraries.
//
// Here is an example of connecting to a PostgreSQL server:
conn := fmt.Sprintf("host=127.0.0.1 port=%d username=foo", tunnel.Local.Port)
db, err := sql.Open("postgres", conn)

// ...
```
