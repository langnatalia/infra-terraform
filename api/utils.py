import os
import json
from typing import Dict, List

def load_config(file_path: str) -> Dict:
    with open(file_path, 'r') as file:
        return json.load(file)

def save_config(file_path: str, config: Dict) -> None:
    with open(file_path, 'w') as file:
        json.dump(config, file, indent=4)

def get_terraform_state_file(working_dir: str) -> str:
    return os.path.join(working_dir, 'terraform.tfstate')

def validate_terraform_config(config: Dict) -> bool:
    required_keys = ['provider', 'resource']
    for key in required_keys:
        if key not in config:
            return False
    return True

def get_resource_ids(config: Dict) -> List[str]:
    resource_ids = []
    for resource in config.get('resource', []):
        resource_ids.append(resource.get('id'))
    return resource_ids