wget -q -O /tmp/ci-scripts "https://github.com/bcaldwell/ci-scripts/releases/download/latest/linux_amd64"

REMOTESHA=$(wget -q -O - https://github.com/bcaldwell/ci-scripts/releases/download/latest/checksums.txt | grep linux_amd64 | awk  '{print $1}')
LOCALSHA=$(sha256sum /tmp/ci-scripts | awk '{print $1}')
if [[ "$REMOTESHA" != "$LOCALSHA" ]]; then
  echo "sha mismatch: ${REMOTESHA} -- ${LOCALSHA}"
  exit 1
fi

chmod +x /tmp/ci-scripts
echo "Moving binary to /usr/local/bin/"

SUDO=''
if command -v sudo; then
    SUDO='sudo'
fi
$SUDO

$SUDO mv /tmp/ci-scripts /usr/local/bin/
