package auth

import "github.com/google/uuid"

type Auth struct {
	ID                 int64     `pg:"id,pk"`
	ClientID           uuid.UUID `pg:"client_id, type:uuid,default:uuid_generate_v4()"`
	ClientSecret       uuid.UUID `pg:"client_secret, type:uuid,default:uuid_generate_v4()"`
	SignatureAlgorithm string    `pg:"signature_algorithm,default:RS256"`
	Description        string    `pg:"description"`
}
