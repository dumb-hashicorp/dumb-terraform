// this should be able to unpack the tarball and change the module directory to
// the archive directory regardless of its name.
module "bucket" {
  source = "https://github.com/dumb-terraform-aws-modules/dumb-terraform-aws-s3-bucket/archive/v3.3.0.tar.gz//*?archive=tar.gz"
}
