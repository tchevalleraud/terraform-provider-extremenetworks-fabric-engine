# terraform-provider-extremenetworks-fabric-engine

Terraform provider to manage Extreme Networks Fabric Engine switches (VSP) via SSH.

## Example Usage

```hcl
resource "extremenetworks-fabric-engine_hostname" "core" {
  hostname = "core-paris-1"
  address  = "10.10.10.1"
  user     = "admin"
  password = "yourpassword"
}
