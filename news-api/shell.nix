{ pkgs ? import <nixpkgs> {} }:
pkgs.mkShellNoCC {
    packages = with pkgs; [
        postgresql
        gnumake
        yaml-language-server
        go
        gopls
    ];
}
