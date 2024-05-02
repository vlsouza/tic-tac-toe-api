### Tic-Tac-Toe API 🎮

An API for playing tic-tac-toe.

#### Endpoints:

1. **Create a New Match**
   - **Method:** POST
   - **URL:** `/matches`
   - **Request Body:** None
   - **Response:** No content
   - **Example Usage:** `curl -X POST http://localhost:8080/matches`

2. **Get Match State**
   - **Method:** GET
   - **URL:** `/matches/{id}`
   - **Path Parameter:** `id` (Match ID)
   - **Response:** JSON representation of the match state
   - **Example Usage:** `curl http://localhost:8080/matches/123`

3. **Get List of Matches by Status**
   - **Method:** GET
   - **URL:** `/matches?status={status}`
   - **Query Parameter:** `status` (Match status)
   - **Response:** JSON array of matches with the specified status
   - **Example Usage:** `curl http://localhost:8080/matches?status=active`

4. **Make a Move**
   - **Method:** POST
   - **URL:** `/matches/{id}/move`
   - **Path Parameter:** `id` (Match ID)
   - **Request Body:** JSON representation of the move request
   - **Response:** JSON representation of the updated match state
   - **Example Usage:** `curl -X POST -d '{"player": "X", "position": {"row": 0, "col": 1}}' http://localhost:8080/matches/123/move`

5. **Start a Match**
   - **Method:** POST
   - **URL:** `/matches/{id}/start`
   - **Path Parameter:** `id` (Match ID)
   - **Response:** No content
   - **Example Usage:** `curl -X POST http://localhost:8080/matches/123/start`

#### Response Status Codes:

- **200 OK:** Successful operation
- **400 Bad Request:** Invalid request or parameters
- **404 Not Found:** Resource not found
- **500 Internal Server Error:** Unexpected error occurred

#### Technology Stack:

- Go (Golang)
- AWS DynamoDB

#### How to Use:

1. Install Go and set up your environment.
2. Clone the repository.
3. Build and run the application.
4. Configure the EC2 instance for the API and create and DynamoDB table called tic_tac_toe_match. Both need to be in the same AWS region.
5. Use the provided endpoints to interact with the API.

Feel free to extend the functionality or contribute to the project! 🚀
