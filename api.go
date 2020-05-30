package ssm

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
	"log"
)

// SSM is a SSM API client.
type SSM struct {
	client ssmiface.SSMAPI
}

/*
Param input for querying a parameter from the store
*/
type Param struct {
	// Name of the parameter
	// example:  MyParameterKey
	// required: true
	Name string
	// WithDecryption is mandatory only if the value is SecureString
	// example:  MyParameterKey
	// required: false
	WithDecryption bool
	//Client
	svc *SSM
}

/*
Params input for querying multiple parameters from the store
*/
type Params struct {
	// Names of the parameters in the form of list
	// required: true
	Names []*string
	// WithDecryption is mandatory only if the value is SecureString
	// example:  MyParameterKey
	// required: false
	WithDecryption bool
	//Client
	svc *SSM
}

/*
sessions creates a new aws session using default credentials and returns a pointer of the session and an error
*/
func sessions() (*session.Session, error) {
	sess, err := session.NewSession()
	svc := session.Must(sess, err)
	return svc, err
}

/*New an exported function used to create a new SSM session with default credentials.*/
func New() *SSM {
	// Create AWS Session
	sess, err := sessions()
	if err != nil {
		log.Println(err)
		return nil
	}
	svc := &SSM{ssm.New(sess)}
	// Return SSM client
	return svc
}

/*Param creates the struct for querying the param store. Type receives a name of the parameter as string and whether the parameter decryption needs to be performed*/
func (s *SSM) Param(name string, decryption bool) *Param {
	return &Param{
		Name:           name,
		WithDecryption: decryption,
		svc:            s,
	}
}

/*Params creates the struct for querying one more items from the param store. Type receives a name of the parameter as list of string and whether the parameter decryption needs to be performed*/
func (s *SSM) Params(names []string, decryption bool) *Params {
	var Names []*string
	for i := range names {
		Names = append(Names, &names[i])
	}
	return &Params{
		Names:          Names,
		WithDecryption: decryption,
		svc:            s,
	}
}

/*GetValue type receives Param as receiver  and returns the value in the form of string and error*/
func (p *Param) GetValue() (string, error) {
	svc := p.svc.client
	parameter, err := svc.GetParameter(&ssm.GetParameterInput{
		Name:           &p.Name,
		WithDecryption: &p.WithDecryption,
	})
	if err != nil {
		return "", err
	}
	value := *parameter.Parameter.Value
	return value, nil
}

/*GetValues type receives Params as receiver  and returns the value in the form of list of string and error*/
func (p *Params) GetValues() ([]string, error) {
	var results []string
	svc := p.svc.client
	parameter, err := svc.GetParameters(&ssm.GetParametersInput{
		Names:          p.Names,
		WithDecryption: &p.WithDecryption,
	})
	if err != nil {
		return results, err
	}
	for i := range parameter.Parameters {
		results = append(results, *parameter.Parameters[i].Value)
	}
	return results, nil
}
