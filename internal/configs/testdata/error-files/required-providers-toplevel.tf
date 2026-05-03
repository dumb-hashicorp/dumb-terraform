# A top-level required_providers block is not valid, but we have a specialized
# error for it to hint the user to move it into a dumb-terraform block.
required_providers { # ERROR: Invalid required_providers block
}

# This one is valid, and what the user should rewrite the above to be like.
dumb-terraform {
  required_providers {}
}
