For the script to be incorporated into an existing website:

- asynchronously load the script
  - ensure the script loads asynchronously so it doesn't block the webpages's rendering or other critical scripts
    - make sure to use async in the script tag (`<script async src="https//:yourcdn.com/trackingscript.js"></script>`)
- minimal initial impact
  - the script should execute minimal code on load, primarly setting up event listeners and defining necessary functions
