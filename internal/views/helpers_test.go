package views

import (
	"testing"

	"github.com/derailed/k9s/internal/config"
	"github.com/derailed/k9s/internal/resource"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func init() {
	zerolog.SetGlobalLevel(zerolog.Disabled)
}

func TestIsTCPPort(t *testing.T) {
	uu := map[string]struct {
		p string
		e bool
	}{
		"tcp": {"80╱TCP", true},
		"udp": {"80╱UDP", false},
	}

	for k, u := range uu {
		t.Run(k, func(t *testing.T) {
			assert.Equal(t, u.e, isTCPPort(u.p))
		})
	}
}

func TestFQN(t *testing.T) {
	uu := map[string]struct {
		ns, n, e string
	}{
		"fullFQN": {"blee", "fred", "blee/fred"},
		"allNS":   {"", "fred", "fred"},
	}

	for k, u := range uu {
		t.Run(k, func(t *testing.T) {
			assert.Equal(t, u.e, fqn(u.ns, u.n))
		})
	}
}

func TestDeltas(t *testing.T) {
	uu := []struct {
		s1, s2, e string
	}{
		{"", "", ""},
		{resource.MissingValue, "", deltaSign},
		{resource.NAValue, "", ""},
		{"fred", "fred", ""},
		{"fred", "blee", deltaSign},
		{"1", "1", ""},
		{"1", "2", plusSign},
		{"2", "1", minusSign},
		{"2m33s", "2m33s", ""},
		{"2m33s", "1m", minusSign},
		{"33s", "1m", plusSign},
		{"10Gi", "10Gi", ""},
		{"10Gi", "20Gi", plusSign},
		{"30Gi", "20Gi", minusSign},
		{"15%", "15%", ""},
		{"20%", "40%", plusSign},
		{"5%", "2%", minusSign},
	}

	for _, u := range uu {
		assert.Equal(t, u.e, deltas(u.s1, u.s2))
	}
}

func TestUrlFor(t *testing.T) {
	uu := map[string]struct {
		cfg      config.BenchConfig
		co, port string
		e        string
	}{
		"empty": {
			config.BenchConfig{}, "c1", "9000", "http://localhost:9000/",
		},
		"path": {
			config.BenchConfig{
				HTTP: config.HTTP{
					Path: "/fred/blee",
				},
			},
			"c1",
			"9000",
			"http://localhost:9000/fred/blee",
		},
		"host/path": {
			config.BenchConfig{
				HTTP: config.HTTP{
					Host: "zorg",
					Path: "/fred/blee",
				},
			},
			"c1",
			"9000",
			"http://zorg:9000/fred/blee",
		},
	}

	for k, u := range uu {
		t.Run(k, func(t *testing.T) {
			assert.Equal(t, u.e, urlFor(u.cfg, u.co, u.port))
		})
	}
}

func TestContainerID(t *testing.T) {
	uu := map[string]struct {
		path, co string
		e        string
	}{
		"plain": {
			"fred/blee", "c1", "fred/blee:c1",
		},
		"podID": {
			"fred/blee-78f8b5d78c-f8588", "c1", "fred/blee:c1",
		},
		"stsID": {
			"fred/blee-1", "c1", "fred/blee:c1",
		},
	}

	for k, u := range uu {
		t.Run(k, func(t *testing.T) {
			assert.Equal(t, u.e, containerID(u.path, u.co))
		})
	}
}

func TestStripPort(t *testing.T) {
	uu := map[string]struct {
		port, e string
	}{
		"full": {
			"fred:8000", "8000",
		},
		"port": {
			"8000", "8000",
		},
		"protocol": {
			"dns:53╱UDP", "53",
		},
	}

	for k, u := range uu {
		t.Run(k, func(t *testing.T) {
			assert.Equal(t, u.e, stripPort(u.port))
		})
	}
}
