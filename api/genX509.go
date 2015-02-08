/* GenX509Cert
Thanks! : http://jan.newmarch.name/go/security/chapter-security.html
 */
package main

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "crypto/x509/pkix"
    "encoding/gob"
    "encoding/pem"
    "fmt"
    "math/big"
    "os"
    "time"
)

func main() {
    random := rand.Reader

    var key rsa.PrivateKey
    loadKey("keys/private.key", &key)

    now := time.Now()
    then := now.Add(100 * 60 * 60 * 24 * 365 * 1000 * 1000 * 1000)
    template := x509.Certificate{
        SerialNumber: big.NewInt(1),
        Subject: pkix.Name{
            CommonName:   "www.i-aor.com",
            Organization: []string{"i-aor.com"},
        },
        //  NotBefore: time.Unix(now, 0).UTC(),
        //  NotAfter:  time.Unix(now+60*60*24*365, 0).UTC(),
        NotBefore: now,
        NotAfter:  then,

        SubjectKeyId: []byte{1, 2, 3, 4},
        KeyUsage:     x509.KeyUsageCertSign | x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,

        BasicConstraintsValid: true,
        IsCA:                  true,
        DNSNames:              []string{"www.i-aor.com", "localhost"},
    }
    derBytes, err := x509.CreateCertificate(random, &template,
        &template, &key.PublicKey, &key)
    checkError(err)

    certCerFile, err := os.Create("keys/www.i-aor.com.cer")
    checkError(err)
    certCerFile.Write(derBytes)
    certCerFile.Close()

    certPEMFile, err := os.Create("keys/www.i-aor.com.pem")
    checkError(err)
    pem.Encode(certPEMFile, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
    certPEMFile.Close()

    keyPEMFile, err := os.Create("keys/private.pem")
    checkError(err)
    pem.Encode(keyPEMFile, &pem.Block{Type: "RSA PRIVATE KEY",
        Bytes: x509.MarshalPKCS1PrivateKey(&key)})
    keyPEMFile.Close()
}

func loadKey(fileName string, key interface{}) {
    inFile, err := os.Open(fileName)
    checkError(err)
    decoder := gob.NewDecoder(inFile)
    err = decoder.Decode(key)
    checkError(err)
    inFile.Close()
}

func checkError(err error) {
    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        os.Exit(1)
    }
}