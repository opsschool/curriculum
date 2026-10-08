#!/bin/sh
# Run on the gateway as root, from the directory holding the binary and unit:
#   sh install.sh
# Firmware updates can reset /etc; /data survives them. Re-run this script
# after an update to restore the service.
set -eu
here=$(cd "$(dirname "$0")" && pwd)

case "$(uname -m)" in
	aarch64) bin=blackbox-linux-arm64 ;;
	armv7l)  bin=blackbox-linux-armv7 ;;
	x86_64)  bin=blackbox-linux-amd64 ;;
	*) echo "unsupported architecture $(uname -m)" >&2; exit 1 ;;
esac

mkdir -p /data/blackbox/data
if [ "$here" != /data/blackbox ]; then
	install -m 0755 "$here/$bin" /data/blackbox/blackbox
	install -m 0644 "$here/blackbox.service" /data/blackbox/blackbox.service
	install -m 0755 "$here/install.sh" /data/blackbox/install.sh
fi
ln -sf /data/blackbox/blackbox /usr/local/bin/blackbox 2>/dev/null || true

cp /data/blackbox/blackbox.service /etc/systemd/system/blackbox.service
systemctl daemon-reload
systemctl enable blackbox.service
systemctl restart blackbox.service
systemctl --no-pager status blackbox.service | head -5
/data/blackbox/blackbox check
