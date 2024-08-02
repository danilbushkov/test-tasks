{ pkgs ? import <nixpkgs> {} }:
pkgs.mkShellNoCC {
    packages = with pkgs; [
        php83
        php83Packages.composer
        phpactor
        vscode-langservers-extracted
        docker
        postgresql
    ];
}
