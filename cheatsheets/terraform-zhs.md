---
title: Terraform
icon: fa-cubes
primary: "#7B42BC"
lang: hcl
locale: zhs
---

## fa-plug Provider 与 Resource

```hcl
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = "us-east-1"
}

resource "aws_instance" "web" {
  ami           = "ami-0c55b159cbfafe1f0"
  instance_type = "t3.micro"

  tags = {
    Name = "web-server"
  }
}
```

## fa-sliders 变量

```hcl
variable "region" {
  type    = string
  default = "us-east-1"
}

variable "instance_type" {
  type        = string
  description = "EC2 实例类型"
  default     = "t3.micro"

  validation {
    condition     = contains(["t3.micro", "t3.small", "t3.medium"], var.instance_type)
    error_message = "必须是有效的 t3 实例类型。"
  }
}

variable "azs" {
  type    = list(string)
  default = ["us-east-1a", "us-east-1b"]
}

variable "tags" {
  type    = map(string)
  default = {}
}
```

## fa-arrow-right-from-bracket 输出

```hcl
output "instance_id" {
  value       = aws_instance.web.id
  description = "EC2 实例 ID"
}

output "public_ip" {
  value = aws_instance.web.public_ip
}

output "vpc_id" {
  value       = aws_vpc.main.id
  description = "VPC ID"
  sensitive   = false
}
```

## fa-database 数据源

```hcl
data "aws_ami" "ubuntu" {
  most_recent = true
  owners      = ["099720109477"]

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-*"]
  }
}

data "aws_vpc" "default" {
  default = true
}

data "aws_subnets" "all" {
  filter {
    name   = "vpc-id"
    values = [data.aws_vpc.default.id]
  }
}

resource "aws_instance" "web" {
  ami = data.aws_ami.ubuntu.id
}
```

## fa-boxes 模块

```hcl
module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "~> 5.0"

  name = "my-vpc"
  cidr = "10.0.0.0/16"

  azs             = ["us-east-1a", "us-east-1b"]
  public_subnets  = ["10.0.1.0/24", "10.0.2.0/24"]
  private_subnets = ["10.0.3.0/24", "10.0.4.0/24"]

  enable_nat_gateway = true
}

module "ec2" {
  source = "./modules/ec2"

  instance_type = "t3.micro"
  subnet_id     = module.vpc.public_subnets[0]
}
```

## fa-hard-drive 状态管理

```bash
terraform state list                        # 列出状态中所有资源
terraform state show aws_instance.web       # 查看资源属性
terraform state mv aws_instance.old aws_instance.new  # 在状态中重命名资源
terraform state rm aws_instance.web         # 从状态中移除资源（不删除基础设施）
terraform state replace-provider hashicorp/aws hashicorp/aws  # 替换 Provider
terraform refresh                           # 同步状态与实际基础设施
terraform import aws_instance.web i-123456  # 将已有资源导入状态
```

## fa-terminal CLI 命令

```bash
terraform init                  # 初始化后端并安装 Provider
terraform init -upgrade         # 升级 Provider 到最新允许版本
terraform plan                  # 预览变更
terraform plan -out=tfplan      # 保存计划到文件
terraform plan -var="region=us-west-2"  # 传递变量
terraform apply                 # 应用变更（需交互确认）
terraform apply -auto-approve   # 自动确认应用变更
terraform apply tfplan          # 应用已保存的计划
terraform destroy               # 销毁所有托管基础设施
terraform fmt                   # 格式化所有 .tf 文件
terraform validate              # 语法和一致性检查
```

## fa-cloud Terraform Cloud

```hcl
terraform {
  cloud {
    organization = "my-org"
    workspaces {
      name = "my-workspace"
    }
  }
}

terraform {
  backend "remote" {
    organization = "my-org"
    workspaces {
      name = "production"
    }
  }
}
```

## fa-code-branch 条件与循环

```hcl
resource "aws_instance" "web" {
  count = var.enable_web ? 1 : 0

  ami           = data.aws_ami.ubuntu.id
  instance_type = var.instance_type
}

resource "aws_eip" "eip" {
  for_each = toset(var.azs)
  domain   = "vpc"
  tags = {
    AZ = each.key
  }
}

locals {
  security_group_rules = [
    { port = 80, cidr = "0.0.0.0/0" },
    { port = 443, cidr = "0.0.0.0/0" },
    { port = 22, cidr = "10.0.0.0/16" },
  ]
}

dynamic "ingress" {
  for_each = local.security_group_rules
  content {
    from_port   = ingress.value.port
    to_port     = ingress.value.port
    protocol    = "tcp"
    cidr_blocks = [ingress.value.cidr]
  }
}
```

## fa-wrench Provisioner

```hcl
resource "aws_instance" "web" {
  ami           = data.aws_ami.ubuntu.id
  instance_type = "t3.micro"

  provisioner "remote-exec" {
    inline = [
      "sudo apt-get update",
      "sudo apt-get install -y nginx",
    ]

    connection {
      type        = "ssh"
      user        = "ubuntu"
      private_key = file("~/.ssh/id_rsa")
      host        = self.public_ip
    }
  }

  provisioner "local-exec" {
    command = "echo ${self.public_ip} > inventory.txt"
  }
}
```

## fa-file-import 导入

```bash
terraform import aws_instance.web i-0abcd1234efgh5678
terraform import aws_s3_bucket.mybucket my-bucket-name
terraform import aws_subnet.private subnet-abc123
terraform import 'module.vpc.aws_subnet.public[0]' subnet-xyz789
terraform import -var="region=us-west-2" aws_instance.web i-123
```

## fa-layer-group Workspace

```bash
terraform workspace new staging       # 创建工作空间
terraform workspace new production    # 创建另一个工作空间
terraform workspace list              # 列出所有工作空间
terraform workspace select staging    # 切换工作空间
terraform workspace show              # 显示当前工作空间
terraform workspace delete staging    # 删除工作空间
```

```hcl
resource "aws_instance" "web" {
  ami           = data.aws_ami.ubuntu.id
  instance_type = terraform.workspace == "production" ? "t3.medium" : "t3.micro"

  tags = {
    Environment = terraform.workspace
  }
}
```

## fa-server 远程后端

```hcl
terraform {
  backend "s3" {
    bucket         = "my-terraform-state"
    key            = "infra/terraform.tfstate"
    region         = "us-east-1"
    encrypt        = true
    dynamodb_table = "terraform-locks"
  }
}

terraform {
  backend "gcs" {
    bucket = "my-terraform-state"
    prefix = "infra"
  }
}

terraform {
  backend "consul" {
    path = "terraform/infra"
  }
}
```

## fa-function 函数

```hcl
locals {
  str_upper  = upper("hello")
  str_lower  = lower("HELLO")
  joined     = join(", ", ["a", "b", "c"])
  splitted   = split(",", "a,b,c")
  trimmed    = trimspace("  hello  ")
  replaced   = replace("hello world", "world", "terraform")

  cidr_subnet = cidrsubnet("10.0.0.0/16", 8, 1)
  lookup_val  = lookup({ a = "1", b = "2" }, "a", "default")
  merge_map   = merge({ a = 1 }, { b = 2 })

  file_content = file("config.yaml")
  template     = templatefile("user-data.sh", { name = "web" })

  json_enc  = jsonencode({ name = "web", port = 80 })
  yaml_enc  = yamlencode({ name = "web", port = 80 })

  sorted = sort(["c", "a", "b"])
  unique = toset(["a", "a", "b"])

  timestamp_val = timestamp()
  formatted     = format("web-%03d", 1)
}
```
