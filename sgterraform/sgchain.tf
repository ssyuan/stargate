# Create  chain server
resource "alicloud_instance" "sgchain" {
  count = var.node_number
  image_id = "${data.alicloud_images.default.images.0.id}"
  internet_charge_type = "PayByBandwidth"
  spot_strategy = "SpotAsPriceGo"
  spot_price_limit = 0.45
  instance_type = "ecs.g6.large"
  system_disk_category = "cloud_efficiency"
  internet_max_bandwidth_out = 10
  password = var.ecs_password
  key_name = alicloud_key_pair.key_pair.id
  depends_on = [alicloud_instance.compile]
  security_groups = [
    "${alicloud_security_group.default.id}"]
  instance_name = "sgchain"
  vswitch_id = "${alicloud_vswitch.vsw.id}"
}

#upload file and  install package
resource "null_resource" "install" {
  count = var.node_number
  depends_on = [
    alicloud_instance.sgchain, alicloud_key_pair.key_pair ]
  ## trigger the instance inited
  triggers = {
    instance = alicloud_instance.sgchain.*.id[count.index]
  }
  ##copy ssh key to chain node
  provisioner "file" {
    source = var.ssh_key_pair_file
    destination = "/opt/"

    connection {
      host        = alicloud_instance.sgchain.*.public_ip[count.index]
      type        = "ssh"
      agent       = false
      user        = "root"
      private_key = file(var.ssh_key_pair_file)
    }
  }

  ##copy config to chain node
  provisioner "file" {
    source = "${var.config_file_path}/val/${count.index}.tar.gz"
    destination = "/opt/config.tar.gz"

    connection {
      host        = alicloud_instance.sgchain.*.public_ip[count.index]
      type        = "ssh"
      agent       = false
      user        = "root"
      private_key = file(var.ssh_key_pair_file)
    }
  }

  ##copy exec to chain node
  provisioner "file" {
    source      = var.exec_file_path
    destination = "/opt/sgchain"

    connection {
      host        = alicloud_instance.sgchain.*.public_ip[count.index]
      type        = "ssh"
      agent       = false
      user        = "root"
      private_key = file(var.ssh_key_pair_file)
    }
  }

  ##start sgchain
  provisioner "remote-exec" {
    connection {
      host = alicloud_instance.sgchain.*.public_ip[count.index]
      type = "ssh"
      agent = false
      user = "root"
      private_key = file(var.ssh_key_pair_file)
    }

    inline = [
      #command
      "mkdir -p /opt/starcoin",
      "cd /opt/starcoin",
      "scp -i ${var.ssh_key_pair_file}  ${alicloud_instance.compile.private_ip}:/starcoin/stargate/target/release/sgchain .",
      "chmod +x sgchain",
      "tar xzvf ../config.tar.gz",
      "./sgchain -f ${count.index}/node.config.toml",
    ]
  }
}