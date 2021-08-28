package auth

import "github.com/google/uuid"

type AuthModel struct {
	ClientID     uuid.UUID `pg:"type:uuid,default:uuid_generate_v4()"`
	ClientSecret uuid.UUID `pg:"type:uuid,default:uuid_generate_v4()"`
	SignatureAlg string    `pg:"signature_algorith,default:RS256"`
	Description  string    `pg:"description"`
}
