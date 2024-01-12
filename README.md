# Bus Reservation App

Welcome to the Bus Reservation app, a comprehensive solution that caters to the needs of users, bus service providers, and app administrators. This application facilitates seamless bus reservations, providing a user-friendly experience for all stakeholders.

## Features

### For Users

- **User Registration and Authentication:**
  - Users can register and authenticate via email OTP validation to access the app.

- **Bus Search and Booking:**
  - Search for buses based on input routes.
  - Add new passengers.
  - Book seats.
  - View passenger details.

- **Booking Management:**
  - Users can cancel bookings.
  - Check seat availability.
  - Obtain details of bus stations.

- **Wallet System:**
  - A wallet system is implemented for users.
  - Provides a convenient payment option.

- **SMS Notifications:**
  - Receive SMS notifications on booking and cancellation events.

### For Bus Service Providers

- **Provider Registration and Verification:**
  - Bus service providers can register.
  - Upon verification by the app admin, they gain access to the app.

- **Station and Bus Management:**
  - Add new stations.
  - View available stations.
  - Add new buses.
  - Edit bus information.
  - Remove buses from the app.

- **Coupon Management:**
  - Providers can offer discounts through coupons.
  - Manage the coupons they provide.

- **Wallet System:**
  - Similar to users, bus service providers have a wallet system.

### For App Admin

- **Comprehensive Admin Rights:**
  - Admin has full control over users, service providers.
  - Can perform any necessary modifications.

- **Chart Management:**
  - Admin can add new charts for buses.
  - Has the authority to cancel a bus.

### Additional Features

- **Payment Options:**
  - Payment options are implemented between the wallet and Razorpay for efficient payment processing.

- **Enhanced Performance:**
  - The use of Go Routines and channels enhances overall application performance.

- **Scheduled Tasks:**
  - Cron is utilized for scheduled tasks, ensuring timely execution of specific functions.

- **Coupon Validation:**
  - Coupons are implemented with proper validation through the use of a cron job.

## Technology Stack

- **Backend Language:** Go
- **Web Framework:** Gin
- **Database:** PostgreSQL
- **Caching:** Redis
- **Scheduled Tasks:** Cron
- **Payment Processing:** Razorpay
- **SMS Notifications:** Twilio
- **Authentication:** JWT
- **Code Linting:** Golint
- **ORM:** GORM

## API Documentation

For detailed API documentation, refer to  [API Documentation](https://documenter.getpostman.com/view/30887078/2s9YsNcq8T).

# Setup and Installation

Follow these steps to set up and run the Bus Reservation app on your local machine.

## 1. Clone the repository:

```bash
git clone https://github.com/AswinManojan/Go_Bus.git
```
## 2. Install Dependencies:
```bash
	go get -u github.com/dgrijalva/jwt-go v3.2.0+incompatible
	go get -u github.com/gin-gonic/gin v1.9.1
	go get -u github.com/go-redis/redis/v8 v8.11.5
	go get -u github.com/joho/godotenv v1.5.1
	go get -u github.com/razorpay/razorpay-go v1.2.0
	go get -u github.com/robfig/cron v1.2.0
	go get -u github.com/twilio/twilio-go v1.15.2
```
## 3. Setup Databse:
CREATE DATABASE gobus;

## 4. Configure environment variables:

DB_CONFIG="host=##### user=##### password= dbname=gobus port=### sslmode=disable"

EMAIL="#######@gmail.com"

APP_PASSWORD="### ### ### ###"

TWILIO_ACCOUNT_SID=#########

TWILIO_AUTH_TOKEN=#########

RAZOR_KEY_ID=#########

RAZOR_SECRET=########

MY_NUMBER="888-888-8888"



### Feel free to reach out for any inquiries or issues. Happy coding!
