package uaa 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Oauth struct {

	/*Providers - Descr: Contains a hash of OpenID Connect/Oauth Identity Providers, the key will be used as the origin key for that provider, followed by key/value pairs. Presence of the userInfoUrl will mark it as an OpenID provider instead of OAuth. Default: <nil>
*/
	Providers interface{} `yaml:"providers,omitempty"`

}