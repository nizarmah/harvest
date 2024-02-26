# buymeacoffee
BuyMeACoffee webhook event handlers.

## Testing
Instructions can help you test this.

### Setup

1. Expose the web app to the internet. We'll use `ngrok`.

    ```bash
    ngrok http whatisbean.local:80 --host-header=rewrite
    ```

2. Create a new webhook at https://www.buymeacoffee.com/webhooks

3. Set the URL to `<ngrok-url>/webhooks/buymeacoffee`

4. Set the events to: `membership.started`, `membership.cancelled`

5. Add the secret into `.env` under `BMC_WEBHOOK_SECRET=`

6. Start bean

### Test

1. Select your webhook at https://www.buymeacoffee.com/webhooks

2. Send a test event with the "Send test" button

3. Ensure the request is logged in `ngrok`

4. Ensure the response is displayed on BuyMeACoffee
