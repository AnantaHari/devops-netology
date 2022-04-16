import sentry_sdk
sentry_sdk.init(
    "https://ec64b0f26df74be9813a456d530c119d@o1202968.ingest.sentry.io/6328539",

    # Set traces_sample_rate to 1.0 to capture 100%
    # of transactions for performance monitoring.
    # We recommend adjusting this value in production.
    traces_sample_rate=1.0,
    environment="development",
    release="myapp@1.1.0",
    debug=False
    
)

if __name__ == "__main__":
    division_by_zero = 1 / 0
