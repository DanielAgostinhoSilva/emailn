#!/usr/bin/env bash
# Ativa os hooks versionados em .githooks/ para este clone do repositório.
# Rode uma vez após clonar:
#   bash .githooks/install.sh
set -euo pipefail

if ! git rev-parse --git-dir >/dev/null 2>&1; then
  echo "Erro: rode este script de dentro de um repositório git." >&2
  exit 1
fi

repo_root="$(git rev-parse --show-toplevel)"
chmod +x "$repo_root/.githooks/commit-msg"
git config core.hooksPath .githooks

echo "OK: git configurado para usar os hooks em .githooks/ (core.hooksPath)."
echo "A partir de agora, 'git commit' valida a mensagem contra o padrão Conventional Commits."