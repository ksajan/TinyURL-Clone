package aerospike

import (
	aerospike "github.com/aerospike/aerospike-client-go/v6"
)

var writePolicy = CreateWritePolicy()
var readPolicy = CreateReadPolicy()
var deletePolicy = CreateDeletePolicy()

func init() {
	client, _ = GetAerospikeClient()
}

func CreateWritePolicy() *aerospike.WritePolicy {
	policy := aerospike.NewWritePolicy(0, 0)
	policy.RecordExistsAction = aerospike.UPDATE
	policy.SendKey = true
	return policy
}

func CreateReadPolicy() *aerospike.BasePolicy {
	policy := aerospike.NewPolicy()
	policy.SocketTimeout = 300
	return policy
}

func CreateDeletePolicy() *aerospike.WritePolicy {
	deletePolicy := aerospike.NewWritePolicy(0, 0)
	deletePolicy.Expiration = 24 * 60 * 60
	return deletePolicy
}
