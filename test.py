import requests
import jwt

class API:
    def __init__(self, base_url):
        self.base_url = base_url

    def register(self, email, password):
        response = requests.post(
            f"{self.base_url}/auth/register",
            json={'email': email, 'password': password}
        )
        return response.json()

    def login(self, email, password):
        response = requests.post(
            f"{self.base_url}/auth/login",
            json={'email': email, 'password': password}
        )
        return response.json()

    def refresh(self, refresh_token):
        response = requests.post(
            f"{self.base_url}/auth/refresh",
            cookies={"refresh_token": refresh_token}
        )
        return response.json()

    def get_user(self, access_token):
        response = requests.get(
            f"{self.base_url}/api/me",
            headers={"Authorization": f"Bearer {access_token}"}
        )
        return response.json()

def main():
    api = API("http://localhost:8080")
    
    # Register
    print("Registering...")
    register_resp = api.register("example@example.com", "example_password")
    print(register_resp)

    # Login
    print("Logging in...")
    login_resp = api.login("example@example.com", "example_password")
    print(login_resp)
    
    # Get user info
    print("Getting user info...")
    user_resp = api.get_user(login_resp['access_token'])
    print(user_resp)
    
    # Refresh token
    print("Refreshing token...")
    refresh_resp = api.refresh(login_resp['refresh_token'])
    print(refresh_resp)

if __name__ == "__main__":
    main()
