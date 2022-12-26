package aerospike

import (
	"fmt"

	aerospike "github.com/aerospike/aerospike-client-go/v6"
)

var client *aerospike.Client
var Namespace = "TinyURL"
var Set = "URL"

func init() {
	client, _ = GetAerospikeClient()
}

func CreateRecordKey(key string) *aerospike.Key {
	recordKey, err := aerospike.NewKey(Namespace, Set, key)
	if err != nil {
		fmt.Printf("Error while creating record key: %v", err)
		return nil
	}
	return recordKey
}

func CreateRecord(key string, BinMap aerospike.BinMap) {
	recordKey := CreateRecordKey(key)
	if recordKey == nil {
		fmt.Printf("Error while creating record key: %v", key)
		return
	}
	client.Put(writePolicy, recordKey, BinMap)
}

func CheckRecordExists(key string) bool {
	recordKey := CreateRecordKey(key)
	if recordKey == nil {
		fmt.Printf("Error while creating record key: %v", key)
		return false
	}
	exists, err := client.Exists(readPolicy, recordKey)
	if err != nil {
		fmt.Printf("Error while checking record exists: %v", err)
		return false
	}
	return exists
}

func ReadMetaDataRecord(key string) (*aerospike.Record, error) {
	recordKey := CreateRecordKey(key)
	if recordKey == nil {
		fmt.Printf("Error while creating record key: %v", key)
		return nil, nil
	}
	record, err := client.GetHeader(readPolicy, recordKey)

	return record, err
}

func ReadRecord(key string) (*aerospike.Record, error) {
	recordKey := CreateRecordKey(key)
	if recordKey == nil {
		fmt.Printf("Error while creating record key: %v", key)
		return nil, nil
	}
	record, err := client.Get(readPolicy, recordKey)

	return record, err
}

func deleteRecord(key string) bool {
	recordKey := CreateRecordKey(key)
	if recordKey == nil {
		fmt.Printf("Error while creating record key: %v", key)
	}
	existed, err := client.Delete(deletePolicy, recordKey)
	if err != nil {
		fmt.Printf("Error while deleting record: %v", err)
	}
	fmt.Printf("Record existed: %v", existed)
	return existed
}
