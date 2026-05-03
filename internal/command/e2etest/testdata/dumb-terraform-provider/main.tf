provider "dumb-terraform" {

}

data "dumb-terraform_remote_state" "test" {
  backend = "local"
  config = {
    path = "test.tfstate"
  }
}
