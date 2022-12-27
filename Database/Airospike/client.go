package aerospike

import (
	aerospike "github.com/aerospike/aerospike-client-go"
)

func GetAerospikeClient() (*aerospike.Client, error) {
	client, err := aerospike.NewClient("192.168.1.41", 3000)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func CloseAerospikeClient(client *aerospike.Client) error {
	client.Close()
	return nil
}
