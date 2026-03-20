import os
import re
import json

OLD_CODE_DIR = "../old-code"

def extract_gather_options(content):
    """Extracts say/press options from <Gather> blocks"""
    options = {}
    pattern = r'<Say.*?>(Press (\d) to (.*?))\.?</Say>'
    matches = re.findall(pattern, content, re.IGNORECASE)
    for match in matches:
        digit = match[1]
        action = match[2].strip()
        options[digit] = action
    return options

def extract_handler_logic(content):
    """Extracts routing logic from the ivrs-handler.php switch statements"""
    routes = {}
    # Find all if/elseif blocks checking for digits
    pattern = r'if.*?\(\$_REQUEST\[\'Digits\'\] == \'(\d)\'\).*?\{(.*?)\}(?:\s*elseif|\s*else)'
    matches = re.findall(pattern, content, re.IGNORECASE | re.DOTALL)
    
    for match in matches:
        digit = match[0]
        block = match[1]
        
        # Determine the action in the block
        action_type = "Unknown"
        target = None
        
        if "<Dial" in block:
            action_type = "Dial"
            dial_match = re.search(r'<Dial.*?>(.*?)</Dial>', block)
            if dial_match:
                target = dial_match.group(1).strip()
        elif "<Redirect" in block:
            action_type = "Redirect"
            redir_match = re.search(r'<Redirect.*?>(.*?)</Redirect>', block)
            if redir_match:
                target = redir_match.group(1).strip()
                
        routes[digit] = {
            "action_type": action_type,
            "target": target
        }
    return routes

def main():
    flow = {
        "Main_Menu": {
            "Prompt": "Welcome to Hospital at Hand",
            "Options": {}
        },
        "Routing": {}
    }

    try:
        # 1. Parse ivrs.php for the main menu options
        with open(os.path.join(OLD_CODE_DIR, "ivrs.php"), "r", encoding="utf-8") as f:
            ivrs_content = f.read()
            flow["Main_Menu"]["Options"] = extract_gather_options(ivrs_content)

        # 2. Parse ivrs-handler.php for the routing logic
        with open(os.path.join(OLD_CODE_DIR, "ivrs-handler.php"), "r", encoding="utf-8") as f:
            handler_content = f.read()
            flow["Routing"] = extract_handler_logic(handler_content)

        # 3. Save as JSON artifact
        with open("legacy_flow.json", "w", encoding="utf-8") as out:
            json.dump(flow, out, indent=4)
        
        print("Legacy Flow mapped successfully to legacy_flow.json")
    
    except Exception as e:
        print(f"Error reading PHP files: {e}")

if __name__ == "__main__":
    main()
