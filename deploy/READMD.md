# consul
## vim /etc/consul.d/consul.hcl
systemctl status consul
/usr/bin/consul agent -config-dir=/etc/consul.d/

## /etc/consul.d/consul.hcl
autopilot = {}
disable_update_check = true
