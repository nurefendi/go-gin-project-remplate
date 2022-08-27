package config

import (
	"net"
	"os"
	"time"

	"github.com/Shopify/sarama"
)

var KAFKA *sarama.Config

func InitKafkaClient() (sarama.AsyncProducer, error) {

    // cf := filepath.Join(certPath, certFile)
    // kf := filepath.Join(certPath, keyFile)

    // Log cert and key path
    // log.Debugln(cf)
    // log.Debugln(kf)

    // Read the cert in
    // certIn, err := ioutil.ReadFile(cf)

    // if err != nil {
    //     log.Error("cannot read cert", err)
    //     return nil, err
    // }

    // Read & decode the encrypted key file with the pass to make tls work
    // keyIn, err := ioutil.ReadFile(kf)
    // if err != nil {
    //     log.Error("cannot read key", err)
    //     return nil, err
    // }

    // Decode and decrypt our PEM block as DER
    // decodedPEM, _ := pem.Decode([]byte(keyIn))
    // decrypedPemBlock, err := x509.DecryptPEMBlock(decodedPEM, []byte("m4d3ups3curity4k4fka?"))
    // if err != nil {
    //     log.Error("cannot decrypt pem block", err)
    //     return nil, err
    // }

    // Parse the DER encoded block as PEM
    // rsaKey, err := x509.ParsePKCS1ParrivateKey(decrypedPemBlock)
    // if err != nil {
    //    log.Error("failed to parse rsa as pem", err)
    //    return nil, err
    // }

    // Marshal the pem encoded RSA key to bytes in memory
    // pemdata := pem.EncodeToMemory(
    //    &pem.Block{
    //         Type:  "RSA PRIVATE KEY",
    //         Bytes: x509.MarshalPKCS1PrivateKey(rsaKey),
    //     },
    // )
    // if err != nil {
    //     log.Error("cannot marshal rsa as pem in memory", err)
    //     return nil, err
    // }

    // Load our decrypted key pair
    // crt, err := tls.X509KeyPair(certIn, pemdata)
    // if err != nil {
    //     log.Error("cannot load key pair", err)
    //     return nil, err
    // }
    config := sarama.NewConfig()
    config.Net.TLS.Enable = false
    // config.Net.TLS.Config = &tls.Config{
	// 	Certificates: []tls.Certificate{crt},
	// 	InsecureSkipVerify: true,
	// }

    // Setting this allows us not to read from successes channel
    config.Producer.Return.Successes = false
    // Setting this allows us not to read from errors channel
    config.Producer.Return.Errors = false
	config.Net.WriteTimeout = 5 * time.Second
	config.Producer.Retry.Max = 0
    client, err := sarama.NewClient([]string{net.JoinHostPort(os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT"))}, config)
    if err != nil {
        return nil, err
    }
    return sarama.NewAsyncProducerFromClient(client)
}