package embed

import _ "embed"

//go:embed internal/config/.env.development.json
var EmbedEnvConfig []byte