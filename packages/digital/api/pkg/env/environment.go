package environment

import (
	_ "embed"
	"encoding/json"
	embed "github.com/cyberpunkattack"
	"os"
)

type Environment interface {
	GetVariable(key string) string
	SetVariable(key, value string) error
}

var SingletonEnvHandler *Handler

type Handler struct {
	TookVariables map[string]string
}

func parseFile() (map[string]any, error) {
	unmarshalledDict := map[string]any{}
//	cwd, err := os.Getwd()
//	if err != nil {
//		return unmarshalledDict, errors.New("roflan.io error! on parseFile getWd")
//	}
//	absolutePath := path.Join(path.Dir(cwd), relativePath)
//	file, err := os.Open("")
//	if err != nil {
//		return unmarshalledDict, errors.New(fmt.Sprintf("roflan.io error! on os.Open, path: %s", absolutePath))
//	}
//	defer file.Close()
//	byteJson, err := io.ReadAll(file)
//	if err != nil {
//		return unmarshalledDict, errors.New(fmt.Sprintf("roflan.io error! on ioutils.ReadAll, path: %s", absolutePath))
//	}
	_ = json.Unmarshal(embed.EmbedEnvConfig, &unmarshalledDict)
	return unmarshalledDict, nil
}

func (env *Handler) GetVariable(key string) string {
	if env.TookVariables[key] == "" {
		return os.Getenv(key)
	} else {
		return env.TookVariables[key]
	}
}

func (env *Handler) SetVariable(key, value string) error {
	err := os.Setenv(key, value)
	env.TookVariables[key] = value
	return err
}

func InitEnvironment() {
	SingletonEnvHandler = &Handler{
		TookVariables: map[string]string{},
	}
	dict, _ := parseFile()

	for key, value := range dict {
		stringRepresentation, isOk := value.(string)
		if isOk {
			_ = os.Setenv(key, stringRepresentation)
			SingletonEnvHandler.TookVariables[key] = stringRepresentation
		}
	}

}

func GEnv() *Handler {
	if SingletonEnvHandler == nil {
		SingletonEnvHandler = new(Handler)
	}
	return SingletonEnvHandler
}
