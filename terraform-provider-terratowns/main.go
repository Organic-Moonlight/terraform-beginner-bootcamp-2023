package main

import (
    "context" 
    "fmt"
    "log"
    "github.com/google/uuid"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag" 
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: Provider,
    })
}

func Provider() *schema.Provider {
    return &schema.Provider{
        ResourcesMap: map[string]*schema.Resource{
            "terratowns_home": Resource(),
        },
        DataSourcesMap: map[string]*schema.Resource{},
        Schema: map[string]*schema.Schema{
            "endpoint": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "The endpoint for the external service",
            },
            "token": {
                Type:        schema.TypeString,
                Sensitive:   true,
                Required:    true,
                Description: "Bearer token for authorization",
            },
            "user_uuid": {
                Type:         schema.TypeString,
                Required:     true,
                Description:  "UUID for configuration",
                ValidateFunc: validateUUID,
            },
        },
        ConfigureContextFunc: providerConfigure,
    }
}

func validateUUID(v interface{}, k string) (ws []string, errors []error) {
    log.Println("validateUUID: start")
    value, ok := v.(string)
    if !ok {
        errors = append(errors, fmt.Errorf("expected type of %s to be string", k))
        return
    }
    if _, err := uuid.Parse(value); err != nil {
        errors = append(errors, fmt.Errorf("invalid UUID format for %s: %w", k, err))
    }
    log.Println("validateUUID: end")
    return
}

type Config struct {
    Endpoint string
    Token    string
    UserUUID string
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
    log.Print("providerConfigure: start")
    config := Config{
        Endpoint: d.Get("endpoint").(string),
        Token:    d.Get("token").(string),
        UserUUID: d.Get("user_uuid").(string),
    }
    log.Print("providerConfigure: end")
    return &config, nil
}

func Resource() *schema.Resource {
    log.Println("Resource: start")
    resource := &schema.Resource{
        CreateContext: resourceHouseCreate,
        ReadContext:   resourceHouseRead,
        UpdateContext: resourceHouseUpdate,
        DeleteContext: resourceHouseDelete,
        Schema: map[string]*schema.Schema{
            // Add your resource schema here
        },
    }
    log.Println("Resource: end")
    return resource
}

func resourceHouseCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
    var diags diag.Diagnostics
    return diags
}

func resourceHouseRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
    var diags diag.Diagnostics
    return diags
}

func resourceHouseUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
    var diags diag.Diagnostics
    return diags
}

func resourceHouseDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
    var diags diag.Diagnostics
    return diags
}