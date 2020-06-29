# go-ssm [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) [![Go Report Card](https://goreportcard.com/badge/github.com/pradeeprajagopal/go-ssm)](https://goreportcard.com/report/github.com/pradeeprajagopal/go-ssm) [![Generic badge](https://img.shields.io/badge/coverage-90-green.svg)](https://shields.io/)


Package ssm provides the client and types for making API requests to Amazon Simple Systems Manager (SSM). 

    import "github.com/pradeeprajagopal/go-ssm"
#### *What is SSM?*

AWS Systems Manager Parameter Store provides secure, hierarchical storage for configuration data management and secrets management

#### *How to use?*

#### 1.Get Parameter value:
```go
package main
import (
    "github.com/pradeeprajagopal/go-ssm"
    "log"
)

func main() {
	svc := ssm.New()
	result,err := svc.Param("ParameterName", true).GetValue()
	if err != nil {
		log.Println(err)
	}
	log.Println(result)
}
```

OUTPUT:
`MyParameterOutput`

#### 2. Get Multiple parameter Value:

```go
	value, _ := ssmsvc.Params(names, true).GetValues()
	log.Println(value)
```

