package views

import (
	"testing"

	"github.com/derailed/k9s/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestAliasView(t *testing.T) {
	v := newAliasView(NewApp(config.NewConfig(ks{})), nil)
	td := v.hydrate()
	v.init(nil, "")

	assert.Equal(t, 3, len(td.Header))
	assert.Equal(t, 31, len(td.Rows))
	assert.Equal(t, "Aliases", v.getTitle())
}
