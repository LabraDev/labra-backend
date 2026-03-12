{
	inputs = {
		nixpkgs.url = "github@nixos/nixpkgs";
	};

	outputs = { self, nixpkgs }: {
		packages.aarch64-darwin.hello = nixpkgs.legacyPackages.aarch64-darwin.hello;
	};
}
