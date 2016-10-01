package jiradata

/////////////////////////////////////////////////////////////////////////
// This Code is Generated by SlipScheme Project:
// https://github.com/coryb/slipscheme
//
// Generated with command:
// slipscheme -pkg jiradata -overwrite ../schemas/SearchResults.json
/////////////////////////////////////////////////////////////////////////
//                            DO NOT EDIT                              //
/////////////////////////////////////////////////////////////////////////

// JSONType defined from schema:
// {
//   "title": "Json Type",
//   "type": "object",
//   "properties": {
//     "custom": {
//       "type": "string"
//     },
//     "customId": {
//       "type": "integer"
//     },
//     "items": {
//       "type": "string"
//     },
//     "system": {
//       "type": "string"
//     },
//     "type": {
//       "type": "string"
//     }
//   }
// }
type JSONType struct {
	Custom   string `json:"custom,omitempty" yaml:"custom,omitempty"`
	CustomID int    `json:"customId,omitempty" yaml:"customId,omitempty"`
	Items    string `json:"items,omitempty" yaml:"items,omitempty"`
	System   string `json:"system,omitempty" yaml:"system,omitempty"`
	Type     string `json:"type,omitempty" yaml:"type,omitempty"`
}