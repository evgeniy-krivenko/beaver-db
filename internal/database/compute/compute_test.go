package compute_test

import (
	"io"
	"log/slog"
	"testing"

	"beaver/internal/database/compute"
	"github.com/stretchr/testify/require"
)

func Test_Query(t *testing.T) {
	dummyLogger := slog.New(slog.NewTextHandler(io.Discard, nil))

	type err struct {
		occured bool
		err     error
	}

	comp := compute.New(dummyLogger)
	cases := []struct {
		name  string
		q     string
		err   err
		query compute.Query
	}{
		{
			name: "invalid query when empty tokens",
			q:    "",
			err: err{
				occured: true,
				err:     compute.ErrInvalidQuery,
			},
		},
		{
			name: "unknown command err",
			q:    "METHOD",
			err: err{
				occured: true,
				err:     compute.ErrUnknownCommand,
			},
		},
		{
			name: "wrong number of arguments for get cmd",
			q:    "GET key value",
			err: err{
				occured: true,
				err:     compute.ErrWrongNumArgs,
			},
		},
		{
			name: "wrong number of arguments for set cmd",
			q:    "SET key",
			err: err{
				occured: true,
				err:     compute.ErrWrongNumArgs,
			},
		},
		{
			name: "wrong number of arguments for del cmd",
			q:    "DEL key value",
			err: err{
				occured: true,
				err:     compute.ErrWrongNumArgs,
			},
		},
		{
			name:  "success to parse set cmd",
			q:     "SET key value",
			query: compute.NewQuery(compute.SetCommandID, []string{"key", "value"}),
		},
		{
			name:  "success to parse get cmd",
			q:     "GET somekey",
			query: compute.NewQuery(compute.GetCommandID, []string{"somekey"}),
		},
		{
			name:  "success to parse del cmd",
			q:     "DEL otherkey",
			query: compute.NewQuery(compute.DelCommandID, []string{"otherkey"}),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			query, err := comp.Parse(c.q)
			if c.err.occured {
				require.Error(t, err)
				require.Empty(t, query)
			} else {
				require.NoError(t, err)
				require.Equal(t, c.query.CommandID(), query.CommandID())
				require.Equal(t, c.query.Args(), query.Args())
			}
		})
	}
}
