package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindByPayloadCID(t *testing.T) {
	client, err := New("https://cid.contact")
	assert.NoError(t, err)

	cid := "bafybeiaasmuisjhbgd6fj5af2gs6t42cvkdkm77njkjea2twdvvpsnvmwu"
	ctx := context.Background()
	result, err := client.FindByPayloadCID(ctx, cid)
	assert.NoError(t, err)
	fmt.Println(result)
}
