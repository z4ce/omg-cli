package director 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type UserManagement struct {

	/*Uaa - Descr: Symmetric key to verify Uaa token Default: <nil>
*/
	Uaa *Uaa `yaml:"uaa,omitempty"`

	/*Local - Descr: List of users that can authenticate with director in non-Uaa mode Default: <nil>
*/
	Local *Local `yaml:"local,omitempty"`

	/*Provider - Descr: User management implementation (local|uaa) Default: local
*/
	Provider interface{} `yaml:"provider,omitempty"`

}