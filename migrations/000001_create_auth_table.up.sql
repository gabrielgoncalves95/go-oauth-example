CREATE TABLE public.auth
(
    id                  SERIAL PRIMARY KEY,
    client_id           uuid                   NOT NULL,
    client_secret       uuid                   NOT NULL,
    description         character varying(255) NULL,
    signature_algorithm character varying(255) NOT NULL DEFAULT 'RS256'
);

CREATE INDEX idx_client_id
    ON public.auth (client_id);


CREATE INDEX idx_secret_id
    ON public.auth (client_secret);