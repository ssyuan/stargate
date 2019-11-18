#compile server
resource "alicloud_instance" "compile" {
  image_id = "${data.alicloud_images.default.images.0.id}"
  internet_charge_type = "PayByBandwidth"
  spot_strategy = "SpotAsPriceGo"
  spot_price_limit = 0.45
  instance_type = "ecs.g6.large"
  system_disk_category = "cloud_efficiency"
  internet_max_bandwidth_out = 10
  password = var.ecs_password
  key_name = alicloud_key_pair.key_pair.id
  security_groups = [
    "${alicloud_security_group.default.id}"]
  instance_name = "sgchain"
  vswitch_id = "${alicloud_vswitch.vsw.id}"
  provisioner "local-exec" {
    command = "apt-get update -y && apt-get upgrade -y && apt-get install -y protobuf-compiler cmake curl clang git && curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y --default-toolchain none && source $HOME/.cargo/env &&mkdir /starcoin && cd /starcoin && git clone https://github.com/starcoinorg/stargate.git && cd stargate && git submodule init && git submodule update && cargo build --release -p sgchain && cd target/release && rm -r build deps incremental"
    #command = "ls /starcoin/stargate/target/release"
    connection {
      host = alicloud_instance.compile.public_ip
      type = "ssh"
      agent = false
      user = "root"
      private_key = file(var.ssh_key_pair_file)
    }
  }
}