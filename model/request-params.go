package model

// swagger:parameters AddConfig
type RequestConfigBody struct {
	// - name: body
	//  in: body
	//  description: name and status
	//  schema:
	//  type: object
	//     "$ref": "#/definitions/Config"
	//  required: true
	Body Config `json:"body"`
}

// swagger:parameters GetOneConfig DeleteOneConfig
type StructId struct {
	// -name: Id
	// in: path
	//  required: true
	Id string `json:"id"`
}

// swagger:parameters GetOneConfig DeleteOneConfig
type StructVersion struct {
	// -name: Version
	// in: path
	//  required: true
	Version string `json:"version"`
}

// swagger:parameters AddConfigGroup
type RequestConfigGroupBody struct {
	// - name: body
	//  in: body
	//  description: name and status
	//  schema:
	//  type: object
	//     "$ref": "#/definitions/ConfigGroup"
	//  required: true
	Body ConfigGroup `json:"body"`
}

// swagger:parameters GetOneConfigGroup RemoveConfigGroup GetAllConfigsInGroupByLabel
type GroupId struct {
	// -name: Id
	// in: path
	//  required: true
	Id string `json:"id"`
}

// swagger:parameters GetOneConfigGroup RemoveConfigGroup GetAllConfigsInGroupByLabel
type GroupVersion struct {
	// -name: Version
	// in: path
	//  required: true
	Version string `json:"version"`
}

// swagger:parameters GetAllConfigsInGroupByLabel
type GroupLabel struct {
	// -name: Label
	// in: path
	// required: true
	Label string `json:"label"`
}
