# go-ssm
Package ssm provides the client and types for making API requests to Amazon Simple Systems Manager (SSM). 

    import "githubdev.dco.elmae/ReleaseEngineering/verita-core/data/ssm"
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

