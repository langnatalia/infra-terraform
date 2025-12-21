import os
import argparse
from terraformimporter import TerraformImporter
from infrastructure import Infrastructure

def main():
    parser = argparse.ArgumentParser(description='Infrastructure deployment tool')
    parser.add_argument('--config', help='Path to configuration file', required=True)
    args = parser.parse_args()

    if not os.path.exists(args.config):
        raise FileNotFoundError(f"Configuration file {args.config} not found")

    config = TerraformImporter.load(args.config)
    infra = Infrastructure(config)

    infra.apply()

if __name__ == "__main__":
    main()