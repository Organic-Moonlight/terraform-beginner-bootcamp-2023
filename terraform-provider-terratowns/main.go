// package main: Declares the package name.
// The main package is special in Go, it's where the execution of the program starts.
package main

// import "fmt" Imports the fmt package, which contains functions for formatted I/O
// fmt is short for format
import (
	// "log"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

// func main(): Defines the main function, the entry point of the app.
// When you run the program, it starts executing from this function.
func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: Provider,
	})
	fmt.Print("Hello, World")
}

func Provider() *schema.Provider{
	var p *schema.Provider
	p = &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{

		},
		DataSourcesMap: map[string]*schema.Resource{

		},
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type: schema.TypeString,
				Required: true,
				Description: "The endpoint for the external service",
			},
			"token": {
				Type: schema.TypeString,
				Sensitive: true, // make the toke as sensitive to hide it in the logs
				Required: true,
				Description: "Bearer token for autorization",
			},
			"user_uuid": {
				Type: schema.TypeString,
				Required: true,
				Description: "UUID for configuration",
				// ValidateFunc: validateUUID,
			},
		},
	}
	// p.ConfigureContextFunc = providerConfigure(p)
	return p
}

// func validateUUID(v interface{}, k string)(ws []string, errors []error){
// 	log.Print('validateUUID:start')
// 	value := v.(string)
// 	if _,err = uuid.Parse(value); err != nil{
// 		errors = append(error, fmt Errorf("invalid UUID format"))
// 	}
// 	log.Print('validateUUID:end')
// }