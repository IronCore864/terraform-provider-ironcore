# Terraform Provider Ironcore

Name of the provider: "ironcore".

This is an example of Terraform provider, providing a way to use terraform to manage local files.

## Build

```
go build
```

## Resources: `ironcore_file`

### Argument Reference:

- name: the name of the file

### Sample Usage

```
resource "ironcore_file" "test" {
  name = "test.txt"
}
```

## Deploy

```
terraform init
terraform apply
```
