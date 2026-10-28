# FOREVER-GO

## About The Project

An exploratory full-stack e-commerce prototype for sustainable, lifetime-guaranteed clothing, from the original project description. It is marked as abandoned in that description.

It has a React and Material UI storefront with a Stripe checkout flow (address form, payment form and success page) and a small Go backend that currently answers "The API is up and running!". Both services run in Docker for local development.

### Built With

- React and Material UI
- Stripe (`@stripe/react-stripe-js`)
- Go
- Docker Compose

## Getting Started

### Prerequisites

- Docker and Docker Compose

### Running

```sh
make start
```

This runs `docker compose up -d --build`. The frontend is served on port 3000 and the backend on port 4500.

To run the frontend alone, from the original `frontend/README.md`:

```sh
cd frontend
npm install
npm start
```

Then open http://localhost:3000/product.

## Roadmap

The project is abandoned and not planned to be developed further.

## License

No license file is present in this repository.

## Contact

Dyrits - https://github.com/Dyrits
