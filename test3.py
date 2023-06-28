import requests
import json
import random
import string

# Replace with your server url
url = 'http://localhost:8080'

# Define a function to generate random string of a certain length
def random_string(length):
    letters = string.ascii_lowercase
    return ''.join(random.choice(letters) for i in range(length))

# Generate a large list of users
users = [
    {"email": f"user{i}@test.com", "password": random_string(8)}
    for i in range(10000)  # 10000 valid users
]

# Add some faulty users
faulty_users = [
    {"email": f"user{i}@test", "password": random_string(8)}  # Invalid email format
    for i in range(10000, 11000)
] + [
    {"email": "", "password": random_string(8)}  # Missing email
    for i in range(11000, 12000)
] + [
    {"email": f"user{i}@test.com", "password": ""}  # Missing password
    for i in range(12000, 13000)
] + [
    {"email": f"user{i}@test.com"}  # Missing password field
    for i in range(13000, 14000)
] + [
    {"password": random_string(8)}  # Missing email field
    for i in range(14000, 15000)
]

users.extend(faulty_users)  # Add faulty users to the main users list

# Register and Login users
for user in users:
    # Register user
    register_endpoint = '/auth/register'
    register_response = requests.post(url + register_endpoint, data=json.dumps(user), headers={'Content-Type': 'application/json'})
    print(f"Register status code for {user.get('email', 'MISSING_EMAIL')}: {register_response.status_code}, Response: {register_response.json()}")

    # Only try to log in if registration was successful
    if register_response.status_code == 201:
        # Correct password login
        login_endpoint = '/auth/login'
        login_response = requests.post(url + login_endpoint, data=json.dumps(user), headers={'Content-Type': 'application/json'})
        print(f"Login status code for {user.get('email', 'MISSING_EMAIL')}: {login_response.status_code}, Response: {login_response.json()}")

        if login_response.status_code == 200:
            access_token = login_response.json().get('access_token')
            refresh_token = login_response.json().get('refresh_token')

            # Call the /auth/refresh endpoint
            refresh_endpoint = '/auth/refresh'
            refresh_response = requests.post(url + refresh_endpoint, headers={'Authorization': f"Bearer {refresh_token}"})
            print(f"Refresh token status code for {user.get('email', 'MISSING_EMAIL')}: {refresh_response.status_code}, Response: {refresh_response.json()}")

            # Call the /api/me endpoint
            me_endpoint = '/api/me'
            me_response = requests.get(url + me_endpoint, headers={'Authorization': f"Bearer {access_token}"})
            print(f"/api/me status code for {user.get('email', 'MISSING_EMAIL')}: {me_response.status_code}, Response: {me_response.json()}")

        # Incorrect password login
        user_wrong_password = {"email": user.get('email', 'MISSING_EMAIL'), "password": random_string(8)}
        login_response_wrong_password = requests.post(url + login_endpoint, data=json.dumps(user_wrong_password), headers={'Content-Type': 'application/json'})
        print(f"Login with wrong password status code for {user.get('email', 'MISSING_EMAIL')}: {login_response_wrong_password.status_code}, Response: {login_response_wrong_password.json()}")

    # Try to register again with the same email
    register_response_again = requests.post(url + register_endpoint, data=json.dumps(user), headers={'Content-Type': 'application/json'})
    print(f"Re-register status code for {user.get('email', 'MISSING_EMAIL')}: {register_response_again.status_code}, Response: {register_response_again.json()}")

# Call the /env endpoint
env_endpoint = '/env'
env_response = requests.get(url + env_endpoint)
print(f"/env endpoint status code: {env_response.status_code}, Response: {env_response.json()}")
