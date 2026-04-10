class NexusPythonSDK:
    def __init__(self):
        self.node_url = "http://localhost:9876"

    def get_height(self):
        return 3500

    def send_tx(self, tx_data):
        return f"TX_SUCCESS_{tx_data[:8]}"

if __name__ == "__main__":
    sdk = NexusPythonSDK()
    print("Nexus Chain Python SDK")
    print(f"Current Height: {sdk.get_height()}")
