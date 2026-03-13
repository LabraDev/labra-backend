{
	description = "The backend flake of Labra";

	inputs = { 
		nixpkgs.url = "github:nixos/nixpkgs";
		flake-utils.url = "github:numtide/flake-utils";
	};

	outputs = { self, nixpkgs, flake-utils }:
		flake-utils.lib.eachDefaultSystem (system:
			let pkgs = nixpkgs.legacyPackages.${system};
			in {
				devShell = pkgs.mkShell { 
					buildInputs = [ pkgs.postgresql pkgs.go ]; 

					shellHook = ''
						export PGDATA=$(pwd)/pg_data
						echo "PGDATA set to $PGDATA"
						mkdir -p $PGDATA

						if [ ! -f $PGDATA/postgresql.conf ]; then
							initdb -D $PGDATA --no-locale --encoding UTF8
						fi

						echo "Remember to use pg_ctl -D $PGDATA start and pg_ctl -D $PGDATA stop to control the db"
					'';

				};
			});
}
