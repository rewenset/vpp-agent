// Code generated by adapter-generator. DO NOT EDIT.

package adapter

import (
	"github.com/golang/protobuf/proto"
	. "github.com/ligato/vpp-agent/plugins/kvscheduler/api"
	"github.com/ligato/vpp-agent/api/models/vpp/nat"
)

////////// type-safe key-value pair with metadata //////////

type DNAT44KVWithMetadata struct {
	Key      string
	Value    *vpp_nat.DNat44
	Metadata interface{}
	Origin   ValueOrigin
}

////////// type-safe Descriptor structure //////////

type DNAT44Descriptor struct {
	Name                 string
	KeySelector          KeySelector
	ValueTypeName        string
	KeyLabel             func(key string) string
	ValueComparator      func(key string, oldValue, newValue *vpp_nat.DNat44) bool
	NBKeyPrefix          string
	WithMetadata         bool
	MetadataMapFactory   MetadataMapFactory
	Validate             func(key string, value *vpp_nat.DNat44) error
	Create               func(key string, value *vpp_nat.DNat44) (metadata interface{}, err error)
	Delete               func(key string, value *vpp_nat.DNat44, metadata interface{}) error
	Update               func(key string, oldValue, newValue *vpp_nat.DNat44, oldMetadata interface{}) (newMetadata interface{}, err error)
	UpdateWithRecreate   func(key string, oldValue, newValue *vpp_nat.DNat44, metadata interface{}) bool
	Retrieve             func(correlate []DNAT44KVWithMetadata) ([]DNAT44KVWithMetadata, error)
	IsRetriableFailure   func(err error) bool
	DerivedValues        func(key string, value *vpp_nat.DNat44) []KeyValuePair
	Dependencies         func(key string, value *vpp_nat.DNat44) []Dependency
	RetrieveDependencies []string /* descriptor name */
}

////////// Descriptor adapter //////////

type DNAT44DescriptorAdapter struct {
	descriptor *DNAT44Descriptor
}

func NewDNAT44Descriptor(typedDescriptor *DNAT44Descriptor) *KVDescriptor {
	adapter := &DNAT44DescriptorAdapter{descriptor: typedDescriptor}
	descriptor := &KVDescriptor{
		Name:                 typedDescriptor.Name,
		KeySelector:          typedDescriptor.KeySelector,
		ValueTypeName:        typedDescriptor.ValueTypeName,
		KeyLabel:             typedDescriptor.KeyLabel,
		NBKeyPrefix:          typedDescriptor.NBKeyPrefix,
		WithMetadata:         typedDescriptor.WithMetadata,
		MetadataMapFactory:   typedDescriptor.MetadataMapFactory,
		IsRetriableFailure:   typedDescriptor.IsRetriableFailure,
		RetrieveDependencies: typedDescriptor.RetrieveDependencies,
	}
	if typedDescriptor.ValueComparator != nil {
		descriptor.ValueComparator = adapter.ValueComparator
	}
	if typedDescriptor.Validate != nil {
		descriptor.Validate = adapter.Validate
	}
	if typedDescriptor.Create != nil {
		descriptor.Create = adapter.Create
	}
	if typedDescriptor.Delete != nil {
		descriptor.Delete = adapter.Delete
	}
	if typedDescriptor.Update != nil {
		descriptor.Update = adapter.Update
	}
	if typedDescriptor.UpdateWithRecreate != nil {
		descriptor.UpdateWithRecreate = adapter.UpdateWithRecreate
	}
	if typedDescriptor.Retrieve != nil {
		descriptor.Retrieve = adapter.Retrieve
	}
	if typedDescriptor.Dependencies != nil {
		descriptor.Dependencies = adapter.Dependencies
	}
	if typedDescriptor.DerivedValues != nil {
		descriptor.DerivedValues = adapter.DerivedValues
	}
	return descriptor
}

func (da *DNAT44DescriptorAdapter) ValueComparator(key string, oldValue, newValue proto.Message) bool {
	typedOldValue, err1 := castDNAT44Value(key, oldValue)
	typedNewValue, err2 := castDNAT44Value(key, newValue)
	if err1 != nil || err2 != nil {
		return false
	}
	return da.descriptor.ValueComparator(key, typedOldValue, typedNewValue)
}

func (da *DNAT44DescriptorAdapter) Validate(key string, value proto.Message) (err error) {
	typedValue, err := castDNAT44Value(key, value)
	if err != nil {
		return err
	}
	return da.descriptor.Validate(key, typedValue)
}

func (da *DNAT44DescriptorAdapter) Create(key string, value proto.Message) (metadata Metadata, err error) {
	typedValue, err := castDNAT44Value(key, value)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Create(key, typedValue)
}

func (da *DNAT44DescriptorAdapter) Update(key string, oldValue, newValue proto.Message, oldMetadata Metadata) (newMetadata Metadata, err error) {
	oldTypedValue, err := castDNAT44Value(key, oldValue)
	if err != nil {
		return nil, err
	}
	newTypedValue, err := castDNAT44Value(key, newValue)
	if err != nil {
		return nil, err
	}
	typedOldMetadata, err := castDNAT44Metadata(key, oldMetadata)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Update(key, oldTypedValue, newTypedValue, typedOldMetadata)
}

func (da *DNAT44DescriptorAdapter) Delete(key string, value proto.Message, metadata Metadata) error {
	typedValue, err := castDNAT44Value(key, value)
	if err != nil {
		return err
	}
	typedMetadata, err := castDNAT44Metadata(key, metadata)
	if err != nil {
		return err
	}
	return da.descriptor.Delete(key, typedValue, typedMetadata)
}

func (da *DNAT44DescriptorAdapter) UpdateWithRecreate(key string, oldValue, newValue proto.Message, metadata Metadata) bool {
	oldTypedValue, err := castDNAT44Value(key, oldValue)
	if err != nil {
		return true
	}
	newTypedValue, err := castDNAT44Value(key, newValue)
	if err != nil {
		return true
	}
	typedMetadata, err := castDNAT44Metadata(key, metadata)
	if err != nil {
		return true
	}
	return da.descriptor.UpdateWithRecreate(key, oldTypedValue, newTypedValue, typedMetadata)
}

func (da *DNAT44DescriptorAdapter) Retrieve(correlate []KVWithMetadata) ([]KVWithMetadata, error) {
	var correlateWithType []DNAT44KVWithMetadata
	for _, kvpair := range correlate {
		typedValue, err := castDNAT44Value(kvpair.Key, kvpair.Value)
		if err != nil {
			continue
		}
		typedMetadata, err := castDNAT44Metadata(kvpair.Key, kvpair.Metadata)
		if err != nil {
			continue
		}
		correlateWithType = append(correlateWithType,
			DNAT44KVWithMetadata{
				Key:      kvpair.Key,
				Value:    typedValue,
				Metadata: typedMetadata,
				Origin:   kvpair.Origin,
			})
	}

	typedValues, err := da.descriptor.Retrieve(correlateWithType)
	if err != nil {
		return nil, err
	}
	var values []KVWithMetadata
	for _, typedKVWithMetadata := range typedValues {
		kvWithMetadata := KVWithMetadata{
			Key:      typedKVWithMetadata.Key,
			Metadata: typedKVWithMetadata.Metadata,
			Origin:   typedKVWithMetadata.Origin,
		}
		kvWithMetadata.Value = typedKVWithMetadata.Value
		values = append(values, kvWithMetadata)
	}
	return values, err
}

func (da *DNAT44DescriptorAdapter) DerivedValues(key string, value proto.Message) []KeyValuePair {
	typedValue, err := castDNAT44Value(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.DerivedValues(key, typedValue)
}

func (da *DNAT44DescriptorAdapter) Dependencies(key string, value proto.Message) []Dependency {
	typedValue, err := castDNAT44Value(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.Dependencies(key, typedValue)
}

////////// Helper methods //////////

func castDNAT44Value(key string, value proto.Message) (*vpp_nat.DNat44, error) {
	typedValue, ok := value.(*vpp_nat.DNat44)
	if !ok {
		return nil, ErrInvalidValueType(key, value)
	}
	return typedValue, nil
}

func castDNAT44Metadata(key string, metadata Metadata) (interface{}, error) {
	if metadata == nil {
		return nil, nil
	}
	typedMetadata, ok := metadata.(interface{})
	if !ok {
		return nil, ErrInvalidMetadataType(key)
	}
	return typedMetadata, nil
}
