package swagger

type Swagger struct {
	SwaggerVersion      string                `json:"swagger,omitempty" yaml:"swagger,omitempty"`
	Infos               Information           `json:"info" yaml:"info"`
	Host                string                `json:"host,omitempty" yaml:"host,omitempty"`
	BasePath            string                `json:"basePath,omitempty" yaml:"basePath,omitempty"`
	Schemes             []string              `json:"schemes,omitempty" yaml:"schemes,omitempty"`
	Consumes            []string              `json:"consumes,omitempty" yaml:"consumes,omitempty"`
	Produces            []string              `json:"produces,omitempty" yaml:"produces,omitempty"`
	Paths               map[string]*Item      `json:"paths" yaml:"paths"`
	Definitions         map[string]Schema     `json:"definitions,omitempty" yaml:"definitions,omitempty"`
	SecurityDefinitions map[string]Security   `json:"securityDefinitions,omitempty" yaml:"securityDefinitions,omitempty"`
	Security            []map[string][]string `json:"security,omitempty" yaml:"security,omitempty"`
	Tags                []Tag                 `json:"tags,omitempty" yaml:"tags,omitempty"`
	ExternalDocs        *ExternalDocs         `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}

type Information struct {
	Title          string `json:"title,omitempty" yaml:"title,omitempty"`
	Description    string `json:"description,omitempty" yaml:"description,omitempty"`
	Version        string `json:"version,omitempty" yaml:"version,omitempty"`
	TermsOfService string `json:"termsOfService,omitempty" yaml:"termsOfService,omitempty"`

	Contact Contact  `json:"contact,omitempty" yaml:"contact,omitempty"`
	License *License `json:"license,omitempty" yaml:"license,omitempty"`
}

type Contact struct {
	Name  string `json:"name,omitempty" yaml:"name,omitempty"`
	URL   string `json:"url,omitempty" yaml:"url,omitempty"`
	EMail string `json:"email,omitempty" yaml:"email,omitempty"`
}

type License struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	URL  string `json:"url,omitempty" yaml:"url,omitempty"`
}

type Item struct {
	Ref     string     `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Get     *Operation `json:"get,omitempty" yaml:"get,omitempty"`
	Put     *Operation `json:"put,omitempty" yaml:"put,omitempty"`
	Post    *Operation `json:"post,omitempty" yaml:"post,omitempty"`
	Delete  *Operation `json:"delete,omitempty" yaml:"delete,omitempty"`
	Options *Operation `json:"options,omitempty" yaml:"options,omitempty"`
	Head    *Operation `json:"head,omitempty" yaml:"head,omitempty"`
	Patch   *Operation `json:"patch,omitempty" yaml:"patch,omitempty"`
}

type Operation struct {
	Tags        []string              `json:"tags,omitempty" yaml:"tags,omitempty"`
	Summary     string                `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string                `json:"description,omitempty" yaml:"description,omitempty"`
	OperationID string                `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Consumes    []string              `json:"consumes,omitempty" yaml:"consumes,omitempty"`
	Produces    []string              `json:"produces,omitempty" yaml:"produces,omitempty"`
	Schemes     []string              `json:"schemes,omitempty" yaml:"schemes,omitempty"`
	Parameters  []Parameter           `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Responses   map[string]Response   `json:"responses,omitempty" yaml:"responses,omitempty"`
	Security    []map[string][]string `json:"security,omitempty" yaml:"security,omitempty"`
	Deprecated  bool                  `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
}

type Parameter struct {
	In          string          `json:"in,omitempty" yaml:"in,omitempty"`
	Name        string          `json:"name,omitempty" yaml:"name,omitempty"`
	Description string          `json:"description,omitempty" yaml:"description,omitempty"`
	Required    bool            `json:"required,omitempty" yaml:"required,omitempty"`
	Schema      *Schema         `json:"schema,omitempty" yaml:"schema,omitempty"`
	Type        string          `json:"type,omitempty" yaml:"type,omitempty"`
	Format      string          `json:"format,omitempty" yaml:"format,omitempty"`
	Items       *ParameterItems `json:"items,omitempty" yaml:"items,omitempty"`
	Default     interface{}     `json:"default,omitempty" yaml:"default,omitempty"`
}

type ParameterItems struct {
	Type             string            `json:"type,omitempty" yaml:"type,omitempty"`
	Format           string            `json:"format,omitempty" yaml:"format,omitempty"`
	Items            []*ParameterItems `json:"items,omitempty" yaml:"items,omitempty"`
	CollectionFormat string            `json:"collectionFormat,omitempty" yaml:"collectionFormat,omitempty"`
	Default          string            `json:"default,omitempty" yaml:"default,omitempty"`
}

type Schema struct {
	Ref         string               `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Title       string               `json:"title,omitempty" yaml:"title,omitempty"`
	Format      string               `json:"format,omitempty" yaml:"format,omitempty"`
	Description string               `json:"description,omitempty" yaml:"description,omitempty"`
	Required    []string             `json:"required,omitempty" yaml:"required,omitempty"`
	Type        string               `json:"type,omitempty" yaml:"type,omitempty"`
	Items       *Schema              `json:"items,omitempty" yaml:"items,omitempty"`
	Properties  map[string]Propertie `json:"properties,omitempty" yaml:"properties,omitempty"`
	Enum        []interface{}        `json:"enum,omitempty" yaml:"enum,omitempty"`
	Example     interface{}          `json:"example,omitempty" yaml:"example,omitempty"`
}

type Propertie struct {
	Ref                  string               `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Title                string               `json:"title,omitempty" yaml:"title,omitempty"`
	Description          string               `json:"description,omitempty" yaml:"description,omitempty"`
	Default              interface{}          `json:"default,omitempty" yaml:"default,omitempty"`
	Type                 string               `json:"type,omitempty" yaml:"type,omitempty"`
	Example              interface{}          `json:"example,omitempty" yaml:"example,omitempty"`
	Required             []string             `json:"required,omitempty" yaml:"required,omitempty"`
	Format               string               `json:"format,omitempty" yaml:"format,omitempty"`
	ReadOnly             bool                 `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`
	Properties           map[string]Propertie `json:"properties,omitempty" yaml:"properties,omitempty"`
	Items                *Propertie           `json:"items,omitempty" yaml:"items,omitempty"`
	AdditionalProperties *Propertie           `json:"additionalProperties,omitempty" yaml:"additionalProperties,omitempty"`
}

type Response struct {
	Description string  `json:"description" yaml:"description"`
	Schema      *Schema `json:"schema,omitempty" yaml:"schema,omitempty"`
	Ref         string  `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}

type Security struct {
	Type             string            `json:"type,omitempty" yaml:"type,omitempty"`
	Description      string            `json:"description,omitempty" yaml:"description,omitempty"`
	Name             string            `json:"name,omitempty" yaml:"name,omitempty"`
	In               string            `json:"in,omitempty" yaml:"in,omitempty"`
	Flow             string            `json:"flow,omitempty" yaml:"flow,omitempty"`
	AuthorizationURL string            `json:"authorizationUrl,omitempty" yaml:"authorizationUrl,omitempty"`
	TokenURL         string            `json:"tokenUrl,omitempty" yaml:"tokenUrl,omitempty"`
	Scopes           map[string]string `json:"scopes,omitempty" yaml:"scopes,omitempty"`
}

type Tag struct {
	Name         string        `json:"name,omitempty" yaml:"name,omitempty"`
	Description  string        `json:"description,omitempty" yaml:"description,omitempty"`
	ExternalDocs *ExternalDocs `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}

type ExternalDocs struct {
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	URL         string `json:"url,omitempty" yaml:"url,omitempty"`
}
