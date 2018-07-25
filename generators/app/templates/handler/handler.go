package handler

import (
	"fmt"

	"bitbucket.org/credomobile/govault"
	"bitbucket.org/credomobile/{{lambda_name}}/process"
	"github.com/aws/aws-lambda-go/events"
)
// CHANGEME
// Parameterize this function based on the type of event you will be handling.
// See https://github.com/aws/aws-lambda-go/tree/master/events for event types to use
func EventHandler() error {
	if !initialized {
		err := initialize()
		if err != nil {
			return err
		}
	}

	logger.Info().Msg(fmt.Sprintf("Lambda %s initialized with ID: %s", lambdaName, lambdaID.String()))

	httpClient, err = consulAndVaultHttpCient()
	if err != nil {
		logger.Fatal().Err(err).Msg("Failure to initialize http client")
		return err
	}

	consulConfig := api.DefaultConfig()
	consulConfig.Address = consulUrl
	consulConfig.HttpClient = httpClient
	consulApiClient, err = api.NewClient(consulConfig)

	consulClient := goconsul.NewClient(consulUrl, httpClient)

	err = consulClient.ReadConfig(lambdaName, config)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failure to read config")
		return err
	}

	tokenRes, err := govault.GetIamBasedVaultToken(httpClient, consulConfiguration.VaultUrl, lambdaName)
	if err != nil {
		logger.Error().Err(err).Msg("Failure to retrieve vault token")
		return err
	}

	vaultClient := govault.NewClient(consulConfiguration.VaultUrl, tokenRes.ClientToken, httpClient)

	creds, err := vaultClient.GetNonRefreshingConnectionCredentials("fakepath", lambdaName)
	if err != nil {
		logger.Error().Err(err).Msg("Failure to get creds from Vault for fake service")
		return err
	}

	// do something with the fake creds you got from vault

	eventProcessor := process.New(logger, nil)

	/* CHANGEME
	 * Replace the "event" variable here with the param that is passed into this function. 
	 * Also replace the record.Event with record.<member> based on the struct of the event you are processing
	 * from https://github.com/aws/aws-lambda-go/tree/master/events 
	 */

	for _, record := range event.Records {
		err := eventProcessor.ProcessRecord(record.Event)
		if err != nil {
			logger.Error().
				Err(err).
				Str("dataBytes", string(record.Event.Data[:])).
				Str("sequenceNumber", record.Event.SequenceNumber).
				Str("partitionKey", record.Event.PartitionKey).
				Msg("Failed to process event")
		}
	}

	return nil
}

func consulAndVaultHttpCient() (*http.Client, error) {
	caCerts, err := ioutil.ReadFile("./vault-cas.crt")
	if err != nil {
		return nil, err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCerts)

	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	return &http.Client{Transport: transport}, nil
}
