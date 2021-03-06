package node

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCPUUsage(t *testing.T) {
	_, err := cpuUsage()
	assert.Equal(t, nil, err)
}

func TestCPUUsageSeparate(t *testing.T) {
	_, _, err := cpuUsageSeparated()
	assert.Equal(t, nil, err)
}
