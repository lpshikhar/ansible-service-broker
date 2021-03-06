package ansibleapp

import (
	"encoding/json"
	"github.com/op/go-logging"
	"github.com/pborman/uuid"
	"gopkg.in/yaml.v2"
)

type Parameters map[string]interface{}
type SpecManifest map[string]*Spec

var AnsibleAppSpecLabel = "com.redhat.ansibleapp.spec"

type ImageData struct {
	Name         string
	Tag          string
	Labels       map[string]string
	Layers       []string
	IsAnsibleApp bool
	Error        error
}

type ParameterDescriptor struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Type        string      `json:"type"`
	Required    bool        `json:"required"`
	Default     interface{} `json:"default"`
}

type Spec struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Bindable    bool   `json:"bindable"`
	Description string `json:"description"`

	// required, optional, unsupported
	Async      string                 `json:"async"`
	Parameters []*ParameterDescriptor `json:"parameters"`
}

func specLogDump(spec *Spec, log *logging.Logger) {
	log.Debug("============================================================")
	log.Debug("Spec: %s", spec.Id)
	log.Debug("============================================================")
	log.Debug("Name: %s", spec.Name)
	log.Debug("Bindable: %t", spec.Bindable)
	log.Debug("Description: %s", spec.Description)
	log.Debug("Async: %s", spec.Async)

	for _, param := range spec.Parameters {
		log.Debug("ParameterDescriptor")
		log.Debug("  Name: %s", param.Name)
		log.Debug("  Description: %s", param.Description)
		log.Debug("  Type: %s", param.Type)
		log.Debug("  Required: %t", param.Required)
		log.Debug("  Default: %s", param.Name)
	}
}

func specsLogDump(specs []*Spec, log *logging.Logger) {
	for _, spec := range specs {
		specLogDump(spec, log)
	}
}

func NewSpecManifest(specs []*Spec) SpecManifest {
	manifest := make(map[string]*Spec)
	for _, spec := range specs {
		manifest[spec.Id] = spec
	}
	return manifest
}

type ServiceInstance struct {
	Id         uuid.UUID   `json:"id"`
	Spec       *Spec       `json:"spec"`
	Parameters *Parameters `json:"parameters"`
}

func LoadJSON(payload string, obj interface{}) error {
	err := json.Unmarshal([]byte(payload), obj)
	if err != nil {
		return err
	}

	return nil
}

func DumpJSON(obj interface{}) (string, error) {
	payload, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	return string(payload), nil
}

func LoadYAML(payload string, obj interface{}) error {
	var err error

	if err = yaml.Unmarshal([]byte(payload), obj); err != nil {
		return err
	}

	return nil
}
