package upcloud

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestUnmarshalPlans tests that Plans and Plan objects unmarshal correctly
func TestUnmarshalPlans(t *testing.T) {
	originalXML := `<?xml version="1.0" encoding="utf-8"?>
<plans>
    <plan>
        <core_number>1</core_number>
        <memory_amount>1024</memory_amount>
        <name>1xCPU-1GB</name>
        <public_traffic_out>1024</public_traffic_out>
        <storage_size>25</storage_size>
        <storage_tier>maxiops</storage_tier>
    </plan>
    <plan>
        <core_number>1</core_number>
        <memory_amount>2048</memory_amount>
        <name>1xCPU-2GB</name>
        <public_traffic_out>2048</public_traffic_out>
        <storage_size>50</storage_size>
        <storage_tier>maxiops</storage_tier>
    </plan>
    <plan>
        <core_number>2</core_number>
        <memory_amount>4096</memory_amount>
        <name>2xCPU-4GB</name>
        <public_traffic_out>4096</public_traffic_out>
        <storage_size>80</storage_size>
        <storage_tier>maxiops</storage_tier>
    </plan>
    <plan>
        <core_number>4</core_number>
        <memory_amount>8192</memory_amount>
        <name>4xCPU-8GB</name>
        <public_traffic_out>5120</public_traffic_out>
        <storage_size>160</storage_size>
        <storage_tier>maxiops</storage_tier>
    </plan>
    <plan>
        <core_number>6</core_number>
        <memory_amount>16384</memory_amount>
        <name>6xCPU-16GB</name>
        <public_traffic_out>6144</public_traffic_out>
        <storage_size>320</storage_size>
        <storage_tier>maxiops</storage_tier>
    </plan>
    <plan>
        <core_number>8/core_number>
        <memory_amount>32768</memory_amount>
        <name>8xCPU-32GB</name>
        <public_traffic_out>7168</public_traffic_out>
        <storage_size>640</storage_size>
        <storage_tier>maxiops</storage_tier>
    </plan>
    <plan>
        <core_number>12</core_number>
        <memory_amount>49152</memory_amount>
        <name>12xCPU-48GB</name>
        <public_traffic_out>9216</public_traffic_out>
        <storage_size>960</storage_size>
        <storage_tier>maxiops</storage_tier>
    </plan>
    <plan>
        <core_number>16</core_number>
        <memory_amount>65536</memory_amount>
        <name>16xCPU-64GB</name>
        <public_traffic_out>10240</public_traffic_out>
        <storage_size>1280</storage_size>
        <storage_tier>maxiops</storage_tier>
    </plan>
    <plan>
        <core_number>20</core_number>
        <memory_amount>98304</memory_amount>
        <name>20xCPU-96GB</name>
        <public_traffic_out>12288</public_traffic_out>
        <storage_size>1920</storage_size>
        <storage_tier>maxiops</storage_tier>
    </plan>
    <plan>
        <core_number>20</core_number>
        <memory_amount>131072</memory_amount>
        <name>20xCPU-128GB</name>
        <public_traffic_out>24576</public_traffic_out>
        <storage_size>2048</storage_size>
        <storage_tier>maxiops</storage_tier>
    </plan>
</plans>`

	plans := Plans{}
	err := xml.Unmarshal([]byte(originalXML), &plans)
	assert.Nil(t, err)
	assert.Len(t, plans.Plans, 4)

	plan := plans.Plans[0]
	assert.Equal(t, 1, plan.CoreNumber)
	assert.Equal(t, 1024, plan.MemoryAmount)
	assert.Equal(t, "1xCPU-1GB", plan.Name)
	assert.Equal(t, 1024, plan.PublicTrafficOut)
	assert.Equal(t, 25, plan.StorageSize)
	assert.Equal(t, StorageTierMaxIOPS, plan.StorageTier)
}
