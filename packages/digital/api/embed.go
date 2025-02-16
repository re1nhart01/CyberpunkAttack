package embed

import _ "embed"

//go:embed internal/config/.env.development.json
var EmbedEnvConfig []byte

//go:embed internal/config/cards.actions.json
var EmbedActionsDeck []byte

//go:embed internal/config/cards.implants.json
var EmbedImplantsDeck []byte

//go:embed internal/config/cards.cyberviruses.json
var EmbedCybervirusesDeck []byte
