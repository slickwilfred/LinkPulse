dotenv.config()

/**
 * 
 */
function sendData(data, path) {
    fetch(process.env.API_URL/path, {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify(data),
        keepalive: true // Important for send-operations in unload events
    }).catch(error => {
        console.error('Failed to send data', error);
    })
}

export default sendData;