package embed

import (
	_ "embed"
)

//go:embed internal/config/.env.development.json
var EmbedEnvConfig []byte

//go:embed internal/config/cards.actions.json
var EmbedActionsDeck []byte

//go:embed internal/config/cards.implants.json
var EmbedImplantsDeck []byte

//go:embed internal/config/actions.map.json
var EmbedActionsMap []byte

//go:embed internal/config/implants.map.json
var EmbedImplantsMap []byte

//go:embed internal/config/actions.counter.json
var EmbedActionsCounter []byte

//go:embed internal/config/implants.counter.json
var EmbedImplantsCounter []byte

//go:embed internal/config/roles.counter.json
var EmbedRolesCounter []byte

//go:embed internal/config/roles.map.json
var EmbedRolesMap []byte
