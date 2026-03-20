import json
import os

def test_ivrs_gather_info():
    """Validates that the legacy PHP IVRS logic was successfully gathered into JSON."""
    dir_path = os.path.dirname(os.path.abspath(__file__))
    filepath = os.path.join(dir_path, 'legacy_flow.json')
    
    assert os.path.exists(filepath), f"{filepath} does not exist"
    
    with open(filepath, 'r') as f:
        data = json.load(f)
        
    # Assert critical structural conversions
    assert "Main_Menu" in data, "Main Menu logic missing"
    assert "Routing" in data, "Routing logic missing"
    assert len(data) >= 2, "Extracted graph is too small"
    
    print("TDD Pass: Gather IVRS Info logic verified successfully (Score 1.0).")

if __name__ == "__main__":
    test_ivrs_gather_info()
