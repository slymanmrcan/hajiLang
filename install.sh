#!/bin/bash
set -e

# Renkler
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}HajiLang Kurulumu Başlatılıyor...${NC}"

# OS tespiti
OS="$(uname -s)"
ARCH="$(uname -m)"

case "$OS" in
    Linux*)     OS_TYPE=linux;;
    Darwin*)    OS_TYPE=darwin;;
    *)          
        echo -e "${RED}Desteklenmeyen işletim sistemi: $OS${NC}"
        exit 1
        ;;
esac

case "$ARCH" in
    x86_64)     ARCH_TYPE=amd64;;
    arm64)      ARCH_TYPE=arm64;;
    aarch64)    ARCH_TYPE=arm64;;
    *)          
        echo -e "${RED}Desteklenmeyen mimari: $ARCH${NC}"
        exit 1
        ;;
esac

echo -e "${YELLOW}Platform: $OS_TYPE-$ARCH_TYPE${NC}"

# GitHub'dan son sürümü al
REPO="slymanmrcan/hajilang"
RELEASE_URL="https://github.com/$REPO/releases/latest/download/hajilang-$OS_TYPE-$ARCH_TYPE"

# Binary'yi indir
echo -e "${YELLOW}İndiriliyor...${NC}"
TMP_FILE="/tmp/hajilang"

if command -v curl &> /dev/null; then
    curl -L "$RELEASE_URL" -o "$TMP_FILE"
elif command -v wget &> /dev/null; then
    wget -O "$TMP_FILE" "$RELEASE_URL"
else
    echo -e "${RED}curl veya wget bulunamadı!${NC}"
    exit 1
fi

# Çalıştırılabilir yap
chmod +x "$TMP_FILE"

# Kurulum yeri
if [ -w "/usr/local/bin" ]; then
    INSTALL_DIR="/usr/local/bin"
    sudo mv "$TMP_FILE" "$INSTALL_DIR/hajilang"
else
    INSTALL_DIR="$HOME/.local/bin"
    mkdir -p "$INSTALL_DIR"
    mv "$TMP_FILE" "$INSTALL_DIR/hajilang"
    
    # PATH kontrolü
    if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
        echo -e "${YELLOW}⚠️  $INSTALL_DIR PATH'te değil!${NC}"
        echo -e "${YELLOW}Şunu ~/.bashrc veya ~/.zshrc'ye ekle:${NC}"
        echo -e "${GREEN}export PATH=\"\$HOME/.local/bin:\$PATH\"${NC}"
    fi
fi

echo -e "${GREEN}✅ HajiLang başarıyla kuruldu!${NC}"
echo -e "${GREEN}Kullanım: hajilang [dosya.haji]${NC}"
echo -e "${GREEN}REPL için: hajilang${NC}"