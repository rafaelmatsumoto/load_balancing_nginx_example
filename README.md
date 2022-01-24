# Load Balancing with NGINX

This project is an example of how to load balance traffic (layer 7) with NGINX.

## Requirements

- [Docker Compose](https://docs.docker.com/compose/).

## Usage

```bash
docker-compose up
```

Then go to: http://localhost, then you should see `APPID = 8000`, if you keep reloading the page this `APPID` will range between 8000, 8001 and 8002

You can check other load balancing implementations by accessing the routes: http://localhost/app1 and http://localhost:app2

## Contributing
Pull requests and suggestions are welcome 😁.

## License
[MIT](https://choosealicense.com/licenses/mit/)

## Author

- [@rafaelmatsumoto](https://github.com/rafaelmatsumoto)
