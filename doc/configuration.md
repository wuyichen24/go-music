## React Proxy
- **Description**: React server needs to know where is the backend server.
- **Location**: react-app/[package.json](../react-app/package.json)
  ```json
  "proxy": "http://127.0.0.1:8080/"
  ```

## Database Connection
- **Description**: Gin server needs to use it to connect to MySQL database.
- **Location**: gin-app/rest/[handler.go](../gin-app/rest/handler.go)
  ```go
  func NewHandler() (HandlerInterface, error) {
      db, err := db.NewORM("mysql", "root:6ytow2-;S3lA@/gomusic")    // Pattern: <username>:<password>@/<schema> 
      // omit other lines
  }
  ```
  
## Stripe API Key
- **Description**: Stripe requires an API key for online credit card charging.
- **Location**: react-app/src/[CreditCards.js](../react-app/src/CreditCards.js)
  ```js
  export default function CreditCardInformation(props) {
    // omit other lines
    return (
      <div>
        <StripeProvider apiKey="pk_test_TYooMQauvdEDq54NiTphI7jx">
        </StripeProvider>
      </div>
    );
  }
  ```
