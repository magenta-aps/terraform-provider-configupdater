#!/usr/bin/env -S nix-shell --run zsh
with (import <nixpkgs> {});
mkShell {
  buildInputs = [
    go
    terraform_1
  ];
}
