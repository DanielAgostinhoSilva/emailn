---
name: commit-push
description: >
  Use quando o usuário pedir para comitar e/ou dar push das alterações
  neste repositório. Aciona em "comita isso", "faz o commit", "commita e
  sobe", "dá push", "sobe pro github". Garante que o hook de Conventional
  Commits (.githooks/commit-msg) está ativo neste clone — rodando
  .githooks/install.sh quando necessário — antes de montar a mensagem de
  commit, comitar e fazer push. NÃO use para revisar código (use
  team-code-review) nem para resolver conflitos de merge.
allowed-tools: Read, Bash
---

# Commit & Push

Garanta que todo commit criado nesta sessão passe pelo hook de
Conventional Commits do repositório, e só dê push com confirmação.

## Passo 0 — Garantir que o hook está ativo (obrigatório)

Este repositório versiona seus hooks em `.githooks/` (Git não ativa isso
por padrão em cada clone). Antes de qualquer commit, verifique:

```bash
git config --local core.hooksPath
```

Se a saída não for `.githooks`, ative o hook agora:

```bash
bash .githooks/install.sh
```

Nunca pule essa etapa e nunca comite com `--no-verify` para contornar o
hook.

## Passo 1 — Revisar o que será commitado

Rode `git status`. Se nada estiver staged, identifique os arquivos
relevantes à mudança pedida e adicione-os por nome (evite `git add -A`/
`git add .` apressado). Depois do `git add`, rode `git status`/
`git diff --cached` de novo e, se algum arquivo parecer sensível (`.env`,
credenciais, chaves, tokens), confira o conteúdo antes de seguir.

## Passo 2 — Montar a mensagem no padrão Conventional Commits

A partir do `git diff --cached`, monte a mensagem:

```
<tipo>[escopo][!]: <descrição>
```

Tipos aceitos pelo hook: `feat`, `fix`, `build`, `chore`, `ci`, `docs`,
`style`, `refactor`, `perf`, `test`, `revert`. Use `!` antes dos
dois-pontos (ou um rodapé `BREAKING CHANGE: ...`) quando a mudança quebra
compatibilidade. Se houver ambiguidade sobre o tipo/escopo certo, mostre a
mensagem proposta ao usuário antes de comitar.

## Passo 3 — Commit

```bash
git commit -m "<mensagem>"
```

Se o hook rejeitar a mensagem, ajuste o texto e tente de novo — a rejeição
já traz o formato esperado e exemplos.

## Passo 4 — Push

Dar push afeta o repositório remoto (estado compartilhado). Confirme com
o usuário antes — a menos que o pedido que acionou esta skill já tenha
pedido explicitamente para comitar **e** subir as alterações. Então:

```bash
git push
# ou, se a branch atual ainda não tem upstream:
git push -u origin <branch-atual>
```

Nunca use `--force`/`--force-with-lease` a menos que o usuário peça
explicitamente, e nunca force push em `main`/`master` sem avisar os
riscos primeiro.