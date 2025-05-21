# Currency Converter CLI Tool

A simple command-line currency converter that uses the ExchangeRate-API to convert between currencies.

---
Author
Vaibhav Gupta — github.com/GVaibhav92


## Features

- Real-time currency conversion using live exchange rates  
- Supports all major world currencies  
- Simple command-line interface  
- Clear output with emoji visualization  

---

## Notes

The program uses a free-tier API key with limited requests.
For production use, consider:

-> Using environment variables for the API key
-> Implementing rate limiting
-> Adding error handling for API quota limits

## Supported Currencies

All major currencies are supported. Common examples include:

USD (US Dollar)
EUR (Euro)
GBP (British Pound)
JPY (Japanese Yen)
INR (Indian Rupee)
CAD (Canadian Dollar)
AUD (Australian Dollar)

For a full list of supported currencies, refer to the ExchangeRate-API documentation.

## Limitations

-> Free API tier has limited requests per month
-> No historical data support
-> No batch conversion capability

## License
This project is open-source and free to use. The ExchangeRate-API may have its own usage terms.
