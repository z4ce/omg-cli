package vault 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type S3 struct {

	/*AccessKey - Descr: AWS access key Default: <nil>
*/
	AccessKey interface{} `yaml:"access_key,omitempty"`

	/*Endpoint - Descr: AWS S3 endpoint Default: <nil>
*/
	Endpoint interface{} `yaml:"endpoint,omitempty"`

	/*Bucket - Descr: S3 bucket name Default: <nil>
*/
	Bucket interface{} `yaml:"bucket,omitempty"`

	/*SessionToken - Descr: AWS session token Default: <nil>
*/
	SessionToken interface{} `yaml:"session_token,omitempty"`

	/*SecretKey - Descr: AWS secret key Default: <nil>
*/
	SecretKey interface{} `yaml:"secret_key,omitempty"`

	/*Region - Descr: AWS region Default: us-east-1
*/
	Region interface{} `yaml:"region,omitempty"`

}