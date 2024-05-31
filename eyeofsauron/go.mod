module github.com/gandalf-network/gandalf-sdk-go/eyeofsauron

go 1.22.1

require (
	github.com/Khan/genqlient v0.7.0
	github.com/btcsuite/btcd/btcec/v2 v2.3.3
	github.com/google/uuid v1.6.0
	github.com/machinebox/graphql v0.2.2
)

require (
	github.com/agnivade/levenshtein v1.1.1 // indirect
	github.com/alexflint/go-arg v1.4.2 // indirect
	github.com/alexflint/go-scalar v1.0.0 // indirect
	github.com/bmatcuk/doublestar/v4 v4.6.1 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	github.com/matryer/is v1.4.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/vektah/gqlparser/v2 v2.5.11 // indirect
	golang.org/x/mod v0.15.0 // indirect
	golang.org/x/tools v0.18.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/machinebox/graphql => github.com/machinebox/graphql v0.2.3-0.20181106130121-3a9253180225

replace github.com/Khan/genqlient v0.7.0 => ./../../../gandalf-gate/genqlient
